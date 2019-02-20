package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webrpc/webrpc/lib/webrpc-go/tests/proto"
)

// end-to-end integration testing of from json schema file, to schema parsing+validatoin,
// code-generation, to server start and client communication -- the whole shabang is tested here.
//
// just run.. `make build && make test`

func TestPing(t *testing.T) {
	err := Client.Ping(context.Background())
	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	{
		code, user, err := Client.GetUser(context.Background(), nil, 12)
		assert.Equal(t, &proto.User{ID: 12, Username: "hihi"}, user)
		assert.Equal(t, uint32(1), code)
		assert.NoError(t, err)
	}

	{
		// Error case, expecting to receive an error
		code, user, err := Client.GetUser(context.Background(), nil, 911)

		// TODO: err should be webrpc.Error type
		// so we can decode the .Code(), .Msg(), etc..

		assert.Equal(t, uint32(0), code)
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not_found")
	}
}
