webrpc golang-file-upload example
=================================

* Server: Go
* Client: Go

An example of file upload and download endpoints built with the webrpc `file`
core type, see [./fileupload.ridl](./fileupload.ridl).

`file` and `[]file` can be used as method input arguments (or as top-level
fields of a struct used as a method input), and a method can return a single
`file` output (a download method).

In generated Go code, `file` maps to `*File`:

```go
type File struct {
	Name        string // filename, optional
	ContentType string // MIME type as claimed by the sender
	Size        int64  // -1 when unknown
	Body        io.ReadCloser
}
```

On the wire:

* Upload methods take a `multipart/form-data` request body: a `json` part
  carrying the JSON-encoded non-file arguments comes first, followed by one
  part per file (`[]file` is sent as repeated parts with the same name).
  The server caps request bodies at `Options.MaxUploadSize` (default 32 MiB).
* Download methods respond with the raw file body; `Content-Type`,
  `Content-Disposition`, and `Content-Length` are set from the returned
  `*File`. Errors still return the standard webrpc JSON error envelope.


## Running the example

* $ `cd _examples/golang-file-upload`
* $ `go run .` -- runs the server on :4242

Then, in another terminal window:

*Upload:*
```
curl -F 'json={"userId":1};type=application/json' -F 'avatar=@some-image.png' http://localhost:4242/rpc/FileService/UploadAvatar
```

*Download:*
```
curl -X POST -H "Content-Type: application/json" -d '{"userId":1}' -OJ http://localhost:4242/rpc/FileService/DownloadAvatar
```

You can also make strongly-typed requests through the generated Go client,
as done in [./example_test.go](./example_test.go). Run the tests with
`go test -v .`.

## TypeScript and JavaScript interoperability

The [./interop](./interop) folder generates a TypeScript client, a TypeScript
server, and a plain-JavaScript client from the same schema and runs an
end-to-end test suite of the file upload/download wire contract in all
directions (TypeScript and JavaScript clients vs Go server, TypeScript and
JavaScript clients vs TypeScript server, and Go client vs TypeScript server):

* $ `make test-interop`

In generated TypeScript code, file inputs are typed `Blob | File`, download
methods resolve to a `WebrpcFile` (metadata + streamed body), and server
handlers receive native `File` objects and return a `Blob | File` from
download methods.

The generated JavaScript client mirrors the TypeScript client's behavior:
file inputs are `Blob`/`File` values, and download methods resolve to a plain
`{ name, contentType, size, body, blob() }` object. The JavaScript Express
server does not support file endpoints yet; generating it from a schema with
file methods fails with a clear error at generation time.
