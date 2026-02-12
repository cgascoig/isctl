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

	// Test that abstract class MoRef dependencies are resolved to concrete classes.
	// resourcepool.Pool has QualificationPolicies which references
	// resource.AbstractResourceQualificationPolicy (an abstract class).
	// The ordering should ensure resourcepool.QualificationPolicy comes before resourcepool.Pool.
	res, err = getOrderedMOs([]rawMO{
		{
			"ClassId":               "resourcepool.Pool",
			"Name":                  "test-pool",
			"QualificationPolicies": []any{"MoRef[isctl-bats-test]"},
		},
		{
			"ClassId": "resourcepool.QualificationPolicy",
			"Name":    "isctl-bats-test",
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, []rawMO{
		{
			"ClassId": "resourcepool.QualificationPolicy",
			"Name":    "isctl-bats-test",
		},
		{
			"ClassId": "resourcepool.Pool",
			"Name":    "test-pool",
			"QualificationPolicies": []any{&oapi.MoRef{
				RelationshipType: "resource.AbstractResourceQualificationPolicy.Relationship",
				Filter:           "Name eq 'isctl-bats-test'",
			}},
		},
	}, res)
}

func TestGetOrderedMOsWithExplicitMoRef(t *testing.T) {
	// Test that explicity configure mo.MoRef works for references
	res, err := getOrderedMOs([]rawMO{
		{
			"ClassId": "resourcepool.Pool",
			"Name":    "test-pool",
			"QualificationPolicies": []any{
				map[string]any{
					"ClassId":    "mo.MoRef",
					"Moid":       "6421194f6f62692d31f53ec5",
					"ObjectType": "resourcepool.QualificationPolicy",
				},
			},
		},
		{
			"ClassId": "resourcepool.QualificationPolicy",
			"Name":    "isctl-bats-test",
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, []rawMO{
		{
			"ClassId": "resourcepool.QualificationPolicy",
			"Name":    "isctl-bats-test",
		},
		{
			"ClassId": "resourcepool.Pool",
			"Name":    "test-pool",
			"QualificationPolicies": []any{
				map[string]any{
					"ClassId":    "mo.MoRef",
					"Moid":       "6421194f6f62692d31f53ec5",
					"ObjectType": "resourcepool.QualificationPolicy",
				},
			},
		},
	}, res)
}
