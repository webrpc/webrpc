# webrpc hello-webrpc-dart

* Server: Shelf (Dart)
* Client: Web Browser (Dart)

A simple example showcasing the use of webrpc dart for fullstack usage using [hello-api.ridl](hello-api.ridl). A more thorough example including flutter usage can be found [here](https://github.com/g5becks/Dart-WebRpc-Example).

## Usage

1. Install dart and run `pub global activate webdev`
2. $ `make bootstrap` -- runs   `pub get` on both ./server and ./client
3. $ `make run-server`
4. $ `make run-client`
5. To regenerate the client and/or server run `make generate`

Visit http://localhost:4444 !

## Notes

The cool thing about webrpc and other schema-driven rpc libraries (like grpc), is that you can
generate a Go client for this node server just by running:

`webrpc-gen -schema=service.ridl -target=go -pkg=proto -client -out=./proto/client.gen.go`

and tada, your Go programs now have full type definitions and network communication to the node
server!