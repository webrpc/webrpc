package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/webrpc/webrpc/_example/golang-sse/proto"
)

//go:generate go run ../../cmd/webrpc-gen -schema=proto/chat.ridl -target=../../gen-golang -pkg=proto -server -client -out=proto/chat.gen.go

func main() {
	port := 4848

	slog.Info(fmt.Sprintf("serving at http://localhost:%v", port))

	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), router())
	if err != nil {
		log.Fatal(err)
	}
}

func router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var reqBody bytes.Buffer
			r.Body = io.NopCloser(io.TeeReader(r.Body, &reqBody))

			var respBody bytes.Buffer
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			ww.Tee(&respBody)

			slog.Info(fmt.Sprintf("req started"),
				slog.String("url", fmt.Sprintf("%v %v", r.Method, r.URL.String())))

			defer func() {
				slog.Info(fmt.Sprintf("req finished HTTP %v", ww.Status()),
					slog.String("url", fmt.Sprintf("%v %v", r.Method, r.URL.String())),
					slog.String("reqBody", reqBody.String()),
					slog.String("respBody", respBody.String()),
				)
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	})
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	rpc := NewChatServer()
	webrpcHandler := proto.NewChatServer(rpc)
	r.Handle("/*", webrpcHandler)

	return r
}
