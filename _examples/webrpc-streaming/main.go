package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog/v3"
	"github.com/webrpc/webrpc/_example/golang-sse/proto"
)

func main() {
	port := 4848

	rpc := NewChatServer()

	slog.SetDefault(rpc.logger)
	slog.Info(fmt.Sprintf("serving at http://localhost:%v", port))

	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), rpc.Router())
	if err != nil {
		log.Fatal(err)
	}
}

func (s *ChatServer) Router() http.Handler {
	r := chi.NewRouter()

	// Request logger
	r.Use(httplog.RequestLogger(s.logger, &httplog.Options{
		Level:           slog.LevelInfo,
		Schema:          httplog.SchemaECS.Concise(true),
		RecoverPanics:   true,
		LogRequestBody:  func(req *http.Request) bool { return true },
		LogResponseBody: func(req *http.Request) bool { return true },
	}))

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Webrpc"},
		ExposedHeaders:   []string{"Link", "Webrpc"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	webrpcHandler := proto.NewChatServer(s)
	r.Handle("/*", webrpcHandler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	return r
}
