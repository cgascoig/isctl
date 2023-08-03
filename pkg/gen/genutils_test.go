package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceArgs(t *testing.T) {
	var s string
	var e error
	s, e = ReplaceArgs("api/{Moid}", []string{"12345"})
	assert.NoError(t, e)
	assert.Equal(t, "api/12345", s)
}
