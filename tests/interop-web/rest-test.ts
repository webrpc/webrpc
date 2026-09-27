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
// Go client <-> Go server is covered by TestRESTInteroperability in tests/.
//
// Usage: make test (or: npm install && npm test)
import assert from 'node:assert/strict'
import http from 'node:http'
import { once } from 'node:events'
import { Readable } from 'node:stream'
import { pipeline } from 'node:stream/promises'
import { TestApiRest, Status } from './restClient.gen'
import { TestApiRestServer, Item, serveTestApiRestRpc } from './restServer.gen'

// createRestService is a TypeScript implementation of the TestApiRest
// service, mirroring the Go implementation in tests/server/rest.go.
const createRestService = (): TestApiRestServer => {
  const items = new Map<number, Item>()

  return {
    async getItem(_ctx, { id, details }) {
      const item = items.get(id)
      if (!item) {
        throw new Error(`no item ${id}`)
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
        throw new Error(`no item ${id}`)
      }
      item.name = name
      return { item }
    },

    async patchItem(_ctx, { id, name }) {
      const item = items.get(id)
      if (!item) {
        throw new Error(`no item ${id}`)
      }
      if (name !== undefined) {
        item.name = name
      }
      return { item }
    },

    async deleteItem(_ctx, { id }) {
      if (!items.delete(id)) {
        throw new Error(`no item ${id}`)
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

    async resetItems() {
      items.clear()
      return {}
    },

    async rpcMethod(_ctx, { id }) {
      const item = items.get(id)
      if (!item) {
        throw new Error(`no item ${id}`)
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
  if (method !== 'GET' && method !== 'HEAD' && method !== 'DELETE' && method !== 'QUERY') {
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

const createRestHttpServer = (): http.Server => {
  const service = createRestService()

  return http.createServer(async (req, res) => {
    try {
      const response = await serveTestApiRestRpc(service, null, toWebRequest(req))
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

const restSuite = async (addr: string) => {
  const api = new TestApiRest(addr, fetch)

  // POST /rpc/items — request fields in the JSON body.
  const { item: created } = await api.createItem({ item: { id: 7, name: 'alice' } })
  assert.equal(created.name, 'alice', 'createItem name mismatch')

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

  // Body routes reject non-JSON requests, so cross-site form posts (sent
  // without a CORS preflight) can't reach handlers — with or without inputs.
  for (const path of ['/rpc/items', '/rpc/items/reset']) {
    const csrf = await fetch(`${addr}${path}`, {
      method: 'POST',
      headers: { 'Content-Type': 'text/plain' },
      body: '{"item":{"id":5,"name":"csrf"}}',
    })
    assert.equal(csrf.status, 400, `text/plain POST ${path} must return 400`)
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

const main = async () => {
  const tsServer = createRestHttpServer()
  tsServer.listen(0, '127.0.0.1')
  await once(tsServer, 'listening')
  const tsAddr = tsServer.address() as import('node:net').AddressInfo

  try {
    console.log('==> TypeScript client -> TypeScript server (REST routes)')
    await restSuite(`http://127.0.0.1:${tsAddr.port}`)
    console.log('    OK')
  } finally {
    tsServer.close()
  }

  console.log('All REST interop tests passed.')
}

main().catch((err) => {
  console.error(err)
  process.exit(1)
})
