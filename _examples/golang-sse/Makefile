all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

generate:
	webrpc-gen -schema=proto/chat.ridl -target=golang -pkg=proto -server -client -out=proto/chat.gen.go

dev-generate:
	../../bin/webrpc-gen -schema=proto/chat.ridl -target=golang -pkg=proto -server -client -out=proto/chat.gen.go

dev-generate-local-templates:
	../../bin/webrpc-gen -schema=proto/chat.ridl -target=golang -pkg=proto -server -client -out=proto/chat.gen.go

run:
	go run ./

test:
	go test -v ./

curl:
	curl -v -X POST -H "Content-Type: application/json" -H "Accept: application/x-ndjson" --data '{"username": "'$(shell whoami)'", "serverTimeoutSec": 2}' http://localhost:4848/rpc/Chat/SubscribeMessages
