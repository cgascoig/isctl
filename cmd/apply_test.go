package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrderedMOs(t *testing.T) {
	res, err := getOrderedMOs([]rawMO{
		{
			"ClassId":      "ntp.Policy",
			"Organization": "MoRef[Name:default]",
		},
		{
			"ClassId": "organization.Organization",
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, []rawMO{
		{
			"ClassId": "organization.Organization",
		},
		{
			"ClassId":      "ntp.Policy",
			"Organization": "MoRef[Name:default]",
		},
	}, res)
}
