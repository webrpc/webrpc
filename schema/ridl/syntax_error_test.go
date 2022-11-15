package ridl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyntaxError(t *testing.T) {
	syntaxErrors := []string{
		`





		####

		####


		namex = myapi`,
		`name`,
		`name = myapi$`,
		`name = myapi1
			name = myapi2`,
		`
				webrpc = v2
				name = foo
				version = v1

				import -
			`,
		`
				webrpc = v2
				name = foo
				version = v1

				import
				$
			`,
	}

	for i := range syntaxErrors {
		_, err := parseString(syntaxErrors[i])
		assert.Error(t, err)
		t.Logf("%v", err)
	}
}
