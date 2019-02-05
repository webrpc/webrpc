TODO
====

## Schema TODO

- [x] general file parsing + structs
- [x] parsing DataType for lists, maps, etc.
- [ ] dont allow dupe message names
- [ ] service arguments
- [ ] 


## Overall


- [x] example stub, hand-made
- [~] example json-rpc for request-response method call
- [ ] JSON-RPC? or dropping the meta-structure around it?

- [ ] schema - get inspiration from TS, https://www.typescriptlang.org/docs/handbook/basic-types.html

- [ ] code-generation of server (Go)
- [ ] code-generation of client (Go)
- [ ] code-generation of client (TypeScript)

- [ ] example json-rpc for streaming from server->client over websockets (and later with http2)
- [ ] example json-rpc for streaming from client->server over websockets (and later with http2)
- [ ] example json-rpc for streaming from server<->client over websockets (and later with http2)
 

- [ ] binary data encoding/decoding, []byte inside of json-rpc ..?
  * see https://stackoverflow.com/questions/53354042/binary-data-inside-json
  * Buffer .toJSON() ..
  * https://github.com/PsychoLlama/bin-json

- [ ] pluggable transport - easily add TCP, QUIC, ..
- [ ] besides JSON, should we support pluggable encoding formats? flatbuffers? msgpack? cbor? other?


- [ ] future: custom IDL for webrpc that resembles TypeScript and generates a webrpc-schema.json file

interface GetUserRequest {
  id?: number {
    structTags: 'db:"id"'
  } // this is an annotation, got some other ideas?
  // + how do we add plugin-specific hints in here..?

  userID?: number
}

.. but, adding some custom stuff by target as well
