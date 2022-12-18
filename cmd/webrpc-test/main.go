package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/webrpc/webrpc/tests/client"
	"github.com/webrpc/webrpc/tests/server"
)

var (
	flags      = flag.NewFlagSet("webrpc-test", flag.ContinueOnError)
	clientFlag = flags.Bool("client", false, "client mode")
	serverFlag = flags.Bool("server", false, "server mode")

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
			log.Fatal(err)
		}

		fmt.Printf("All tests ran successfully against %v\n", *urlFlag)
		os.Exit(0)
	}

	if err := serverFlags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	err := serverMode()
	if err != nil {
		log.Fatal(err)
	}
}

func serverMode() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", *portFlag),
		Handler: server.NewComplexApiServer(&server.ComplexServer{}),
	}

	if *timeoutFlag > 0 {
		go func() {
			timeoutCtx, cancel := context.WithTimeout(context.Background(), *timeoutFlag)
			defer cancel()

			select {
			case <-timeoutCtx.Done():
				gracefulShutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				srv.Shutdown(gracefulShutdownCtx)
			}
		}()
	}

	return srv.ListenAndServe()
}
