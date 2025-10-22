generate:
	webrpc-gen -schema=../service.ridl -target=golang -pkg=main -server -client -service="Example" -out=./example.gen.go

test:
	go test -v -race ./...

run:
	go run .

generate-local:
	webrpc-gen -schema=../service.ridl -target=../../../../gen-golang -pkg=main -server -client -out=./example.gen.go
