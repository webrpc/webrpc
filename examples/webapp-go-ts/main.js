let out = require('./example.gen.js')

const fetch = require('node-fetch');

let svc = new out.ExampleService('http://127.0.01:4242', fetch)

svc.Ping({}).then((res) => {
  console.log({res})
}).catch(() => {
  console.log({err})
})
