package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webrpc/webrpc/tests/client"
	"github.com/webrpc/webrpc/tests/server"
)

func TestGetComplexStructure(t *testing.T) {
	srv := httptest.NewServer(server.NewComplexApiServer(&server.ComplexServer{}))
	defer srv.Close()

	err := client.TestComplexData(srv.URL)
	assert.NoError(t, err)
}
