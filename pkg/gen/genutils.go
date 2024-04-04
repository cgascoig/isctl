package gen

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/icza/dyno"
	log "github.com/sirupsen/logrus"

	"github.com/cgascoig/isctl/pkg/oapi"
	"github.com/cgascoig/isctl/pkg/util"
)

var momorefCache = map[oapi.MoRef]map[string]any{}

func GetMoMoRef(client *util.IsctlClient, moref *oapi.MoRef) (map[string]any, error) {
	log.Debugf("Looking up Mo by MoRef %v", *moref)

	if mo, ok := momorefCache[*moref]; ok {
		log.Trace("Returning MoMoRef from cache")
		return mo, nil
	}

	op := GetOperationForRelationship(moref.RelationshipType)
	if op == nil {
		return nil, fmt.Errorf("no operation for relationship %s", moref.RelationshipType)
	}

	res, err := op.Execute(client, nil, map[string]string{"filter": moref.Filter})
	if err != nil {
		return nil, fmt.Errorf("error executing lookup query: %v", err)
	}

	moid, ok := util.GetMoid(res)
	if !ok {
		return nil, fmt.Errorf("GetMoMoRef: unable to get moid")
	}

	classId, ok := getClassId(res)
	if !ok {
		return nil, fmt.Errorf("GetMoMoRef: ubable to get classID")
	}

	log.Debugf("Got Moid and ClassId: %s, %s", moid, classId)

	ret := map[string]any{
		"ClassId":    "mo.MoRef",
		"Moid":       moid,
		"ObjectType": classId,
	}

	momorefCache[*moref] = ret

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
