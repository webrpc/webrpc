package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	// CORS middleware for dev/demo purposes. Adjust origins/methods in production.
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           int((12 * time.Hour).Seconds()),
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	webrpcHandler := NewExampleServer(&ExampleServiceRPC{
		avatars: map[uint64]storedAvatar{},
	})

	r.Handle("/*", webrpcHandler)

	return http.ListenAndServe(":3000", r)
}

type ExampleServiceRPC struct {
	mu      sync.Mutex
	avatars map[uint64]storedAvatar
}

type storedAvatar struct {
	name        string
	contentType string
	data        []byte
}

func (s *ExampleServiceRPC) Ping(ctx context.Context, counter *BigInt) error {
	return nil
}

func (s *ExampleServiceRPC) GetUser(ctx context.Context, userID uint64) (uint32, *User, error) {
	if userID == 911 {
		return 0, nil, ErrWebrpcEndpoint.WithCausef("user is forbidden")
	}

	return 200, &User{
		ID:       userID,
		Username: fmt.Sprintf("user-%d", userID),
		Role:     Kind_USER,
		Balance:  NewBigInt(31337),
		Extra: &Extra{
			Info:   "additional user info",
			Amount: NewBigInt(5678),
			Points: []BigInt{NewBigInt(100), NewBigInt(200), NewBigInt(300)},
		},
	}, nil
}

func (s *ExampleServiceRPC) UploadAvatar(ctx context.Context, userID uint64, avatar *WebrpcFile) (uint64, string, error) {
	if avatar == nil {
		return 0, "", ErrWebrpcBadRequest.WithCausef("missing avatar file")
	}
	defer avatar.Body.Close()

	data, err := io.ReadAll(avatar.Body)
	if err != nil {
		return 0, "", ErrWebrpcBadRequest.WithCausef("failed to read avatar: %w", err)
	}

	s.mu.Lock()
	s.avatars[userID] = storedAvatar{name: avatar.Name, contentType: avatar.ContentType, data: data}
	s.mu.Unlock()

	return uint64(len(data)), avatar.Name, nil
}

func (s *ExampleServiceRPC) DownloadAvatar(ctx context.Context, userID uint64) (*WebrpcFile, error) {
	s.mu.Lock()
	avatar, ok := s.avatars[userID]
	s.mu.Unlock()

	if !ok {
		return nil, ErrWebrpcBadRoute.WithCausef("no avatar for user %d", userID)
	}

	return &WebrpcFile{
		Name:        avatar.name,
		ContentType: avatar.contentType,
		Size:        int64(len(avatar.data)),
		Body:        io.NopCloser(bytes.NewReader(avatar.data)),
	}, nil
}

func (s *ExampleServiceRPC) GetArticle(ctx context.Context, req GetArticleRequest) (*GetArticleResponse, error) {
	fmt.Println("GetArticle called with ByBN", req.ByBN.String())

	return &GetArticleResponse{
		Title:    fmt.Sprintf("Article %d", req.ArticleID),
		Content:  PtrTo(fmt.Sprintf("Hello, this is the content for article %d", req.ArticleID)),
		LargeNum: NewBigInt(req.ByBN.AsInt().Int64() * 2),
	}, nil
}
