package ridl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnescapeStringErrorNamesOffendingChar(t *testing.T) {
	_, err := unescapeString(`a\qb`)
	require.Error(t, err)
	assert.EqualError(t, err, `unexpected character 'q' after backslash`)
}
