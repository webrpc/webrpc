package gen

import (
	"bytes"
	"os"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

func Generate(proto *schema.WebRPCSchema, target string, refreshCache bool, opts map[string]interface{}) (string, error) {
	target = getBuiltInTarget(target)

	tmpl, err := loadTemplates(proto, target, refreshCache, opts)
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
		WebrpcGenCommand string
		WebrpcTarget     string
		Opts             map[string]interface{}
	}{
		proto,
		schemaHash,
		VERSION,
		strings.Join(os.Args, " "),
		target,
		opts,
	}
	if isLocalDir(target) {
		vars.WebrpcTarget = "custom"
	}

	// Generate the template
	var b bytes.Buffer
	err = tmpl.ExecuteTemplate(&b, "main", vars)
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
