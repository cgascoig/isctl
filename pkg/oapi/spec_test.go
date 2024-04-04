package oapi

import (
	"testing"

	"github.com/icza/dyno"
	"github.com/stretchr/testify/assert"
)

func TestLazyLoadSpec(t *testing.T) {
	spec = lazyLoadSpec()
	assert.NotNil(t, spec)

	oapi, err := dyno.Get(spec, "openapi")
	assert.NoError(t, err)
	assert.Equal(t, "3.0.2", oapi)
}
