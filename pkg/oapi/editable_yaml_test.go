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
			"Moid":         "6421194f6f62692d31f53ec5",
			"ClassId":      "fabric.EthNetworkGroupPolicy",         // Top-level ClassId preserved
			"ObjectType":   "fabric.EthNetworkGroupPolicy",         // Top-level ObjectType preserved
			"Organization": "MoRef[Moid:5ddec4226972652d33548943]", // MoRef collapsed to shorthand
			"Description":  "",
			"Name":         "COMMON-NET-GRP",
			"Tags":         []any{},
			"VlanSettings": map[string]any{
				// ClassId/ObjectType omitted because fabric.VlanSettings matches declared type
				"AllowedVlans": "1-4093",
				"NativeVlan":   1,
				"QinqEnabled":  false,
				"QinqVlan":     2,
			},
		}, filtered)
	})

	t.Run("access.Policy MoRef collapsing and ClassId omission", func(t *testing.T) {
		mo := map[string]any{
			"ClassId":    "access.Policy",
			"ObjectType": "access.Policy",
			"Moid":       "aabbccdd11223344",
			"Name":       "test-access-policy",
			"AddressType": map[string]any{
				"ClassId":    "access.AddressType",
				"ObjectType": "access.AddressType",
				"EnableIpV4": true,
				"EnableIpV6": false,
			},
			"ConfigurationType": map[string]any{
				"ClassId":            "access.ConfigurationType",
				"ObjectType":         "access.ConfigurationType",
				"ConfigureInband":    true,
				"ConfigureOutOfBand": false,
			},
			"InbandIpPool": map[string]any{
				"ClassId":    "mo.MoRef",
				"ObjectType": "ippool.Pool",
				"Moid":       "deadbeef12345678",
			},
			"Organization": map[string]any{
				"ClassId":    "mo.MoRef",
				"ObjectType": "organization.Organization",
				"Moid":       "5ddec4226972652d33548943",
			},
			"Profiles": []any{
				map[string]any{
					"ClassId":    "mo.MoRef",
					"ObjectType": "server.Profile",
					"Moid":       "profile001moid",
				},
			},
			"CreateTime": "2024-01-01T00:00:00Z", // Read-only, should be filtered
		}

		filtered, err := FilterWritableProperties(mo, "access.Policy")
		require.NoError(t, err)
		require.NotNil(t, filtered)

		// Top-level ClassId/ObjectType preserved
		assert.Equal(t, "access.Policy", filtered["ClassId"])
		assert.Equal(t, "access.Policy", filtered["ObjectType"])
		assert.Equal(t, "aabbccdd11223344", filtered["Moid"])
		assert.Equal(t, "test-access-policy", filtered["Name"])
		assert.NotContains(t, filtered, "CreateTime")

		// MoRefs should be collapsed to shorthand, omitting ObjectType when it matches declared type
		assert.Equal(t, "MoRef[Moid:deadbeef12345678]", filtered["InbandIpPool"])
		assert.Equal(t, "MoRef[Moid:5ddec4226972652d33548943]", filtered["Organization"])

		// Profiles array items: server.Profile differs from declared policy.AbstractConfigProfile, so type is included
		profiles, ok := filtered["Profiles"].([]any)
		require.True(t, ok)
		require.Len(t, profiles, 1)
		assert.Equal(t, "MoRef:server.Profile[Moid:profile001moid]", profiles[0])

		// AddressType should not have ClassId/ObjectType (declared type matches)
		addressType, ok := filtered["AddressType"].(map[string]any)
		require.True(t, ok)
		assert.NotContains(t, addressType, "ClassId")
		assert.NotContains(t, addressType, "ObjectType")
		assert.Equal(t, true, addressType["EnableIpV4"])
		assert.Equal(t, false, addressType["EnableIpV6"])

		// ConfigurationType should not have ClassId/ObjectType (declared type matches)
		configType, ok := filtered["ConfigurationType"].(map[string]any)
		require.True(t, ok)
		assert.NotContains(t, configType, "ClassId")
		assert.NotContains(t, configType, "ObjectType")
		assert.Equal(t, true, configType["ConfigureInband"])
		assert.Equal(t, false, configType["ConfigureOutOfBand"])
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

func TestStripEmptyFields(t *testing.T) {
	t.Run("removes nil values", func(t *testing.T) {
		m := map[string]any{"a": nil, "b": "hello"}
		result := StripEmptyFields(m)
		assert.NotContains(t, result, "a")
		assert.Equal(t, "hello", result["b"])
	})

	t.Run("removes empty strings", func(t *testing.T) {
		m := map[string]any{"empty": "", "nonempty": "hi"}
		result := StripEmptyFields(m)
		assert.NotContains(t, result, "empty")
		assert.Equal(t, "hi", result["nonempty"])
	})

	t.Run("removes empty slices", func(t *testing.T) {
		m := map[string]any{"empty": []any{}, "nonempty": []any{"x"}}
		result := StripEmptyFields(m)
		assert.NotContains(t, result, "empty")
		assert.Contains(t, result, "nonempty")
	})

	t.Run("removes empty maps", func(t *testing.T) {
		m := map[string]any{"empty": map[string]any{}, "nonempty": map[string]any{"k": "v"}}
		result := StripEmptyFields(m)
		assert.NotContains(t, result, "empty")
		assert.Contains(t, result, "nonempty")
	})

	t.Run("removes maps that become empty after recursive stripping", func(t *testing.T) {
		m := map[string]any{
			"nested": map[string]any{"inner": ""},
		}
		result := StripEmptyFields(m)
		assert.NotContains(t, result, "nested")
	})

	t.Run("preserves false boolean values", func(t *testing.T) {
		m := map[string]any{"flag": false, "other": true}
		result := StripEmptyFields(m)
		assert.Equal(t, false, result["flag"])
		assert.Equal(t, true, result["other"])
	})

	t.Run("preserves zero integer values", func(t *testing.T) {
		m := map[string]any{"count": 0, "size": 42}
		result := StripEmptyFields(m)
		assert.Equal(t, 0, result["count"])
		assert.Equal(t, 42, result["size"])
	})

	t.Run("preserves non-empty nested structures", func(t *testing.T) {
		m := map[string]any{
			"nested": map[string]any{"key": "value", "empty": ""},
		}
		result := StripEmptyFields(m)
		nested, ok := result["nested"].(map[string]any)
		require.True(t, ok)
		assert.Equal(t, "value", nested["key"])
		assert.NotContains(t, nested, "empty")
	})

	t.Run("strips empty items from slices", func(t *testing.T) {
		m := map[string]any{
			"items": []any{"a", "", "b"},
		}
		result := StripEmptyFields(m)
		items, ok := result["items"].([]any)
		require.True(t, ok)
		assert.Equal(t, []any{"a", "b"}, items)
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

		yamlBytes, err := FormatEditableYAML(mo, "ntp.Policy", false)
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
		yamlBytes, err := FormatEditableYAML(mo, "nonexistent.Class", false)
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

		yamlBytes, err := FormatEditableYAML(mo, "ntp.Policy", false)
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

		yamlBytes, err := FormatEditableYAML(mo, "policy.AbstractConfigProfile", false)
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

		yamlBytes, err := FormatEditableYAML(mo, "ntp.Policy", false)
		require.NoError(t, err)
		require.NotEmpty(t, yamlBytes)

		yamlStr := string(yamlBytes)
		assert.NotContains(t, yamlStr, "Name:") // "Name" is identity but not in 'mo'
		assert.Contains(t, yamlStr, "Enabled: true")
		assert.Contains(t, yamlStr, "ClassId: ntp.Policy")
		assert.Contains(t, yamlStr, "ObjectType: ntp.Policy")
	})

	t.Run("strips empty fields by default", func(t *testing.T) {
		mo := map[string]any{
			"Name":        "test-policy",
			"Description": "",
			"Enabled":     true,
			"NtpServers":  []any{},
			"Moid":        "12345",
			"ClassId":     "ntp.Policy",
			"ObjectType":  "ntp.Policy",
		}

		yamlBytes, err := FormatEditableYAML(mo, "ntp.Policy", false)
		require.NoError(t, err)
		yamlStr := string(yamlBytes)
		assert.NotContains(t, yamlStr, "Description:")
		assert.NotContains(t, yamlStr, "NtpServers:")
		assert.Contains(t, yamlStr, "Name: test-policy")
		assert.Contains(t, yamlStr, "Enabled: true")
	})

	t.Run("includes empty fields when includeEmptyFields is true", func(t *testing.T) {
		mo := map[string]any{
			"Name":        "test-policy",
			"Description": "",
			"Enabled":     true,
			"NtpServers":  []any{},
			"Moid":        "12345",
			"ClassId":     "ntp.Policy",
			"ObjectType":  "ntp.Policy",
		}

		yamlBytes, err := FormatEditableYAML(mo, "ntp.Policy", true)
		require.NoError(t, err)
		yamlStr := string(yamlBytes)
		assert.Contains(t, yamlStr, "Description:")
		assert.Contains(t, yamlStr, "NtpServers:")
		assert.Contains(t, yamlStr, "Name: test-policy")
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

func TestFilterNestedValueResolvedIdentity(t *testing.T) {
	meta, err := GetMeta()
	require.NoError(t, err)

	t.Run("uses _ResolvedIdentity when present", func(t *testing.T) {
		mo := map[string]any{
			"ClassId":    "ntp.Policy",
			"ObjectType": "ntp.Policy",
			"Moid":       "aabbccddee1122334455aabb",
			"Name":       "test",
			"Organization": map[string]any{
				"ClassId":           "mo.MoRef",
				"Moid":              "5ddec4226972652d33548943",
				"ObjectType":        "organization.Organization",
				"_ResolvedIdentity": "Name:default",
			},
		}
		filtered, err := FilterWritableProperties(mo, "ntp.Policy")
		require.NoError(t, err)
		assert.Equal(t, "MoRef[Name:default]", filtered["Organization"])
	})

	t.Run("falls back to Moid when _ResolvedIdentity absent", func(t *testing.T) {
		mo := map[string]any{
			"ClassId":    "ntp.Policy",
			"ObjectType": "ntp.Policy",
			"Moid":       "aabbccddee1122334455aabb",
			"Name":       "test",
			"Organization": map[string]any{
				"ClassId":    "mo.MoRef",
				"Moid":       "5ddec4226972652d33548943",
				"ObjectType": "organization.Organization",
			},
		}
		filtered, err := FilterWritableProperties(mo, "ntp.Policy")
		require.NoError(t, err)
		assert.Equal(t, "MoRef[Moid:5ddec4226972652d33548943]", filtered["Organization"])
	})

	t.Run("splitBalancedCommas handles nested brackets", func(t *testing.T) {
		_ = meta // avoid unused var
		parts := splitBalancedCommas("VlanId:100,EthNetworkPolicy:MoRef:fabric.EthNetworkPolicy[Name:my-policy]")
		assert.Equal(t, []string{"VlanId:100", "EthNetworkPolicy:MoRef:fabric.EthNetworkPolicy[Name:my-policy]"}, parts)
	})

	t.Run("splitBalancedCommas handles simple string", func(t *testing.T) {
		parts := splitBalancedCommas("Name:default,Account:acc1")
		assert.Equal(t, []string{"Name:default", "Account:acc1"}, parts)
	})
}
