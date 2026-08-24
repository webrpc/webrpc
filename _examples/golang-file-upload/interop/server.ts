// TypeScript implementation of the FileService schema, mirroring the Go
// server in ../main.go. Used by test.ts to prove both sides of the file
// upload/download wire contract.
import http from 'node:http'
import { Readable } from 'node:stream'
import { pipeline } from 'node:stream/promises'
import { FileServiceServer, serveFileServiceRpc, WebrpcBadRequestError, WebrpcBadRouteError } from './server.gen'

export const createFileService = (): FileServiceServer => {
  const avatars = new Map<number, File>()

  return {
    // File arguments arrive as native File objects parsed from the
    // multipart/form-data request body.
    async uploadAvatar(_ctx, { userId, avatar }) {
      if (avatar.size === 0) {
        throw new WebrpcBadRequestError({ cause: 'empty avatar file' })
      }
      const file = avatar instanceof File ? avatar : new File([avatar], 'avatar', { type: avatar.type })
      avatars.set(userId, file)
      return { size: file.size }
    },

    async uploadAttachments(_ctx, { attachments }) {
      return { count: attachments.length }
    },

    // Download methods return the file to be streamed as the raw response
    // body; a File carries the filename for the Content-Disposition header.
    async downloadAvatar(_ctx, { userId }) {
      const avatar = avatars.get(userId)
      if (!avatar) {
        throw new WebrpcBadRouteError({ cause: `no avatar for user ${userId}` })
      }
      return avatar
    },

    async stampAvatar(_ctx, { userId, stamp }) {
      const avatar = avatars.get(userId)
      if (!avatar) {
        throw new WebrpcBadRouteError({ cause: `no avatar for user ${userId}` })
      }
      return new File([avatar, stamp], `stamped-${avatar.name}`, { type: avatar.type })
    },

    async listAvatars() {
      return { userIds: [...avatars.keys()].sort((a, b) => a - b) }
    },
  }
}

// createFileServiceHttpServer hosts the webrpc service on a plain node:http
// server, adapting IncomingMessage/ServerResponse to the web-standard
// Request/Response pair the generated handler works with.
export const createFileServiceHttpServer = (): http.Server => {
  const fileService = createFileService()

  return http.createServer(async (req, res) => {
    try {
      const request = toWebRequest(req)
      const response = await serveFileServiceRpc(fileService, null, request)
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
