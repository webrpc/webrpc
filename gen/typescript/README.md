# webrpc-gen Typescript templates

This repo contains the templates used by the `webrpc-gen` cli to code-generate
webrpc Typescript server and client code.

This generator, from a webrpc schema/design file will code-generate:

1. Client -- an isomorphic/universal Typescript client to speak to a webrpc server using the
   provided schema. This client is compatible with any webrpc server language (ie. Go, nodejs, etc.).
   As the client is isomorphic, means you can use this within a Web browser or use the client in a
   server like nodejs -- both without needing any dependencies. I suggest to read the generated TS
   output of the generated code, and you shall see, its nothing fancy, just the sort of thing you'd
   write by hand.

2. Server -- a nodejs Typescript server handler. See examples.

## Features

### Query Keys for React Query / SWR

The generated client includes a `queryKey` property with type-safe query key generators for each endpoint. This makes it easy to use with popular data-fetching libraries:

```typescript
import { useQuery } from '@tanstack/react-query'
import { Example } from './client.gen'

const client = new Example('http://localhost:3000', fetch)

function UserProfile({ userId }) {
  const { data } = useQuery({
    queryKey: client.queryKey.getUser({ userId }),
    queryFn: ({ signal }) => client.getUser({ userId }, undefined, signal)
  })

  return <div>{data?.user.name}</div>
}
```

The query keys follow the pattern `[ServiceName, methodName, request?]` and are fully type-safe with `as const` assertions.

### Enum style: `enum` vs `union`

TypeScript best practices are moving away from `enum` declarations — they emit runtime
code that can't be erased, which is incompatible with the `erasableSyntaxOnly` tsconfig
option and tooling that expects type-only syntax to disappear at build time.

The `-enumStyle` option lets you choose how schema enums are code-generated:

| `-enumStyle` value | Output |
| ------------------ | ------ |
| `enum` (default) | Traditional TypeScript `export enum` |
| `union` | `as const` object + union type (erasable, no runtime `enum`) |

Given this schema:

```ridl
enum Kind: uint32
  - USER
  - ADMIN

enum Intent: string
  - openSession
  - closeSession
```

`-enumStyle=enum` (default) generates:

```typescript
export enum Kind {
  USER = 'USER',
  ADMIN = 'ADMIN'
}

export enum Intent {
  openSession = 'openSession',
  closeSession = 'closeSession'
}
```

`-enumStyle=union` generates:

```typescript
export const Kind = {
  USER: 'USER',
  ADMIN: 'ADMIN',
} as const
export type Kind = (typeof Kind)[keyof typeof Kind]

export const Intent = {
  openSession: 'openSession',
  closeSession: 'closeSession',
} as const
export type Intent = (typeof Intent)[keyof typeof Intent]
```

The `union` style produces a value (the `const` object) and a type that share the same
name, so it's *mostly* a drop-in replacement anywhere the enum was used — function
parameters, interface fields, and client/server code all reference the type by the same
identifier. You still get autocompletion and access to members via `Kind.USER`, and
nothing is left behind at runtime beyond a plain object.

There are two behavioral differences from a real `enum` to be aware of:

- **No reverse mapping for numeric enums.** A numeric TypeScript `enum` lets you look up a
  member name by its value (`WebrpcErrorCodes[1000]` → `'Unauthorized'`). A `const` object
  has no reverse entries, so `WebrpcErrorCodes[1000]` is `undefined`. Consumer code that maps
  an error code back to its name this way must be updated. (Schema enums are unaffected — they
  are generated with string values regardless of their backing type.)
- **Members can't be used directly as types.** With an `enum` you can write `type T = Kind.USER`.
  In the `union` style `Kind.USER` is a value, so use `type T = typeof Kind.USER` instead.

Passing any value other than `enum` or `union` (e.g. `-enumStyle=foo`) prints an error and exits.

## Usage

```
webrpc-gen -schema=example.ridl -target=typescript -server -client -out=./example.gen.ts
```

or

```
webrpc-gen -schema=example.ridl -target=github.com/webrpc/gen-typescript@v0.7.0 -server -client -out=./example.gen.ts
```

or

```
webrpc-gen -schema=example.ridl -target=./local-templates-on-disk -server -client -out=./example.gen.ts
```

As you can see, the `-target` supports default `typescript`, any git URI, or a local folder :)

### Set custom template variables

Change any of the following values by passing `-option="Value"` CLI flag to `webrpc-gen`.

| webrpc-gen -option | Description                             | Default value | Version |
| ------------------ | --------------------------------------- | ------------- | ------- |
| `-client`          | generate client code                    | `false`       | v0.0.1  |
| `-server`          | generate server code                    | `false`       | v0.0.1  |
| `-webrpcHeader`    | send Webrpc header in all HTTP requests | `true`        | v0.15.0 |
| `-schemaHash=false` | don't emit schema hash + version consts | `true`        | v0.28.0 |
| `-enumStyle`       | enum codegen style: `enum` or `union`   | `enum`        | v0.29.0 |

**Note:** Generated code requires ES2022+ runtime environment.

## LICENSE

[MIT LICENSE](./LICENSE)
