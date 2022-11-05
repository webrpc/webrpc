package gen

import (
	"fmt"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

var tsFieldTypeMap = map[schema.DataType]string{
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

func tsFieldType(in *schema.VarType) (string, error) {
	switch in.Type {
	case schema.T_Map:
		typK, ok := tsFieldTypeMap[in.Map.Key]
		if !ok {
			return "", fmt.Errorf("unknown type mapping %v", in.Map.Key)
		}
		typV, err := tsFieldType(in.Map.Value)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("{[key: %s]: %s}", typK, typV), nil

	case schema.T_List:
		z, err := tsFieldType(in.List.Elem)
		if err != nil {
			return "", err
		}
		return "Array<" + z + ">", nil

	case schema.T_Struct:
		return in.Struct.Name, nil

	default:
		if tsFieldTypeMap[in.Type] != "" {
			return tsFieldTypeMap[in.Type], nil
		}
	}
	return "", fmt.Errorf("could not represent type: %#v", in)
}

func tsMethodArgumentInputInterfaceName(in *schema.Method) string {
	if len(in.Service.Schema.Services) == 1 {
		return fmt.Sprintf("%s%s", in.Name, "Args")
	} else {
		return fmt.Sprintf("%s%s%s", in.Service.Name, in.Name, "Args")
	}
}

func tsMethodArgumentOutputInterfaceName(in *schema.Method) string {
	if len(in.Service.Schema.Services) == 1 {
		return fmt.Sprintf("%s%s", in.Name, "Return")
	} else {
		return fmt.Sprintf("%s%s%s", in.Service.Name, in.Name, "Return")
	}
}

func tsMethodInputs(in *schema.Method) (string, error) {
	inputs := []string{}
	if len(in.Inputs) > 0 {
		inputs = append(inputs, fmt.Sprintf("args: %s", tsMethodArgumentInputInterfaceName(in)))
	}
	inputs = append(inputs, "headers?: object")
	return strings.Join(inputs, ", "), nil
}

func tsMethodOutputs(in *schema.Method) (string, error) {
	return fmt.Sprintf("Promise<%s>", tsMethodArgumentOutputInterfaceName(in)), nil
}

func tsInterfaceName(in schema.VarName) (string, error) {
	s := string(in)
	return s, nil
}

func tsServiceInterfaceName(in schema.VarName) (string, error) {
	s := string(in)
	return s, nil
}

func tsNewOutputArgResponse(in *schema.MethodArgument) (string, error) {
	z, err := tsFieldType(in.Type)
	if err != nil {
		return "", err
	}

	typ := fmt.Sprintf("<%s>", z)
	line := fmt.Sprintf("%s: %s(_data.%s)", in.Name, typ, in.Name)

	return line, nil
}

func tsMethodArgType(in *schema.MethodArgument) string {
	z, err := tsFieldType(in.Type)

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
