webrpc Javascript (ES6) generator
=================================

> NOTE: the javascript generator is almost identical to the typescript one,
> just without the types.. we could have even omitted this generator
> and just relied on `tsc` to compile to es5/es6 JS versions, but
> for ease of use for devs that don't use TS, we offer this generator.

This generator, from a webrpc schema/design file will code-generate:

1. Client -- an isomorphic/universal Javascript client to speak to a webrpc server using the
provided schema. This client is compatible with any webrpc server language (ie. Go, nodejs, etc.).
As the client is isomorphic, means you can use this within a Web browser or use the client in a 
server like nodejs -- both without needing any dependencies. I suggest to read the generated TS
output of the generated code, and you shall see, its nothing fancy, just the sort of thing you'd
write by hand.

2. Server -- a nodejs Javascript server handler. See examples.

## Usage

```
webrpc-gen -schema=example.ridl -target=javascript -server -client -out=./example.gen.ts
```

or 

```
webrpc-gen -schema=example.ridl -target=github.com/webrpc/gen-javascript@v0.7.0 -server -client -out=./example.gen.js
```

or

```
webrpc-gen -schema=example.ridl -target=./local-templates-on-disk -server -client -out=./example.gen.js
```

### Set custom template variables
Change any of the following values by passing `-option="Value"` CLI flag to `webrpc-gen`.

| webrpc-gen -option   | Description                | Default value              |
|----------------------|----------------------------|----------------------------|
| `-client`            | generate client code       | unset (`false`)            |
| `-server`            | generate server code       | unset (`false`)            |
| `-exports=false`     | disable "exports" in code  | enabled (`true`)           |

Example:
```
webrpc-gen -schema=example.ridl -target=javascript -server -client -exports=false -out=./example.gen.js
```

## LICENSE

[MIT LICENSE](./LICENSE)
