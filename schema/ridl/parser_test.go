package ridl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newStringParser(src string) (*parser, error) {
	return newParser([]byte(src))
}

func TestParserTopLevelDefinitions(t *testing.T) {
	type expectation struct {
		left  string
		right string
	}

	{
		p, err := newStringParser("\n\n\n\n# a comment\n")
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)
		assert.Zero(t, len(p.root.Children()))
	}

	{
		p, err := newStringParser("\n\n\n\nwebrpc =   v1 # a comment\n")
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)
		assert.NotZero(t, len(p.root.Children()))

		assert.Equal(t, "webrpc", p.root.Definitions()[0].Left().String())
		assert.Equal(t, "v1", p.root.Definitions()[0].Right().String())
	}

	{
		p, err := newStringParser("\n\n\n\nwebrpc=v1 #a comment\n")
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.Equal(t, "webrpc", p.root.Definitions()[0].Left().String())
		assert.Equal(t, "v1", p.root.Definitions()[0].Right().String())
	}

	{
		p, err := newStringParser("\n\n\n\nwebrpc\t=\tv1. # a comment\n")
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.Equal(t, "webrpc", p.root.Definitions()[0].Left().String())
		assert.Equal(t, "v1.", p.root.Definitions()[0].Right().String())
	}

	{
		p, err := newStringParser("\n\n\n\nwebrpc = v1 .# a comment\nname=FOO")
		assert.NoError(t, err)

		err = p.run()
		assert.Error(t, err, "should have failed, we have garbage after webrpc definition")

		assert.Equal(t, 1, len(p.root.Children()))
	}

	{
		p, err := newStringParser("\n\n\n\nwebrpc = v1 # a comment\t\t\n\n\tname   = \tEXAMPLE_SERVICE   \n\nversion = v1.2.3")
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		expectations := []expectation{
			{"webrpc", "v1"},
			{"name", "EXAMPLE_SERVICE"},
			{"version", "v1.2.3"},
		}

		for i, e := range expectations {
			assert.Equal(t, e.left, p.root.Definitions()[i].Left().String())
			assert.Equal(t, e.right, p.root.Definitions()[i].Right().String())
		}
	}

	{
		p, err := newStringParser("\n\n\n\nwebrpc = v1 # a comment\n\tname =foo-bar#")
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		expectations := []expectation{
			{"webrpc", "v1"},
			{"name", "foo-bar#"},
		}

		for i, e := range expectations {
			assert.Equal(t, e.left, p.root.Definitions()[i].Left().String())
			assert.Equal(t, e.right, p.root.Definitions()[i].Right().String())
		}
	}

}

func TestParserError(t *testing.T) {
	p, err := newStringParser(`
		error 12345 InvalidUsername "username is invalid" HTTP 401
		error 12345 InvalidUsername "username is invalid" HTTP 1 # comment
		error 12345 InvalidUsername InvalidUsername
		error 12345 InvalidUsername InvalidUsername # comment
		error 45678 Unauthorized "unauthorized access" HTTP 401
		error 45678 Unauthorized "unauthorized access" HTTP 401 # comment
		error 45678 Unauthorized Unauthorized HTTP 401
		error 45678 Unauthorized Unauthorized HTTP 401 # comment
	`)
	assert.NoError(t, err)

	err = p.run()
	assert.NoError(t, err)

	if !assert.Equal(t, 8, len(p.root.Errors())) {
		for _, e := range p.root.Errors() {
			t.Logf("%v", e.message)
		}
	}
}

func TestParserErrorInvalid(t *testing.T) {
	tt := []string{
		`error`,
		`error WRONG`,
		`error 12345`,
		`error 12345 Unauthorized`, // missing <message>
		`error 12345 Unauthorized unauthorized access`, // missing quotes for multi-word <message>
		`error 12345 Unauthorized "unauthorized access" WRONG 401`,
		`error 12345 Unauthorized "unauthorized access" HTTP STATUS`,
		`error 12345 Unauthorized "unauthorized access" HTTP 401 EXTRA`,
	}

	for _, str := range tt {
		p, err := newStringParser(str)
		assert.NoError(t, err)

		err = p.run()
		assert.Error(t, err)

		if !assert.Equal(t, 0, len(p.root.Errors())) {
			for _, e := range p.root.Errors() {
				t.Logf("%v", e.message)
			}
		}
	}
}

func TestParserImport(t *testing.T) {
	{
		p, err := newStringParser(`import`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.Zero(t, len(p.root.Imports()))
	}

	{
		p, err := newStringParser(`
			# comment
			import #comment
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.Zero(t, len(p.root.Imports()))
	}

	{
		p, err := newStringParser(`
			import #comment
			# comment
			# comment
			# comment
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.Zero(t, len(p.root.Imports()))
	}

	{
		p, err := newStringParser(`
			import packageName.ridl
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.NotZero(t, len(p.root.Children()))
		if assert.Equal(t, 1, len(p.root.Imports())) {
			assert.Equal(t, "packageName.ridl", p.root.Imports()[0].Path().String())
		}
	}

	{
		p, err := newStringParser(`
			import # comment
			-foo
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		if assert.Equal(t, 1, len(p.root.Imports())) {
			assert.Equal(t, "foo", p.root.Imports()[0].Path().String())
		}
	}

	{
		p, err := newStringParser(`
			import # comment
			- foo # comment

			# comment
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		if assert.Equal(t, 1, len(p.root.Imports())) {
			assert.Equal(t, "foo", p.root.Imports()[0].Path().String())
		}
	}

	{
		p, err := newStringParser(`
			import # comment
		# comment
			#comment
			- foo
			                      # comment with spaces
			  - bar
			- baz                 # # # comment
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		imports := p.root.Imports()
		if assert.Equal(t, 3, len(imports)) {
			assert.Equal(t, "foo", imports[0].Path().String())
			assert.Equal(t, "bar", imports[1].Path().String())
			assert.Equal(t, "baz", imports[2].Path().String())
		}
	}

	{
		p, err := newStringParser(`
			import # comment
			# comment
				-        foo.ridl # comment
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.Equal(t, 1, len(p.root.Children()))
		if assert.Equal(t, 1, len(p.root.Imports())) {
			assert.Equal(t, "foo.ridl", p.root.Imports()[0].Path().String())
		}
	}

	{
		p, err := newStringParser(`
			import # comment
			# comment
				- # comment
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.Error(t, err)

		assert.Zero(t, len(p.root.Children()))
		assert.Zero(t, len(p.root.Imports()))
	}

	{
		p, err := newStringParser(`
			import # comment
			# comment
				- x # # comment
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.NotZero(t, len(p.root.Children()))
		imports := p.root.Imports()
		if assert.Equal(t, 1, len(imports)) {
			assert.Equal(t, "x", imports[0].Path().String())
		}
	}

	{
		p, err := newStringParser(`
			# comment
			 # comment
			import #comment
				-        foo.ridl
					- ./path/to/bar.ridl
					# comment
				- baz_- #comment
				- ../parent.ridl
				# comment
	`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		imports := p.root.Imports()
		if assert.Equal(t, 4, len(imports)) {
			assert.Equal(t, "foo.ridl", imports[0].Path().String())
			assert.Equal(t, "./path/to/bar.ridl", imports[1].Path().String())
			assert.Equal(t, "baz_-", imports[2].Path().String())
			assert.Equal(t, "../parent.ridl", imports[3].Path().String())
		}
	}

	{
		p, err := newStringParser(`
			import
				-        foo.ridl
				- t
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)
	}

	{
		p, err := newStringParser(`
			import /path/to/file.ridl
				- Member
				- Name # hello

			import /path/to/file2.ridl # hello
			# comment
				- Member
				- Name # hello
				#comment
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		imports := p.root.Imports()
		if assert.Equal(t, 2, len(imports)) {
			assert.Equal(t, "/path/to/file.ridl", imports[0].Path().String())
			assert.Equal(t, "Member", imports[0].Members()[0].String())
			assert.Equal(t, "Name", imports[0].Members()[1].String())

			assert.Equal(t, "/path/to/file2.ridl", imports[1].Path().String())
			assert.Equal(t, "Member", imports[1].Members()[0].String())
			assert.Equal(t, "Name", imports[1].Members()[1].String())
		}
	}

	{
		p, err := newStringParser(`
			import
				-        "/path /to foo.ridl"
				-
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.Error(t, err, "expecting import")

		imports := p.root.Imports()
		if assert.Equal(t, 1, len(imports)) {
			assert.Equal(t, "/path /to foo.ridl", imports[0].Path().String())
		}
	}

	{
		p, err := newStringParser(`
			import
				-        "/path /to foo.ridl"
				# coment
				- skywalker.ridl
				- "  " #comment
				- " \nABC \" DEF "
				# comment
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		imports := p.root.Imports()
		if assert.Equal(t, 4, len(imports)) {
			assert.Equal(t, "/path /to foo.ridl", imports[0].Path().String())
			assert.Equal(t, "skywalker.ridl", imports[1].Path().String())
			assert.Equal(t, "  ", imports[2].Path().String())
			assert.Equal(t, " \nABC \" DEF ", imports[3].Path().String())
		}
	}

	{
		p, err := newStringParser(`
			# api.ridl
			webrpc = v1

			import "./users.ridl" #comment
			import "./users # .ridl" #comment

			import #comment
				- uno
				- dos

			service API
				 - FindUser(id: uint64) => (user: User)
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.Equal(t, "webrpc", p.root.Definitions()[0].Left().String())
		assert.Equal(t, "v1", p.root.Definitions()[0].Right().String())

		imports := p.root.Imports()
		if assert.Equal(t, 4, len(imports)) {
			assert.Equal(t, "./users.ridl", imports[0].Path().String())
			assert.Equal(t, "./users # .ridl", imports[1].Path().String())
			assert.Equal(t, "uno", imports[2].Path().String())
			assert.Equal(t, "dos", imports[3].Path().String())
		}
	}

	{
		p, err := newStringParser(`
			import  "../contacts/proto/contacts.ridl"
				# - Counter # only Counter is imported from the contacts.ridl schema, or add more with more -'s
				# - Message2
				# - Service1
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.Equal(t, "../contacts/proto/contacts.ridl", p.root.Imports()[0].Path().String())
	}
}

func TestParserEnum(t *testing.T) {
	{
		p, err := newStringParser(`
			enum
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.Error(t, err, "expecting enum name")
	}

	{
		p, err := newStringParser(`
			enum Foo
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.Error(t, err, "expecting enum type")
	}

	{
		p, err := newStringParser(`
			enum Foo:
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.Error(t, err, "expecting enum type")
	}

	{
		p, err := newStringParser(`
			enum Foo: uint
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)
	}

	{
		p, err := newStringParser(`
			enum Foo: uint # comment
			# comment
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)
	}

	{
		p, err := newStringParser(`
			import
				- "hello"
			# comment


# comment

			enum Foo: int # comment
			# comment
				- Value # comment
 #ccomment
				- value2 # comment

#comment
				# comment
			## cmment


			enum Bar: int
				-  USER
				- ADMIN
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		imports := p.root.Imports()
		if assert.Equal(t, 1, len(imports)) {
			assert.Equal(t, "hello", imports[0].Path().String())
		}

		enums := p.root.Enums()
		if assert.Equal(t, 2, len(enums)) {
			assert.Equal(t, "Value", enums[0].values[0].Left().String())
			assert.Equal(t, "value2", enums[0].values[1].Left().String())

			assert.Equal(t, "USER", enums[1].values[0].Left().String())
			assert.Equal(t, "ADMIN", enums[1].values[1].Left().String())
		}
	}

	{
		p, err := newStringParser(`
			import
			# comment
				- "hello"

# comment

			enum Foo: string # comment
			# comment
				- Value = "SKY" # comment

				- value2 = WALKER # comment


				# comment
			## cmment


			enum Bar: int
				- USER
				- ADMIN

				# this is a comment
				# yep
				enum Kind: uint32
					- USER = 1             # comment
					- ADMIN = 2            # comment..

				# or.. just..
				enum Kind2: uint32
					- USER                 # aka, = 0
					- ADMIN                # aka, = 1
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		imports := p.root.Imports()
		if assert.Equal(t, 1, len(imports)) {
			assert.Equal(t, "hello", imports[0].Path().String())
		}

		enums := p.root.Enums()
		if assert.Equal(t, 4, len(enums)) {
			assert.Equal(t, "Value", enums[0].values[0].Left().String())
			assert.Equal(t, "SKY", enums[0].values[0].Right().String())

			assert.Equal(t, "value2", enums[0].values[1].Left().String())
			assert.Equal(t, "WALKER", enums[0].values[1].Right().String())

			assert.Equal(t, "USER", enums[1].values[0].Left().String())
			assert.Equal(t, "ADMIN", enums[1].values[1].Left().String())

			assert.Equal(t, "USER", enums[3].values[0].Left().String())
			assert.Equal(t, "ADMIN", enums[3].values[1].Left().String())
		}
	}
}

func TestParserStruct(t *testing.T) {
	{
		p, err := newStringParser(`
			struct
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.Error(t, err, "expecting struct name")
	}

	{
		p, err := newStringParser(`
			struct Empty
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)
	}

	{
		p, err := newStringParser(`
		struct Role
			- name: string
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		structs := p.root.Structs()
		if assert.Equal(t, 1, len(structs)) {
			assert.Equal(t, "name", structs[0].Fields()[0].Left().String())
			assert.Equal(t, "string", structs[0].Fields()[0].Right().String())
		}
	}

	{
		p, err := newStringParser(`
		struct ComplexTypes
			- arrayOfStrings: []string
			- arrayOfArrayOfStrings: [][]string
			- arrayOfArrayOfArrayOfStrings: [][][]string
			# KURSOR
			- stringMapOfStrings: map<string,string> # COMMENT
			- stringMapOfUint64: map<string,uint64> # COMMENT
			- stringMapOfArrayOfUint64: map<string,[]uint64>
			- stringMapOfArrayOfStringMapOfUint64s: map<string,[]map<string,uint64>>
			# # comment
			- stringMapOfArrayOfStringMapOfStringOfUint64: map<string,[]map<map<string,string>,uint64>># comment
			# comment
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		structFields := p.root.Structs()[0].Fields()
		if assert.Equal(t, 8, len(structFields)) {
			assert.Equal(t, "[]string", structFields[0].Right().String())
			assert.Equal(t, "[][]string", structFields[1].Right().String())
			assert.Equal(t, "[][][]string", structFields[2].Right().String())
			assert.Equal(t, "map<string,string>", structFields[3].Right().String())
			assert.Equal(t, "map<string,uint64>", structFields[4].Right().String())
			assert.Equal(t, "map<string,[]uint64>", structFields[5].Right().String())
			assert.Equal(t, "map<string,[]map<string,uint64>>", structFields[6].Right().String())
			assert.Equal(t, "map<string,[]map<map<string,string>,uint64>>", structFields[7].Right().String())
		}
	}

	{
		p, err := newStringParser(`
		struct Role # comment
		# comment
			- name: string # comment
			- age: uint32 # comment
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		structFields := p.root.Structs()[0].Fields()
		if assert.Equal(t, 2, len(structFields)) {
			assert.Equal(t, "name", structFields[0].Left().String())
			assert.Equal(t, "string", structFields[0].Right().String())

			assert.Equal(t, "age", structFields[1].Left().String())
			assert.Equal(t, "uint32", structFields[1].Right().String())
		}
	}

	{
		p, err := newStringParser(`
		struct Role
			- name: string
				# comment
				+ go.tag.db = id # comment
				# comment
				+ json = id
				# comment
    + go.tag.json = created_at,omitempty
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		structFields := p.root.Structs()[0].Fields()
		if assert.Equal(t, 1, len(structFields)) {
			assert.Equal(t, "name", structFields[0].Left().String())
			assert.Equal(t, "string", structFields[0].Right().String())

			assert.Equal(t, "go.tag.db", structFields[0].Meta()[0].Left().String())
			assert.Equal(t, "id", structFields[0].Meta()[0].Right().String())

			assert.Equal(t, "json", structFields[0].Meta()[1].Left().String())
			assert.Equal(t, "id", structFields[0].Meta()[1].Right().String())

			assert.Equal(t, "go.tag.json", structFields[0].Meta()[2].Left().String())
			assert.Equal(t, "created_at,omitempty", structFields[0].Meta()[2].Right().String())
		}
	}

	{
		p, err := newStringParser(`
		struct User
			- ID: uint64
				+json = id
				+go.tag.db = id

			- username: string
				+ json= USERNAME #comment
				+ go.tag.db =username # commen
				# comment

			- role: string
				+ go.tag.db = -

			- createdAt?: timestamp
				+ json = created_at
				+ go.tag.json = created_at,omitempty
				+ go.tag.db = created_at
				+ go.tag.other = created_at,omitempty


		struct Notice
			- msg:string
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		structFields := p.root.Structs()[0].Fields()
		if assert.Equal(t, 4, len(structFields)) {
			assert.Equal(t, "ID", structFields[0].Left().String())
			assert.Equal(t, "uint64", structFields[0].Right().String())

			assert.Equal(t, "json", structFields[0].Meta()[0].Left().String())
			assert.Equal(t, "id", structFields[0].Meta()[0].Right().String())

			assert.Equal(t, "go.tag.db", structFields[0].Meta()[1].Left().String())
			assert.Equal(t, "id", structFields[0].Meta()[1].Right().String())

			assert.Equal(t, "username", structFields[1].Left().String())
			assert.Equal(t, "string", structFields[1].Right().String())

			assert.Equal(t, "json", structFields[1].Meta()[0].Left().String())
			assert.Equal(t, "USERNAME", structFields[1].Meta()[0].Right().String())

			assert.Equal(t, "go.tag.db", structFields[1].Meta()[1].Left().String())
			assert.Equal(t, "username", structFields[1].Meta()[1].Right().String())

			assert.Equal(t, "role", structFields[2].Left().String())
			assert.Equal(t, "string", structFields[2].Right().String())

			assert.Equal(t, "go.tag.db", structFields[2].Meta()[0].Left().String())
			assert.Equal(t, "-", structFields[2].Meta()[0].Right().String())

			assert.Equal(t, "createdAt", structFields[3].Left().String())
			assert.Equal(t, "timestamp", structFields[3].Right().String())

			assert.Equal(t, "json", structFields[3].Meta()[0].Left().String())
			assert.Equal(t, "created_at", structFields[3].Meta()[0].Right().String())

			assert.Equal(t, "go.tag.json", structFields[3].Meta()[1].Left().String())
			assert.Equal(t, "created_at,omitempty", structFields[3].Meta()[1].Right().String())

			assert.Equal(t, "go.tag.db", structFields[3].Meta()[2].Left().String())
			assert.Equal(t, "created_at", structFields[3].Meta()[2].Right().String())

			assert.Equal(t, "go.tag.other", structFields[3].Meta()[3].Left().String())
			assert.Equal(t, "created_at,omitempty", structFields[3].Meta()[3].Right().String())
		}
	}
}

func TestParserService(t *testing.T) {
	{
		p, err := newStringParser(`
			service
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.Error(t, err, "expecting service name")
	}

	{
		p, err := newStringParser(`
			service Empty
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)
	}

	{
		p, err := newStringParser(`
			service Pinger
				- Ping()
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)
	}

	{
		p, err := newStringParser(`
			service Pinger
				- Ping() => ()
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)
	}

	{
		p, err := newStringParser(`
			service Demo
				- DemoService(in: input) => (out: output)
				- DemoService2(in: map<string,map<string,uint64>>) => (out: map<string,uint64>)
				- DemoService3(i_n?: map<string,map<string,uint64>>) => (out: map<string,uint64>)
				- DemoService3(in?: map<string,map<string,uint64>>) => (ou_t?: map<string,uint64>)
				- stream Method(index: uint64, filter: string) => stream (t: time, data: string)
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)
	}
}

func TestParserServiceSuccint(t *testing.T) {
	p, err := newStringParser(`
		struct FlattenRequest
			- name: string
				+ go.tag.db = name
			- amount: uint64
				+ go.tag.db = amount

		struct FlattenResponse
			- id: uint64
				+ go.field.name = ID
			- count: uint64
				+ json = counter

		service Demo
			- DemoService(in: input) => (out: output)
			- Flatten(FlattenRequest) => (FlattenResponse)
	`)
	assert.NoError(t, err)

	err = p.run()
	assert.NoError(t, err)
}

func TestParserExamples(t *testing.T) {
	{
		p, err := newStringParser(`
			# contacts.ridl
			webrpc = v1

			struct Contact
				- id: int
				- name: string

			struct Counter
				- counter: int

			service ContactsService
				- GetContact(id: int) => (contact: Contact)
				- Version() => (details: any)
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)
	}

	{
		p, err := newStringParser(`
			# api.ridl
			webrpc = v1

			import "../contacts/proto/contacts.ridl"

			struct PingResponse
				- pong: string
				- counter: Counter # Counter is available here from the import

			service API
				- Ping() => (resp: PingResponse)
				- proxy ContactsService.GetContact # proxied rpc method :)
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)
	}
}

func TestParseStructComments(t *testing.T) {
	p, err := newStringParser(`
		# Defines role in our application
		struct Role # => role
          # role name line first
          # role name line second
		  # role name line third
		  - name: string # role name
		  - perms: []string # permissions
		  #- codesSubmitted: bool
		  #  + go.tag.db = codes_submitted

		  - countryCode: string # ie. US
			+ go.tag.db = country_code
		`)
	assert.NoError(t, err)

	err = p.run()
	assert.NoError(t, err)

	structNode, ok := p.root.node.children[0].(*StructNode)
	if !ok {
		t.Errorf("expected type StructNode")
	}

	assert.Equal(t, "Defines role in our application\n=> role", structNode.comment)

	assert.Equal(t, "role name line first\nrole name line second\nrole name line third\nrole name", structNode.fields[0].comment)
	assert.Equal(t, "permissions", structNode.fields[1].comment)
	assert.Equal(t, "ie. US", structNode.fields[2].comment)
}

func TestParseServiceComments(t *testing.T) {
	p, err := newStringParser(`
		# Contacts service line 1
		# Contacts service line 2
		service ContactsService  # Contacts service line 3
			#! skip this line
			# GetContact gives you contact for specific id
			# see https://www.example.com/?first=1&second=12#help
			- GetContact(id: int) => (contact: Contact)
			# Version returns you current deployed version
			# 
			#! skip this line as its internal comment
			#! skip more lines
			#! more
			@deprecated Version2
			- Version() => (details: any)
			# test comment two lines above
			
			- Version2() => (details: any)
		`)
	assert.NoError(t, err)

	err = p.run()
	assert.NoError(t, err)

	serviceNode, ok := p.root.node.children[0].(*ServiceNode)
	if !ok {
		t.Errorf("expected type ServiceNode")
	}

	assert.Equal(t, "Contacts service line 1\nContacts service line 2\nContacts service line 3", serviceNode.comment)
	assert.Equal(t, "GetContact gives you contact for specific id\nsee https://www.example.com/?first=1&second=12#help", serviceNode.methods[0].comment)
	assert.Equal(t, "Version returns you current deployed version\n", serviceNode.methods[1].comment)
	assert.Equal(t, "", serviceNode.methods[2].comment)
}

func TestParseAnnotations(t *testing.T) {
	p, err := newStringParser(`
		service ContactsService
			# Version returns you current deployed version
			@auth:x-access-key
			@deprecated:Version2 @internal @acl:admin
			- Version() => (details: any)
			@internal
			- Version2() => (details: any)
		`)
	assert.NoError(t, err)

	err = p.run()
	assert.NoError(t, err)

	serviceNode, ok := p.root.node.children[0].(*ServiceNode)
	if !ok {
		t.Errorf("expected type ServiceNode")
	}

	require.Len(t, serviceNode.methods[0].annotations, 4)
	require.Equal(t, "deprecated", serviceNode.methods[0].annotations[0].AnnotationType().String())
	require.Equal(t, "Version2", serviceNode.methods[0].annotations[0].Value().String())

	require.Equal(t, "internal", serviceNode.methods[0].annotations[1].AnnotationType().String())

	require.Equal(t, "acl", serviceNode.methods[0].annotations[2].AnnotationType().String())
	require.Equal(t, "admin", serviceNode.methods[0].annotations[2].Value().String())

	require.Equal(t, "auth", serviceNode.methods[0].annotations[3].AnnotationType().String())
	require.Equal(t, "x-access-key", serviceNode.methods[0].annotations[3].Value().String())

	require.Len(t, serviceNode.methods[1].annotations, 1)
	require.Equal(t, "internal", serviceNode.methods[1].annotations[0].AnnotationType().String())
}
