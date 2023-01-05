package gen

import (
	"fmt"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

func toString(v interface{}) string {
	switch t := v.(type) {
	case schema.VarType:
		return t.String()
	case *schema.VarType:
		if t != nil {
			return t.String()
		}
		panic(fmt.Sprintf("toString(): nil %T", v))
	case schema.Type:
		return t.Kind
	case *schema.Type:
		return t.Kind
	case string:
		return t
	case map[string]interface{}:
		var b strings.Builder
		for k, v := range t {
			b.WriteString(fmt.Sprintf("%v=%v\n", k, v))
		}
		return b.String()
	default:
		panic(fmt.Sprintf("toString(): unknown arg type %T", v))
	}
}

func applyStringFunction(fnName string, fn func(string) string) func(v interface{}) string {
	return func(v interface{}) string {
		switch t := v.(type) {
		case string:
			return fn(t)
		default:
			panic(fmt.Errorf("%v: unknown arg type %T", fnName, v))
		}
	}
}
