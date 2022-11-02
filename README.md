<img alt="webrpc" src="https://github.com/webrpc/webrpc/raw/master/webrpc.png" width="1024" />

webrpc is a schema-driven approach to writing backend servers for the Web. Write your server's
api interface in a schema format of [RIDL](./_examples/golang-basics/example.ridl) or [JSON](./_examples/golang-basics/example.webrpc.json),
and then run `webrpc-gen` to generate the networking source code for your server and client apps. From the schema,
webrpc-gen will generate application base class types/interfaces, JSON encoders, and networking code. In doing
so, it's able to generate fully functioning and typed client libraries to communicate with your server. Enjoy
strongly-typed Web services and never having to write an API client library again.

Under the hood, webrpc is a Web service meta-protocol, schema and code-generator tool for
simplifying the development of backend services for modern Web applications.

| Generator                                              | Description           | Schema | Client | Server |
|--------------------------------------------------------|-----------------------|--------|--------|--------|
| [golang](https://github.com/webrpc/gen-golang)         | Go 1.14+              | v1     | ✅     | ✅     |
| [typescript](https://github.com/webrpc/gen-typescript) | TypeScript            | v1     | ✅     | ✅     |
| [javascript](https://github.com/webrpc/gen-javascript) | JavaScript (ES6)      | v1     | ✅     | ✅     |
| [openapi](https://github.com/webrpc/gen-openapi)       | OpenAPI 3.x (Swagger) | v1     | ✅ [*](https://github.com/swagger-api/swagger-codegen#overview) | ✅ [*](https://github.com/swagger-api/swagger-codegen#overview) |

..contribute more! [webrpc generators](./gen/) are just Go templates (similar to [Hugo](https://gohugo.io/templates/) or [Helm](https://helm.sh/docs/chart_best_practices/templates/)).

## Quick example

Here is an example webrpc schema in RIDL format (a new documentation-like format introduced by webrpc)

```
webrpc = v1

name = your-app
version = v0.1.0

message User
  - id: uint64
  - username: string
  - createdAt?: timestamp

message UsersQueryFilter
  - page?: uint32
  - name?: string
  - location?: string

service ExampleService
  - Ping()
  - Status() => (status: bool)
  - GetUserByID(userID: uint64) => (user: User)
  - IsOnline(user: User) => (online: bool)
  - ListUsers(q?: UsersQueryFilter) => (page: uint32, users: []User)
```

WebRPC is a design/schema-driven approach to writing backend servers. Write your server's
api interface in a schema format of [RIDL](./_examples/golang-basics/example.ridl) or
[JSON](./_examples/golang-basics/example.webrpc.json) format and run `webrpc-gen` to generate
source code for your target language.

For example, to generate webrpc server+client code -- run:

```
bin/webrpc-gen -schema=example.ridl -target=go -pkg=main -server -client -out=./example.gen.go
```

and see the generated `./example.gen.go` file of types, server and client in Go. This is essentially
how the `golang-basics` example was built.


### More example apps

* [hello-webrpc](./_examples/hello-webrpc) - webrpc service with Go server and Javascript webapp
* [hello-webrpc-ts](./_examples/hello-webrpc-ts) - webrpc service with Go server and Typescript webapp
* [golang-basics](./_examples/golang-basics) - webrpc service with Go server and Go client
* [golang-nodejs](./_examples/golang-nodejs) - webrpc service with Go server and nodejs (Javascript ES6) client
* [node-ts](./_examples/node-ts) - webrpc service with nodejs server and Typescript webapp client


## Why

**TLDR;** its much simpler + faster to write and consume a webrpc service than traditional approaches
like a REST api or gRPC service.

  1. Code-generate your client libraries in full -- never write another API client again
  2. Compatible with the Web. A Webrpc server is just a HTTP/HTTPS server that speaks JSON, and thus
     all existing browsers, http clients, load balancers, proxies, caches, and tools work
     out of the box (versus gRPC). cURL "just works".
  3. Be more productive, write more correct systems.

---

Writing a Web service / microservice takes a lot of work and time. REST is making me tired.
There are many pieces to build -- designing the routes of your service, agreeing on conventions
for the routes with your team, the request payloads, the response payloads, writing the actual server
logic, routing the methods and requests to the server handlers, implementing the handlers, and
then writing a client library for your desired language so it can speak to your Web
service. Yikes, it's a lot of work. Want to add an additional field or handler? yea, you
have to go through the entire cycle. And what about type-safety across the wire?

webrpc automates a lot the work for you. Now from a single [webrpc schema file](./schema/README.md),
you can use the `webrpc-gen` cli to generate source code for:
* Strongly-typed request / response data payloads for your target language
* Strongly-typed server interface and methods on the service, aka the RPC methods
* Complete client library to communicate with the web service


## Design / architecture

webrpc services speak JSON, as our goals are to build services that communicate with webapps.
We optimize for developer experience, ease of use and productivity when building backends
for modern webapps. However, webrpc also works great for service<->service communication,
but it won't be as fast as gRPC in that scenario, but I'd be surprised to hear if for the majority
of cases that this would be a bottleneck or costly tradeoff.

webrpc is heavily inspired by gRPC and Twirp. It is architecturally the same and has a similar
workflow, but simpler. In fact, the webrpc schema is similar in design to protobuf, as
in we have messages and rpc methods, but the type system is arguably more flexible and
code-gen tooling is simpler. The [webrpc schema](./schema/README.md) is a documentation-like
language for describing a server's api interface and the type system within is inspired by Go,
Typescript and WASM.

We've been thinking about webrpc's design for years, and were happy to see gRPC and Twirp
come onto the scene and pave the way with some great patterns. Over the years and after writing
dozens of backends for Javascript-based Webapps and native mobile apps, and even built prior
libraries like [chi](https://github.com/go-chi/chi), a HTTP router for Go -- we asked ourselves: 

Why have "Rails" and "Django" been such productive frameworks for writing webapps? And the answer
we came to is that its productive because the server and client are the same program,
running in the same process on the same computer. Rails/Django/others like it, when rendering
client-state can just call a function in the same program, the client and the server
are within the same domain and same state -- everything is a function-call away. Compare this to
modern app development such as writing a React.js SPA or a native iOS mobile app, where the app
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
especially if you have to write it by hand and then maintain robust code in both the client and
the server. Ahh, I just want to call a function on my server from my app! Save yourself the work
and time, and code-generate it instead - Enter gRPC / Twirp .. and now, webrpc :) 


Future goals/work:
1. Add RPC streaming support for client/server
2. More code generators.. for Rust, Python, ..


## Getting started

1.  * for Go v1.17+: `go install github.com/webrpc/webrpc/cmd/webrpc-gen`
    * for earlier versions: `go get -u github.com/webrpc/webrpc/cmd/webrpc-gen`  
2. Write+design a [webrpc schema file](./_examples/golang-basics/example.ridl) for your Web service
3. Run the code-generator to create your server interface and client, ie.
  * `webrpc-gen -schema=example.ridl -target=go -pkg=service -server -client -out=./service/proto.gen.go`
  * `webrpc-gen -schema=example.ridl -target=ts -pkg=client -client -out=./web/client.ts`
4. Implement the handlers for your server -- of course, it can't guess the server logic :)

another option is to copy the [hello-webrpc](./_examples/hello-webrpc) example, and adapt for your own webapp and server.


## Schema

The webrpc schema type system is inspired by Go and TypeScript, and is simple and flexible enough
to cover the wide variety of language targets, designed to target RPC communication with Web
applications and other Web services.

High-level features:

  * RIDL, aka RPC IDL, aka "RPC interface design language", format - a documentation-like schema
  format for describing a server application.
  * JSON schema format is also supported if you prefer to write tools to target webrpc's code-gen tools
  * Type system inspired by Go + Typescript
    * integers, floats, byte, bool, any, null, date/time
    * lists (multi-dimensional arrays supported too)
    * maps (with nesting / complex structures)
    * structs / objects
      * optional fields, default values, and pluggable code-generation for a language target
    * enums

For more information please see the [schema readme](./schema/README.md).


## Building from source / making your own code-generator

### Dev

1. Install Go 1.17+
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


## Authors

* [Peter Kieltyka](https://github.com/pkieltyka)
* [José Carlos Nieto](https://github.com/xiam)
* [Vojtech Vitek](https://github.com/VojtechVitek)
* ..and full list of [contributors](https://github.com/webrpc/webrpc/graphs/contributors)!


## Credits

* [Twirp authors](https://github.com/twitchtv/twirp) for making twirp. Much of the webrpc-go
library comes from the twirp project.
* [gRPC authors](https://grpc.io), for coming up with the overall architecture and patterns
for code-generating the bindings between client and server from a common IDL.


## License

MIT
