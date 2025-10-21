import { HttpHandler } from './helpers'
import { randomUUID } from 'node:crypto'

// Middleware for logging requests & responses
export const withLogging = (next: HttpHandler): HttpHandler => {
  return async (ctx, req, res) => {
    const traceId = ctx.reqId
    console.log(`[REQ ${traceId}] ${req.method} ${req.url}`)
    try {
      await next(ctx, req, res)
    } finally {
      const end = performance.now()
      const durationMs = end - ctx.start
      console.log(`[RES ${traceId}] ${req.method} ${req.url} -> ${res.statusCode} (${durationMs.toFixed(3)}ms)`)
    }
  }
}

// Middleware for adding a trace id header
export const withTrace = (next: HttpHandler): HttpHandler => {
  return async (ctx, req, res) => {
    ctx.set('traceId', randomUUID())
    res.setHeader('X-Trace-Id', ctx.reqId)
    await next(ctx, req, res)
  }
}

// Middleware for permissive CORS handling (allow all origins)
// - Sets standard Access-Control-* headers
// - Responds immediately to OPTIONS preflight with 204 No Content
// - Leaves actual request handling to downstream middleware/handler
//
// NOTE: make sure to adjust for your security needs! This is just a demo
// cors middleware.
export const withCors = (next: HttpHandler): HttpHandler => {
  return async (ctx, req, res) => {
    // Apply headers early (unless already sent)
    if (!res.headersSent) {
      res.setHeader('Access-Control-Allow-Origin', '*')
      res.setHeader('Access-Control-Allow-Methods', 'GET,POST,PUT,PATCH,DELETE,OPTIONS')
      // If the client asked for specific headers, echo them; otherwise provide a reasonable default list
      const reqHeaders = req.headers['access-control-request-headers']
      res.setHeader('Access-Control-Allow-Headers', reqHeaders || 'Content-Type, Authorization')
      res.setHeader('Access-Control-Max-Age', '86400') // 24h caching of preflight
    }

    if (req.method === 'OPTIONS') {
      // Preflight: finalize quickly
      if (!res.headersSent) {
        res.writeHead(204)
      }
      res.end()
      return
    }

    await next(ctx, req, res)
  }
}
