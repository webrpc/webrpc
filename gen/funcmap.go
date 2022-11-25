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
		"stderrPrintf": stderrPrintf, // v0.7.0
		"exit":         exit,         // v0.7.0
		"minVersion":   minVersion,   // v0.7.0

		// Dictionary, aka runtime map[string]interface{}.
		"dict":   dict,   // v0.7.0
		"get":    get,    // v0.7.0
		"set":    set,    // v0.7.0
		"exists": exists, // v0.7.0

		// Schema type helpers.
		"isBasicType":  isBasicType,  // v0.7.0
		"isStructType": isStructType, // v0.7.0
		"isEnumType":   isEnumType,   // v0.7.0
		"isMapType":    isMapType,    // v0.7.0
		"isListType":   isListType,   // v0.7.0
		"mapKeyType":   mapKeyType,   // v0.7.0
		"mapValueType": mapValueType, // v0.7.0
		"listElemType": listElemType, // v0.7.0

		// String utils.
		"join":       strings.Join,                                    // v0.7.0
		"split":      split,                                           // v0.7.0
		"first":      first,                                           // v0.7.0
		"last":       last,                                            // v0.7.0
		"sort":       sortFn,                                          // v0.8.0
		"in":         in,                                              // v0.7.0
		"default":    defaultFn,                                       // v0.7.0
		"coalesce":   coalesce,                                        // v0.7.0
		"ternary":    ternary,                                         // v0.7.0
		"hasPrefix":  strings.HasPrefix,                               // v0.7.0
		"hasSuffix":  strings.HasSuffix,                               // v0.7.0
		"trimPrefix": strings.TrimPrefix,                              // v0.8.0
		"trimSuffix": strings.TrimSuffix,                              // v0.8.0
		"toLower":    applyStringFunction("toLower", strings.ToLower), // v0.7.0
		"toUpper":    applyStringFunction("toLower", strings.ToUpper), // v0.7.0
		"firstLetterToLower": applyStringFunction("firstLetterToLower", func(input string) string { // v0.7.0
			if input == "" {
				return ""
			}
			return strings.ToLower(input[:1]) + input[1:]
		}),
		"firstLetterToUpper": applyStringFunction("firstLetterToUpper", func(input string) string { // v0.7.0
			if input == "" {
				return ""
			}
			return strings.ToUpper(input[:1]) + input[1:]
		}),
		"camelCase":  applyStringFunction("camelCase", textcase.CamelCase),   // v0.7.0
		"pascalCase": applyStringFunction("pascalCase", textcase.PascalCase), // v0.7.0
		"snakeCase":  applyStringFunction("snakeCase", textcase.SnakeCase),   // v0.7.0
		"kebabCase":  applyStringFunction("kebabCase", textcase.KebabCase),   // v0.7.0
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
