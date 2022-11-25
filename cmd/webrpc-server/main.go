package main

import (
	"net/http"
	"github.com/webrpc/webrpc/tests/server"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/complex/rpc/ComplexApi/GetComplex", http.StripPrefix("/complex", server.NewComplexApiServer(&server.ComplexServer{})))
	mux.Handle("/complex/rpc/ComplexApi/SendComplex", http.StripPrefix("/complex", server.NewComplexApiServer(&server.ComplexServer{})))

	// shutdown process through http call
	mux.Handle("/shutdown", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		os.Exit(0)
	}))

	err := http.ListenAndServe(":9999", mux)
	if err != nil {
		panic(err)
	}
}
