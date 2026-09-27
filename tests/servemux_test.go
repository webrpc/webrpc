package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webrpc/webrpc/tests/client"
	"github.com/webrpc/webrpc/tests/server"
)

// TestServeMuxMounting mounts every method of both services onto a plain Go
// 1.22+ http.ServeMux via the patterns returned by Methods() — REST routes as
// "GET /rpc/items/{id}", plain RPC methods as their bare POST path — and runs
// both full client suites through the mux.
//
// Caveat this test accepts: requests rejected by the mux itself (unknown
// path, wrong verb) get Go's plain-text 404/405 instead of the webrpc JSON
// error envelope; the envelope applies once the mux has matched a pattern.
func TestServeMuxMounting(t *testing.T) {
	mux := http.NewServeMux()
	srv := server.Server{
		TestApi:     &server.TestServer{},
		TestApiRest: server.NewTestRESTServer(),
	}
	for _, m := range srv.Methods(nil) {
		mux.Handle(m.Path, m.Handler)
	}

	ts := httptest.NewServer(mux)
	defer ts.Close()

	if err := client.RunTests(context.Background(), ts.URL); err != nil {
		t.Errorf("RPC suite via mux: %v", err)
	}
	if err := client.RunRESTTests(context.Background(), ts.URL); err != nil {
		t.Errorf("REST suite via mux: %v", err)
	}
}
