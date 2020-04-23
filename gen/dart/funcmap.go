package dart

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/webrpc/webrpc/schema"
)

var fieldTypeMap = map[schema.DataType]string{
	schema.T_Uint:      "int",
	schema.T_Uint8:     "int",
	schema.T_Uint16:    "int",
	schema.T_Uint32:    "int",
	schema.T_Uint64:    "int",
	schema.T_Int:       "int",
	schema.T_Int8:      "int",
	schema.T_Int16:     "int",
	schema.T_Int32:     "int",
	schema.T_Int64:     "int",
	schema.T_Float32:   "double",
	schema.T_Float64:   "double",
	schema.T_String:    "String",
	schema.T_Timestamp: "DateTime",
	schema.T_Null:      "null",
	schema.T_Any:       "dynamic",
	schema.T_Byte:      "String",
	schema.T_Bool:      "bool",
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
		return fmt.Sprintf("Map<%s, %s>", typK, typV), nil

	case schema.T_List:
		z, err := fieldType(in.List.Elem)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("List<%s>", z), nil

	case schema.T_Struct:
		return in.Struct.Name, nil

	default:
		if fieldTypeMap[in.Type] != "" {
			return fieldTypeMap[in.Type], nil
		}
	}
	return "", fmt.Errorf("could not represent type: %#v", in)
}

func isStructField(in *schema.VarType) bool {
	return in.Type == schema.T_Struct
}

func isListField(in *schema.VarType) bool {
	return in.Type == schema.T_List
}

func isMapField(in *schema.VarType) bool {
	return in.Type == schema.T_Map
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

func methodName(in interface{}) string {
	v, _ := downcaseName(in)
	return v
}
func isEnum(t schema.MessageType) bool {
	return t == "enum"
}

func isStruct(t schema.MessageType) bool {
	return t == "struct"
}

func lastField(msg schema.Message) schema.MessageField {
	return *msg.Fields[len(msg.Fields)-1]
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

func templateFuncMap(proto *schema.WebRPCSchema) map[string]interface{} {
	return map[string]interface{}{
		"fieldType":         fieldType,
		"methodName":        methodName,
		"isEnum":            isEnum,
		"exportedJSONField": exportedJSONField,
		"exportableField":   exportableField,
		"isStruct":          isStruct,
		"lastField":         lastField,
		"isStructField":     isStructField,
		"isListField":       isListField,
		"isMapField":        isMapField,
	}
}
