//go:generate ../../../bin/webrpc-gen -schema=sample.webrpc.json -target=go -pkg=proto -server -client -out=./proto/proto.gen.go
package golang_test

import "testing"

func TestStuff(t *testing.T) {

}
