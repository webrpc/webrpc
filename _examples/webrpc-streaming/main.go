package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", port),
		Handler: rpc.Router(),
	}

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		defer cancel()

		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	slog.Info("shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rpc.close(shutdownCtx)

	// Finish active connections.
	err := server.Shutdown(shutdownCtx)
	if err != nil {
		log.Fatal(fmt.Errorf("shut down HTTP server: %w", err))
	}

	slog.Info("shut down")
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
