// End-to-end interoperability tests for the webrpc `file` core type.
//
// Runs the same client test suite (plain JSON method, single file upload,
// repeated []file upload, file download, combined upload + download, and the
// JSON error envelope) in three directions:
//
//   1. TypeScript client -> Go server        (go run ..)
//   2. TypeScript client -> TypeScript server (server.ts, in-process)
//   3. Go client         -> TypeScript server (go run .. client <url>)
//
// Usage: pnpm test
import assert from 'node:assert/strict'
import { spawn, ChildProcess } from 'node:child_process'
import { randomBytes } from 'node:crypto'
import { once } from 'node:events'
import { FileService, WebrpcFile, WebrpcError } from './client.gen'
import { createFileServiceHttpServer } from './server'

const GO_SERVER_PORT = 14242

const fileBytes = async (file: WebrpcFile): Promise<Uint8Array> => {
  return new Uint8Array(await (await file.blob()).arrayBuffer())
}

const concat = (a: Uint8Array, b: Uint8Array): Uint8Array => {
  const out = new Uint8Array(a.length + b.length)
  out.set(a, 0)
  out.set(b, a.length)
  return out
}

// clientSuite runs the full TypeScript client test suite against a running
// FileService server (Go or TypeScript) at the given address.
export const clientSuite = async (addr: string) => {
  const client = new FileService(addr, fetch)
  const userId = 1234

  // Plain JSON method on a service that also transfers files.
  const before = await client.listAvatars()
  assert.deepEqual(before.userIds, [], 'expected no avatars yet')

  // Single file upload alongside a scalar argument.
  const avatarBytes = new Uint8Array(randomBytes(1024))
  const avatar = new File([avatarBytes], 'me.png', { type: 'image/png' })
  const { size } = await client.uploadAvatar({ userId, avatar })
  assert.equal(size, avatarBytes.length, 'uploaded avatar size mismatch')

  // Repeated []file upload.
  const { count } = await client.uploadAttachments({
    userId,
    attachments: [
      new File(['attachment a'], 'a.txt', { type: 'text/plain' }),
      new File(['attachment b'], 'b.txt', { type: 'text/plain' }),
      new Blob(['attachment c (a bare Blob)'], { type: 'text/plain' }),
    ],
  })
  assert.equal(count, 3, 'expected 3 attachments')

  // File download: metadata from the headers, bytes compared to the upload.
  const download = await client.downloadAvatar({ userId })
  assert.equal(download.name, 'me.png', 'downloaded avatar name mismatch')
  assert.equal(download.contentType, 'image/png', 'downloaded avatar content type mismatch')
  assert.equal(download.size, avatarBytes.length, 'downloaded avatar size mismatch')
  assert.deepEqual(await fileBytes(download), avatarBytes, 'downloaded avatar bytes mismatch')

  // Combined upload + download in one method.
  const stampBytes = new Uint8Array(randomBytes(256))
  const stamped = await client.stampAvatar({
    userId,
    stamp: new File([stampBytes], 'stamp.bin', { type: 'application/octet-stream' }),
  })
  assert.equal(stamped.name, 'stamped-me.png', 'stamped avatar name mismatch')
  assert.equal(stamped.contentType, 'image/png', 'stamped avatar content type mismatch')
  assert.deepEqual(await fileBytes(stamped), concat(avatarBytes, stampBytes), 'stamped avatar bytes mismatch')

  const after = await client.listAvatars()
  assert.deepEqual(after.userIds, [userId], 'expected one avatar')

  // Errors still arrive as the standard webrpc JSON error envelope, even on
  // file download methods.
  await assert.rejects(
    client.downloadAvatar({ userId: 999999 }),
    (err: unknown) => err instanceof WebrpcError && err.code === -2, // WebrpcBadRoute
    'expected WebrpcBadRoute error for a missing avatar',
  )
}

const waitForServer = async (addr: string, timeoutMs = 30_000) => {
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

const stop = async (child: ChildProcess) => {
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

const run = (name: string, cmd: string, args: string[], env: Record<string, string> = {}): ChildProcess => {
  const child = spawn(cmd, args, {
    cwd: '..',
    env: { ...process.env, ...env },
    stdio: ['ignore', 'inherit', 'inherit'],
    detached: true,
  })
  child.on('error', (err) => {
    console.error(`failed to start ${name}: ${err.message}`)
    process.exit(1)
  })
  return child
}

const main = async () => {
  // 1. TypeScript client vs the Go example server.
  console.log('--- interop 1/3: TypeScript client -> Go server')
  const goServer = run('go server', 'go', ['run', '.'], { PORT: String(GO_SERVER_PORT) })
  try {
    await waitForServer(`http://localhost:${GO_SERVER_PORT}`)
    await clientSuite(`http://localhost:${GO_SERVER_PORT}`)
  } finally {
    await stop(goServer)
  }
  console.log('OK')

  // 2. TypeScript client vs the TypeScript server.
  console.log('--- interop 2/3: TypeScript client -> TypeScript server')
  const tsServer = createFileServiceHttpServer()
  await new Promise<void>((resolve) => tsServer.listen(0, resolve))
  const tsAddress = tsServer.address()
  if (tsAddress === null || typeof tsAddress === 'string') {
    throw new Error('unexpected TypeScript server address')
  }
  const tsServerUrl = `http://localhost:${tsAddress.port}`
  await clientSuite(tsServerUrl)
  console.log('OK')

  // 3. Go client vs the TypeScript server (fresh service state).
  console.log('--- interop 3/3: Go client -> TypeScript server')
  const tsServer2 = createFileServiceHttpServer()
  await new Promise<void>((resolve) => tsServer2.listen(0, resolve))
  const tsAddress2 = tsServer2.address()
  if (tsAddress2 === null || typeof tsAddress2 === 'string') {
    throw new Error('unexpected TypeScript server address')
  }
  const goClient = run('go client', 'go', ['run', '.', 'client', `http://localhost:${tsAddress2.port}`])
  const [exitCode] = (await once(goClient, 'exit')) as [number | null]
  assert.equal(exitCode, 0, 'go client test suite failed')
  console.log('OK')

  tsServer.close()
  tsServer2.close()
  console.log('all interop tests passed')
}

main().catch((err) => {
  console.error(err)
  process.exit(1)
})
