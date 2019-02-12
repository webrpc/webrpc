package ridl

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func SkipTestLexer(t *testing.T) {
	buf := `


		ridl						 =  v1

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
	ridl = v1
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
	version = v0.1.1
	service = hello-webrpc
	`
		_, err := Parse(buf)
		assert.NoError(t, err)
	}
}

func TestHeadersImport(t *testing.T) {
	{
		input := `
	ridl = v1
		version = v0.1.1
	service = hello-webrpc

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
}

func SkipTestParse(t *testing.T) {
	fp, err := os.Open("example1.ridl")
	assert.NoError(t, err)

	buf, err := ioutil.ReadAll(fp)
	assert.NoError(t, err)

	schema, err := Parse(string(buf))
	assert.NoError(t, err)

	log.Printf("buf: %v", string(buf))
	log.Printf("schema: %v", schema)
}
