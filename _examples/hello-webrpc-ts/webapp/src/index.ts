import { fetch as polyfetch } from 'whatwg-fetch'
import * as client from './client.gen'

const svc = new client.ExampleService('http://127.0.01:4242', polyfetch)

// Expecting "true"
console.log('[A] webrpc -- calling Ping() rpc method (expecting true):')

svc.Ping().then(res => {
  console.log('[A]', {res})
}).catch(err => {
  console.log('[A]', {err})
})

// Expecting an error
console.log('[B] webrpc -- calling GetUser() rpc method of an unknown user (expecting a 404):')

svc.GetUser(new client.GetUserRequest({userID: 911})).then((user) => {
  console.log('[B]', {user})
}).catch(err => {
  console.log('[B]', {err})
})

// Expecting some user data
console.log('[C] webrpc -- calling GetUser() rpc method (expecting User object):')

svc.GetUser(new client.GetUserRequest({userID: 966})).then((user) => {
  console.log('[C]', {user})
  console.log('[C] welcome user ID', user.ID, 'with username', user.Username)
}).catch(err => {
  console.log('[C]', {err})
})
