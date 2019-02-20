const polyfetch = require('node-fetch')

async function main() {
  const {
    ExampleService,
    GetUserRequest
  } = await import('./client.gen.mjs')

  const svc = new ExampleService('http://127.0.0.1:4242', polyfetch)

  // Expecting "true"
  console.log('[A] webrpc -- calling Ping() rpc method (expecting true):')

  svc.Ping().then(resp => {
    console.log('[A]', {resp})
  }).catch((err) => {
    console.log('[A]', {err})
  })

  // Expecting an error
  console.log('[B] webrpc -- calling GetUser() rpc method of an unknown user (expecting a 404):')

  svc.GetUser({ req: new GetUserRequest({userID: 911}) }).then(resp => {
    console.log('[B]', {resp})
  }).catch((err, res) => {
    console.log('[B]', {err})
  })

  // Expecting some user data
  console.log('[C] webrpc -- calling GetUser() rpc method (expecting User object):')

  svc.GetUser({ req: new GetUserRequest({userID: 966}) }).then(resp => {
    console.log('[C]', {resp})
  }).catch((err, res) => {
    console.log('[C]', {err})
  })

}

main()
