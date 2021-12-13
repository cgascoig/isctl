package main

import (
	"testing"

	openapi "github.com/cgascoig/intersight-go-sdk/intersight"
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
			in:  openapi.MoVersionContext{},
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

func TestSetMoRefSelector(t *testing.T) {
	tests := []struct {
		in       interface{}
		selector string
		out      interface{}
	}{
		{
			in:       &openapi.OrganizationOrganizationRelationship{},
			selector: "Name eq 'test'",
			out: &openapi.OrganizationOrganizationRelationship{
				MoMoRef: &openapi.MoMoRef{
					ClassId:    "mo.MoRef",
					Selector:   optionalStr("Name eq 'test'"),
					ObjectType: "organization.Organization",
				},
			},
		},
	}

	for _, test := range tests {
		setMoRefSelector(test.in, test.selector)
		assert.Equal(t, test.in, test.out)
	}
}
