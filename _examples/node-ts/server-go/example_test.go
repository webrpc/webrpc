package main

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	client ExampleClient
)

func init() {
	go func() {
		startServer()
	}()

	client = NewExampleClient("http://0.0.0.0:3000", &http.Client{
		Timeout: time.Duration(2 * time.Second),
	})
	time.Sleep(time.Millisecond * 500)

}

func TestPing(t *testing.T) {
	err := client.Ping(context.Background(), nil)
	require.NoError(t, err)
}

// func TestGetUser(t *testing.T) {
// 	{
// 		arg1 := map[string]string{"a": "1"}
// 		resp, err := client.GetUserV2(context.Background(), GetUserRequest{UserID: 12, Prefs: arg1, ByBN: NewBigInt(123)})
// 		intent := Intent_openSession
// 		kind := Kind_ADMIN

// 		assert.Equal(t, uint32(200), resp.Code)
// 		assert.Equal(t, &User{ID: 12, Username: "hihi", Intent: intent, Kind: kind}, resp.User)
// 		assert.Equal(t, NewBigInt(31337), resp.LargeNum)
// 		assert.Equal(t, int64(31337), resp.LargeNum.AsInt().Int64())
// 		assert.NoError(t, err)
// 	}

// 	{
// 		// Error case, expecting to receive an error
// 		resp, err := client.GetUserV2(context.Background(), GetUserRequest{UserID: 911})

// 		assert.ErrorAs(t, err, &ErrUserNotFound)
// 		assert.Nil(t, resp)
// 		// assert.Equal(t, uint32(0), resp.Code)
// 		assert.Error(t, err)
// 		assert.Contains(t, err.Error(), "not found")
// 	}

// 	{
// 		name, user, err := client.FindUser(context.Background(), &SearchFilter{Q: "joe"})
// 		assert.Equal(t, "joe", name)
// 		assert.Equal(t, &User{ID: 123, Username: "joe"}, user)
// 		assert.NoError(t, err)
// 	}
// }
