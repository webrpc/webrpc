package main

import (
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
	clientFlag  = flags.Bool("client", false, "client mode")
	serverFlag  = flags.Bool("server", false, "server mode")
	versionFlag = flags.Bool("version", false, "print version and exit")
	printSchema = flags.Bool("print-schema", false, "print RIDL schema used for testing client/server communication")

	// -client
	clientFlags = flag.NewFlagSet("webrpc-test -client", flag.ExitOnError)
	urlFlag     = clientFlags.String("url", "http://localhost:9988", "run client against given server URL")

	// -server
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

	if *printSchema {
		fmt.Println(tests.GetSchema())
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

		err := client.RunTests(*urlFlag)
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
