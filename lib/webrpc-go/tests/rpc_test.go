package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webrpc/webrpc/lib/webrpc-go/tests/proto"
)

func TestPing(t *testing.T) {
	resp, err := Client.Ping(context.Background())
	assert.Equal(t, false, *resp)
	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	{
		resp, err := Client.GetUser(context.Background(), &proto.GetUserRequest{
			UserID: 12,
		})
		assert.Equal(t, &proto.User{ID: 1, Username: "hihi"}, resp)
		assert.NoError(t, err)
	}

	{
		// Error case, expecting to receive an error
		resp, err := Client.GetUser(context.Background(), &proto.GetUserRequest{
			UserID: 911,
		})

		// TODO: err should be webrpc.Error type
		// so we can decode the .Code(), .Msg(), etc..

		assert.Nil(t, resp)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not_found")
	}
}
