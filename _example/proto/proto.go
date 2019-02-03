//go:generate webrpc-gen -schema=proto.json -target=go -pkg=proto -server -client -websockets=on -out=.
package proto
