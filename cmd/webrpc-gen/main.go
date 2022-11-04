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

	// Collect CLI -flags and custom -Flags (template options).
	cliFlags, templateOpts, err := collectFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse CLI flags: %v", err)
		os.Exit(1)
	}

	flags.Parse(cliFlags)

	if *versionFlag {
		fmt.Printf("webrpc %s\n", gen.VERSION)
		os.Exit(0)
	}

	if *schemaFlag == "" {
		fmt.Fprintln(os.Stderr, "oops, you must pass a -schema flag, see -h for help/usage")
		os.Exit(1)
	}

	// Parse+validate the webrpc schema file
	schema, err := webrpc.ParseSchemaFile(*schemaFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse schema file: %v", err)
		os.Exit(1)
	}

	// Test the schema file (useful for ridl files)
	if *testFlag {
		jout, err := schema.ToJSON(true)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Println(jout)
		os.Exit(0)
	}

	// Code-gen targets
	if *targetFlag == "" {
		fmt.Fprintln(os.Stderr, "oops, you must pass a -target flag, see -h for help/usage")
		os.Exit(1)
	}

	protoGen, err := gen.Generate(schema, *targetFlag, *refreshCache, templateOpts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// Write output to stdout
	if *outFlag == "" || *outFlag == "stdout" {
		fmt.Println(protoGen)
		os.Exit(0)
	}

	// Write output to a file
	outfile := *outFlag
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if outfile[0:1] != "/" {
		outfile = filepath.Join(cwd, outfile)
	}

	outdir := filepath.Dir(outfile)
	if _, err := os.Stat(outdir); os.IsNotExist(err) {
		err := os.MkdirAll(outdir, 0755)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}

	err = ioutil.WriteFile(outfile, []byte(protoGen), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func collectFlags(flags []string) (cliFlags []string, templateOpts map[string]interface{}, err error) {
	templateOpts = map[string]interface{}{}

	for _, flag := range flags {
		name, value, _ := strings.Cut(flag, "=")
		if !strings.HasPrefix(name, "-") {
			return nil, nil, fmt.Errorf("option %q is invalid (expected -name=value)", flag)
		}
		name = strings.TrimLeft(name, "-")
		if len(name) == 0 {
			return nil, nil, fmt.Errorf("option %q is invalid (expected -name=value)", flag)
		}

		if name[0] >= 'A' && name[0] <= 'Z' {
			// -Flag
			templateOpts[name] = value
		} else {
			// -flag
			// Support webrpc-gen v0.6.0 flags in new templates.
			switch name {
			case "pkg", "client", "server":
				templateOpts[strings.Title(name)] = value
			case "extra":
				if value == "noexports" {
					templateOpts["Export"] = false
				}
			default:
				cliFlags = append(cliFlags, flag)
			}
		}
	}

	return
}
