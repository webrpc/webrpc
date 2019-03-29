//go:generate ../../bin/webrpc-gen -schema=example.ridl -target=go -pkg=main -server -client -out=./example.gen.go
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/webrpc/webrpc/lib/webrpc-go"
)

// NOTES: after considering this design, it feels like streaming adds a lot of unnecessary
// complexity to webrpc and therefore, at this point I don't see the benefits to including it here.
// instead, devs can use webrpc's generated message types, and write their own streaming code

func main() {
	go func() {
		err := startServer()
		if err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(1000 * time.Millisecond)

	// TODO: client can pass headers too..
	// such as for the auth of opening the connection..

	client, err := NewExampleServiceWSClient("localhost:4242")
	if err != nil {
		log.Fatal("dial:", err)
	}

	time.Sleep(100 * time.Millisecond)

	b, err := client.Ping(context.Background())
	if err != nil {
		log.Fatal("ping", err)
	}
	log.Println("ping", b)
}

func startServer() error {
	// TODO: for websocket transport -- what do we do about the middleware
	// stack... ugh.... doesn't belong here.
	// 1. logging
	// 2. requestID (interesting..)
	// 3. recoverer from panic ..
	// 4. telemetry.. etc..
	// 5. rate-limiting..
	// 6. caching ..

	// options..
	// we come up with RequestHandler
	// or, WebRPCHandler ......
	//

	// TODO: look at go-kit
	// and think about http2 streaming, how are headers etc. sent each request/middleware?

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// on WS Open, we can pass a request header, perhaps
			// we should be doing this, and passing the auth token, or other?
			// or, we need client to call Auth() and we track this in server ..?
			log.Println("mw.........")
			next.ServeHTTP(w, r)
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	webrpcHandler := NewExampleServiceServer(&ExampleServiceRPC{})
	// TODO: attach middlewares here..........?

	r.Handle("/*", webrpcHandler)

	return http.ListenAndServe(":4242", r)
}

type ExampleServiceRPC struct {
}

func (s *ExampleServiceRPC) Ping(ctx context.Context) (bool, error) {
	return true, nil
}

func (s *ExampleServiceRPC) GetUser(ctx context.Context, userID uint64) (*User, error) {
	if userID == 911 {
		return nil, webrpc.ErrorNotFound("unknown userID %d", 911)
		// return nil, webrpc.Errorf(webrpc.ErrNotFound, "unknown userID %d", 911)
		// return nil, webrpc.WrapError(webrpc.ErrNotFound, err, "unknown userID %d", 911)
	}

	return &User{
		Id:       userID,
		Username: "hihi",
	}, nil
}

func (s *ExampleServiceRPC) Auth(ctx context.Context, hash string) (bool, error) {
	return false, nil
}

func (s *ExampleServiceRPC) Chat(ctx context.Context, msg *ChatMsg) (*ChatMsg, error) {
	return nil, nil
}
