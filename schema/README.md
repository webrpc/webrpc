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
  - [List (Array)](#list-array)
  - [Map](#map)
  - [Enum](#enum)
  - [Struct (Message)](#struct-message)

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

- `timestamp` - for date/time


## List (Array)

- form: `[]<type>`
- ie.
  * `[]string`
  * `[]uint8`
  * `[][]string`
  * ..


## Map

- form: `map<key,value>`
- ie.
  * `map<string,any>`
  * `map<string,map<string,any>>`
  * `map<string,[]uint8>`
  * `map<int64,[]string>`
  * `map<string,User>` - where `User` is a struct type defined in schema


## Enum

- enum, see examples


## Struct (Message)

- think of a struct as a JavaScript or JSON object
- struct has 0..N fields
  - field can be `optional`
  - fields are by default required, unless made optional
  - fields always return default values by default, ie. default of int is 0, string is "", etc. (like in Go)
    - otherwise someone should make it optional which will have it be nullable
