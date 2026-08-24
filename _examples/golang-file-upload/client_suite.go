package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// runClientSuite runs the Go client test suite against a running FileService
// server (Go or TypeScript) at the given address. It exercises a plain JSON
// method, single and repeated file uploads, a file download, a combined
// upload + download method, and the JSON error envelope. Used by
// interop/test.ts to verify the TypeScript server against the Go client.
func runClientSuite(addr string) error {
	client := NewFileServiceClient(addr, &http.Client{})
	ctx := context.Background()
	const userId = 777

	// Plain JSON method on a service that also transfers files.
	userIds, err := client.ListAvatars(ctx)
	if err != nil {
		return fmt.Errorf("ListAvatars: %w", err)
	}
	if len(userIds) != 0 {
		return fmt.Errorf("ListAvatars: expected no avatars yet, got %v", userIds)
	}

	// Single file upload alongside a scalar argument.
	avatarData := make([]byte, 1024)
	if _, err := rand.Read(avatarData); err != nil {
		return err
	}
	size, err := client.UploadAvatar(ctx, userId, &File{
		Name:        "gopher.png",
		ContentType: "image/png",
		Size:        int64(len(avatarData)),
		Body:        io.NopCloser(strings.NewReader(string(avatarData))),
	})
	if err != nil {
		return fmt.Errorf("UploadAvatar: %w", err)
	}
	if size != uint64(len(avatarData)) {
		return fmt.Errorf("UploadAvatar: expected size %d, got %d", len(avatarData), size)
	}

	// Repeated []file upload.
	count, err := client.UploadAttachments(ctx, userId, []*File{
		{Name: "a.txt", ContentType: "text/plain", Size: 1, Body: io.NopCloser(strings.NewReader("a"))},
		{Name: "b.txt", ContentType: "text/plain", Size: 1, Body: io.NopCloser(strings.NewReader("b"))},
	})
	if err != nil {
		return fmt.Errorf("UploadAttachments: %w", err)
	}
	if count != 2 {
		return fmt.Errorf("UploadAttachments: expected 2 attachments, got %d", count)
	}

	// File download: metadata from the headers, bytes compared to the upload.
	avatar, err := client.DownloadAvatar(ctx, userId)
	if err != nil {
		return fmt.Errorf("DownloadAvatar: %w", err)
	}
	data, err := io.ReadAll(avatar.Body)
	avatar.Body.Close()
	if err != nil {
		return fmt.Errorf("DownloadAvatar: reading body: %w", err)
	}
	if string(data) != string(avatarData) {
		return fmt.Errorf("DownloadAvatar: content mismatch")
	}
	if avatar.Name != "gopher.png" || avatar.ContentType != "image/png" {
		return fmt.Errorf("DownloadAvatar: unexpected file metadata: name=%q contentType=%q", avatar.Name, avatar.ContentType)
	}
	if avatar.Size >= 0 && avatar.Size != int64(len(avatarData)) {
		return fmt.Errorf("DownloadAvatar: unexpected size %d", avatar.Size)
	}

	// Combined upload + download in one method.
	stampData := make([]byte, 256)
	if _, err := rand.Read(stampData); err != nil {
		return err
	}
	stamped, err := client.StampAvatar(ctx, userId, &File{
		Name:        "stamp.bin",
		ContentType: "application/octet-stream",
		Size:        int64(len(stampData)),
		Body:        io.NopCloser(strings.NewReader(string(stampData))),
	})
	if err != nil {
		return fmt.Errorf("StampAvatar: %w", err)
	}
	stampedData, err := io.ReadAll(stamped.Body)
	stamped.Body.Close()
	if err != nil {
		return fmt.Errorf("StampAvatar: reading body: %w", err)
	}
	if string(stampedData) != string(avatarData)+string(stampData) {
		return fmt.Errorf("StampAvatar: content mismatch")
	}
	if stamped.Name != "stamped-gopher.png" {
		return fmt.Errorf("StampAvatar: unexpected file name %q", stamped.Name)
	}

	userIds, err = client.ListAvatars(ctx)
	if err != nil {
		return fmt.Errorf("ListAvatars: %w", err)
	}
	if len(userIds) != 1 || userIds[0] != userId {
		return fmt.Errorf("ListAvatars: expected [%d], got %v", userId, userIds)
	}

	// Errors still arrive as the standard webrpc JSON error envelope, even on
	// file download methods.
	_, err = client.DownloadAvatar(ctx, 999999)
	rpcErr, ok := err.(WebRPCError)
	if !ok {
		return fmt.Errorf("DownloadAvatar: expected WebRPCError, got %T: %v", err, err)
	}
	if rpcErr.Code != ErrWebrpcBadRoute.Code {
		return fmt.Errorf("DownloadAvatar: unexpected error code: %v", rpcErr)
	}

	return nil
}
