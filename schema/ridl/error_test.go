package ridl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError(t *testing.T) {
	syntaxErrors := []string{
		`namex = myapi`,
		`name`,
		`name = myapi$`,
		`name = myapi1
		name = myapi2`,
		`
			webrpc = v1
			name = foo
			version = v1

			import -
		`,
		`
			webrpc = v1
			name = foo
			version = v1

			import
			$
		`,
	}

	for i := range syntaxErrors {
		_, err := Parse(syntaxErrors[i])
		assert.Error(t, err)
		t.Logf("%v", err)
	}
}
