import Fastify, { FastifyRequest } from 'fastify'
import { Readable } from 'node:stream'
import { pipeline } from 'node:stream/promises'
import { ExampleServer, serveExampleRpc, Kind, WebrpcEndpointError, WebrpcBadRouteError } from './server.gen'

// ---------------------------------------------------------------------------
// Example rpc service
// ---------------------------------------------------------------------------
interface RequestContext {
  reqId: string
}

// In-memory storage of the uploaded avatars.
const avatars = new Map<number, File>()

// NOTE: you can also implement the service methods in a class.
const exampleService: ExampleServer<RequestContext> = {
  async ping() {
    return {}
  },

  async getUser(ctx, { userId }) {
    if (userId === 911) {
      throw new WebrpcEndpointError({ cause: 'User 911 is forbidden' })
    }

    return {
      code: 200,
      user: {
        id: userId,
        USERNAME: `user-${userId}`,
        role: Kind.USER,
        meta: { env: 'dev', reqId: ctx.reqId },
        balance: BigInt(31337),
        extra: {
          info: 'additional user info',
          amount: BigInt(5678),
          points: [BigInt(100), BigInt(200), BigInt(300)],
        }
      }
    }
  },

  async getArticle(ctx, { articleId, byBN }) {
    return {
      title: `Article ${articleId}`,
      content: `Hello, this is the content for article ${articleId}. (req ${ctx.reqId})`,
      largeNum: byBN * BigInt(2),
    }
  },

  // File upload: the avatar argument arrives as a native File parsed from the
  // multipart/form-data request body.
  async uploadAvatar(_ctx, { userId, avatar }) {
    const file = avatar instanceof File ? avatar : new File([avatar], `avatar-${userId}`, { type: avatar.type })
    avatars.set(userId, file)
    return { size: file.size, name: file.name }
  },

  // File download: return a Blob or File; a File's name is sent to the client
  // via the Content-Disposition header.
  async downloadAvatar(_ctx, { userId }) {
    const avatar = avatars.get(userId)
    if (!avatar) {
      throw new WebrpcBadRouteError({ cause: `no avatar for user ${userId}` })
    }
    return avatar
  }
}

// ---------------------------------------------------------------------------
// Fastify http service
// ---------------------------------------------------------------------------
const fastify = Fastify({ logger: true })

// Pass request bodies through untouched: the generated webrpc handler parses
// JSON and multipart/form-data bodies itself from the web-standard Request.
fastify.removeAllContentTypeParsers()
fastify.addContentTypeParser('*', (_req, payload, done) => done(null, payload))

// toWebRequest adapts fastify's raw node:http request to a web-standard
// Request, streaming the request body.
const toWebRequest = (req: FastifyRequest): Request => {
  const method = (req.raw.method || 'GET').toUpperCase()
  const headers = new Headers()
  for (const [key, value] of Object.entries(req.raw.headers)) {
    if (value === undefined) continue
    if (Array.isArray(value)) {
      for (const v of value) headers.append(key, v)
    } else {
      headers.set(key, value)
    }
  }
  const init: RequestInit & { duplex?: 'half' } = { method, headers }
  if (method !== 'GET' && method !== 'HEAD') {
    init.body = Readable.toWeb(req.raw) as unknown as BodyInit
    init.duplex = 'half'
  }
  return new Request(`http://${req.raw.headers.host || 'localhost'}${req.raw.url || '/'}`, init)
}

// Health route
fastify.get('/health', async () => ({ ok: true, time: new Date().toISOString() }))

// RPC mount and entrypoint (/v1 is the schema's basepath)
fastify.all('/v1/*', async (req, reply) => {
  const ctx: RequestContext = { reqId: req.id }
  const response = await serveExampleRpc(exampleService, ctx, toWebRequest(req))

  // Take over the raw response to stream the web-standard Response back.
  reply.hijack()
  if (response === null) {
    reply.raw.writeHead(404, { 'Content-Type': 'text/plain' })
    reply.raw.end('Not Found\n')
    return
  }
  const headers: Record<string, string> = {}
  response.headers.forEach((value, key) => {
    headers[key] = value
  })
  reply.raw.writeHead(response.status, headers)
  if (response.body) {
    await pipeline(Readable.fromWeb(response.body as unknown as import('node:stream/web').ReadableStream), reply.raw)
  } else {
    reply.raw.end()
  }
})

// Start server
fastify.listen({ port: 3000 }, (err, address) => {
  if (err) {
    fastify.log.error(err)
    process.exit(1)
  }
  console.log(`Fastify server running at ${address}`)
})
