package golang

import (
	"errors"
	"fmt"
	"sort"
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
		return "*" + in.Struct.Name, nil
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
		return "", nil // noop because by default struct uses '*' prefix
	default:
		if fieldTypeMap[field.Type.Type] != "" {
			return "*", nil
		}
	}
	return "", fmt.Errorf("could not represent type: %#v", field)
}

func fieldTypeDef(in *schema.MessageField) (string, error) {
	goFieldType := ""

	meta := in.Meta
	for kk := range meta {
		for k, v := range meta[kk] {
			if k == "go.field.type" {
				goFieldType = fmt.Sprintf("%v", v)
			}
		}
	}

	if goFieldType != "" {
		return goFieldType, nil
	}

	return fieldType(in.Type)
}

func fieldTags(in *schema.MessageField) (string, error) {
	fieldTags := map[string]interface{}{}

	jsonFieldName, err := downcaseName(in.Name)
	if err != nil {
		return "", err
	}
	fieldTags["json"] = fmt.Sprintf("%s", jsonFieldName)

	goTagJSON := ""

	meta := in.Meta
	for kk := range meta {
		for k, v := range meta[kk] {

			switch {
			case k == "json":
				if goTagJSON == "" {
					fieldTags["json"] = fmt.Sprintf("%v", v)
				}

			case strings.HasPrefix(k, "go.tag.json"):
				goTagJSON = fmt.Sprintf("%v", v)
				if !strings.HasPrefix(goTagJSON, fmt.Sprintf("%v", fieldTags["json"])) {
					return "", errors.New("go.tag.json is invalid, it must match the json fieldname")
				}
				fieldTags[k[7:]] = fmt.Sprintf("%v", v)

			case strings.HasPrefix(k, "go.tag."):
				if k == "go.tag.json" {
					goTagJSON = fmt.Sprintf("%v", v)
				}
				fieldTags[k[7:]] = fmt.Sprintf("%v", v)
			}

		}
	}

	tagKeys := []string{}
	for k, _ := range fieldTags {
		if k != "json" {
			tagKeys = append(tagKeys, k)
		}
	}
	sort.StringSlice(tagKeys).Sort()
	tagKeys = append([]string{"json"}, tagKeys...)

	tags := []string{}
	for _, k := range tagKeys {
		tags = append(tags, fmt.Sprintf(`%s:"%v"`, k, fieldTags[k]))
	}

	return "`" + strings.Join(tags, " ") + "`", nil
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

func methodArgName(in *schema.MethodArgument) string {
	name := string(in.Name)
	if name == "" && in.Type != nil {
		name = in.Type.String()
	}
	if name != "" {
		return name
	}
	return ""
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

func methodInputs(in []*schema.MethodArgument) (string, error) {
	inputs := []string{"ctx context.Context"}
	for i := range in {
		inputs = append(inputs, fmt.Sprintf("%s %s", methodArgName(in[i]), methodArgType(in[i])))
	}
	return strings.Join(inputs, ", "), nil
}

func methodOutputs(in []*schema.MethodArgument) (string, error) {
	outputs := []string{}
	for i := range in {
		outputs = append(outputs, methodArgType(in[i]))
	}
	outputs = append(outputs, "error")
	return strings.Join(outputs, ", "), nil
}

func methodArgNames(in []*schema.MethodArgument) (string, error) {
	inputs := []string{}
	for i := range in {
		inputs = append(inputs, fmt.Sprintf("%s", methodArgName(in[i])))
	}
	return strings.Join(inputs, ", "), nil
}

func argsList(in []*schema.MethodArgument, prefix string) (string, error) {
	ins := []string{}
	for i := range in {
		ins = append(ins, fmt.Sprintf("%s%d", prefix, i))
	}
	return strings.Join(ins, ", "), nil
}

func commaIfLen(in []*schema.MethodArgument) string {
	if len(in) > 0 {
		return ","
	}
	return ""
}

func isStruct(t schema.MessageType) bool {
	return t == "struct"
}

func exportedField(in *schema.MessageField) (string, error) {
	s := string(in.Name)
	s = strings.ToUpper(s[0:1]) + s[1:]

	nameTag := "go.field.name"
	for k := range in.Meta {
		for k, v := range in.Meta[k] {
			if k == nameTag {
				s = fmt.Sprintf("%v", v)
			}
		}
	}

	return s, nil
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

func isEnum(t schema.MessageType) bool {
	return t == "enum"
}

func hasFieldType(proto *schema.WebRPCSchema) func(fieldType string) (bool, error) {
	return func(fieldType string) (bool, error) {
		return proto.HasFieldType(fieldType)
	}
}

func templateFuncMap(proto *schema.WebRPCSchema) map[string]interface{} {
	return map[string]interface{}{
		"serviceMethodName":     serviceMethodName,
		"serviceMethodJSONName": serviceMethodJSONName,
		"hasFieldType":          hasFieldType(proto),
		"fieldTags":             fieldTags,
		"fieldType":             fieldType,
		"fieldOptional":         fieldOptional,
		"fieldTypeDef":          fieldTypeDef,
		"newClientServiceName":  newClientServiceName,
		"newServerServiceName":  newServerServiceName,
		"constPathPrefix":       constPathPrefix,
		"countMethods":          countMethods,
		"clientServiceName":     clientServiceName,
		"serverServiceName":     serverServiceName,
		"methodInputs":          methodInputs,
		"methodOutputs":         methodOutputs,
		"methodArgName":         methodArgName,
		"methodArgType":         methodArgType,
		"methodArgNames":        methodArgNames,
		"argsList":              argsList,
		"commaIfLen":            commaIfLen,
		"isStruct":              isStruct,
		"isEnum":                isEnum,
		"exportedField":         exportedField,
		"downcaseName":          downcaseName,
	}
}
