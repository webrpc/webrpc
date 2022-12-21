package gen

import (
	"fmt"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

// Returns true if given type is core type.
func isCoreType(v interface{}) bool {
	_, isCoreType := schema.CoreTypeFromString[toString(v)]
	return isCoreType
}

// Returns true if given type is struct.
func isStructType(v interface{}) bool {
	return !isCoreType(v) && !isListType(v) && !isMapType(v) && !isEnumType(v)
}

// Returns true if given type is enum.
func isEnumType(v interface{}) bool {
	return toString(v) == "enum"
}

// Returns true if given type is list (ie. `[]T`).
func isListType(v interface{}) bool {
	return strings.HasPrefix(toString(v), "[]")
}

// Return true if given type is map (ie. map<T1,T2>).
func isMapType(v interface{}) bool {
	key, value, found := stringsCut(toString(v), ",")
	return found && strings.HasPrefix(key, "map<") && strings.HasSuffix(value, ">")
}

// Returns given map's key type (ie. `T1` from `map<T1,T2>`)
func mapKeyType(v interface{}) string {
	str := toString(v)
	key, value, found := stringsCut(str, ",")
	if !found || !strings.HasPrefix(key, "map<") || !strings.HasSuffix(value, ">") {
		panic(fmt.Errorf("mapKeyValue: expected map<Type1,Type2>, got %v", str))
	}
	return strings.TrimPrefix(key, "map<")
}

// Returns given map's value type (ie. `T2` from `map<T1,T2>`)
func mapValueType(v interface{}) string {
	str := toString(v)
	key, value, found := stringsCut(str, ",")
	if !found || !strings.HasPrefix(key, "map<") || !strings.HasSuffix(value, ">") {
		panic(fmt.Errorf("mapKeyValue: expected map<Type1,Type2>, got %v", str))
	}
	return strings.TrimSuffix(value, ">")
}

// Returns list's element type (ie. `T` from `[]T`)
func listElemType(v interface{}) string {
	str := toString(v)
	if !strings.HasPrefix(str, "[]") {
		panic(fmt.Errorf("listElemType: expected []Type, got %v", str))
	}
	return strings.TrimPrefix(str, "[]")
}
