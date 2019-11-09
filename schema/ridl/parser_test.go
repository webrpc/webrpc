package ridl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		assert.Equal(t, "packageName.ridl", p.root.Imports()[0].Path().String())
	}

	{
		p, err := newStringParser(`
			import # comment
			-foo
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.Equal(t, "foo", p.root.Imports()[0].Path().String())
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

		assert.Equal(t, "foo", p.root.Imports()[0].Path().String())
	}

	{
		p, err := newStringParser(`
			import # comment
		# comment
			#comment
			- foo
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.Equal(t, "foo", p.root.Imports()[0].Path().String())
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

		assert.Equal(t, 1, len(p.root.Imports()))
		assert.Equal(t, 1, len(p.root.Children()))
		assert.Equal(t, "foo.ridl", p.root.Imports()[0].Path().String())
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
		assert.Equal(t, "x", p.root.Imports()[0].Path().String())
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
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.Equal(t, 3, len(p.root.Imports()))

		assert.Equal(t, "foo.ridl", p.root.Imports()[0].Path().String())
		assert.Equal(t, "./path/to/bar.ridl", p.root.Imports()[1].Path().String())
		assert.Equal(t, "baz_-", p.root.Imports()[2].Path().String())
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

		assert.Equal(t, "/path/to/file.ridl", p.root.Imports()[0].Path().String())
		assert.Equal(t, "Member", p.root.Imports()[0].Members()[0].String())
		assert.Equal(t, "Name", p.root.Imports()[0].Members()[1].String())

		assert.Equal(t, "/path/to/file2.ridl", p.root.Imports()[1].Path().String())
		assert.Equal(t, "Member", p.root.Imports()[1].Members()[0].String())
		assert.Equal(t, "Name", p.root.Imports()[1].Members()[1].String())
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
		assert.Equal(t, "/path /to foo.ridl", p.root.Imports()[0].Path().String())
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
		assert.Equal(t, "/path /to foo.ridl", p.root.Imports()[0].Path().String())
		assert.Equal(t, "skywalker.ridl", p.root.Imports()[1].Path().String())
		assert.Equal(t, "  ", p.root.Imports()[2].Path().String())
		assert.Equal(t, " \nABC \" DEF ", p.root.Imports()[3].Path().String())
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

		assert.Equal(t, "./users.ridl", p.root.Imports()[0].Path().String())
		assert.Equal(t, "./users # .ridl", p.root.Imports()[1].Path().String())
		assert.Equal(t, "uno", p.root.Imports()[2].Path().String())
		assert.Equal(t, "dos", p.root.Imports()[3].Path().String())
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
		assert.Equal(t, "hello", p.root.Imports()[0].Path().String())

		assert.Equal(t, "Value", p.root.Enums()[0].values[0].Left().String())
		assert.Equal(t, "value2", p.root.Enums()[0].values[1].Left().String())

		assert.Equal(t, "USER", p.root.Enums()[1].values[0].Left().String())
		assert.Equal(t, "ADMIN", p.root.Enums()[1].values[1].Left().String())
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
		assert.Equal(t, "hello", p.root.Imports()[0].Path().String())

		assert.Equal(t, "Value", p.root.Enums()[0].values[0].Left().String())
		assert.Equal(t, "SKY", p.root.Enums()[0].values[0].Right().String())

		assert.Equal(t, "value2", p.root.Enums()[0].values[1].Left().String())
		assert.Equal(t, "WALKER", p.root.Enums()[0].values[1].Right().String())

		assert.Equal(t, "USER", p.root.Enums()[1].values[0].Left().String())
		assert.Equal(t, "ADMIN", p.root.Enums()[1].values[1].Left().String())

		assert.Equal(t, "USER", p.root.Enums()[3].values[0].Left().String())
		assert.Equal(t, "ADMIN", p.root.Enums()[3].values[1].Left().String())
	}
}

func TestParserMessage(t *testing.T) {
	{
		p, err := newStringParser(`
			message
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.Error(t, err, "expecting message name")
	}

	{
		p, err := newStringParser(`
			message Empty
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)
	}

	{
		p, err := newStringParser(`
		message Role
			- name: string
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.Equal(t, "name", p.root.Messages()[0].Fields()[0].Left().String())
		assert.Equal(t, "string", p.root.Messages()[0].Fields()[0].Right().String())
	}

	{
		p, err := newStringParser(`
		message ComplexTypes
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

		assert.Equal(t, "[]string", p.root.Messages()[0].Fields()[0].Right().String())
		assert.Equal(t, "[][]string", p.root.Messages()[0].Fields()[1].Right().String())
		assert.Equal(t, "[][][]string", p.root.Messages()[0].Fields()[2].Right().String())
		assert.Equal(t, "map<string,string>", p.root.Messages()[0].Fields()[3].Right().String())
		assert.Equal(t, "map<string,uint64>", p.root.Messages()[0].Fields()[4].Right().String())
		assert.Equal(t, "map<string,[]uint64>", p.root.Messages()[0].Fields()[5].Right().String())
		assert.Equal(t, "map<string,[]map<string,uint64>>", p.root.Messages()[0].Fields()[6].Right().String())
		assert.Equal(t, "map<string,[]map<map<string,string>,uint64>>", p.root.Messages()[0].Fields()[7].Right().String())
	}

	{
		p, err := newStringParser(`
		message Role # comment
		# comment
			- name: string # comment
			- age: uint32 # comment
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.Equal(t, "name", p.root.Messages()[0].Fields()[0].Left().String())
		assert.Equal(t, "string", p.root.Messages()[0].Fields()[0].Right().String())

		assert.Equal(t, "age", p.root.Messages()[0].Fields()[1].Left().String())
		assert.Equal(t, "uint32", p.root.Messages()[0].Fields()[1].Right().String())
	}

	{
		p, err := newStringParser(`
		message Role
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

		assert.Equal(t, "name", p.root.Messages()[0].Fields()[0].Left().String())
		assert.Equal(t, "string", p.root.Messages()[0].Fields()[0].Right().String())

		assert.Equal(t, "go.tag.db", p.root.Messages()[0].Fields()[0].Meta()[0].Left().String())
		assert.Equal(t, "id", p.root.Messages()[0].Fields()[0].Meta()[0].Right().String())

		assert.Equal(t, "json", p.root.Messages()[0].Fields()[0].Meta()[1].Left().String())
		assert.Equal(t, "id", p.root.Messages()[0].Fields()[0].Meta()[1].Right().String())

		assert.Equal(t, "go.tag.json", p.root.Messages()[0].Fields()[0].Meta()[2].Left().String())
		assert.Equal(t, "created_at,omitempty", p.root.Messages()[0].Fields()[0].Meta()[2].Right().String())
	}

	{
		p, err := newStringParser(`
		message User
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


		message Notice
			- msg:string
		`)
		assert.NoError(t, err)

		err = p.run()
		assert.NoError(t, err)

		assert.Equal(t, "ID", p.root.Messages()[0].Fields()[0].Left().String())
		assert.Equal(t, "uint64", p.root.Messages()[0].Fields()[0].Right().String())

		assert.Equal(t, "json", p.root.Messages()[0].Fields()[0].Meta()[0].Left().String())
		assert.Equal(t, "id", p.root.Messages()[0].Fields()[0].Meta()[0].Right().String())

		assert.Equal(t, "go.tag.db", p.root.Messages()[0].Fields()[0].Meta()[1].Left().String())
		assert.Equal(t, "id", p.root.Messages()[0].Fields()[0].Meta()[1].Right().String())

		assert.Equal(t, "username", p.root.Messages()[0].Fields()[1].Left().String())
		assert.Equal(t, "string", p.root.Messages()[0].Fields()[1].Right().String())

		assert.Equal(t, "json", p.root.Messages()[0].Fields()[1].Meta()[0].Left().String())
		assert.Equal(t, "USERNAME", p.root.Messages()[0].Fields()[1].Meta()[0].Right().String())

		assert.Equal(t, "go.tag.db", p.root.Messages()[0].Fields()[1].Meta()[1].Left().String())
		assert.Equal(t, "username", p.root.Messages()[0].Fields()[1].Meta()[1].Right().String())

		assert.Equal(t, "role", p.root.Messages()[0].Fields()[2].Left().String())
		assert.Equal(t, "string", p.root.Messages()[0].Fields()[2].Right().String())

		assert.Equal(t, "go.tag.db", p.root.Messages()[0].Fields()[2].Meta()[0].Left().String())
		assert.Equal(t, "-", p.root.Messages()[0].Fields()[2].Meta()[0].Right().String())

		assert.Equal(t, "createdAt", p.root.Messages()[0].Fields()[3].Left().String())
		assert.Equal(t, "timestamp", p.root.Messages()[0].Fields()[3].Right().String())

		assert.Equal(t, "json", p.root.Messages()[0].Fields()[3].Meta()[0].Left().String())
		assert.Equal(t, "created_at", p.root.Messages()[0].Fields()[3].Meta()[0].Right().String())

		assert.Equal(t, "go.tag.json", p.root.Messages()[0].Fields()[3].Meta()[1].Left().String())
		assert.Equal(t, "created_at,omitempty", p.root.Messages()[0].Fields()[3].Meta()[1].Right().String())

		assert.Equal(t, "go.tag.db", p.root.Messages()[0].Fields()[3].Meta()[2].Left().String())
		assert.Equal(t, "created_at", p.root.Messages()[0].Fields()[3].Meta()[2].Right().String())

		assert.Equal(t, "go.tag.other", p.root.Messages()[0].Fields()[3].Meta()[3].Left().String())
		assert.Equal(t, "created_at,omitempty", p.root.Messages()[0].Fields()[3].Meta()[3].Right().String())
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

func TestParserExamples(t *testing.T) {
	{
		p, err := newStringParser(`
			# contacts.ridl
			webrpc = v1

			message Contact
				- id: int
				- name: string

			message Counter
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

			message PingResponse
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
