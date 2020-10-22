webrpc golang-streaming example
============================

* Server: Go
* Client: Go

A simple example of a Go web service built using webrpc.

The process of developing something like this is..

1. Start with your webrpc schema file, in this case, [./example.ridl](./example.ridl) in RIDL format.. or,
you can also write your schema in JSON format like so, [./example.webrpc.json](./example.webrpc.json). RIDL is simpler :)
2. Design your schema file and think about the methods calls clients will need to make
to your service
3. Write the "services" section of the schema file
4. From the inputs and outputs for the function definitions, start writing the "messages"
section of the data types needed in your program.
5. Run the code generator to build the server and client:
  * `webrpc-gen -schema=example.ridl -server -client -target=go -pkg main -out=example.gen.go`
  * however, in this example we put it inside a `go:generate`, so you can run `go generate .`
6. Write your server ([./main.go](./main.go)) and implement the `ExampleServiceRPC` interface type
that was created by the code generator, and located in the [gen'd file](./example.gen.go).
7. Enjoy!

Generating a Typescript client will soon be available by running:
* `webrpc-gen -schema=example.ridl -target=ts -pkg=example -client -out=./example-client.ts`
* check out the [hello-webrpc](../hello-webrpc). 

## Testing the example

You can run the tests if you want with `go test -v .`.


### Running the example

* $ `cd _examples/golang-streaming`

* $ `go run .` -- runs the server

..or, if you want to see the internals, lets run some curl commands manually. In another terminal
window, run some raw curl commands:

*Request:*
```
curl -v -X POST -H"Content-Type: application/json" -v -d '{}' http://localhost:4242/rpc/ExampleService/Download
```

*Response:*
```
{"base64":"hiii send 3"}
```

### How it works

Please read all of the source in this folder :) including example_test.go
