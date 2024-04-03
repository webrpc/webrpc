package gen

import (
	"embed"

	golang "github.com/webrpc/gen-golang"
)

// Baked-in templates officially supported by webrpc-gen tooling.
var DefaultTemplates = map[string]embed.FS{
	"golang": golang.FS,
}
