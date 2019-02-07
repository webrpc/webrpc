//go:generate ../../bin/webrpc-gen -schema=example.webrpc.json -target=go -pkg=main -server -client -out=./example.gen.go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
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

	webrpcHandler := NewExampleServiceServer(&ExampleServiceRPC{})
	r.Handle("/*", webrpcHandler)

	return http.ListenAndServe(":4242", r)
}

// TODO: lets move the service entire to tests

type ExampleServiceRPC struct {
}

func (s *ExampleServiceRPC) Ping(ctx context.Context) (*bool, error) {
	resp := false
	return &resp, nil
}

func (s *ExampleServiceRPC) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	if req.UserID == 911 {
		return nil, webrpc.ErrorNotFound("unknown userID %d", 911)
		// return nil, webrpc.Errorf(webrpc.ErrNotFound, "unknown userID %d", 911)
		// return nil, webrpc.WrapError(webrpc.ErrNotFound, err, "unknown userID %d", 911)
	}

	return &User{
		ID:       req.UserID,
		Username: "hihi",
	}, nil
}
