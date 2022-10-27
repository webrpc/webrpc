all:
	@echo "please read Makefile source or README to see available commands"

generate:
	webrpc-gen -schema=example.webrpc.json -target=golang -pkg=main -server -out=./server/server.gen.go
	webrpc-gen -schema=example.webrpc.json -target=javascript -client -out=./client/client.gen.mjs

run-server:
	go run ./server

run-client:
	cd ./client && npm start
