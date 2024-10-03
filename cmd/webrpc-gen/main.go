package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/webrpc/webrpc"
	"github.com/webrpc/webrpc/gen"
	"github.com/webrpc/webrpc/schema"
)

var (
	flags            = flag.NewFlagSet("webrpc-gen", flag.ExitOnError)
	versionFlag      = flags.Bool("version", false, "print version and exit")
	schemaFlag       = flags.String("schema", "", "webrpc input schema file, ie. proto.ridl or proto.json (required)")
	targetFlag       = flags.String("target", "", targetUsage())
	outFlag          = flags.String("out", "", "generated output file (default stdout)")
	fmtFlag          = flags.Bool("fmt", true, "format generated code")
	refreshCacheFlag = flags.Bool("refreshCache", false, "refresh webrpc cache")
	silentFlag       = flags.Bool("silent", false, "silence gen summary")
	serviceFlag      = flags.String("service", "", "generate passed service only separated by comma, empty string means all services")
	ignoreFlag       = flags.String("ignore", "", "ignore ridl elements with specific annotations separated by commas")
)

func main() {
	// Collect CLI -flags and custom template -options.
	cliFlags, templateOpts, err := collectCliArgs(flags, os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse CLI flags: %v\n", err)
		os.Exit(1)
	}

	setTargetFlagsUsage(templateOpts)
	flags.Parse(cliFlags)

	if *versionFlag {
		fmt.Println("webrpc-gen", webrpc.VERSION)
		os.Exit(0)
	}

	if *schemaFlag == "" {
		fmt.Fprintf(os.Stderr, "-schema flag is required\n\n")
		flags.Usage()
		os.Exit(1)
	}

	// Parse+validate the webrpc schema file
	s, err := webrpc.ParseSchemaFile(*schemaFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse %s:%v\n", *schemaFlag, err)
		os.Exit(1)
	}

	// Code-gen targets
	if *targetFlag == "" {
		fmt.Fprintf(os.Stderr, "-target flag is required\n\n")
		flags.Usage()
		os.Exit(1)
	}

	if strings.Trim(*serviceFlag, " ") != "" {
		s = schema.MatchServices(s, strings.Split(*serviceFlag, ","))
	}

	if strings.Trim(*ignoreFlag, " ") != "" {
		ignoreAnnotations := strings.Split(*ignoreFlag, ",")

		ignoreAnnotationMap := map[string]struct{}{}
		for _, annonation := range ignoreAnnotations {
			_, annotationName, ok := strings.Cut(annonation, "@")
			if !ok {
				fmt.Fprintf(os.Stderr, "ignore annotations must start with @\n\n")
				os.Exit(1)
			}

			ignoreAnnotationMap[annotationName] = struct{}{}
		}

		s = schema.IgnoreMethodsWithAnnotations(s, ignoreAnnotationMap)
	}

	config := &gen.Config{
		RefreshCache:    *refreshCacheFlag,
		Format:          *fmtFlag,
		TemplateOptions: templateOpts,
	}

	genOutput, err := gen.Generate(s, *targetFlag, config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// Write output to stdout
	if *outFlag == "" || *outFlag == "stdout" {
		fmt.Println(genOutput.Code)
		os.Exit(0)
	}

	// Write output to a file
	err = writeOutfile(*outFlag, []byte(genOutput.Code))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// Print gen report
	if *silentFlag {
		os.Exit(0)
	}

	fmt.Println("=======================================")
	fmt.Println("|      webrpc generated summary       |")
	fmt.Println("=======================================")
	fmt.Println(" webrpc-gen version :", webrpc.VERSION)
	fmt.Println(" target             :", genOutput.TmplVersion)
	if !genOutput.IsLocal {
		fmt.Println(" target cache       :", genOutput.TmplDir)
		cacheAge := "now (refreshed)"
		if genOutput.CacheAge > 0 {
			cacheAge = fmt.Sprintf("%v", genOutput.CacheAge.Truncate(time.Second))
			if genOutput.CacheRefreshErr != nil {
				cacheAge += fmt.Sprintf(" (failed to refresh: %v)", genOutput.CacheRefreshErr)
			}
		}
		fmt.Println(" target cache age   :", cacheAge)
	}
	fmt.Println(" schema file        :", *schemaFlag)
	fmt.Println(" output file        :", *outFlag)
	if genOutput.FormatErr != nil {
		fmt.Println(" format error       :", genOutput.FormatErr)
		os.Exit(1)
	}
}

func collectCliArgs(flags *flag.FlagSet, args []string) (cliFlags []string, templateOpts map[string]interface{}, err error) {
	templateOpts = map[string]interface{}{}

	for _, arg := range args {
		//name, value, _ := strings.Cut(arg, "=") // Added in Go 1.18.
		name, value := arg, ""
		if i := strings.Index(arg, "="); i >= 0 {
			name = arg[:i]
			value = arg[i+1:]
		}

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

func writeOutfile(outfile string, protoGen []byte) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if outfile[0:1] != "/" {
		outfile = filepath.Join(cwd, outfile)
	}

	outdir := filepath.Dir(outfile)
	if _, err := os.Stat(outdir); os.IsNotExist(err) {
		err := os.MkdirAll(outdir, 0755)
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(outfile, []byte(protoGen), 0644)
	if err != nil {
		return err
	}

	return nil
}

func targetUsage() string {
	var b strings.Builder
	fmt.Fprintln(&b, "target code generator (required)")
	fmt.Fprintln(&b, "built-in targets:")
	for _, target := range gen.EmbeddedTargetNames {
		fmt.Fprintf(&b, "  -target=%s\n", target)
	}
	fmt.Fprintln(&b, "  -target=json (prints schema in JSON)")
	fmt.Fprintln(&b, "  -target=debug (prints schema and template variables incl. Go type information)")
	fmt.Fprintln(&b, "remote git repo:")
	for _, target := range gen.EmbeddedTargetNames {
		fmt.Fprintf(&b, "  -target=%s\n", gen.EmbeddedTargets[target].ImportTag)
	}
	fmt.Fprintln(&b, "local folder:")
	fmt.Fprintln(&b, "  -target=../local-go-templates-on-disk")
	fmt.Fprintln(&b, "  (see https://github.com/webrpc/webrpc/tree/master/gen)")

	return b.String()
}

func setTargetFlagsUsage(templateOpts map[string]interface{}) {
	flags.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -schema=<file.(ridl|json)> -target=<target> -out=<file> [...targetOpts]\n", flags.Name())
		flags.PrintDefaults()
		fmt.Fprintf(os.Stderr, "See https://github.com/webrpc/webrpc for more info.\n")

		if *targetFlag != "" {
			fmt.Fprintf(os.Stderr, "\nTarget generator usage:\n")
			templateHelp, err := gen.Generate(&schema.WebRPCSchema{}, *targetFlag, &gen.Config{TemplateOptions: templateOpts})
			if err != nil {
				fmt.Fprintln(os.Stderr, templateHelp.Code)
			} else {
				fmt.Fprintf(os.Stderr, "failed to render -help: %v\n", err)
			}
		}
	}
}
