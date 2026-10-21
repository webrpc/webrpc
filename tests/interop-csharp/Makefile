# Use webrpc-gen binary from the root of the project.
export PATH = $(shell echo $$PWD/../../bin:$$PATH)

.PHONY: all build generate test

all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

build:
	make -C ../../ build

generate: build
	webrpc-gen -schema=../schema/test.ridl -target=csharp -client -namespace=Interop.Test -out=./lib/stj/test.gen.cs
	webrpc-gen -schema=./csharp.ridl -target=csharp -client -namespace=Interop.CSharp -out=./lib/stj/csharp.gen.cs
	# Newtonsoft.Json with internal visibility, as embedded in a Unity SDK. The
	# CSharp schema covers every serializer-specific code path.
	webrpc-gen -schema=./csharp.ridl -target=csharp -client -json=newtonsoft -visibility=internal -namespace=Interop.CSharp -out=./lib/newtonsoft/csharp.gen.cs
	webrpc-gen -schema=./csharp.ridl -target=golang -pkg=main -server -out=./server/server.gen.go

# End-to-end interoperability tests of the C# client generator (requires the .NET SDK).
test:
	./test.sh
