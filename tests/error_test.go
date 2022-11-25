package tests

import (
	"context"
	"github.com/webrpc/webrpc/tests/server"
	"testing"
	"net/http/httptest"
	"github.com/webrpc/webrpc/tests/client"
	"net/http"
	"github.com/stretchr/testify/assert"
)

type notFoundServer struct{}

func (*notFoundServer) GetError(ctx context.Context) (*server.Empty, error) {
	return nil, server.ErrorNotFound("resource was not found")
}

type invalidArgumentServer struct{}

func (*invalidArgumentServer) GetError(ctx context.Context) (*server.Empty, error) {
	return nil, server.ErrorInvalidArgument("name", "can't be empty")
}

func TestNotFound(t *testing.T) {
	srv := httptest.NewServer(server.NewErrorApiServer(&notFoundServer{}))
	defer srv.Close()

	httpClient := client.NewErrorApiClient(srv.URL, &http.Client{})

	_, err := httpClient.GetError(context.Background())

	if err, ok := err.(client.Error); ok {
		assert.Equal(t, client.ErrNotFound, err.Code())
		assert.Equal(t, "resource was not found", err.Msg())
		err.Msg()
	} else {
		assert.Fail(t, "error must be type of client.Error, actual %v", err)
	}
}

func TestInvalidArgument(t *testing.T) {
	srv := httptest.NewServer(server.NewErrorApiServer(&invalidArgumentServer{}))
	defer srv.Close()

	httpClient := client.NewErrorApiClient(srv.URL, &http.Client{})

	_, err := httpClient.GetError(context.Background())

	if err, ok := err.(client.Error); ok {
		assert.Equal(t, client.ErrInvalidArgument, err.Code())
		assert.Equal(t, "name can't be empty", err.Msg())
	} else {
		assert.Fail(t, "error must be type of client.Error, actual %v", err)
	}
}
