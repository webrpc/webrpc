import http, { IncomingMessage, ServerResponse } from 'node:http'
import { HttpHandler, createHttpEntrypoint, createWebrpcServerHandler, RequestContext, composeHttpHandler, sendJson } from './helpers'
import { Kind, ExampleServer, serveExampleRpc, WebrpcEndpointError } from './server.gen'
import { withLogging, withTrace, withCors } from './middleware'

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

    // First try RPC routing (/rpc/*)
    if (url?.startsWith('/rpc/')) {
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
