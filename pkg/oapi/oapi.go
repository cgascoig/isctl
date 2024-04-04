package oapi

import (
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

type MoRef struct {
	Filter           string
	RelationshipType string
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
		return nil
	}

	r = regexp.MustCompile(`MoRef\[\$filter:(.+)\]`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
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

	r = regexp.MustCompile(`MoRef:([\w\.]+)\[(\w+):([0-9A-Za-z_\-\.]+)\]`)

	m = r.FindStringSubmatch(moref)
	if m != nil {
		return &MoRef{
			Filter:           fmt.Sprintf("%s eq '%s'", m[2], m[3]),
			RelationshipType: canonicaliseRelationshipType(m[1]),
		}
	}

	r = regexp.MustCompile(`MoRef\[(\w+):([0-9A-Za-z_\-\.]+)\]`)

	m = r.FindStringSubmatch(moref)
	if m != nil {
		return &MoRef{
			Filter:           fmt.Sprintf("%s eq '%s'", m[1], m[2]),
			RelationshipType: defaultRelationshipType,
		}
	}

	r = regexp.MustCompile(`^MoRef:([\w\.]+)\[([0-9A-Za-z_\-\.]+)\]`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
		return &MoRef{
			Filter:           fmt.Sprintf("Name eq '%s'", m[2]),
			RelationshipType: canonicaliseRelationshipType(m[1]),
		}
	}

	r = regexp.MustCompile(`^MoRef\[([0-9A-Za-z_\-\.]+)\]`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
		return &MoRef{
			Filter:           fmt.Sprintf("Name eq '%s'", m[1]),
			RelationshipType: defaultRelationshipType,
		}
	}

	r = regexp.MustCompile(`^\s*([0-9A-Za-z_\-\.]+)\s*$`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
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

func CanonicaliseMoRefs(o *map[string]any, baseSchemaName string) {
	schema := getSchema(baseSchemaName)
	if schema == nil {
		return
	}

	canonicaliseMoRefs(o, schema)
}

func SchemaNameToClassId(ref string) string {
	// If this is a fully qualified schema path just strip out the name
	pathRe := regexp.MustCompile(`^#/components/schemas/(.*)$`)
	m := pathRe.FindStringSubmatch(ref)
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
