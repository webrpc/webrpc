all:
	@echo "please read Makefile source to see available commands"

generate:
	../../bin/webrpc-gen -schema=example.webrpc.json -target=ts -server -out=./server/server.gen.ts
	../../bin/webrpc-gen -schema=example.webrpc.json -target=ts -client -out=./webapp/client.gen.ts

run-server:
	# go run ./server

run-client:
	cd ./client && npm start
