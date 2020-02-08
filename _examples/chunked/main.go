package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	// stream support with chunked mode

	// terminating chunk..
	// 0\r\n
	// \r\n

	// TODO: send 1024 byts of arbitrary data before sending
	// the useful data

	r.Get("/chunked", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			panic("expected http.ResponseWriter to be an http.Flusher")
		}

		// w.Header().Set("X-Content-Type-Options", "nosniff")
		for {
			select {
			case <-r.Context().Done():
				fmt.Println("done")
				return
			default:
			}

			// for i := 1; i <= 10; i++ {
			fmt.Fprintf(w, "Chunk\n")
			flusher.Flush() // Trigger "chunked" encoding and send a chunk...
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("sending..")
		}
	})

	http.ListenAndServe(":3333", r)
}
