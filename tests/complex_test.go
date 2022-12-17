package tests

import (
	"context"
	"testing"
	"net/http"
	"github.com/webrpc/webrpc/tests/client"
	"github.com/stretchr/testify/assert"
)

func TestGetComplexStructure(t *testing.T) {
	httpClient := client.NewComplexApiClient("http://localhost:9999/complex", &http.Client{})

	// test getting data from client
	resp, err := httpClient.GetComplex(context.Background())
	assert.NoError(t, err)

	// test sending the same data to server
	err = httpClient.SendComplex(context.Background(), resp)
	assert.NoError(t, err)
}
