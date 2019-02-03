package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/webrpc/webrpc-go"
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

// TODO: lets move the service entire to tests

type ExampleService struct {
}

func (s *ExampleService) Ping(ctx context.Context) (*bool, error) {
	resp := false
	return &resp, nil
}

func (s *ExampleService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.User, error) {
	if req.UserID == 911 {
		// let's intentionally fail on this one
		// TODO: lets return proper error code of 404
		return nil, webrpc.Error(errors.New("unknown userID"))
	}

	return &proto.User{
		ID:       1,
		Username: "hihi",
	}, nil
}
