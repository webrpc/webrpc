package gen

import (
	"errors"
	"fmt"
	"sort"
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

func goHasGoFieldType(proto *schema.WebRPCSchema) func(goFieldType string) (bool, error) {
	return func(goFieldType string) (bool, error) {
		return proto.HasFieldType(goFieldType)
	}
}
