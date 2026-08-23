# gen-swift

This repo contains the templates used by the `webrpc-gen` CLI to code-generate
webrpc Swift client code.

## Scope

`gen-swift` is client-only.

- It generates Swift client code and runtime helpers.
- It does not generate server handlers or server runtime code.
- Passing `-server` is rejected as an unsupported option.

## Usage

`gen-swift` is a built-in `webrpc-gen` target:

```sh
webrpc-gen -schema=example.ridl -target=swift -client -out=./ExampleClient.swift
```

Or point `-target` at a local checkout of this repo:

```sh
webrpc-gen -schema=example.ridl -target=./local-templates-on-disk -client -out=./ExampleClient.swift
```

## Generated Surface

Generated output includes:

- schema constants
- WebRPC header/version metadata helpers
- DTOs and enums
- transport/error helpers
- service metadata helpers
- high-level async client methods

The generated client targets Swift 5.9+. The built-in `URLSession` transport is intended
for modern Apple concurrency platforms such as iOS 15+ and macOS 12+.

Low-level helpers remain visible:

- `ServiceAPI.basePath`
- `ServiceAPI.Method.path`
- `ServiceAPI.Method.urlPath`
- `ServiceAPI.Method.encodeRequest(...)`
- `ServiceAPI.Method.decodeResponse(...)`

## Tooling

This repo pins the published webrpc tool module in `tools/go.mod` using Go tool
dependencies.

Use the pinned tools with:

```sh
go -C tools tool webrpc-gen
go -C tools tool webrpc-test
```

## Testing

Use the fast generator/unit suite for normal development:

```sh
go test ./...
```

Run the shared Swift integration package separately when you want real generated
Swift compilation and runtime coverage:

```sh
./Tests/test.sh
```

That script clears stale SwiftPM build artifacts, regenerates
`Tests/Sources/Generated/Generated.swift` from `Tests/test.ridl`, and runs `swift test`
against a shared SwiftPM package.

The real external-schema integration remains available as an opt-in Go test:

```sh
WEBRPC_SWIFT_EXTERNAL_SCHEMA=1 go test ./...
```

Regenerate and build the tracked client example with:

```sh
make -C _examples/ClientExample generate
make -C _examples/ClientExample build
make -C _examples diff
```

CI runs the fast Go suite on every PR and push to `master`, and runs the shared Swift
integration suite and tracked example checks on `macos-latest`.

## Options

### Embedding A Generated Client

Generated clients are public top-level declarations by default, which is appropriate
when applications use the generated client directly. Libraries that embed a generated
client as an implementation detail can use `-namespace` to avoid top-level name
collisions and `-visibility=internal` to keep generated models, transport types, and
helpers out of the host library's public Swift API.

These options are opt-in. Existing generator output remains public and un-namespaced
unless explicitly configured otherwise.

Namespace values must use ASCII identifier syntax: they start with a letter or `_`
and contain only letters, digits, or `_`.

| webrpc-gen option | Description | Default |
| --- | --- | --- |
| `-client` | generate client code | unset (`false`) |
| `-webrpcHeader` | send the standard `Webrpc` header on client requests | `true` |
| `-schemaHash=false` | omit the schema hash + version constants from generated output | `true` |
| `-visibility=public\|internal` | set generated declaration access | `public` |
| `-namespace=<SwiftIdentifier>` | wrap declarations in a namespace enum | none |
