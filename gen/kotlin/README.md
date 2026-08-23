# gen-kotlin

This repo contains the templates used by the `webrpc-gen` cli to code-generate
webrpc Kotlin client code.

This generator, from a webrpc schema/design file will code-generate:

1. Client -- a Kotlin client (via an injected transport, with optional provided OkHttp transport support and
`kotlinx.serialization`) to speak to a webrpc server using the
provided schema. This client is compatible with any webrpc server language (ie. Go, nodejs, etc.).

## Dependencies

Generated clients depend on `kotlinx.coroutines` and
`kotlinx.serialization`. Add these to the project that consumes the generated
file.

Add this to `build.gradle.kts`:
```kotlin
plugins {
    kotlin("plugin.serialization") version "<your-kotlin-version>"
}

dependencies {
    val coroutinesVersion = "1.7.3"
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:$coroutinesVersion")

    val serializationVersion = "1.6.3"
    implementation("org.jetbrains.kotlinx:kotlinx-serialization-json:$serializationVersion")
}
```

Generated clients depend on a `WebRpcTransport` abstraction.

If you want the generated file to also include the provided
`OkHttpWebRpcTransport`, generate with:

```sh
webrpc-gen -schema=example.ridl -target=kotlin -client -okhttpTransport=true -out=./example.gen.kt
```

and add:

```kotlin
dependencies {
    val okhttpVersion = "4.12.0"
    implementation("com.squareup.okhttp3:okhttp:$okhttpVersion")
}
```

The sample client project in `_examples/ClientExample/client` and the test
harness in `Tests` use Kotlin `1.8.21` on JVM `11`.

Generated output also exposes low-level method helpers for custom flows:

- schema-aware service symbols, for example `WaasWalletApi` / `WaasWalletClient`
- `SchemaServiceApi.basePath`
- `SchemaServiceApi.Method.path`
- `SchemaServiceApi.Method.urlPath`
- `SchemaServiceApi.Method.encodeRequest(...)`
- `SchemaServiceApi.Method.decodeResponse(...)`

## Usage

The built-in target is the simplest option:

```
webrpc-gen -schema=example.ridl -target=kotlin -client -out=./example.gen.kt
```

If you want to load this generator from git instead of the built-in target, use
an explicit git ref:

```
webrpc-gen -schema=example.ridl -target=github.com/webrpc/gen-kotlin@tags/v0.3.1 -client -out=./example.gen.kt
```

You can also point `-target` at a local folder:

```
webrpc-gen -schema=example.ridl -target=./local-templates-on-disk -client -out=./example.gen.kt
```

As you can see, `-target` supports the built-in `kotlin` target, a git URI, or
a local folder.

## Tooling

This repo pins the published webrpc tool module in `tools/go.mod` using Go tool
dependencies. Using the pinned tools from this repo requires Go `1.24+`.

Pinned tool versions:

- `github.com/webrpc/webrpc v0.37.2`
- `tool github.com/webrpc/webrpc/cmd/webrpc-gen`
- `tool github.com/webrpc/webrpc/cmd/webrpc-test`

The generator itself currently supports WebRPC schema version `v1`.

Use the pinned tools from this repo with:

```sh
go -C tools tool webrpc-gen
go -C tools tool webrpc-test
```

To update both to the latest published version and write that version back into `tools/go.mod`:

```sh
go -C tools get -tool github.com/webrpc/webrpc/cmd/webrpc-gen@latest github.com/webrpc/webrpc/cmd/webrpc-test@latest
```

## Development

Repository checks I verified while updating this README:

```sh
go test ./...
make -C _examples/ClientExample generate
```

### Target options

Change any of the following values by passing the target option to
`webrpc-gen`.

| webrpc-gen -option              | Description                | Default value              |
|---------------------------------|----------------------------|----------------------------|
| `-client`                       | generate client code       | unset (`false`)            |
| `-okhttpTransport=%bool%`       | include optional OkHttp transport | `false`            |
| `-packageName=%package name%`   | define package name        | `io.webrpc.client`         |

## LICENSE

[MIT LICENSE](./LICENSE.md)
