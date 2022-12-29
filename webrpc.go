package webrpc

import (
	"fmt"
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
	switch ext {
	case ".json":
		json, err := os.ReadFile(absolutePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read %q: %w", path, err)
		}

		return schema.ParseSchemaJSON(json)

	case ".ridl":
		// Use root FS to allow RIDL file imports from parent directories,
		// ie. import ../../common.ridl.
		rootFS := os.DirFS("/")

		parser := ridl.NewParser(rootFS, absolutePath[1:])
		return parser.Parse()

	default:
		return nil, fmt.Errorf("error! invalid extension, %s: %w", ext, err)
	}
}
