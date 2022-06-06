generate:
	go generate ./...

generate-from-ridl:
	../../bin/webrpc-gen -schema=example.ridl -target=go -pkg=main -server -client -out=./example.gen.go

generate-from-json:
	../../bin/webrpc-gen -schema=example.webrpc.json -target=go -pkg=main -server -client -out=./example.gen.go
