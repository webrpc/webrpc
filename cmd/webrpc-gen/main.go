package main

import "fmt"

func main() {
	fmt.Println("webrpc-gen")

}

/*

USAGE/flags
===========

-schema     -- webrpc schema file
-pkg        -- package name used for target language,
-websocket  -- enable websocket support, ie. if we code-gen websocket transport stuff too (flag, on/off)
-target     -- target language: go, ts, ...
-client     -- flag, code-generate client
-server     -- flag, code-generate server
-out        -- Directory to output the ie. <pkg>.gen.go file, if empty will print to stdout

extra flags, target-lang specific ... how to pass those?
-fastjson -- flag to use optimized json encoding, withou reflection.
							ie. for Go, use: https://github.com/pquerna/ffjson
							or, .. https://github.com/valyala/quicktemplate#use-cases


*/

// webrpc-gen is used to code-generate the server interfaces, handlers and message payload types with serialization.
// this single util will code-generate for multiple language targets, ie. Go, TypeScript, and one day.. node_js node_ts
// rust, ruby, python, etc......
