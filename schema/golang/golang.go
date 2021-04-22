package golang

import (
	"go/types"
	"path"

	"github.com/pkg/errors"
	"github.com/webrpc/webrpc/schema"
	"golang.org/x/tools/go/packages"
)

func NewParser(r *schema.Reader) *parser {
	return &parser{
		schema:          &schema.WebRPCSchema{},
		parsedTypes:     map[types.Type]*schema.VarType{},
		parsedTypeNames: map[string]struct{}{},

		// TODO: Change this to map[*types.Package]string so we can rename duplicated pkgs?
		resolvedImports: map[string]struct{}{
			// Initial schema file's package name artificially set by golang.org/x/tools/go/packages.
			"command-line-arguments": struct{}{},

			// The following imports are already defined in the Go template.
			"context":       struct{}{},
			"encoding/json": struct{}{},
			"fmt":           struct{}{},
			"io/ioutil":     struct{}{},
			"net/http":      struct{}{},
			"time":          struct{}{},
			"strings":       struct{}{},
			"bytes":         struct{}{},
			"errors":        struct{}{},
			"io":            struct{}{},
			"net/url":       struct{}{},
		},
	}
}

// Parses Go source file and returns WebRPC schema.
//
// This parser was designed to run sequentially, without any concurrency, so we can leverage
// maps to cache the already parsed types, while not having to deal with sync primitives.
type parser struct {
	schema *schema.WebRPCSchema

	// Cache for already parsed types, to improve performance & so we can traverse circular dependencies.
	parsedTypes     map[types.Type]*schema.VarType
	parsedTypeNames map[string]struct{}

	inlineMode      bool // When traversing `json:",inline"`, we don't want to store the struct type as WebRPC message.
	resolvedImports map[string]struct{}

	schemaPkgName string // Shema file's package name.
}

// Parses Go source file and return WebRPC schema.
func (p *parser) Parse(filePath string) (*schema.WebRPCSchema, error) {
	cfg := &packages.Config{
		Dir:  path.Dir(filePath),
		Mode: packages.NeedName | packages.NeedImports | packages.NeedTypes | packages.NeedFiles | packages.NeedDeps | packages.NeedSyntax,
	}

	schemaPkg, err := packages.Load(cfg, filePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load packages")
	}
	if len(schemaPkg) != 1 {
		return nil, errors.Errorf("failed to load initial package (len=%v)", len(schemaPkg))
	}

	p.schemaPkgName = schemaPkg[0].Name

	err = p.parsePkgInterfaces(schemaPkg[0].Types.Scope())
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse Go interfaces")
	}

	return p.schema, nil
}
