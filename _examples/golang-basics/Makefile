all:
	@echo "please read Makefile source or README to see available commands"

generate:
	webrpc-gen -schema=example.ridl -target=go -pkg=main -server -client -out=./example.gen.go

dev-generate:
	../../bin/webrpc-gen -schema=example.ridl -target=go -pkg=main -server -client -out=./example.gen.go
