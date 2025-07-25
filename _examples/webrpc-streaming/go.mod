module github.com/webrpc/webrpc/_example/golang-sse

go 1.23

replace github.com/golang-cz/ringbuf => ../../../../golang-cz/ringbuf

toolchain go1.24.2

require (
	github.com/go-chi/chi/v5 v5.1.0
	github.com/go-chi/cors v1.2.1
	github.com/go-chi/httplog/v3 v3.2.2
	github.com/golang-cz/devslog v0.0.15
	github.com/golang-cz/ringbuf v0.0.4
	github.com/stretchr/testify v1.8.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
