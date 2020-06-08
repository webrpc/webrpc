//
// hello-webrpc-ts example program.
// --
// see the `hello-api.ridl` schema where the API is defined
// and the client.gen.ts file is code-generated which we use below..
//

// import { fetch as polyfetch } from 'whatwg-fetch'

// import * as polyfetch from 'node-fetch' // for now..

// import * as client from './client-xhr.gen'
import * as client from './client.gen'


let fetch: client.Fetch

// if (window.fetch) {
//   fetch = window.fetch.bind(window)
// } else {
//   fetch = polyfetch
// }

fetch = window.fetch.bind(window)
// fetch = polyfetch


const appFetch = (input: RequestInfo, init?: RequestInit): Promise<Response> => {
  return new Promise<Response>((resolve, reject) => {
    // before the request is made..
    init!.headers = { ...init!.headers, 'X-Custom': 'yes' }

    fetch(input, init).then(resp => {
      // after the request has been made..
      resolve(resp)
    }).catch(err => {
      // request error
      reject(err)
    })
  })
}



const api = new client.ExampleService('http://localhost:4242', appFetch)
// const api = new client.ExampleService('https://pkgrok.0xhorizon.net', appFetch)
// const api = new client.ExampleService('https://pkgrok.ngrok.io', appFetch)


// const appFetch: Fetch = async (input: RequestInfo, init?: RequestInit) => {
//   try {
//     const response = await polyfetch(input, init)

//     if (!response.ok) {
//       const json = await response.json()

//       throw new Error(json.msg)
//     }

//     return wrapResponse(response)
//   } catch (err) {
//     throw err
//   }
// }



async function main() {

  //
  // Ping the API
  //
  console.log('[A] webrpc -- calling Ping() rpc method (expecting true):')

  try {
    const resp = await api.ping()
    console.log('[A]', resp)
  } catch (err) {
    console.log('[A]', {err})
  }


  //
  // Get a user from the API
  //
  console.log('[B] webrpc -- calling GetUser() rpc method (expecting User object):')

  try {
    const { user }: client.GetUserReturn = await api.getUser({ id: 966 })
    console.log('[B]', user)
    console.log('[B] welcome user ID', user.id, 'with username', user.USERNAME)
  } catch (err) {
    console.log('[B]', {err})
  }

  //
  // Get an unknown user from the API -- we expect a 404 from the server
  //
  console.log('[C] webrpc -- calling GetUser() rpc method of an unknown user (expecting a 404):')

  try {
    const resp = await api.getUser({ id: 911 })
    console.log('[C]', resp.user)
  } catch (err) {
    console.log('[C]', {err})
  }


  //
  // Get a user from the API -- different code syntax example
  //
  console.log('[D] webrpc -- calling Download() rpc streaming method:')

  const stream = await api.download({ file: "hiithereeee" })

  console.log('yes..?')

  // hmm, could keep these separate?
  // await stream.open()
  // stream.onopen(() => {
  // })

  setTimeout(() => {
    // console.log('closing........')
    // stream.close()
  }, 3000)

  stream.ondata((data: client.DownloadReturn) => {
    // do something with data..
    console.log('YES, now we talkin.......', data)
  })
  stream.onclose((err) => {
    // err is empty if no error ..

    // connection has been closed..

    // how can we reconnect?

    // stream.open({ file: "hii" }, { headers... })
    // stream.open() // will use original input ..?
    // and fail if stream isn't closed first..
  })

  // stream.onError((err) => {
  //   // received an error, next up the connection will close..
  // })

  // await stream.close() // close the connection


}

main()
