---
name: webrpc-ridl-schema
description: Write and edit webrpc .ridl schema files and run webrpc-gen code generation. Use when working with RIDL syntax, webrpc schemas, generated *.gen.go / *.gen.ts files, or designing webrpc API contracts (Go/TypeScript/JavaScript/Kotlin/Dart/OpenAPI).
---

# webrpc RIDL schemas and codegen

RIDL defines an HTTP API contract; `webrpc-gen` generates typed servers and clients
from it. POST-only with JSON bodies — never design path or query params.

## Rules
- Edit `.ridl` sources, never `*.gen.*` output — rerun codegen instead.
- Identifiers are case-sensitive. Run `ridlfmt` after edits. No circular imports.
- Prefer succinct method signatures (one request struct in, one response struct out);
  never mix succinct and multi-arg forms in one method.
- Common mistakes: missing `webrpc = v1`; deprecated `message` keyword (use `struct`);
  forgetting `?` on nullable fields; omitting `HTTP <status>` on non-400 errors.

## Header (required) and imports
```
webrpc = v1
name = <schema-name>
version = <schema-version>
basepath = /rpc

import "path/to/file.ridl"    # everything (quotes optional)
import "../shared/ids.ridl"   # named members only
  - OrgID
```
Imports merge into one flat schema. An import with no referenced member can be
intentional (pulls types into the generated package) — keep those in codegen roots only.

## Running codegen
Install from releases, or pin as a Go tool dependency and run via
`go run github.com/webrpc/webrpc/cmd/webrpc-gen`:
```
webrpc-gen -schema=main.ridl -target=golang -pkg=proto -server -client -out=./proto/api.gen.go
webrpc-gen -schema=app/dashboard.ridl -target=typescript -client -enumStyle=union -out=./src/rpc/proto/api.gen.ts
```
Golang flags: `-errorStackTrace=true`, `-json=sonic`. Pin targets: `-target=golang@v0.14.0`.

## Multi-file layout (recommended for larger projects)
```
schema/
  main.ridl          # Go codegen root: imports errors.ridl + every app/*.ridl
  errors.ridl        # all errors in one place
  app/<app>.ridl     # per-app TS codegen roots
  service/<app>/     # one service block per domain
  model/             # shared structs, typed ids, enums
```
Go codegen runs on `main.ridl` (one unified package); TS codegen on `app/<app>.ridl`
so each webapp sees only its own surface.

## Types
`#` starts a comment; adjacent comments attach to the next definition as docs.
Core: `byte`, `bool`, `any`, `null`, `string`, `timestamp`, `uint8`-`64`, `int8`-`64`,
`float32/64`, `[]T`, `map<K,V>`.
```
struct User
  - id: uint64
  - name?: string    # optional; fields are required unless marked ?

enum SortOrder: uint32   # values default by index
  - DESC
  - ASC

type WebhookID: string          # alias: JSON stays string;
  + go.alias = typeid.UUID      # Go uses this type (typed IDs, uuid.UUID, ...)
  + go.type.import = github.com/example/typeid
```

## Services, errors, annotations
```
error 1000 Unauthorized "unauthorized" HTTP 401    # status defaults to 400
error 4001 InvalidTransition "invalid state" HTTP 409

service Example
  @deprecated:"use NewMethod instead"
  @access:s2s
  - GetUser(GetUserRequest) => (GetUserResponse)   # succinct form, preferred
  - Ping() => (version: string)                    # no-payload inline return is fine
```
- Group error codes by range (1xxx auth, 2xxx validation, 3xxx not-found, 4xxx business).
- Go: a `WebRPCError` must reach the transport unwrapped (the server type-asserts it) —
  use `proto.ErrXxx.WithCause(err)`, not `fmt.Errorf` chains.
- Custom annotations (`@access:`, `@perms:`) are readable at runtime: middleware calls
  `proto.MethodCtx(ctx)` then `Annotation("access")` / `HasAnnotation(...)`.

## Field metadata
```
  - id: string
    + go.field.name = ID       # rename the Go field
    + go.field.type = uint64   # override the Go type
    + go.tag.db = id           # any go.tag.<name> becomes a struct tag
  - featureIndex: int
    + json = -                 # server-only: dropped from Go json AND the TS interface
```
- `+ json = -` vs `+ go.tag.json = -`: the latter is Go-only — the field stays in the TS
  interface but is always `undefined` at runtime (the TS type lies). Use `go.tag.json`
  only for Go tag options like `,omitempty`.
- The generator title-cases names naively (`userId` → `UserId`); fix acronyms with
  `go.field.name` (`ID`, `URL`, `URI`, `API`, `JSON`, `HTTP`, `RPC`, `IP`, `DB`, ...).

### Inferred Go types — override only to get a *different* type
| RIDL | Go |
|---|---|
| primitives, `[]T`, `map<K,V>` | direct equivalents; `timestamp` → `time.Time` |
| `name?: <simple>` | pointer (`*string`); `name?: []T` stays `[]T` (already nil-able) |
| struct fields | pointers: `*S`, `[]*S`, `map[K]*S` (succinct method arg: value) |
| enum / alias | value; pointer when optional |

Useful overrides: a different type (`json.RawMessage`, typed IDs), value-element slices
(`[]S`), non-pointer optionals. Overrides matching the table are noise — delete them.

## Custom templates
`-target` accepts a local Go text/template directory, and fields carry arbitrary
metadata — combine them to generate validators, permission constants, etc.:
```
  - email: string
    + validate = "required,maxlen=320,email"
```
```
webrpc-gen -schema=main.ridl -target=./schema/templates/validate-go -pkg=proto -out=./proto/validate.gen.go
```

## References
`_examples/golang-basics/example.ridl`, `schema/README.md`, `schema/ridl/README.md`.
Official targets: `github.com/webrpc/gen-{golang,typescript,javascript,kotlin,dart,openapi}`.
