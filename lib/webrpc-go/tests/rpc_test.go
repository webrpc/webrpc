//go:generate ../../../bin/webrpc-gen -schema=sample.webrpc.json -target=go -pkg=proto -server -client -out=./proto/proto.gen.go
package golang_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-chi/chi"
	"github.com/webrpc/webrpc/lib/webrpc-go"
	"github.com/webrpc/webrpc/lib/webrpc-go/tests/proto"
)

func NewWebService() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	r.Handle("/rpc/ExampleRPC/*", proto.NewExampleRPCServer(&exampleRPC{}))
	r.Handle("/rpc/AnotherRPC/*", proto.NewAnotherRPCServer(&anotherRPC{}))

	return r
}

type exampleRPC struct{}

func (s *exampleRPC) Ping(ctx context.Context) (*bool, error) {
	resp := false
	return &resp, nil
}

func (s *exampleRPC) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.User, error) {
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

type anotherRPC struct{}

func (s *anotherRPC) Owner(ctx context.Context) (*proto.User, error) {
	return &proto.User{
		ID:       42,
		Username: "magic",
	}, nil
}

func TestStuff(t *testing.T) {

}
