package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/webrpc/webrpc-go/_example/proto"
)

func main() {
	err := startServer()
	if err != nil {
		log.Fatal(err)
	}
}

func startServer() error {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi post"))
	})

	webrpcHandler := proto.NewExampleServiceServer(&ExampleService{})
	r.Handle("/*", webrpcHandler)

	return http.ListenAndServe(":4242", r)
}

type ExampleService struct {
}

func (s *ExampleService) Ping(ctx context.Context) (*bool, error) {
	resp := false
	return &resp, nil
}

func (s *ExampleService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.User, error) {
	return &proto.User{
		ID:       1,
		Username: "hihi",
	}, nil
}
