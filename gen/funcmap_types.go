package gen

import (
	"fmt"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

// Return true if given type is a base webrpc type (not a custom struct).
// See https://github.com/webrpc/webrpc/tree/master/schema#type-system.
func isBaseType(v interface{}) bool {
	str := toString(v)
	_, ok := schema.DataTypeFromString[str]

	return ok && !isMapType(v) && !isArrayType(v)
}

// TODO: Consider removing.
func isStruct(t schema.MessageType) bool {
	return t == "struct"
}

// TODO: Consider removing.
func isEnum(t schema.MessageType) bool {
	return t == "enum"
}

func isMapType(v interface{}) bool {
	str := toString(v)
	key, value, found := strings.Cut(str, ",")
	return found && strings.HasPrefix(key, "map<") && strings.HasSuffix(value, ">")
}

func isArrayType(v interface{}) bool {
	str := toString(v)
	return strings.HasPrefix(str, "[]")
}

// Expects webrpc map type, ie. "map<Type1,Type2>".
// Returns the key type, ie. "Type1".
func mapKeyType(v interface{}) string {
	str := toString(v)
	key, value, found := strings.Cut(str, ",")
	if !found || !strings.HasPrefix(key, "map<") || !strings.HasSuffix(value, ">") {
		panic(fmt.Errorf("mapKeyValue: expected map<Type1,Type2>, got %v", str))
	}
	return strings.TrimPrefix(key, "map<")
}

// Expects webrpc map type, ie. "map<Type1,Type2>".
// Returns the value type, ie. "Type2".
func mapValueType(v interface{}) string {
	str := toString(v)
	key, value, found := strings.Cut(str, ",")
	if !found || !strings.HasPrefix(key, "map<") || !strings.HasSuffix(value, ">") {
		panic(fmt.Errorf("mapKeyValue: expected map<Type1,Type2>, got %v", str))
	}
	return strings.TrimSuffix(value, ">")
}

// Expects webrpc array of types, ie. "[]Type".
// Returns the item, ie. "Type".
func arrayItemType(v interface{}) string {
	str := toString(v)
	if !strings.HasPrefix(str, "[]") {
		panic(fmt.Errorf("arrayItemType: expected []Type, got %v", str))
	}
	return strings.TrimPrefix(str, "[]")
}
