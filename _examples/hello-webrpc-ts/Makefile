tools:
	GO111MODULE=off go get -u github.com/goware/webify

generate: generate-server generate-client

generate-server:
	../../bin/webrpc-gen -schema=hello-api.ridl -target=go -pkg=main -server -out=./api-server/hello_api.gen.go

generate-client:
	../../bin/webrpc-gen -schema=hello-api.ridl -target=ts -client -out=./webapp/src/client.gen.ts

serve-webapp:
	cd webapp && npm start

start-api-server:
	go run ./api-server
