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
	targetFlag := flags.String("target", "", fmt.Sprintf("target generator (required), ie. golang@v0.6.0"))
	outFlag := flags.String("out", "", "generated output file, default: stdout")
	refreshCache := flags.Bool("refreshCache", false, "refresh webrpc cache")
	testFlag := flags.Bool("test", false, "test schema parsing (skips code-gen)")

	var cliFlags []string
	templateOpts := map[string]interface{}{}

	// Collect custom template -Options and CLI -flags.
	for _, flag := range os.Args[1:] {
		name, value, _ := strings.Cut(flag, "=")
		if !strings.HasPrefix(name, "-") {
			fmt.Printf("oops, option %q is invalid (format must be -name=value)\n", flag)
			os.Exit(1)
		}
		name = strings.TrimLeft(name, "-")
		if strings.ToUpper(name[:1]) == name[:1] {
			templateOpts[name] = value
		} else {
			switch name {
			case "pkg", "client", "server", "extra":
				// Support obsolete flags (v0.6.0) in new templates.
				templateOpts[strings.Title(name)] = value
			default:
				cliFlags = append(cliFlags, flag)
			}
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

	protoGen, err := gen.Generate(schema, *targetFlag, *refreshCache, templateOpts)
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
