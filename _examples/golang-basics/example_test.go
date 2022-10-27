package main

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	client ExampleService
)

func init() {
	go func() {
		startServer()
	}()

	client = NewExampleClient("http://0.0.0.0:4242", &http.Client{
		Timeout: time.Duration(2 * time.Second),
	})
	time.Sleep(time.Millisecond * 500)

}

func TestPing(t *testing.T) {
	err := client.Ping(context.Background())
	assert.NoError(t, err)
}

func TestStatus(t *testing.T) {
	resp, err := client.Status(context.Background())
	assert.Equal(t, true, resp)
	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	{
		arg1 := map[string]string{"a": "1"}
		code, user, err := client.GetUser(context.Background(), arg1, 12)
		assert.Equal(t, uint32(200), code)
		assert.Equal(t, &User{ID: 12, Username: "hihi"}, user)
		assert.NoError(t, err)
	}

	{
		// Error case, expecting to receive an error
		code, user, err := client.GetUser(context.Background(), nil, 911)

		assert.True(t, errors.Is(err, ErrUserNotFound))
		assert.Nil(t, user)
		assert.Equal(t, uint32(0), code)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not found")

		rpcErr, ok := err.(RPCError)
		require.True(t, ok)
		assert.Equal(t, ErrUserNotFound.Code, rpcErr.Code)
		assert.Contains(t, rpcErr.Unwrap().Error(), "911")
	}

	{
		name, user, err := client.FindUser(context.Background(), &SearchFilter{Q: "joe"})
		assert.Equal(t, "joe", name)
		assert.Equal(t, &User{ID: 123, Username: "joe"}, user)
		assert.NoError(t, err)
	}
}
