package ridl

import (
	"reflect"
	"runtime"
)

func stateFnName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}
