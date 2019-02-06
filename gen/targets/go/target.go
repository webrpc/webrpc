package gen

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/webrpc/webrpc/schema"
)

func serviceMethodName(in schema.VarName) (string, error) {
	s := string(in)
	return strings.ToLower(s[0:1]) + s[1:], nil
}

func fieldTags(in []schema.MessageFieldMeta) (string, error) {
	if len(in) < 1 {
		return "", nil
	}
	fieldTags := []string{}
	for kk := range in {
		for k := range in[kk] {
			fieldTags = append(fieldTags, fmt.Sprintf(`%s:"%v"`, k, in[kk][k]))
		}
	}
	return "`" + strings.Join(fieldTags, " ") + "`", nil
}

func newClientServiceName(in schema.VarName) (string, error) {
	return "New" + string(in) + "Client", nil
}

func fieldType(in schema.VarType) (string, error) {
	s := in.String()
	return s, nil
}

func constPathPrefix(in schema.VarName) (string, error) {
	return string(in) + "PathPrefix", nil
}

func countMethods(in []*schema.Method) (string, error) {
	return strconv.Itoa(len(in)), nil
}

func clientServiceName(in schema.VarName) (string, error) {
	s := string(in)
	return strings.ToLower(s[0:1]) + s[1:] + "Client", nil
}

func methodInputName(in *schema.MethodArgument) string {
	name := string(in.Name)
	if name != "" {
		return name
	}
	if in.Type != nil {
		return in.Type.String()
	}
	return ""
}

func methodInputType(in *schema.MethodArgument) string {
	// TODO: how do we know when we have a pointer?
	return in.Type.String()
}

func methodInputs(in []*schema.MethodArgument) (string, error) {
	inputs := []string{"ctx context.Context"}

	for i := range in {
		inputs = append(inputs, fmt.Sprintf("%s %s", methodInputName(in[i]), methodInputType(in[i])))
	}

	return strings.Join(inputs, ", "), nil
}

func methodOutputs(in []*schema.MethodArgument) (string, error) {
	return "", nil
}

func isStruct(t schema.MessageType) bool {
	return t == "struct"
}

var templateFuncMap = map[string]interface{}{
	"serviceMethodName":    serviceMethodName,
	"fieldTags":            fieldTags,
	"fieldType":            fieldType,
	"newClientServiceName": newClientServiceName,
	"constPathPrefix":      constPathPrefix,
	"countMethods":         countMethods,
	"clientServiceName":    clientServiceName,
	"methodInputs":         methodInputs,
	"methodOutputs":        methodOutputs,
	"isStruct":             isStruct,
}

func compile(s *schema.WebRPCSchema) ([]byte, error) {
	tpl, err := template.New("").
		Funcs(templateFuncMap).
		Parse(goTemplate)

	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	if err := tpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
