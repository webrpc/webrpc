package tests

//go:generate webrpc-gen -schema=./schema/test.ridl -target=json -out=./schema/test.gen.json
//go:generate webrpc-gen -schema=./schema/test.ridl -target=golang -pkg=client -client -out=./client/client.gen.go
//go:generate webrpc-gen -schema=./schema/test.ridl -target=golang -pkg=server -server -out=./server/server.gen.go

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webrpc/webrpc/tests/client"
	"github.com/webrpc/webrpc/tests/server"
)

func TestInteroperability(t *testing.T) {
	srv := httptest.NewServer(server.NewTestApiServer(&server.TestServer{}))
	defer srv.Close()

	err := client.RunTests(context.Background(), srv.URL)
	assert.NoError(t, err)
}
