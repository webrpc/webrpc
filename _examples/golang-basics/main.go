//go:generate ../../bin/webrpc-gen -schema=example.ridl -target=go -pkg=main -server -client -out=./example.gen.go
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	r.Handle("/*", webrpcHandler)

	return http.ListenAndServe(":4242", r)
}

type ExampleServiceRPC struct {
}

func init() {
	// err := RPCError{Code: 123}
	// // hmm...
	// if errors.Is(err, RPCError{}) {
	// 	fmt.Println("yep")
	// }

	err := ErrorWithCause(ErrInvalidName, Errorf("yes: %w", ErrMemoryFull))
	if errors.Is(err, ErrInvalidName) {
		fmt.Println("yep")
	}
	if errors.Is(err, ErrMemoryFull) {
		fmt.Println("ahaha")
	}
	spew.Dump(err)
	fmt.Println("??", err.Unwrap())

	// err := Errorf("hiiii: %w", ErrInvalidName)
	// if errors.Is(err, ErrInvalidName) {
	// 	fmt.Println("kkff..")
	// }
}

func (s *ExampleServiceRPC) Ping(ctx context.Context) error {
	return ErrInvalidName
	// return Errorf("dfdfdf")

	// return ErrorWithCause(ErrInvalidName, err)
	// return ErrorWithCause(Errorf(), err)

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

func (s *ExampleServiceRPC) GetUser(ctx context.Context, header map[string]string, userID uint64) (uint32, *User, error) {
	// HMM.. so, what are the http status codes we'll use..? lets copy json-rpc
	// which is..
	// 200 for OK
	// 400 for any general error
	// 404 is rpc method is not found
	// 401/403 for unauthorized requests..
	// 500 of course can happen

	if userID == 911 {
		// return 0, nil, WrapError(ErrInternal, errors.New("bad"), "app msg here")
		return 0, nil, Errorf("hihi")
		// return 0, nil, ErrorNotFound("unknown userID %d", 911)
		// return 0, nil, Errorf(ErrNotFound, "unknown userID %d", 911)
		// return 0, nil, WrapError(ErrNotFound, nil, "unknown userID %d", 911)
	}

	return 200, &User{
		ID:       userID,
		Username: "hihi",
	}, nil
}

func (s *ExampleServiceRPC) FindUser(ctx context.Context, f *SearchFilter) (string, *User, error) {
	name := f.Q
	return f.Q, &User{
		ID: 123, Username: name,
	}, nil
}
