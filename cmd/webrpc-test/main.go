package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/webrpc/webrpc"
	"github.com/webrpc/webrpc/tests"
	"github.com/webrpc/webrpc/tests/client"
	"github.com/webrpc/webrpc/tests/server"
)

var (
	flags       = flag.NewFlagSet("webrpc-test", flag.ContinueOnError)
	clientFlag  = flags.Bool("client", false, "run client tests")
	serverFlag  = flags.Bool("server", false, "run test server")
	versionFlag = flags.Bool("version", false, "print version and exit")
	printSchema = flags.Bool("print-schema", false, "obsolete flag (use -printRIDL)") // Obsolete.
	printRIDL   = flags.Bool("printRIDL", false, "print schema in RIDL")
	printJSON   = flags.Bool("printJSON", false, "print schema in JSON")

	// webrpc-test -client -url=http://localhost:9988
	clientFlags = flag.NewFlagSet("webrpc-test -client", flag.ExitOnError)
	urlFlag     = clientFlags.String("url", "http://localhost:9988", "run client against given server URL")

	// webrpc-test -server -port=9988 -timeout=1m
	serverFlags = flag.NewFlagSet("webrpc-test -server", flag.ExitOnError)
	portFlag    = serverFlags.Int("port", 9988, "run server at given port")
	timeoutFlag = serverFlags.Duration("timeout", time.Minute, "exit after given timeout")
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "-server or -client flag is required\n")
		os.Exit(1)
	}

	_ = flags.Parse(os.Args[1:2])

	if *versionFlag {
		fmt.Println("webrpc-test", webrpc.VERSION)
		os.Exit(0)
	}

	if *printRIDL || *printSchema {
		fmt.Println(tests.GetRIDLSchema())
		os.Exit(0)
	}

	if *printJSON {
		fmt.Println(tests.GetJSONSchema())
		os.Exit(0)
	}

	if !*serverFlag && !*clientFlag {
		fmt.Fprintf(os.Stderr, "-server or -client flag is required\n")
		os.Exit(1)
	}

	if *clientFlag {
		if err := clientFlags.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}

		err := client.RunTests(context.Background(), *urlFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	if err := serverFlags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	server, err := server.RunTestServer(fmt.Sprintf("0.0.0.0:%v", *portFlag), *timeoutFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	if err := server.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
