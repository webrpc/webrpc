<img alt="webrpc" src="https://github.com/webrpc/webrpc/raw/master/webrpc.png" width="1024" />

webrpc is a Web service meta-protocol, schema and code-generator tool for simplifying
the development of backend services for modern Web applications.

Enjoy strongly-typed Web services and never having to write an API client library again.

Current code-generation language targets:
* [Go](./gen/golang)
* [Typescript](./gen/typescript)
* [Javascript](./gen/javascript)
* .. contribute more! they're just templates


## Why

Writing a Web service / microservice takes a lot of work and time. There are many pieces
to build -- designing the routes of your service, agreeing on conventions for the routes
with your team, the request payloads, the response payloads, writing the actual server logic,
routing the methods and requests to the server handlers, implementing the handlers, and
then writing a client library for your desired language so it can speak to your Web
service. Yikes, it's a lot of work. Want to add an additional field or handler? yea, you
have to go through the entire cycle. And what about type-safety across the wire?

webrpc automates a lot the work for you. Now from a single [webrpc schema file](./schema/README.md),
you can use the `webrpc-gen` cli to generate source code for:
* Strongly-typed request / response data payloads for your target language
* Strongly-typed server interface and methods on the service, aka the RPC methods
* Complete client library to communicate with the web service

See examples:
* [hello-webrpc](./_examples/hello-webrpc) - webrpc service with Go server and Javascript webapp
* [golang-basics](./_examples/golang-basics) - Go webrpc service with generated Go client
* [golang-nodejs](./_examples/golang-nodejs) - Go webrpc service with generated Typescript client running with nodejs


## Design / architecture

webrpc services speak JSON, as our goals are to build services that communicate with webapps.
We optimize for developer experience, ease of use and productivity when building backends
for modern webapps. However, webrpc also works great for service<->service communication,
but it won't be as fast as gRPC in that scenario, but I'd be surprised hear, for majority cases
of cases, for this to be a bottleneck or costly tradeoff.

webrpc is heavily inspired by gRPC and Twirp. It is architecturally the same and has a similar
workflow, but simpler. In fact, the webrpc schema is similar in design to protobuf, as
in we have messages and rpc methods, but the type system is arguably more flexible and
code-gen tooling is simpler. The [webrpc type system](./schema/README.md) is inspired by Go,
Typescript and WASM.

We've been thinking about webrpc's design for years, and were happy to see gRPC and Twirp
come onto the scene and pave the way with some great patterns. Over the years and after writing
dozens of backends for Javascript-based Webapps and native mobile apps, and even built prior
libraries like [chi](https://github.com/go-chi/chi), a HTTP router for Go -- we asked ourselves: 

Why have "Rails" and "Django" been such productive frameworks for writing webapps? And the answer
we came to is that its productive because the server and client are the same program,
running in the same process on the same computer. Rails/Django/others like it, when rendering
client-state can just call a function on the in the same program, the client and the server
are within the same domain and same state - everything is a function-call away. Compare this to
modern app development, as writing a React.js SPA or a native iOS mobile app, where the app
speaks to an external API server with now the huge added effort to bridge data/runtime from
one namespace (the app) to an entirely other namespace (the server). It's too much work and
takes too much time, and is too brittle. There is a better way! instead of writing the code..
just generate it. If we generate all of the code to native objects in both app/server,
suddenly, we can make a remote service once again feel like calling a method on the same
program running on the same computer/process. Remote-Procedure-Call works!

Finally, we'd like to compare generated RPC services (gRPC/Twirp/webrpc/other) to the most
common pattern to writing services by "making a RESTful API", where the machinery is similar
to RPC services. Picture the flow of data when a client calls out to a server -- from a client
runtime proxy-object, we encode that object, send it over the wire, the server decodes it into
a server runtime proxy-object, the server handler queries the db, returns a proxy object,
encodes it, and sends the function return data over the wire again. That is a ton of work,
especially if you have to write it by hand and then maintain robust code in both namespaces --
aka REST. But, I just want to call a function on my server! Save yourself the work and time,
and code-generate it instead - Enter gRPC / Twirp .. and now, webrpc :) 


Future goals/work:
1. Add RPC streaming support from client/server using websockets
2. Create an [IDL](https://en.wikipedia.org/wiki/IDL_(programming_language)) as a frontend
to the webrpc json schema file. Our idea is it could look similar to a Typescript definition
file that would generate this JSON file, which in-turn would be used by the code-generators
to generate server+client code - we have an open issue [here](https://github.com/webrpc/webrpc/issues/1).


## Getting started

1. `go get -u github.com/webrpc/webrpc/cmd/webrpc-gen`
2. Write+design a [webrpc schema file](./_examples/golang-basics/example.webrpc.json) for your Web service
3. Run the code-generator to create your server interface and client, ie.
  * `webrpc-gen -schema=example.webrpc.json -target=go -pkg=service -server -client -out=./service/proto.gen.go`
  * `webrpc-gen -schema=example.webrpc.json -target=ts -pkg=client -client -out=./web/client.ts`
4. Implement the handlers for your server -- of course, it can't guess the server logic :)

another option is copy the [hello-webrpc](./_examples/hello-webrpc) example, and adapt for your own webapp and server.


## Schema

The schema for webrpc we've designed is inspired by Go and TypeScript type systems, and
is simple enough to cover the wide variety of language targets, designed to target RPC
communication with Web applications and other Web services.

High-level schema features:

  * integers, floats, byte, bool, any, null, date/time
  * lists (multi-dimensional arrays supported too)
  * maps (with nesting / complex structures)
  * structs / objects
    * optional fields, default values, and pluggable code-generation for a language target
  * enums

For more information please see the [schema readme](./schema/README.md).


## Building from source / making your own code-generator

### Dev

1. Install Go 1.11+
2. $ `go get -u github.com/webrpc/webrpc/...`
3. $ `make tools`
4. $ `make build`
5. $ `make test`
6. $ `go install ./cmd/webrpc-gen`


### Writing your own code-generator

Some tips..

1. Copy `gen/golang` to `gen/<yourtargetlang>` and start writing templates
2. Write an example service and use `make build` to regenerate 
3. Write tests, TDD is a great approach to confirm things work


## We're hiring!

Our team at https://horizongames.net is building Arcadeum.net, a distributed
network and platform for blockchain based video games :) built for Ethereum.

If you're passionate about distributed systems, cryptography, privacy, and
writing awesome network infrastructure to help power the Arcadeum network, please
write to us, hello at arcadeum.net


## License

MIT
