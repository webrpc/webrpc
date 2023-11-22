package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//go:generate go run ../../cmd/webrpc-gen -schema=chat.ridl -target=../../gen-golang -pkg=main -server -client -out=chat.gen.go

func main() {
	err := startServer(4848)
	if err != nil {
		log.Fatal(err)
	}
}

func startServer(port int) error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	webrpcHandler := NewChatServer(&ChatServer{})
	r.Handle("/*", webrpcHandler)

	return http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}

type ChatServer struct {
}

func (s *ChatServer) SendMessage(ctx context.Context, authorName string, text string) error {
	return nil
}

func (s *ChatServer) GetNewMessages(ctx context.Context) (*Message, error) {
	return &Message{}, nil
}
