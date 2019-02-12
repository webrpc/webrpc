package ridl

import (
	"reflect"
	"runtime"
)

func stateName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}
