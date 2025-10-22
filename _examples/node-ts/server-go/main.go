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
		// Allow all origins for local development; replace with specific domains as needed.
		AllowedOrigins: []string{"*"},
		// Allow standard methods; include OPTIONS for preflight.
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// Allow headers commonly used in APIs.
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		// MaxAge indicates how long (in seconds) the results of a preflight request can be cached.
		MaxAge: int((12 * time.Hour).Seconds()),
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
		return nil, ErrorWithCause(ErrUserNotFound, fmt.Errorf("unknown user id %d", req.UserID))
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
			Balance:  NewBigInt(1234),
			Extra: &Extra{
				Info:   "some info",
				Amount: NewBigInt(5678),
				Points: []BigInt{NewBigInt(10), NewBigInt(20), NewBigInt(30)},
			},
		},
	}, nil
}

func (s *ExampleServiceRPC) GetUserV2(ctx context.Context, req GetUserRequest) (*GetUserResponse, error) {
	if req.UserID == 911 {
		return nil, ErrUserNotFound
	}
	if req.UserID == 31337 {
		return nil, ErrorWithCause(ErrUserNotFound, fmt.Errorf("unknown user id %d", req.UserID))
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
		Profile:  "https://www.google.com/images/john-doe.jpg",
		LargeNum: NewBigInt(31337),
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
