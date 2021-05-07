package golang

import (
	"bytes"
	"embed"
	"text/template"

	"github.com/pkg/errors"
	"github.com/webrpc/webrpc/gen"
	"github.com/webrpc/webrpc/schema"
)

func init() {
	gen.Register("go", &generator{})
}

//go:embed templates/*
var templatesFS embed.FS

type generator struct{}

func (g *generator) Gen(proto *schema.WebRPCSchema, opts gen.TargetOptions) (string, error) {
	// Load templates
	tmpl, err := template.
		New("webrpc-gen-go").
		Funcs(templateFuncMap(proto)).
		ParseFS(templatesFS, "*.go.tmpl")
	if err != nil {
		return "", errors.Wrap(err, "failed to parse golang templates")
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
		TargetOpts gen.TargetOptions
	}{
		proto, schemaHash, opts,
	}

	// generate the template
	genBuf := bytes.NewBuffer(nil)
	err = tmpl.ExecuteTemplate(genBuf, "proto", vars)
	if err != nil {
		return "", err
	}

	// return string(genBuf.Bytes()), nil

	src, err := FormatSource(genBuf.Bytes())
	if err != nil {
		return "", errors.Errorf("gofmt is failing to format the Go code because: %v", err)
	}

	return string(src), nil
}
