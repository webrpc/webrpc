all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

generate:
	webrpc-gen -schema=example.webrpc.json -target=golang -pkg=main -server -out=./server/server.gen.go
	webrpc-gen -schema=example.webrpc.json -target=javascript -client -out=./client/client.gen.mjs

run-server:
	go run ./server

run-client:
	cd ./client && npm start
