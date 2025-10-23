webrpc node-ts
==============

* Server: Nodejs (TypeScript)
* Client: Web Browser (Typescript)

example of generating a webrpc server and client from [service.ridl](./service.ridl) schema.

## Usage

1. Install nodejs + pnpm
2. $ `make bootstrap` -- runs pnpm on both ./server and ./webapp
3. $ `make run-server`
4. $ `make run-client`

Visit http://localhost:4444 !

## Recommendation

You can find webrpc nodejs server examples in a variety of frameworks including:
* ./server -- just vanilla nodejs (no framework)
* ./server-hono -- built on hono library
* ./server-fastify -- built on fastify library

Our recommendation is to use hono over fastify, as hono is much cleaner. The cool thing
is that the server.gen.ts generated webrpc code is the same in all examples, its just
the http service router (ie. none/hono/fastify) is different.

## Notes

The cool thing about webrpc and other schema-driven rpc libraries (like grpc), is that you can
generate a Go client for this node server just by running:

`webrpc-gen -schema=service.ridl -target=golang -pkg=proto -client -out=./proto/client.gen.go`

and tada, your Go programs now have full type definitions and network communication to the node
server!
