import * as client from './client.gen'
import * as polyfetch from 'node-fetch'

const api = new client.ExampleService('http://127.0.01:4242', polyfetch)

console.log('trying')
