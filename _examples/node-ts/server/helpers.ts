import type { IncomingMessage, ServerResponse } from 'node:http'
import { Readable } from 'node:stream'
import { pipeline } from 'node:stream/promises'
import { randomUUID } from 'node:crypto'

export type HttpHandler<C = RequestContext> = (ctx: C, req: IncomingMessage, res: ServerResponse) => Promise<void>

export interface RequestContext {
  // request id
  reqId: string

  // timestamp when request started
  start: number

  // AbortSignal that fires if client disconnects or server cancels work
  abort: AbortSignal

  // Arbitrary key/value bag for middleware & handlers
  data: Map<string, unknown>
  set<T = unknown>(key: string, value: T): void
  get<T = unknown>(key: string): T | undefined

  // internal abort controller instance (non-enumerable at runtime)
  _controller?: AbortController

  // raw HTTP request & response objects
  httpRequest: IncomingMessage
  httpResponse: ServerResponse
}

export const createRequestContext = (req: IncomingMessage, res: ServerResponse): RequestContext => {
  const start = performance.now()
  const controller = new AbortController()
  const ctx: RequestContext = {
    reqId: randomUUID(),
    start,
    abort: controller.signal,
    data: new Map(),
    set(key, value) { this.data.set(key, value) },
    get(key) { return this.data.get(key) as any },
    httpRequest: req,
    httpResponse: res,
  }

  // Make _controller non-enumerable to keep logs clean
  Object.defineProperty(ctx, '_controller', { value: controller, enumerable: false, writable: false })

  // Make httpRequest & httpResponse non-enumerable to keep logs clean
  Object.defineProperty(ctx, 'httpRequest', { value: req, enumerable: false, writable: false })
  Object.defineProperty(ctx, 'httpResponse', { value: res, enumerable: false, writable: false })

  return ctx
}

// Function that, given a service + ctx + web-standard Request, either handles
// the RPC and returns a web-standard Response, or null on pattern mismatch.
export type ServeWebrpcFn<S, C extends RequestContext> = (service: S, ctx: C, request: Request) => Promise<Response | null>

export const createHttpEntrypoint = (handler: HttpHandler) => {
  return async (req: IncomingMessage, res: ServerResponse) => {
    const ctx = createRequestContext(req, res)

    const abort = () => {
      const controller: AbortController | undefined = (ctx as any)._controller
      if (controller && !controller.signal.aborted) controller.abort()
    }
    req.once('aborted', abort)
    res.once('close', abort)

    try {
      await handler(ctx, req, res)
    } catch (err: any) {
      console.error(`[ERR ${ctx.reqId}]`, err?.message || err)
      if (!res.headersSent) {
        res.writeHead(500, { 'Content-Type': 'application/json' })
      }
      if (!res.writableEnded) {
        const body = ctx.abort.aborted
          ? { msg: 'client closed request', reqId: ctx.reqId }
          : { msg: 'internal error', reqId: ctx.reqId }
        res.end(JSON.stringify(body))
      }
    }
  }
}

export const createWebrpcServerHandler = <S, C extends RequestContext>(service: S, serveRpc: ServeWebrpcFn<S, C>): HttpHandler<C> => {
  return async (ctx: C, req: IncomingMessage, res: ServerResponse): Promise<void> => {
    // Adapt node:http to the web-standard Request the generated webrpc
    // handler consumes. The generated handler reads and parses the body
    // itself (JSON, or multipart/form-data for file upload methods).
    const request = toWebRequest(req, ctx.abort)

    const response = await serveRpc(service, ctx, request)
    if (response == null) {
      res.writeHead(404, { 'Content-Type': 'text/plain' })
      res.end('Not Found\n')
      return
    }

    await sendWebResponse(res, response)
  }
}

// toWebRequest adapts a node:http IncomingMessage to a web-standard Request,
// streaming the request body.
export const toWebRequest = (req: IncomingMessage, signal?: AbortSignal): Request => {
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
  const init: RequestInit & { duplex?: 'half' } = { method, headers, signal: signal ?? null }
  if (method !== 'GET' && method !== 'HEAD') {
    init.body = Readable.toWeb(req) as unknown as BodyInit
    init.duplex = 'half'
  }
  return new Request(`http://${req.headers.host || 'localhost'}${req.url || '/'}`, init)
}

// sendWebResponse writes a web-standard Response to a node:http
// ServerResponse, streaming the response body.
export const sendWebResponse = async (res: ServerResponse, response: Response): Promise<void> => {
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

// Generic middleware composer. Applies middleware array in order so that
// composeHttpHandler([a, b], h) => a(b(h)).
export const composeHttpHandler = <C = RequestContext>(
  middleware: Array<(next: HttpHandler<C>) => HttpHandler<C>>,
  handler: HttpHandler<C>
): HttpHandler<C> => {
  return middleware.reduceRight<HttpHandler<C>>((acc, mw) => mw(acc), handler)
}

// Simple JSON helper (typed) – narrows headers & body
export const sendJson = (res: ServerResponse, status: number, body: unknown): void => {
  const payload = JSON.stringify(body)
  res.writeHead(status, {
    "Content-Type": "application/json",
    "Content-Length": Buffer.byteLength(payload)
  })
  res.end(payload)
}
