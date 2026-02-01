# Use webrpc-gen binary from the root of the project.
export PATH = $(shell echo $$PWD/../../bin:$$PATH)

all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

build:
	make -C ../../ build

format:
	go run github.com/webrpc/ridlfmt -w example.ridl

generate: build
	webrpc-gen -schema=example.ridl -target=golang -pkg=main -server -client -service="Example" -out=./example.gen.go
	webrpc-gen -schema=example.ridl -target=golang -pkg=admin -server -client -service="Admin" -ignore="@deprecated" -out=./admin/admin.gen.go
	webrpc-gen -schema=example.ridl -target=golang -pkg=internal -server -client -service="Example,Admin" -ignore="@public,@deprecated" -out=./internal/internal.gen.go

generate-local-templates: build
	webrpc-gen -schema=example.ridl -target=golang -pkg=main -server -client -service="Example" -ignore="@deprecated" -out=./example.gen.go
	webrpc-gen -schema=example.ridl -target=golang -pkg=admin -server -client -service="Admin" -ignore="@deprecated" -out=./admin/admin.gen.go
	webrpc-gen -schema=example.ridl -target=golang -pkg=internal -server -client -service="Example,Admin" -ignore="@public,@deprecated" -out=./internal/internal.gen.go

test:
	go test -v -race ./...
