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
	waitFlag    = flags.Bool("waitForServer", false, "wait for server to be ready")
	versionFlag = flags.Bool("version", false, "print version and exit")
	printSchema = flags.Bool("print-schema", false, "obsolete flag (use -printRIDL)") // Obsolete.
	printRIDL   = flags.Bool("printRIDL", false, "print schema in RIDL")
	printJSON   = flags.Bool("printJSON", false, "print schema in JSON")

	// webrpc-test -client -url=http://localhost:9988
	clientFlags   = flag.NewFlagSet("webrpc-test -client", flag.ExitOnError)
	clientUrlFlag = clientFlags.String("url", "http://localhost:9988", "run client against given server URL")

	// webrpc-test -server -port=9988 -timeout=1m
	serverFlags       = flag.NewFlagSet("webrpc-test -server", flag.ExitOnError)
	serverPortFlag    = serverFlags.Int("port", 9988, "run server at given port")
	serverTimeoutFlag = serverFlags.Duration("timeout", time.Minute, "exit after given timeout")

	// webrpc-test -waitForServer -url=http://localhost:9988 -timeout=1m
	waitFlags       = flag.NewFlagSet("webrpc-test -waitForServer", flag.ExitOnError)
	waitUrlFlag     = waitFlags.String("url", "http://localhost:9988", "run client against given server URL")
	waitTimeoutFlag = waitFlags.Duration("timeout", time.Minute, "exit after given timeout")
)

func main() {
	if len(os.Args) >= 2 {
		if err := flags.Parse(os.Args[1:2]); err != nil {
			os.Exit(1)
		}
	}

	switch {
	case *versionFlag:
		fmt.Println("webrpc-test", webrpc.VERSION)
		os.Exit(0)

	case *printRIDL, *printSchema:
		fmt.Println(tests.GetRIDLSchema())
		os.Exit(0)

	case *printJSON:
		fmt.Println(tests.GetJSONSchema())
		os.Exit(0)

	case *clientFlag:
		if err := clientFlags.Parse(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := client.RunTests(context.Background(), *clientUrlFlag); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		os.Exit(0)

	case *serverFlag:
		if err := serverFlags.Parse(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		server, err := server.RunTestServer(fmt.Sprintf("0.0.0.0:%v", *serverPortFlag), *serverTimeoutFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := server.Wait(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		os.Exit(0)

	case *waitFlag:
		if err := waitFlags.Parse(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		start := time.Now()
		if err := client.Wait(*waitUrlFlag, *waitTimeoutFlag); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stderr, "wait: test server ready in %v\n", time.Since(start).Round(time.Millisecond))
		os.Exit(0)

	default:
		flags.Usage()
		os.Exit(1)
	}
}
