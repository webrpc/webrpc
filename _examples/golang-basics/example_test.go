package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
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

// streamingExampleService overrides UploadAvatar to consume the avatar as a
// stream (hash + count) without buffering it, recording what it saw.
type streamingExampleService struct {
	*ExampleServiceRPC

	mu        sync.Mutex
	sum       []byte
	size      int64
	tempFiles []string // multipart spool files present while the handler ran
}

func (s *streamingExampleService) UploadAvatar(ctx context.Context, req UploadAvatarRequest) (*UploadAvatarResponse, error) {
	if req.Avatar == nil {
		return nil, ErrWebrpcBadRequest.WithCausef("missing avatar file")
	}
	defer req.Avatar.Body.Close()

	// The old implementation spooled the parts to multipart-* files in
	// os.TempDir before calling the handler; snapshot what is there now.
	tempFiles, _ := filepath.Glob(filepath.Join(os.TempDir(), "multipart-*"))

	h := sha256.New()
	n, err := io.Copy(h, req.Avatar.Body)
	if err != nil {
		return nil, ErrWebrpcBadRequest.WithCausef("failed to read avatar: %w", err)
	}

	s.mu.Lock()
	s.sum = h.Sum(nil)
	s.size = n
	s.tempFiles = tempFiles
	s.mu.Unlock()

	return &UploadAvatarResponse{Size: uint64(n)}, nil
}

// TestLargeStreamingUpload uploads 64 MiB (above the 32 MiB default cap, so
// the cap is raised via Options.MaxUploadSize) and asserts the handler
// receives every byte streamed off the wire: the server no longer spools
// multipart parts to temp files, which is verified by watching os.TempDir
// for the multipart-* files the old spooling implementation would create.
func TestLargeStreamingUpload(t *testing.T) {
	const uploadSize = 64 << 20

	svc := &streamingExampleService{ExampleServiceRPC: &ExampleServiceRPC{}}
	srv := httptest.NewServer(NewExampleServer(svc, &Options{MaxUploadSize: uploadSize + (1 << 20)}))
	defer srv.Close()

	tempFilesBefore := multipartTempFiles(t)

	payload := make([]byte, uploadSize)
	for i := range payload {
		payload[i] = byte(i * 31)
	}
	expectedSum := sha256.Sum256(payload)

	c := NewExampleClient(srv.URL, &http.Client{Timeout: 60 * time.Second})
	resp, err := c.UploadAvatar(context.Background(), UploadAvatarRequest{
		UserId: 1,
		Avatar: &File{
			Name:        "big.bin",
			ContentType: "application/octet-stream",
			Size:        uploadSize,
			Body:        io.NopCloser(bytes.NewReader(payload)),
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, uint64(uploadSize), resp.Size)

	svc.mu.Lock()
	assert.Equal(t, expectedSum[:], svc.sum)
	assert.Equal(t, int64(uploadSize), svc.size)
	assert.Equal(t, tempFilesBefore, svc.tempFiles, "upload must not spool multipart temp files")
	svc.mu.Unlock()
}

// multipartTempFiles lists mime/multipart's spool files in os.TempDir.
func multipartTempFiles(t *testing.T) []string {
	t.Helper()
	files, err := filepath.Glob(filepath.Join(os.TempDir(), "multipart-*"))
	assert.NoError(t, err)
	return files
}
