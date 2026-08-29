package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/webrpc/webrpc/_example/golang-basics/admin"
)

func main() {
	err := startServer()
	if err != nil {
		log.Fatal(err)
	}
}

func startServer() error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	webrpcHandler := NewExampleServer(&ExampleServiceRPC{})
	webrpcHandler.OnError = func(r *http.Request, err *WebRPCError) {
		m, ok := MethodCtx(r.Context())

		if ok {
			_, ok = m.Annotations()["deprecated"]
			if ok {
				fmt.Println(r.URL.Path, "deprecated")
			}
		}
	}
	webrpcHandler.OnRequest = func(w http.ResponseWriter, r *http.Request) error {
		m, ok := MethodCtx(r.Context())
		if !ok {
			return fmt.Errorf("could not find method context for request method: %s", r.URL.Path)
		}

		newEndpoint, ok := m.Annotations()["deprecated"]
		if ok {
			return fmt.Errorf(
				"endpoint %s has been deprecated in favor of endpoint %s",
				r.URL.Path,
				newEndpoint,
			)
		}

		return nil
	}

	r.Handle("/admin/*", admin.NewAdminServer(&AdminServiceRPC{}))
	r.Handle("/*", webrpcHandler)

	addr := ":4242"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	return http.ListenAndServe(addr, r)
}

type AdminServiceRPC struct{}

func (*AdminServiceRPC) Auth(ctx context.Context) (string, string, error) {
	return "jwt", "admin", nil
}

func (s *AdminServiceRPC) Status(ctx context.Context) (bool, error) {
	return true, nil
}

func (s *AdminServiceRPC) Version(ctx context.Context) (*admin.Version, error) {
	return &admin.Version{
		WebrpcVersion: WebRPCVersion(),
		SchemaVersion: WebRPCSchemaVersion(),
		SchemaHash:    WebRPCSchemaHash(),
	}, nil
}

type storedAvatar struct {
	name        string
	contentType string
	data        []byte
}

type ExampleServiceRPC struct {
	mu      sync.Mutex
	avatars map[uint64]storedAvatar
}

func (s *ExampleServiceRPC) Ping(ctx context.Context) error {
	return nil
}

func (s *ExampleServiceRPC) Status(ctx context.Context) (bool, error) {
	return true, nil
}

func (s *ExampleServiceRPC) Version(ctx context.Context) (*Version, error) {
	return &Version{
		WebrpcVersion: WebRPCVersion(),
		SchemaVersion: WebRPCSchemaVersion(),
		SchemaHash:    WebRPCSchemaHash(),
	}, nil
}

func (s *ExampleServiceRPC) GetUser(ctx context.Context, req GetUserRequest) (*GetUserResponse, error) {
	if req.UserID == 911 {
		return nil, ErrUserNotFound
	}

	if req.UserID == 31337 {
		return nil, ErrUserNotFound.WithCausef("unknown user id %d", req.UserID)
	}

	kind := Kind_ADMIN
	intent := Intent_openSession

	return &GetUserResponse{
		Code: 200,
		User: &User{
			ID:       req.UserID,
			Username: "hihi",
			Kind:     kind,
			Intent:   intent,
		},
	}, nil
}

func (s *ExampleServiceRPC) GetUserV2(ctx context.Context, req GetUserRequest) (*GetUserResponse, error) {
	if req.UserID == 911 {
		return nil, ErrUserNotFound
	}
	if req.UserID == 31337 {
		return nil, ErrUserNotFound.WithCausef("unknown user id %d", req.UserID)
	}

	kind := Kind_ADMIN
	intent := Intent_openSession

	return &GetUserResponse{
		Code: 200,
		User: &User{
			ID:       req.UserID,
			Username: "hihi",
			Kind:     kind,
			Intent:   intent,
		},
		Profile: "https://www.google.com/images/john-doe.jpg",
	}, nil
}

func (s *ExampleServiceRPC) FindUser(ctx context.Context, f *SearchFilter) (string, *User, error) {
	name := f.Q
	return f.Q, &User{
		ID: 123, Username: name,
	}, nil
}

func (s *ExampleServiceRPC) GetIntents(ctx context.Context) ([]Intent, error) {
	return []Intent{Intent_openSession, Intent_closeSession, Intent_validateSession}, nil
}

func (s *ExampleServiceRPC) CountIntents(ctx context.Context, userID uint64) (map[Intent]uint32, error) {
	return map[Intent]uint32{
		Intent_openSession:     1,
		Intent_closeSession:    2,
		Intent_validateSession: 3,
	}, nil
}

// UploadAvatar receives the avatar as a *WebrpcFile parsed from the multipart
// request body, alongside the JSON-encoded scalar fields of the request.
func (s *ExampleServiceRPC) UploadAvatar(ctx context.Context, req UploadAvatarRequest) (*UploadAvatarResponse, error) {
	if req.Avatar == nil {
		return nil, ErrWebrpcBadRequest.WithCausef("missing avatar file")
	}
	defer req.Avatar.Body.Close()

	data, err := io.ReadAll(req.Avatar.Body)
	if err != nil {
		return nil, ErrWebrpcBadRequest.WithCausef("failed to read avatar: %w", err)
	}

	s.mu.Lock()
	if s.avatars == nil {
		s.avatars = map[uint64]storedAvatar{}
	}
	s.avatars[req.UserId] = storedAvatar{name: req.Avatar.Name, contentType: req.Avatar.ContentType, data: data}
	s.mu.Unlock()

	return &UploadAvatarResponse{Size: uint64(len(data))}, nil
}

// DownloadAvatar returns a *WebrpcFile that is streamed to the client as the raw
// response body, with the file's metadata in the response headers.
func (s *ExampleServiceRPC) DownloadAvatar(ctx context.Context, req DownloadAvatarRequest) (*WebrpcFile, error) {
	s.mu.Lock()
	avatar, ok := s.avatars[req.UserId]
	s.mu.Unlock()

	if !ok {
		return nil, ErrUserNotFound.WithCausef("no avatar for user %d", req.UserId)
	}

	return &WebrpcFile{
		Name:        avatar.name,
		ContentType: avatar.contentType,
		Size:        int64(len(avatar.data)),
		Body:        io.NopCloser(bytes.NewReader(avatar.data)),
	}, nil
}
