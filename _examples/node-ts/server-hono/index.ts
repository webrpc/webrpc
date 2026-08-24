import { Hono } from 'hono'
import type { Context } from 'hono'
import { serve } from '@hono/node-server'
import { logger } from 'hono/logger'
import { cors } from 'hono/cors'
import { ExampleServer, serveExampleRpc, Kind, WebrpcEndpointError, WebrpcBadRouteError } from './server.gen'
import { randomUUID } from 'node:crypto'


// ---------------------------------------------------------------------------
// Example rpc service
// ---------------------------------------------------------------------------
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
        meta: { env: 'dev', reqId: ctx.var.traceId },
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
    console.log('getArticle byBN:', byBN)
    return {
      title: `Article ${articleId}`,
      content: `Hello, this is the content for article ${articleId}. (req ${ctx.var.traceId})`,
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
// Http service with middleware and request context
// ---------------------------------------------------------------------------
// Define typed variable bag for Hono Context via Env generics.
// Env = { Variables: { ... }, Bindings?: ..., BasePath?: string }
interface ServiceEnv {
  Variables: {
    traceId: string
    pingedAt?: string
  }
}

type RequestContext = Context<ServiceEnv>

// Hono app setup with typed env
const app = new Hono<ServiceEnv>()

// Built-in concise request logger (method/path/status/duration)
app.use('*', logger())

// Tracing + request id enrichment middleware
app.use('*', async (c: RequestContext, next: () => Promise<void>) => {
  const traceId = randomUUID()
  c.set('traceId', traceId)
  await next()
  c.res.headers.set('X-Trace-Id', traceId)
})

// CORS middleware (simple manual implementation). Adjust origins as needed.
app.use('*', cors({
  origin: (origin) => origin ?? '*',
  allowMethods: ['GET','POST','PUT','PATCH','DELETE','OPTIONS'],
  allowHeaders: ['Content-Type','Authorization','Webrpc'],
  exposeHeaders: [],
  credentials: true,
  maxAge: 86400
}))

// Root route
app.get('/', (c: RequestContext) => c.text(`Hello world (req ${c.var.traceId})`))

// Health route
app.get('/health', (c: RequestContext) => c.json({ ok: true, time: new Date().toISOString(), traceId: c.var.traceId }))

// RPC mount and entrypoint. Hono handlers already speak web-standard
// Request/Response, so the generated webrpc handler plugs in directly.
app.all('/v1/*', async (ctx: RequestContext) => {
  const response = await serveExampleRpc(exampleService, ctx, ctx.req.raw)
  return response ?? ctx.notFound()
})

// Start server
serve({ fetch: app.fetch, port: 3000 }, (info: { port: number }) => {
  console.log(`Hono server running at http://localhost:${info.port}`)
})
