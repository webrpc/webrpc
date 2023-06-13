all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

install:
	go mod vendor && go mod tidy
	cd client && npm install

generate:
	webrpc-gen -schema=example.webrpc.json -target=golang -pkg=main -server -out=./server/server.gen.go
	webrpc-gen -schema=example.webrpc.json -target=javascript -client -out=./client/client.gen.mjs

run-server:
	go run lsserver/cmd/api.go

run-client:
	cd ./client && npm start
