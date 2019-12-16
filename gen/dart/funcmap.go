package dart

import (
	"github.com/pkg/errors"
	"github.com/webrpc/webrpc/schema"
	"strings"
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
func templateFuncMap(proto *schema.WebRPCSchema) map[string]interface{} {
	return map[string]interface{}{
		"methodName": methodName,
	}
}
