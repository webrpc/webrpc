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

func assertHandlerCtx(t *testing.T, ctx context.Context, method, service, path string) {
	t.Helper()
	if got := MethodNameFromContext(ctx); got != method {
		t.Errorf("MethodNameFromContext = %q, want %q", got, method)
	}
	if got := ServiceNameFromContext(ctx); got != service {
		t.Errorf("ServiceNameFromContext = %q, want %q", got, service)
	}
	child, cancel := context.WithCancel(ctx)
	defer cancel()
	if got := MethodNameFromContext(child); got != method {
		t.Errorf("MethodNameFromContext(child) = %q, want %q", got, method)
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

// A path absent from methods reports no name.
func TestCtxNamesEmptyOnBadRoute(t *testing.T) {
	r := httptest.NewRequest("GET", "/rpc/NoSuchService/NoSuchMethod", nil)
	parent := context.Background()
	ctx := &reqCtx{Context: parent, w: httptest.NewRecorder(), r: r}
	if got := ServiceNameFromContext(ctx); got != "" {
		t.Errorf("ServiceNameFromContext on bad route = %q, want \"\"", got)
	}
	if got := MethodNameFromContext(ctx); got != "" {
		t.Errorf("MethodNameFromContext on bad route = %q, want \"\"", got)
	}
	if got := MethodNameFromContext(parent); got != "" {
		t.Errorf("MethodNameFromContext without a request = %q, want \"\"", got)
	}
}

// Injecting the request, not a name, lets MethodCtx resolve a real method.
func TestMethodCtxResolvesInjectedRequest(t *testing.T) {
	r := httptest.NewRequest("POST", "/rpc/TestApiRest/RpcMethod", nil)
	ctx := context.WithValue(context.Background(), HTTPRequestCtxKey, r)
	m, ok := MethodCtx(ctx)
	if !ok || m.Name() != "RpcMethod" || m.Service() != "TestApiRest" {
		t.Errorf("MethodCtx = %v, %v, want RpcMethod on TestApiRest", m, ok)
	}
}

// A request without a URL must not panic the name helpers.
func TestMethodCtxNilURL(t *testing.T) {
	ctx := context.WithValue(context.Background(), HTTPRequestCtxKey, &http.Request{})
	if m, ok := MethodCtx(ctx); ok || m != nil {
		t.Errorf("MethodCtx = %v, %v, want nil, false", m, ok)
	}
	if got := MethodNameFromContext(ctx); got != "" {
		t.Errorf("MethodNameFromContext = %q, want \"\"", got)
	}
	if got := ServiceNameFromContext(ctx); got != "" {
		t.Errorf("ServiceNameFromContext = %q, want \"\"", got)
	}
}
