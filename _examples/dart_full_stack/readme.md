# WebRpc Dart

* Shelf Server

* Flutter Client

* Generated Bloc Library code

  This example showcases using webrpc to generate a shelf server and a client including [bloc library](https://bloclibrary.dev/#/) using the example.ridl schema file.

## Usage

1. Install go-task using the installation for your os found [here](https://taskfile.dev/#/installation)
2. $ task bootstrap -- runs pub get and build_runner for client and server
3. $ task run_server
4. $ task run_client
5. To generate a client without bloc library code delete the argument -extra="flutter_bloc" from the build_client task
6. To generate a client with bloc library for angular dart add the line -extra="angular_bloc" to the build_client task
7. To change the name of the part files generated for [freezed](https://pub.dev/packages/freezed) and [json_serializable](https://pub.dev/packages/json_serializable) change the -pkg argument to match the name of file used in the part directive.

## Notes

The cool thing about webrpc and other schema-driven rpc libraries (like grpc), is that you can generate a Go client for this node server just by running:

```
webrpc-gen -schema=service.ridl -target=go -pkg=proto -client -out=./proto/client.gen.go
```

and tada, your Go programs now have full type definitions and network communication to the node server!

