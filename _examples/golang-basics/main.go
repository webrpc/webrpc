package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/webrpc/webrpc/_example/golang-basics/admin"

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
	webrpcHandler.OnError = func(r *http.Request, err *WebRPCError) {
		m, ok := MethodCtx(r.Context())

		if ok {
			_, ok = m.Annotations["deprecated"]
			if ok {
				fmt.Println(r.URL.Path, "deprecated")
			}
		}
	}
	webrpcHandler.OnRequest = func(w http.ResponseWriter, r *http.Request) error {
		m, ok := MethodCtx(r.Context())
		if !ok {
			return fmt.Errorf("could not find method context for request method: %s\n", r.URL.Path)
		}

		newEndpoint, ok := m.Annotations["deprecated"]
		if ok {
			return fmt.Errorf(
				"endpoint %s has been deprecated in favor of endpoint %s\n",
				r.URL.Path,
				newEndpoint,
			)
		}

		return nil
	}

	r.Handle("/admin/*", admin.NewAdminServiceServer(&AdminServiceRPC{}))
	r.Handle("/*", webrpcHandler)

	return http.ListenAndServe(":4242", r)
}

type AdminServiceRPC struct{}

func (*AdminServiceRPC) Auth(ctx context.Context) (string, string, error) {
	return "jwt", "admin", nil
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

	kind := Kind_ADMIN
	intent := Intent_openSession

	return 200, &User{
		ID:       userID,
		Username: "hihi",
		Kind:     kind,
		Intent:   intent,
	}, nil
}

func (s *ExampleServiceRPC) GetUserV2(ctx context.Context, header map[string]string, userID uint64) (uint32, *User, string, error) {
	if userID == 911 {
		return 0, nil, "", ErrUserNotFound
	}
	if userID == 31337 {
		return 0, nil, "", ErrorWithCause(ErrUserNotFound, fmt.Errorf("unknown user id %d", userID))
	}

	kind := Kind_ADMIN
	intent := Intent_openSession

	return 200, &User{
		ID:       userID,
		Username: "hihi",
		Kind:     kind,
		Intent:   intent,
	}, "https://www.google.com/images/john-doe.jpg", nil
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
