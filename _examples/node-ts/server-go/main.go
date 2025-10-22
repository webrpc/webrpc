package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
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

	webrpcHandler := NewExampleServer(&ExampleServiceRPC{})

	r.Handle("/*", webrpcHandler)

	return http.ListenAndServe(":3000", r)
}

type ExampleServiceRPC struct{}

func (s *ExampleServiceRPC) Ping(ctx context.Context) error {
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

func (s *ExampleServiceRPC) GetArticle(ctx context.Context, req GetArticleRequest) (*GetArticleResponse, error) {

	return &GetArticleResponse{
		Title:    fmt.Sprintf("Article %d", req.ArticleID),
		Content:  PtrTo(fmt.Sprintf("Hello, this is the content for article %d", req.ArticleID)),
		LargeNum: NewBigInt(999999999999),
	}, nil
}
