all:
	@echo "please read Makefile source to see available commands"

generate:
	../../bin/webrpc-gen -schema=example.webrpc.json -target=go -pkg=main -server -out=./server/server.gen.go
	../../bin/webrpc-gen -schema=example.webrpc.json -target=js -client -out=./client/client.gen.mjs

run-server:
	go run ./server

run-client:
	cd ./client && npm start
