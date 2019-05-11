webrpc golang-nodejs example
============================

* Server: Go
* Client: Nodejs (Javascript)

NOTE: please use node v11.9.0 which has the new experimental es module support, so we can
just import es6 code as you'd expect to work.

## Usage

1. `$ cd client; npm install` -- need to install npm modules for node client, we need the `node-fetch` polyfill
2. `$ make run-server` -- in one terminal, to start the Go server
3. `$ make run-client` -- in another terminal, to run the Nodejs client to make some RPC calls
