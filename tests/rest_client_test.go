package tests

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/webrpc/webrpc/tests/client"
)

// capturingClient records each request and its body, then answers 200 with "{}".
type capturingClient struct {
	reqs   []*http.Request
	bodies []string
}

func (c *capturingClient) Do(req *http.Request) (*http.Response, error) {
	c.reqs = append(c.reqs, req)
	body := ""
	if req.Body != nil {
		b, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		body = string(b)
	}
	c.bodies = append(c.bodies, body)
	return &http.Response{StatusCode: 200, Header: http.Header{}, Body: io.NopCloser(strings.NewReader("{}")), Request: req}, nil
}

func TestRESTClientQueryParams(t *testing.T) {
	hc := &capturingClient{}
	api := client.NewTestApiRestClient("http://example.test", hc)
	if _, err := api.ListItems(context.Background(), client.ListItemsRequest{Tags: []string{"a b", "c"}, Limit: 5}); err != nil {
		t.Fatal(err)
	}
	req := hc.reqs[0]
	if req.Method != "GET" || req.URL.Path != "/rpc/items" {
		t.Fatalf("got %s %s, want GET /rpc/items", req.Method, req.URL.Path)
	}
	if got := req.URL.Query()["tags"]; len(got) != 2 || got[0] != "a b" || got[1] != "c" {
		t.Errorf("tags query = %q, want repeated keys [a b c]", got)
	}
	if got := req.URL.Query().Get("limit"); got != "5" {
		t.Errorf("limit query = %q, want 5", got)
	}
	if _, ok := req.URL.Query()["q"]; ok {
		t.Errorf("absent optional q was sent: %s", req.URL.RawQuery)
	}
	if req.Body != nil || req.GetBody != nil {
		t.Errorf("GET sent a body")
	}
}

func TestRESTClientPathAndOptionalQuery(t *testing.T) {
	hc := &capturingClient{}
	api := client.NewTestApiRestClient("http://example.test", hc)
	details := true
	if _, err := api.GetItem(context.Background(), client.GetItemRequest{Id: 7, Details: &details}); err != nil {
		t.Fatal(err)
	}
	if _, err := api.GetItem(context.Background(), client.GetItemRequest{Id: 8}); err != nil {
		t.Fatal(err)
	}
	if u := hc.reqs[0].URL; u.Path != "/rpc/items/7" || u.Query().Get("details") != "true" {
		t.Errorf("got %s, want /rpc/items/7?details=true", u.String())
	}
	if u := hc.reqs[1].URL; u.Path != "/rpc/items/8" || u.RawQuery != "" {
		t.Errorf("got %s, want /rpc/items/8 with no query", u.String())
	}
}

func TestRESTClientPathParamsKeepBaseAndTypedText(t *testing.T) {
	hc := &capturingClient{}
	api := client.NewTestApiRestClient("http://example.test/api/", hc)
	ts := time.Date(2026, 9, 3, 12, 30, 15, 0, time.UTC)
	if _, err := api.EchoParams(context.Background(), "hello world/&?=", -42, true, 1.5, ts, client.Status_NOT_AVAILABLE, nil, nil); err != nil {
		t.Fatal(err)
	}
	u := hc.reqs[0].URL
	if u.Path != "/api/rpc/items/echo/params" {
		t.Errorf("path = %q, want /api/rpc/items/echo/params", u.Path)
	}
	if got := u.Query().Get("str"); got != "hello world/&?=" {
		t.Errorf("str query = %q", got)
	}
	if got := u.Query().Get("ts"); got != "2026-09-03T12:30:15Z" {
		t.Errorf("ts query = %q, want the RFC 3339 text form", got)
	}
	wantStatus, _ := client.Status_NOT_AVAILABLE.MarshalText()
	if got := u.Query().Get("status"); got != string(wantStatus) {
		t.Errorf("status query = %q, want %q", got, wantStatus)
	}
}

func TestRESTClientBodyRouteSetsGetBody(t *testing.T) {
	hc := &capturingClient{}
	api := client.NewTestApiRestClient("http://example.test", hc)
	if _, err := api.CreateItem(context.Background(), client.CreateItemRequest{Item: &client.Item{Id: 5, Name: "alice"}}); err != nil {
		t.Fatal(err)
	}
	req := hc.reqs[0]
	if req.Method != "POST" || req.URL.Path != "/rpc/items" || req.URL.RawQuery != "" {
		t.Fatalf("got %s %s?%s, want POST /rpc/items", req.Method, req.URL.Path, req.URL.RawQuery)
	}
	if req.GetBody == nil {
		t.Fatal("GetBody is nil: a redirect or retry could not resend the body")
	}
	again, err := req.GetBody()
	if err != nil {
		t.Fatal(err)
	}
	b, _ := io.ReadAll(again)
	if !strings.Contains(string(b), `"name":"alice"`) || req.Header.Get("Content-Type") != "application/json" {
		t.Errorf("GetBody = %s, Content-Type = %q", b, req.Header.Get("Content-Type"))
	}
}

func TestRESTClientQueryVerbSendsBody(t *testing.T) {
	hc := &capturingClient{}
	api := client.NewTestApiRestClient("http://example.test", hc)
	if _, err := api.QueryItems(context.Background(), "carol"); err != nil {
		t.Fatal(err)
	}
	req := hc.reqs[0]
	if req.Method != "QUERY" || req.URL.Path != "/rpc/items/search" {
		t.Fatalf("got %s %s, want QUERY /rpc/items/search", req.Method, req.URL.Path)
	}
	if !strings.Contains(hc.bodies[0], `"q":"carol"`) {
		t.Errorf("QUERY body = %q, want the q field", hc.bodies[0])
	}
	if req.GetBody == nil || req.Header.Get("Content-Type") != "application/json" {
		t.Errorf("QUERY must set GetBody and a JSON Content-Type")
	}
}

func TestRESTClientMultiplePathParams(t *testing.T) {
	hc := &capturingClient{}
	api := client.NewTestApiRestClient("http://example.test", hc)
	if _, err := api.GetChild(context.Background(), 1, 2); err != nil {
		t.Fatal(err)
	}
	if got := hc.reqs[0].URL.Path; got != "/rpc/items/1/children/2" {
		t.Errorf("path = %q, want /rpc/items/1/children/2", got)
	}
}

func TestRESTClientCallsCanonicalRouteOnly(t *testing.T) {
	hc := &capturingClient{}
	api := client.NewTestApiRestClient("http://example.test", hc)
	if _, err := api.LookupItem(context.Background(), 5); err != nil {
		t.Fatal(err)
	}
	if _, err := api.FindItems(context.Background(), "x"); err != nil {
		t.Fatal(err)
	}
	if req := hc.reqs[0]; req.Method != "GET" || req.URL.Path != "/rpc/items/lookup/5" {
		t.Errorf("LookupItem: got %s %s, want GET /rpc/items/lookup/5", req.Method, req.URL.Path)
	}
	if req := hc.reqs[1]; req.Method != "GET" || req.URL.Path != "/rpc/items/find" || req.URL.Query().Get("q") != "x" {
		t.Errorf("FindItems: got %s %s, want GET /rpc/items/find?q=x", req.Method, req.URL.String())
	}
}
