package ridl

import (
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestLexerBasic(t *testing.T) {
	buf := `


		ridl						 =  v1

				+     foo=bar

					-baz   = 56 # a comment

													version=                    v0.0.1


foo=bar`

	tokens, err := tokenize(buf)
	assert.NoError(t, err)

	log.Printf("buf: %v", string(buf))
	spew.Dump(tokens)
}

func TestLexerBig(t *testing.T) {
	buf := `ridl = v1

	name = hello-webrpc
	version = v0.0.1
	
	
	import
		- ./blah.ridl
		- ./abc.json
	
	
	# this is a comment
	# yep
	enum Kind: uint32
		- USER = 1             # comment
		- ADMIN = 2            # comment..
	
	# or.. just..
	enum Kind: uint32
		- USER                 # aka, = 0
		- ADMIN                # aka, = 1
	
	
	message Empty
	
	
	message GetUserRequest
		- userID: uint64
	
	
	message User
		- ID: uint64          # hmm.. TODO, think about the name of the filed we want..? should it represent the JSON field name? prob
			+ json = id         # if so, we then need to add the json thing here .....
			+ go.tag.db = id
	
		- meta: map<string,int64>

		- username: string
			+ json = USERNAME
			+ go.tag.db = username
	
		- createdAt?: timestamp
			+ json = created_at,omitempty
			+ go.tag.db = created_at
	
	
	message Role
	- name: string

	- users: map<uint64,User>

	- perms: []string

	service ExampleService
	
		# comment can go here
		# too .. :)
		- Ping(): bool
	
		- GetUser(GetUserRequest): User
	
		- Recv(stream bytes)
	
		- Send(): stream bytes
	
		- SendAndRecv(stream bytes): stream bytes
	
		- Broadcast(): stream Notice	
`

	tokens, err := tokenize(buf)
	assert.NoError(t, err)

	log.Printf("buf: %v", string(buf))
	spew.Dump(tokens)

}
