all:
	@echo "please read Makefile source or README to see available commands"

generate:
	webrpc-gen -schema=example.webrpc.json -target=go -pkg=main -server -out=./server/server.gen.go
	webrpc-gen -schema=example.webrpc.json -target=js -client -out=./client/client.gen.mjs

run-server:
	go run ./server

run-client:
	cd ./client && npm start
