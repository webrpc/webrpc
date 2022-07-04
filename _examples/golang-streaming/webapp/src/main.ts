import * as client from './client.gen'

let fetch: client.Fetch = global.fetch


//--

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


const api = new client.ExampleService('http://127.0.0.1:4242', appFetch)

async function main() {

  console.log('[B] webrpc -- calling GetUser() rpc method (expecting User object):')

  try {
    const { user }: client.GetUserReturn = await api.getUser({ id: 966 })
    console.log('[B]', user)
    console.log('[B] welcome user ID', user.id, 'with username', user.username)
  } catch (err) {
    console.log('[B]', {err})
  }

  //
  // Get a user from the API -- different code syntax example
  //
  console.log('[D] webrpc -- calling Download() rpc streaming method:')

  // TODO: what is the client interface...?
  // check eventsource api, and check websocket api too..

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

}

main()
