TODO
====

## Schema TODO

- [x] general file parsing + structs
- [x] parsing DataType for lists, maps, etc.
- [x] dont allow dupe message names
- [x] service parsing + verification
- [x] VarName casing methods
- [ ] service arguments
- [ ] VarName parsing + verification, check against regexp a-z0-9_
- [ ] Enum design.. review
- [ ] enum parsing + verification

## Next

- [ ] json-rpc, to be or not to be?
- [ ] request body via query params option (more cachable)


## Overall

- [x] example stub, hand-made
- [~] example json-rpc for request-response method call
- [ ] JSON-RPC? or dropping the meta-structure around it?

- [x] schema - get inspiration from TS, https://www.typescriptlang.org/docs/handbook/basic-types.html

- [ ] code-generation of server (Go)
- [ ] code-generation of client (Go)
- [ ] code-generation of client (TypeScript)

- [ ] example json-rpc for streaming from server->client over websockets (and later with http2)
- [ ] example json-rpc for streaming from client->server over websockets (and later with http2)
- [ ] example json-rpc for streaming from server<->client over websockets (and later with http2)
 