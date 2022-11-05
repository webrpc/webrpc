package gen

import (
	"fmt"
	"os"
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
		"stderrPrintf": stderrPrintf,
		"exit":         exit,
		"minVersion":   minVersion,

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
		"join":      strings.Join,
		"split":     split,
		"first":     first,
		"last":      last,
		"in":        in,
		"default":   defaultFn,
		"coalesce":  coalesce,
		"ternary":   ternary,
		"hasPrefix": strings.HasPrefix,
		"hasSuffix": strings.HasSuffix,
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
		"isEnum":     isEnum,
		"isStruct":   isStruct,
		"commaIfLen": commaIfLen,
		"listComma":  listComma,
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
		"jsFieldType":                         jsFieldType,
	}
}

// Similar to "printf" but instead of writing into the generated
// output file, stderrPrintf writes to webrpc-gen CLI stderr.
// Useful for printing template errors or for template debugging.
func stderrPrintf(format string, a ...interface{}) error {
	_, err := fmt.Fprintf(os.Stderr, format, a...)
	return err
}

// Terminate template execution. Useful for fatal errors.
func exit(code int) error {
	os.Exit(code)
	return nil
}

// Returns true if any of the given values match the first value.
func in(first interface{}, values ...interface{}) bool {
	for _, value := range values {
		if reflect.DeepEqual(first, value) {
			return true
		}
	}
	return false
}

// Returns defaultValue, if given value is empty.
func defaultFn(value interface{}, defaultValue interface{}) interface{} {
	val := reflect.ValueOf(value)
	if !val.IsValid() || val.IsZero() {
		return defaultValue
	}

	return value
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

// Ternary if-else. Returns first value if true, second value if false.
func ternary(boolean interface{}, first interface{}, second interface{}) interface{} {
	if toBool(boolean) {
		return first
	}
	return second
}

func toBool(in interface{}) bool {
	switch v := in.(type) {
	case bool:
		return v
	case string:
		if in == "true" {
			return true
		}
		if in == "false" {
			return false
		}
		panic(fmt.Sprintf("unexpected boolean %q", in))
	default:
		panic(fmt.Sprintf("unexpected boolean %v", v))
	}
}

func minVersion(version string, minVersion string) bool {
	major, minor, err := parseMajorMinorVersion(version)
	if err != nil {
		panic(fmt.Sprintf("minVersion: unexpected version %q", version))
	}

	minMajor, minMinor, err := parseMajorMinorVersion(minVersion)
	if err != nil {
		panic(fmt.Sprintf("minVersion: unexpected min version %q", minVersion))
	}

	if minMajor > major {
		return false
	}

	if minMinor > minor {
		return false
	}

	return true
}

func parseMajorMinorVersion(version string) (major int, minor int, err error) {
	version = strings.TrimPrefix(version, "v")
	parts := strings.Split(version, ".")

	major, err = strconv.Atoi(parts[0])
	if err != nil {
		return
	}

	if len(parts) > 1 {
		minor, err = strconv.Atoi(parts[1])
		if err != nil {
			return
		}
	}

	return
}

// appendComma
func commaIfLen(in []*schema.MethodArgument) string {
	if len(in) > 0 {
		return ","
	}
	return ""
}

// prependComma
func listComma(item int, count int) string {
	if item+1 < count {
		return ", "
	}
	return ""
}
