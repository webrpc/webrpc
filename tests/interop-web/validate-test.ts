// Request validation of the generated TypeScript server (validate.ridl):
// valid requests reach the service, invalid ones are rejected with HTTP 400.
import assert from 'node:assert/strict'
import { serveValidateRpc } from './validate.gen'

const check = async (body: object): Promise<number> => {
  const request = new Request('http://localhost/rpc/Validate/Check', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  })
  const response = await serveValidateRpc({ check: async () => ({ ok: true }) }, null, request)
  return response!.status
}

const valid = {
  kinds: ['USER', 'ADMIN'],
  tiers: ['free', 'pro'],
  matrix: [['a', 'b'], []],
  maps: [{ a: 'b' }, {}],
  items: [{ id: 1 }],
  anything: { a: [1, 'b'] },
}

const main = async () => {
  assert.equal(await check(valid), 200, 'valid request')

  const invalid: { [name: string]: object } = {
    'int enum by field name, not its + json wire value': { ...valid, kinds: ['Admin'] },
    'string enum by field name, not its value': { ...valid, tiers: ['Pro'] },
    'number in a nested string array': { ...valid, matrix: [['a', 1]] },
    'string in an array of maps': { ...valid, maps: ['a'] },
    'invalid struct in an array': { ...valid, items: [{ id: 'one' }] },
  }
  for (const [name, body] of Object.entries(invalid)) {
    assert.equal(await check(body), 400, name)
  }

  console.log('All request validation tests passed.')
}

main().catch((err) => {
  console.error(err)
  process.exit(1)
})
