package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTypeToIntersightClassIDMapping(t *testing.T) {
	td := TemplateData{
		Models: map[string]Model{
			"OrganizationOrganization": Model{
				Parents: []string{"OrganizationOrganizationAllOf"},
			},
			"OrganizationOrganizationAllOf": Model{
				Vars: []Var{
					{
						Name:    "ClassId",
						Default: "organization.Organization",
					},
				},
			},
		},
	}

	assert.Equal(t, map[string]string{"OrganizationOrganization": "organization.Organization", "OrganizationOrganizationAllOf": "organization.Organization"}, td.GetTypeToIntersightClassIDMapping())
}
