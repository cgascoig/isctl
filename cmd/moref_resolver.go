package main

import (
	"fmt"
	"strings"

	"github.com/icza/dyno"
	log "github.com/sirupsen/logrus"

	"github.com/cgascoig/isctl/pkg/gen"
	"github.com/cgascoig/isctl/pkg/oapi"
	"github.com/cgascoig/isctl/pkg/util"
)

// MoRefResolver resolves MoRef Moids to human-readable identity strings
// using identity constraints from meta.json.
type MoRefResolver struct {
	client *util.IsctlClient
	meta   *oapi.Meta
	cache  map[string]string // "objectType/moid" -> resolved identity string
}

// NewMoRefResolver creates a new MoRefResolver.
func NewMoRefResolver(client *util.IsctlClient) (*MoRefResolver, error) {
	meta, err := oapi.GetMeta()
	if err != nil {
		return nil, fmt.Errorf("failed to load metadata: %w", err)
	}
	return &MoRefResolver{
		client: client,
		meta:   meta,
		cache:  make(map[string]string),
	}, nil
}

// ResolveResult recursively walks result (map[string]any or []any), finds mo.MoRef
// maps, and annotates them with a "_ResolvedIdentity" key containing the
// identity-based representation. Falls back to Moid if resolution fails.
func (r *MoRefResolver) ResolveResult(result any) any {
	switch v := result.(type) {
	case map[string]any:
		if classId, ok := v["ClassId"].(string); ok && classId == "mo.MoRef" {
			if objectType, ok := v["ObjectType"].(string); ok {
				if moid, ok := v["Moid"].(string); ok {
					if identity, ok := r.resolveIdentity(objectType, moid); ok {
						v["_ResolvedIdentity"] = identity
					}
				}
			}
			return v
		}
		for k, val := range v {
			v[k] = r.ResolveResult(val)
		}
		return v
	case []any:
		for i, item := range v {
			v[i] = r.ResolveResult(item)
		}
		return v
	default:
		return result
	}
}

// resolveIdentity resolves a Moid to an identity string for the given objectType.
// Returns ("", false) if the class has no identity constraints or resolution fails.
func (r *MoRefResolver) resolveIdentity(objectType, moid string) (string, bool) {
	cacheKey := objectType + "/" + moid
	if cached, ok := r.cache[cacheKey]; ok {
		return cached, true
	}

	constraints := r.meta.GetIdentityConstraints(objectType)
	if len(constraints) == 0 {
		return "", false
	}

	// Fetch the object to get its identity field values
	op := gen.GetGetOperationForClassID(objectType)
	if op == nil {
		log.Debugf("readable-morefs: no GET operation for %s, falling back to Moid", objectType)
		return "", false
	}

	// Build $select with identity fields + Moid
	selectFields := append([]string{"Moid"}, constraints...)
	res, err := op.Execute(r.client, nil, map[string]string{
		"filter": fmt.Sprintf("Moid eq '%s'", moid),
		"select": strings.Join(selectFields, ","),
	})
	if err != nil {
		log.Debugf("readable-morefs: error fetching %s/%s: %v", objectType, moid, err)
		return "", false
	}

	// Extract the single result
	results, err := dyno.GetSlice(res, "Results")
	if err != nil || len(results) != 1 {
		log.Debugf("readable-morefs: expected 1 result for %s/%s, got error or wrong count", objectType, moid)
		return "", false
	}
	obj, ok := results[0].(map[string]any)
	if !ok {
		return "", false
	}

	// Build the identity string from constraint fields
	var parts []string
	for _, field := range constraints {
		// Skip Account (matches buildIdentityFilter behaviour in apply.go)
		if field == "Account" {
			continue
		}

		_, isRef := r.meta.GetRefType(objectType, field)
		if isRef {
			// Relationship field — resolve it recursively
			refMap, err := dyno.GetMapS(obj, field)
			if err != nil {
				log.Debugf("readable-morefs: could not get relationship field %s: %v", field, err)
				continue
			}
			refMoid, ok1 := refMap["Moid"].(string)
			refObjectType, ok2 := refMap["ObjectType"].(string)
			if !ok1 || !ok2 {
				continue
			}
			// Try to recursively resolve the referenced object's identity
			innerIdentity, innerOk := r.resolveIdentity(refObjectType, refMoid)
			if innerOk {
				parts = append(parts, fmt.Sprintf("%s:MoRef:%s[%s]", field, refObjectType, innerIdentity))
			} else {
				parts = append(parts, fmt.Sprintf("%s:MoRef:%s[Moid:%s]", field, refObjectType, refMoid))
			}
		} else {
			// Scalar field
			val, err := dyno.Get(obj, field)
			if err != nil {
				log.Debugf("readable-morefs: could not get field %s: %v", field, err)
				continue
			}
			parts = append(parts, fmt.Sprintf("%s:%v", field, val))
		}
	}

	if len(parts) == 0 {
		return "", false
	}

	identity := strings.Join(parts, ",")
	r.cache[cacheKey] = identity
	return identity, true
}
