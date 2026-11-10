package streaming

//go:generate webrpc-gen -schema=../../interop-python/streaming.ridl -target=golang -pkg=streaming -client -server -json=v2 -out=./streaming.gen.go
