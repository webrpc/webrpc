package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

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
		return nil, ErrorNotFound("fdf")
	}

	return &User{
		Id:       userID,
		Username: "hihi",
	}, nil
}

// func (s *ExampleServiceRPC) Upload(ctx context.Context, reader UploadReader) (bool, error) {
// 	return false, nil
// }

func (s *ExampleServiceRPC) Download(ctx context.Context, file string, stream DownloadStreamWriter) error {
	// TODO: the middleware.Logger in chi isn't suitable here as well, it will be reading all of this, wrapping it etc.
	// and we don't want that.. ideally chi logger after amount of bytes stops tracking and drops count

	i := 0
	for {
		err := stream.Data(fmt.Sprintf("hiii send %d", i))
		if err == ErrStreamClosed {
			fmt.Println("client is gone.....")
			return nil
		}

		if err != nil {
			fmt.Println("ERR!!", err)
			return nil
		}
		time.Sleep(1 * time.Second)
		if i >= 6 {
			break
		}

		i += 1
	}

	// NOTE: stream will close on its own when we return actually, nice.
	// we can close ourselves too if we want though
	// stream.Close()

	return nil

}
