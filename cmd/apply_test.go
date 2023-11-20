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
				"MoRef:BiosPolicyRelationship[cgascoig-bios-policy]",
				"MoRef:IamLdapPolicyRelationship[cgascoig-ldap-policy]",
			},
			"Organization": "default",
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
				"MoRef:BiosPolicyRelationship[cgascoig-bios-policy]",
				"MoRef:IamLdapPolicyRelationship[cgascoig-ldap-policy]",
			},
			"Organization": "default",
		},
	}, res)
	assert.True(t, isOrder1 || isOrder2)
}
