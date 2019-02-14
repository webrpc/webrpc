hello-webrpc-ts
===============

* Client: Web Browser (Typescript)
* Server: Go

Simple client+server app with Go api backend (server) and Typescript Webapp (client).

1. `$ cd client && npm install`
2. `$ make serve-webapp` - build+start webapp via webpack dev server at http://localhost:4444/
3. `$ make start-api-server` - to start the Go server at http://localhost:4242/
4. Open your browser to https://localhost:4444/ and open your console, and see rpc calls, tada

webrpc comes with its own schema design language called RIDL, which stands for "RPC interface
design language" :) it reads and feels like documentation, but it very flexible. See
[hello-api.ridl](./hello-api.ridl) for the RIDL file for this service.

as well, webrpc supports a json-formatted schema with the identical functionality as the RIDL format.
See here, [hello-api.webrpc.json](./hello-api.webrpc.json).
