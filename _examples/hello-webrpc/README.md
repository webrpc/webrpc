hello-webrpc-ts
===============

* Server: Go
* Client: Web Browser (Typescript)

Simple client+server app with Go api backend (server) and Typescript Webapp (client).

1. `$ make bootstrap` - installs node modules for the webapp (client)
2. `$ make run-server` - to start the Go server at http://localhost:4242/
3. `$ make run-client` - build+start webapp via webpack dev server at http://localhost:4444/
4. Open your browser to https://localhost:4444/ and open your console, and see rpc calls, tada

webrpc comes with its own schema design language called RIDL, which stands for "RPC interface
design language" :) it reads and feels like documentation, but it very flexible. See
[hello-api.ridl](./hello-api.ridl) for the RIDL file for this service.

as well, webrpc supports a json-formatted schema with the identical functionality as the RIDL format.
See here, [hello-api.webrpc.json](./hello-api.webrpc.json).
