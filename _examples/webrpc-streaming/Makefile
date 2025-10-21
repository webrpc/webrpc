all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

generate:
	webrpc-gen -schema=proto/chat.ridl -target=golang -pkg=proto -server -client -out=proto/chat.gen.go
	webrpc-gen -schema=proto/chat.ridl -target=typescript@v0.20.1 -client -out=webapp/src/rpc.gen.ts

dev-generate:
	../../bin/webrpc-gen -schema=proto/chat.ridl -target=golang -pkg=proto -server -client -out=proto/chat.gen.go
	../../bin/webrpc-gen -schema=proto/chat.ridl -target=typescript -client -out=webapp/src/rpc.gen.ts

dev-generate-local-templates:
	../../bin/webrpc-gen -schema=proto/chat.ridl -target=golang -pkg=proto -server -client -out=proto/chat.gen.go
	../../bin/webrpc-gen -schema=proto/chat.ridl -target=typescript -client -out=webapp/src/rpc.gen.ts

run-server:
	go run ./

run-webapp:
	cd webapp && pnpm run dev

test:
	go test -v ./

curl:
	curl -v -X POST -H "Content-Type: application/json" -H "Accept: application/x-ndjson" --data '{"username": "'$(shell whoami)'", "serverTimeoutSec": 2}' http://localhost:4848/rpc/Chat/SubscribeMessages
