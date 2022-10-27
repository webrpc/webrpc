package gen

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/golang-cz/textcase"
	"github.com/webrpc/webrpc/schema"
)

// Template functions are part of webrpc-gen API. Keep backward-compatible.
func templateFuncMap(proto *schema.WebRPCSchema, opts map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		// Flow.
		"printfStderr": printfStderr,
		"exit":         exit,

		// Dictionary, aka runtime map[string]interface{}.
		"dict":   dict,
		"get":    get,
		"set":    set,
		"exists": exists,

		"first": first, // rename to coalesce?

		// Type helpers.
		"isBaseType":    isBaseType,
		"isMapType":     isMapType,
		"isArrayType":   isArrayType,
		"mapKeyType":    mapKeyType,
		"mapValueType":  mapValueType,
		"arrayItemType": arrayItemType,

		// String utils.
		"str":       str,
		"join":      join,
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
		// Golang specific template functions.
		"goHasGoFieldType": goHasGoFieldType(proto),
		"goFieldTags":      goFieldTags,
		"goFieldType":      goFieldType,
		"goFieldOptional":  goFieldOptional,
		"goFieldTypeDef":   goFieldTypeDef,
		"goMethodArgType":  goMethodArgType,
		"goExportedField":  goExportedField,

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

// Base webrpc type is anything but map, array or custom message type.
func isBaseType(v interface{}) bool {
	str := str(v)
	_, ok := schema.DataTypeFromString[str]
	return ok
}

func isMapType(v interface{}) bool {
	str := str(v)
	key, value, found := strings.Cut(str, ",")
	return found && strings.HasPrefix(key, "map<") && strings.HasSuffix(value, ">")
}

func isArrayType(v interface{}) bool {
	str := str(v)
	return strings.HasPrefix(str, "[]")
}

// Expects webrpc map type, ie. "map<Type1,Type2>".
// Returns the key type, ie. "Type1".
func mapKeyType(v interface{}) string {
	str := str(v)
	key, value, found := strings.Cut(str, ",")
	if !found || !strings.HasPrefix(key, "map<") || !strings.HasSuffix(value, ">") {
		panic(fmt.Errorf("mapKeyValue: expected map<Type1,Type2>, got %v", str))
	}
	return strings.TrimPrefix(key, "map<")
}

// Expects webrpc map type, ie. "map<Type1,Type2>".
// Returns the value type, ie. "Type2".
func mapValueType(v interface{}) string {
	str := str(v)
	key, value, found := strings.Cut(str, ",")
	if !found || !strings.HasPrefix(key, "map<") || !strings.HasSuffix(value, ">") {
		panic(fmt.Errorf("mapKeyValue: expected map<Type1,Type2>, got %v", str))
	}
	return strings.TrimSuffix(value, ">")
}

// Expects webrpc array of types, ie. "[]Type".
// Returns the item, ie. "Type".
func arrayItemType(v interface{}) string {
	str := str(v)
	if !strings.HasPrefix(str, "[]") {
		panic(fmt.Errorf("arrayItemType: expected []Type, got %v", str))
	}
	return strings.TrimPrefix(str, "[]")
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

func set(m map[string]interface{}, key string, value interface{}) string {
	m[key] = value
	return ""
}

// TODO: Support slices too?
func exists(m map[string]interface{}, key string) bool {
	_, ok := m[key]
	return ok
}

// Returns first non-empty value.
func first(v ...interface{}) interface{} {
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
