// hello-webrpc v1.0.0 1769ce5a249c8ed4e4dab8320a7d67779eae0664
// --
// Code generated by webrpc-gen with javascript generator. DO NOT EDIT.
//
// webrpc-gen -schema=hello-api.ridl -target=javascript -exports=false -client -out=./webapp/client.gen.js

// WebRPC description and code-gen version
const WebRPCVersion = "v1"

// Schema version of your RIDL schema
const WebRPCSchemaVersion = "v1.0.0"

// Schema hash generated from your RIDL schema
const WebRPCSchemaHash = "1769ce5a249c8ed4e4dab8320a7d67779eae0664"

//
// Types
//

var Kind;
(function (Kind) {
  Kind["USER"] = "USER"
  Kind["ADMIN"] = "ADMIN"
})(Kind || (Kind = {}))

class Empty {
  constructor(_data) {
    this._data = {}
    if (_data) {
    }
  }
  
  toJSON() {
    return this._data
  }
}

class User {
  constructor(_data) {
    this._data = {}
    if (_data) {
      this._data['id'] = _data['id']
      this._data['USERNAME'] = _data['USERNAME']
      this._data['created_at'] = _data['created_at']
      this._data['updatedAt'] = _data['updatedAt']
      this._data['deletedAt'] = _data['deletedAt']
    }
  }
  
  get id() {
    return this._data['id']
  }
  set id(value) {
    this._data['id'] = value
  }
  
  get USERNAME() {
    return this._data['USERNAME']
  }
  set USERNAME(value) {
    this._data['USERNAME'] = value
  }
  
  get created_at() {
    return this._data['created_at']
  }
  set created_at(value) {
    this._data['created_at'] = value
  }
  
  get updatedAt() {
    return this._data['updatedAt']
  }
  set updatedAt(value) {
    this._data['updatedAt'] = value
  }
  
  get deletedAt() {
    return this._data['deletedAt']
  }
  set deletedAt(value) {
    this._data['deletedAt'] = value
  }
  
  toJSON() {
    return this._data
  }
}

  
//
// Client
//

class ExampleService {
  constructor(hostname, fetch) {
    this.path = '/rpc/ExampleService/'
    this.hostname = hostname
    this.fetch = (input, init) => fetch(input, init)
  }

  url(name) {
    return this.hostname + this.path + name
  }
  
  ping = (headers) => {
    return this.fetch(
      this.url('Ping'),
      createHTTPRequest({}, headers)
    ).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          status: (_data.status)
        }
      })
    })
  }
  
  getUser = (args, headers) => {
    return this.fetch(
      this.url('GetUser'),
      createHTTPRequest(args, headers)
    ).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          user: new User(_data.user)
        }
      })
    })
  }
  
}

  
const createHTTPRequest = (body = {}, headers = {}) => {
  return {
    method: 'POST',
    headers: { ...headers, 'Content-Type': 'application/json' },
    body: JSON.stringify(body || {})
  }
}

const buildResponse = (res) => {
  return res.text().then(text => {
    let data
    try {
      data = JSON.parse(text)
    } catch(err) {
      throw { code: 'unknown', msg: `expecting JSON, got: ${text}`, status: res.status }
    }
    if (!res.ok) {
      throw data // webrpc error response
    }
    return data
  })
}
