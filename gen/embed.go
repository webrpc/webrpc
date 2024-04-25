package gen

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/webrpc/webrpc"

	dart "github.com/webrpc/gen-dart"
	golang "github.com/webrpc/gen-golang"
	javascript "github.com/webrpc/gen-javascript"
	kotlin "github.com/webrpc/gen-kotlin"
	openapi "github.com/webrpc/gen-openapi"
	typescript "github.com/webrpc/gen-typescript"
)

// Embedded templates officially supported by webrpc-gen tooling.
// Versioning is managed via Go modules (go.mod file).
//
// To propose a new officially supported github.com/webrpc/gen-* template,
// please submit an issue at https://github.com/webrpc/webrpc/issues/new.
var embeddedTargetFS = map[string]embed.FS{
	"golang":     golang.FS,
	"typescript": typescript.FS,
	"javascript": javascript.FS,
	"openapi":    openapi.FS,
	"kotlin":     kotlin.FS,
	"dart":       dart.FS,
}

// The values are computed in init() function based on go.mod file.
var (
	EmbeddedTargetNames = []string{}
	EmbeddedTargets     = map[string]EmbeddedTarget{}
)

type EmbeddedTarget struct {
	Name      string
	Version   string
	ImportTag string
	FS        embed.FS
}

func init() {
	// Parse target versions from go.mod file
	scanner := bufio.NewScanner(strings.NewReader(webrpc.GoModFile))
	for scanner.Scan() {
		//	github.com/webrpc/gen-golang v0.14.2 // comment
		line := scanner.Text()
		prefix := "\tgithub.com/webrpc/gen-"
		if !strings.HasPrefix(line, prefix) {
			// NOTE: Use strings.CutPrefix() once we decide to bump go.mod to Go 1.20.
			continue
		}
		parts := strings.Split(line[len(prefix):], " ")
		if len(parts) < 2 {
			continue
		}
		name, version := parts[0], parts[1]
		fs, ok := embeddedTargetFS[name]
		if !ok {
			fmt.Fprintf(os.Stderr, "%s embed FS not found", name)
			continue
		}
		EmbeddedTargetNames = append(EmbeddedTargetNames, name)

		for _, gen := range []string{
			name,                                // golang
			fmt.Sprintf("%s@%s", name, version), // golang@v0.14.2
			fmt.Sprintf("github.com/webrpc/gen-%s", name),             // github.com/webrpc/gen-golang
			fmt.Sprintf("github.com/webrpc/gen-%s@%s", name, version), // github.com/webrpc/gen-golang@v0.14.2
		} {
			EmbeddedTargets[gen] = EmbeddedTarget{
				Name:      name,
				ImportTag: fmt.Sprintf("github.com/webrpc/gen-%s@%s", name, version),
				Version:   version,
				FS:        fs,
			}
		}
	}
}
