package javascript

import (
	"errors"
	"fmt"
	"strings"

	"github.com/webrpc/webrpc/gen"
	"github.com/webrpc/webrpc/schema"
)

var fieldTypeMap = map[schema.DataType]string{
	schema.T_Uint8:     "number",
	schema.T_Uint16:    "number",
	schema.T_Uint32:    "number",
	schema.T_Uint64:    "number",
	schema.T_Int8:      "number",
	schema.T_Int16:     "number",
	schema.T_Int32:     "number",
	schema.T_Int64:     "number",
	schema.T_Float32:   "number",
	schema.T_Float64:   "number",
	schema.T_String:    "string",
	schema.T_Timestamp: "string",
	schema.T_Null:      "null",
	schema.T_Any:       "any",
	schema.T_Byte:      "string",
	schema.T_Bool:      "boolean",
}

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

func methodName(in interface{}) string {
	v, _ := downcaseName(in)
	return v
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

	if len(in) > 0 {
		inputs = append(inputs, fmt.Sprintf("args"))
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

func downcaseName(v interface{}) (string, error) {
	downFn := func(s string) string {
		if s == "" {
			return ""
		}
		return strings.ToLower(s[0:1]) + s[1:]
	}
	switch t := v.(type) {
	case schema.VarName:
		return downFn(string(t)), nil
	case string:
		return downFn(t), nil
	default:
		return "", errors.New("downcaseFieldName, unknown arg type")
	}
}

func listComma(item int, count int) string {
	if item+1 < count {
		return ", "
	}
	return ""
}

func newOutputArgResponse(in *schema.MethodArgument) (string, error) {
	z, err := fieldConcreteType(in.Type)
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

func exportKeyword(opts gen.TargetOptions) func() string {
	return func() string {
		if opts.Extra == "noexports" {
			return ""
		} else {
			return "export "
		}
	}
}

func serverServiceName(in schema.VarName) (string, error) {
	s := string(in)
	return strings.ToLower(s[0:1]) + s[1:] + "Server", nil
}

func jsFieldType(in *schema.VarType) (string, error) {
	switch in.Type {
	case schema.T_Map:
		return "object", nil

	case schema.T_Struct:
		return in.Struct.Name, nil

	default:
		if fieldTypeMap[in.Type] != "" {
			return fieldTypeMap[in.Type], nil
		}
	}

	return "", fmt.Errorf("could not represent type: %#v", in)
}

func serviceInterfaceName(in schema.VarName) (string, error) {
	s := string(in)
	return s, nil
}

func templateFuncMap(opts gen.TargetOptions) map[string]interface{} {
	return map[string]interface{}{
		"constPathPrefix":      constPathPrefix,
		"methodName":           methodName,
		"methodInputs":         methodInputs,
		"isStruct":             isStruct,
		"isEnum":               isEnum,
		"listComma":            listComma,
		"exportedField":        exportedField,
		"exportedJSONField":    exportedJSONField,
		"newOutputArgResponse": newOutputArgResponse,
		"exportKeyword":        exportKeyword(opts),
		"serverServiceName":    serverServiceName,
		"jsFieldType":          jsFieldType,
		"serviceInterfaceName": serviceInterfaceName,
	}
}
