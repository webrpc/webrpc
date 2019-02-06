package main

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/webrpc/webrpc/examples/golang-basics/proto"
)

var (
	client proto.ExampleService
)

// func TestMain()

func init() {
	go func() {
		startServer()
	}()

	client = proto.NewExampleServiceClient("http://0.0.0.0:4242", &http.Client{
		Timeout: time.Duration(2 * time.Second),
	})

}

func TestPing(t *testing.T) {
	resp, err := client.Ping(context.Background())
	assert.Equal(t, false, *resp)
	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	{
		resp, err := client.GetUser(context.Background(), &proto.GetUserRequest{
			UserID: 12,
		})
		assert.Equal(t, &proto.User{ID: 1, Username: "hihi"}, resp)
		assert.NoError(t, err)
	}

	{
		// Error case, expecting to receive an error
		resp, err := client.GetUser(context.Background(), &proto.GetUserRequest{
			UserID: 911,
		})

		// TODO: err should be webrpc.Error type
		// so we can decode the .Code(), .Msg(), etc..

		assert.Nil(t, resp)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not_found")
	}
}
