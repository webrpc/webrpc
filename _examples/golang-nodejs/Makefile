all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

install:
	go mod vendor && go mod tidy
	cd client && npm install

generate:
	webrpc-gen -schema=example.webrpc.json -target=golang -pkg=rpc -server -out=./server/pkg/rpc/server.gen.go
	webrpc-gen -schema=example.webrpc.json -target=javascript -client -out=./client/client.gen.mjs

run-server:
	go run ./server/cmd/api

run-client:
	cd ./client && pnpm run dev
