package golang

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

var fieldTypeMap = map[schema.DataType]string{
	schema.T_Uint8:     "uint8",
	schema.T_Uint16:    "uint16",
	schema.T_Uint32:    "uint32",
	schema.T_Uint64:    "uint64",
	schema.T_Int8:      "int8",
	schema.T_Int16:     "int16",
	schema.T_Int32:     "int32",
	schema.T_Int64:     "int64",
	schema.T_Float32:   "float32",
	schema.T_Float64:   "float64",
	schema.T_String:    "string",
	schema.T_Timestamp: "time.Time",
	schema.T_Null:      "struct{}",
	schema.T_Any:       "interface{}",
	schema.T_Byte:      "byte",
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
		return in.Struct.Name, nil
	default:
		if fieldTypeMap[in.Type] != "" {
			return fieldTypeMap[in.Type], nil
		}
	}
	return "", fmt.Errorf("could not represent type: %#v", in)
}

func fieldOptional(field *schema.MessageField) (string, error) {
	if !field.Optional {
		return "", nil
	}
	switch field.Type.Type {
	case schema.T_Map:
		return "", nil // noop
	case schema.T_List:
		return "", nil // noop
	case schema.T_Struct:
		return "*", nil
	default:
		if fieldTypeMap[field.Type.Type] != "" {
			return "*", nil
		}
	}
	return "", fmt.Errorf("could not represent type: %#v", field)
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
	if name == "" && in.Type != nil {
		name = in.Type.String()
	}
	if name != "" {
		return strings.ToLower(name[0:1]) + name[1:]
	}
	return ""
}

func methodInputType(in *schema.MethodArgument) string {
	// TODO: how do we know when we have a pointer?
	z, err := fieldType(in.Type)
	if err != nil {
		panic(err.Error())
	}
	return "*" + z
}

func methodInputs(in []*schema.MethodArgument) (string, error) {
	inputs := []string{"ctx context.Context"}

	for i := range in {
		inputs = append(inputs, fmt.Sprintf("%s %s", methodInputName(in[i]), methodInputType(in[i])))
	}

	return strings.Join(inputs, ", "), nil
}

// func methodInterfaceInputs(in []*schema.MethodArgument) (string, error) {
// 	inputs := []string{"ctx context.Context"}

// 	for i := range in {
// 		inputs = append(inputs, fmt.Sprintf("%s", methodInputType(in[i])))
// 	}

// 	return strings.Join(inputs, ", "), nil
// }

func methodOutputs(in []*schema.MethodArgument) (string, error) {
	outputs := []string{}

	for i := range in {
		outputs = append(outputs, methodInputType(in[i]))
	}

	outputs = append(outputs, "error")

	return strings.Join(outputs, ", "), nil
}

func isStruct(t schema.MessageType) bool {
	return t == "struct"
}

func exportedField(in schema.VarName) (string, error) {
	s := string(in)
	return strings.ToUpper(s[0:1]) + s[1:], nil
}

func isEnum(t schema.MessageType) bool {
	return t == "enum"
}

var templateFuncMap = map[string]interface{}{
	"serviceMethodName":     serviceMethodName,
	"serviceMethodJSONName": serviceMethodJSONName,
	"fieldTags":             fieldTags,
	"fieldType":             fieldType,
	"fieldOptional":         fieldOptional,
	"newClientServiceName":  newClientServiceName,
	"newServerServiceName":  newServerServiceName,
	"constPathPrefix":       constPathPrefix,
	"countMethods":          countMethods,
	"clientServiceName":     clientServiceName,
	"serverServiceName":     serverServiceName,
	// "methodInterfaceInputs": methodInterfaceInputs,
	"methodInputs":    methodInputs,
	"methodOutputs":   methodOutputs,
	"methodInputName": methodInputName,
	"isStruct":        isStruct,
	"isEnum":          isEnum,
	"exportedField":   exportedField,
}
