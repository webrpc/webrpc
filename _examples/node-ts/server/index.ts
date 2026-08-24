import http, { IncomingMessage, ServerResponse } from 'node:http'
import { HttpHandler, createHttpEntrypoint, createWebrpcServerHandler, RequestContext, composeHttpHandler, sendJson } from './helpers'
import { Kind, ExampleServer, serveExampleRpc, WebrpcEndpointError, WebrpcBadRouteError } from './server.gen'
import { withLogging, withTrace, withCors } from './middleware'

// In-memory storage of the uploaded avatars.
const avatars = new Map<number, File>()

// ExampleServer RPC implementation of the webrpc service definition
const exampleService: ExampleServer<RequestContext> = {
  async ping() {
    return {}
  },

  async getUser(ctx, { userId }) {
    const traceId = ctx.get<string>('traceId') || ''

		if (userId === 911) {
			throw new WebrpcEndpointError({ cause: 'User 911 is forbidden' })
		}

    return {
      code: 200,
      user: {
        id: userId,
        USERNAME: `user-${userId}`,
        role: Kind.USER,
        meta: { env: 'dev', reqId: ctx.reqId, traceId },
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
      content: `Hello, this is the content for article ${articleId}. (req ${ctx.reqId})`,
      largeNum: byBN * BigInt(2),
    }
  },

  // File upload: the avatar argument arrives as a native File parsed from the
  // multipart/form-data request body.
  async uploadAvatar(ctx, { userId, avatar }) {
    const file = avatar instanceof File ? avatar : new File([avatar], `avatar-${userId}`, { type: avatar.type })
    avatars.set(userId, file)
    console.log(`stored avatar ${file.name} (${file.type}, ${file.size} bytes) for user ${userId} (req ${ctx.reqId})`)
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

// NOTE: here is another demo of how you can implement the ExampleServer RPC interface using a class.
// --
// class ExampleService implements ExampleServer<RequestContext> {
//   async ping(): Promise<{}> {
//     return {}
//   }
//   async getUser(ctx: RequestContext, req: GetUserRequest): Promise<GetUserResponse> {
//     const traceId = ctx.get<string>('traceId') || ''
//     return {
//       code: 200,
//       user: {
//         id: args.userId,
//         USERNAME: `user-${args.userId}`,
//         role: Kind.USER,
//         meta: { env: 'dev', reqId: ctx.reqId, traceId },
//       }
//     }
//   }
//   async getArticle(ctx: RequestContext, req: GetArticleRequest): Promise<GetArticleResponse> {
//     return {
//       title: `Article ${args.articleId}`,
//       content: `Hello, this is the content for article ${args.articleId}. (req ${ctx.reqId})`
//     }
//   }
// }

// Main routes of the service
const routes = (): HttpHandler  => {
  // const exampleService = new ExampleService()
  const rpcHandler = createWebrpcServerHandler(exampleService, serveExampleRpc)

  // Return the actual request handler (async because we use await inside)
  return async (ctx: RequestContext, req: IncomingMessage, res: ServerResponse): Promise<void> => {
    const url = req.url

    // First try RPC routing (/v1/*, the schema's basepath)
    if (url?.startsWith('/v1/')) {
      await rpcHandler(ctx, req, res)
      return
    }

    // Other routes
    switch (url) {
      case "/": {
        res.writeHead(200, { "Content-Type": "text/plain" })
        res.end(`Hello world (req ${ctx.reqId})\n`)
        return
      }
      case "/json": {
        sendJson(res, 200, { ok: true, time: new Date().toISOString(), reqId: ctx.reqId })
        return
      }
      default: {
        res.writeHead(404, { "Content-Type": "text/plain" })
        res.end("Not Found\n")
        return
      }
    }
  }
}

// Compose middleware chain and primary routes entrypoint handler
const handler = composeHttpHandler([withLogging, withTrace, withCors], routes())

// Node http server bootstrap (ensure handler is passed)
http.createServer(createHttpEntrypoint(handler)).listen(3000, () => {
  console.log("Server running at http://localhost:3000/")
})
