// End-to-end interoperability tests for webrpc REST routes between the
// TypeScript client and TypeScript server generators.
//
// The client (restClient.gen.ts) and server (restServer.gen.ts) are generated
// from the TestApiRest service of the tests schema (tests/schema/test.ridl),
// whose methods declare REST routes (verb + path template) instead of the
// default webrpc dispatch. The suite covers path params, query params
// (scalars, repeated keys, absent optionals), JSON bodies on POST/PUT/PATCH,
// the QUERY verb, static-over-param route precedence, the coexisting plain
// RPC method, and the error envelope (405 with Allow, 400 coercion, 404).
//
// The matrix also runs the TypeScript client against the Go REST server
// (webrpc-test -server -rest) and the Go client (webrpc-test -client -rest)
// against this TypeScript server. Go <-> Go is TestRESTInteroperability in tests/.
//
// Usage: make test (or: npm install && npm test)
import assert from 'node:assert/strict'
import { ChildProcess, spawn } from 'node:child_process'
import http from 'node:http'
import { once } from 'node:events'
import { Readable } from 'node:stream'
import { pipeline } from 'node:stream/promises'
import { TestApiRest, Status, WebrpcBadRequestError } from './restClient.gen'
import { TestApiRestServer, Item, UnexpectedValueError, serveTestApiRestRpc } from './restServer.gen'
import { ParamRest, WebrpcBadRequestError as ParamBadRequestError } from './paramRestClient.gen'
import { ParamRestServer, serveParamRestRpc } from './paramRestServer.gen'

// Go REST server and client, built by `make build-test` and found on PATH.
const GO_REST_SERVER_PORT = 14245

// createRestService is a TypeScript implementation of the TestApiRest
// service, mirroring the Go implementation in tests/server/rest.go.
const createRestService = (): TestApiRestServer => {
  const items = new Map<number, Item>()

  return {
    async getItem(_ctx, { id, details }) {
      const item = items.get(id)
      if (!item) {
        throw new UnexpectedValueError({ cause: `no item ${id}` })
      }
      return { item: { ...item, name: details ? `${item.name} (details)` : item.name } }
    },

    async listItems(_ctx, { q, tags, limit }) {
      const echo = `q=${q !== undefined} tags=[${(tags ?? []).join(' ')}] limit=${limit}`
      return { items: [{ id: 0, name: echo }, ...items.values()] }
    },

    async createItem(_ctx, { item }) {
      items.set(item.id, item)
      return { item }
    },

    async updateItem(_ctx, { id, name }) {
      const item = items.get(id)
      if (!item) {
        throw new UnexpectedValueError({ cause: `no item ${id}` })
      }
      item.name = name
      return { item }
    },

    async patchItem(_ctx, { id, name }) {
      const item = items.get(id)
      if (!item) {
        throw new UnexpectedValueError({ cause: `no item ${id}` })
      }
      if (name !== undefined) {
        item.name = name
      }
      return { item }
    },

    async deleteItem(_ctx, { id }) {
      if (!items.delete(id)) {
        throw new UnexpectedValueError({ cause: `no item ${id}` })
      }
      return {}
    },

    async queryItems(_ctx, { q }) {
      return { items: [...items.values()].filter((item) => item.name.includes(q)) }
    },

    async echoParams(_ctx, { str, num, flag, fl, ts, status, tags, opt }) {
      const echo = `str=${str} num=${num} flag=${flag} fl=${fl} ts=${ts} status=${status} tags=[${tags.join(' ')}] opt=${opt ?? '<nil>'}`
      return { echo }
    },

    async getChild(_ctx, { id, childId }) {
      return { path: `/items/${id}/children/${childId}` }
    },

    async getItemStatus(_ctx, { id }) {
      const item = items.get(id)
      if (!item) {
        throw new UnexpectedValueError({ cause: `no item ${id}` })
      }
      return { status: `latest status of item ${id}: ${item.name}` }
    },

    async lookupItem(_ctx, { id }) {
      const item = items.get(id)
      if (!item) {
        throw new UnexpectedValueError({ cause: `no item ${id}` })
      }
      return { item }
    },

    async findItems(_ctx, { q }) {
      return { items: [...items.values()].filter((item) => item.name.includes(q)) }
    },

    async resetItems() {
      items.clear()
      return {}
    },

    async rpcMethod(_ctx, { id }) {
      const item = items.get(id)
      if (!item) {
        throw new UnexpectedValueError({ cause: `no item ${id}` })
      }
      return { item }
    },
  }
}

// node:http <-> web-standard Request/Response adapters, mirroring test.ts.
const toWebRequest = (req: http.IncomingMessage): Request => {
  const method = (req.method || 'GET').toUpperCase()
  const headers = new Headers()
  for (const [key, value] of Object.entries(req.headers)) {
    if (value === undefined) continue
    if (Array.isArray(value)) {
      for (const v of value) headers.append(key, v)
    } else {
      headers.set(key, value)
    }
  }
  const init: RequestInit & { duplex?: 'half' } = { method, headers }
  if (method !== 'GET' && method !== 'HEAD' && method !== 'DELETE') {
    init.body = Readable.toWeb(req) as unknown as BodyInit
    init.duplex = 'half'
  }
  return new Request(`http://${req.headers.host || 'localhost'}${req.url || '/'}`, init)
}

const sendWebResponse = async (res: http.ServerResponse, response: Response): Promise<void> => {
  const headers: Record<string, string> = {}
  response.headers.forEach((value, key) => {
    headers[key] = value
  })
  res.writeHead(response.status, headers)
  if (response.body) {
    await pipeline(Readable.fromWeb(response.body as unknown as import('node:stream/web').ReadableStream), res)
  } else {
    res.end()
  }
}

// createParamService echoes the decoded params, so tests can check each field's value.
const createParamService = (): ParamRestServer => ({
  async probe(_ctx, args) {
    return { echo: JSON.stringify(args) }
  },
  async make(_ctx, { constructor: name }) {
    return { echo: name }
  },
  async tag(_ctx, { name }) {
    return { echo: name }
  },
  async numbers(_ctx, args) {
    return { echo: JSON.stringify(args) }
  },
  async search(_ctx, args) {
    return { echo: JSON.stringify(args) }
  },
  async legacy(_ctx, { name }) {
    return { echo: name }
  },
  async both(_ctx, args) {
    return { echo: JSON.stringify(args) }
  },
  async wild(_ctx, { x }) {
    return { echo: `wild:${x}` }
  },
})

const createHttpServer = (handle: (request: Request) => Promise<Response | null>): http.Server => {
  return http.createServer(async (req, res) => {
    try {
      const response = await handle(toWebRequest(req))
      if (response === null) {
        res.writeHead(404, { 'Content-Type': 'text/plain' })
        res.end('Not Found\n')
        return
      }
      await sendWebResponse(res, response)
    } catch (err) {
      console.error(err)
      if (!res.headersSent) {
        res.writeHead(500, { 'Content-Type': 'text/plain' })
      }
      if (!res.writableEnded) {
        res.end('Internal Server Error\n')
      }
    }
  })
}

// echoQuery is a valid EchoParams query string, with overrides applied.
const echoQuery = (overrides: { [key: string]: string } = {}): string =>
  new URLSearchParams({ str: 's', num: '1', flag: 'true', fl: '1.5', ts: '2026-09-03T12:30:15Z', status: 'NOT_AVAILABLE', tags: 'x', ...overrides }).toString()

// restSuite runs the REST suite against a server. The server is the TypeScript
// one, or the Go one (-rest). Both must agree on status codes and error names.
// Where they differ by design, the check is per server: Go's 64-bit integers
// cover the full range, TypeScript's are limited to safe integers.
const restSuite = async (addr: string, server: 'typescript' | 'go') => {
  const api = new TestApiRest(addr, fetch)

  // POST /rpc/items — request fields in the JSON body; the route declares
  // `+ @success 201`, so the success status must be 201, not the default 200.
  const create = await fetch(`${addr}/rpc/items`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ item: { id: 7, name: 'alice' } }),
  })
  assert.equal(create.status, 201, 'createItem must answer the declared @success status')
  const created = (await create.json()) as { item: Item }
  assert.equal(created.item.name, 'alice', 'createItem name mismatch')

  // GET /rpc/items/{id} — path param plus optional query param.
  const { item: got } = await api.getItem({ id: 7, details: true })
  assert.equal(got.name, 'alice (details)', 'getItem details mismatch')

  // GET /rpc/items — query params: repeated keys, required scalar, absent optional.
  const { items: listed } = await api.listItems({ tags: ['a', 'b'], limit: 5 })
  assert.ok(
    listed.some((item) => item.name === 'q=false tags=[a b] limit=5'),
    `listItems echo mismatch: ${JSON.stringify(listed)}`,
  )

  // PUT /rpc/items/{id} — path param plus JSON body.
  const { item: updated } = await api.updateItem({ id: 7, name: 'bob' })
  assert.equal(updated.name, 'bob', 'updateItem name mismatch')

  // PATCH /rpc/items/{id} — optional body field present, then absent.
  const { item: patched } = await api.patchItem({ id: 7, name: 'carol' })
  assert.equal(patched.name, 'carol', 'patchItem name mismatch')
  const { item: unpatched } = await api.patchItem({ id: 7 })
  assert.equal(unpatched.name, 'carol', 'patchItem without name must keep the value')

  // QUERY /rpc/items/search — the experimental QUERY verb.
  const { items: found } = await api.queryItems({ q: 'carol' })
  assert.equal(found.length, 1, 'queryItems count mismatch')

  // GET /rpc/items/echo/params — every scalar query param type round-trips;
  // the static route must win over GET /rpc/items/{id}/....
  const { echo } = await api.echoParams({
    str: 'hello world/&?=',
    num: -42,
    flag: true,
    fl: 1.5,
    ts: '2026-09-03T12:30:15Z',
    status: Status.NOT_AVAILABLE,
    tags: ['x y', 'z'],
  })
  assert.equal(
    echo,
    'str=hello world/&?= num=-42 flag=true fl=1.5 ts=2026-09-03T12:30:15Z status=NOT_AVAILABLE tags=[x y z] opt=<nil>',
    'echoParams mismatch',
  )

  // GET /rpc/items/{id}/children/{childId} — multiple path params.
  const { path } = await api.getChild({ id: 1, childId: 2 })
  assert.equal(path, '/items/1/children/2', 'getChild path mismatch')

  // GET /rpc/items/{id}/status/latest — literal segments after the path param.
  const { status } = await api.getItemStatus({ id: 7 })
  assert.equal(status, 'latest status of item 7: carol', 'getItemStatus mismatch')

  // A different literal after the param is a different path: 404.
  const statusMiss = await fetch(`${addr}/rpc/items/7/wrong`)
  assert.equal(statusMiss.status, 404, 'wrong trailing literal must not match')

  // LookupItem: the client calls the canonical GET; the alias at its RPC path is a POST with a JSON body.
  const { item: looked } = await api.lookupItem({ id: 7 })
  const lookAlias = await fetch(`${addr}/rpc/TestApiRest/LookupItem`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: '{"id": 7}',
  })
  assert.equal(lookAlias.status, 200, 'the RPC-path alias must work')
  assert.deepEqual((await lookAlias.json()).item, looked, 'canonical and alias must agree')

  // FindItems: GET reads the query string, QUERY reads the JSON body.
  const { items: foundGet } = await api.findItems({ q: 'carol' })
  const foundQuery = await fetch(`${addr}/rpc/items/find`, {
    method: 'QUERY',
    headers: { 'Content-Type': 'application/json' },
    body: '{"q": "carol"}',
  })
  assert.equal(foundQuery.status, 200, 'QUERY alias must work')
  assert.deepEqual((await foundQuery.json()).items, foundGet, 'GET and QUERY must agree')
  assert.equal(foundGet.length, 1, 'findItems count mismatch')

  // A method without a route keeps the default webrpc dispatch.
  const { item: rpcItem } = await api.rpcMethod({ id: 7 })
  assert.equal(rpcItem.name, 'carol', 'rpcMethod name mismatch')

  // A routed method answers only on its REST route.
  const viaRpc = await fetch(`${addr}/rpc/TestApiRest/GetItem`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: '{"id": 7}',
  })
  assert.equal(viaRpc.status, 404, 'routed method must not answer on its RPC path')
  await viaRpc.body?.cancel()

  // The RPC path of a plain method accepts only POST; GET gets 405, not a REST match.
  if (server === 'typescript') {
    const routedGet = await fetch(`${addr}/rpc/TestApiRest/GetItem`)
    assert.equal(routedGet.status, 404, 'a routed RPC path must 404 for GET too')
    await routedGet.body?.cancel()
    const plainGet = await fetch(`${addr}/rpc/TestApiRest/RpcMethod`)
    assert.equal(plainGet.status, 405, 'a plain RPC path must reject GET')
    assert.equal(plainGet.headers.get('Allow'), 'POST')
    await plainGet.body?.cancel()
  }

  // HEAD answers like GET, without a body.
  const head = await fetch(`${addr}/rpc/items/7`, { method: 'HEAD' })
  assert.equal(head.status, 200, 'HEAD on a GET route must succeed')
  assert.equal(await head.text(), '', 'HEAD must not send a body')

  // Empty segments are not dropped: a trailing or doubled slash is another path.
  for (const path of ['/rpc/items/7/', '/rpc/items//7']) {
    const res = await fetch(`${addr}${path}`)
    assert.equal(res.status, 404, `${path} must not match a route`)
    await res.body?.cancel()
  }

  // Fixed segments compare after percent-decoding.
  const encodedFixed = await fetch(`${addr}/rpc/%69tems/7`)
  assert.equal(encodedFixed.status, 200, 'a percent-encoded fixed segment must match')
  await encodedFixed.body?.cancel()

  // Path and query values are range-checked (uint64, uint32, int64).
  const overflow = await fetch(`${addr}/rpc/items/18446744073709551615`)
  assert.equal(overflow.status, 400)
  if (server === 'typescript') {
    assert.match(await overflow.text(), /invalid path param 'id'/)
  } else {
    // Go accepts any uint64; the id is then not found, which is also a 400.
    await overflow.body?.cancel()
  }
  const limitMax = await fetch(`${addr}/rpc/items?tags=a&limit=4294967295`)
  assert.equal(limitMax.status, 200, 'uint32 max must be accepted')
  await limitMax.body?.cancel()
  const limitOver = await fetch(`${addr}/rpc/items?tags=a&limit=4294967296`)
  assert.equal(limitOver.status, 400)
  assert.match(await limitOver.text(), /invalid query param 'limit'/)
  const numMax = await fetch(`${addr}/rpc/items/echo/params?${echoQuery({ num: '9007199254740991' })}`)
  assert.equal(numMax.status, 200, 'the largest safe int64 must be accepted')
  assert.match(String((await numMax.json()).echo), /num=9007199254740991 /)
  const numUnsafe = await fetch(`${addr}/rpc/items/echo/params?${echoQuery({ num: '9007199254740993' })}`)
  if (server === 'typescript') {
    assert.equal(numUnsafe.status, 400)
    assert.match(await numUnsafe.text(), /invalid query param 'num'/)
  } else {
    assert.equal(numUnsafe.status, 200, 'Go int64 takes values beyond the safe integer range')
    await numUnsafe.body?.cancel()
  }

  // A missing required list binds as an empty list, not an error.
  const noTags = await fetch(`${addr}/rpc/items?limit=1`)
  assert.equal(noTags.status, 200, 'a missing required list must not fail')
  assert.equal((await noTags.json()).items[0]?.name, 'q=false tags=[] limit=1')

  // A JSON null body is read as {}, like Go; other non-object bodies are a 400.
  const patchWith = (body: string) =>
    fetch(`${addr}/rpc/items/7`, { method: 'PATCH', headers: { 'Content-Type': 'application/json' }, body })
  const nullPatch = await patchWith('null')
  const emptyPatch = await patchWith('{}')
  assert.equal(nullPatch.status, 200, 'a null JSON body must succeed')
  assert.equal(emptyPatch.status, 200, 'an empty JSON object must succeed')
  assert.deepEqual(await nullPatch.json(), await emptyPatch.json(), 'a null body must behave like {}')
  for (const body of ['[]', '5', '"x"']) {
    const res = await patchWith(body)
    assert.equal(res.status, 400, `a ${body} JSON body must return 400`)
    assert.equal((await res.json()).error, 'WebrpcBadRequest', `a ${body} JSON body must be WebrpcBadRequest`)
  }

  // Allow lists HEAD wherever GET is allowed.
  const allowHead = await fetch(`${addr}/rpc/items/7`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: '{}',
  })
  assert.match(allowHead.headers.get('Allow') ?? '', /\bGET\b.*\bHEAD\b/, 'Allow must list HEAD with GET')
  await allowHead.body?.cancel()

  // The client sends the path value as one encoded segment, so a placeholder-like value stays literal.
  const sent: string[] = []
  const recording = (input: RequestInfo, init?: RequestInit): Promise<Response> => {
    sent.push(String(input))
    return fetch(input, init)
  }
  const recorded = new TestApiRest(addr, recording)
  await assert.rejects(recorded.getChild({ id: '{childId}' as unknown as number, childId: 2 }), 'placeholder id must be rejected by the server')
  assert.equal(sent.at(-1), `${addr}/rpc/items/%7BchildId%7D/children/2`, 'path values must substitute in one pass')

  // The client rejects path values that are empty or dot segments before sending anything.
  const before = sent.length
  for (const id of ['.', '..', '']) {
    await assert.rejects(recorded.getItem({ id: id as unknown as number }), WebrpcBadRequestError, `path value "${id}" must be rejected`)
  }
  assert.equal(sent.length, before, 'dot and empty path values must not be sent')

  // DELETE /rpc/items/{id}.
  await api.deleteItem({ id: 7 })
  await assert.rejects(api.getItem({ id: 7 }), 'getItem after deleteItem must fail')

  // Path params overwrite conflicting JSON body fields (decode order:
  // body -> path).
  await api.createItem({ item: { id: 9, name: 'nine' } })
  const overwrite = await fetch(`${addr}/rpc/items/9`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: '{"id": 1, "name": "dave"}',
  })
  assert.equal(overwrite.status, 200, 'body/path conflict PUT must succeed')
  await overwrite.body?.cancel()
  const { item: overwritten } = await api.getItem({ id: 9 })
  assert.equal(overwritten.name, 'dave', 'path param must win over the body id')
  await assert.rejects(api.getItem({ id: 1 }), 'the body id must not create item 1')

  // Wrong verb on an existing path: error envelope with an Allow header.
  const badVerb = await fetch(`${addr}/rpc/items/7`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: '{}',
  })
  assert.notEqual(badVerb.status, 200, 'wrong verb must not return 200')
  assert.ok(badVerb.headers.get('Allow'), 'wrong verb must return an Allow header')
  await badVerb.body?.cancel()

  // Path param that fails type coercion: HTTP 400 in the error envelope.
  const badParam = await fetch(`${addr}/rpc/items/not-a-number`)
  assert.equal(badParam.status, 400, 'bad path param must return 400')
  await badParam.body?.cancel()

  // A body route with body fields rejects non-JSON requests, so cross-site
  // form posts (sent without a CORS preflight) can't reach handlers. A route
  // with zero body fields drops the check, so its body is parsed as JSON and
  // unknown fields are ignored.
  for (const path of ['/rpc/items', '/rpc/items/reset']) {
    const csrf = await fetch(`${addr}${path}`, {
      method: 'POST',
      headers: { 'Content-Type': 'text/plain' },
      body: '{"item":{"id":5,"name":"csrf"}}',
    })
    assert.equal(csrf.status, path === '/rpc/items/reset' ? 200 : 400, `text/plain POST ${path} status mismatch`)
    await csrf.body?.cancel()
  }
  await assert.rejects(api.getItem({ id: 5 }), 'the text/plain POST must not have created an item')

  // A body route without inputs still works from the generated client.
  await api.resetItems()

  // A missing required query param is a 400, not a missing field.
  const noLimit = await fetch(`${addr}/rpc/items?tags=a`)
  assert.equal(noLimit.status, 400, 'missing required query param must return 400')
  assert.match(await noLimit.text(), /missing query param 'limit'/)

  // An empty required list sends no query keys and decodes as empty.
  const { items: emptyList } = await api.listItems({ tags: [], limit: 1 })
  assert.equal(emptyList[0]?.name, 'q=false tags=[] limit=1', 'empty tags list mismatch')
}

// paramSuite covers the param parsing rules of the Go server: bool spellings,
// sized integer bounds, and field names that shadow Object.prototype members.
const paramSuite = async (addr: string) => {
  const api = new ParamRest(addr, fetch)
  const probeUrl = (query: string) => fetch(`${addr}/rpc/params/probe?${query}`)
  const valid = 'constructor=c&i8=-1&u8=1&i32=1&flag=1'

  const args = { constructor: 'c', toString: 't', i8: -128, u8: 255, i32: -2147483648, flag: true }
  const { echo } = await api.probe(args)
  assert.deepEqual(JSON.parse(echo), args, 'fields named constructor and toString must bind')

  for (const v of ['1', 't', 'T', 'TRUE', 'true', 'True']) {
    const res = await probeUrl(valid.replace('flag=1', `flag=${v}`))
    assert.equal(res.status, 200, `flag=${v} must parse as true`)
    assert.equal(JSON.parse((await res.json()).echo).flag, true)
  }
  for (const v of ['0', 'f', 'F', 'FALSE', 'false', 'False']) {
    const res = await probeUrl(valid.replace('flag=1', `flag=${v}`))
    assert.equal(res.status, 200, `flag=${v} must parse as false`)
    assert.equal(JSON.parse((await res.json()).echo).flag, false)
  }
  for (const v of ['yes', 'on', '2', '']) {
    const res = await probeUrl(valid.replace('flag=1', `flag=${v}`))
    assert.equal(res.status, 400, `flag=${v} must be rejected`)
    assert.match(await res.text(), /invalid query param 'flag'/)
  }

  const bounds: [string, string, boolean][] = [
    ['i8=127', 'i8', true],
    ['i8=128', 'i8', false],
    ['i8=-129', 'i8', false],
    ['i8=1.5', 'i8', false],
    ['u8=256', 'u8', false],
    ['u8=-1', 'u8', false],
    ['i32=2147483647', 'i32', true],
    ['i32=2147483648', 'i32', false],
  ]
  for (const [param, name, ok] of bounds) {
    const res = await probeUrl(valid.replace(new RegExp(`${param.split('=')[0]}=[^&]*`), param))
    if (ok) {
      assert.equal(res.status, 200, `${param} must be accepted`)
    } else {
      assert.equal(res.status, 400, `${param} must be rejected`)
      assert.match(await res.text(), new RegExp(`invalid query param '${name}'`))
      continue
    }
    await res.body?.cancel()
  }

  const missing = await probeUrl('i8=1&u8=1&i32=1&flag=1')
  assert.equal(missing.status, 400)
  assert.match(await missing.text(), /missing query param 'constructor'/)

  // A body field named constructor is required like any other field.
  const make = await fetch(`${addr}/rpc/params/make`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: '{}',
  })
  assert.equal(make.status, 400, 'a missing body field named constructor must return 400')
  assert.match(await make.text(), /missing argument `constructor`/)
  const made = await api.make({ constructor: 'm' })
  assert.equal(made.echo, 'm')

  // A path value is one decoded segment, so an encoded %2e%2e is a literal name.
  const tagged = await api.tag({ name: '%2e%2e' })
  assert.equal(tagged.echo, '%2e%2e', 'an encoded dot segment must round-trip')
  for (const name of ['.', '..', '']) {
    await assert.rejects(api.tag({ name }), ParamBadRequestError, `path value "${name}" must be rejected`)
  }

  // Floats are finite decimals; float32 is bounded by its max, float64 only by finiteness.
  const numbers = (overrides: { [key: string]: string }) =>
    fetch(`${addr}/rpc/params/numbers?${new URLSearchParams({ f32: '1.5', f64: '1.5', ts: '2026-09-03T12:30:15Z', ...overrides })}`)
  const acceptedNumbers: [string, string][] = [['f32', '1'], ['f32', '1.'], ['f32', '.5'], ['f32', '-1.5e3'], ['f32', '+2E-4'], ['f32', '3.4028234663852886e38'], ['f64', '1e308']]
  for (const [key, value] of acceptedNumbers) {
    const res = await numbers({ [key]: value })
    assert.equal(res.status, 200, `${key}=${value} must be accepted`)
    await res.body?.cancel()
  }
  const rejectedNumbers: [string, string][] = [['f32', '0x10'], ['f32', ' 1 '], ['f32', '1e'], ['f32', 'NaN'], ['f32', 'Inf'], ['f32', '1_0'], ['f32', '3.5e38'], ['f64', '0x10'], ['f64', '1e400']]
  for (const [key, value] of rejectedNumbers) {
    const res = await numbers({ [key]: value })
    assert.equal(res.status, 400, `${key}=${JSON.stringify(value)} must be rejected`)
    assert.match(await res.text(), new RegExp(`invalid query param '${key}'`))
  }

  // Timestamps are RFC 3339 as in Go: a real date, uppercase T and Z, and a zone.
  for (const ts of ['2026-09-03T12:30:15.123Z', '2024-02-29T00:00:00+05:30', '2026-09-03T12:30:15-00:00']) {
    const res = await numbers({ ts })
    assert.equal(res.status, 200, `ts=${ts} must be accepted`)
    await res.body?.cancel()
  }
  for (const ts of ['2026-02-30T00:00:00Z', '2026-02-29T00:00:00Z', '2026-04-31T00:00:00Z', '2026-13-01T00:00:00Z', '2026-09-03T24:00:00Z', '2026-09-03T23:59:60Z', '2026-09-03 12:30:15Z', '2026-09-03T12:30:15z', '2026-09-03T12:30:15', '2026-09-03T12:30Z', '2026-09-03T12:30:15+25:00']) {
    const res = await numbers({ ts })
    assert.equal(res.status, 400, `ts=${ts} must be rejected`)
    assert.match(await res.text(), /invalid query param 'ts'/)
  }

  // A query field binds by its JSON name, dashed_term, which is also its TS field name.
  const searched = await api.search({ term: 't', dashed_term: 'd' })
  assert.deepEqual(JSON.parse(searched.echo), { term: 't', dashed_term: 'd' }, 'dashed_term must bind')
  const wire = await fetch(`${addr}/rpc/params/search?term=t&dashed=x`)
  assert.deepEqual(JSON.parse((await wire.json()).echo), { term: 't' }, 'the RIDL field name must not bind')
}

// aliasSuite covers methods with several routes: the client calls the canonical
// route, and server-only aliases (including a legacy RPC path) bind by their own verb.
const aliasSuite = async (addr: string) => {
  const sent: string[] = []
  const recording = (input: RequestInfo, init?: RequestInit): Promise<Response> => {
    sent.push(`${init?.method} ${input}`)
    return fetch(input, init)
  }
  const api = new ParamRest(addr, recording)
  const json = { 'Content-Type': 'application/json' }

  // The canonical route is the first one: GET with a query string.
  const viaClient = await api.legacy({ name: 'x' })
  assert.equal(viaClient.echo, 'x')
  assert.equal(sent.at(-1), `GET ${addr}/rpc/params/legacy?name=x`, 'the client must call the canonical route')

  // The alias equal to the method's RPC path is a normal REST route: POST with a JSON body.
  const viaAlias = await fetch(`${addr}/rpc/ParamRest/Legacy`, { method: 'POST', headers: json, body: '{"name": "x"}' })
  assert.equal(viaAlias.status, 200, 'the RPC-path alias must work')
  assert.deepEqual(await viaAlias.json(), viaClient, 'canonical and alias must agree')
  const aliasBadVerb = await fetch(`${addr}/rpc/ParamRest/Legacy`)
  assert.equal(aliasBadVerb.status, 405, 'GET on the alias path must be a 405 from the route table')
  await aliasBadVerb.body?.cancel()

  // GET reads query params, QUERY reads the JSON body, same implementation.
  const viaGet = await api.both({ q: 'a', n: 2 })
  assert.equal(sent.at(-1), `GET ${addr}/rpc/params/both?q=a&n=2`, 'the client must call the first route')
  const viaQuery = await fetch(`${addr}/rpc/params/both`, { method: 'QUERY', headers: json, body: '{"q": "a", "n": 2}' })
  assert.equal(viaQuery.status, 200, 'QUERY route must work')
  assert.deepEqual(await viaQuery.json(), viaGet, 'GET and QUERY must agree')
  const allowBoth = await fetch(`${addr}/rpc/params/both`, { method: 'POST', headers: json, body: '{}' })
  assert.equal(allowBoth.status, 405)
  assert.equal(allowBoth.headers.get('Allow'), 'GET, HEAD, QUERY', 'Allow must be sorted lexically')
  await allowBoth.body?.cancel()

  // An alias on an old RPC name and a single-segment param route must not be shadowed by the RPC fast path.
  const oldName = await fetch(`${addr}/rpc/ParamRest/OldMake`, { method: 'POST', headers: json, body: '{"constructor": "c"}' })
  assert.equal(oldName.status, 200, 'an alias on an old RPC name must work')
  assert.deepEqual(await oldName.json(), { echo: 'c' })
  const wild = await fetch(`${addr}/rpc/ParamRest/zzz`, { method: 'POST', headers: json, body: '{}' })
  assert.equal(wild.status, 200, 'a POST /ParamRest/{x} route must match')
  assert.deepEqual(await wild.json(), { echo: 'wild:zzz' })

  const queryNoBody = await fetch(`${addr}/rpc/params/both?q=a&n=2`, { method: 'QUERY', headers: json, body: '{}' })
  assert.equal(queryNoBody.status, 400, 'QUERY must not read query params')
  await queryNoBody.body?.cancel()
}

const waitForServer = async (addr: string, timeoutMs = 60_000) => {
  const deadline = Date.now() + timeoutMs
  for (;;) {
    try {
      await fetch(addr, { method: 'HEAD' })
      return
    } catch {
      if (Date.now() > deadline) {
        throw new Error(`server at ${addr} did not come up within ${timeoutMs}ms`)
      }
      await new Promise((resolve) => setTimeout(resolve, 100))
    }
  }
}

// runWebrpcTest runs the Go webrpc-test binary and fails on a non-zero exit.
const runWebrpcTest = async (args: string[]) => {
  const child = spawn('webrpc-test', args, { stdio: 'inherit' })
  const [code] = await once(child, 'exit')
  if (code !== 0) {
    throw new Error(`webrpc-test ${args.join(' ')} exited with ${code}`)
  }
}

// withGoRestServer runs the Go REST test server for the duration of fn.
const withGoRestServer = async (fn: (url: string) => Promise<void>) => {
  const goServer: ChildProcess = spawn('webrpc-test', ['-server', '-rest', `-port=${GO_REST_SERVER_PORT}`, '-timeout=10m'], {
    stdio: ['ignore', 'inherit', 'inherit'],
  })
  try {
    const url = `http://127.0.0.1:${GO_REST_SERVER_PORT}`
    await waitForServer(url)
    await fn(url)
  } finally {
    if (goServer.exitCode === null) {
      goServer.kill('SIGKILL')
      await once(goServer, 'exit')
    }
  }
}

const main = async () => {
  const restService = createRestService()
  const paramService = createParamService()
  const tsServer = createHttpServer((request) => serveTestApiRestRpc(restService, null, request))
  const paramServer = createHttpServer((request) => serveParamRestRpc(paramService, null, request))
  tsServer.listen(0, '127.0.0.1')
  paramServer.listen(0, '127.0.0.1')
  await Promise.all([once(tsServer, 'listening'), once(paramServer, 'listening')])
  const tsAddr = tsServer.address() as import('node:net').AddressInfo
  const paramAddr = paramServer.address() as import('node:net').AddressInfo

  try {
    console.log('==> TypeScript client -> TypeScript server (REST routes)')
    await restSuite(`http://127.0.0.1:${tsAddr.port}`, 'typescript')
    console.log('    OK')
    console.log('==> TypeScript client -> TypeScript server (REST param parsing)')
    await paramSuite(`http://127.0.0.1:${paramAddr.port}`)
    console.log('    OK')

    console.log('==> TypeScript client -> TypeScript server (REST aliases)')
    await aliasSuite(`http://127.0.0.1:${paramAddr.port}`)
    console.log('    OK')

    console.log('==> TypeScript client -> Go server (REST routes)')
    await withGoRestServer(async (goServerUrl) => restSuite(goServerUrl, 'go'))
    console.log('    OK')

    // The Go client runs its own REST suite against the TypeScript server.
    console.log('==> Go client -> TypeScript server (REST routes)')
    await runWebrpcTest(['-client', '-rest', '-skip=PatchItem(nil)', `-url=http://127.0.0.1:${tsAddr.port}`])
    console.log('    OK')
  } finally {
    tsServer.close()
    paramServer.close()
  }

  console.log('All REST interop tests passed.')
}

main().catch((err) => {
  console.error(err)
  process.exit(1)
})
