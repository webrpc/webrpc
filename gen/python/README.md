# webrpc Python generator

The templates used by `webrpc-gen` to code-generate webrpc Python clients and
servers.

It emits stdlib `@dataclass` models (no runtime dependency for the types), an
[`httpx`](https://www.python-httpx.org/) client, and a framework-agnostic WSGI
server.

## Usage

```sh
webrpc-gen -schema=example.ridl -target=python -client -server -out=./example_api.py
```

## Options

| webrpc-gen option | Description | Default |
| --- | --- | --- |
| `-client` | generate the `httpx` client | unset |
| `-server` | generate the WSGI server and the service `Protocol` | unset |
| `-schemaHash=false` | omit the schema version and hash constants | `true` |
| `-webrpcHeader=false` | don't send the `Webrpc` version header on requests | `true` |

## Output

- **Types**: a `@dataclass` with generated `from_dict` / `to_dict`. Enums become
  `enum.Enum`; aliases become type assignments.
- **Errors**: a `WebRPCError` subclass per webrpc and schema error, and a code
  registry to rebuild them on the client.
- **Client**: one `<Service>Client`. Methods are `snake_case`, return typed
  values and raise the matching error on a non-2xx response.
- **Server**: a `<Service>` `typing.Protocol` to implement, and a
  `<Service>Server` WSGI app that routes requests, deserializes inputs and
  serializes outputs.

Requires Python 3.11+. The client requires `httpx`.

## Known gaps

- Streaming methods, the `file` core type and `[]byte` are not supported: the
  generator fails when the schema uses them.
- REST routes (`+ GET /path`) are ignored; every method is served and called
  as `POST <basepath><Service>/<Method>`.
- Enums are serialized as their wire value, like the Go and TypeScript
  generators: the value of a string enum, the member name (or its `json`
  meta) of a numeric one.
- Timestamps keep microsecond precision: the nanoseconds of a Go timestamp are
  truncated on decode.
- Request validation differs from the Go server in a few places. An unknown
  enum value and a missing or null non-optional timestamp are a
  `400 WebrpcBadRequest`, where Go decodes the zero value. Integer width and
  sign are not range-checked (a `uint8` of 300 is accepted). On Python 3.11 and
  later, a date-only timestamp is accepted.
- Alias-typed struct fields are deserialized through the alias name, for
  scalar aliases and aliases of user types alike.

## Tests

`tests/interop-python` runs the generated code against the Go reference
implementation (`webrpc-test`): a generated Python client against the Go
server, and the Go client against a generated Python WSGI server. The
interop schema is `tests/schema/test.ridl` without its `file` methods and
REST service.
