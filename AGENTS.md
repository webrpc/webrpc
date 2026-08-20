# webrpc — agent guide

webrpc is a schema-driven codegen tool: an API schema in RIDL (or JSON) generates
strongly-typed servers and clients (Go, TypeScript, JavaScript, Kotlin, Dart, OpenAPI)
that speak HTTP + JSON.

This file is for agents contributing to **this repo** (the schema parser and codegen
CLI). For helping a user *write RIDL schemas*, see `SKILL.md` instead.

## Layout

- `cmd/webrpc-gen/` — codegen CLI; `cmd/webrpc-test/` — interoperability test tool
- `schema/` — schema model and parsers; `schema/ridl/` is the RIDL parser
- `gen/` — template-driven generator core (code generators are Go templates)
- `_examples/`, `tests/` — runnable examples and interop schema; generated output is checked in
- Codegen templates live in **separate repos** (`webrpc/gen-golang`, `gen-typescript`, ...).
  Template bugs are fixed there, not here.

## Commands

- `make build` — build `./bin/webrpc-gen`
- `make test` — full suite (Go tests + interoperability suite)
- `make generate` — regenerate all examples and tests using latest templates
- `make diff` — fail if the working tree has changes (CI runs this after generate)
- `make update-ridl-test-golden-examples` — refresh RIDL parser golden JSON files
- `make dep-upgrade-templates` — bump all `gen-*` template dependencies

## Rules

- Never hand-edit generated files (`*.gen.*`, generated output under `_examples/` and
  `tests/`). Change the schema or the upstream `gen-*` template repo, then `make generate`.
- After changing parser or schema code: `make generate && make diff` to confirm generated
  output is deterministic and intentional.
- Golden files under `schema/ridl/_example/` are updated only via
  `make update-ridl-test-golden-examples` — never by hand.
- `make generate` deliberately strips version strings from generated files to avoid
  churn; don't "fix" that.
- The `message` keyword is deprecated (renamed to `struct` in v0.9.0) — keep the parser
  error message intact.
- RIDL files are formatted with [ridlfmt](https://github.com/webrpc/ridlfmt).

## Testing a parser/codegen change end-to-end

1. `make build`
2. `./bin/webrpc-gen -schema=_examples/golang-basics/example.ridl -target=golang -pkg=main -server -client -out=/tmp/out.gen.go`
3. `make test`
