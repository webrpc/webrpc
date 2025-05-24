package webrpc

import (
	"fmt"
	"io/fs"
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
		// Use root FS to allow RIDL file imports from parent directories,
		// ie. import ../../common.ridl.

		// Support Windows paths. Currently only supports paths on the same volume.
		root := filepath.VolumeName(absolutePath) + "/"

		basePath, _ := os.Getwd()

		path, err := filepath.Rel(basePath, absolutePath)
		if err != nil {
			return nil, fmt.Errorf("get relative path %q: %w", path, err)
		}

		rootFS := rootFS{
			FS:       os.DirFS(root),
			BasePath: basePath,
		}

		return ridl.NewParser(&rootFS, path).Parse()

	default:
		return nil, fmt.Errorf("invalid schema file extension %q", ext)
	}
}

type rootFS struct {
	fs.FS
	BasePath string
}

func (r *rootFS) Open(path string) (fs.File, error) {
	return r.FS.Open(strings.TrimPrefix(filepath.Join(r.BasePath, path), "/"))
}
