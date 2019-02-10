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

2. Server (TODO) -- a nodejs-compatible Javascript server handler. TODO / still needs to be written,
but shouldn't be too hard, its just a simple interface that would use the same types are the client
code. For someone interested to writing it, look at how we reuse types in the Go generator between
client and server.


## webrpc-gen extras

The `webrpc-gen` tool accepts an `-extra` cli flag, which is passed down to the generator, and
we use it for some basic toggling during code-generation. For `javascript` generator, we make
es6 code that includes module "exports" by default, but, if you'd like to use the client in a
UMD scenario without having to use babel or something like it, you can pass `-extra=noexports`
option to the cli.. like we do in the [hello-webrpc]() example. 

ie..

```
webrpc-gen -schema=hello-api.webrpc.json -target=js -extra=noexports -client -out=./webapp/client.gen.js
```
