all:
	@echo "please read Makefile source or README to see available commands"

tools:
	GO111MODULE=off go get -u github.com/goware/webify

generate: generate-server generate-client

generate-server:
	../../bin/webrpc-gen -schema=hello-api.ridl -target=dart -pkg=hello_api.gen -server -out=./server/lib/hello_api.gen.dart && cd server && pub run build_runner build --delete-conflicting-outputs && dartfmt -w bin lib

generate-client:
	../../bin/webrpc-gen -schema=hello-api.ridl -target=dart -client -out=./client/lib/client.gen.dart -pkg=client.gen && cd client && pub run build_runner build --delete-conflicting-outputs && dartfmt -w lib

bootstrap:
	cd server && pub get && cd ../client && pub get

run-server:
	go run ./server

run-client:
	cd webapp && npm start
