package typescript

import (
	"fmt"
	"strconv"
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
	schema.T_Bool:      "bool",
}

func serviceMethodName(in schema.VarName) (string, error) {
	s := string(in)
	return "serve" + strings.ToUpper(s[0:1]) + s[1:], nil
}

func serviceMethodJSONName(in schema.VarName) (string, error) {
	s := string(in)
	return "serve" + strings.ToUpper(s[0:1]) + s[1:] + "JSON", nil
}

func fieldTags(in []schema.MessageFieldMeta) (string, error) {
	if len(in) < 1 {
		return "", nil
	}
	fieldTags := []string{}
	for kk := range in {
		for k := range in[kk] {
			switch {
			case strings.HasPrefix(k, "go.tag."):
				fieldTags = append(fieldTags, fmt.Sprintf(`%s:"%v"`, k[7:], in[kk][k]))
			case k == "json":
				fieldTags = append(fieldTags, fmt.Sprintf(`json:"%v"`, in[kk][k]))
			}
		}
	}
	return "`" + strings.Join(fieldTags, " ") + "`", nil
}

func newServerServiceName(in schema.VarName) (string, error) {
	return "New" + string(in) + "Server", nil
}

func newClientServiceName(in schema.VarName) (string, error) {
	return "New" + string(in) + "Client", nil
}

func fieldConcreteType(in *schema.VarType) (string, error) {
	switch in.Type {
	case schema.T_Map:
		z, err := fieldType(in.Map.Value)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("map[%v]%s", in.Map.Key, z), nil
	case schema.T_List:
		z, err := fieldType(in.List.Elem)
		if err != nil {
			return "", err
		}
		return "[]" + z, nil
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
		return fmt.Sprintf("map[%v]%s", in.Map.Key, z), nil
	case schema.T_List:
		z, err := fieldType(in.List.Elem)
		if err != nil {
			return "", err
		}
		return "[]" + z, nil
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

func countMethods(in []*schema.Method) (string, error) {
	return strconv.Itoa(len(in)), nil
}

func clientServiceName(in schema.VarName) (string, error) {
	s := string(in)
	return strings.ToLower(s[0:1]) + s[1:] + "Client", nil
}

func serverServiceName(in schema.VarName) (string, error) {
	s := string(in)
	return strings.ToLower(s[0:1]) + s[1:] + "Server", nil
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

	inputs = append(inputs, "headers: object = {}")

	return strings.Join(inputs, ", "), nil
}

func methodOutputs(in []*schema.MethodArgument) (string, error) {
	outputs := []string{}

	for i := range in {
		outputs = append(outputs, methodInputType(in[i]))
	}

	return fmt.Sprintf("Promise<%s>", strings.Join(outputs, ", ")), nil
}

func isStruct(t schema.MessageType) bool {
	return t == "struct"
}

func exportedField(in schema.VarName) (string, error) {
	s := string(in)
	return strings.ToUpper(s[0:1]) + s[1:], nil
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

func optional(in bool) (string, error) {
	if in {
		return "?", nil
	}
	return "", nil
}

func newResponseConcreteType(in schema.MethodArgument) (string, error) {
	z, err := fieldConcreteType(in.Type)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("new %s", z), nil
}

var templateFuncMap = map[string]interface{}{
	"serviceMethodName":       serviceMethodName,
	"serviceMethodJSONName":   serviceMethodJSONName,
	"fieldTags":               fieldTags,
	"fieldType":               fieldType,
	"newClientServiceName":    newClientServiceName,
	"newServerServiceName":    newServerServiceName,
	"newResponseConcreteType": newResponseConcreteType,
	"constPathPrefix":         constPathPrefix,
	"countMethods":            countMethods,
	"clientServiceName":       clientServiceName,
	"interfaceName":           interfaceName,
	"serverServiceName":       serverServiceName,
	"methodInputs":            methodInputs,
	"methodOutputs":           methodOutputs,
	"isStruct":                isStruct,
	"isEnum":                  isEnum,
	"optional":                optional,
	"serviceInterfaceName":    serviceInterfaceName,
	"exportedField":           exportedField,
}
