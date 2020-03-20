Streaming TODO
==============


# Go client

- [x] timeout for first chunk response.. should be no timeout, or the keep-alive * 2
- [x] reading + discarding keep-alive packets
- [x] ErrStreamClosed, not really an error? the stream is just done.. what is right client interface here?
- [x] returning { "error": .. } responses, and closing..
- [x] review application-level errors (Download handler returning err)
- [x] review transport-level errors, ie. server disconnecting..
- [x] reconnect pattern..?
- [x] client ErrConnectionLost -- losing connection to server
- [x] panics from server?
- [ ] upload code


# Go server

- [x] succint code: single streamWriter? smaller impl..?
- [x] ErrorPayload + Error review, kinda ugly..
- [x] panics inside of a handler..?
- [ ] upload code



# TS client

- [ ] ..





# Misc

- [ ] chi middleware.Logger -- hmmpf.. when handling streams?
