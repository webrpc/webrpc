package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Custom"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

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

func (s *ExampleServiceRPC) GetUser(ctx context.Context, userID uint64) (*User, error) {
	if userID == 911 {
		return nil, ErrorNotFound("user doest exist")
	}

	return &User{
		Id:       userID,
		Username: "hihi",
	}, nil
}

func (s *ExampleServiceRPC) Download(ctx context.Context, file string, stream DownloadStreamWriter) error {
	// TODO: the middleware.Logger in chi isn't suitable here as well, it will be reading all of this, wrapping it etc.
	// and we don't want that.. ideally chi logger after amount of bytes stops tracking and drops count

	req, _ := ctx.Value(HTTPRequestCtxKey).(*http.Request)
	fmt.Println("==??", req.Header.Get("X-Hi"))

	fmt.Println("req, file:", file)

	i := 0
	for {
		err := stream.Data(fmt.Sprintf("hiii send %d", i))
		if err == ErrStreamClosed {
			fmt.Println("client is gone.....")
			return nil
		}

		// if i == 2 {
		// 	panic("oh shit..")
		// }

		if err != nil {
			fmt.Println("ERR!!", err)
			return nil
		}
		time.Sleep(1 * time.Second)
		if i >= 100 {
			break
		}

		i += 1
	}

	// NOTE: stream will close on its own when we return actually, nice.
	// we can close ourselves too if we want though
	// stream.Close()

	return nil

}
