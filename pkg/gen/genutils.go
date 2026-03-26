package gen

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"sync"

	"github.com/icza/dyno"
	log "github.com/sirupsen/logrus"

	"github.com/cgascoig/isctl/pkg/oapi"
	"github.com/cgascoig/isctl/pkg/util"
)

var (
	momorefCache      = map[string]map[string]any{}
	momorefCacheMutex sync.RWMutex
)

func GetMoMoRef(client *util.IsctlClient, moref *oapi.MoRef) (map[string]any, error) {
	log.Debugf("Looking up Mo by MoRef %v", *moref)

	if moref.Moid != "" {
		ret := map[string]any{
			"ClassId": "mo.MoRef",
			"Moid":    moref.Moid,
		}
		if moref.RelationshipType != "" {
			ret["ObjectType"] = getClassIDFromRelationship(moref.RelationshipType)
		}
		return ret, nil
	}

	// Handle multi-field identity MoRefs by building a compound filter
	if len(moref.IdentityFields) > 0 {
		parts := []string{}
		for _, field := range moref.IdentityFields {
			if field.Ref != nil {
				resolved, err := GetMoMoRef(client, field.Ref)
				if err != nil {
					return nil, fmt.Errorf("error resolving identity field %s: %v", field.Name, err)
				}
				moid, err := dyno.GetString(resolved, "Moid")
				if err != nil {
					return nil, fmt.Errorf("error getting Moid for identity field %s: %v", field.Name, err)
				}
				parts = append(parts, fmt.Sprintf("%s/Moid eq '%s'", field.Name, moid))
			} else {
				parts = append(parts, fmt.Sprintf("%s eq '%s'", field.Name, field.Value))
			}
		}
		// Build a modified moref with the compound filter and fall through to filter-based resolution
		moref = &oapi.MoRef{
			Filter:           strings.Join(parts, " and "),
			RelationshipType: moref.RelationshipType,
		}
	}

	momorefCacheMutex.RLock()
	mo, ok := momorefCache[fmt.Sprintf("%v", *moref)]
	momorefCacheMutex.RUnlock()
	if ok {
		log.Trace("Returning MoMoRef from cache")
		return mo, nil
	}

	filter := moref.Filter

	// Here we look up the moid of the organisation if the MoRef has the organization set
	if moref.Organization != "" {
		orgMoRef := oapi.CanonicaliseMoRef(moref.Organization, "organization.Organization.Relationship")
		resolvedOrgMoRef, err := GetMoMoRef(client, orgMoRef)
		if err != nil {
			return nil, fmt.Errorf("error finding organization: %v", err)
		}

		orgMoid, err := dyno.GetString(resolvedOrgMoRef, "Moid")
		if err != nil {
			return nil, fmt.Errorf("error finding organization: %v", err)
		}

		filter = fmt.Sprintf("%s and Organization/Moid eq '%s'", filter, orgMoid)
	}

	op := GetOperationForRelationship(moref.RelationshipType)
	if op == nil {
		// The relationship type has no direct GET operation - check if it's an abstract class
		ret, err := resolveAbstractMoRef(client, moref, filter)
		if err != nil {
			return nil, err
		}
		momorefCacheMutex.Lock()
		momorefCache[fmt.Sprintf("%v", *moref)] = ret
		momorefCacheMutex.Unlock()
		return ret, nil
	}

	res, err := op.Execute(client, nil, map[string]string{"filter": filter})
	if err != nil {
		return nil, fmt.Errorf("error executing lookup query: %v", err)
	}

	moid, ok := util.GetMoid(res)
	if !ok {
		return nil, fmt.Errorf("GetMoMoRef: unable to get moid")
	}

	classId, ok := getClassId(res)
	if !ok {
		return nil, fmt.Errorf("GetMoMoRef: unable to get classID")
	}

	log.Debugf("Got Moid and ClassId: %s, %s", moid, classId)

	ret := map[string]any{
		"ClassId":    "mo.MoRef",
		"Moid":       moid,
		"ObjectType": classId,
	}

	momorefCacheMutex.Lock()
	momorefCache[fmt.Sprintf("%v", *moref)] = ret
	momorefCacheMutex.Unlock()

	return ret, nil
}

// TODO: Refactor this to remove duplicate code in getMoid
// getClassId takes a "<objecttype>.List" structure, checks there was exactly 1 match and returns the ClassId of that match
func getClassId(res interface{}) (string, bool) {
	log.Tracef("getClassId for: %#v", res)

	resMap, ok := res.(map[string]any)
	if !ok {
		log.Tracef("getClassId: res is not map[string]any")
		return "", false
	}

	resList, err := dyno.GetSlice(resMap, "Results")
	if err != nil {
		log.Tracef("getClassId: res does not contain list of results: %v", err)
		return "", false
	}

	if len(resList) != 1 {
		log.Tracef("getClassId: res does not contain list of exactly 1 item")
		return "", false
	}

	classId, err := dyno.GetString(resList, 0, "ClassId")
	if err != nil {
		log.Tracef("getClassId: could not get ClassId from first item: %v", err)
		return "", false
	}

	return classId, true
}

// resolveAbstractMoRef attempts to resolve a MoRef for an abstract class by
// trying all concrete implementations. Returns an error if zero or multiple
// matches are found across different concrete classes.
func resolveAbstractMoRef(client *util.IsctlClient, moref *oapi.MoRef, filter string) (map[string]any, error) {
	classID := getClassIDFromRelationship(moref.RelationshipType)
	log.Debugf("No direct operation for %s, checking if it's an abstract class", classID)

	m, err := oapi.GetMeta()
	if err != nil {
		return nil, fmt.Errorf("no operation for relationship %s (failed to load metadata: %v)", moref.RelationshipType, err)
	}

	if m.IsConcreteClass(classID) {
		return nil, fmt.Errorf("no operation for relationship %s", moref.RelationshipType)
	}

	implementations := m.GetConcreteImplementations(classID)
	if len(implementations) == 0 {
		return nil, fmt.Errorf("no operation for relationship %s (no concrete implementations found for %s)", moref.RelationshipType, classID)
	}

	log.Debugf("Found %d concrete implementations of %s: %v", len(implementations), classID, implementations)

	type candidate struct {
		moid    string
		classId string
	}

	var candidates []candidate
	var triedClasses []string

	for _, implClass := range implementations {
		op := GetGetOperationForClassID(implClass)
		if op == nil {
			log.Debugf("No GET operation for concrete class %s, skipping", implClass)
			continue
		}

		triedClasses = append(triedClasses, implClass)

		res, err := op.Execute(client, nil, map[string]string{"filter": filter})
		if err != nil {
			log.Debugf("Error querying concrete class %s: %v, skipping", implClass, err)
			continue
		}

		moid, ok := util.GetMoid(res)
		if !ok {
			// Zero or multiple results for this class - skip
			log.Debugf("No unique match in concrete class %s, skipping", implClass)
			continue
		}

		cid, ok := getClassId(res)
		if !ok {
			log.Debugf("Could not get ClassId from result for %s, skipping", implClass)
			continue
		}

		candidates = append(candidates, candidate{moid: moid, classId: cid})
	}

	switch len(candidates) {
	case 0:
		return nil, fmt.Errorf("no matching object found for any concrete implementation of %s (tried: %s)", classID, strings.Join(triedClasses, ", "))
	case 1:
		log.Debugf("Resolved abstract MoRef to %s (Moid: %s)", candidates[0].classId, candidates[0].moid)
		return map[string]any{
			"ClassId":    "mo.MoRef",
			"Moid":       candidates[0].moid,
			"ObjectType": candidates[0].classId,
		}, nil
	default:
		matchedClasses := make([]string, len(candidates))
		for i, c := range candidates {
			matchedClasses[i] = c.classId
		}
		return nil, fmt.Errorf("ambiguous MoRef: matched objects in multiple concrete classes: %s — use MoRef:ClassName[NAME] to specify the concrete type", strings.Join(matchedClasses, ", "))
	}
}

func ReplaceArgs(s string, args []string) (string, error) {
	var err error = nil

	re := regexp.MustCompile(`{\w+}`)

	ret := re.ReplaceAllStringFunc(s, func(s string) string {
		if len(args) < 1 {
			err = fmt.Errorf("ReplaceArgs: insufficient args supplied")
			return ""
		}

		r := args[0]
		args = args[1:]
		return r
	})

	if err != nil {
		return "", err
	}

	return ret, nil
}

func EncodeQueryParams(queryParams map[string]string) string {
	vals := url.Values{}
	for k, v := range queryParams {
		vals.Add(fmt.Sprintf("$%s", k), v)
	}

	return vals.Encode()
}

func ExecuteWithPagination(op *Operation, batchSize int, client *util.IsctlClient, args []string, queryParams map[string]string) (any, error) {
	log.Trace("Starting ExecuteWithPagination")

	skip := 0
	var ret any

	for {
		log.Infof("Auto pagination requesting batch, skip=%d, top=%d", skip, batchSize)
		queryParams["top"] = fmt.Sprint(batchSize)
		queryParams["skip"] = fmt.Sprint(skip)
		res, err := op.Execute(client, args, queryParams)
		if err != nil {
			return nil, fmt.Errorf("error while executing query for batch skip=%d, top=%d: %v", skip, batchSize, err)
		}
		var newCount int
		ret, newCount, err = appendResults(ret, res)
		if err != nil {
			return nil, fmt.Errorf("error while executing query for batch skip=%d, top=%d: %v", skip, batchSize, err)
		}

		log.Tracef("ExecuteWithPagination finished batch, added %d results", newCount)

		if newCount == 0 {
			log.Infof("Auto pagination finished: no more results")
			return ret, nil
		}

		skip += batchSize
	}
}

// returns cur+new along with the count of new items or error
func appendResults(cur, new any) (any, int, error) {
	newS, err := dyno.GetSlice(new, "Results")
	if err != nil {
		return nil, 0, fmt.Errorf("no Results to append: %v", err)
	}

	count := len(newS)

	if cur == nil {
		return new, count, nil
	}

	curS, err := dyno.GetSlice(cur, "Results")
	if err != nil {
		return nil, 0, fmt.Errorf("no Results to append to: %v", err)
	}

	dyno.Set(cur, append(curS, newS...), "Results")
	return cur, count, nil
}
