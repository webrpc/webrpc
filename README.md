webrpc
======

webrpc, is a lightweight client+server RPC framework tailored for modern Web apps.

The work has been inspired by the gRPC and Twirp projects.


## Features

  * From a simple app.webrpc.json schema file (see examples/golang-basics/example.webrpc.json), for
  your target language of choice (Go and TypeScript currently supported), code generate:
    * Server message types, with encoding/decoding
    * Server HTTP handler, interfaces and routing
    * Complete client library in Go or TypeScript (extensible to support other language targets too)
  * Extensible -- code-generate for your own source and target languages
  * Pluggable -- write plugins to modify the code-generation for your target language, see "meta" field in schema
  * Extensible transports -- currently support HTTP and HTTP2, (WebSocket support is coming)


## Building from source

1. Install Go 1.11+
2. $ `go get -u github.com/webrpc/webrpc`
3. $ `make tools`
4. $ `make build`
5. $ `make test`
6. $ `go install ./cmd/webrpc-gen`


## Schema

The schema for webrpc we've designed is inspired by TypeScript and Go, and is simple enough
to cover the wide variety of language targets, designed to target RPC communication with
Web applications.

High-level features:

  * integers, floats, byte, bool, any, null, date/time
  * lists (multi-dimensional arrays supported too)
  * maps (with nesting / complex structures)
  * structs / objects
    * optional fields, default values, and pluggable code-generation for a language target
  * enums

For more information please see the [schema readme].


## License

MIT
