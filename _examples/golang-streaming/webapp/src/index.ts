//
// hello-webrpc-ts example program.
// --
// see the `hello-api.ridl` schema where the API is defined
// and the client.gen.ts file is code-generated which we use below..
//

import * as client from './client.gen'

// const appFetch = (input: RequestInfo, init?: RequestInit): Promise<Response> => {
//   return new Promise<Response>((resolve, reject) => {
//     // before the request is made..

//     if (!init) {
//       init = {}
//     }

//     init.headers = { ...init!.headers, 'X-Custom': 'yes' }
//     init.mode = 'cors'

//     window.fetch(input, init).then(resp => {
//       // after the request has been made..
//       resolve(resp)
//     }).catch(err => {
//       // request error
//       reject(err)
//     })
//   })
// }



const api = new client.ExampleService('http://localhost:4242', window.fetch)
// const api = new client.ExampleService('https://ba48911374e9.ngrok.io', window.fetch)

api.rpcDefaultHeaders = {
  'X-Custom': 'v123'
}


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
    console.log('[B] welcome user ID', user.id, 'with username', user.username)
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

  // we can add .ondata(data, options)
  // where options we can do autoRetry: true .. default..

  stream.ondata((data: client.DownloadReturn) => {
    // do something with data..
    console.log('YES, now we talkin.......', data)
  })

  // TODO: lets check what websocket looks like..
  // maybe we have .onerror and .onclose ? 

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
