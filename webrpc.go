package webrpc

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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

	// open file
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	ext := filepath.Ext(path)
	if ext == ".json" {
		// TODO: implement ParseSchemaJSON with io.Reader or read contents
		// before passing them.
		contents, err := ioutil.ReadAll(fp)
		if err != nil {
			return nil, err
		}

		return schema.ParseSchemaJSON(contents)
	} else if ext == ".ridl" {
		rdr := ridl.NewParser(schema.NewReader(fp, path))
		s, err := rdr.Parse()
		if err != nil {
			return nil, err
		}

		return s, nil
	} else {
		return nil, fmt.Errorf("error! invalid extension, %s: %w", ext, err)
	}
}
