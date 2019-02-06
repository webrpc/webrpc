package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/webrpc/webrpc/gen"
	"github.com/webrpc/webrpc/schema"
)

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

func main() {
	f := `../../examples/golang-basics/example.webrpc.json`
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}

	schema, err := schema.ParseSchema(data)
	if err != nil {
		log.Fatal(err)
	}

	// spew.Dump(schema)

	// fmt.Println("")
	// fmt.Println("")

	// fmt.Println(schema.Messages[3].Fields[5].Type.String())

	tpl, err := gen.Gen(schema)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tpl)
}
