package javascript

import (
	"fmt"
	"strings"

	"github.com/webrpc/webrpc/gen"
	"github.com/webrpc/webrpc/schema"
)

func fieldConcreteType(in *schema.VarType) (string, error) {
	switch in.Type {
	case schema.T_Struct:
		return in.Struct.Name, nil
	default:
		return "", nil
	}
	return "", fmt.Errorf("could not represent type: %#v", in)
}

func constPathPrefix(in schema.VarName) (string, error) {
	return string(in) + "PathPrefix", nil
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

func methodInputs(in []*schema.MethodArgument) (string, error) {
	inputs := []string{}

	for i := range in {
		_ = i
		inputs = append(inputs, fmt.Sprintf("params"))
	}

	inputs = append(inputs, "headers")

	return strings.Join(inputs, ", "), nil
}

func isStruct(t schema.MessageType) bool {
	return t == "struct"
}

func exportedField(in schema.VarName) (string, error) {
	s := string(in)
	return strings.ToUpper(s[0:1]) + s[1:], nil
}

func exportedJSONField(in schema.MessageField) (string, error) {
	for i := range in.Meta {
		for k := range in.Meta[i] {
			if k == "json" {
				s := strings.Split(fmt.Sprintf("%v", in.Meta[i][k]), ",")
				if len(s) > 0 {
					return s[0], nil
				}
			}
		}
	}
	return string(in.Name), nil
}

func isEnum(t schema.MessageType) bool {
	return t == "enum"
}

func newResponseConcreteType(in schema.MethodArgument) (string, error) {
	z, err := fieldConcreteType(in.Type)
	if err != nil {
		return "", err
	}

	switch in.Type.Type {
	case schema.T_Struct:
		return fmt.Sprintf("new %s", z), nil
	}
	return "", nil
}

func exportKeyword(opts gen.TargetOptions) func() string {
	return func() string {
		if opts.Extra == "noexports" {
			return ""
		} else {
			return "export "
		}
	}
}

func templateFuncMap(opts gen.TargetOptions) map[string]interface{} {
	return map[string]interface{}{
		"newResponseConcreteType": newResponseConcreteType,
		"constPathPrefix":         constPathPrefix,
		"methodInputs":            methodInputs,
		"isStruct":                isStruct,
		"isEnum":                  isEnum,
		"exportedField":           exportedField,
		"exportedJSONField":       exportedJSONField,
		"exportKeyword":           exportKeyword(opts),
	}
}
