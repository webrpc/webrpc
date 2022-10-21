package gen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/golang-cz/textcase"
	"github.com/webrpc/webrpc/schema"
)

// Template functions are part of webrpc-gen API. Keep backward-compatible.
func templateFuncMap(proto *schema.WebRPCSchema, opts TargetOptions) map[string]interface{} {
	return map[string]interface{}{

		// String manipulation.
		"toLower": applyStringFunction("toLower", strings.ToLower),
		"toUpper": applyStringFunction("toLower", strings.ToUpper),
		"firstLetterToLower": applyStringFunction("firstLetterToLower", func(input string) string {
			if input == "" {
				return ""
			}
			return strings.ToLower(input[:1]) + input[1:]
		}),
		"firstLetterToUpper": applyStringFunction("firstLetterToUpper", func(input string) string {

			return input
		}),
		"camelCase":  applyStringFunction("camelCase", textcase.CamelCase),
		"pascalCase": applyStringFunction("pascalCase", textcase.PascalCase),
		"snakeCase":  applyStringFunction("snakeCase", textcase.SnakeCase),
		"kebabCase":  applyStringFunction("kebabCase", textcase.KebabCase),

		// Generic template functions.
		"constPathPrefix": constPathPrefix,
		"countMethods":    countMethods,
		"commaIfLen":      commaIfLen,
		"isStruct":        isStruct,
		"isEnum":          isEnum,
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

		// JavaScript specific template functions.
		"jsExportKeyword":        jsExportKeyword(opts),
		"jsMethodInputs":         jsMethodInputs,
		"jsNewOutputArgResponse": jsNewOutputArgResponse,
		"jsServiceInterfaceName": jsServiceInterfaceName,
		"jsExportedJSONField":    jsExportedJSONField,
	}
}

func applyStringFunction(fnName string, fn func(string) string) func(v interface{}) (string, error) {
	return func(v interface{}) (string, error) {
		switch t := v.(type) {
		case schema.VarName:
			return fn(string(t)), nil
		case string:
			return fn(t), nil
		default:
			return "", fmt.Errorf("%v: unknown arg type %T", fnName, v)
		}
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

// func camelCase(v interface{}) (string, error) {
// 	switch t := v.(type) {
// 	case schema.VarName:
// 		return textcase.CamelCase(string(t)), nil
// 	case string:
// 		return textcase.CamelCase(t), nil
// 	default:
// 		return "", errors.New("camelCase, unknown arg type")
// 	}
// }

// func pascalCase(v interface{}) (string, error) {
// 	switch t := v.(type) {
// 	case schema.VarName:
// 		return textcase.PascalCase(string(t)), nil
// 	case string:
// 		return textcase.PascalCase(t), nil
// 	default:
// 		return "", errors.New("PascalCase, unknown arg type")
// 	}
// }

// func snakeCase(v interface{}) (string, error) {
// 	switch t := v.(type) {
// 	case schema.VarName:
// 		return textcase.SnakeCase(string(t)), nil
// 	case string:
// 		return textcase.SnakeCase(t), nil
// 	default:
// 		return "", errors.New("snake_case, unknown arg type")
// 	}
// }

// func kebabCase(v interface{}) (string, error) {
// 	switch t := v.(type) {
// 	case schema.VarName:
// 		return textcase.KebabCase(string(t)), nil
// 	case string:
// 		return textcase.KebabCase(t), nil
// 	default:
// 		return "", errors.New("snake_case, unknown arg type")
// 	}
// }

// func toLower(v interface{}) (string, error) {
// 	switch t := v.(type) {
// 	case schema.VarName:
// 		return strings.ToLower(string(t)), nil
// 	case string:
// 		return strings.ToLower(t), nil
// 	default:
// 		return "", errors.New("snake_case, unknown arg type")
// 	}
// }

func constPathPrefix(in schema.VarName) (string, error) {
	return string(in) + "PathPrefix", nil
}

func countMethods(in []*schema.Method) (string, error) {
	return strconv.Itoa(len(in)), nil
}
