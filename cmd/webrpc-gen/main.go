package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/webrpc/webrpc"
	"github.com/webrpc/webrpc/gen"
)

var flags = flag.NewFlagSet("webrpc-gen", flag.ExitOnError)

func main() {
	versionFlag := flags.Bool("version", false, "print webrpc version and exit")
	schemaFlag := flags.String("schema", "", "webrpc schema file (required)")
	targetFlag := flags.String("target", "", fmt.Sprintf("target generator for webrpc library generation (required), ie. github.com/webrpc/gen-golang@v0.6.0"))
	outFlag := flags.String("out", "", "generated output file, default: stdout")
	testFlag := flags.Bool("test", false, "test schema parsing (skips code-gen)")

	// Obsolete flags (v0.6.0).
	pkgFlag := flags.String("pkg", "proto", "generated package name for target language, default: proto")
	clientFlag := flags.Bool("client", false, "enable webrpc client library generation, default: off")
	serverFlag := flags.Bool("server", false, "enable webrpc server library generation, default: off")
	targetExtra := flags.String("extra", "", "target language extra/custom options")

	var cliFlags []string
	opts := map[string]interface{}{
		// Support the obsolete webrpc-gen flags (v0.6.0) in the new templates (v0.7.0+).
		"Pkg":    *pkgFlag,
		"Client": fmt.Sprintf("%v", clientFlag),
		"Server": fmt.Sprintf("%v", serverFlag),
		"Extra":  *targetExtra,
	}

	// Collect CLI -flags and template -Options.
	for _, flag := range os.Args[1:] {
		name, value, _ := strings.Cut(flag, "=")
		if !strings.HasPrefix(name, "-") {
			fmt.Printf("oops, option %q is invalid (format must be -name=value)\n", flag)
			os.Exit(1)
		}
		name = strings.TrimLeft(name, "-")
		if strings.ToUpper(name[:1]) == name[:1] {
			opts[name] = value
		} else {
			cliFlags = append(cliFlags, flag)
		}
	}

	flags.Parse(cliFlags)

	if *versionFlag {
		fmt.Printf("webrpc %s\n", gen.VERSION)
		os.Exit(0)
	}

	if *schemaFlag == "" {
		fmt.Println("oops, you must pass a -schema flag, see -h for help/usage")
		os.Exit(1)
	}

	// Parse+validate the webrpc schema file
	schema, err := webrpc.ParseSchemaFile(*schemaFlag)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Test the schema file (useful for ridl files)
	if *testFlag {
		jout, err := schema.ToJSON(true)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(jout)
		os.Exit(0)
	}

	// Code-gen targets
	if *targetFlag == "" {
		fmt.Println("oops, you must pass a -target flag, see -h for help/usage")
		os.Exit(1)
	}

	protoGen, err := gen.Generate(schema, *targetFlag, opts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Write output
	if *outFlag != "" && *outFlag != "stdout" {
		outfile := *outFlag
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if outfile[0:1] != "/" {
			outfile = filepath.Join(cwd, outfile)
		}

		outdir := filepath.Dir(outfile)
		if _, err := os.Stat(outdir); os.IsNotExist(err) {
			err := os.MkdirAll(outdir, 0755)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}

		err = ioutil.WriteFile(outfile, []byte(protoGen), 0644)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	} else {
		fmt.Println(protoGen)
	}
}
