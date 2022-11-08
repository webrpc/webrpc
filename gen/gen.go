package gen

import (
	"bytes"
	"os"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

type Config struct {
	RefreshCache    bool
	Format          bool
	TemplateOptions map[string]interface{}
}

func Generate(proto *schema.WebRPCSchema, target string, config *Config) (string, error) {
	target = getOldTarget(target)

	tmpl, err := loadTemplates(proto, target, config)
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
		config.TemplateOptions,
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

	if config.Format && isGolangTarget(target) {
		return formatGoSource(b.Bytes())
	}

	return b.String(), nil
}
