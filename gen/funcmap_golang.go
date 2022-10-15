package gen

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

var goFieldTypeMap = map[schema.DataType]string{
	schema.T_Uint:      "uint",
	schema.T_Uint8:     "uint8",
	schema.T_Uint16:    "uint16",
	schema.T_Uint32:    "uint32",
	schema.T_Uint64:    "uint64",
	schema.T_Int:       "int",
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

func goServiceMethodName(in schema.VarName) (string, error) {
	s := string(in)
	return "serve" + strings.ToUpper(s[0:1]) + s[1:], nil
}

func goServiceMethodJSONName(in schema.VarName) (string, error) {
	s := string(in)
	return "serve" + strings.ToUpper(s[0:1]) + s[1:] + "JSON", nil
}

func goNewServerServiceName(in schema.VarName) (string, error) {
	return "New" + string(in) + "Server", nil
}

func goNewClientServiceName(in schema.VarName) (string, error) {
	return "New" + string(in) + "Client", nil
}

func goFieldType(in *schema.VarType) (string, error) {
	switch in.Type {
	case schema.T_Map:
		typK, ok := goFieldTypeMap[in.Map.Key]
		if !ok {
			return "", fmt.Errorf("unknown type mapping %v", in.Map.Key)
		}
		typV, err := goFieldType(in.Map.Value)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("map[%v]%s", typK, typV), nil

	case schema.T_List:
		z, err := goFieldType(in.List.Elem)
		if err != nil {
			return "", err
		}
		return "[]" + z, nil

	case schema.T_Struct:
		return "*" + in.Struct.Name, nil

	default:
		if goFieldTypeMap[in.Type] != "" {
			return goFieldTypeMap[in.Type], nil
		}
	}
	return "", fmt.Errorf("could not represent type: %#v", in)
}

func goFieldOptional(field *schema.MessageField) (string, error) {
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
		if goFieldTypeMap[field.Type.Type] != "" {
			return "*", nil
		}
	}
	return "", fmt.Errorf("could not represent type: %#v", field)
}

func goFieldTypeDef(in *schema.MessageField) (string, error) {
	gogoFieldType := ""

	meta := in.Meta
	for kk := range meta {
		for k, v := range meta[kk] {
			if k == "go.field.type" {
				gogoFieldType = fmt.Sprintf("%v", v)
			}
		}
	}

	if gogoFieldType != "" {
		return gogoFieldType, nil
	}

	return goFieldType(in.Type)
}

func goFieldTags(in *schema.MessageField) (string, error) {
	goFieldTags := map[string]interface{}{}

	jsonFieldName, err := downcaseName(in.Name)
	if err != nil {
		return "", err
	}
	goFieldTags["json"] = fmt.Sprintf("%s", jsonFieldName)

	goTagJSON := ""

	meta := in.Meta
	for kk := range meta {
		for k, v := range meta[kk] {

			switch {
			case k == "json":
				if goTagJSON == "" {
					goFieldTags["json"] = fmt.Sprintf("%v", v)
				}

			case strings.HasPrefix(k, "go.tag.json"):
				goTagJSON = fmt.Sprintf("%v", v)
				if !strings.HasPrefix(goTagJSON, fmt.Sprintf("%v", goFieldTags["json"])) {
					return "", errors.New("go.tag.json is invalid, it must match the json fieldname")
				}
				goFieldTags[k[7:]] = fmt.Sprintf("%v", v)

			case strings.HasPrefix(k, "go.tag."):
				if k == "go.tag.json" {
					goTagJSON = fmt.Sprintf("%v", v)
				}
				goFieldTags[k[7:]] = fmt.Sprintf("%v", v)
			}

		}
	}

	tagKeys := []string{}
	for k, _ := range goFieldTags {
		if k != "json" {
			tagKeys = append(tagKeys, k)
		}
	}
	sort.StringSlice(tagKeys).Sort()
	tagKeys = append([]string{"json"}, tagKeys...)

	tags := []string{}
	for _, k := range tagKeys {
		tags = append(tags, fmt.Sprintf(`%s:"%v"`, k, goFieldTags[k]))
	}

	return "`" + strings.Join(tags, " ") + "`", nil
}

func constPathPrefix(in schema.VarName) (string, error) {
	return string(in) + "PathPrefix", nil
}

func countMethods(in []*schema.Method) (string, error) {
	return strconv.Itoa(len(in)), nil
}

func goClientServiceName(in schema.VarName) (string, error) {
	s := string(in)
	return strings.ToLower(s[0:1]) + s[1:] + "Client", nil
}

func goServerServiceName(in schema.VarName) (string, error) {
	s := string(in)
	return strings.ToLower(s[0:1]) + s[1:] + "Server", nil
}

func goMethodArgName(in *schema.MethodArgument) string {
	name := string(in.Name)
	if name == "" && in.Type != nil {
		name = in.Type.String()
	}
	if name != "" {
		return name
	}
	return ""
}

func goMethodArgType(in *schema.MethodArgument) string {
	z, err := goFieldType(in.Type)
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

func goMethodInputs(in []*schema.MethodArgument) (string, error) {
	inputs := []string{"ctx context.Context"}
	for i := range in {
		inputs = append(inputs, fmt.Sprintf("%s %s", goMethodArgName(in[i]), goMethodArgType(in[i])))
	}
	return strings.Join(inputs, ", "), nil
}

func goMethodOutputs(in []*schema.MethodArgument) (string, error) {
	outputs := []string{}
	for i := range in {
		outputs = append(outputs, goMethodArgType(in[i]))
	}
	outputs = append(outputs, "error")
	return strings.Join(outputs, ", "), nil
}

func goMethodArgNames(in []*schema.MethodArgument) (string, error) {
	inputs := []string{}
	for i := range in {
		inputs = append(inputs, fmt.Sprintf("%s", goMethodArgName(in[i])))
	}
	return strings.Join(inputs, ", "), nil
}

func goArgsList(in []*schema.MethodArgument, prefix string) (string, error) {
	ins := []string{}
	for i := range in {
		ins = append(ins, fmt.Sprintf("%s%d", prefix, i))
	}
	return strings.Join(ins, ", "), nil
}
func goExportedField(in *schema.MessageField) (string, error) {
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

func goHasGoFieldType(proto *schema.WebRPCSchema) func(goFieldType string) (bool, error) {
	return func(goFieldType string) (bool, error) {
		return proto.HasFieldType(goFieldType)
	}
}
