package webrpc

import (
	"os/exec"
	"strings"
)

// VERSION of webrpc tooling and webrpc-gen Template Functions API.
// Available as {{.WebrpcGenVersion}} variable in Go templates.
//
// The actual value is injected into the binary in the release CI step.
var VERSION = ""

func init() {
	if VERSION == "" {
		VERSION = getVersion()
	}
}

// getVersion tries to get version of the github.com/webrpc/webrpc module from
// current go.mod file. This is useful when running webrpc-gen from within
// another Go module using `go run github.com/webrpc/webrpc/cmd/webrpc-gen'.
func getVersion() string {
	out, err := exec.Command("go", "list", "-m", "github.com/webrpc/webrpc").Output()
	if err != nil {
		return "development"
	}

	// $ go list -m github.com/webrpc/webrpc
	// github.com/webrpc/webrpc v0.13.1\n
	parts := strings.Split(string(out), " ")
	if len(parts) < 2 {
		return "development"
	}

	// v0.13.1
	return parts[1]
}
