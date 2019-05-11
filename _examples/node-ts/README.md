webrpc node-ts
==============

* Server: Nodejs (TypeScript)
* Client: Web Browser (Typescript)

example of generating a webrpc server and client from [service.ridl](./service.ridl) schema.

## Usage

1. Install nodejs + yarn
2. $ `make bootstrap` -- runs yarn on both ./server and ./webapp
3. $ `make run-server`
4. $ `make run-client`

Visit http://localhost:4444 !


## Notes

The cool thing about webrpc and other schema-driven rpc libraries (like grpc), is that you can
generate a Go client for this node server just by running:

`webrpc-gen -schema=service.ridl -target=go -pkg=proto -client -out=./proto/client.gen.go`

and tada, your Go programs now have full type definitions and network communication to the node
server!
