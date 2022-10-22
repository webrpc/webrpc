package gen

import (
	"bytes"
	"net/http"
	"text/template"

	"github.com/pkg/errors"
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

func Generate(proto *schema.WebRPCSchema, templatesFS http.FileSystem, opts TargetOptions) (string, error) {
	// Load templates
	tmpl := template.New("webrpc-gen").Funcs(templateFuncMap(proto, opts))
	tmpl, err := vfstemplate.ParseGlob(templatesFS, tmpl, "/*.tmpl")
	if err != nil {
		return "", errors.Wrap(err, "failed to parse Go templates")
	}

	// generate deterministic schema hash of the proto file
	schemaHash, err := proto.SchemaHash()
	if err != nil {
		return "", err
	}

	// template vars
	vars := struct {
		*schema.WebRPCSchema
		SchemaHash string
		TargetOpts TargetOptions
	}{
		proto, schemaHash, opts,
	}

	// generate the template
	genBuf := bytes.NewBuffer(nil)
	err = tmpl.ExecuteTemplate(genBuf, "proto", vars)
	if err != nil {
		return "", err
	}

	return string(genBuf.Bytes()), nil
}
