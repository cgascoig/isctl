package main

import (
	"testing"

	"github.com/cgascoig/isctl/openapi"
	"github.com/stretchr/testify/assert"
)

type moRefTest struct {
	in  interface{}
	out bool
}

func TestIsMoRef(t *testing.T) {
	var tests []moRefTest
	tests = []moRefTest{
		{
			in:  "1234",
			out: false,
		},
		{
			in:  openapi.OrganizationOrganizationRelationship{},
			out: true,
		},
		{
			in:  openapi.NtpPolicy{},
			out: false,
		},
		{
			in: openapi.MoVersionContext{},
			out: false,
		},
		{
			in:  &openapi.OrganizationOrganizationRelationship{},
			out: true,
		},
		{
			in:  &openapi.NtpPolicy{},
			out: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, isMoRef(test.in))
	}
}
