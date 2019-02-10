const polyfetch = require('node-fetch')

async function main() {
  const {
    ExampleService,
    GetUserRequest
  } = await import('./client.gen.mjs')

  const svc = new ExampleService('http://127.0.01:4242', polyfetch)

  // Expecting "true"
  console.log('[A] webrpc -- calling Ping() rpc method (expecting true):')

  svc.Ping().then((res) => {
    console.log('[A]', {res})
  }).catch((err) => {
    console.log('[A]', {err})
  })

  // Expecting an error
  console.log('[B] webrpc -- calling GetUser() rpc method of an unknown user (expecting a 404):')

  svc.GetUser(new GetUserRequest({userID: 911})).then((user) => {
    console.log('[B]', {user})
  }).catch((err, res) => {
    console.log('[B]', {err})
  })

  // Expecting some user data
  console.log('[C] webrpc -- calling GetUser() rpc method (expecting User object):')

  svc.GetUser(new GetUserRequest({userID: 966})).then((user) => {
    console.log('[C]', {user})
  }).catch((err, res) => {
    console.log('[C]', {err})
  })

}

main()
