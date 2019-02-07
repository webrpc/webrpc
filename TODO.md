TODO
====

## Schema TODO

- [x] general file parsing + structs
- [x] parsing DataType for lists, maps, etc.
- [x] dont allow dupe message names
- [x] service parsing + verification
- [x] VarName casing methods
- [x] VarName parsing + verification, check against regexp a-z0-9_
- [x] service arguments
- [x] timestamp type
- [x] Enum design.. review
- [x] enum parsing + verification
- [ ] rpc methods, do we need argument names for inputs?
- [ ] rpc methods, do we need argument names for outputs?
- [ ] improve Go client error response code, some TODO's in there..


## Go gen

- [ ] should message types be objects? .. ie "User" in schema -> *User in code..
- [x] can our current gen pkg generate multiple services?


## Next

- [ ] streaming request payload -- we will need a sequence number here
- [ ] streaming response payload -- we will need a sequence number and status code here

- [ ] request body via query params option (more cachable)
    * is this needed..? perhaps CloudFlare workers can help us instead?


## Overall

- [x] example stub, hand-made
- [x] schema - get inspiration from TS, https://www.typescriptlang.org/docs/handbook/basic-types.html

- [x] code-generation of server (Go)
- [x] code-generation of client (Go)
- [ ] code-generation of client (TypeScript)

- [ ] example streaming from server->client over websockets (and later with http2)
- [ ] example streaming from client->server over websockets (and later with http2)
- [ ] example streaming from server<->client over websockets (and later with http2)
 

## Folder structure (ideas..)

### Option A monorepo: github.com/webrpc/webrpc -- what we have right now..

* /cmd/webrpc-gen       - cli util for all generation
* /schema               - webrpc schema parser+validator
* /gen/golang           - code-gen templates for Go client+server via schema ast
* /gen/typescript       - code-gen templates for TS client+server
* /gen/<targetlang>     - ...
* /pkgs/webrpc-go       - Lib for Go client+server
* /pkgs/webrpc-js       - Lib for JS client (Web/node) + server (node)
* /pkgs/webrpc-ts       - Lib for TS client (Web/node) + server (node)
* /pkgs/webrpc-rust     - ..
* /pkgs/<targetlang>

NOTE: we may only need a runtime for the server parts..
and clients could be self-contained from code generated source (maybe).. websockets? http2? other encodings? .. would prob be seperarate dependencies..


### Option B multiple repos

* github.com/webrpc/webrpc
  * /cmd/webrpc-gen       - cli util for all generation
  * /schema               - webrpc schema parser+validator
  * /gen/golang           - code-gen templates for Go client+server via schema ast
  * /gen/typescript       - ...
  * /gen/rust
  * /gen/<targetlang>

* github.com/webrpc/webrpc-go
  * the runtime for webrpc for Go client or servers

* github.com/webrpc/webrpc-ts-server (for nodejs only)
  * the runtime for webrpc node server

* github.com/webrpc/webrpc-ts-client (for Web/node)
  * NOTE: I don't think we'll need this, the entire client lib can fit in the
    code-generated source I believe..

* github.com/webrpc/webrpc-rust
  * the runtime for webrpc for Rust client or servers


