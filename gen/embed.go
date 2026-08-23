package gen

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/webrpc/webrpc"
)

// Templates officially supported by webrpc-gen tooling, embedded from the
// gen/<target> folders of this repository and released atomically with
// the webrpc-gen binary.
//
// To propose a new officially supported target template,
// please submit an issue at https://github.com/webrpc/webrpc/issues/new.
//
//go:embed golang/*.go.tmpl typescript/*.go.tmpl javascript/*.go.tmpl openapi/*.go.tmpl kotlin/*.go.tmpl dart/*.go.tmpl
var templatesFS embed.FS

var EmbeddedTargetNames = []string{
	"golang",
	"typescript",
	"javascript",
	"openapi",
	"kotlin",
	"dart",
}

// The values are computed in init() function.
var EmbeddedTargets = map[string]EmbeddedTarget{}

type EmbeddedTarget struct {
	Name      string
	Version   string
	ImportTag string
	FS        fs.FS
}

func init() {
	for _, name := range EmbeddedTargetNames {
		subFS, err := fs.Sub(templatesFS, name)
		if err != nil {
			panic(fmt.Errorf("failed to load embedded %s templates: %w", name, err))
		}

		target := EmbeddedTarget{
			Name:      name,
			Version:   webrpc.VERSION,
			ImportTag: fmt.Sprintf("%s@%s", name, webrpc.VERSION),
			FS:        subFS,
		}

		for _, alias := range []string{
			name, // golang
			fmt.Sprintf("%s@%s", name, webrpc.VERSION), // golang@v0.45.0
			// Legacy import paths of the former standalone template repos.
			fmt.Sprintf("github.com/webrpc/gen-%s", name),                    // github.com/webrpc/gen-golang
			fmt.Sprintf("github.com/webrpc/gen-%s@%s", name, webrpc.VERSION), // github.com/webrpc/gen-golang@v0.45.0
		} {
			EmbeddedTargets[alias] = target
		}
	}
}
