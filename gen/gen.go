package gen

import (
	"bytes"
	"context"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/posener/gitfs"
	"github.com/shurcooL/httpfs/text/vfstemplate"
	"github.com/webrpc/webrpc/schema"
)

type TargetOptions struct {
	PkgName   string
	Client    bool
	Server    bool
	Extra     string
	Websocket bool
}

func Generate(proto *schema.WebRPCSchema, target string, opts TargetOptions) (string, error) {
	target = getBuiltInTarget(target)

	var err error
	tmpl := template.New("webrpc-gen").Funcs(templateFuncMap(proto, opts))

	// Load templates
	if isLocalDir(target) {
		// from local directory
		tmpl, err = tmpl.ParseGlob("*.tmpl")
		if err != nil {
			return "", errors.Wrapf(err, "failed to load templates from %s", target)
		}
	} else {
		// from remote git directory
		remoteFS, err := gitfs.New(context.Background(), target)
		if err != nil {
			return "", errors.Wrapf(err, "failed to load templates from remote git repository %s", target)
		}

		tmpl, err = vfstemplate.ParseGlob(remoteFS, tmpl, "/*.tmpl")
		if err != nil {
			return "", errors.Wrap(err, "failed to parse Go templates")
		}
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

	return string(buf.Bytes()), nil
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

func isLocalDir(target string) bool {
	return strings.HasPrefix(target, "/") || strings.HasPrefix(target, ".")
}
