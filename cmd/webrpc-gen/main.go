package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/webrpc/webrpc/schema"
)

var (
	flags = flag.NewFlagSet("webrpc-gen", flag.ExitOnError)

	schemaFlag    = flags.String("schema", "", "webrpc schema file (required)")
	targetFlag    = flags.String("target", "", "target language for webrpc library generation, options: [go, ts] (required)")
	pkgFlag       = flags.String("pkg", "proto", "generated package name for target language, default: proto")
	outFlag       = flags.String("out", "", "generated output file, default: stdout")
	clientFlag    = flags.Bool("client", false, "enable webrpc client library generation, default: off")
	serverFlag    = flags.Bool("server", false, "enable webrpc server library generation, default: off")
	websocketFlag = flags.Bool("websocket", false, "enable websocket transport generation, default: off")
)

func main() {
	flags.Parse(os.Args[1:])

	if *schemaFlag == "" {
		fmt.Println("oops, you must pass a -schema flag, see -h for help/usage")
		os.Exit(1)
	}
	if *targetFlag == "" {
		fmt.Println("oops, you must pass a -target flag, see -h for help/usage")
		os.Exit(1)
	}
	if *targetFlag != "go" && *targetFlag != "ts" {
		fmt.Println("oops, you passed an invalid -target flag, see -h for help/usage")
		os.Exit(1)
	}

	// Parse the schema file
	schema, err := schema.ParseSchemaFile(*schemaFlag)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	spew.Dump(schema)
}
