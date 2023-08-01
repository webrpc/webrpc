package gen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/webrpc/webrpc"
	"github.com/webrpc/webrpc/schema"
)

type Config struct {
	RefreshCache    bool
	Format          bool
	TemplateOptions map[string]interface{}
}

type GenOutput struct {
	Code string
	TemplateSource
	FormatErr error
}

func Generate(proto *schema.WebRPCSchema, target string, config *Config) (out *GenOutput, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%v\n\tcommand failed: %w", getWebrpcGenCommand(), err)
		}
	}()

	// Generate deterministic schema hash of the proto file
	schemaHash, err := proto.SchemaHash()
	if err != nil {
		return nil, err
	}

	// Template vars
	vars := struct {
		*schema.WebRPCSchema
		SchemaHash       string
		WebrpcGenVersion string
		WebrpcGenCommand string
		WebrpcTarget     string
		WebrpcErrors     []*schema.Error
		Opts             map[string]interface{}
	}{
		proto,
		schemaHash,
		webrpc.VERSION,
		getWebrpcGenCommand(),
		target,
		WebrpcErrors,
		config.TemplateOptions,
	}
	if isLocalDir(target) {
		vars.WebrpcTarget = "custom"
	}

	genOutput := &GenOutput{}

	// Built-in targets
	switch target {
	case "json":
		genJSON, err := proto.ToJSON()
		if err != nil {
			return genOutput, err
		}
		genOutput.Code = genJSON
		return genOutput, nil

	case "debug":
		genOutput.Code = spew.Sdump(vars)
		return genOutput, nil
	}

	// webrpc-gen v0.6.0
	target = getOldTarget(target)

	tmpl, tmplSource, err := loadTemplates(proto, target, config)
	if err != nil {
		return genOutput, err
	}
	genOutput.TemplateSource = *tmplSource

	// Generate the template
	var b bytes.Buffer
	err = tmpl.ExecuteTemplate(&b, "main", vars)
	if err != nil {
		return genOutput, err
	}

	if config.Format && isGolangTarget(target) {
		genOutput.Code, genOutput.FormatErr = formatGoSource(b.Bytes())
	} else {
		genOutput.Code = b.String()
	}

	return genOutput, nil
}

func getWebrpcGenCommand() string {
	cmd := filepath.Base(os.Args[0])
	if len(os.Args) > 1 {
		cmd = fmt.Sprintf("%s %s", cmd, strings.Join(os.Args[1:], " "))
	}
	return cmd
}
