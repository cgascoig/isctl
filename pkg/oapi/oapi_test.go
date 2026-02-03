package oapi

import (
	"encoding/json"
	"strings"
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
				"Targets": []any{
					map[string]any{
						"Organization": "MoRef:organization.Organization[default]",
					},
				},
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
				"Targets": []any{
					map[string]any{
						"Organization": &MoRef{
							Filter:           "Name eq 'default'",
							RelationshipType: "organization.Organization",
						},
					},
				},
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
		{
			moref:                   "MoRef:organization.Organization[default]",
			defaultRelationshipType: "",
			res: &MoRef{
				Filter:           "Name eq 'default'",
				RelationshipType: "organization.Organization",
			},
		},
		{
			moref:                   "MoRef[default\\test]",
			defaultRelationshipType: "ntp.Policy",
			res: &MoRef{
				Filter:           "Name eq 'test'",
				RelationshipType: "ntp.Policy",
				Organization:     "default",
			},
		},
		{
			moref:                   "MoRef:iam.EndPointUser[default\\test]",
			defaultRelationshipType: "",
			res: &MoRef{
				Filter:           "Name eq 'test'",
				RelationshipType: "iam.EndPointUser",
				Organization:     "default",
			},
		},
		{
			moref:                   "MoRef:hcl.OperatingSystem[Red Hat Enterprise Linux 7.6]",
			defaultRelationshipType: "hcl.OperatingSystem.Relationship",
			res: &MoRef{
				Filter:           "Name eq 'Red Hat Enterprise Linux 7.6'",
				RelationshipType: "hcl.OperatingSystem",
			},
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

func TestClassIdHasProperty(t *testing.T) {
	assert.True(t, ClassIdHasProperty("ntp.Policy", "Organization"))
	assert.False(t, ClassIdHasProperty("vnic.EthIf", "Organization"))
	assert.False(t, ClassIdHasProperty("organization.Organization", "Organization"))
	assert.True(t, ClassIdHasProperty("fabric.Vlan", "VlanId"))
}

func TestSchemaToVarsWithExample(t *testing.T) {
	// Mock a schema that uses a known complex type
	// We'll use ippool.IpV4Block again as we know it works from the previous test
	schema := map[string]any{
		"properties": map[string]any{
			"MyBlock": map[string]any{
				"$ref":        "#/components/schemas/ippool.IpV4Block",
				"description": "My custom block",
			},
			"MyList": map[string]any{
				"type": "array",
				"items": map[string]any{
					"$ref": "#/components/schemas/ippool.IpV4Config",
				},
				"description": "My list of configs",
			},
		},
	}

	vars := schemaToVars(schema)
	assert.Len(t, vars, 2)

	for _, v := range vars {
		if v.Name == "MyBlock" {
			assert.Contains(t, v.Usage, "My custom block")
			assert.Contains(t, v.Usage, "{")
			assert.Contains(t, v.Usage, "\"From\": \"<First IPv4 address of the block.>\"")
			assert.NotContains(t, v.Usage, "\"ClassId\"")
			assert.NotContains(t, v.Usage, "\"ObjectType\"")
			// Check for newline separation
			assert.True(t, strings.Contains(v.Usage, "My custom block\nExample: \n{"), "Usage should contain description followed by 'Example:' and the example starting on new line")
		}
		if v.Name == "MyList" {
			assert.Contains(t, v.Usage, "My list of configs")
			// Depending on exact json encoding, it might be [\n  {
			assert.Contains(t, v.Usage, "[\n  {")
			assert.Contains(t, v.Usage, "\"Gateway\": \"<IP address of the default IPv4 gateway.>\"")
			assert.NotContains(t, v.Usage, "\"ClassId\"")
			assert.NotContains(t, v.Usage, "\"ObjectType\"")
			assert.True(t, strings.HasSuffix(v.Usage, "\n]\n"), "Array example should end with ']\\n'")
		}
	}
}

func TestGetSchemaExample(t *testing.T) {
	// We rely on the embedded spec, so we pick a known schema.
	// ippool.IpV4Config is mentioned by the user.
	schemaName := "#/components/schemas/ippool.IpV4Config"

	propSchema := map[string]any{"$ref": schemaName}
	exampleJSON, err := GetSchemaExample(propSchema)
	assert.NoError(t, err)
	assert.NotEmpty(t, exampleJSON)

	// Validate it's valid JSON
	var result map[string]any
	err = json.Unmarshal([]byte(exampleJSON), &result)
	assert.NoError(t, err)

	// Check for expected fields based on user request
	assert.Contains(t, result, "Gateway")
	assert.Contains(t, result, "Netmask")
	assert.Contains(t, result, "PrimaryDns")
	assert.Contains(t, result, "SecondaryDns")

	// Verify excluded fields
	assert.NotContains(t, result, "ClassId")
	assert.NotContains(t, result, "ObjectType")

	// Verify values are descriptions (strings starting with < and ending with >)
	for k, v := range result {
		s, ok := v.(string)
		if assert.True(t, ok, "Value for %s should be a string", k) {
			assert.True(t, (len(s) > 0 && s[0] == '<' && s[len(s)-1] == '>'), "Value for %s should be a description wrapped in <>: %s", k, s)
		}
	}

	// fmt.Printf("Generated Example:\n%s\n", exampleJSON)

	// Test case for nested schema
	t.Run("ippool.IpV4Block", func(t *testing.T) {
		schemaName := "#/components/schemas/ippool.IpV4Block"
		propSchema := map[string]any{"$ref": schemaName}
		exampleJSON, err := GetSchemaExample(propSchema)
		assert.NoError(t, err)
		assert.NotEmpty(t, exampleJSON)

		var result map[string]any
		err = json.Unmarshal([]byte(exampleJSON), &result)
		assert.NoError(t, err)

		// Check top level simple fields
		// Verify excluded fields
		assert.NotContains(t, result, "ClassId")
		assert.NotContains(t, result, "ObjectType")
		assert.Contains(t, result, "From")
		assert.Contains(t, result, "To")
		assert.Contains(t, result, "Size")

		// Verify excluded fields
		assert.NotContains(t, result, "ClassId")
		assert.NotContains(t, result, "ObjectType")

		// Check nested object
		val, ok := result["IpV4Config"]
		assert.True(t, ok, "IpV4Config should be present")

		nested, ok := val.(map[string]any)
		assert.True(t, ok, "IpV4Config should be a map")

		if ok {
			assert.Contains(t, nested, "Gateway")
			assert.Contains(t, nested, "Netmask")
			assert.Contains(t, nested, "PrimaryDns")
			assert.Contains(t, nested, "SecondaryDns")
		}

		// fmt.Printf("Generated Example for IpV4Block:\n%s\n", exampleJSON)
	})
}
