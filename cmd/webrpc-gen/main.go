package main

import (
	"flag"
	"fmt"
	"gitlab.com/logement3d/webrpc"
	"gitlab.com/logement3d/webrpc/gen"
	_ "gitlab.com/logement3d/webrpc/gen/golang"
	_ "gitlab.com/logement3d/webrpc/gen/javascript"
	_ "gitlab.com/logement3d/webrpc/gen/typescript"
	schema2 "gitlab.com/logement3d/webrpc/schema"
	"io/ioutil"
	"os"
	"path/filepath"
)

var flags = flag.NewFlagSet("webrpc-gen", flag.ExitOnError)

func main() {
	versionFlag := flags.Bool("version", false, "print webrpc version and exit")
	schemaFlag := flags.String("schema", "", "webrpc schema file (required)")
	pkgFlag := flags.String("pkg", "proto", "generated package name for target language, default: proto")
	outFlag := flags.String("out", "", "generated output file, default: stdout")
	testFlag := flags.Bool("test", false, "test schema parsing (skips code-gen)")
	clientFlag := flags.Bool("client", false, "enable webrpc client library generation, default: off")
	serverFlag := flags.Bool("server", false, "enable webrpc server library generation, default: off")

	// registered/available target languages
	targets := []string{}
	for k, _ := range gen.Generators {
		targets = append(targets, k)
	}
	targetFlag := flags.String("target", "", fmt.Sprintf("target language for webrpc library generation, %s (required)", targets))
	targetExtra := flags.String("extra", "", "target language extra/custom options")

	flags.Parse(os.Args[1:])

	if *versionFlag {
		fmt.Printf("webrpc %s\n", webrpc.VERSION)
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

	cpy, err := webrpc.ParseSchemaFile(*schemaFlag)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	cpy2, err := webrpc.ParseSchemaFile(*schemaFlag)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, msg := range cpy.Messages {
		msg.Name = "Partial" + msg.Name
		for _, field := range msg.Fields {
			field.Optional = true
		}
		keys := schema2.MessageField{
			Name: "keys",
			Type: &schema2.VarType{
				Type: schema2.T_List,
				List: &schema2.VarListType{
					Elem: &schema2.VarType{
						Type:   schema2.T_String,
						List:   nil,
						Map:    nil,
						Struct: nil,
					},
				},
				Map:    nil,
				Struct: nil,
			},
			Optional: false,
			Value:    "",
			Meta:     nil,
		}
		msg.Fields = append(msg.Fields, &keys)
	}

	schema.Messages = []*schema2.Message{}

	for i, msg := range cpy.Messages {
		schema.Messages = append(schema.Messages, cpy2.Messages[i])
		schema.Messages = append(schema.Messages, msg)
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

	targetLang := *targetFlag
	if _, ok := gen.Generators[targetLang]; !ok {
		fmt.Printf("oops, you passed an invalid -target flag, try one of registered generators: %s\n", targets)
		os.Exit(1)
	}

	// Call our target code-generator
	generator := gen.GetGenerator(*targetFlag)
	if generator == nil {
		fmt.Printf("error! unable to find generator for target '%s'\n", *targetFlag)
		os.Exit(1)
	}

	targetOpts := gen.TargetOptions{
		PkgName: *pkgFlag,
		Client:  *clientFlag,
		Server:  *serverFlag,
		Extra:   *targetExtra,
	}

	protoGen, err := generator.Gen(schema, targetOpts)
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
