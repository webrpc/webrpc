# webrpc C# generator

The templates used by `webrpc-gen` to code-generate webrpc C# clients for any
webrpc server (Go, TypeScript, ...).

The generator is client-only: `-server` is rejected.

## Usage

```sh
webrpc-gen -schema=example.ridl -target=csharp -client -namespace=Example.Api -out=./Example.gen.cs
```

Without `-client`, only the schema types are generated.

## Options

| webrpc-gen option | Description | Default |
| --- | --- | --- |
| `-client` | generate the service clients | unset (types only) |
| `-namespace=<Namespace>` | C# namespace of the generated code, eg. `Example.Api` | the schema name in PascalCase |
| `-json=stj\|newtonsoft` | JSON library: `System.Text.Json` or `Newtonsoft.Json` | `stj` |
| `-visibility=public\|internal` | accessibility of the generated declarations | `public` |
| `-webrpcHeader=false` | don't send the `Webrpc` version header on requests | `true` |
| `-schemaHash=false` | omit the schema version and hash constants | `true` |

Use `-visibility=internal` to embed a generated client in a library as an
implementation detail. Each generated file carries its own runtime helpers, so
generate each schema into its own namespace.

## Requirements

- C# 8 or later. The generated code targets `netstandard2.0` and `netstandard2.1`,
  so it runs on .NET (Core) 2.0+, .NET Framework 4.6.2+ and Unity.
- `-json=stj`: System.Text.Json 6.0 or later, built into .NET 6+, or the
  [System.Text.Json](https://www.nuget.org/packages/System.Text.Json) package.
- `-json=newtonsoft`: [Newtonsoft.Json](https://www.nuget.org/packages/Newtonsoft.Json) 13.
- Streaming methods return `IAsyncEnumerable<T>`, built into `netstandard2.1`
  and .NET Core 3.0+; add the
  [Microsoft.Bcl.AsyncInterfaces](https://www.nuget.org/packages/Microsoft.Bcl.AsyncInterfaces)
  package on `netstandard2.0` and .NET Framework.

## Generated code

```ridl
service Example
  - GetUser(userId: uint64) => (user: User)
  - FindUsers(FindUsersRequest) => (FindUsersResponse)
  - Upload(name: string, data: file) => (size: uint64)
  - Watch(userId: uint64) => stream (event: Event)
```

```csharp
using var http = new HttpClient();
var api = new ExampleClient("https://api.example.com", http);

User user = await api.GetUserAsync(1);
FindUsersResponse found = await api.FindUsersAsync(new FindUsersRequest { Query = "jane" });

ulong size = await api.UploadAsync("avatar.png", WebrpcFile.FromBytes(bytes, "avatar.png", "image/png"));

await foreach (Event ev in api.WatchAsync(1, cancellationToken))
{
    // ...
}
```

For every service, the generator emits an `I{Service}Client` interface and a
`{Service}Client` class with one `{Method}Async` method per schema method. Every
method takes a `CancellationToken`.

| Method | C# signature |
| --- | --- |
| no outputs | `Task MethodAsync(args...)` |
| one output | `Task<T> MethodAsync(args...)` |
| several outputs | `Task<{Service}Client.{Method}Response> MethodAsync(args...)` |
| succinct, `Method(Req) => (Resp)` | `Task<Resp> MethodAsync(Req request)` |
| file download, `=> (data: file)` | `Task<WebrpcFile> MethodAsync(args...)` |
| streaming, `=> stream (...)` | `IAsyncEnumerable<T> MethodAsync(args...)` |

- **REST routes**: methods with `+ GET /path/{param}` routes call their
  canonical (first) route. Path and query params are filled from the arguments,
  and the other fields go in the JSON body for POST/PUT/PATCH/QUERY routes.
- **Files**: upload methods send `multipart/form-data`, streaming each file's
  `Body`. A downloaded `WebrpcFile` streams the HTTP response, so dispose it.
- **Streams**: the request is sent when the enumeration starts. Each message
  is yielded as it arrives, and the enumeration ends when the server closes the
  stream. Cancel the token to stop listening.
- **Deprecated methods** (`+ @deprecated = "..."`) are marked `[Obsolete]`.

### Types

| webrpc | C# |
| --- | --- |
| `bool` | `bool` |
| `uint8`, `byte` | `byte` |
| `uint16`, `uint32`, `uint64`, `uint` | `ushort`, `uint`, `ulong`, `uint` |
| `int8`, `int16`, `int32`, `int64`, `int` | `sbyte`, `short`, `int`, `long`, `int` |
| `float32`, `float64` | `float`, `double` |
| `string` | `string` |
| `timestamp` | `DateTimeOffset` (ISO 8601 in UTC on the wire) |
| `bigint` | `System.Numerics.BigInteger` (decimal string on the wire) |
| `any`, `null` | `object?`: decoded as a `JsonElement` (System.Text.Json), or a `JToken` or primitive (Newtonsoft.Json) |
| `[]byte`, `[]uint8` | `byte[]` (base64 on the wire, like Go) |
| `[]T` | `List<T>` |
| `map<K,V>` | `Dictionary<K,V>` |
| `file` | `WebrpcFile` |
| struct | `partial class` |
| enum | `readonly partial struct` (see below) |
| type alias | the aliased type |

Optional fields and arguments are nullable (`int?`, `string?`), and are left
out of the JSON when null. Required strings, lists and maps are initialized
to empty values. Fields with `+ json = name` are renamed on the wire, and fields
with `+ json = -` are left out.

Identifiers that are C# keywords are escaped (`@class`), and members that would
clash with their type or with `System.Object` members get a `_` suffix (`ToString_`),
as do arguments named like the generated parameters (`cancellationToken_`).

### Enums

Enums are value types that keep values unknown to the client. A server can add
enum values without breaking older clients:

```csharp
Status status = item.Status;
if (status == Status.ACTIVE) { /* ... */ }
switch (status.Value) { case "ACTIVE": /* ... */ break; }
if (!status.IsKnown) { /* a value added to the server schema later */ }
```

`Status.KnownValues` lists the schema values. `default(Status)` is the zero
value: the `0` member of an integer enum, or `""` for a string enum.

### Errors

Errors are thrown as `WebrpcException`, with `Code`, `Name`, `Message`,
`HttpStatus` and `Cause`. Each webrpc and schema error has its own subclass,
eg. `catch (UserNotFoundException)`. Client failures are thrown as
`WebrpcRequestFailedException` (no response) and `WebrpcBadResponseException`
(an unexpected response). A cancelled token throws `OperationCanceledException`.

### JSON

`WebrpcJson.Serialize` and `WebrpcJson.Deserialize` encode the generated types
in the wire format, using the serializer options (`WebrpcJson.Options`, or
`WebrpcJson.Settings` with Newtonsoft.Json) the clients use.

## Customizing requests

Clients send requests with the given `HttpClient`, so customize them with a
`DelegatingHandler`, eg. to add auth headers or sign requests:

```csharp
class AuthHandler : DelegatingHandler
{
    protected override Task<HttpResponseMessage> SendAsync(HttpRequestMessage request, CancellationToken cancellationToken)
    {
        request.Headers.Authorization = new AuthenticationHeaderValue("Bearer", token);
        return base.SendAsync(request, cancellationToken);
    }
}

var http = new HttpClient(new AuthHandler { InnerHandler = new HttpClientHandler() });
```

## Unity

- Use `-json=newtonsoft` with Unity's
  [Newtonsoft Json](https://docs.unity3d.com/Packages/com.unity.nuget.newtonsoft-json@latest)
  package (`com.unity.nuget.newtonsoft-json`).
- Serializers use reflection. With IL2CPP managed code stripping, preserve the
  generated types, eg. with a `link.xml` that preserves the assembly holding the
  generated code.
- On WebGL, where the default `HttpClientHandler` can't open sockets, pass an
  `HttpClient` with a custom `HttpMessageHandler` that sends the requests with
  `UnityWebRequest`. This setup is not covered by the tests here.

## Testing

[tests/interop-csharp](../../tests/interop-csharp) runs the generated clients
against the webrpc interoperability test servers (`webrpc-test`, RPC and REST
suites) and a Go server of a C#-specific schema, once with System.Text.Json and
once with Newtonsoft.Json and `-visibility=internal`. It also builds the
generated code without warnings for `netstandard2.0` with C# 8. It runs as part
of `make test` and requires the .NET SDK:

```sh
make -C tests/interop-csharp test
```
