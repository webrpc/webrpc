package target

import "github.com/webrpc/webrpc-go/schema"

type Generator interface {
	Gen(proto *schema.WebRPCSchema) (string, error)
}
