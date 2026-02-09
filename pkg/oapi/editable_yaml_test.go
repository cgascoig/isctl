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

		// Should not include read-only system properties
		assert.NotContains(t, filtered, "Moid")
		assert.NotContains(t, filtered, "ClassId")
		assert.NotContains(t, filtered, "ObjectType")
		assert.NotContains(t, filtered, "CreateTime")
	})

	t.Run("returns error for invalid classId", func(t *testing.T) {
		mo := map[string]any{"Name": "test"}
		filtered, err := FilterWritableProperties(mo, "nonexistent.Class")
		assert.Error(t, err)
		assert.Nil(t, filtered)
	})

	t.Run("returns empty map when no writable properties present", func(t *testing.T) {
		mo := map[string]any{
			"Moid":       "12345",
			"ClassId":    "ntp.Policy",
			"ObjectType": "ntp.Policy",
		}

		filtered, err := FilterWritableProperties(mo, "ntp.Policy")
		require.NoError(t, err)
		require.NotNil(t, filtered)
		// Should be an empty map since we only passed read-only properties
		assert.Empty(t, filtered)
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

		// Should not include read-only properties
		assert.NotContains(t, yamlStr, "Moid")
		assert.NotContains(t, yamlStr, "ClassId")
	})

	t.Run("returns error for invalid classId", func(t *testing.T) {
		mo := map[string]any{"Name": "test"}
		yamlBytes, err := FormatEditableYAML(mo, "nonexistent.Class")
		assert.Error(t, err)
		assert.Nil(t, yamlBytes)
	})
}
