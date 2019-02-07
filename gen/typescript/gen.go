//go:generate statik -src=./templates -dest=. -f -Z -p=embed
package typescript

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/rakyll/statik/fs"
	"github.com/webrpc/webrpc/gen"
	_ "github.com/webrpc/webrpc/gen/typescript/embed"
	"github.com/webrpc/webrpc/schema"
)

func init() {
	gen.Register("ts", &generator{})
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
		New("webrpc-gen-ts").
		Funcs(templateFuncMap)
	for _, tmplData := range templates {
		_, err = tmpl.Parse(tmplData)
		if err != nil {
			return "", err
		}
	}

	// Generate the template
	genBuf := bytes.NewBuffer(nil)
	err = tmpl.ExecuteTemplate(genBuf, "proto", proto)
	if err != nil {
		return "", err
	}

	return string(genBuf.Bytes()), nil
}

func getTemplates() (map[string]string, error) {
	data := map[string]string{}

	statikFS, err := fs.New()
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
