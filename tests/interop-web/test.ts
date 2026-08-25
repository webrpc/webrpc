// End-to-end interoperability tests for the webrpc `file` core type between
// the Go and TypeScript/JavaScript generators.
//
// The TypeScript client (client.gen.ts), JavaScript client (client.gen.mjs)
// and TypeScript server (server.gen.ts) are generated from the golang-basics
// example schema (_examples/golang-basics/example.ridl). The same client
// suite (plain JSON method, multipart file upload, raw-body file download
// with a byte-compare against the upload, and the JSON error envelope) runs
// in four directions:
//
//   1. TypeScript client -> Go server (go run ../../_examples/golang-basics)
//   2. JavaScript client -> Go server (same instance)
//   3. TypeScript client -> TypeScript server (in-process, node:http)
//   4. JavaScript client -> TypeScript server (same instance)
//
// Go client <-> Go server is covered by the golang-basics example test.
//
// Usage: make test (or: npm install && npm test)
import assert from 'node:assert/strict'
import { ChildProcess, spawn } from 'node:child_process'
import { randomBytes } from 'node:crypto'
import { once } from 'node:events'
import http from 'node:http'
import { Readable } from 'node:stream'
import { pipeline } from 'node:stream/promises'
import { Example, WebrpcFile } from './client.gen'
import { ExampleServer, Intent, Kind, UserNotFoundError, serveExampleRpc } from './server.gen'

const GO_SERVER_PORT = 14243
const GO_SERVER_DIR = '../../_examples/golang-basics'

const fileBytes = async (file: WebrpcFile): Promise<Uint8Array> => {
  return new Uint8Array(await (await file.blob()).arrayBuffer())
}

// clientSuite runs the file upload/download test suite against a running
// Example server (Go or TypeScript) at the given address. The client is the
// generated TypeScript client, or the generated JavaScript client (which has
// the same call surface, but no type declarations).
const clientSuite = async (client: Example, userId: number) => {
  // Plain JSON method on a service that also transfers files.
  const { status } = await client.status()
  assert.equal(status, true, 'expected status to be true')

  // File upload: a multipart request carrying the JSON arguments and the file.
  const avatarBytes = new Uint8Array(randomBytes(1024))
  const avatar = new File([avatarBytes], 'me.png', { type: 'image/png' })
  const { size } = await client.uploadAvatar({ userId, avatar })
  assert.equal(size, avatarBytes.length, 'uploaded avatar size mismatch')

  // File download: metadata from the headers, bytes compared to the upload.
  const download = await client.downloadAvatar({ userId })
  assert.equal(download.name, 'me.png', 'downloaded avatar name mismatch')
  assert.equal(download.contentType, 'image/png', 'downloaded avatar content type mismatch')
  assert.equal(download.size, avatarBytes.length, 'downloaded avatar size mismatch')
  assert.deepEqual(await fileBytes(download), avatarBytes, 'downloaded avatar bytes mismatch')

  // File download with metadata outputs: the file streams as the response
  // body while the remaining outputs arrive in the Webrpc-Response header.
  const withMeta = await client.downloadAvatarWithMeta({ userId })
  assert.equal(withMeta.size, avatarBytes.length, 'downloadAvatarWithMeta size mismatch')
  assert.equal(withMeta.contentType, 'image/png', 'downloadAvatarWithMeta contentType mismatch')
  assert.equal(withMeta.avatar.contentType, 'image/png', 'downloadAvatarWithMeta file contentType mismatch')
  assert.deepEqual(await fileBytes(withMeta.avatar), avatarBytes, 'downloadAvatarWithMeta bytes mismatch')

  // Errors still arrive as the standard webrpc JSON error envelope, even on
  // file download methods.
  await assert.rejects(
    client.downloadAvatar({ userId: 999999 }),
    (err: unknown) => (err as { code?: number }).code === 1000, // UserNotFound
    'expected UserNotFound error for a missing avatar',
  )
}

// createExampleService is a TypeScript implementation of the Example service,
// mirroring the Go implementation in the golang-basics example.
const createExampleService = (): ExampleServer => {
  const avatars = new Map<number, File>()
  const user = { id: 1, USERNAME: 'hihi', role: 'user', kind: Kind.ADMIN, intent: Intent.openSession }

  return {
    async ping() { return {} },
    async status() { return { status: true } },
    async version() { return { version: { webrpcVersion: 'v1', schemaVersion: 'v1.0.0', schemaHash: '' } } },
    async getUser() { return { code: 200, user, profile: '' } },
    async getUserV2() { return { code: 200, user, profile: '' } },
    async findUser(_ctx, { s }) { return { name: s.q, user: { ...user, USERNAME: s.q } } },
    async getIntents() { return { intents: [Intent.openSession, Intent.closeSession] } },
    async countIntents() { return { count: { [Intent.openSession]: 1 } } },

    // The avatar arrives as a native File parsed from the multipart request
    // body, alongside the JSON-decoded scalar fields of the request.
    async uploadAvatar(_ctx, { userId, avatar }) {
      const file = avatar instanceof File ? avatar : new File([avatar], 'avatar', { type: avatar.type })
      avatars.set(userId, file)
      return { size: file.size }
    },

    // Download methods return the file to be streamed as the raw response
    // body; a File carries the filename for the Content-Disposition header.
    async downloadAvatar(_ctx, { userId }) {
      const avatar = avatars.get(userId)
      if (!avatar) {
        throw new UserNotFoundError({ cause: `no avatar for user ${userId}` })
      }
      return avatar
    },

    // Download methods may also declare metadata outputs alongside the file;
    // the generated handler puts them in the Webrpc-Response header.
    async downloadAvatarWithMeta(_ctx, { userId }) {
      const avatar = avatars.get(userId)
      if (!avatar) {
        throw new UserNotFoundError({ cause: `no avatar for user ${userId}` })
      }
      return { avatar, size: avatar.size, contentType: avatar.type }
    },
  }
}

// createExampleHttpServer hosts the webrpc service on a plain node:http
// server, adapting IncomingMessage/ServerResponse to the web-standard
// Request/Response pair the generated handler works with.
const createExampleHttpServer = (): http.Server => {
  const service = createExampleService()

  return http.createServer(async (req, res) => {
    try {
      const response = await serveExampleRpc(service, null, toWebRequest(req))
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
  if (method !== 'GET' && method !== 'HEAD') {
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

const stopGoServer = async (child: ChildProcess) => {
  if (child.exitCode === null && child.pid !== undefined) {
    // Kill the whole process group: `go run` spawns the compiled binary as a
    // child process which would otherwise outlive it.
    try {
      process.kill(-child.pid, 'SIGKILL')
    } catch {
      child.kill('SIGKILL')
    }
    await once(child, 'exit')
  }
}

const main = async () => {
  // The generated plain-JavaScript client, running the same suite. It ships
  // without type declarations, hence the dynamic import.
  // @ts-expect-error client.gen.mjs has no type declarations
  const jsClient = await import('./client.gen.mjs')

  // 1. + 2. TypeScript and JavaScript clients vs the Go example server.
  const goServer = spawn('go', ['run', '.'], {
    cwd: GO_SERVER_DIR,
    env: { ...process.env, PORT: String(GO_SERVER_PORT) },
    stdio: ['ignore', 'inherit', 'inherit'],
    detached: true,
  })
  goServer.on('error', (err) => {
    console.error(`failed to start go server: ${err.message}`)
    process.exit(1)
  })
  try {
    const goServerUrl = `http://localhost:${GO_SERVER_PORT}`
    await waitForServer(goServerUrl)
    console.log('--- interop 1/4: TypeScript client -> Go server')
    await clientSuite(new Example(goServerUrl, fetch), 1001)
    console.log('OK')
    console.log('--- interop 2/4: JavaScript client -> Go server')
    await clientSuite(new jsClient.Example(goServerUrl, fetch), 1002)
    console.log('OK')
  } finally {
    await stopGoServer(goServer)
  }

  // 3. + 4. TypeScript and JavaScript clients vs the TypeScript server.
  const tsServer = createExampleHttpServer()
  await new Promise<void>((resolve) => tsServer.listen(0, resolve))
  const tsAddress = tsServer.address()
  if (tsAddress === null || typeof tsAddress === 'string') {
    throw new Error('unexpected TypeScript server address')
  }
  const tsServerUrl = `http://localhost:${tsAddress.port}`
  try {
    console.log('--- interop 3/4: TypeScript client -> TypeScript server')
    await clientSuite(new Example(tsServerUrl, fetch), 1003)
    console.log('OK')
    console.log('--- interop 4/4: JavaScript client -> TypeScript server')
    await clientSuite(new jsClient.Example(tsServerUrl, fetch), 1004)
    console.log('OK')
  } finally {
    tsServer.close()
  }

  console.log('all interop tests passed')
}

main().catch((err) => {
  console.error(err)
  process.exit(1)
})
