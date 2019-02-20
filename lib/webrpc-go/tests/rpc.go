//go:generate ../../../bin/webrpc-gen -schema=sample.ridl -target=go -pkg=proto -server -client -out=./proto/proto.gen.go
package tests

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/webrpc/webrpc/lib/webrpc-go"
	"github.com/webrpc/webrpc/lib/webrpc-go/tests/proto"
)

func NewRPCHandler() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	r.Handle("/rpc/ExampleRPC/*", proto.NewExampleRPCServer(&exampleRPC{}))
	r.Handle("/rpc/AnotherRPC/*", proto.NewAnotherRPCServer(&anotherRPC{}))

	return r
}

type exampleRPC struct{}

func (s *exampleRPC) Ping(ctx context.Context) error {
	return nil
}

func (s *exampleRPC) Status(ctx context.Context) (bool, error) {
	return true, nil
}

func (s *exampleRPC) GetUser(ctx context.Context, header map[string]string, userID uint64) (uint32, *proto.User, error) {
	if userID == 911 {
		return 0, nil, webrpc.ErrorNotFound("unknown userID %d", 911)
		// return nil, webrpc.Errorf(webrpc.ErrNotFound, "unknown userID %d", 911)
		// return nil, webrpc.WrapError(webrpc.ErrNotFound, err, "unknown userID %d", 911)
	}

	return 1, &proto.User{
		ID:       userID,
		Username: "hihi",
	}, nil
}

type anotherRPC struct{}

func (s *anotherRPC) Owner(ctx context.Context, q *string, id *uint64, desc string) (uint32, *proto.User, []string, error) {
	return 100, &proto.User{
		ID:       42,
		Username: "magic",
	}, []string{"hi"}, nil
}
