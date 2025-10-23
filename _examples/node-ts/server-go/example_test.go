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

func TestGetUserSuccess(t *testing.T) {
	code, user, err := client.GetUser(context.Background(), 123)
	require.NoError(t, err)
	require.Equal(t, uint32(200), code)
	require.NotNil(t, user)
	require.Equal(t, uint64(123), user.ID)
	require.Equal(t, "user-123", user.Username)
	require.Equal(t, Kind_USER, user.Role)
	// BigInt underlying string value
	require.Equal(t, "31337", user.Balance.String())
	require.NotNil(t, user.Extra)
	require.Equal(t, "additional user info", user.Extra.Info)
	require.Equal(t, "5678", user.Extra.Amount.String())
	require.Len(t, user.Extra.Points, 3)
	require.Equal(t, "100", user.Extra.Points[0].String())
}

func TestGetUserForbidden(t *testing.T) {
	_, _, err := client.GetUser(context.Background(), 911)
	require.Error(t, err)
	// Expect a WebRPCError with endpoint error cause
	var rpcErr WebRPCError
	require.ErrorAs(t, err, &rpcErr)
	require.Equal(t, ErrWebrpcEndpoint.Code, rpcErr.Code)
}

func TestGetArticle(t *testing.T) {
	req := GetArticleRequest{ArticleID: 7, ByBN: NewBigInt(42)}
	resp, err := client.GetArticle(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "Article 7", resp.Title)
	require.NotNil(t, resp.Content)
	require.Contains(t, *resp.Content, "article 7")
	// LargeNum should be ByBN * 2 (42 * 2)
	require.Equal(t, "84", resp.LargeNum.String())
}
