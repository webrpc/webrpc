package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/webrpc/webrpc/lib/webrpc-go"
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

	webrpcHandler := NewExampleServiceServer(&ExampleServiceRPC{})
	r.Handle("/*", webrpcHandler)

	log.Printf("Starting webrpc server on localhost:4242")

	return http.ListenAndServe(":4242", r)
}

type ExampleServiceRPC struct {
}

func (s *ExampleServiceRPC) Ping(ctx context.Context) (bool, error) {
	return true, nil
}

func (s *ExampleServiceRPC) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	if req.UserID == 911 {
		return nil, webrpc.ErrorNotFound("unknown userID %d", 911)
	}

	return &User{
		ID:       req.UserID,
		Username: "hihi",
	}, nil
}
