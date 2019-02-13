package ridl

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRidlHeader(t *testing.T) {
	{
		buf := `
	ridl = v1
	name = myapp
	version = 1
	`
		_, err := Parse(buf)
		assert.NoError(t, err)
	}
	{
		buf := `
	ridl = v0
	ridl = v1
	`
		_, err := Parse(buf)
		assert.Error(t, err)
	}
}

func TestHeaders(t *testing.T) {
	{
		buf := `
	ridl = v1

	name = hello-webrpc
	version = v0.1.1
	`
		_, err := Parse(buf)
		assert.NoError(t, err)
	}
}

func TestImport(t *testing.T) {
	{
		input := `
		ridl = v1
		name = a
		version = b

		import
		- foo
			- bar
		`
		schema, err := Parse(input)
		assert.NoError(t, err)

		log.Printf("schema: %v", schema)

		buf, err := json.Marshal(schema)
		assert.NoError(t, err)
		log.Printf("schema JSON: %v", string(buf))

	}

	{
		input := `
	ridl = v1
	name = hello-webrpc
	version = v0.1.1 # version number

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
	ridl = v1
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
	ridl = v1
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
	ridl = v1
	name = hello-webrpc
		version = v0.1.1

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
	 ridl = v1

	 name = hello-webrpc
	 version = v0.1.1

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
	ridl = v1

	name = hello-webrpc
		version = v0.1.1


	message Simple # with a-comment an,d meta fields
		- ID: uint32
	- Field2: uint64 # one two #t
			+ json = field_2 # a comment
				+ go.tag.db = field_2


	message Simple2 # with a-comment an,d meta fields


	# complex message
	message Complex

		- username: string # user's name

		-	dict: map<string,Simple>
			+ json = cool

		- list: []string

		- more: []Simple

		- createdAt?: timestamp

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
	ridl = v1
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
	ridl = v1
		version = v0.1.1
	name = hello-webrpc

	service Simple
	-	Ping(): bool
	`
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
	log.Printf("schema: %v", schema)
}
