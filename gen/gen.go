package gen

import (
	"bytes"
	"os"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

func Generate(proto *schema.WebRPCSchema, target string, refreshCache bool, opts map[string]interface{}) (string, error) {
	target = getOldTarget(target)

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

	fmt, _ := opts["fmt"].(bool)
	delete(opts, "fmt")

	// Generate the template
	var b bytes.Buffer
	err = tmpl.ExecuteTemplate(&b, "main", vars)
	if err != nil {
		return "", err
	}

	if fmt && isGolangTarget(target) {
		return formatGoSource(b.Bytes())
	}

	return b.String(), nil
}
