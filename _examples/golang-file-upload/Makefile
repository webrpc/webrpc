# Use webrpc-gen binary from the root of the project.
export PATH = $(shell echo $$PWD/../../bin:$$PATH)

all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

build:
	make -C ../../ build

generate: build
	webrpc-gen -schema=fileupload.ridl -target=golang -pkg=main -server -client -out=./fileupload.gen.go

test:
	go test -v -race ./...

run:
	go run .
