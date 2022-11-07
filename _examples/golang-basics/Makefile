all:
	@echo "please read Makefile source or README to see available commands"

generate:
	webrpc-gen -schema=example.ridl -target=golang -Pkg=main -Server -Client -out=./example.gen.go

dev-generate:
	../../bin/webrpc-gen -schema=example.ridl -target=golang -Pkg=main -Server -Client -out=./example.gen.go

dev-generate-local-templates:
	../../bin/webrpc-gen -schema=example.ridl -target=../../../gen-golang -Pkg=main -Server -Client -out=./example.gen.go
