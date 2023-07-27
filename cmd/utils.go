package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"regexp"

	"github.com/icza/dyno"
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

func readBody(bodyFormat string, bodyParamMap interface{}) error {
	if bodyFormat == "json" {
		// Gather body from JSON on stdin.
		fmt.Println("Waiting for JSON body: ")
		err := json.NewDecoder(os.Stdin).Decode(bodyParamMap)
		if err != nil {
			return fmt.Errorf("Decoding JSON: %v", err)
		}

		log.Tracef("After JSON parse, bodyParamMap: %v", bodyParamMap)
	} else if bodyFormat == "yaml" {
		// Gather body from YAML on stdin.
		fmt.Println("Waiting for YAML body: ")
		err := yaml.NewDecoder(os.Stdin).Decode(bodyParamMap)
		if err != nil {
			return fmt.Errorf("Decoding YAML: %v", err)
		}

		log.Tracef("After YAML parse, bodyParamMap: %v", bodyParamMap)
	} else {
		return fmt.Errorf("Unknown request body format: %s", bodyFormat)
	}

	return nil
}

// Parse the MoRef string and return:
// (filter string, relationshipDataType string, ok bool)
func parseMoRef(moref string) (string, string, bool) {
	log.Tracef("Parsing moRef %s", moref)
	var r *regexp.Regexp
	var m []string

	r = regexp.MustCompile(`MoRef\[\$filter:(.+)\]`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
		return m[1], "", true
	}

	r = regexp.MustCompile(`MoRef:(\w+)\[\$filter:(.+)\]`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
		return m[2], m[1], true
	}

	r = regexp.MustCompile(`MoRef:(\w+)\[(\w+):([0-9A-Za-z_\-\.]+)\]`)

	m = r.FindStringSubmatch(moref)
	if m != nil {
		return fmt.Sprintf("%s eq '%s'", m[2], m[3]), m[1], true
	}

	r = regexp.MustCompile(`MoRef\[(\w+):([0-9A-Za-z_\-\.]+)\]`)

	m = r.FindStringSubmatch(moref)
	if m != nil {
		return fmt.Sprintf("%s eq '%s'", m[1], m[2]), "", true
	}

	r = regexp.MustCompile(`^MoRef:(\w+)\[([0-9A-Za-z_\-\.]+)\]`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
		return fmt.Sprintf("Name eq '%s'", m[2]), m[1], true
	}

	r = regexp.MustCompile(`^MoRef\[([0-9A-Za-z_\-\.]+)\]`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
		return fmt.Sprintf("Name eq '%s'", m[1]), "", true
	}

	r = regexp.MustCompile(`^\s*([0-9A-Za-z_\-\.]+)\s*$`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
		return fmt.Sprintf("Name eq '%s'", m[1]), "", true
	}

	return "", "", false
}

func setMoMoRefByName(client *isctlClient, relationship string, name string) map[string]any {
	return setMoMoRefByFilter(client, relationship, fmt.Sprintf("Name eq '%s'", name))
}

func getReferredTypeName(t reflect.Type) string {
	var referredTypeName string
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == "MoMoRef" {
			continue
		}
		referredTypeName = t.Field(i).Type.Elem().Name()
	}

	return referredTypeName
}

func setMoMoRefByFilter(client *isctlClient, relationship string, filter string) map[string]any {

	log.Debugf("Looking up MoMoRef %s with filter %s", relationship, filter)

	moref := map[string]any{
		"ClassId": "mo.MoRef",
	}
	op := getOperationForRelationship(relationship)
	if op == nil {
		log.Fatalf("FATAL: No operation for relationship: %s", relationship)
	}
	res, _, err := op.Execute(client, nil, map[string]string{"filter": filter})
	if err != nil {
		log.Errorf("Error executing lookup query: %v", err)
		return nil
	}

	moid, ok := getMoid(res)
	if !ok {
		return nil
	}
	classId, ok := getClassId(res)
	if !ok {
		return nil
	}

	log.Debugf("Got Moid and ClassId: %s, %s", moid, classId)

	moref["Moid"] = moid
	moref["ObjectType"] = classId

	return moref
}

// getMoid takes a "<objecttype>.List" structure, checks there was exactly 1 match and returns the Moid of that match
func getMoid(res interface{}) (string, bool) {
	results, err := dyno.GetSlice(res, "Results")
	if err != nil {
		log.Tracef("getMoid: no Results field (original res: %+v)", res)
		return "", false
	}

	if len(results) != 1 {
		log.Tracef("getMoid: number of results doesn't exactly equal 1")
		return "", false
	}

	moid, err := dyno.GetString(results, 0, "Moid")
	if err != nil {
		log.Tracef("getMoid: single result doesn't have Moid field")
		return "", false
	}

	return moid, true
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

func safeStringP(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

type ComplexValue struct {
	js string
}

func (v *ComplexValue) String() string {
	if v == nil {
		return ""
	}
	return v.js
}

func (v *ComplexValue) Set(s string) error {
	v.js = s
	return nil
}

func (v *ComplexValue) Type() string {
	return "json"
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

func relationshipToIntersightClassId(relationship string) string {
	// Strip "Relationship" at the end if needed
	r := regexp.MustCompile(`Relationship$`)
	relationship = r.ReplaceAllString(relationship, "")

	// Strip "[]" at the start if needed
	r2 := regexp.MustCompile(`^\[\]`)
	relationship = r2.ReplaceAllString(relationship, "")

	if cls, ok := goTypeNameToIntersightClassID[relationship]; ok {
		return cls
	}

	return ""
}
