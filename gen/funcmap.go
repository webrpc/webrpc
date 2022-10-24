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

// Dictionary, a map of string to string with a default value.
type dict struct {
	m            map[string]string
	defaultValue string
}

func newDict(defaultValue string) *dict {
	return &dict{
		m:            map[string]string{},
		defaultValue: defaultValue,
	}
}

func set(d *dict, key string, value string) string {
	if d == nil {
		panic(fmt.Sprintf("set: dict is nil, did you call $dict:=newDict ?"))
	}

	d.m[key] = value
	return ""
}

func get(d *dict, key interface{}) string {
	switch t := key.(type) {
	case string:
		return d.get(t)
	case schema.VarName:
		return d.get(string(t))
	case schema.VarType:
		return d.get(t.Type.String())
	case *schema.VarType:
		return d.get(t.Type.String())
	default:
		panic(fmt.Sprintf("get: unknown type %T", key))
	}
}

func (d *dict) get(key string) string {
	if d == nil {
		return ""
	}
	v, ok := d.m[key]
	if !ok {
		return d.defaultValue
	}
	return v
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
