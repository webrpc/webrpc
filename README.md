webrpc
======

webrpc, is simple/lightweight RPC framework for webapps via JSON-RPC.

think of it as a much simpler gRPC, but using JSON for serialization, and the JSON-RPC 2.0
form for making remote server calls. See https://www.jsonrpc.org/specification

## Features

  * From a simple proto.json schema file (see _example/proto.json), for your language
    of choice (ie. Go), code generate:
    * Server message types, with encoding/decoding
    * Server HTTP handler, interfaces and routing
    * Complete client library in Go and TypeScript (or other languages)
  * Extensible, pluggable -- code-generate for your own source and target languages
  * Extensible transports, supporting first HTTP, HTTP2 and WebSockets


## Schema

  * primitives
  * objects
  * lists
  * optional fields, etc.
  * its just JSON! https://www.json.org


## License

MIT
