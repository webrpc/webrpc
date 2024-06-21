all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

tools:
	GO111MODULE=off go get -u github.com/goware/webify

generate: generate-server generate-client

generate-server:
	webrpc-gen -schema=hello-api.ridl -target=golang -pkg=main -server -out=./server/hello_api.gen.go

generate-client:
	webrpc-gen -schema=hello-api.ridl -target=javascript -exports=false -client -out=./webapp/client.gen.js

run-server:
	go run ./server

run-client:
	go run github.com/goware/webify -host=localhost -port=4444 -dir=./webapp

