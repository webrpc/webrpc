hello-webrpc 
============

Simple client+server app with Go api backend (server) and Javascript Webapp (client).

1. `$ make tools` - to download `webify` cli to serve the 'webapp/' local files
2. `$ make serve-webapp` - to serve webapp files at http://localhost:4444/
3. `$ make start-api-server` - to start the Go server at http://localhost:4242/
4. Open your browser to https://localhost:4444/ and open your console, and see rpc calls, tada

see the [hello-api.webrpc.json](./hello-api.webrpc.json) service schema for how
the sausage is made.
