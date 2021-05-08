package javascript

import (
	"bytes"
	"embed"
	"text/template"

	"github.com/pkg/errors"
	"github.com/webrpc/webrpc/gen"
	"github.com/webrpc/webrpc/schema"
)

func init() {
	gen.Register("js", &generator{})
}

//go:embed templates/*.js.tmpl
var templatesFS embed.FS

type generator struct{}

func (g *generator) Gen(proto *schema.WebRPCSchema, opts gen.TargetOptions) (string, error) {
	// Load templates
	tmpl, err := template.
		New("webrpc-gen-js").
		Funcs(templateFuncMap(opts)).
		ParseFS(templatesFS, "templates/*.js.tmpl")
	if err != nil {
		return "", errors.Wrap(err, "failed to parse javascript templates")
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

	// Generate the template
	genBuf := bytes.NewBuffer(nil)
	err = tmpl.ExecuteTemplate(genBuf, "proto", vars)
	if err != nil {
		return "", err
	}

	return string(genBuf.Bytes()), nil
}
