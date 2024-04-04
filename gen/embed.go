package gen

import (
	"embed"

	golang "github.com/webrpc/gen-golang"
	javascript "github.com/webrpc/gen-javascript"
	kotlin "github.com/webrpc/gen-kotlin"
	openapi "github.com/webrpc/gen-openapi"
	typescript "github.com/webrpc/gen-typescript"
)

// Embedded templates officially supported by webrpc-gen tooling.
// Versioning is managed via Go modules (go.mod file).
//
// NOTE: We only allow officially supported github.com/webrpc/gen-* templates here.
//
// Please, submit new templates at https://github.com/webrpc/webrpc/issues/new.
var DefaultEmbeddedTemplates = map[string]embed.FS{
	"golang":     golang.FS,
	"typescript": typescript.FS,
	"javascript": javascript.FS,
	"openapi":    openapi.FS,
	"kotlin":     kotlin.FS,
}
