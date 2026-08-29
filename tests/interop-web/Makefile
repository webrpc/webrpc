# Use webrpc-gen binary from the root of the project.
export PATH = $(shell echo $$PWD/../../bin:$$PATH)

RIDL = ../../_examples/golang-basics/example.ridl

all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

build:
	make -C ../../ build

generate: build
	webrpc-gen -schema=$(RIDL) -target=typescript -client -service=Example -out=./client.gen.ts
	webrpc-gen -schema=$(RIDL) -target=typescript -server -service=Example -out=./server.gen.ts
	webrpc-gen -schema=$(RIDL) -target=javascript -client -service=Example -out=./client.gen.mjs

# End-to-end interoperability tests of the file upload/download wire contract
# between the Go and TypeScript/JavaScript generators.
test:
	npm install --no-audit --no-fund --no-package-lock
	npm test
