package gen

import (
	"bytes"
	"fmt"
	"go/format"
	"strconv"
	"strings"
	"text/template"

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
	schema.T_Timestamp: "*time.Time",
	schema.T_Null:      "struct{}",
	schema.T_Any:       "interface{}",
	schema.T_Byte:      "byte",
	schema.T_Bool:      "bool",
}

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

func exportedField(in schema.VarName) (string, error) {
	s := string(in)
	return strings.ToUpper(s[0:1]) + s[1:], nil
}

func isEnum(t schema.MessageType) bool {
	return t == "enum"
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
	"isEnum":               isEnum,
	"exportedField":        exportedField,
}

func compileTemplate(src string, s *schema.WebRPCSchema) ([]byte, error) {
	tpl, err := template.New("").
		Funcs(templateFuncMap).
		Parse(src)

	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	if err := tpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func compile(s *schema.WebRPCSchema) ([]byte, error) {
	buf, err := compileTemplate(goTemplate, s)
	if err != nil {
		return nil, err
	}
	return format.Source(buf)
}
