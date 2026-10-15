module github.com/webrpc/webrpc/_example/golang-sse

go 1.23

// replace github.com/golang-cz/ringbuf => ../../../../golang-cz/ringbuf

toolchain go1.24.2

require (
	github.com/go-chi/chi/v5 v5.3.2
	github.com/go-chi/cors v1.2.2
	github.com/go-chi/httplog/v3 v3.5.0
	github.com/golang-cz/devslog v0.0.17
	github.com/golang-cz/ringbuf v0.6.0
	github.com/stretchr/testify v1.12.1
)

require go.yaml.in/yaml/v3 v3.0.5 // indirect
