# Use webrpc-gen binary from the root of the project.
export PATH = $(shell echo $$PWD/../../bin:$$PATH)

all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

build:
	make -C ../../ build

generate: generate-server generate-client

generate-server: build
	webrpc-gen -schema=hello-api.ridl -target=golang -pkg=main -server -out=./server/hello_api.gen.go

generate-client: build
	webrpc-gen -schema=hello-api.ridl -target=typescript -client -out=./webapp/client.gen.ts

bootstrap:
	rm -rf webapp/node_modules
	cd webapp && pnpm install

run-server:
	go run ./server

run-client:
	cd webapp && pnpm run dev
