package gen

import (
	"bytes"
	"path/filepath"

	"github.com/webrpc/webrpc/schema"
)

type TargetOptions struct {
	PkgName      string
	Client       bool
	Server       bool
	Extra        string
	OutFilename  string
	RefreshCache bool
}

func Generate(proto *schema.WebRPCSchema, target string, opts TargetOptions) (string, error) {
	target = getBuiltInTarget(target)

	tmpl, err := loadTemplates(proto, target, opts)
	if err != nil {
		return "", err
	}

	// Generate deterministic schema hash of the proto file
	schemaHash, err := proto.SchemaHash()
	if err != nil {
		return "", err
	}

	// Template vars
	vars := struct {
		*schema.WebRPCSchema
		SchemaHash       string
		WebrpcGenVersion string
		Opts             map[string]interface{}
		Cmd              string
	}{
		proto,
		schemaHash,
		VERSION,
		opts,
		strings.Join(os.Args, " "),
	}

	// Generate the template
	var b bytes.Buffer
	err = tmpl.ExecuteTemplate(&b, "proto", vars)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

// Backward compatibility with webrpc-gen v0.6.0.
func getBuiltInTarget(target string) string {
	switch target {
	case "go":
		return "github.com/webrpc/gen-golang@v0.6.0"
	case "ts":
		return "github.com/webrpc/gen-typescript@v0.6.0"
	case "js":
		return "github.com/webrpc/gen-javascript@v0.6.0"
	}
	return target
}
