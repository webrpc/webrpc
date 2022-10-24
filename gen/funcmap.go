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
		"dict": dict,
		"set":  set,
		"get":  get,
		"str":  str,

		// String manipulation.
		"toLower": applyStringFunction("toLower", strings.ToLower),
		"toUpper": applyStringFunction("toLower", strings.ToUpper),
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

		// String utils.
		"hasPrefix": strings.HasPrefix,
		"hasSuffix": strings.HasSuffix,

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

func str(v interface{}) string {
	switch t := v.(type) {
	case schema.VarName:
		return string(t)
	case schema.VarType:
		return t.Type.String()
	case *schema.VarType:
		return t.Type.String()
	case schema.MessageType:
		return string(t)
	case string:
		return t
	default:
		panic(fmt.Sprintf("str: unknown arg type %T", v))
	}
}

func applyStringFunction(fnName string, fn func(string) string) func(v interface{}) string {
	return func(v interface{}) string {
		switch t := v.(type) {
		case schema.VarName:
			return fn(string(t))
		case string:
			return fn(t)
		default:
			panic(fmt.Errorf("%v: unknown arg type %T", fnName, v))
		}
	}
}

// Create new dictionary.
func dict(pairs ...interface{}) map[string]interface{} {
	if len(pairs)%2 == 1 {
		panic("dict must be created with even number of parameters (key:value pairs)")
	}

	m := map[string]interface{}{}
	for i := 0; i < len(pairs); i += 2 {
		key, ok := pairs[i].(string)
		if !ok {
			panic(fmt.Sprintf("dict argument(%v) must be string key", i))
		}
		m[key] = pairs[i+1]
	}

	return m
}

func set(m map[string]interface{}, key string, value interface{}) string {
	m[key] = value
	return ""
}

func get(m map[string]interface{}, key interface{}) interface{} {
	switch t := key.(type) {
	case string:
		return m[t]
	case schema.VarName:
		return m[string(t)]
	case schema.VarType:
		return m[t.Type.String()]
	case *schema.VarType:
		return m[t.Type.String()]
	default:
		panic(fmt.Sprintf("get: unknown type %T", key))
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

func constPathPrefix(in schema.VarName) (string, error) {
	return string(in) + "PathPrefix", nil
}

func countMethods(in []*schema.Method) (string, error) {
	return strconv.Itoa(len(in)), nil
}
