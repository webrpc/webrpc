package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/webrpc/webrpc"
	"github.com/webrpc/webrpc/gen"
	"github.com/webrpc/webrpc/schema"
)

var flags = flag.NewFlagSet("webrpc-gen", flag.ExitOnError)

func main() {
	versionFlag := flags.Bool("version", false, "print webrpc version and exit")
	schemaFlag := flags.String("schema", "", "webrpc schema file (required)")
	targetFlag := flags.String("target", "", fmt.Sprintf("target generator (required), ie. golang@v0.7.0"))
	outFlag := flags.String("out", "", "generated output file, default: stdout")
	fmtFlag := flags.Bool("fmt", true, "format generated code")
	refreshCacheFlag := flags.Bool("refreshCache", false, "refresh webrpc cache")
	testFlag := flags.Bool("test", false, "test schema parsing (skips code-gen)")

	// Collect CLI -flags and custom template -options.
	cliFlags, templateOpts, err := collectCliArgs(flags, os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse CLI flags: %v\n", err)
		os.Exit(1)
	}

	flags.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -schema=<file.[ridl|json]> -target=<target> -out=<file> [...targetOpts]\n", flags.Name())
		flags.PrintDefaults()
		fmt.Fprintf(os.Stderr, "See https://github.com/webrpc/webrpc for more info.\n")

		if *targetFlag != "" {
			fmt.Fprintf(os.Stderr, "\nTarget generator usage:\n")
			templateHelp, err := gen.Generate(&schema.WebRPCSchema{}, *targetFlag, &gen.Config{TemplateOptions: templateOpts})
			if err != nil {
				fmt.Fprintln(os.Stderr, templateHelp)
			} else {
				fmt.Fprintf(os.Stderr, "failed to render -help: %v\n", err)
			}
		}
	}

	flags.Parse(cliFlags)

	if *versionFlag {
		fmt.Printf("webrpc %s\n", gen.VERSION)
		os.Exit(0)
	}

	if *schemaFlag == "" {
		fmt.Fprintf(os.Stderr, "-schema flag is required\n\n")
		flags.Usage()
		os.Exit(1)
	}

	// Parse+validate the webrpc schema file
	schema, err := webrpc.ParseSchemaFile(*schemaFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse schema file: %v\n", err)
		os.Exit(1)
	}

	// Test the schema file (useful for ridl files)
	if *testFlag {
		out, err := schema.ToJSON(true)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Println(out)
		os.Exit(0)
	}

	// Code-gen targets
	if *targetFlag == "" {
		fmt.Fprintf(os.Stderr, "-target flag is required\n\n")
		flags.Usage()
		os.Exit(1)
	}

	config := &gen.Config{
		RefreshCache:    *refreshCacheFlag,
		Format:          *fmtFlag,
		TemplateOptions: templateOpts,
	}

	protoGen, err := gen.Generate(schema, *targetFlag, config)
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

	err = os.WriteFile(outfile, []byte(protoGen), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func collectCliArgs(flags *flag.FlagSet, args []string) (cliFlags []string, templateOpts map[string]interface{}, err error) {
	templateOpts = map[string]interface{}{}

	for _, arg := range args {
		name, value, _ := strings.Cut(arg, "=")
		if !strings.HasPrefix(name, "-") {
			return nil, nil, fmt.Errorf("option %q is invalid (expected -name=value)", arg)
		}
		name = strings.TrimLeft(name, "-")
		if len(name) == 0 {
			return nil, nil, fmt.Errorf("option %q is invalid (expected -name=value)", arg)
		}

		if flags.Lookup(name) != nil {
			cliFlags = append(cliFlags, arg)
		} else if name == "h" || name == "help" {
			cliFlags = append(cliFlags, arg)
			templateOpts["help"] = ""
		} else {
			templateOpts[name] = value
		}

		// Support webrpc-gen v0.6.0 -target=js -extra=noexports flag.
		if name == "extra" && value == "noexports" {
			templateOpts["export"] = "false"
		}
	}

	return
}
