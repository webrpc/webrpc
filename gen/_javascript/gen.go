//go:generate go run github.com/goware/statik -src=./templates -dest=. -f -Z -p=embed
package javascript

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/goware/statik/fs"
	"github.com/webrpc/webrpc/gen"
	"github.com/webrpc/webrpc/gen/javascript/embed"
	"github.com/webrpc/webrpc/schema"
)

func init() {
	gen.Register("js", &generator{})
}

type generator struct{}

func (g *generator) Gen(proto *schema.WebRPCSchema, opts gen.TargetOptions) (string, error) {
	// Get templates from `embed` asset package
	// NOTE: make sure to `go generate` whenever you change the files in `templates/` folder
	templates, err := getTemplates()
	if err != nil {
		return "", err
	}

	// Load templates
	tmpl := template.
		New("webrpc-gen-js").
		Funcs(templateFuncMap(opts))

	for _, tmplData := range templates {
		_, err = tmpl.Parse(tmplData)
		if err != nil {
			return "", err
		}
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

func getTemplates() (map[string]string, error) {
	data := map[string]string{}

	statikFS, err := fs.New(embed.Asset)
	if err != nil {
		return nil, err
	}

	fs.Walk(statikFS, "/", func(path string, info os.FileInfo, err error) error {
		if path == "/" {
			return nil
		}
		f, err := statikFS.Open(path)
		if err != nil {
			return err
		}
		buf, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		data[path] = string(buf)
		return nil
	})

	return data, nil
}
