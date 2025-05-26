package webrpc

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
		root, _, err := getRootPath()
		if err != nil {
			return nil, fmt.Errorf("get root path: %w", err)
		}

		schema, err := ridl.NewParser(os.DirFS(root), filepath.ToSlash(absolutePath[len(root):])).Parse()
		if err != nil {
			return nil, err
		}

		// Convert absolute paths to relative paths.
		if root, wd, _ := getRootPath(); wd != "" {
			basePath := strings.TrimPrefix(wd, root)
			schema.Filename, _ = filepath.Rel(basePath, schema.Filename)
			for _, t := range schema.Types {
				if t.Filename != "" {
					t.Filename, _ = filepath.Rel(basePath, t.Filename)
				}
			}
		}

		return schema, nil

	default:
		return nil, fmt.Errorf("invalid schema file extension %q", ext)
	}
}

func getRootPath() (root, wd string, err error) {
	if wd, err = os.Getwd(); err != nil {
		return "", "", fmt.Errorf("get working directory: %w", err)
	}

	// Use root FS to allow RIDL file imports from parent directories,
	// ie. import ../../common.ridl.
	root = "/"
	// Support Windows paths. Currently only supports paths on the same volume.
	if volume := filepath.VolumeName(wd); volume != "" {
		root = volume + "/"
	}
	return root, wd, nil
}
