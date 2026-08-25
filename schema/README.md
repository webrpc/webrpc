webrpc schema <!-- omit in toc -->
=============

webrpc is a design/schema driven approach to writing backend servers, with fully-generated
client libraries. Write your schema, and it will generate strongly-typed bindings between
your server and client. The type system is described below.

Some example webrpc schemas:
  * from the _examples/, here is a schema in [RIDL](../_examples/golang-basics/example.ridl) or
  in [JSON](../_examples/golang-basics/example.webrpc.json)
  * ..find more in ./_examples

- [Type system](#type-system)
  - [Core types](#core-types)
    - [Integers](#integers)
    - [Floats](#floats)
    - [Strings](#strings)
    - [Timestamps (date/time)](#timestamps-datetime)
    - [Files](#files)
  - [List](#list)
  - [Map](#map)
  - [Enum](#enum)
  - [Struct](#struct)

# Type system

## Core types

- `byte` (aka uint8)
- `bool`
- `any`
- `null`


### Integers

- `uint8`
- `uint16`
- `uint32`
- `uint64`

- `int8`
- `int16`
- `int32`
- `int64`


### Floats

- `float32`
- `float64`


### Strings

- `string`


### Timestamps (date/time)

- `timestamp` - for date/time (serialized to [ECMA Script ISO 8601 format](https://tc39.es/ecma262/multipage/numbers-and-dates.html#sec-date-time-string-format): `YYYY-MM-DDTHH:mm:ss.sssZ`)


### Files

- `file` - for upload/download endpoints (transferred as raw bytes over the wire, not JSON)
- `file` and `[]file` may be used only as:
  * the type of a method input argument, or
  * the type of a top-level field of a struct used as a method input
- a method may return at most one `file` output (a download method)
- a download method may declare metadata outputs alongside its file:

  ```ridl
  - DownloadAvatar(userId: uint64) => (avatar: file, size: uint64, etag: string)
  ```

  The response body carries the raw file, so the remaining outputs are
  JSON-encoded into the `Webrpc-Response` header instead. Keep them small:
  header limits are enforced by servers, proxies and CDNs, and generated
  servers reject metadata over 4 KiB. Non-ASCII characters are escaped, since
  HTTP header values are ASCII.
- `file` is not allowed in nested structs, maps, aliases, or streaming methods

Metadata outputs are supported by the golang, typescript and javascript
targets. The kotlin, dart and swift clients support plain `file` downloads and
fail at generation time on a download method that declares metadata.


## List

- List represents a JSON array over the wire
- form: `[]<type>`
- ie.
  * `[]string`
  * `[]uint8`
  * `[][]string`
  * ..


## Map

- Map represents a JSON object with 0..N properties (key:value pairs) over the wire
- form: `map<key,value>`
- ie.
  * `map<string,any>`
  * `map<string,map<string,any>>`
  * `map<string,[]uint8>`
  * `map<int64,[]string>`
  * `map<string,User>` - where `User` is a struct type defined in schema


## Enum

- enum, see examples


## Struct

- struct represents a JSON object over the wire
- struct has 0..N fields
  - field can be `optional`
  - fields are by default required, unless made optional
  - fields always return default values by default, ie. default of int is 0, string is "", etc. (like in Go)
    - otherwise someone should make it optional which will have it be nullable
