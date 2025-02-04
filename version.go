package webrpc

import (
	_ "embed"
	"os/exec"
	"strings"
)

// Version of webrpc-gen tooling.
// Available as {{.WebrpcGenVersion}} variable in Go templates.
//
// The value can be injected during the build step:
//
//	go build -ldflags="-s -w -X github.com/webrpc/webrpc.VERSION=$$(git describe --tags)" ./cmd/webrpc-gen
//
// Alternatively, the value is inferred in the runtime from:
//  1. go.mod file (useful when running webrpc-gen as a tool from another Go module)
//  2. git history
var VERSION = ""

func init() {
	if VERSION == "" {
		VERSION = inferRuntimeVersion()
	}
}

func inferRuntimeVersion() string {
	// $ go list -m -f {{.Version}} github.com/webrpc/webrpc
	// v0.22.0\n
	if out, _ := exec.Command("go", "list", "-m", "-f", "{{.Version}}", "github.com/webrpc/webrpc").Output(); len(out) > 0 {
		return strings.TrimSpace(string(out))
	}

	// $ git describe --tags
	// v0.15.1-6-g550333d\n
	if out, _ := exec.Command("git", "describe", "--tags").Output(); len(out) > 0 {
		return strings.TrimSpace(string(out))
	}

	return "unknown"
}

// Embed go.mod file, so we can parse version of the embedded generators in the runtime.
//
//go:embed go.mod
var GoModFile string
