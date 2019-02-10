tools:
	GO111MODULE=off go get -u github.com/goware/webify

generate: generate-server generate-client

generate-server:
	../../bin/webrpc-gen -schema=hello-api.webrpc.json -target=go -pkg=main -server -out=./api-server/hello_api.gen.go

generate-client:
	../../bin/webrpc-gen -schema=hello-api.webrpc.json -target=js -extra=noexports -client -out=./webapp/client.gen.js

serve-webapp:
	webify -port=4444 -dir=./webapp

start-api-server:
	go run ./api-server
