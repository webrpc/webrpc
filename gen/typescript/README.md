webrpc Typescript generator
===========================

This generator, from a webrpc schema/design file will code-generate:

1. Client -- an isomorphic/universal Typescript client to speak to a webrpc server using the
provided schema. This client is compatible with any webrpc server language (ie. Go, nodejs, etc.).
As the client is isomorphic, means you can use this within a Web browser or use the client in a 
server like nodejs -- both without needing any dependencies. I suggest to read the generated TS
output of the generated code, and you shall see, its nothing fancy, just the sort of thing you'd
write by hand.

2. Server -- a nodejs Typescript server handler. See examples.
