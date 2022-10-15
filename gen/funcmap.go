package gen

import (
	"errors"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

func templateFuncMap(proto *schema.WebRPCSchema) map[string]interface{} {
	return map[string]interface{}{
		// Generic template functions.
		"constPathPrefix": constPathPrefix,
		"countMethods":    countMethods,
		"commaIfLen":      commaIfLen,
		"isStruct":        isStruct,
		"isEnum":          isEnum,
		"downcaseName":    downcaseName,
		"listComma":       listComma,

		// Golang specific template functions.
		"goServiceMethodName":     goServiceMethodName,
		"goServiceMethodJSONName": goServiceMethodJSONName,
		"goHasGoFieldType":        goHasGoFieldType(proto),
		"goFieldTags":             goFieldTags,
		"goFieldType":             goFieldType,
		"goFieldOptional":         goFieldOptional,
		"goFieldTypeDef":          goFieldTypeDef,
		"goNewClientServiceName":  goNewClientServiceName,
		"goNewServerServiceName":  goNewServerServiceName,
		"goClientServiceName":     goClientServiceName,
		"goServerServiceName":     goServerServiceName,
		"goMethodInputs":          goMethodInputs,
		"goMethodOutputs":         goMethodOutputs,
		"goMethodArgName":         goMethodArgName,
		"goMethodArgType":         goMethodArgType,
		"goMethodArgNames":        goMethodArgNames,
		"goArgsList":              goArgsList,
		"goExportedField":         goExportedField,

		// TypeScript specific template functions.
		"tsFieldType":                         tsFieldType,
		"tsInterfaceName":                     tsInterfaceName,
		"tsMethodName":                        tsMethodName,
		"tsMethodInputs":                      tsMethodInputs,
		"tsMethodOutputs":                     tsMethodOutputs,
		"tsNewOutputArgResponse":              tsNewOutputArgResponse,
		"tsMethodArgumentInputInterfaceName":  tsMethodArgumentInputInterfaceName,
		"tsMethodArgumentOutputInterfaceName": tsMethodArgumentOutputInterfaceName,
		"tsServiceInterfaceName":              tsServiceInterfaceName,
		"tsExportableField":                   tsExportableField,
		"tsExportedField":                     tsExportedField,
		"tsExportedJSONField":                 tsExportedJSONField,
		"jsFieldType":                         jsFieldType,
	}
}

func commaIfLen(in []*schema.MethodArgument) string {
	if len(in) > 0 {
		return ","
	}
	return ""
}

func listComma(item int, count int) string {
	if item+1 < count {
		return ", "
	}
	return ""
}

func isStruct(t schema.MessageType) bool {
	return t == "struct"
}

func isEnum(t schema.MessageType) bool {
	return t == "enum"
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
