package gen

import (
	"bytes"
	"text/template"

	"github.com/webrpc/webrpc/schema"
)

// the webrpc-gen tool goes here.
// starting with support for Go server+client and TypeScript client

func Gen(proto *schema.WebRPCSchema) (string, error) {

	vars := struct {
		Name string
	}{
		"joe",
	}

	t, err := template.New("go.client").Parse(GoClientTemplate)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)

	err = t.Execute(buf, vars)
	if err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}
