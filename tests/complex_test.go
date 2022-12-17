package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webrpc/webrpc/tests/client"
	"github.com/webrpc/webrpc/tests/server"
)

func TestGetComplexStructure(t *testing.T) {
	srv := httptest.NewServer(server.NewComplexApiServer(&server.ComplexServer{}))
	defer srv.Close()

	httpClient := client.NewComplexApiClient(srv.URL, &http.Client{})

	// test getting data from client
	resp, err := httpClient.GetComplex(context.Background())
	assert.NoError(t, err)

	// test sending the same data to server
	err = httpClient.SendComplex(context.Background(), resp)
	assert.NoError(t, err)
}
