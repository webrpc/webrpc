package typescript

import (
	"errors"
	"fmt"
	"strings"

	"gitlab.com/logement3d/webrpc/schema"
)

var fieldTypeMap = map[schema.DataType]string{
	schema.T_Uint:      "number",
	schema.T_Uint8:     "number",
	schema.T_Uint16:    "number",
	schema.T_Uint32:    "number",
	schema.T_Uint64:    "number",
	schema.T_Int:       "number",
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

func jsFieldType(in *schema.VarType) (string, error) {
	switch in.Type {
	case schema.T_Map:
		return "object", nil

	case schema.T_List:
		z, err := fieldType(in.List.Elem)

		if err != nil {
			return "", err
		}

		return z + "[]", nil

	case schema.T_Struct:
		return in.Struct.Name, nil

	default:
		if fieldTypeMap[in.Type] != "" {
			return fieldTypeMap[in.Type], nil
		}
	}

	return "", fmt.Errorf("could not represent type: %#v", in)
}

func fieldType(in *schema.VarType) (string, error) {
	switch in.Type {
	case schema.T_Map:
		typK, ok := fieldTypeMap[in.Map.Key]
		if !ok {
			return "", fmt.Errorf("unknown type mapping %v", in.Map.Key)
		}
		typV, err := fieldType(in.Map.Value)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("{[key: %s]: %s}", typK, typV), nil

	case schema.T_List:
		z, err := fieldType(in.List.Elem)
		if err != nil {
			return "", err
		}
		return "Array<" + z + ">", nil

	case schema.T_Struct:
		return in.Struct.Name, nil

	default:
		if fieldTypeMap[in.Type] != "" {
			return fieldTypeMap[in.Type], nil
		}
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

func methodInputType(in *schema.MethodArgument) string {
	z, _ := fieldType(in.Type)
	return z
}

func methodArgumentInputInterfaceName(in *schema.Method) string {
	if len(in.Service.Schema.Services) == 1 {
		return fmt.Sprintf("%s%s", in.Name, "Args")
	} else {
		return fmt.Sprintf("%s%s%s", in.Service.Name, in.Name, "Args")
	}
}

func methodArgumentOutputInterfaceName(in *schema.Method) string {
	if len(in.Service.Schema.Services) == 1 {
		return fmt.Sprintf("%s%s", in.Name, "Return")
	} else {
		return fmt.Sprintf("%s%s%s", in.Service.Name, in.Name, "Return")
	}
}

func methodInputs(in *schema.Method) (string, error) {
	inputs := []string{}
	if len(in.Inputs) > 0 {
		inputs = append(inputs, fmt.Sprintf("args: %s", methodArgumentInputInterfaceName(in)))
	}
	inputs = append(inputs, "headers?: object")
	return strings.Join(inputs, ", "), nil
}

func methodOutputs(in *schema.Method) (string, error) {
	return fmt.Sprintf("Promise<%s>", methodArgumentOutputInterfaceName(in)), nil
}

func methodName(in interface{}) string {
	v, _ := downcaseName(in)
	return v
}

func isStruct(t schema.MessageType) bool {
	return t == "struct"
}

func exportedField(in schema.VarName) (string, error) {
	return string(in), nil
}

func exportableField(in schema.MessageField) bool {
	for _, meta := range in.Meta {
		for k := range meta {
			if k == "json" {
				if meta[k] == "-" {
					return false
				}
			}
		}
	}
	return true
}

func exportedJSONField(in schema.MessageField) (string, error) {
	for _, meta := range in.Meta {
		for k := range meta {
			if k == "json" {
				s := strings.Split(fmt.Sprintf("%v", meta[k]), ",")
				return s[0], nil
			}
		}
	}
	return string(in.Name), nil
}

func interfaceName(in schema.VarName) (string, error) {
	s := string(in)
	return s, nil
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

func serviceInterfaceName(in schema.VarName) (string, error) {
	s := string(in)
	return s, nil
}

func newOutputArgResponse(in *schema.MethodArgument) (string, error) {
	z, err := fieldType(in.Type)
	if err != nil {
		return "", err
	}

	typ := fmt.Sprintf("<%s>", z)
	line := fmt.Sprintf("%s: %s(_data.%s)", in.Name, typ, in.Name)

	return line, nil
}

func serverServiceName(in schema.VarName) (string, error) {
	s := string(in)
	return strings.ToLower(s[0:1]) + s[1:] + "Server", nil
}

func methodArgType(in *schema.MethodArgument) string {
	z, err := fieldType(in.Type)

	if err != nil {
		panic(err.Error())
	}

	var prefix string
	typ := in.Type.Type

	if in.Optional {
		prefix = "*"
	}
	if typ == schema.T_Struct {
		prefix = "" // noop, as already pointer applied elsewhere
	}
	if typ == schema.T_List || typ == schema.T_Map {
		prefix = ""
	}

	return prefix + z
}

var templateFuncMap = map[string]interface{}{
	"fieldType":                         fieldType,
	"constPathPrefix":                   constPathPrefix,
	"interfaceName":                     interfaceName,
	"methodName":                        methodName,
	"methodInputs":                      methodInputs,
	"methodOutputs":                     methodOutputs,
	"methodArgumentInputInterfaceName":  methodArgumentInputInterfaceName,
	"methodArgumentOutputInterfaceName": methodArgumentOutputInterfaceName,
	"isStruct":                          isStruct,
	"isEnum":                            isEnum,
	"listComma":                         listComma,
	"serviceInterfaceName":              serviceInterfaceName,
	"exportableField":                   exportableField,
	"exportedField":                     exportedField,
	"exportedJSONField":                 exportedJSONField,
	"newOutputArgResponse":              newOutputArgResponse,
	"downcaseName":                      downcaseName,
	"serverServiceName":                 serverServiceName,
	"methodArgType":                     methodArgType,
	"jsFieldType":                       jsFieldType,
}
