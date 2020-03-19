Streaming TODO
==============


# Go client

- [x] timeout for first chunk response.. should be no timeout, or the keep-alive * 2
- [x] reading + discarding keep-alive packets
- [ ] returning { "error": .. } responses, and closing..
- [ ] ErrStreamClosed, not really an error? the stream is just done.. what is right client interface here?
- [ ] review application-level errors (Download handler returning err)
- [ ] review transport-level errors, ie. server disconnecting..
- [ ] reconnect pattern..?
- [ ] upload code
- [ ] connection scenarios
      - [ ] server disconnects
      - [ ] client disconnects
      - [ ] server finished stream transmission
      - [ ] client finishes stream retrieval
      - [ ] client wants to reconnect after disconnect


# Go server

- [ ] succint code: single streamWriter? smaller impl..?
- [ ] ErrorPayload + Error review, kinda ugly..
- [ ] upload code



# TS client

- [ ] ..





# Misc

- [ ] chi middleware.Logger -- hmmpf.. when handling streams?
