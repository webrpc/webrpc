let example = require('./example.gen.js')

const fetch = require('node-fetch');

let svc = new example.ExampleService('http://127.0.01:4242', fetch)

// Expecting "true"
svc.Ping().then((res) => {
  console.log({res})
}).catch(() => {
  console.log({err})
})

// Expecting an error
svc.GetUser(new example.GetUserRequest({userID: 911})).then((user) => {
  console.log({user})
}).catch((err, res) => {
  console.log({err})
})

// Expecting some user data
svc.GetUser(new example.GetUserRequest({userID: 966})).then((user) => {
  console.log({user})
}).catch((err, res) => {
  console.log({err})
})
