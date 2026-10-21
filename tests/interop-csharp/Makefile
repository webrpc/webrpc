# Use webrpc-gen binary from the root of the project.
export PATH = $(shell echo $$PWD/../../bin:$$PATH)

.PHONY: all build generate test

all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

build:
	make -C ../../ build

generate: build
	for json in stj newtonsoft; do \
		webrpc-gen -schema=../schema/test.ridl -target=csharp -client -json=$$json -namespace=Interop.Test -out=./lib/$$json/test.gen.cs || exit 1; \
		webrpc-gen -schema=../../_examples/golang-basics/example.ridl -target=csharp -client -json=$$json -namespace=Interop.Example -out=./lib/$$json/example.gen.cs || exit 1; \
		webrpc-gen -schema=./csharp.ridl -target=csharp -client -json=$$json -namespace=Interop.CSharp -out=./lib/$$json/csharp.gen.cs || exit 1; \
	done
	# Compile check of the -visibility=internal output and of the disabled header options.
	webrpc-gen -schema=./csharp.ridl -target=csharp -client -namespace=Interop.Internal -visibility=internal -schemaHash=false -webrpcHeader=false -out=./lib/stj/internal.gen.cs
	webrpc-gen -schema=./csharp.ridl -target=golang -pkg=main -server -out=./server/server.gen.go

# End-to-end interoperability tests of the C# client generator (requires the .NET SDK).
test:
	./test.sh
