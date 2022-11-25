package gen

import (
	"fmt"
	"reflect"
)

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
