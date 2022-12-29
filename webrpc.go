package webrpc

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/webrpc/webrpc/schema"
	"github.com/webrpc/webrpc/schema/ridl"
)

func ParseSchemaFile(path string) (*schema.WebRPCSchema, error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path %q: %w", path, err)
	}

	ext := filepath.Ext(absolutePath)
	if ext == ".json" {
		// ensure schema file exists
		if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
			return nil, err
		}

		// open file
		fp, err := os.Open(absolutePath)
		if err != nil {
			return nil, err
		}
		defer fp.Close()

		// TODO: implement ParseSchemaJSON with io.Reader or read contents
		// before passing them.
		contents, err := ioutil.ReadAll(fp)
		if err != nil {
			return nil, err
		}

		return schema.ParseSchemaJSON(contents)
	} else if ext == ".ridl" {
		// We need root to allow RIDL imports from parent directories (import ../file.ridl)
		rootFS := os.DirFS("/")

		parser := ridl.NewParser(rootFS, absolutePath[1:])

		s, err := parser.Parse()
		if err != nil {
			return nil, err
		}

		return s, nil
	} else {
		return nil, fmt.Errorf("error! invalid extension, %s: %w", ext, err)
	}
}
