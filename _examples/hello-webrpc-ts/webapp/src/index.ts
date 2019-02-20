import { fetch as polyfetch } from 'whatwg-fetch'
import * as client from './client.gen'

const svc = new client.ExampleService('http://127.0.01:4242', polyfetch)

// Expecting "true"
console.log('[A] webrpc -- calling Ping() rpc method (expecting true):')

svc.ping().then(resp => {
  console.log('[A]', resp.status)
}).catch(err => {
  console.log('[A]', {err})
})

// Expecting an error
console.log('[B] webrpc -- calling GetUser() rpc method of an unknown user (expecting a 404):')

svc.getUser({userID: 911}).then(resp => {
  console.log('[B]', resp.user)
}).catch((err: client.WebRPCError) => {
  console.log('[B]', {err})
})

// Expecting some user data, using object destructuring syntax for response
console.log('[C] webrpc -- calling GetUser() rpc method (expecting User object):')

svc.getUser({userID: 966}).then(({ user }: client.IExampleServiceGetUserReturn) => {
  console.log('[C]', user)
  console.log('[C] welcome user ID', user.id, 'with username', user.USERNAME)
}).catch(err => {
  console.log('[C]', {err})
})

// Expecting some user data - again, with differen response syntax
console.log('[D] webrpc -- calling GetUser() rpc method (expecting User object):')

svc.getUser({userID: 1337}).then(resp => {
  const user = resp.user
  console.log('[D]', user)
  console.log('[D] welcome user ID', user.id, 'with username', user.USERNAME)
}).catch(err => {
  console.log('[D]', {err})
})
