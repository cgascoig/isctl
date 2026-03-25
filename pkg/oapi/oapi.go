package oapi

import (
	"bytes"
	"encoding/json"
	"fmt"

	"regexp"
	"strings"

	"github.com/icza/dyno"
	log "github.com/sirupsen/logrus"
)

func getObjectProperty(propName string, obj any) map[string]any {
	props, err := dyno.GetMapS(obj, "properties")
	if err == nil {
		prop, err := dyno.GetMapS(props, propName)
		if err == nil {
			return prop
		}
	}
	return nil
}

func getSchemaProperty(propName string, schema map[string]any) map[string]any {
	// is the schema an "object" itself
	if typ, err := dyno.GetString(schema, "type"); err == nil && typ == "object" {
		prop := getObjectProperty(propName, schema)
		if prop != nil {
			return prop
		}
	}

	// go through the allOfs
	allOf, err := dyno.GetSlice(schema, "allOf")
	if err == nil {
		for _, v := range allOf {

			// this item in allOf is a literal object
			typ, err := dyno.GetString(v, "type")
			if err == nil && typ == "object" {
				prop := getObjectProperty(propName, v)
				if prop != nil {
					return prop
				}
			}

			// this item in allOf is a ref
			ref, err := dyno.GetString(v, "$ref")
			if err == nil {
				refSchema := getSchema(ref)
				prop := getSchemaProperty(propName, refSchema)
				if prop != nil {
					return prop
				}
			}
		}
	}

	return nil
}

// IdentityField represents a single field in a multi-field identity-based MoRef.
type IdentityField struct {
	Name  string
	Value string // scalar value, or empty if Ref is set
	Ref   *MoRef // nested MoRef reference (for relationship fields)
}

type MoRef struct {
	Moid             string // When set, this is a direct Moid reference (no API lookup needed)
	Filter           string
	RelationshipType string
	Organization     string
	IdentityFields   []IdentityField // for multi-field identity-based MoRefs
}

// splitBalancedCommas splits s at commas that are not inside square brackets.
func splitBalancedCommas(s string) []string {
	var parts []string
	depth := 0
	start := 0
	for i, ch := range s {
		switch ch {
		case '[':
			depth++
		case ']':
			depth--
		case ',':
			if depth == 0 {
				parts = append(parts, s[start:i])
				start = i + 1
			}
		}
	}
	parts = append(parts, s[start:])
	return parts
}

// extractBracketedContent extracts the content between the first '[' and its matching ']' in s.
// Returns content and whether it was found.
func extractBracketedContent(s string) (string, bool) {
	start := strings.Index(s, "[")
	if start == -1 {
		return "", false
	}
	depth := 0
	for i := start; i < len(s); i++ {
		switch s[i] {
		case '[':
			depth++
		case ']':
			depth--
			if depth == 0 {
				return s[start+1 : i], true
			}
		}
	}
	return "", false
}

// parseMultiFieldMoRef attempts to parse a multi-field identity MoRef of the form
// MoRef:type[K1:V1,K2:V2,...] where values may themselves be MoRef strings.
// Returns nil if the string does not match this format.
func parseMultiFieldMoRef(moref string) *MoRef {
	// Must start with MoRef: followed by a type
	r := regexp.MustCompile(`^MoRef:([\w\.]+)\[`)
	m := r.FindStringSubmatch(moref)
	if m == nil {
		return nil
	}
	relType := canonicaliseRelationshipType(m[1])

	content, ok := extractBracketedContent(moref)
	if !ok {
		return nil
	}

	// Split content at top-level commas
	parts := splitBalancedCommas(content)

	// Must have at least 2 parts to be a multi-field MoRef
	// (single-field cases are handled by the existing regex patterns)
	if len(parts) < 2 {
		return nil
	}

	var fields []IdentityField
	for _, part := range parts {
		// Split at first colon to get key:value
		idx := strings.Index(part, ":")
		if idx == -1 {
			return nil // malformed
		}
		key := part[:idx]
		val := part[idx+1:]

		field := IdentityField{Name: key}
		if strings.HasPrefix(val, "MoRef") {
			nested := CanonicaliseMoRef(val, "")
			if nested == nil {
				return nil
			}
			field.Ref = nested
		} else {
			field.Value = val
		}
		fields = append(fields, field)
	}

	return &MoRef{
		RelationshipType: relType,
		IdentityFields:   fields,
	}
}

func canonicaliseRelationshipType(rt string) string {
	r := regexp.MustCompile(`^([[:upper:]]+[[:lower:]]+)([[:upper:]]+[[:alpha:]]+)Relationship$`)
	m := r.FindStringSubmatch(rt)
	if m != nil {
		return fmt.Sprintf("%s.%s.Relationship", strings.ToLower(m[1]), m[2])
	}
	return rt
}

func CanonicaliseMoRef(moref string, defaultRelationshipType string) *MoRef {
	var r *regexp.Regexp
	var m []string

	r = regexp.MustCompile(`^[[:xdigit:]]{24}$`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
		if defaultRelationshipType != "" {
			return &MoRef{Moid: moref}
		}
		return nil
	}

	r = regexp.MustCompile(`MoRef\[\$filter:(.+)\]`)
	m = r.FindStringSubmatch(moref)
	if m != nil && defaultRelationshipType != "" {
		return &MoRef{
			Filter:           m[1],
			RelationshipType: defaultRelationshipType,
		}
	}

	r = regexp.MustCompile(`MoRef:([\w\.]+)\[\$filter:(.+)\]`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
		return &MoRef{
			Filter:           m[2],
			RelationshipType: canonicaliseRelationshipType(m[1]),
		}
	}

	// Multi-field identity MoRef: MoRef:type[K1:V1,K2:V2,...] — must be checked before single-field
	if mf := parseMultiFieldMoRef(moref); mf != nil {
		return mf
	}

	r = regexp.MustCompile(`MoRef:([\w\.]+)\[(\w+):([0-9A-Za-z_\-\.\s]+)\]`)

	m = r.FindStringSubmatch(moref)
	if m != nil {
		// Direct Moid reference with explicit type — no API lookup needed
		if m[2] == "Moid" && regexp.MustCompile(`^[[:xdigit:]]{24}$`).MatchString(m[3]) {
			return &MoRef{
				Moid:             m[3],
				RelationshipType: canonicaliseRelationshipType(m[1]),
			}
		}
		return &MoRef{
			Filter:           fmt.Sprintf("%s eq '%s'", m[2], m[3]),
			RelationshipType: canonicaliseRelationshipType(m[1]),
		}
	}

	r = regexp.MustCompile(`MoRef\[(\w+):([0-9A-Za-z_\-\.\s]+)\]`)

	m = r.FindStringSubmatch(moref)
	if m != nil {
		// Direct Moid reference without explicit type — no API lookup needed
		if m[1] == "Moid" && regexp.MustCompile(`^[[:xdigit:]]{24}$`).MatchString(m[2]) {
			return &MoRef{Moid: m[2]}
		}
		if defaultRelationshipType != "" {
			return &MoRef{
				Filter:           fmt.Sprintf("%s eq '%s'", m[1], m[2]),
				RelationshipType: defaultRelationshipType,
			}
		}
	}

	r = regexp.MustCompile(`^MoRef:([\w\.]+)\[([0-9A-Za-z_\-\.\s]+)\]`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
		return &MoRef{
			Filter:           fmt.Sprintf("Name eq '%s'", m[2]),
			RelationshipType: canonicaliseRelationshipType(m[1]),
		}
	}

	r = regexp.MustCompile(`^MoRef\[([[:xdigit:]]{24})\]$`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
		return &MoRef{Moid: m[1]}
	}

	r = regexp.MustCompile(`^MoRef\[([0-9A-Za-z_\-\.\s]+)\]`)
	m = r.FindStringSubmatch(moref)
	if m != nil && defaultRelationshipType != "" {
		return &MoRef{
			Filter:           fmt.Sprintf("Name eq '%s'", m[1]),
			RelationshipType: defaultRelationshipType,
		}
	}

	// MoRef:ntp.Policy[default\test]
	r = regexp.MustCompile(`^MoRef:([\w\.]+)\[([0-9A-Za-z_\-\.]+)\\([0-9A-Za-z_\-\.\s]+)\]`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
		return &MoRef{
			Filter:           fmt.Sprintf("Name eq '%s'", m[3]),
			RelationshipType: canonicaliseRelationshipType(m[1]),
			Organization:     m[2],
		}
	}

	// MoRef[default\test]
	r = regexp.MustCompile(`^MoRef\[([0-9A-Za-z_\-\.]+)\\([0-9A-Za-z_\-\.\s]+)\]`)
	m = r.FindStringSubmatch(moref)
	if m != nil && defaultRelationshipType != "" {
		return &MoRef{
			Filter:           fmt.Sprintf("Name eq '%s'", m[2]),
			RelationshipType: defaultRelationshipType,
			Organization:     m[1],
		}
	}

	r = regexp.MustCompile(`^\s*([0-9A-Za-z_\-\.\s]+)\s*$`)
	m = r.FindStringSubmatch(moref)
	if m != nil && defaultRelationshipType != "" {
		return &MoRef{
			Filter:           fmt.Sprintf("Name eq '%s'", m[1]),
			RelationshipType: defaultRelationshipType,
		}
	}

	return nil
}

func canonicaliseMoRefs(o *map[string]any, s map[string]any) {
	relationshipRegExp := regexp.MustCompile(`/?([0-9A-Za-z\.]*\.Relationship)$`)

	for propName, val := range *o {
		prop := getSchemaProperty(propName, s)

		// is the property a relationship?
		if ref, err := dyno.GetString(prop, "$ref"); err == nil {
			if m := relationshipRegExp.FindStringSubmatch(ref); len(m) == 2 {
				// if val is a string we will attempt to annotate it, otherwise leave untouched
				if val, ok := val.(string); ok {
					(*o)[propName] = CanonicaliseMoRef(val, m[1])
				}
			}

		} else {
			if typ, err := dyno.GetString(prop, "type"); err == nil {
				if typ == "array" {
					if ref, err := dyno.GetString(prop, "items", "$ref"); err == nil {
						if val, ok := val.([]any); ok {
							for i, v := range val {
								if v, ok := v.(string); ok {
									val[i] = CanonicaliseMoRef(v, SchemaNameToClassId(ref))
								}
							}
						}
					} else {
						log.Trace("canonicaliseMoRefs: unsupported array")
					}
				}
			}
		}
	}
}

func canonicaliseMoRefsWithoutSchema(body any) any {
	switch body := body.(type) {
	case map[string]any:
		for k, v := range body {
			new := canonicaliseMoRefsWithoutSchema(v)
			body[k] = new
		}
		return body
	case []any:
		for i, v := range body {
			new := canonicaliseMoRefsWithoutSchema(v)
			body[i] = new
		}
		return body
	case string:
		new := CanonicaliseMoRef(body, "")
		if new != nil {
			return new
		}
	}

	return body
}

func CanonicaliseMoRefs(o *map[string]any, baseSchemaName string) {
	schema := getSchema(baseSchemaName)
	if schema == nil {
		return
	}

	canonicaliseMoRefs(o, schema)

	// Do a secondary pass to force expand any remaining "MoRef:<type>[target]" refs
	newO := canonicaliseMoRefsWithoutSchema(*o)
	if newO, ok := newO.(map[string]any); ok {
		*o = newO
	}
}

var schemaNameRegexp = regexp.MustCompile(`^#/components/schemas/(.*)$`)

func SchemaNameToClassId(ref string) string {
	// If this is a fully qualified schema path just strip out the name
	m := schemaNameRegexp.FindStringSubmatch(ref)
	if m != nil {
		return m[1]
	}
	return ref
}

func getSchema(schemaName string) map[string]any {
	schemaName = SchemaNameToClassId(schemaName)

	s, err := dyno.Get(lazyLoadSpec(), "components", "schemas", schemaName)
	if err != nil {
		log.Debug("schema not found", "schema_name", schemaName)
		return nil
	}

	if s, ok := s.(map[string]any); ok {
		return s
	}

	return nil
}

func ClassIdHasProperty(classId, propertyName string) bool {
	schema := getSchema(classId)
	if schema == nil {
		return false
	}

	return getSchemaProperty(propertyName, schema) != nil
}

// GetSchemaExample returns a JSON example/template for the given schema property.
func GetSchemaExample(propSchema map[string]any) (string, error) {
	var refSchemaName string
	isArray := false

	if ref, err := dyno.GetString(propSchema, "$ref"); err == nil {
		refSchemaName = ref
	} else if dtStr, err := dyno.GetString(propSchema, "type"); err == nil && dtStr == "array" {
		if arrayTypeRef, err := dyno.GetString(propSchema, "items", "$ref"); err == nil {
			refSchemaName = arrayTypeRef
			isArray = true
		}
	}

	if refSchemaName == "" {
		return "", fmt.Errorf("property is not a complex type or array of complex type")
	}

	if strings.HasSuffix(refSchemaName, ".Relationship") {
		return "", nil
	}

	schema := getSchema(refSchemaName)
	if schema == nil {
		return "", fmt.Errorf("schema not found: %s", refSchemaName)
	}

	example, err := generateExample(schema)
	if err != nil {
		return "", err
	}

	if len(example) == 0 {
		return "", nil
	}

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	if isArray {
		if err := encoder.Encode([]any{example}); err != nil {
			return "", err
		}
	} else {
		if err := encoder.Encode(example); err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func generateExample(schema map[string]any) (map[string]any, error) {
	example := make(map[string]any)
	properties := getAllProperties(schema)

	for name, prop := range properties {
		if name == "ClassId" || name == "ObjectType" {
			continue
		}

		// Check for specific type or ref
		if ref, err := dyno.GetString(prop, "$ref"); err == nil {
			// It's a reference to another schema
			subSchemaName := ref
			subSchema := getSchema(subSchemaName)
			if subSchema != nil {
				subExample, err := generateExample(subSchema)
				if err == nil {
					example[name] = subExample
					continue
				}
			}
		}

		if desc, err := dyno.GetString(prop, "description"); err == nil {
			example[name] = fmt.Sprintf("<%s>", desc)
		} else {
			example[name] = "<No description>"
		}
	}
	return example, nil
}

func getAllProperties(schema map[string]any) map[string]any {
	props := make(map[string]any)

	// Check for "properties" directly in the schema
	if p, err := dyno.GetMapS(schema, "properties"); err == nil {
		for k, v := range p {
			props[k] = v
		}
	}

	// Check "allOf" for inherited properties
	if allOf, err := dyno.GetSlice(schema, "allOf"); err == nil {
		for _, v := range allOf {
			var subSchema map[string]any

			if s, ok := v.(map[string]any); ok {
				// Check if it's a ref
				if ref, err := dyno.GetString(s, "$ref"); err == nil {
					subSchema = getSchema(ref)
				} else {
					subSchema = s
				}
			}

			if subSchema != nil {
				subProps := getAllProperties(subSchema)
				for k, v := range subProps {
					// We prioritize properties from the base schema, but if they are already present
					// (from the child schema, or a previous allOf), we keep the existing one.
					// This assumes that the child properties are processed before allOf,
					// or that we want to keep the first occurrence.
					// In JSON Schema, if properties are defined in both, they are both validated.
					// Here we just want a map of all possible properties.
					if _, exists := props[k]; !exists {
						props[k] = v
					}
				}
			}
		}
	}
	return props
}
