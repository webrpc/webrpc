package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/webrpc/webrpc/_example/golang-nodejs/server/pkg/rpc"
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

	webrpcHandler := rpc.NewExampleServiceServer(&rpc.ExampleServiceRPC{})
	r.Handle("/*", webrpcHandler)

	log.Printf("Starting webrpc server on localhost:4242")

	return http.ListenAndServe(":4242", r)
}
