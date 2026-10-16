package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webrpc/webrpc/tests/server"
)

// discardWriter is a reusable ResponseWriter, so the benchmark measures the server alone.
type discardWriter struct {
	header http.Header
}

func (w *discardWriter) Header() http.Header         { return w.header }
func (w *discardWriter) WriteHeader(int)             {}
func (w *discardWriter) Write(p []byte) (int, error) { return len(p), nil }

// newGetItemRequest returns a server with item 7 and a GET request for it.
func newGetItemRequest(tb testing.TB) (http.Handler, *http.Request, *discardWriter) {
	impl := server.NewTestRESTServer()
	if _, err := impl.CreateItem(context.Background(), server.CreateItemRequest{Item: &server.Item{Id: 7, Name: "alice"}}); err != nil {
		tb.Fatal(err)
	}
	srv := server.NewTestApiRestServer(impl)
	return srv, httptest.NewRequest("GET", "/rpc/items/7", nil), &discardWriter{header: http.Header{}}
}

// TestRestGetItemAllocs guards the allocation budget that BenchmarkRestGetItem reports.
func TestRestGetItemAllocs(t *testing.T) {
	srv, req, w := newGetItemRequest(t)
	allocs := testing.AllocsPerRun(100, func() {
		clear(w.header)
		srv.ServeHTTP(w, req)
	})
	if allocs > 14 {
		t.Errorf("GET /rpc/items/7: %v allocs per request, want <= 14", allocs)
	}
}

func BenchmarkRestGetItem(b *testing.B) {
	srv, req, w := newGetItemRequest(b)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		clear(w.header)
		srv.ServeHTTP(w, req)
	}
}
