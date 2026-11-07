package v2

//go:generate webrpc-gen -schema=../schema/test.ridl -target=golang -pkg=v2 -client -server -json=v2 -out=./v2.gen.go
