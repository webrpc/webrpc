//go:generate webrpc-gen -schema=example.webrpc.json -target=go -pkg=proto -server -client -websockets=on -out=.
package proto

// this file is useful if someone wants to add some of their own manual code
// perhaps methods, or some special marshalling or whatever
