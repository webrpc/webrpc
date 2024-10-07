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

type TemplateVars struct {
	*schema.WebRPCSchema
	SchemaHash       string
	WebrpcGenVersion string
	WebrpcGenCommand string
	WebrpcGenHeader  string
	WebrpcTarget     string
	WebrpcErrors     []*schema.Error
	Opts             map[string]interface{}
	TmplVersion      string
	TmplTarget       string
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

	vars := TemplateVars{
		WebRPCSchema:     proto,
		SchemaHash:       schemaHash,
		WebrpcGenVersion: webrpc.VERSION,
		WebrpcGenCommand: getWebrpcGenCommand(),
		WebrpcTarget:     target,
		WebrpcErrors:     WebrpcErrors,
		Opts:             config.TemplateOptions,
	}
	if isLocalDir(target) {
		vars.WebrpcTarget = target
	}

	genOutput := &GenOutput{}

	// Built-in targets
	switch target {
	case "json":
		genJSON, err := proto.ToJSON()
		if err != nil {
			return genOutput, err
		}
		genOutput.TmplVersion = target
		genOutput.IsLocal = true
		genOutput.Code = genJSON
		return genOutput, nil

	case "debug":
		debug := spew.NewDefaultConfig()
		debug.DisableMethods = true
		debug.DisablePointerAddresses = true
		debug.Indent = "\t"
		debug.SortKeys = true
		genOutput.TmplVersion = target
		genOutput.IsLocal = true
		genOutput.Code = debug.Sdump(vars)
		return genOutput, nil
	}

	// webrpc-gen v0.6.0
	target = getOldTarget(target)

	tmpl, tmplSource, err := loadTemplates(proto, target, config)
	if err != nil {
		return genOutput, err
	}
	genOutput.TemplateSource = *tmplSource

	v := strings.Split(tmplSource.TmplVersion, "/")
	tmplTarget, tmplVersion, _ := strings.Cut(v[len(v)-1], "@")

	vars.TmplVersion = tmplVersion
	vars.TmplTarget = tmplTarget
	vars.WebrpcGenHeader = webRPCGenHeader(vars)

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

func webRPCGenHeader(t TemplateVars) string {
	webrpcGenVersion := fmt.Sprintf("webrpc@%s", t.WebrpcGenVersion)

	tmplVersion := t.TmplVersion
	if tmplVersion == "" {
		tmplVersion = "unknown"
	}
	tmpl := fmt.Sprintf("%s@%s", t.TmplTarget, tmplVersion)

	schemaVersion := fmt.Sprintf("%s@%s", t.SchemaName, t.SchemaVersion)

	return fmt.Sprintf("%s;%s;%s", webrpcGenVersion, tmpl, schemaVersion)
}
