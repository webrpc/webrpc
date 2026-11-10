//go:build go1.27 || goexperiment.jsonv2

package v2

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// createItemStub implements only CreateItem; the other methods panic if called.
type createItemStub struct {
	TestApiRestServer
}

func (createItemStub) CreateItem(_ context.Context, req CreateItemRequest) (*CreateItemResponse, error) {
	return &CreateItemResponse{Item: req.Item}, nil
}

func TestRestHandlerRequestBody(t *testing.T) {
	svc := NewTestApiRestServer(createItemStub{})

	tests := []struct {
		name     string
		body     string
		wantCode int
	}{
		{name: "empty body", body: "", wantCode: http.StatusCreated},
		{name: "whitespace-only body", body: " \n\t ", wantCode: http.StatusCreated},
		{name: "valid body", body: `{"item":{"id":1,"name":"x"}}`, wantCode: http.StatusCreated},
		{name: "trailing data", body: `{} x`, wantCode: http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/rpc/items", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			svc.ServeHTTP(rec, req)

			assert.Equal(t, tt.wantCode, rec.Code, rec.Body.String())
		})
	}
}
