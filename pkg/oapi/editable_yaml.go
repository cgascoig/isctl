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
	// Always include ClassId
	writableSet["ClassId"] = true
	writableSet["ObjectType"] = true
	writableSet["Moid"] = true

	for _, name := range writableProps {
		writableSet[name] = true
	}
	for _, name := range writableRels {
		writableSet[name] = true
	}

	// Always include identity constraints
	identityConstraints := meta.GetIdentityConstraints(classId)
	for _, name := range identityConstraints {
		writableSet[name] = true
	}

	// Filter the MO to only include writable properties
	filtered := make(map[string]any)
	for key, value := range mo {
		if writableSet[key] {
			// Recursively filter nested objects, passing classId and key for context
			filtered[key] = filterNestedValue(value, meta, classId, key)
		}
	}

	return filtered, nil
}

// filterNestedValue recursively filters nested objects and slices.
// parentClassId and propName provide context about where this value appears,
// used to look up declared types for ClassId/ObjectType omission.
func filterNestedValue(value any, meta *Meta, parentClassId string, propName string) any {
	switch v := value.(type) {
	case map[string]any:
		// Collapse MoRef objects to shorthand syntax
		if classId, ok := v["ClassId"].(string); ok && classId == "mo.MoRef" {
			objectType, _ := v["ObjectType"].(string)
			expectedType := meta.GetPropertyOrRelationshipType(parentClassId, propName)
			includeType := objectType != "" && objectType != expectedType
			// Use pre-resolved identity string if available (from --readable-morefs)
			if identity, ok := v["_ResolvedIdentity"].(string); ok {
				if includeType {
					return fmt.Sprintf("MoRef:%s[%s]", objectType, identity)
				}
				return fmt.Sprintf("MoRef[%s]", identity)
			}
			if moid, ok := v["Moid"].(string); ok {
				if includeType {
					return fmt.Sprintf("MoRef:%s[Moid:%s]", objectType, moid)
				}
				return fmt.Sprintf("MoRef[Moid:%s]", moid)
			}
		}
		return filterNestedMap(v, meta, parentClassId, propName)
	case []any:
		return filterNestedSlice(v, meta, parentClassId, propName)
	default:
		return value
	}
}

// filterNestedMap filters a nested map, recursively filtering its contents
// based on the ObjectType's class metadata.
// parentClassId and propName are the caller's context, used to look up the
// declared type of this object for ClassId/ObjectType omission.
func filterNestedMap(m map[string]any, meta *Meta, parentClassId string, propName string) map[string]any {
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
			// Always include ObjectType and ClassId for nested objects
			writableSet["ObjectType"] = true
			writableSet["ClassId"] = true
			writableSet["Moid"] = true

			for _, name := range writableProps {
				writableSet[name] = true
			}
			for _, name := range writableRels {
				writableSet[name] = true
			}

			// Always include identity constraints for nested objects
			identityConstraints := meta.GetIdentityConstraints(objectType)
			for _, name := range identityConstraints {
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

		// Recursively filter nested values using current object's type for child context
		filtered[key] = filterNestedValue(value, meta, objectType, key)
	}

	// Omit ClassId/ObjectType if the nested object's ClassId matches the declared type
	expectedType := meta.GetPropertyOrRelationshipType(parentClassId, propName)
	if expectedType != "" {
		if classId, ok := filtered["ClassId"].(string); ok && classId == expectedType {
			delete(filtered, "ClassId")
			delete(filtered, "ObjectType")
		}
	}

	return filtered
}

// filterNestedSlice filters each item in a slice
func filterNestedSlice(s []any, meta *Meta, parentClassId string, propName string) []any {
	filtered := make([]any, len(s))
	for i, item := range s {
		filtered[i] = filterNestedValue(item, meta, parentClassId, propName)
	}
	return filtered
}

// StripEmptyFields recursively removes entries from a map where the value is
// nil, an empty string, an empty slice, or an empty map (after recursive stripping).
// Values like false and 0 are preserved as they are meaningful.
func StripEmptyFields(m map[string]any) map[string]any {
	result := make(map[string]any)
	for k, v := range m {
		stripped := stripEmptyValue(v)
		if stripped != nil {
			result[k] = stripped
		}
	}
	return result
}

// stripEmptyValue returns nil if the value should be considered empty, otherwise
// returns the value (recursively stripped if it is a map or slice).
func stripEmptyValue(v any) any {
	if v == nil {
		return nil
	}
	switch val := v.(type) {
	case string:
		if val == "" {
			return nil
		}
		return val
	case map[string]any:
		stripped := StripEmptyFields(val)
		if len(stripped) == 0 {
			return nil
		}
		return stripped
	case []any:
		if len(val) == 0 {
			return nil
		}
		result := make([]any, 0, len(val))
		for _, item := range val {
			stripped := stripEmptyValue(item)
			if stripped != nil {
				result = append(result, stripped)
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	default:
		return v
	}
}

// FormatEditableYAML filters a managed object to only include writable properties
// and marshals it to YAML format. This produces YAML suitable for editing by users.
// When includeEmptyFields is false (the default), empty/null fields are omitted.
func FormatEditableYAML(mo map[string]any, classId string, includeEmptyFields bool) ([]byte, error) {
	filtered, err := FilterWritableProperties(mo, classId)
	if err != nil {
		return nil, err
	}

	if !includeEmptyFields {
		filtered = StripEmptyFields(filtered)
	}

	yamlBytes, err := yaml.Marshal(filtered)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal to YAML: %w", err)
	}

	return yamlBytes, nil
}
