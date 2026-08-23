package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFileUploadDownload(t *testing.T) {
	server := httptest.NewServer(NewFileServiceServer(&FileServiceRPC{
		avatars: map[uint64]storedFile{},
	}))
	defer server.Close()

	client := NewFileServiceClient(server.URL, &http.Client{})
	ctx := context.Background()

	// Upload a single file alongside a scalar argument.
	content := "not really a png"
	size, err := client.UploadAvatar(ctx, 1234, &File{
		Name:        "me.png",
		ContentType: "image/png",
		Size:        int64(len(content)),
		Body:        io.NopCloser(strings.NewReader(content)),
	})
	if err != nil {
		t.Fatal(err)
	}
	if size != uint64(len(content)) {
		t.Fatalf("expected size %d, got %d", len(content), size)
	}

	// Upload multiple files at once.
	count, err := client.UploadAttachments(ctx, 1234, []*File{
		{Name: "a.txt", ContentType: "text/plain", Size: 1, Body: io.NopCloser(strings.NewReader("a"))},
		{Name: "b.txt", ContentType: "text/plain", Size: 1, Body: io.NopCloser(strings.NewReader("b"))},
	})
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Fatalf("expected 2 attachments, got %d", count)
	}

	// Download the avatar back.
	avatar, err := client.DownloadAvatar(ctx, 1234)
	if err != nil {
		t.Fatal(err)
	}
	defer avatar.Body.Close()

	data, err := io.ReadAll(avatar.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != content {
		t.Fatalf("expected %q, got %q", content, data)
	}
	if avatar.Name != "me.png" || avatar.ContentType != "image/png" || avatar.Size != int64(len(content)) {
		t.Fatalf("unexpected file metadata: %+v", avatar)
	}

	// Errors still arrive as the standard webrpc JSON error envelope.
	_, err = client.DownloadAvatar(ctx, 9999)
	rpcErr, ok := err.(WebRPCError)
	if !ok {
		t.Fatalf("expected WebRPCError, got %T: %v", err, err)
	}
	if rpcErr.Code != ErrWebrpcBadRoute.Code {
		t.Fatalf("unexpected error code: %v", rpcErr)
	}
}
