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

func BenchmarkRestGetItem(b *testing.B) {
	impl := server.NewTestRESTServer()
	if _, err := impl.CreateItem(context.Background(), server.CreateItemRequest{Item: &server.Item{Id: 7, Name: "alice"}}); err != nil {
		b.Fatal(err)
	}
	srv := server.NewTestApiRestServer(impl)
	req := httptest.NewRequest("GET", "/rpc/items/7", nil)
	w := &discardWriter{header: http.Header{}}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		clear(w.header)
		srv.ServeHTTP(w, req)
	}
}
