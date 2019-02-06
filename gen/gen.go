package gen

import (
	"github.com/webrpc/webrpc/gen/golang"
	"github.com/webrpc/webrpc/schema"
)

func Gen(proto *schema.WebRPCSchema) (string, error) {
	return golang.Gen(proto)
}
