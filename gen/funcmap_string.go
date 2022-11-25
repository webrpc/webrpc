package gen

import (
	"fmt"
	"sort"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

func toString(v interface{}) string {
	switch t := v.(type) {
	case schema.VarName:
		return string(t)
	case schema.VarType:
		return t.String()
	case *schema.VarType:
		return t.String()
	case schema.MessageType:
		return string(t)
	case string:
		return t
	case map[string]interface{}:
		var b strings.Builder
		for k, v := range t {
			b.WriteString(fmt.Sprintf("%v=%v\n", k, v))
		}
		return b.String()
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

func sortFn(slice []string) []string {
	sorted := make([]string, len(slice))
	copy(sorted, slice)
	sort.Strings(sorted)
	return sorted
}

func split(sep string, str string) []string {
	return strings.Split(str, sep)
}

func first(elems []string) (string, error) {
	if len(elems) == 0 {
		return "", fmt.Errorf("first: no elements in string array")
	}
	return elems[0], nil
}

func last(elems []string) (string, error) {
	if len(elems) == 0 {
		return "", fmt.Errorf("last: no elements in string array")
	}
	return elems[len(elems)-1], nil
}
