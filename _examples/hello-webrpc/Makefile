all:
	@echo "please read Makefile source or README to see available commands"

tools:
	GO111MODULE=off go get -u github.com/goware/webify

generate: generate-server generate-client

generate-server:
	../../bin/webrpc-gen -schema=hello-api.ridl -target=go -pkg=main -server -out=./server/hello_api.gen.go

generate-client:
	../../bin/webrpc-gen -schema=hello-api.ridl -target=js -extra=noexports -client -out=./webapp/client.gen.js

run-server:
	go run ./server

run-client:
	webify -port=4444 -dir=./webapp

