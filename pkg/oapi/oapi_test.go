package oapi

import (
	"testing"

	"github.com/icza/dyno"
	"github.com/stretchr/testify/assert"
)

func TestCanonicaliseMoRefs(t *testing.T) {
	tests := []struct {
		in         map[string]any
		baseSchema string
		expected   map[string]any
	}{
		{
			in: map[string]any{
				"Organization": "default",
				"NtpServers":   []string{"1.1.1.1", "2.2.2.2"},
			},
			baseSchema: "ntp.Policy",
			expected: map[string]any{
				"Organization": &MoRef{
					Filter:           "Name eq 'default'",
					RelationshipType: "organization.Organization.Relationship",
				},
				"NtpServers": []string{"1.1.1.1", "2.2.2.2"},
			},
		},
		{
			in: map[string]any{
				"Organization": "MoRef[default]",
			},
			baseSchema: "ntp.Policy",
			expected: map[string]any{
				"Organization": &MoRef{
					Filter:           "Name eq 'default'",
					RelationshipType: "organization.Organization.Relationship",
				},
			},
		},
		{
			in: map[string]any{
				"Organization": "MoRef:ntp.Policy[default]",
			},
			baseSchema: "ntp.Policy",
			expected: map[string]any{
				"Organization": &MoRef{
					Filter:           "Name eq 'default'",
					RelationshipType: "ntp.Policy",
				},
			},
		},
		{
			in: map[string]any{
				"ClusterIpPools": []any{"MoRef[ip-pool-1]"},
			},
			baseSchema: "kubernetes.ClusterProfile",
			expected: map[string]any{
				"ClusterIpPools": []any{
					&MoRef{
						Filter:           "Name eq 'ip-pool-1'",
						RelationshipType: "ippool.Pool.Relationship",
					},
				},
			},
		},
		{
			in: map[string]any{
				"ClassId":      "bulk.MoCloner",
				"Organization": "default",
				"Sources":      []any{"MoRef:ServerProfileTemplateRelationship[OCP-BM]"},
			},
			baseSchema: "bulk.MoCloner",
			expected: map[string]any{
				"ClassId": "bulk.MoCloner",
				"Organization": &MoRef{
					Filter:           "Name eq 'default'",
					RelationshipType: "organization.Organization.Relationship",
				},
				"Sources": []any{&MoRef{
					Filter:           "Name eq 'OCP-BM'",
					RelationshipType: "server.ProfileTemplate.Relationship",
				}},
			},
		},
	}

	for _, test := range tests {
		CanonicaliseMoRefs(&(test.in), test.baseSchema)
		assert.Equal(t, test.expected, test.in)
	}
}

func TestGetSchemaProperty(t *testing.T) {
	s := getSchema("ntp.Policy")
	p := getSchemaProperty("Enabled", s)
	assert.NotNil(t, p)
	typ, _ := dyno.GetString(p, "type")
	assert.Equal(t, "boolean", typ)

	p = getSchemaProperty("Moid", s)
	assert.NotNil(t, p)
	typ, _ = dyno.GetString(p, "type")
	assert.Equal(t, "string", typ)

	p = getSchemaProperty("Organization", s)
	assert.NotNil(t, p)
	ref, _ := dyno.GetString(p, "$ref")
	assert.Equal(t, "#/components/schemas/organization.Organization.Relationship", ref)
}

func TestGetSchema(t *testing.T) {
	s := getSchema("aaa.RetentionPolicy")

	assert.NotNil(t, s)
	title, err := dyno.Get(s, "title")
	assert.NoError(t, err)
	assert.Equal(t, "Audit Log Retention Policy", title)

	s = getSchema("fjowisfjoir")
	assert.Nil(t, s)
}

func TestCanonicaliseMoRef(t *testing.T) {
	tests := []struct {
		moref                   string
		defaultRelationshipType string
		res                     *MoRef
	}{
		{
			moref:                   "MoRef[Name:isctl-test]",
			defaultRelationshipType: "organisation.Organisation.Relationship",
			res: &MoRef{
				Filter:           "Name eq 'isctl-test'",
				RelationshipType: "organisation.Organisation.Relationship",
			},
		},
		{
			moref:                   "MoRef[$filter:DeviceHostname eq 'XYZ' and PlatformType eq 'IWE']",
			defaultRelationshipType: "organisation.Organisation.Relationship",
			res: &MoRef{
				Filter:           "DeviceHostname eq 'XYZ' and PlatformType eq 'IWE'",
				RelationshipType: "organisation.Organisation.Relationship",
			},
		},
		{
			moref:                   "MoRef:KubernetesVirtualMachineInfraConfigPolicyRelationship[$filter:DeviceHostname eq 'XYZ' and PlatformType eq 'IWE']",
			defaultRelationshipType: "organisation.Organisation.Relationship",
			res: &MoRef{
				Filter:           "DeviceHostname eq 'XYZ' and PlatformType eq 'IWE'",
				RelationshipType: "kubernetes.VirtualMachineInfraConfigPolicy.Relationship",
			},
		},
		{
			moref:                   "{\"ClassId\": \"Organization\"}",
			defaultRelationshipType: "organisation.Organisation.Relationship",
			res:                     nil,
		},
		{
			moref:                   "cg-k8s-1.18.2",
			defaultRelationshipType: "organisation.Organisation.Relationship",
			res: &MoRef{
				Filter:           "Name eq 'cg-k8s-1.18.2'",
				RelationshipType: "organisation.Organisation.Relationship",
			},
		},
		{
			moref:                   "cg-k8s-1.18[2",
			defaultRelationshipType: "organisation.Organisation.Relationship",
			res:                     nil,
		},
		{
			moref:                   "MoRef[default]",
			defaultRelationshipType: "organisation.Organisation.Relationship",
			res: &MoRef{
				Filter:           "Name eq 'default'",
				RelationshipType: "organisation.Organisation.Relationship",
			},
		},
		{
			moref:                   "MoRef:KubernetesVirtualMachineInfrastructureProviderRelationship[default]",
			defaultRelationshipType: "organisation.Organisation.Relationship",
			res: &MoRef{
				Filter:           "Name eq 'default'",
				RelationshipType: "kubernetes.VirtualMachineInfrastructureProvider.Relationship",
			},
		},
		{
			moref:                   "MoRef:KubernetesInfrastructureProviderRelationship[Name:default]",
			defaultRelationshipType: "organisation.Organisation.Relationship",
			res: &MoRef{
				Filter:           "Name eq 'default'",
				RelationshipType: "kubernetes.InfrastructureProvider.Relationship",
			},
		},
		{
			moref:                   "MoRef[Name:cg-k8s-1.18.2]",
			defaultRelationshipType: "organisation.Organisation.Relationship",
			res: &MoRef{
				Filter:           "Name eq 'cg-k8s-1.18.2'",
				RelationshipType: "organisation.Organisation.Relationship",
			},
		},
		{
			moref:                   "59c84e4a16267c0001c23428",
			defaultRelationshipType: "organisation.Organisation.Relationship",
			res:                     nil,
		},
	}

	for _, test := range tests {
		res := CanonicaliseMoRef(test.moref, test.defaultRelationshipType)
		assert.Equal(t, test.res, res)
	}
}

func TestCanonicaliseRelationshipType(t *testing.T) {
	assert.Equal(t, "kubernetes.VirtualMachineInfrastructureProvider.Relationship", canonicaliseRelationshipType("KubernetesVirtualMachineInfrastructureProviderRelationship"))
	assert.Equal(t, "organization.Organization.Relationship", canonicaliseRelationshipType("OrganizationOrganizationRelationship"))

	assert.Equal(t, "organization.Organization.Relationship", canonicaliseRelationshipType("organization.Organization.Relationship"))
}
