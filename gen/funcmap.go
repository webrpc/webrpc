package gen

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/golang-cz/textcase"
	"github.com/webrpc/webrpc/schema"
)

// Template functions are part of webrpc-gen API. Keep backward-compatible.
func templateFuncMap(proto *schema.WebRPCSchema, opts map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		// Template flow.
		"printfStderr": printfStderr,
		"exit":         exit,

		// Dictionary, aka runtime map[string]interface{}.
		"dict":   dict,
		"get":    get,
		"set":    set,
		"exists": exists,

		// Schema type helpers.
		"isBaseType":    isBaseType,
		"isMapType":     isMapType,
		"isArrayType":   isArrayType,
		"mapKeyType":    mapKeyType,
		"mapValueType":  mapValueType,
		"arrayItemType": arrayItemType,

		// String utils.
		"str":       str,
		"join":      join,
		"coalesce":  coalesce,
		"hasPrefix": strings.HasPrefix,
		"hasSuffix": strings.HasSuffix,
		"indent":    indent,
		"toLower":   applyStringFunction("toLower", strings.ToLower),
		"toUpper":   applyStringFunction("toLower", strings.ToUpper),
		"firstLetterToLower": applyStringFunction("firstLetterToLower", func(input string) string {
			if input == "" {
				return ""
			}
			return strings.ToLower(input[:1]) + input[1:]
		}),
		"firstLetterToUpper": applyStringFunction("firstLetterToUpper", func(input string) string {
			if input == "" {
				return ""
			}
			return strings.ToUpper(input[:1]) + input[1:]
		}),
		"camelCase":  applyStringFunction("camelCase", textcase.CamelCase),
		"pascalCase": applyStringFunction("pascalCase", textcase.PascalCase),
		"snakeCase":  applyStringFunction("snakeCase", textcase.SnakeCase),
		"kebabCase":  applyStringFunction("kebabCase", textcase.KebabCase),

		// OBSOLETE generic template functions.
		"constPathPrefix": constPathPrefix,
		"countMethods":    countMethods,
		"commaIfLen":      commaIfLen,
		"isStruct":        isStruct,
		"isEnum":          isEnum,
		"listComma":       listComma,
		"downcaseName": applyStringFunction("downcaseName", func(input string) string {
			if input == "" {
				return ""
			}
			return strings.ToLower(input[:1]) + input[1:]
		}),
		"tsMethodName": applyStringFunction("tsMethodName", func(input string) string {
			if input == "" {
				return ""
			}
			return strings.ToLower(input[:1]) + input[1:]
		}),
		"jsMethodName": applyStringFunction("jsMethodName", func(input string) string {
			if input == "" {
				return ""
			}
			return strings.ToLower(input[:1]) + input[1:]
		}),

		// OBSOLETE
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

		// OBSOLETE
		// JavaScript specific template functions.
		"jsExportKeyword":        jsExportKeyword(opts),
		"jsMethodInputs":         jsMethodInputs,
		"jsNewOutputArgResponse": jsNewOutputArgResponse,
		"jsServiceInterfaceName": jsServiceInterfaceName,
		"jsExportedJSONField":    jsExportedJSONField,
	}
}

// Returns first non-empty value.
func coalesce(v ...interface{}) interface{} {
	for _, v := range v {
		val := reflect.ValueOf(v)
		if !val.IsValid() || val.IsZero() {
			continue
		}
		return v
	}
	return ""
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

func constPathPrefix(in schema.VarName) (string, error) {
	return string(in) + "PathPrefix", nil
}

func countMethods(in []*schema.Method) (string, error) {
	return strconv.Itoa(len(in)), nil
}
