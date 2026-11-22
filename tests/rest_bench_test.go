package tests

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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
	if allocs > 8 {
		t.Errorf("GET /rpc/items/7: %v allocs per request, want <= 8", allocs)
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

// BenchmarkRestGetChildren exercises the two-path-param route.
func BenchmarkRestGetChildren(b *testing.B) {
	srv := server.NewTestApiRestServer(server.NewTestRESTServer())
	req := httptest.NewRequest("GET", "/rpc/items/7/children/3", nil)
	w := &discardWriter{header: http.Header{}}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		clear(w.header)
		srv.ServeHTTP(w, req)
	}
}

// BenchmarkRestCreateItem exercises a body route; the body must be fresh per
// iteration, so the request is rebuilt each time (same cost before and after).
func BenchmarkRestCreateItem(b *testing.B) {
	srv := server.NewTestApiRestServer(server.NewTestRESTServer())
	w := &discardWriter{header: http.Header{}}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "/rpc/items", strings.NewReader(`{"item":{"id":7,"name":"alice"}}`))
		req.Header.Set("Content-Type", "application/json")
		clear(w.header)
		srv.ServeHTTP(w, req)
	}
}

// BenchmarkRPCGetEmpty guards the plain RPC dispatch path against regressions.
func BenchmarkRPCGetEmpty(b *testing.B) {
	srv := server.NewTestApiServer(&server.TestServer{})
	w := &discardWriter{header: http.Header{}}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "/rpc/TestApi/GetEmpty", strings.NewReader("{}"))
		req.Header.Set("Content-Type", "application/json")
		clear(w.header)
		srv.ServeHTTP(w, req)
	}
}

// BenchmarkRpcPost measures the plain RPC POST path with a rewound body per request.
func BenchmarkRpcPost(b *testing.B) {
	srv := server.NewTestApiRestServer(newRpcPostImpl(b))
	benchRpcPost(b, srv, context.Background())
}

// BenchmarkRpcPostReadCtx adds per-request ctx reads in the service method.
func BenchmarkRpcPostReadCtx(b *testing.B) {
	srv := server.NewTestApiRestServer(readCtxServer{newRpcPostImpl(b)})
	benchRpcPost(b, srv, outerValueCtx())
}

// BenchmarkRestGetItemReadCtx adds per-request ctx reads in the service method.
func BenchmarkRestGetItemReadCtx(b *testing.B) {
	impl := server.NewTestRESTServer()
	if _, err := impl.CreateItem(context.Background(), server.CreateItemRequest{Item: &server.Item{Id: 7, Name: "alice"}}); err != nil {
		b.Fatal(err)
	}
	srv := server.NewTestApiRestServer(readCtxServer{impl})
	req := httptest.NewRequest("GET", "/rpc/items/7", nil).WithContext(outerValueCtx())
	w := &discardWriter{header: http.Header{}}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		clear(w.header)
		srv.ServeHTTP(w, req)
	}
}

// benchUnrelatedKey is an outer middleware key that webrpc does not know about.
type benchUnrelatedKey struct{}

// benchCtxSink keeps the ctx reads observable so the compiler cannot drop them.
var benchCtxSink int

// readCtxServer reads the handler ctx once per call, like user code on the request path.
type readCtxServer struct {
	*server.TestRESTServer
}

func (s readCtxServer) GetItem(ctx context.Context, req server.GetItemRequest) (*server.GetItemResponse, error) {
	readBenchCtx(ctx)
	return s.TestRESTServer.GetItem(ctx, req)
}

func (s readCtxServer) RpcMethod(ctx context.Context, id uint64) (*server.Item, error) {
	readBenchCtx(ctx)
	return s.TestRESTServer.RpcMethod(ctx, id)
}

func readBenchCtx(ctx context.Context) {
	if _, ok := server.MethodCtx(ctx); ok {
		benchCtxSink++
	}
	benchCtxSink += len(server.ServiceNameFromContext(ctx)) + len(server.MethodNameFromContext(ctx))
	if server.RequestFromContext(ctx) != nil {
		benchCtxSink++
	}
	if server.ResponseWriterFromContext(ctx) != nil {
		benchCtxSink++
	}
	if v, _ := ctx.Value(benchUnrelatedKey{}).(string); v != "" {
		benchCtxSink++
	}
}

// outerValueCtx stands in for an outer middleware value, set once outside the loop.
func outerValueCtx() context.Context {
	return context.WithValue(context.Background(), benchUnrelatedKey{}, "outer")
}

func newRpcPostImpl(tb testing.TB) *server.TestRESTServer {
	impl := server.NewTestRESTServer()
	if _, err := impl.CreateItem(context.Background(), server.CreateItemRequest{Item: &server.Item{Id: 7, Name: "alice"}}); err != nil {
		tb.Fatal(err)
	}
	return impl
}

func benchRpcPost(b *testing.B, srv http.Handler, ctx context.Context) {
	body := bytes.NewReader([]byte(`{"id":7}`))
	req := httptest.NewRequest("POST", "/rpc/TestApiRest/RpcMethod", body).WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	w := &discardWriter{header: http.Header{}}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = body.Seek(0, io.SeekStart)
		clear(w.header)
		srv.ServeHTTP(w, req)
	}
}
