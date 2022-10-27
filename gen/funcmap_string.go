package gen

import (
	"fmt"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

func indent(str string, prefix string) string {
	return prefix + strings.ReplaceAll(str, "\n", "\n"+prefix)
}

func str(v interface{}) string {
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

func join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}
