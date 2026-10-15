package client

import (
	"errors"
	"net/url"
	"testing"
)

func TestRestURLBuild(t *testing.T) {
	tests := []struct {
		name    string
		path    map[string]string
		query   url.Values
		route   string
		want    string
		wantErr bool
	}{
		{name: "escapes a slash and a space", path: map[string]string{"id": "a/b c"}, route: "/items/{id}", want: "http://h/items/a%2Fb%20c"},
		{name: "placeholder in a value is not substituted again", path: map[string]string{"a": "{b}", "b": "x"}, route: "/items/{a}/{b}", want: "http://h/items/%7Bb%7D/x"},
		{name: "percent sequences are escaped, not decoded", path: map[string]string{"a": "%7Bb%7D", "b": "x"}, route: "/items/{a}/{b}", want: "http://h/items/%257Bb%257D/x"},
		{name: "query is appended", path: map[string]string{"id": "7"}, query: url.Values{"tags": {"x", "y"}}, route: "/items/{id}", want: "http://h/items/7?tags=x&tags=y"},
		{name: "empty value is rejected", path: map[string]string{"id": ""}, route: "/items/{id}", wantErr: true},
		{name: "dot is rejected", path: map[string]string{"id": "."}, route: "/items/{id}", wantErr: true},
		{name: "dot dot is rejected", path: map[string]string{"id": ".."}, route: "/items/{id}", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := restURL{query: tt.query}
			for name, v := range tt.path {
				u.setPath(name, v)
			}
			got, err := u.build("http://h", tt.route)
			if tt.wantErr {
				var rpcErr WebRPCError
				if !errors.As(err, &rpcErr) || rpcErr.Name != ErrWebrpcBadRequest.Name {
					t.Fatalf("err = %v, want WebrpcBadRequest", err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("build() = %q, want %q", got, tt.want)
			}
		})
	}
}
