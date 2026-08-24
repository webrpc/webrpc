// End-to-end test suite for the generated plain-JavaScript webrpc client
// (client.gen.mjs) against a running FileService server (Go or TypeScript):
// plain JSON method, single file upload, repeated []file upload, file
// download with metadata + byte-compare, combined upload + download, and the
// JSON error envelope.
//
// Runs on stock Node (>= 20, for the global fetch/FormData/Blob/File) with no
// transpile step, which also proves the generated client is valid plain JS.
//
// Usage: node client_suite.mjs <server url>
import assert from 'node:assert/strict'
import { randomBytes } from 'node:crypto'
import { FileService } from './client.gen.mjs'

const fileBytes = async (file) => {
  return new Uint8Array(await (await file.blob()).arrayBuffer())
}

const concat = (a, b) => {
  const out = new Uint8Array(a.length + b.length)
  out.set(a, 0)
  out.set(b, a.length)
  return out
}

const clientSuite = async (addr) => {
  const client = new FileService(addr, fetch)
  // A userId distinct from the TypeScript suite's, plus relative listAvatars
  // assertions, so both suites can run against the same server instance.
  const userId = 4321

  // Plain JSON method on a service that also transfers files.
  const before = await client.listAvatars()
  assert.ok(!before.userIds.includes(userId), 'expected no avatar for this user yet')

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
  assert.ok(after.userIds.includes(userId), 'expected an avatar for this user')

  // Errors still arrive as the standard webrpc JSON error envelope, even on
  // file download methods; the JavaScript client throws the decoded envelope.
  await assert.rejects(
    client.downloadAvatar({ userId: 999999 }),
    (err) => err.code === -2 && typeof err.msg === 'string', // WebrpcBadRoute
    'expected WebrpcBadRoute error envelope for a missing avatar',
  )
}

const main = async () => {
  const addr = process.argv[2]
  if (!addr) {
    console.error('usage: node client_suite.mjs <server url>')
    process.exit(1)
  }
  await clientSuite(addr)
}

main().catch((err) => {
  console.error(err)
  process.exit(1)
})
