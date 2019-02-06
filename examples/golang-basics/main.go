package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/webrpc/webrpc/examples/golang-basics/proto"
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
		return nil, webrpc.ErrorNotFound("unknown userID %d", 911)
		// return nil, webrpc.Errorf(webrpc.ErrNotFound, "unknown userID %d", 911)
		// return nil, webrpc.WrapError(webrpc.ErrNotFound, err, "unknown userID %d", 911)
	}

	return &proto.User{
		ID:       1,
		Username: "hihi",
	}, nil
}
