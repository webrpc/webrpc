package ridl

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	buf := `


		webrpc						 =  v1

				+     foo=bar

					-baz   = 56 # a comment

													version=                    v0.0.1


foo=bar`

	tokens, err := tokenize(buf)
	assert.NoError(t, err)

	log.Printf("buf: %v", string(buf))
	log.Printf("tokens: %v", tokens)
}

func TestRidlHeader(t *testing.T) {
	{
		buf := `
		webrpc = v1
	`
		_, err := Parse(buf)
		assert.Error(t, err, `"version" is required`)
	}
	{
		buf := `
	webrpc = v0
	webrpc = v1
	`
		_, err := Parse(buf)
		assert.Error(t, err, `should not be able to declare "ridl" twice`)
	}
}

func TestHeaders(t *testing.T) {
	{
		buf := `
		webrpc = v1
	version = v0.1.1

	name= hello-webrpc
	`
		_, err := Parse(buf)
		assert.NoError(t, err)
	}
}

func TestImport(t *testing.T) {
	{
		input := `
		webrpc = v1
			version = v0.1.1
	name = hello-webrpc

		import
		- foo # ko ment
		# ko ment

			- bar
			# comment
		`
		schema, err := Parse(input)
		log.Printf("ERR: %v", err)
		assert.NoError(t, err)

		log.Printf("schema: %v", schema)
		buf, err := json.Marshal(schema)
		assert.NoError(t, err)
		log.Printf("schema JSON: %v", string(buf))

	}

	{
		input := `
		webrpc = v1
		version = v0.1.1 # version number
	name		 = hello-webrpc

	import # import line
	- foo1 # foo-comment with spaces
		- bar2 # # # bar-comment
	`
		schema, err := Parse(input)
		assert.NoError(t, err)

		log.Printf("schema: %v", schema)

		buf, err := json.MarshalIndent(schema, "", "  ")
		assert.NoError(t, err)
		log.Printf("schema JSON: %v", string(buf))
	}
}

func TestEnum(t *testing.T) {
	{
		input := `
		webrpc = v1
		version = v0.1.1
	name = hello-webrpc

					# this is a comment
						# yep
					enum Kind:uint32
						- USER = 1             # comment
						- ADMIN = 2            # comment..

				# or.. just..
				enum	 KindTwo:		uint32
					- USER                 # aka, = 0
					- ADMIN         # aka, = 1
					- OTHER
	`
		schema, err := Parse(input)
		assert.NoError(t, err)

		log.Printf("schema: %v", schema)

		buf, err := json.MarshalIndent(schema, "", "  ")
		assert.NoError(t, err)
		log.Printf("schema JSON: %v", string(buf))

	}
}

func TestMessages(t *testing.T) {
	{
		input := `
		webrpc = v1
		version = v0.1.1
	name = hello-webrpc

	message Empty
	`
		schema, err := Parse(input)
		assert.NoError(t, err)

		log.Printf("schema: %v", schema)

		buf, err := json.MarshalIndent(schema, "", "  ")
		assert.NoError(t, err)
		log.Printf("schema JSON: %v", string(buf))

	}

	{
		input := `
		webrpc = v1
		version = v0.1.1
	name = hello-webrpc

	message Empty # with a, comment
	`
		schema, err := Parse(input)
		assert.NoError(t, err)

		log.Printf("schema: %v", schema)

		buf, err := json.MarshalIndent(schema, "", "  ")
		assert.NoError(t, err)
		log.Printf("schema JSON: %v", string(buf))

	}

	{
		input := `
		webrpc = v1
		version = v0.1.1
	name = hello-webrpc

	message Simple # with a, comment
		- ID: uint32
	`
		schema, err := Parse(input)
		assert.NoError(t, err)

		log.Printf("schema: %v", schema)

		buf, err := json.MarshalIndent(schema, "", "  ")
		assert.NoError(t, err)
		log.Printf("schema JSON: %v", string(buf))

	}

	{
		input := `
		webrpc = v1
		version = v0.1.1
	name = hello-webrpc

	message Simple # with a-comment an,d meta fields
		- ID: uint32
	- Field2: uint64 # one two #t
			+ json = field_2 # a comment
				+ go.tag.db = field_2
	- Field3: bool
		+ go.tag.db = - # omits the field from db


	message Simple2 # with a-comment an,d meta fields
	`
		schema, err := Parse(input)
		assert.NoError(t, err)

		log.Printf("schema: %v", schema)

		buf, err := json.MarshalIndent(schema, "", "  ")
		assert.NoError(t, err)
		log.Printf("schema JSON: %v", string(buf))
	}

	{
		input := `
		webrpc = v1
		version = v0.1.1
	name = hello-webrpc

	message Simple # with a-comment an,d meta fields
		- ID: uint32
	- Field2: map<string, string> # one two #t
			+ json = field_2 # a comment
				+ go.tag.db = field_2
	- Field3: []bool # one two #t
			+ json = field_2 # a comment
				+ go.tag.db = field_2
	- Field4: [][][]bool # one two #t
			+ json = field_2 # a comment
				+ go.tag.db = field_2

	message Simple2 # with a-comment an,d meta fields
	`
		schema, err := Parse(input)
		assert.NoError(t, err)

		log.Printf("schema: %v", schema)

		buf, err := json.MarshalIndent(schema, "", "  ")
		assert.NoError(t, err)
		log.Printf("schema JSON: %v", string(buf))
	}
}

func TestService(t *testing.T) {
	{
		input := `
		webrpc = v1
		version = v0.1.1
	name = hello-webrpc

	service Empty
	`
		schema, err := Parse(input)
		assert.NoError(t, err)

		log.Printf("schema: %v", schema)

		buf, err := json.MarshalIndent(schema, "", "  ")
		assert.NoError(t, err)
		log.Printf("schema JSON: %v", string(buf))
	}

	{
		input := `
		webrpc = v1
		version = v0.1.1
	name = hello-webrpc

	service Simple
	- Ping()
	-	Status() => (status: bool)
	-	StatusStream(q: string) => (status: stream bool)`
		schema, err := Parse(input)
		assert.NoError(t, err)

		log.Printf("schema: %v", schema)

		buf, err := json.MarshalIndent(schema, "", "  ")
		assert.NoError(t, err)
		log.Printf("schema JSON: %v", string(buf))
	}

	{
		input := `
		webrpc = v1
		version = v0.1.1
	name = hello-webrpc

	service Simple
	-	Ping(code: stream uint32) => (code: bool)
	-	PingStream(text: string) => (code: stream bool)`
		schema, err := Parse(input)
		assert.NoError(t, err)

		log.Printf("schema: %v", schema)

		buf, err := json.MarshalIndent(schema, "", "  ")
		assert.NoError(t, err)
		log.Printf("schema JSON: %v", string(buf))
	}
}

func TestParse(t *testing.T) {
	fp, err := os.Open("example1.ridl")
	assert.NoError(t, err)

	buf, err := ioutil.ReadAll(fp)
	assert.NoError(t, err)

	schema, err := Parse(string(buf))
	assert.NoError(t, err)

	log.Printf("buf: %v", string(buf))

	jout, err := schema.ToJSON(true)
	assert.NoError(t, err)
	log.Printf("schema JSON: %v", jout)
}
