package main

import (
	"testing"

	"github.com/cgascoig/isctl/pkg/oapi"
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
			"ClassId": "ntp.Policy",
			"Organization": &oapi.MoRef{
				RelationshipType: "organization.Organization.Relationship",
				Filter:           "Name eq 'default'",
			},
		},
	}, res)

	res, err = getOrderedMOs([]rawMO{
		{
			"ClassId": "server.ProfileTemplate",
			"Name":    "test",
			"PolicyBucket": []any{
				"MoRef:BiosPolicyRelationship[cgascoig-bios-policy]",
				"MoRef:IamLdapPolicyRelationship[cgascoig-ldap-policy]",
			},
			"Organization": "default",
		},
		{
			"ClassId": "bios.Policy",
		},
		{
			"ClassId": "iam.LdapPolicy",
		},
	})

	assert.Nil(t, err)

	// Because (in this example) the order of bios.Policy and iam.LdapPolicy is non-deterministic we check for either ordering
	isOrder1 := assert.ObjectsAreEqual([]rawMO{
		{
			"ClassId": "bios.Policy",
		},
		{
			"ClassId": "iam.LdapPolicy",
		},
		{
			"ClassId": "server.ProfileTemplate",
			"Name":    "test",
			"PolicyBucket": []any{
				&oapi.MoRef{
					RelationshipType: "bios.Policy.Relationship",
					Filter:           "Name eq 'cgascoig-bios-policy'",
				},
				&oapi.MoRef{
					RelationshipType: "iam.LdapPolicy.Relationship",
					Filter:           "Name eq 'cgascoig-ldap-policy'",
				},
			},
			"Organization": &oapi.MoRef{
				RelationshipType: "organization.Organization.Relationship",
				Filter:           "Name eq 'default'",
			},
		},
	}, res)
	isOrder2 := assert.ObjectsAreEqual([]rawMO{
		{
			"ClassId": "iam.LdapPolicy",
		},
		{
			"ClassId": "bios.Policy",
		},
		{
			"ClassId": "server.ProfileTemplate",
			"Name":    "test",
			"PolicyBucket": []any{
				&oapi.MoRef{
					RelationshipType: "bios.Policy.Relationship",
					Filter:           "Name eq 'cgascoig-bios-policy'",
				},
				&oapi.MoRef{
					RelationshipType: "iam.LdapPolicy.Relationship",
					Filter:           "Name eq 'cgascoig-ldap-policy'",
				},
			},
			"Organization": &oapi.MoRef{
				RelationshipType: "organization.Organization.Relationship",
				Filter:           "Name eq 'default'",
			},
		},
	}, res)
	assert.True(t, isOrder1 || isOrder2)

	res, err = getOrderedMOs([]rawMO{
		{
			"ClassId":      "bulk.MoCloner",
			"Organization": "default",
			"Sources":      []any{"MoRef:ServerProfileTemplateRelationship[OCP-BM]"},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, []rawMO{
		{
			"ClassId": "bulk.MoCloner",
			"Organization": &oapi.MoRef{
				RelationshipType: "organization.Organization.Relationship",
				Filter:           "Name eq 'default'",
			},
			"Sources": []any{&oapi.MoRef{
				RelationshipType: "server.ProfileTemplate.Relationship",
				Filter:           "Name eq 'OCP-BM'",
			}},
		},
	}, res)
}
