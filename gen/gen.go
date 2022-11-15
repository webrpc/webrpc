package gen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
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
		getWebrpcGenCommand(),
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

func getWebrpcGenCommand() string {
	cmd := filepath.Base(os.Args[0])
	if len(os.Args) > 1 {
		cmd = fmt.Sprintf("%s %s", cmd, strings.Join(os.Args[1:], " "))
	}
	return cmd
}
