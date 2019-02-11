package typescript

import (
	"fmt"
	"strings"

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
	case schema.T_Map:
		z, err := fieldType(in.Map.Value)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Map<%v,%s>", in.Map.Key, z), nil
	case schema.T_List:
		z, err := fieldType(in.List.Elem)
		if err != nil {
			return "", err
		}
		return "Array<" + z + ">", nil
	case schema.T_Struct:
		// TODO: add in.Struct.Message to global structs
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
		z, err := fieldType(in.Map.Value)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Map<%v,%s>", in.Map.Key, z), nil
	case schema.T_List:
		z, err := fieldType(in.List.Elem)
		if err != nil {
			return "", err
		}
		return "Array<" + z + ">", nil
	case schema.T_Struct:
		// TODO: add in.Struct.Message to global structs
		return "I" + in.Struct.Name, nil
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

func methodInputs(in []*schema.MethodArgument) (string, error) {
	inputs := []string{}

	for i := range in {
		inputs = append(inputs, fmt.Sprintf("params: %s", methodInputType(in[i])))
	}

	inputs = append(inputs, "headers: object")

	return strings.Join(inputs, ", "), nil
}

func methodOutputs(in []*schema.MethodArgument) (string, error) {
	outputs := []string{}

	for i := range in {
		z, _ := fieldConcreteType(in[i].Type)
		outputs = append(outputs, z)
	}

	return fmt.Sprintf("Promise<%s>", strings.Join(outputs, ", ")), nil
}

func isStruct(t schema.MessageType) bool {
	return t == "struct"
}

func exportedField(in schema.VarName) (string, error) {
	return string(in), nil
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
	return "I" + s, nil
}

func isEnum(t schema.MessageType) bool {
	return t == "enum"
}

func serviceInterfaceName(in schema.VarName) (string, error) {
	s := string(in)
	return "I" + s + "Service", nil
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
	return fmt.Sprintf("<%s>", z), nil
}

var templateFuncMap = map[string]interface{}{
	"fieldType":               fieldType,
	"newResponseConcreteType": newResponseConcreteType,
	"constPathPrefix":         constPathPrefix,
	"interfaceName":           interfaceName,
	"methodInputs":            methodInputs,
	"methodOutputs":           methodOutputs,
	"isStruct":                isStruct,
	"isEnum":                  isEnum,
	"serviceInterfaceName":    serviceInterfaceName,
	"exportedField":           exportedField,
	"exportedJSONField":       exportedJSONField,
}
