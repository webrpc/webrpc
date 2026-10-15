package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// recordingServer captures what the generated REST handlers pass to the implementation.
type recordingServer struct {
	*TestRESTServer
	listReq ListItemsRequest
	method  *method
}

func (s *recordingServer) ListItems(ctx context.Context, req ListItemsRequest) (*ListItemsResponse, error) {
	s.listReq = req
	s.method, _ = MethodCtx(ctx)
	return s.TestRESTServer.ListItems(ctx, req)
}

func serveTest(t *testing.T, h http.Handler, method, target, contentType, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(method, target, strings.NewReader(body))
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestRESTUncleanPathIsNotRedirected(t *testing.T) {
	srv := NewTestApiRestServer(NewTestRESTServer())
	for _, target := range []string{"/rpc/items/../items/7", "/rpc/items//7", "/rpc/items/./7"} {
		rec := serveTest(t, srv, "GET", target, "", "")
		if rec.Code != http.StatusNotFound {
			t.Errorf("GET %s: status %d, want 404", target, rec.Code)
		}
		if loc := rec.Header().Get("Location"); loc != "" {
			t.Errorf("GET %s: unexpected redirect to %q", target, loc)
		}
	}
}

func TestRESTAllowListsHeadWithGet(t *testing.T) {
	srv := NewTestApiRestServer(NewTestRESTServer())
	rec := serveTest(t, srv, "POST", "/rpc/items/7", "application/json", "{}")
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("POST /rpc/items/7: status %d, want 405", rec.Code)
	}
	allow := rec.Header().Get("Allow")
	if !strings.Contains(allow, "GET, HEAD") {
		t.Errorf("Allow = %q, want GET and HEAD together", allow)
	}
}

func TestRESTOnRequestSeesRESTRequest(t *testing.T) {
	var verb, path string
	srv := NewTestApiRestServer(NewTestRESTServer(), &Options{
		OnRequest: func(w http.ResponseWriter, r *http.Request) error {
			verb, path = r.Method, r.URL.Path
			return nil
		},
	})
	serveTest(t, srv, "PUT", "/rpc/items/7", "application/json", `{"name":"x"}`)
	if verb != "PUT" || path != "/rpc/items/7" {
		t.Errorf("OnRequest saw %s %s, want PUT /rpc/items/7", verb, path)
	}
}

func TestRESTNullBodyKeepsPathParam(t *testing.T) {
	impl := NewTestRESTServer()
	if _, err := impl.CreateItem(context.Background(), CreateItemRequest{Item: &Item{Id: 7, Name: "alice"}}); err != nil {
		t.Fatal(err)
	}
	srv := NewTestApiRestServer(impl)
	rec := serveTest(t, srv, "PATCH", "/rpc/items/7", "application/json", "null")
	if rec.Code != http.StatusOK {
		t.Fatalf("PATCH with null body: status %d, body %s", rec.Code, rec.Body)
	}
	if !strings.Contains(rec.Body.String(), `"name":"alice"`) {
		t.Errorf("PATCH with null body: unexpected response %s", rec.Body)
	}
}

func TestRESTMissingRequiredListBindsEmptySlice(t *testing.T) {
	impl := &recordingServer{TestRESTServer: NewTestRESTServer()}
	srv := NewTestApiRestServer(impl)
	rec := serveTest(t, srv, "GET", "/rpc/items?limit=1", "", "")
	if rec.Code != http.StatusOK {
		t.Fatalf("GET /rpc/items?limit=1: status %d, body %s", rec.Code, rec.Body)
	}
	if impl.listReq.Tags == nil || len(impl.listReq.Tags) != 0 {
		t.Errorf("missing required list bound to %#v, want empty non-nil slice", impl.listReq.Tags)
	}
	if impl.method == nil || impl.method.Name() != "ListItems" {
		t.Errorf("MethodCtx = %v, want method ListItems", impl.method)
	}
}

func TestRESTScalarOverflowIsBadRequest(t *testing.T) {
	type overflowRequest struct{ N uint8 }
	srv := NewTestApiRestServer(NewTestRESTServer())
	mux := http.NewServeMux()
	mux.Handle("GET /n/{n}", restHandler("Overflow", false,
		func(ctx context.Context, req overflowRequest) (struct{}, error) { return struct{}{}, nil },
		func(p *restParams, req *overflowRequest) error { return restPath(p, "n", "uint8", &req.N, restUint) },
		srv.sendErrorJSON, &srv.OnRequest))

	if rec := serveTest(t, mux, "GET", "/n/300", "", ""); rec.Code != http.StatusBadRequest {
		t.Errorf("GET /n/300 (uint8): status %d, want 400", rec.Code)
	}
	if rec := serveTest(t, mux, "GET", "/n/255", "", ""); rec.Code != http.StatusOK {
		t.Errorf("GET /n/255 (uint8): status %d, want 200", rec.Code)
	}
}

func TestRESTHeadOnGetRouteHasNoBody(t *testing.T) {
	impl := NewTestRESTServer()
	if _, err := impl.CreateItem(context.Background(), CreateItemRequest{Item: &Item{Id: 7, Name: "alice"}}); err != nil {
		t.Fatal(err)
	}
	rec := serveTest(t, NewTestApiRestServer(impl), "HEAD", "/rpc/items/7", "", "")
	if rec.Code != http.StatusOK {
		t.Errorf("HEAD /rpc/items/7: status %d, want 200", rec.Code)
	}
	if rec.Body.Len() != 0 {
		t.Errorf("HEAD /rpc/items/7: body %q, want none", rec.Body)
	}
}
