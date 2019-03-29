websocket transport for webrpc
==============================

1. Open (connect)
2. Close (disconnect)
3. Retry (re-connect)

4. payload format (json-rpc?)

5. communications
  * client ------> server
  * client <------ server
  * client <-----> server

6. streaming conventions
  * when to end the stream? grpc will pass a tail value of sorts.. is this even interesting to us?
  * perhaps we only/always support bi-directional streaming? therefore they're always open..?
    * stream client ------> server :: use case ?
    * stream client <------ server :: use case ?
    

## Implementation

webrpc is a request/response comm model.

* the websocket package adds a new -websocket flag to the generator

* which will make a websocket transport available for all methods, but as well
will permit the `stream` keyword to work in  sort of pub/sub model in uni or bi-directon

* rpc method which are stream must error when communicating over transport that doesn't support
streaming, such as HTTP/1.1


## Approach

1. Write the code by hand for what we want..

2. then code-generate it... bottom-up approach
