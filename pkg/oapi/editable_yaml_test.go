package oapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetClassMeta(t *testing.T) {
	meta, err := GetMeta()
	require.NoError(t, err)

	t.Run("returns ClassMeta for valid classId", func(t *testing.T) {
		cm := meta.GetClassMeta("ntp.Policy")
		require.NotNil(t, cm)
		assert.Equal(t, "ntp.Policy", cm.Name)
		assert.True(t, cm.IsConcrete)
	})

	t.Run("returns nil for invalid classId", func(t *testing.T) {
		cm := meta.GetClassMeta("nonexistent.Class")
		assert.Nil(t, cm)
	})
}

func TestGetWritablePropertyNames(t *testing.T) {
	meta, err := GetMeta()
	require.NoError(t, err)

	t.Run("returns writable property names for valid class", func(t *testing.T) {
		names := meta.GetWritablePropertyNames("ntp.Policy")
		require.NotNil(t, names)
		// NTP Policy should have some writable properties
		assert.Contains(t, names, "Enabled")
		assert.Contains(t, names, "Name")
		assert.Contains(t, names, "NtpServers")
	})

	t.Run("returns nil for invalid class", func(t *testing.T) {
		names := meta.GetWritablePropertyNames("nonexistent.Class")
		assert.Nil(t, names)
	})

	t.Run("does not include read-only properties", func(t *testing.T) {
		names := meta.GetWritablePropertyNames("ntp.Policy")
		// Check that common read-only system properties are not included
		for _, name := range names {
			assert.NotEqual(t, "Moid", name)
			assert.NotEqual(t, "ClassId", name)
			assert.NotEqual(t, "ObjectType", name)
		}
	})
}

func TestGetWritableRelationshipNames(t *testing.T) {
	meta, err := GetMeta()
	require.NoError(t, err)

	t.Run("returns writable relationship names for valid class", func(t *testing.T) {
		names := meta.GetWritableRelationshipNames("ntp.Policy")
		require.NotNil(t, names)
		// NTP Policy has Profiles as a writable relationship
		assert.Contains(t, names, "Profiles")
	})

	t.Run("returns nil for invalid class", func(t *testing.T) {
		names := meta.GetWritableRelationshipNames("nonexistent.Class")
		assert.Nil(t, names)
	})
}

func TestFilterWritableProperties(t *testing.T) {
	t.Run("filters to only writable properties", func(t *testing.T) {
		mo := map[string]any{
			"Name":       "test-policy",
			"Enabled":    true,
			"NtpServers": []string{"1.1.1.1"},
			"Moid":       "12345",
			"ClassId":    "ntp.Policy",
			"ObjectType": "ntp.Policy",
			"CreateTime": "2024-01-01T00:00:00Z",
		}

		filtered, err := FilterWritableProperties(mo, "ntp.Policy")
		require.NoError(t, err)
		require.NotNil(t, filtered)

		// Should include writable properties
		assert.Contains(t, filtered, "Name")
		assert.Contains(t, filtered, "Enabled")
		assert.Contains(t, filtered, "NtpServers")
		assert.Contains(t, filtered, "Moid")

		// Should include ClassId and ObjectType now
		assert.Contains(t, filtered, "ClassId")
		assert.Contains(t, filtered, "ObjectType")

		// Should not include read-only system properties
		assert.NotContains(t, filtered, "CreateTime")
	})

	t.Run("filters to only writable properties recursively", func(t *testing.T) {
		mo := map[string]any{
			"AccountMoid":     "59c84e4a16267c0001c23428",
			"Ancestors":       []any{},
			"ClassId":         "fabric.EthNetworkGroupPolicy",
			"CreateTime":      "2023-03-27T04:19:27.936Z",
			"Description":     "",
			"DomainGroupMoid": "5b25418d7a7662743465cf72",
			"ModTime":         "2023-03-27T04:19:27.937Z",
			"Moid":            "6421194f6f62692d31f53ec5",
			"Name":            "COMMON-NET-GRP",
			"ObjectType":      "fabric.EthNetworkGroupPolicy",
			"Organization": map[string]any{
				"ClassId":    "mo.MoRef",
				"Moid":       "5ddec4226972652d33548943",
				"ObjectType": "organization.Organization",
				"link":       "https://intersight.com/api/v1/organization/Organizations/5ddec4226972652d33548943",
			},
			"Owners": []any{
				"59c84e4a16267c0001c23428",
			},
			"PermissionResources": []any{
				map[string]any{
					"ClassId":    "mo.MoRef",
					"Moid":       "5ddec4226972652d33548943",
					"ObjectType": "organization.Organization",
					"link":       "https://intersight.com/api/v1/organization/Organizations/5ddec4226972652d33548943",
				},
			},
			"SharedScope": "",
			"Tags":        []any{},
			"VlanSettings": map[string]any{
				"AllowedVlans": "1-4093",
				"ClassId":      "fabric.VlanSettings",
				"NativeVlan":   1,
				"ObjectType":   "fabric.VlanSettings",
				"QinqEnabled":  false,
				"QinqVlan":     2,
			},
		}

		filtered, err := FilterWritableProperties(mo, "fabric.EthNetworkGroupPolicy")
		require.NoError(t, err)
		require.NotNil(t, filtered)

		assert.Equal(t, map[string]any{
			"Moid":       "6421194f6f62692d31f53ec5",
			"ClassId":    "fabric.EthNetworkGroupPolicy", // ClassId should be preserved
			"ObjectType": "fabric.EthNetworkGroupPolicy", // ObjectType should be preserved
			"Organization": map[string]any{
				"Moid":       "5ddec4226972652d33548943",
				"ClassId":    "mo.MoRef",
				"ObjectType": "organization.Organization",
			},
			"Description": "",
			"Name":        "COMMON-NET-GRP",
			"Tags":        []any{},
			"VlanSettings": map[string]any{
				"AllowedVlans": "1-4093",
				"ClassId":      "fabric.VlanSettings", // Nested ClassId preserved
				"NativeVlan":   1,
				"QinqEnabled":  false,
				"QinqVlan":     2,
				"ObjectType":   "fabric.VlanSettings",
			},
		}, filtered)
	})

	t.Run("returns error for invalid classId", func(t *testing.T) {
		mo := map[string]any{"Name": "test"}
		filtered, err := FilterWritableProperties(mo, "nonexistent.Class")
		assert.Error(t, err)
		assert.Nil(t, filtered)
	})

	t.Run("returns correct map when some properties present", func(t *testing.T) {
		mo := map[string]any{
			"Moid":       "12345",
			"ClassId":    "ntp.Policy",
			"ObjectType": "ntp.Policy",
			"CreateTime": "2023-03-27T04:19:27.936Z",
		}

		filtered, err := FilterWritableProperties(mo, "ntp.Policy")
		require.NoError(t, err)
		require.NotNil(t, filtered)
		// Should contain ClassId and ObjectType even if no writable properties
		assert.Contains(t, filtered, "ClassId")
		assert.Contains(t, filtered, "ObjectType")
		assert.Contains(t, filtered, "Moid")
		assert.Len(t, filtered, 3)
	})
}

func TestFormatEditableYAML(t *testing.T) {
	t.Run("produces valid YAML with only writable properties", func(t *testing.T) {
		mo := map[string]any{
			"Name":       "test-policy",
			"Enabled":    true,
			"NtpServers": []string{"1.1.1.1", "2.2.2.2"},
			"Moid":       "12345",
			"ClassId":    "ntp.Policy",
			"ObjectType": "ntp.Policy",
		}

		yamlBytes, err := FormatEditableYAML(mo, "ntp.Policy")
		require.NoError(t, err)
		require.NotEmpty(t, yamlBytes)

		yamlStr := string(yamlBytes)
		// Should include writable properties
		assert.Contains(t, yamlStr, "Name")
		assert.Contains(t, yamlStr, "test-policy")
		assert.Contains(t, yamlStr, "Enabled")
		assert.Contains(t, yamlStr, "NtpServers")

		// Should include ClassId
		assert.Contains(t, yamlStr, "ClassId")
		assert.Contains(t, yamlStr, "ntp.Policy")
		assert.Contains(t, yamlStr, "ObjectType")
		assert.Contains(t, yamlStr, "Moid")

		// Should not include read-only properties
		assert.NotContains(t, yamlStr, "CreateTime")
	})

	t.Run("returns error for invalid classId", func(t *testing.T) {
		mo := map[string]any{"Name": "test"}
		yamlBytes, err := FormatEditableYAML(mo, "nonexistent.Class")
		assert.Error(t, err)
		assert.Nil(t, yamlBytes)
	})

	t.Run("includes identity properties for a class with identity", func(t *testing.T) {
		mo := map[string]any{
			"Name":       "test-policy",
			"Enabled":    true,
			"Moid":       "12345",
			"ClassId":    "ntp.Policy",
			"ObjectType": "ntp.Policy",
		}

		yamlBytes, err := FormatEditableYAML(mo, "ntp.Policy")
		require.NoError(t, err)
		require.NotEmpty(t, yamlBytes)

		yamlStr := string(yamlBytes)
		assert.Contains(t, yamlStr, "Name: test-policy") // Identity property
		assert.Contains(t, yamlStr, "Enabled: true")     // Writable property
		assert.Contains(t, yamlStr, "ClassId: ntp.Policy")
		assert.Contains(t, yamlStr, "ObjectType: ntp.Policy")
		assert.Contains(t, yamlStr, "Moid")
	})

	t.Run("includes identity properties for a class with composite identity", func(t *testing.T) {
		mo := map[string]any{
			"Name":       "test-profile",
			"Type":       "instance",
			"Moid":       "12345",
			"ClassId":    "policy.AbstractConfigProfile",
			"ObjectType": "policy.AbstractConfigProfile",
		}

		yamlBytes, err := FormatEditableYAML(mo, "policy.AbstractConfigProfile")
		require.NoError(t, err)
		require.NotEmpty(t, yamlBytes)

		yamlStr := string(yamlBytes)
		assert.Contains(t, yamlStr, "Name: test-profile") // Identity property
		assert.Contains(t, yamlStr, "Type: instance")     // Identity property
		assert.Contains(t, yamlStr, "ClassId: policy.AbstractConfigProfile")
		assert.Contains(t, yamlStr, "ObjectType: policy.AbstractConfigProfile")
		assert.Contains(t, yamlStr, "Moid")
	})

	t.Run("does not include identity properties if not present in original object", func(t *testing.T) {
		mo := map[string]any{
			"Enabled":    true,
			"Moid":       "12345",
			"ClassId":    "ntp.Policy",
			"ObjectType": "ntp.Policy",
		}

		yamlBytes, err := FormatEditableYAML(mo, "ntp.Policy")
		require.NoError(t, err)
		require.NotEmpty(t, yamlBytes)

		yamlStr := string(yamlBytes)
		assert.NotContains(t, yamlStr, "Name:") // "Name" is identity but not in 'mo'
		assert.Contains(t, yamlStr, "Enabled: true")
		assert.Contains(t, yamlStr, "ClassId: ntp.Policy")
		assert.Contains(t, yamlStr, "ObjectType: ntp.Policy")
	})
}

func TestFilterWritablePropertiesIncludesIdentity(t *testing.T) {
	// fabric.Vlan has "VlanId" and "EthNetworkPolicy" as identity constraints.
	input := map[string]any{
		"ClassId": "fabric.Vlan",
		"VlanId":  100,
		"EthNetworkPolicy": map[string]any{
			"Moid": "some-moid",
		},
		"Name":       "test-vlan",            // Writable
		"CreateTime": "2023-01-01T00:00:00Z", // ReadOnly
	}

	filtered, err := FilterWritableProperties(input, "fabric.Vlan")
	assert.NoError(t, err)

	// ClassId must be present
	assert.Contains(t, filtered, "ClassId")
	assert.Equal(t, "fabric.Vlan", filtered["ClassId"])

	// Identity fields must be present
	assert.Contains(t, filtered, "VlanId")
	assert.Contains(t, filtered, "EthNetworkPolicy")

	// Writable fields must be present
	assert.Contains(t, filtered, "Name")

	// ReadOnly fields (non-identity) must be absent
	assert.NotContains(t, filtered, "CreateTime")
}
