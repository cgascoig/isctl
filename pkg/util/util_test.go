package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMoid(t *testing.T) {
	var moid string
	var ok bool

	moid, ok = GetMoid(map[string]interface{}{
		"ObjectType": "ntp.Policy.List",
		"Results": []any{
			map[string]any{
				"Moid": "12345",
			},
		},
	})
	assert.Equal(t, true, ok)
	assert.Equal(t, "12345", moid)
}
