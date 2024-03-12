package webrpc

import (
	"os/exec"
	"strings"
)

// VERSION of webrpc tooling and webrpc-gen Template Functions API.
// Available as {{.WebrpcGenVersion}} variable in Go templates.
//
// The value is injected during `go build' in the release CI step.
var VERSION = ""

func init() {
	if VERSION == "" {
		VERSION = getRuntimeVersion()
	}
}

// getRuntimeVersion tries to infer webrpc version
//  1. from the current go.mod file, which is useful when running webrpc-gen from
//     another Go module using `go run github.com/webrpc/webrpc/cmd/webrpc-gen'.
//  2. from the current git history.
func getRuntimeVersion() string {
	// $ go list -m github.com/webrpc/webrpc
	// github.com/webrpc/webrpc v0.15.1\n
	if out, _ := exec.Command("go", "list", "-m", "github.com/webrpc/webrpc").Output(); len(out) > 0 {
		parts := strings.Split(strings.TrimSpace(string(out)), " ")
		if len(parts) >= 2 {
			return parts[1]
		}
	}

	// $ git describe
	// v0.15.1-6-g550333d\n
	if out, _ := exec.Command("git", "describe").Output(); len(out) > 0 {
		return strings.TrimSpace(string(out))
	}

	return "unknown"
}
