package gen

import (
	"strings"

	"github.com/Masterminds/sprig/v3"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang-cz/textcase"
)

// Template functions are part of webrpc-gen API. Keep backward-compatible.
func templateFuncMap(opts map[string]interface{}) map[string]interface{} {
	f := sprig.FuncMap()
	extra := map[string]interface{}{
		// Template flow, errors, debugging.
		"stderrPrint":  stderrPrint,  // v0.13.0
		"stderrPrintf": stderrPrintf, // v0.7.0
		"exit":         exit,         // v0.7.0
		"minVersion":   minVersion,   // v0.7.0
		"dump":         spew.Sdump,   // v0.13.0
		"hasField":     hasField,     // v0.13.0

		// Schema type helpers.
		"isBasicType":  isCoreType,   // v0.7.0 (deprecated in v0.9.0)
		"isCoreType":   isCoreType,   // v0.9.0
		"isStructType": isStructType, // v0.7.0
		"isEnumType":   isEnumType,   // v0.7.0
		"isMapType":    isMapType,    // v0.7.0
		"isListType":   isListType,   // v0.7.0
		"mapKeyType":   mapKeyType,   // v0.7.0
		"mapValueType": mapValueType, // v0.7.0
		"listElemType": listElemType, // v0.7.0

		// Dictionary (map[string]any).
		"dict":   dict,   // v0.7.0
		"get":    get,    // v0.7.0
		"set":    set,    // v0.7.0
		"exists": exists, // v0.7.0

		// Generic utils.
		"lastIndex": lastIndex, // v0.18.0
		"array":     array,     // v0.11.2 (string support since v0.8.0)
		"append":    appendFn,  // v0.11.2 (string support since v0.7.0)
		"first":     first,     // v0.11.2 (string support since v0.7.0)
		"last":      last,      // v0.11.2 (string support since v0.7.0)
		"sort":      sortFn,    // v0.11.2 (string support since v0.8.0)
		"coalesce":  coalesce,  // v0.7.0
		"default":   defaultFn, // v0.7.0
		"in":        in,        // v0.7.0
		"ternary":   ternary,   // v0.7.0

		// String utils.
		"join":       join,                                            // v0.7.0
		"split":      split,                                           // v0.7.0
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
		"firstWordToLower": applyStringFunction("firstWordToLower", func(input string) string { // v0.29.0
			if input == "" {
				return ""
			}
			if input[0] >= 'a' && input[0] <= 'z' {
				return input
			}

			letters := []rune(input)
			length := len(letters)
			if length == 0 {
				return ""
			}

			isUpper := func(r rune) bool { return r >= 'A' && r <= 'Z' }
			isLower := func(r rune) bool { return r >= 'a' && r <= 'z' }
			isDigit := func(r rune) bool { return r >= '0' && r <= '9' }

			j := 0
			if isUpper(letters[0]) {
				for j < length && isUpper(letters[j]) {
					j++
				}
				if j > 1 && j < length && isLower(letters[j]) && (j+1) < length && isLower(letters[j+1]) {
					j--
				}
				for j < length && (isLower(letters[j]) || isDigit(letters[j])) {
					j++
				}
			} else {
				for j < length && (isLower(letters[j]) || isDigit(letters[j])) {
					j++
				}
			}

			// Lowercase the first word prefix and append the rest unchanged.
			prefix := strings.ToLower(string(letters[:j]))
			return prefix + string(letters[j:])
		}),
		"camelCase":  applyStringFunction("camelCase", textcase.CamelCase),   // v0.7.0
		"pascalCase": applyStringFunction("pascalCase", textcase.PascalCase), // v0.7.0
		"snakeCase":  applyStringFunction("snakeCase", textcase.SnakeCase),   // v0.7.0
		"kebabCase":  applyStringFunction("kebabCase", textcase.KebabCase),   // v0.7.0
		"replaceAll": strings.ReplaceAll,
	}

	for k, v := range extra {
		f[k] = v
	}

	return f
}
