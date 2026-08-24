package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	client ExampleClient
)

// func TestMain()

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

func TestDeprecatedUserEndpoint(t *testing.T) {
	arg1 := map[string]string{"a": "1"}

	_, err := client.GetUser(context.Background(), GetUserRequest{UserID: 42, Prefs: arg1})

	assert.Error(t, err)
}

func TestGetUser(t *testing.T) {
	{
		arg1 := map[string]string{"a": "1"}
		resp, err := client.GetUserV2(context.Background(), GetUserRequest{UserID: 12, Prefs: arg1})
		intent := Intent_openSession
		kind := Kind_ADMIN

		assert.Equal(t, uint32(200), resp.Code)
		assert.Equal(t, &User{ID: 12, Username: "hihi", Intent: intent, Kind: kind}, resp.User)
		assert.NoError(t, err)
	}

	{
		// Error case, expecting to receive an error
		resp, err := client.GetUserV2(context.Background(), GetUserRequest{UserID: 911})

		assert.ErrorAs(t, err, &ErrUserNotFound)
		assert.Nil(t, resp)
		// assert.Equal(t, uint32(0), resp.Code)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	}

	{
		name, user, err := client.FindUser(context.Background(), &SearchFilter{Q: "joe"})
		assert.Equal(t, "joe", name)
		assert.Equal(t, &User{ID: 123, Username: "joe"}, user)
		assert.NoError(t, err)
	}
}

func TestUploadDownloadAvatar(t *testing.T) {
	ctx := context.Background()
	content := []byte("not really a png \x00\x01\xff")

	resp, err := client.UploadAvatar(ctx, UploadAvatarRequest{
		UserId: 77,
		Avatar: &File{
			Name:        "me.png",
			ContentType: "image/png",
			Size:        int64(len(content)),
			Body:        io.NopCloser(bytes.NewReader(content)),
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, uint64(len(content)), resp.Size)

	avatar, err := client.DownloadAvatar(ctx, DownloadAvatarRequest{UserId: 77})
	assert.NoError(t, err)
	defer avatar.Body.Close()

	data, err := io.ReadAll(avatar.Body)
	assert.NoError(t, err)
	assert.Equal(t, content, data)
	assert.Equal(t, "me.png", avatar.Name)
	assert.Equal(t, "image/png", avatar.ContentType)
	assert.Equal(t, int64(len(content)), avatar.Size)

	// Errors on download methods still arrive as the webrpc JSON error envelope.
	_, err = client.DownloadAvatar(ctx, DownloadAvatarRequest{UserId: 999999})
	assert.ErrorAs(t, err, &ErrUserNotFound)
}
