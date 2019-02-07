package tests

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/webrpc/webrpc/lib/webrpc-go/tests/proto"
)

var (
	testServerTimeout         = time.Second * 60       // Max time the test server is allowed to run.
	testServerStartupTimeout  = time.Millisecond * 500 // Max time to wait for the test server to start up.
	testClientResponseTimeout = time.Second * 2        // Max time a client is going to wait for the server to respond.
	someCtx                   = context.Background()

	Server *http.Server
	Client proto.ExampleRPC
)

func init() {
	// Get next available port
	var port int
	{
		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			log.Fatal(err)
		}
		port = listener.Addr().(*net.TCPAddr).Port
		listener.Close()
		time.Sleep(time.Duration(250))
	}
	addr := fmt.Sprintf("0.0.0.0:%d", port)

	rpcErrs := make(chan error)
	go func() {
		rpcErrs <- http.ListenAndServe(addr, NewRPCHandler())
	}()
	rpcServerStartTimeout := time.NewTimer(testServerStartupTimeout)

	Client = proto.NewExampleRPCClient("http://"+addr, &http.Client{
		Timeout: testClientResponseTimeout,
	})

	start := time.Now()
	for {
		_, err := Client.Ping(context.Background())
		if err == nil {
			break
		}

		select {
		case <-rpcServerStartTimeout.C:
			log.Fatalf("server has failed to start within a reasonable time (giving up after waiting %v)", time.Now().Sub(start))
		default:
			continue
		}
	}
}
