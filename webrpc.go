package webrpc

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/webrpc/webrpc/schema"
	"github.com/webrpc/webrpc/schema/ridl"
)

func ParseSchemaFile(schemaFilePath string) (*schema.WebRPCSchema, error) {
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

	ext := filepath.Ext(path)
	if ext == ".json" {
		return schema.ParseSchemaJSON(contents)
	} else if ext == ".ridl" {
		return ridl.Parse(string(contents))
	} else {
		return nil, errors.Errorf("error! invalid extension, %s", ext)
	}
}
