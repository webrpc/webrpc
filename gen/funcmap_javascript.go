package gen

import (
	"fmt"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

func jsExportKeyword(opts TargetOptions) func() string {
	return func() string {
		if opts.Extra == "noexports" {
			return ""
		} else {
			return "export "
		}
	}
}

func jsMethodName(in interface{}) string {
	v, _ := downcaseName(in)
	return v
}

func jsMethodInputs(in []*schema.MethodArgument) (string, error) {
	inputs := []string{}

	if len(in) > 0 {
		inputs = append(inputs, fmt.Sprintf("args"))
	}

	inputs = append(inputs, "headers")

	return strings.Join(inputs, ", "), nil
}

func jsNewOutputArgResponse(in *schema.MethodArgument) (string, error) {
	z, err := jsFieldConcreteType(in.Type)
	if err != nil {
		return "", err
	}

	typ := ""
	switch in.Type.Type {
	case schema.T_Struct:
		typ = fmt.Sprintf("new %s", z)
	default:
		typ = ""
		// 	typ = fmt.Sprintf("<%s>", z)
	}

	line := fmt.Sprintf("%s: %s(_data.%s)", in.Name, typ, in.Name)

	return line, nil
}

func jsFieldConcreteType(in *schema.VarType) (string, error) {
	switch in.Type {
	case schema.T_Struct:
		return in.Struct.Name, nil
	default:
		return "", nil
	}
	return "", fmt.Errorf("could not represent type: %#v", in)
}

func jsServiceInterfaceName(in schema.VarName) (string, error) {
	s := string(in)
	return s, nil
}

func jsExportedJSONField(in schema.MessageField) (string, error) {
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
