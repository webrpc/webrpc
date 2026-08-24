# Use webrpc-gen binary from the root of the project.
export PATH = $(shell echo $$PWD/../../bin:$$PATH)

all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

build:
	make -C ../../ build

generate: build
	webrpc-gen -schema=fileupload.ridl -target=golang -pkg=main -server -client -out=./fileupload.gen.go
	webrpc-gen -schema=fileupload.ridl -target=typescript -client -out=./interop/client.gen.ts
	webrpc-gen -schema=fileupload.ridl -target=typescript -server -out=./interop/server.gen.ts
	webrpc-gen -schema=fileupload.ridl -target=javascript -client -out=./interop/client.gen.mjs

test:
	go test -v -race ./...

# End-to-end interoperability tests of the file upload/download wire contract
# between the Go and TypeScript generators, in both directions.
test-interop:
	cd interop && pnpm install && pnpm test

run:
	go run .
