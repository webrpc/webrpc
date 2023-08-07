package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	return http.ListenAndServe(":4242", r)
}

type ExampleServiceRPC struct {
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

func (s *ExampleServiceRPC) GetUser(ctx context.Context, header map[string]string, userID uint64) (uint32, *User, error) {
	if userID == 911 {
		return 0, nil, ErrUserNotFound
	}
	if userID == 31337 {
		return 0, nil, ErrorWithCause(ErrUserNotFound, fmt.Errorf("unknown user id %d", userID))
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

// func (s *ExampleServiceRPC) SearchUsers(ctx context.Context, username *string, query string) ([]*User, error) {
// 	return nil, nil
// }

// NOTE: this is what we want the generator to produce, but it doesn't yet
func (s *ExampleServiceRPC) SearchUsers(ctx context.Context, username *string, query QueryString) ([]User, error) {
	if query == "admin" {
		return []User{
			{ID: 1, Username: "admin"},
		}, nil
	} else if username != nil && *username == "guest" {
		return []User{
			{ID: 2, Username: "guest"},
		}, nil
	} else {
		return []User{
			{ID: 3, Username: "some user"},
		}, nil
	}
}

// QueryString is a special type override which we use in example.ridl with
// + go.field.type = QueryString
type QueryString string
