WebRPC Schema
=============

## Basics

- byte (aka uint8)
- bool
- any
- null


## Numbers

- integer (aka int64)

- uint8
- uint16
- uint32
- uint64

- int8
- int16
- int32
- int64


## Floats

- float (aka float64)

- float32
- float64


## Strings

- string


## Lists

- <type>[]

- string[]

- uint8[]


## Map

- map<key, value>
  * ie. map<string, any>

- map<string, map<string, any>>

- map<string, integer>


## Enums

- enum of strings
- enum of integers


## Binary (future / v2)

- blob
  * TODO: https://github.com/PsychoLlama/bin-json might have some ideas for us


## Objects aka Messages

- object
  * ie. a struct, a class, etc.
  * think of it just as a Javascript object or JSON object

### Some notes on objects / messages

- fields of an object can be `optional`
- fields of an object are by default required, unless made optional
- fields of an object always return default values by default, ie. default of int is 0, string is "", etc. (like in Go)
  - otherwise someone should make it optional which will have it be nullable


## Services aka RPC methods

- So.. gRPC's service/rpc method definitions like `rpc Ping(Empty) returns (Empty);` or any method, set
a requirement that a method must have one request argument and one response value. That is,
`rpc Ping() returns (Empty);` is not allowed, nor is `rpc Ping(Empty) returns ();`, or any combination
.. it must have one request input and one response output
- In webrpc, I think it's okay if we follow a similar requirement for the v0.1 of the implementation
but certainly id like to go beyond that and have no arguments, or  a list of arguments, etc.
lets save it for v0.2 though
- In webrpc, we might be okay to take an empty input list like `rpc Ping() returns (bool);` but
we probably always need a return type.
