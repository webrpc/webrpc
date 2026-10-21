---
name: webrpc-ridl-schema
description: Write and edit webrpc .ridl schema files and run webrpc-gen code generation. Use when working with RIDL syntax, webrpc schemas, generated *.gen.go / *.gen.ts files, or designing webrpc API contracts (Go/TypeScript/JavaScript/Kotlin/Dart/C#/OpenAPI).
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
  - GetUser(GetUserRequest) => (GetUserResponse)   # succinct form, preferred
    + @deprecated = "use NewMethod instead"
    + @access = s2s
  - Ping() => (version: string)                    # no-payload inline return is fine
```
- Group error codes by range (1xxx auth, 2xxx validation, 3xxx not-found, 4xxx business).
- A method block holds `+ @name[ = value]` annotations, `+ VERB /path` routes,
  `+ success <codes>` and `+ error <Name>` option lines; all of them interleave
  freely below the `- Method` line. Anything else in the block is a parse error.
  One annotation per line: the value after `=` follows struct-tag rules — a bare value is a
  single word (backslashes stay literal), so quote anything with spaces, a literal `#`, or escapes. Above-method
  annotations are deprecated, warn, and will be removed in a future release.
- A method lists the errors it can throw with the `errors A|B` clause on the `- Method`
  line, or with one or more `+ error A|B` option lines below it. Both forms merge into one
  list. Names are case-sensitive and must match a schema `error` declaration; a repeated
  name is a parse error.
- Go: a `WebRPCError` must reach the transport unwrapped (the server type-asserts it) —
  use `proto.ErrXxx.WithCause(err)`, not `fmt.Errorf` chains.
- Custom annotations (`@access`, `@perms`) are readable at runtime: middleware calls
  `proto.MethodCtx(ctx)` then `Annotation("access")` / `HasAnnotation(...)`.

## REST routes
```
basepath = /api

service Users
  - GetUser(GetUserRequest) => (GetUserResponse)
      + GET /users/{id}
      + POST /Users/GetUser      # extra route: server-only alias, keeps the old RPC path
```
- `+ VERB /path` goes under the method. Paths are absolute; the full route is `{basepath}{path}`.
- A method can have several routes. The first is canonical: the generated client calls only it.
- Any route replaces the default RPC path (`POST {basepath}{Service}/{Method}`).
  List it as a route to keep it.
- `{id}` binds to a required scalar request field. `GET`/`DELETE` send other fields as query
  params; `POST`/`PUT`/`PATCH`/`QUERY` send them in the JSON body.
- An empty body is a zero-value request and needs no `Content-Type` — empty means zero
  bytes: a whitespace-only body is a payload and still requires `application/json`, then
  decodes as the zero-value request. A route with zero body fields drops the check
  entirely: any Content-Type is accepted, and a non-empty body is parsed as JSON (unknown
  fields ignored; malformed JSON is a 400). A route *with* body fields still requires
  `application/json` for non-empty bodies: a cross-site form POST (`text/plain`) triggers
  no CORS preflight, so this check is the CSRF guard. Note the classic RPC dispatcher
  requires `application/json` on every request; the relaxed REST rules assume
  header/token auth — cookie-auth deployments need their own CSRF defense.
- `+ success <codes>` declares the method's allowed success statuses as a
  pipe-separated set, e.g. `+ success 200|201`. It applies to every route of the
  method, wherever the `+ success` line sits in the method block. One per method;
  it requires at least one route. The first is the default; no `+ success` line
  means 200. Codes must be 2xx and unique, and 204/205 are allowed
  only for methods with an empty response (no outputs or an empty struct); such routes
  never carry a response body or `Content-Type`. The WriteHeader responsibility is fixed
  at codegen time: on a single-code route (or no `+ success` line) the framework writes
  the declared status and the handler must not; on a multi-code route the Go handler
  picks one with `ResponseWriterFromContext(ctx).WriteHeader(code)` — not validated, the
  declared set is the contract. The TypeScript server picks it the same way via the
  generated `withStatus(status, result)` helper (an untagged result uses the declared
  default).
- Unit-testing a Go handler that writes a status needs a ResponseWriter in ctx: inject
  one with `context.WithValue(ctx, HTTPResponseWriterCtxKey, httptest.NewRecorder())`.
- Routed endpoints ignore unknown fields in the JSON body. A strict-decode opt-in is
  proposed in [#530](https://github.com/webrpc/webrpc/issues/530).
- Only the golang and typescript targets support routes. Other targets exit with an error.

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
Official targets: `gen/{golang,typescript,javascript,kotlin,dart,csharp,openapi}` (embedded in `webrpc-gen`).
