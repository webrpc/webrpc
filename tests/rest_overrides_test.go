package tests

//go:generate webrpc-gen -schema=./schema/rest_overrides.ridl -target=golang -pkg=client -client -out=./overrides/client/client.gen.go
//go:generate webrpc-gen -schema=./schema/rest_overrides.ridl -target=golang -pkg=server -server -out=./overrides/server/server.gen.go

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/webrpc/webrpc/tests/overrides/client"
	"github.com/webrpc/webrpc/tests/overrides/server"
	"github.com/webrpc/webrpc/tests/overrides/types"
)

// recordingClient keeps the escaped URL of every request it sends.
type recordingClient struct {
	urls []string
}

func (c *recordingClient) Do(req *http.Request) (*http.Response, error) {
	c.urls = append(c.urls, req.URL.EscapedPath()+"?"+req.URL.RawQuery)
	return http.DefaultClient.Do(req)
}

func TestRESTOverridesRoundTrip(t *testing.T) {
	srv := httptest.NewServer(server.NewRestOverridesServer(server.Impl{}))
	defer srv.Close()

	rec := &recordingClient{}
	api := client.NewRestOverridesClient(srv.URL, rec)
	ctx := context.Background()

	for _, key := range []string{"a//b", "https://x", "abc"} {
		got, err := api.Search(ctx, client.SearchRequest{Key: types.Label(key), Tag: "ABC", Limit: 3})
		if err != nil {
			t.Fatalf("Search(%q): %v", key, err)
		}
		if want := fmt.Sprintf("key=%s tag=abc limit=3", key); got.Echo != want {
			t.Errorf("Search(%q): echo %q, want %q", key, got.Echo, want)
		}
	}

	if err := api.Index(ctx); err != nil {
		t.Errorf("Index(): %v", err)
	}
	if got, err := api.Scale(ctx, 1.5); err != nil || got != 3 {
		t.Errorf("Scale(1.5) = %v, %v; want 3", got, err)
	}

	// Alias route with a trailing literal after the param: dispatch and binding.
	resp, err := http.Get(srv.URL + "/scale/2.5/half")
	if err != nil {
		t.Fatalf("GET /scale/2.5/half: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || string(body) != `{"value":5}` {
		t.Errorf("GET /scale/2.5/half = %d %s; want 200 {\"value\":5}", resp.StatusCode, body)
	}

	for _, u := range rec.urls {
		if strings.Contains(u, "%22") {
			t.Errorf("wire URL %q has a JSON-quoted value", u)
		}
	}
	if want := "/by/abc?n=3&t=ABC"; !strings.Contains(strings.Join(rec.urls, " "), want) {
		t.Errorf("wire URLs %v do not contain %q", rec.urls, want)
	}
}

func TestRESTOverridesRawRequests(t *testing.T) {
	srv := httptest.NewServer(server.NewRestOverridesServer(server.Impl{}))
	defer srv.Close()

	noRedirect := &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error {
		return http.ErrUseLastResponse
	}}
	tests := []struct {
		name   string
		path   string
		status int
		body   string
	}{
		{"raw string value", "/by/abc?n=1", http.StatusOK, "key=abc"},
		{"text override decodes raw value", "/by/x?n=1&t=XyZ", http.StatusOK, "tag=xyz"},
		{"escaped slashes stay in one segment", "/by/a%2F..%2Fb?n=1", http.StatusOK, "key=a/../b"},
		{"escaped https value", "/by/https:%2F%2Fx?n=1", http.StatusOK, "key=https://x"},
		{"raw dot segments are 404, not a redirect", "/by/../x?n=1", http.StatusNotFound, ""},
		{"float32 overflow is 400", "/scale/1e40", http.StatusBadRequest, "is not a valid float32"},
		{"NaN is 400", "/scale/NaN", http.StatusBadRequest, "is not a valid float32"},
		{"root route serves /", "/", http.StatusOK, ""},
		{"root route is exact", "/anything/else", http.StatusNotFound, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := noRedirect.Get(srv.URL + tt.path)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			if resp.StatusCode != tt.status {
				t.Errorf("status %d, want %d (body %s)", resp.StatusCode, tt.status, body)
			}
			if loc := resp.Header.Get("Location"); loc != "" {
				t.Errorf("unexpected redirect to %q", loc)
			}
			if !strings.Contains(string(body), tt.body) {
				t.Errorf("body %s does not contain %q", body, tt.body)
			}
		})
	}
}
