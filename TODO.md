webrpc v2 TODO
==============

- [x] finalize error handling in the draft
- [x] update ridl parser
- [x] confirm type and struct names must start with upper case
- [x] rename `type` keyword to `alias`
- [x] error if we try to `alias` with a struct or enum type, as wouldn't make sense

--

- [x] update golang-streaming with new errors support on Go side + example.ridl
  - [ ] above is looking good, but could use another review, and testing
        certain situations like streamlost, closed, writing error on the stream, etc.
- [ ] finalize webapp/ source in client.gen.ts with streaming + errors, etc.


--
or

- [ ] add streaming to example again in golang-basics
- [ ] finalize Go client and server, ie. what we'd use for gen-golang

- [ ] finalize ts client in golang-streaming with new errors too


- [ ] update generator for go
- [ ] in generator, for golden ridl, Another.Flatten(etc) .. and Another.GetAccountBalance, should we copy over "meta" at this point, or, look it up during generation..? at schema parse time I think..

- [ ] separate out generators.. ie. `github.com/webrpc/gen-golang , gen-ts, gen-js, gen-csharp, gen-XXX


Streaming TODO
==============
- [ ] stream, then error logic
- [ ] retry logic
- [ ] webapp test
- [ ] node v18 test

- [ ] update example
- [ ] update generator

