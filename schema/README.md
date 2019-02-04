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


## Objects

- object
  * ie. a struct.. etc....... TODO..


## Enums

- enum of strings
- enum of integers


## Binary (future / v2)

- blob
  * TODO: https://github.com/PsychoLlama/bin-json might have some ideas for us
