import type { IncomingMessage, ServerResponse } from 'node:http'
import { randomUUID } from 'node:crypto'
import { WebrpcHeader, WebrpcHeaderValue } from './server.gen'

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

// Function that, given a service + ctx + url + body, either handles the RPC and returns a result, or null if pattern mismatch.
export type ServeWebrpcFn<S, C extends RequestContext> = (service: S, ctx: C, urlPath: string, body: any) => Promise<WebrpcResult | null>

export interface WebrpcResult {
  method: string
  status: number
  headers: Record<string, string>
  body: any
}

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
    const url = req.url || ''
    if (!url.startsWith('/rpc/')) return // not our RPC route; caller will continue

    // Accept both GET & POST/PUT/PATCH for simplicity (GET => empty object)
    const methodVerb = (req.method || 'GET').toUpperCase()
    let rawBody = ''
    if (methodVerb === 'POST') {
      rawBody = await new Promise<string>((resolve, reject) => {
        let data = ''
        const onData = (chunk: Buffer) => { data += chunk.toString('utf8') }
        const onEnd = () => { cleanup(); resolve(data) }
        const onError = (err: Error) => { cleanup(); reject(err) }
        const onAborted = () => { cleanup(); reject(new Error('request aborted')) }

        const cleanup = () => {
          req.off('data', onData)
          req.off('end', onEnd)
          req.off('error', onError)
          req.off('aborted', onAborted)
        }

        req.on('data', onData)
        req.on('end', onEnd)
        req.on('error', onError)
        req.on('aborted', onAborted)

        // If already aborted before listeners attached
        if ((req as any).aborted || ctx.abort.aborted) {
          cleanup()
          reject(new Error('request aborted'))
        }
      })
    }

    let parsed: any = {}
    if (rawBody.length > 0) {
      try {
        parsed = JSON.parse(rawBody)
      } catch (e: any) {
        const status = 400
        const body = { msg: 'invalid JSON body', status, code: '' }
        res.writeHead(status, { [WebrpcHeader]: WebrpcHeaderValue, 'Content-Type': 'application/json' })
        res.end(JSON.stringify(body))
        return
      }
    }

    const result = await serveRpc(service, ctx, url, parsed)
    if (result == null) return // pattern mismatch (shouldn't happen due to prefix check)

    const payload = JSON.stringify(result.body ?? {})
    res.writeHead(result.status, {
      ...result.headers,
      'Content-Length': Buffer.byteLength(payload)
    })
    res.end(payload)
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
