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
			// Recursively filter nested objects
			filtered[key] = filterNestedValue(value, meta)
		}
	}

	return filtered, nil
}

// filterNestedValue recursively filters nested objects and slices
func filterNestedValue(value any, meta *Meta) any {
	switch v := value.(type) {
	case map[string]any:
		return filterNestedMap(v, meta)
	case []any:
		return filterNestedSlice(v, meta)
	default:
		return value
	}
}

// filterNestedMap filters a nested map, recursively filtering its contents
// based on the ObjectType's class metadata
func filterNestedMap(m map[string]any, meta *Meta) map[string]any {
	filtered := make(map[string]any)

	// Check if this nested object has an ObjectType that we can use to filter
	objectType, hasObjectType := m["ObjectType"].(string)

	var writableSet map[string]bool
	if hasObjectType {
		// Try to get writable properties for this nested class
		writableProps := meta.GetWritablePropertyNames(objectType)
		writableRels := meta.GetWritableRelationshipNames(objectType)

		if writableProps != nil || writableRels != nil {
			writableSet = make(map[string]bool)
			for _, name := range writableProps {
				writableSet[name] = true
			}
			for _, name := range writableRels {
				writableSet[name] = true
			}
		}
	}

	for key, value := range m {
		// If we have metadata for this class, filter to writable properties
		if writableSet != nil {
			if !writableSet[key] {
				continue
			}
		}

		// Recursively filter nested values
		filtered[key] = filterNestedValue(value, meta)
	}

	return filtered
}

// filterNestedSlice filters each item in a slice
func filterNestedSlice(s []any, meta *Meta) []any {
	filtered := make([]any, len(s))
	for i, item := range s {
		filtered[i] = filterNestedValue(item, meta)
	}
	return filtered
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
