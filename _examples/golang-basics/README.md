# webrpc golang-basics example

A simple example of a Go web service built using webrpc.

The process of developing something like this is..

1. Start with your webrpc schema file, in this case, [./example.webrpc.json](./example.webrpc.json)
2. Design your schema file and think about the methods calls clients will need to make
to your service
3. Write the "services" section of the schema file
4. From the inputs and outputs for the function definitions, start writing the "messages"
section of the data types needed in your program.
5. Run the code generator to build the server and client:
  * `webrpc-gen -schema=example.webrpc.json -target=go -pkg=main -server -client -out=./example.gen.go`
  * however, in this example we put it inside a `go:generate`, so you can run `go generate .`
6. Write your server ([./main.go](./main.go)) and implement the `ExampleServiceRPC` interface type
that was created by the code generator, and located in the [gen'd file](./example.gen.go).
7. Enjoy!

Next steps, you can generate a Typescript client by running:
* `webrpc-gen -schema=example.webrpc.json -target=ts -pkg=example -client -out=./example-client.ts`
* check out the [hello-webrpc](../hello-webrpc) for an example with a Webapp client talking to a webrpc backend


## Testing the example

You can run the tests if you want with `go test -v .`.


### Running the example

* $ `cd _examples/golang-basics`

* $ `go run .` -- runs the server

You can make strongly-typed requests to the server through the generated Go client
as done inside of example_test.go, as in..

```go
resp, err := client.GetUser(context.Background(), &GetUserRequest{
  UserID: 1234,
})
spew.Dump(err)
spew.Dump(resp)
```

..or, if you want to see the internals, lets run some curl commands manually. In another terminal
window, run some raw curl commands:

*Request:*
```
curl -v -X POST -H"Content-Type: application/json" -v -d '{"userID":1234}' http://localhost:4242/rpc/ExampleService/GetUser
```

*Response:*
```
{"id":1234,"USERNAME":"hihi"}
```


### How it works

Please read all of the source in this folder :) including example_test.go
