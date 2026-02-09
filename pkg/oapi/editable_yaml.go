package oapi

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// FilterWritableProperties filters a managed object to only include properties
// and relationships with ApiAccess == "ReadWrite". Returns a new map containing
// only the writable fields. This is useful for edit operations where users should
// only modify writable properties.
func FilterWritableProperties(mo map[string]any, classId string) (map[string]any, error) {
	meta, err := GetMeta()
	if err != nil {
		return nil, fmt.Errorf("failed to load metadata: %w", err)
	}

	writableProps := meta.GetWritablePropertyNames(classId)
	writableRels := meta.GetWritableRelationshipNames(classId)

	if writableProps == nil && writableRels == nil {
		return nil, fmt.Errorf("class %s not found in metadata", classId)
	}

	// Create a set for O(1) lookup
	writableSet := make(map[string]bool)
	for _, name := range writableProps {
		writableSet[name] = true
	}
	for _, name := range writableRels {
		writableSet[name] = true
	}

	// Filter the MO to only include writable properties
	filtered := make(map[string]any)
	for key, value := range mo {
		if writableSet[key] {
			filtered[key] = value
		}
	}

	return filtered, nil
}

// FormatEditableYAML filters a managed object to only include writable properties
// and marshals it to YAML format. This produces YAML suitable for editing by users.
func FormatEditableYAML(mo map[string]any, classId string) ([]byte, error) {
	filtered, err := FilterWritableProperties(mo, classId)
	if err != nil {
		return nil, err
	}

	yamlBytes, err := yaml.Marshal(filtered)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal to YAML: %w", err)
	}

	return yamlBytes, nil
}
