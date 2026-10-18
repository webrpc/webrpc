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

// TestRESTContentTypeRelaxation pins the relaxed Content-Type rules: an empty
// body bypasses the check on any body route, a route with zero body fields
// drops the check entirely, and a route with body fields keeps it as a CSRF
// guard for non-empty bodies.
func TestRESTContentTypeRelaxation(t *testing.T) {
	impl := NewTestRESTServer()
	if _, err := impl.CreateItem(context.Background(), CreateItemRequest{Item: &Item{Id: 7, Name: "alice"}}); err != nil {
		t.Fatal(err)
	}
	srv := NewTestApiRestServer(impl)

	// (b) POST with body fields, empty body, and no Content-Type: accepted with a zero-value request.
	if rec := serveTest(t, srv, "PATCH", "/rpc/items/7", "", ""); rec.Code != http.StatusOK {
		t.Errorf("PATCH /rpc/items/7 with empty body and no Content-Type: status %d, body %s, want 200", rec.Code, rec.Body)
	} else if !strings.Contains(rec.Body.String(), `"name":"alice"`) {
		t.Errorf("PATCH /rpc/items/7 with empty body: body %s, want unchanged item", rec.Body)
	}

	// (c) POST with body fields, non-empty body, and a wrong Content-Type: rejected (CSRF guard intact).
	if rec := serveTest(t, srv, "PATCH", "/rpc/items/7", "text/plain", `{"name":"csrf"}`); rec.Code != http.StatusBadRequest {
		t.Errorf("PATCH /rpc/items/7 with text/plain body: status %d, want 400", rec.Code)
	}

	// (d) POST with body fields, valid JSON, and the correct Content-Type: unchanged.
	if rec := serveTest(t, srv, "PATCH", "/rpc/items/7", "application/json", `{"name":"bob"}`); rec.Code != http.StatusOK {
		t.Errorf("PATCH /rpc/items/7 with JSON body: status %d, body %s, want 200", rec.Code, rec.Body)
	}

	// (a) POST with zero body fields and no Content-Type: accepted.
	if rec := serveTest(t, srv, "POST", "/rpc/items/reset", "", ""); rec.Code != http.StatusOK {
		t.Errorf("POST /rpc/items/reset with no Content-Type: status %d, want 200", rec.Code)
	}

	// A zero-body-field route with a non-empty body parses it; unknown fields are ignored.
	if rec := serveTest(t, srv, "POST", "/rpc/items/reset", "text/plain", `{"unknown":1}`); rec.Code != http.StatusOK {
		t.Errorf("POST /rpc/items/reset with text/plain body: status %d, want 200", rec.Code)
	}
	// But a malformed body still 400s even on a zero-body-field route.
	if rec := serveTest(t, srv, "POST", "/rpc/items/reset", "text/plain", `{"unknown":`); rec.Code != http.StatusBadRequest {
		t.Errorf("POST /rpc/items/reset with malformed body: status %d, want 400", rec.Code)
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
	mux.Handle("GET /n/{n}", restHandler("GET /n/{n}", "", "Overflow", false, false, 200,
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

// TestRESTDeclaredSuccessStatus pins `+ @success 201`: the route writes the
// declared 2xx status on success, while routes without @success keep 200.
func TestRESTDeclaredSuccessStatus(t *testing.T) {
	impl := NewTestRESTServer()
	srv := NewTestApiRestServer(impl)

	rec := serveTest(t, srv, "POST", "/rpc/items", "application/json", `{"item":{"id":7,"name":"alice"}}`)
	if rec.Code != http.StatusCreated {
		t.Errorf("POST /rpc/items: status %d, want 201 (@success)", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), `"name":"alice"`) {
		t.Errorf("POST /rpc/items: unexpected response %s", rec.Body)
	}

	// A route without @success keeps the default 200.
	rec = serveTest(t, srv, "GET", "/rpc/items/7", "", "")
	if rec.Code != http.StatusOK {
		t.Errorf("GET /rpc/items/7: status %d, want 200", rec.Code)
	}
}

func TestRESTAliasRoutesServeSameMethod(t *testing.T) {
	impl := NewTestRESTServer()
	impl.items[7] = &Item{Id: 7, Name: "alice"}
	srv := NewTestApiRestServer(impl)

	for _, tc := range []struct{ method, target, contentType, body string }{
		{"GET", "/rpc/items/lookup/7", "", ""},
		{"POST", "/rpc/TestApiRest/LookupItem", "application/json", `{"id":7}`},
		{"GET", "/rpc/items/find?q=ali", "", ""},
		{"QUERY", "/rpc/items/find", "application/json", `{"q":"ali"}`},
	} {
		rec := serveTest(t, srv, tc.method, tc.target, tc.contentType, tc.body)
		if rec.Code != 200 || !strings.Contains(rec.Body.String(), `"name":"alice"`) {
			t.Errorf("%s %s: status %d, body %s", tc.method, tc.target, rec.Code, rec.Body)
		}
	}

	// A body verb on the GET-only alias path is not served.
	if rec := serveTest(t, srv, "POST", "/rpc/items/find", "application/json", `{"q":"ali"}`); rec.Code != 405 {
		t.Errorf("POST /rpc/items/find: status %d, want 405", rec.Code)
	}
}

// TestRESTTrailingLiteralRoute pins the route shape with a literal segment
// after a path param: the suffix must match exactly, the param must bind, and
// the handler must run on the service receiver (the shadowed-receiver bug).
func TestRESTTrailingLiteralRoute(t *testing.T) {
	impl := NewTestRESTServer()
	if _, err := impl.CreateItem(context.Background(), CreateItemRequest{Item: &Item{Id: 7, Name: "alice"}}); err != nil {
		t.Fatal(err)
	}
	srv := NewTestApiRestServer(impl)

	rec := serveTest(t, srv, "GET", "/rpc/items/7/status/latest", "", "")
	if rec.Code != http.StatusOK {
		t.Fatalf("GET /rpc/items/7/status/latest: status %d, body %s", rec.Code, rec.Body)
	}
	if want := `"status":"latest status of item 7: alice"`; !strings.Contains(rec.Body.String(), want) {
		t.Errorf("GET /rpc/items/7/status/latest: body %s, want containing %s", rec.Body, want)
	}

	// A different literal after the param is a different path.
	if rec := serveTest(t, srv, "GET", "/rpc/items/7/wrong", "", ""); rec.Code != http.StatusNotFound {
		t.Errorf("GET /rpc/items/7/wrong: status %d, want 404", rec.Code)
	}

	// A wrong method still resolves the path to its Allow set.
	rec = serveTest(t, srv, "DELETE", "/rpc/items/7/status/latest", "", "")
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("DELETE /rpc/items/7/status/latest: status %d, want 405", rec.Code)
	}
	if allow := rec.Header().Get("Allow"); allow != "GET, HEAD" {
		t.Errorf("DELETE /rpc/items/7/status/latest: Allow = %q, want GET, HEAD", allow)
	}
}

func TestRESTWebrpcMethodsHasEntryPerRoute(t *testing.T) {
	m := WebrpcMethods(ServiceTestApiRest)
	for _, key := range []string{"GET /rpc/items/lookup/{id}", "POST /rpc/TestApiRest/LookupItem", "GET /rpc/items/find", "QUERY /rpc/items/find"} {
		if _, ok := m[key]; !ok {
			t.Errorf("WebrpcMethods() missing %q", key)
		}
	}
}

// TestRESTEncodedSlashDoesNotHitStaticRoute pins ServeMux's segment rule: a
// %2F inside a segment is not a path separator, so the encoded form must not
// unwrap into a static route match.
func TestRESTEncodedSlashDoesNotHitStaticRoute(t *testing.T) {
	srv := NewTestApiRestServer(NewTestRESTServer())
	for _, target := range []string{"/rpc%2Fitems", "/rpc/items%2F7"} {
		rec := serveTest(t, srv, "GET", target, "", "")
		if rec.Code != http.StatusNotFound {
			t.Errorf("GET %s: status %d, want 404", target, rec.Code)
			continue
		}
		if !strings.Contains(rec.Body.String(), `"error":"WebrpcBadRoute"`) {
			t.Errorf("GET %s: body %s, want the webrpc 404 envelope", target, rec.Body)
		}
	}
}

// TestRESTEncodedSegmentsMatch pins segment-wise percent-decoding: encoded
// literal segments still match their route, and encoded path params bind
// their decoded value.
func TestRESTEncodedSegmentsMatch(t *testing.T) {
	impl := NewTestRESTServer()
	if _, err := impl.CreateItem(context.Background(), CreateItemRequest{Item: &Item{Id: 7, Name: "alice"}}); err != nil {
		t.Fatal(err)
	}
	srv := NewTestApiRestServer(impl)

	rec := serveTest(t, srv, "GET", "/rpc/%69tems/7", "", "")
	if rec.Code != http.StatusOK {
		t.Errorf("GET /rpc/%%69tems/7: status %d, want 200 (encoded literal segment)", rec.Code)
	}

	rec = serveTest(t, srv, "GET", "/rpc/items/%37", "", "")
	if rec.Code != http.StatusOK {
		t.Fatalf("GET /rpc/items/%%37: status %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), `"id":7`) {
		t.Errorf("GET /rpc/items/%%37: body %s, want item 7", rec.Body)
	}
}

// TestRESTAllowUnionsOverlappingRoutes pins the Allow union: a path served by
// both a static route and a param route reports every method that would
// dispatch, not just the first matching entry.
func TestRESTAllowUnionsOverlappingRoutes(t *testing.T) {
	srv := NewTestApiRestServer(NewTestRESTServer())
	for _, tc := range []struct{ path, want string }{
		// GET+QUERY static routes overlap the DELETE/PUT/PATCH param route.
		{"/rpc/items/find", "DELETE, GET, HEAD, PATCH, PUT, QUERY"},
		{"/rpc/items/reset", "DELETE, GET, HEAD, PATCH, POST, PUT"},
		// No overlap: the param route needs one trailing segment.
		{"/rpc/items/echo/params", "GET, HEAD"},
	} {
		rec := serveTest(t, srv, "OPTIONS", tc.path, "", "")
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("OPTIONS %s: status %d, want 405", tc.path, rec.Code)
			continue
		}
		if allow := rec.Header().Get("Allow"); allow != tc.want {
			t.Errorf("OPTIONS %s: Allow = %q, want %q", tc.path, allow, tc.want)
		}
	}
}

// TestRESTSegHelpersAdjacentParams pins adjacent-param matching: an empty
// middle literal cuts on the segment separator, so /a/{x}/{y} matches
// /rpc/a/v/w but not /rpc/a/v.
func TestRESTSegHelpersAdjacentParams(t *testing.T) {
	x, y, ok := restSeg2("/rpc/a/v/w", "/rpc/a/", "")
	if !ok || x != "v" || y != "w" {
		t.Errorf("restSeg2 adjacent params: got (%q, %q, %v), want (v, w, true)", x, y, ok)
	}
	if _, _, ok := restSeg2("/rpc/a/v", "/rpc/a/", ""); ok {
		t.Error("restSeg2 adjacent params: short path must not match")
	}
	x, ok = restSeg1("/rpc/items/%37", "/rpc/items/")
	if !ok || x != "7" {
		t.Errorf("restSeg1 unescape: got (%q, %v), want (7, true)", x, ok)
	}
}
