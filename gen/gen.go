package gen

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/posener/gitfs"
	"github.com/shurcooL/httpfs/text/vfstemplate"
	"github.com/webrpc/webrpc/schema"
)

type TargetOptions struct {
	PkgName     string
	Client      bool
	Server      bool
	Extra       string
	OutFilename string
}

func Generate(proto *schema.WebRPCSchema, target string, opts TargetOptions) (string, error) {
	target = getBuiltInTarget(target)

	tmpl, err := LoadTemplates(proto, target, opts)
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
		SchemaHash string
		TargetOpts TargetOptions
	}{
		proto, schemaHash, opts,
	}

	// Generate the template
	buf := bytes.NewBuffer(nil)
	err = tmpl.ExecuteTemplate(buf, "proto", vars)
	if err != nil {
		return "", err
	}

	out := buf.Bytes()

	// Auto-format certain extensions
	if filepath.Ext(opts.OutFilename) == ".go" {
		out, _ = FormatGoSource(out)
	}

	return string(out), nil
}

func LoadTemplates(proto *schema.WebRPCSchema, target string, opts TargetOptions) (*template.Template, error) {
	var err error
	tmpl := template.New("webrpc-gen").Funcs(templateFuncMap(proto, opts))

	// Load templates
	if isLocalDir(target) {
		// from local directory
		tmpl, err = tmpl.ParseGlob(filepath.Join(target, "/*.tmpl"))
		if err != nil {
			return nil, fmt.Errorf("failed to load templates from %s: %w", target, err)
		}
	} else {
		// from remote git directory
		remoteFS, err := gitfs.New(context.Background(), inferRemoteTarget(target))
		if err != nil {
			return nil, fmt.Errorf("failed to load templates from remote git repository %s: %w", target, err)
		}

		tmpl, err = vfstemplate.ParseGlob(remoteFS, tmpl, "/*.tmpl")
		if err != nil {
			return nil, fmt.Errorf("failed to parse templates: %w", err)
		}
	}

	return tmpl, nil
}

func isLocalDir(target string) bool {
	return strings.HasPrefix(target, "/") || strings.HasPrefix(target, ".")
}

func inferRemoteTarget(target string) string {
	// extra check to ensure its not a local dir
	if isLocalDir(target) {
		return target
	}

	// determine if a url is passed or just a gen-XXX name
	parts := strings.Split(target, "/")

	// just a name, so by convention assume the default target of the webrpc org
	if len(parts) == 1 {
		return fmt.Sprintf("github.com/webrpc/gen-%s", strings.ToLower(target))
	}

	// accept the target as is
	return target
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
