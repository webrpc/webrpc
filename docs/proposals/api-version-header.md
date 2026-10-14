# Proposal: `Api-Version` header

Status: draft, for discussion. No code changes in this PR.

## Summary

Generated clients send one `Api-Version` request header. It carries the API (schema) version the
client was built against, followed by optional metadata tokens for the SDK and the codegen.
Generated servers answer with their own `Api-Version`. Servers learn which API version, SDK and
tooling each caller uses. Clients learn which version the server runs. The existing `Webrpc` header
stays as it is, and can be retired later.

```
Request:   Api-Version: v26.9.12-a1b2c3d; my-sdk@v1.4.0; webrpc/typescript@v0.20.3
Response:  Api-Version: v26.10.13-9f8e7d6; webrpc/golang@v0.22.0
```

## Motivation

`Webrpc: webrpc@v0.47.0;golang@v0.47.0;example-api@v0.1.0` tells a server which generator built the
caller. It falls short in three ways:

- **It is webrpc-only.** With REST routes ([#517](https://github.com/webrpc/webrpc/pull/517)),
  generated clients send `Webrpc` on REST calls too. Callers that don't use webrpc-generated code
  cannot, though: curl, third-party SDKs, and OpenAPI-generated clients. A public API needs one
  documented header that every caller can send.
- **SDK metadata has no place.** An SDK that wraps a generated client cannot say "I am
  `my-sdk@v1.4.0`" without a second header.
- **The schema version is rarely real.** Most schemas keep a static `version = v0.1.0`.
  `-schemaVersion` / `$WEBRPC_SCHEMA_VERSION` exists, but nothing encourages a build-derived value.

A typical consumer is an API gateway that records the API version and the client of every request
for usage analytics, and later pins behaviour to the version.

## Format

```
Api-Version = api-version *( ";" token )
token       = name "@" version          ; name may contain "/", e.g. webrpc/typescript
```

- **`api-version`** is the schema version, taken from RIDL `version` or overridden by
  `-schemaVersion` / `$WEBRPC_SCHEMA_VERSION`. webrpc treats it as opaque. A build-derived value
  is recommended, such as `v<YY>.<M>.<D>-<short sha>` of the last schema change, so that it
  changes exactly when the contract does.
- **Tokens** are optional. Order is not significant. Generated code appends
  `webrpc/<target>@<generator version>`, and applications may append their own, for example the
  SDK name and version.
- **Parsing is tolerant.** A malformed token is skipped, and never fails the request. Servers must
  bound what they store, because tokens are caller-controlled.

## Generated code

Opt-in, behind a template option `-apiVersionHeader=true` (default `false` at first).

**Clients (Go, TypeScript first)**

- Send `Api-Version: <schemaVersion>; webrpc/<target>@<version>` on every request.
- Let the caller append tokens:
  - Go: a client option, for example `WithAPIVersionTokens("my-sdk@v1.4.0")`;
  - TypeScript: a constructor option, `apiVersionTokens: string[]`.
- Expose the server's response value. Go uses a response helper. TypeScript uses an `onResponse`
  hook, or the parsed value on errors.

**Servers (Go first)**

- Set the response `Api-Version: <schemaVersion>; webrpc/<target>@<version>`.
- Parse the request header into the context:
  - `APIVersionFromContext(ctx)` returns the version and the tokens;
  - `ParseAPIVersion(string)` is exported, so gateways and middleware can reuse it.
- A missing header is valid, and means the caller did not say. Making the header required is a
  server policy decision, not a codegen one.

**Shared**

- Generators expose the same `.WebrpcHeader`-style template variable, `.APIVersionHeader`, so
  every target emits an identical value.
- `make generate` keeps stripping version strings from checked-in output (see `AGENTS.md`), now
  including the new constant.
- Interop tests cover a Go client against a TypeScript server, and the reverse, with tokens.

## REST routes (#517)

`Api-Version` works on REST routes without extra work:

- **Client.** `doHTTPRequest` sets the header before `restClient.Do` rewrites the request into its
  REST verb and path, and the request headers carry over.
- **Server.** The REST mux clones the request (`r.Clone`) before calling `serveRPC`, so the
  request header reaches `APIVersionFromContext`. The response header is set in `serveRPC`,
  which REST routes also go through.
- **OpenAPI.** Documenting the header per operation becomes useful once #517's follow-up exports
  real verbs and paths.

## Relationship to other headers

- **`Webrpc`.** It is unchanged and still emitted by default. Once `Api-Version` is on by default
  in every target, `Webrpc` can be deprecated in a later release.
- **`User-Agent`.** It is not touched. It belongs to the browser, runtime or tool (Chrome, curl,
  Node.js, Go).
- **CORS.** Browser deployments must allow `Api-Version` as a request header, and expose it as a
  response header. Document this in each target's README.

## Rollout

1. Agree on the format and the option name (this PR).
2. `gen/golang` and `gen/typescript` emit and parse the header behind `-apiVersionHeader`, with
   interop tests.
3. The other targets (JavaScript, Kotlin, Dart, Swift) follow. OpenAPI documents the header as a
   parameter.
4. Flip the default to `true`, then deprecate `Webrpc`.

## Open questions

- With REST routes, should method annotations drive deprecation? For example,
  `@deprecated:2026-01-15` and `@sunset:2026-07-01` would make generated servers send
  `Deprecation` (RFC 9745, `@<unix>`) and `Sunset` (RFC 8594), and OpenAPI would mark the
  operation `deprecated: true`. chi [v5.5.0](https://github.com/go-chi/chi/releases/tag/v5.5.0)
  now ships `middleware.Deprecation` and `middleware.Sunset`, so generated servers should emit the
  same values. Servers already on chi can use that middleware directly today. That might be a
  separate proposal.

- Should the token be `webrpc/<target>@<target version>`, `webrpc@<webrpc-gen version>`, or both?
  The `Webrpc` header carries both today.
- Should the schema name appear (`example-api@<version>`), or is the API version enough once it is
  build-derived?
- Should a server that reads `Api-Version` stop sending `Webrpc`, or should both stay until the
  default flips?
