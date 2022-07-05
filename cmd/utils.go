package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"regexp"

	openapi "github.com/cgascoig/intersight-go-sdk/intersight"
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

//Parse the MoRef string and return:
// (filter string, relationshipDataType string, ok bool)
func parseMoRef(moref string) (string, string, bool) {
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

func setMoMoRefByName(client *openapi.APIClient, v interface{}, relationship string, name string) bool {
	return setMoMoRefByFilter(client, v, relationship, fmt.Sprintf("Name eq '%s'", name))
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

func setMoMoRefByFilter(client *openapi.APIClient, v interface{}, relationship string, filter string) bool {
	if !isMoRef(v) {
		return false
	}

	log.Debugf("Looking up MoMoRef %s with filter %s", relationship, filter)

	moref := new(openapi.MoMoRef)
	moref.ClassId = "mo.MoRef"

	op := getOperationForRelationship(relationship)
	if op == nil {
		log.Fatalf("FATAL: No operation for relationship: %s", relationship)
	}
	res, _, err := op.Execute(client, nil, map[string]string{"filter": filter})
	if err != nil {
		log.Errorf("Error executing lookup query: %v", err)
		return false
	}

	moid, ok := getMoid(res)
	if !ok {
		return false
	}

	log.Debugf("Got Moid %s", moid)

	moref.Moid = &moid

	val := reflect.ValueOf(v).Elem()

	// Here we are trying to find the type of the field other than MoMoRef (e.g. OrganizationOrganization)
	// i.e. the type that this MoRef refers to
	t := val.Type()
	referredTypeName := getReferredTypeName(t)

	moref.ObjectType = goTypeNameToIntersightClassID[referredTypeName]

	val.FieldByName("MoMoRef").Set(reflect.ValueOf(moref))

	return true
}

// getMoid takes a "<objecttype>.List" structure, checks there was exactly 1 match and returns the Moid of that match
func getMoid(res interface{}) (string, bool) {
	val := reflect.Indirect(reflect.ValueOf(res))
	valType := val.Type()
	if valType.Kind() != reflect.Struct {
		log.Tracef("getMoid: res not struct")
		return "", false
	}

	// Find the ...List field
	r := regexp.MustCompile(`List$`)
	var listFieldName string
	for i := 0; i < valType.NumField(); i++ {
		if r.MatchString(valType.Field(i).Name) {
			listFieldName = valType.Field(i).Name
		}
	}

	if listFieldName == "" {
		log.Tracef("getMoid: no ...List field")
		return "", false
	}

	listStruct := reflect.Indirect(val.FieldByName(listFieldName))

	if listStruct.Kind() != reflect.Struct {
		log.Tracef("getMoid: ..List not struct")
		return "", false
	}

	_, ok := listStruct.Type().FieldByName("Results")
	if !ok {
		log.Tracef("getMoid: no Results field (original res: %+v)", res)
		return "", false
	}

	results := listStruct.FieldByName("Results")
	if results.Kind() != reflect.Slice && results.Kind() != reflect.Array {
		log.Tracef("getMoid: Results field not slice/array")
		return "", false
	}

	if results.Len() != 1 {
		log.Tracef("getMoid: number of results doesn't exactly equal 1")
		return "", false
	}

	result := results.Index(0)
	if result.Kind() != reflect.Struct {
		log.Tracef("getMoid: single result is not struct")
		return "", false
	}

	if _, ok := result.Type().FieldByName("Moid"); !ok {
		log.Tracef("getMoid: single result doesn't have Moid field")
		return "", false
	}

	return reflect.Indirect(result.FieldByName("Moid")).String(), true
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
