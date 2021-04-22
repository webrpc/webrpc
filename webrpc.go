package webrpc

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/webrpc/webrpc/schema"
	"github.com/webrpc/webrpc/schema/golang"
	"github.com/webrpc/webrpc/schema/ridl"
)

const VERSION = "v0.6.0"

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
			return nil, errors.Wrap(err, "failed to parse RIDL schema")
		}

		return s, nil
	} else if ext == ".go" {
		goParser := golang.NewParser(schema.NewReader(fp, path))
		s, err := goParser.Parse(path)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse Go schema")
		}

		return s, nil
	} else {
		return nil, errors.Errorf("unknown schema file extension .%s", ext)
	}
}
