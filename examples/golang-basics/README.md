# webrpc golang-basics example

A simple example of a Go web service built using webrpc.

The process of developing something like this is..

1. 

## Testing the example

You can run the tests if you want with `go test -v .`.


### Running the example

* $ `cd examples/golang-basics`

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
{"id":1234,"USERNAME":"hihi","created_at":"0001-01-01T00:00:00Z"}
```


### How it works

Please read all of the source in this folder :) including example_test.go
