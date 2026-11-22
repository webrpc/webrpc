package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ctxProbe checks the handler ctx of two methods, then delegates as before.
type ctxProbe struct {
	*TestRESTServer
	check func(ctx context.Context, method string)
}

func (p *ctxProbe) GetItem(ctx context.Context, req GetItemRequest) (*GetItemResponse, error) {
	p.check(ctx, "GetItem")
	return p.TestRESTServer.GetItem(ctx, req)
}

func (p *ctxProbe) RpcMethod(ctx context.Context, id uint64) (*Item, error) {
	p.check(ctx, "RpcMethod")
	return p.TestRESTServer.RpcMethod(ctx, id)
}

type unrelatedKey struct{}

// serveWithOuterValue wraps h with a value under a key webrpc does not know about.
func serveWithOuterValue(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), unrelatedKey{}, "outer"))
		h.ServeHTTP(w, r)
	})
}

// assertHandlerCtx reads the exported keys straight from ctx, not via the accessors.
func assertHandlerCtx(t *testing.T, ctx context.Context, method, service, path string) {
	t.Helper()
	if got := ctx.Value(MethodNameCtxKey); got != method {
		t.Errorf("ctx.Value(MethodNameCtxKey) = %v, want %q", got, method)
	}
	if got := ctx.Value(ServiceNameCtxKey); got != service {
		t.Errorf("ctx.Value(ServiceNameCtxKey) = %v, want %q", got, service)
	}
	if r, _ := ctx.Value(HTTPRequestCtxKey).(*http.Request); r == nil || r.URL.Path != path {
		t.Errorf("ctx.Value(HTTPRequestCtxKey) = %v, want request for %q", r, path)
	}
	if w, _ := ctx.Value(HTTPResponseWriterCtxKey).(http.ResponseWriter); w == nil {
		t.Error("ctx.Value(HTTPResponseWriterCtxKey) is nil")
	}
	if got := ctx.Value(unrelatedKey{}); got != "outer" {
		t.Errorf("outer middleware value = %v, want \"outer\"", got)
	}
	if got := MethodNameFromContext(context.WithValue(ctx, MethodNameCtxKey, "override")); got != "override" {
		t.Errorf("MethodNameFromContext on derived ctx = %q, want \"override\"", got)
	}
}

func TestHandlerCtxKeepsWebrpcValues(t *testing.T) {
	calls := 0
	probe := &ctxProbe{TestRESTServer: NewTestRESTServer()}
	probe.check = func(ctx context.Context, method string) {
		calls++
		switch method {
		case "GetItem":
			assertHandlerCtx(t, ctx, "GetItem", "TestApiRest", "/rpc/items/7")
		case "RpcMethod":
			assertHandlerCtx(t, ctx, "RpcMethod", "TestApiRest", "/rpc/TestApiRest/RpcMethod")
		}
	}
	srv := serveWithOuterValue(NewTestApiRestServer(probe))

	serveTest(t, srv, "GET", "/rpc/items/7", "", "")
	serveTest(t, srv, "POST", "/rpc/TestApiRest/RpcMethod", "application/json", `{"id":7}`)
	if calls != 2 {
		t.Errorf("handler ran %d times, want 2", calls)
	}
}

func TestHandlerCtxCancelsWithRequest(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())
	defer cancel()
	probe := &ctxProbe{TestRESTServer: NewTestRESTServer()}
	probe.check = func(ctx context.Context, _ string) {
		cancel()
		select {
		case <-ctx.Done():
		default:
			t.Error("handler ctx not cancelled when the request ctx was cancelled")
		}
	}
	req := httptest.NewRequest("GET", "/rpc/items/7", nil).WithContext(parent)
	NewTestApiRestServer(probe).ServeHTTP(httptest.NewRecorder(), req)
}
