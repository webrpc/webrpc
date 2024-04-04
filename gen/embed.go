package gen

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/webrpc/webrpc"

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
}

// The values are computed in init() function based on go.mod file.
var (
	EmbeddedTargets        = []string{}
	EmbeddedTargetFS       = map[string]embed.FS{}
	EmbeddedTargetVersions = map[string]string{}
)

func init() {
	// Parse target versions from go.mod file
	scanner := bufio.NewScanner(strings.NewReader(webrpc.GoModFile))
	for scanner.Scan() {
		//	github.com/webrpc/gen-golang v0.14.2 // comment
		after, ok := strings.CutPrefix(scanner.Text(), "\tgithub.com/webrpc/gen-")
		if !ok {
			continue
		}
		parts := strings.Split(after, " ")
		if len(parts) < 2 {
			continue
		}
		gen, version := parts[0], parts[1]
		fs, ok := embeddedTargetFS[gen]
		if !ok {
			fmt.Fprintf(os.Stderr, "%s embed FS not found", gen)
			continue
		}
		EmbeddedTargets = append(EmbeddedTargets, gen)
		EmbeddedTargetVersions[gen] = version

		// Cache embedded template for golang, golang@v0.14.2, github.com/webrpc/gen-golang, github.com/webrpc/gen-golang@v0.14.2
		EmbeddedTargetFS[gen] = fs
		EmbeddedTargetFS[fmt.Sprintf("%s@%s", gen, version)] = fs
		EmbeddedTargetFS[fmt.Sprintf("github.com/webrpc/gen-%s", gen)] = fs
		EmbeddedTargetFS[fmt.Sprintf("github.com/webrpc/gen-%s@%s", gen, version)] = fs
	}
}
