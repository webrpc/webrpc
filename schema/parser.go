package schema

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type Parser interface {
	Parse(schema *WebRPCSchema) error
}

var NameWhitelistRexp = regexp.MustCompile(`(?i)^[a-zA-Z0-9]+(\w*[a-zA-Z0-9]+)?$`)

func ParseSchema(data []byte) (*WebRPCSchema, error) {
	var schema *WebRPCSchema
	err := json.Unmarshal(data, &schema)
	if err != nil {
		return nil, err
	}

	err = schema.Parse(nil)
	if err != nil {
		return schema, err
	}

	return schema, nil
}

func ParseSchemaFile(schemaFilePath string) (*WebRPCSchema, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	var path string
	if schemaFilePath[0:1] == "/" {
		path = schemaFilePath
	} else {
		path = filepath.Join(cwd, schemaFilePath)
	}

	// ensure schema file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	// read file contents
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return ParseSchema(contents)
}

func IsValidArgName(s string) bool {
	// TODO: fix this regexp
	if !NameWhitelistRexp.MatchString(s) {
		return false
	}
	return true
}
