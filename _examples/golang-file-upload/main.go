//go:generate webrpc-gen -schema=fileupload.ridl -target=golang -pkg=main -server -client -out=./fileupload.gen.go
package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"sync"
)

func main() {
	// `go run . client http://localhost:4243` runs the client test suite
	// against a running FileService server instead (see interop/test.ts).
	if len(os.Args) > 2 && os.Args[1] == "client" {
		if err := runClientSuite(os.Args[2]); err != nil {
			log.Fatal(err)
		}
		log.Println("go client suite: all tests passed")
		return
	}

	addr := ":4242"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	err := startServer(addr)
	if err != nil {
		log.Fatal(err)
	}
}

func startServer(addr string) error {
	webrpcHandler := NewFileServiceServer(&FileServiceRPC{
		avatars: map[uint64]storedFile{},
	}, &Options{
		MaxUploadSize: 8 << 20, // 8 MiB; defaults to DefaultMaxUploadSize (32 MiB) when unset
	})

	log.Printf("serving on %s", addr)
	return http.ListenAndServe(addr, webrpcHandler)
}

type storedFile struct {
	name        string
	contentType string
	data        []byte
}

// FileServiceRPC implements the FileServiceServer interface generated from
// the fileupload.ridl schema.
type FileServiceRPC struct {
	mu      sync.Mutex
	avatars map[uint64]storedFile
}

func (s *FileServiceRPC) UploadAvatar(ctx context.Context, userId uint64, avatar *File) (uint64, error) {
	if avatar == nil {
		return 0, ErrWebrpcBadRequest.WithCausef("missing avatar file")
	}
	defer avatar.Body.Close()

	data, err := io.ReadAll(avatar.Body)
	if err != nil {
		return 0, ErrWebrpcBadRequest.WithCausef("failed to read avatar: %w", err)
	}

	s.mu.Lock()
	s.avatars[userId] = storedFile{name: avatar.Name, contentType: avatar.ContentType, data: data}
	s.mu.Unlock()

	log.Printf("stored avatar %q (%s, %d bytes) for user %d", avatar.Name, avatar.ContentType, len(data), userId)
	return uint64(len(data)), nil
}

func (s *FileServiceRPC) UploadAttachments(ctx context.Context, userId uint64, attachments []*File) (uint32, error) {
	var count uint32
	for _, attachment := range attachments {
		size, err := io.Copy(io.Discard, attachment.Body)
		attachment.Body.Close()
		if err != nil {
			return count, ErrWebrpcBadRequest.WithCausef("failed to read attachment %q: %w", attachment.Name, err)
		}
		log.Printf("received attachment %q (%s, %d bytes) for user %d", attachment.Name, attachment.ContentType, size, userId)
		count++
	}
	return count, nil
}

func (s *FileServiceRPC) DownloadAvatar(ctx context.Context, userId uint64) (*File, error) {
	s.mu.Lock()
	avatar, ok := s.avatars[userId]
	s.mu.Unlock()

	if !ok {
		return nil, ErrWebrpcBadRoute.WithCausef("no avatar for user %d", userId)
	}

	return &File{
		Name:        avatar.name,
		ContentType: avatar.contentType,
		Size:        int64(len(avatar.data)),
		Body:        io.NopCloser(bytes.NewReader(avatar.data)),
	}, nil
}

func (s *FileServiceRPC) StampAvatar(ctx context.Context, userId uint64, stamp *File) (*File, error) {
	if stamp == nil {
		return nil, ErrWebrpcBadRequest.WithCausef("missing stamp file")
	}
	stampData, err := io.ReadAll(stamp.Body)
	stamp.Body.Close()
	if err != nil {
		return nil, ErrWebrpcBadRequest.WithCausef("failed to read stamp: %w", err)
	}

	s.mu.Lock()
	avatar, ok := s.avatars[userId]
	s.mu.Unlock()

	if !ok {
		return nil, ErrWebrpcBadRoute.WithCausef("no avatar for user %d", userId)
	}

	stamped := append(append([]byte{}, avatar.data...), stampData...)
	return &File{
		Name:        "stamped-" + avatar.name,
		ContentType: avatar.contentType,
		Size:        int64(len(stamped)),
		Body:        io.NopCloser(bytes.NewReader(stamped)),
	}, nil
}

func (s *FileServiceRPC) ListAvatars(ctx context.Context) ([]uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userIds := make([]uint64, 0, len(s.avatars))
	for userId := range s.avatars {
		userIds = append(userIds, userId)
	}
	sort.Slice(userIds, func(i, j int) bool { return userIds[i] < userIds[j] })
	return userIds, nil
}
