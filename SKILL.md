---
name: webrpc-ridl-schema
description: webrpc schema and RIDL syntax
---

# webrpc schema and RIDL syntax skill for LLMs

## When to use this skill
Use this skill when the user needs to work with RIDL files.

The RIDL file defines schema for HTTP client/server communications (browser-to-server or service-to-service).
The `webrpc-gen` codegen generates a REST-like API with JSON messages, a subset of REST API conventions, which:
- is not RESTful
- always uses POST method
- uses JSON body only (no path params or query params)

## Do and don't
- Prefer editing `.ridl` files; do not manually edit generated `*.gen.*` files.
- Keep syntax strict: identifiers are case‑sensitive; spacing is mostly flexible.
- Use `ridlfmt` after edits when possible.
- Prefer succinct method signatures for request/response structs.
- Avoid mixing succinct and multi‑arg method signatures.
- Don't use circular imports.

## Schema header (required)
```
webrpc = v1
name = <schema-name>
version = <schema-version>
basepath = <api-base-path>
```

## Imports
```
import "path/to/file.ridl"
import "path/to/file.ridl" (TypeA, ServiceB)
```
Imports are merged into the current schema and can be limited to named members.

## Comments
- `#` starts a line comment.
- Adjacent comment lines attach to the next definition as doc comments.

## Types
Core types:
`byte`, `bool`, `any`, `null`, `string`, `timestamp`,
`uint8/16/32/64`, `int8/16/32/64`, `float32/64`.

List:
`[]Type`, `[][]Type`

Map:
`map<key,value>`

Struct (message):
```
struct User
  - id: uint64
  - name?: string  # optional field
```
Fields are required unless marked optional (`?`).

Enum:
```
enum SortOrder: uint32
  - DESC = 0
  - ASC  = 1
```
If no explicit value is provided, enum values default by index.

## Services and methods
```
service Example
  - Ping()
  - GetUser(userId: uint64) => (user: User)
```

Succinct form (single struct arg/return only, preferred):
```
service Example
  - Ping(PingRequest) => (PingResponse)
  - GetUser(GetUserRequest) => (GetUserResponse)
```
Do not mix succinct form with multi‑arg form in a single method.

## Errors
```
error 100 RateLimited "too many requests" HTTP 429
error 200 UserNotFound "user not found"
```
Default HTTP status is 400 when omitted.

## Annotations
Methods can be annotated:
```
@deprecated:"use NewMethod instead"
@stampede:3s
```

## Field metadata
Fields can include metadata lines:
```
  - id: string
    + go.field.name = ID
    + go.field.type = uint64
    + go.tag.json = id,string
    + go.tag.db = id,omitempty
  - createdAt: timestamp
    + go.tag.db = created_at,omitempty
  - updatedAt: timestamp
    + go.tag.db = updated_at,omitempty
  - deletedAt?: timestamp
    + go.tag.db = deleted_at,omitempty
  - featureIndex: int
    + json = - # internal server-only field (omitted in clients)
```

## Good references in this repo
- `_examples/golang-basics/example.ridl` (small, clear schema)
- `schema/README.md` (type system)
- `schema/ridl/README.md` (errors and RIDL notes)

## webrpc-gen codegen targets (upstream Go template repositories)
- https://github.com/webrpc/gen-golang
- https://github.com/webrpc/gen-typescript
- https://github.com/webrpc/gen-javascript
- https://github.com/webrpc/gen-kotlin
- https://github.com/webrpc/gen-dart
- https://github.com/webrpc/gen-openapi

## Common mistakes
- Missing `webrpc = v1`.
- Using deprecated `message` keyword instead of `struct`.
- Mixing succinct and multi‑arg method signatures.
- Forgetting optional `?` on nullable fields.
