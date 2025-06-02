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
		root := "/"
		// Support Windows paths. Currently only supports paths on the same volume.
		if volume := filepath.VolumeName(path); volume != "" {
			root = volume + "/"
		}

		path := filepath.ToSlash(absolutePath[len(root):])
		schema, err := ridl.NewParser(os.DirFS(root), root, path).Parse()
		if err != nil {
			return nil, err
		}

		return schema, nil

	default:
		return nil, fmt.Errorf("invalid schema file extension %q", ext)
	}
}
