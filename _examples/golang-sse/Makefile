all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

generate:
	webrpc-gen -schema=example.ridl -target=golang -pkg=main -server -client -out=./example.gen.go

dev-generate:
	../../bin/webrpc-gen -schema=example.ridl -target=golang -pkg=main -server -client -out=./example.gen.go

dev-generate-local-templates:
	../../bin/webrpc-gen -schema=example.ridl -target=golang -pkg=main -server -client -out=./example.gen.go
