package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"encoding/json"

	"github.com/PaesslerAG/jsonpath"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

func resultHandler(result interface{}, httpResponse *http.Response, err error) {
	if err != nil {
		log.Printf("HTTP Response status text: %v", httpResponse.Status)
		log.Fatalf("ERROR: %v\n", err)
	}

	result, err = applyJSONPathFilter(result, jsonPathFilter)
	if err != nil {
		log.Fatalf("ERROR applying jsonPath filter: %v", err)
	}

	switch outputFormat := strings.ToLower(viper.GetString(keyOutputConfigKey)); {
	case outputFormat == "yaml":
		printResultYAML(result)
	case outputFormat == "json":
		printResultJSON(result)
	default:
		// printResultHuman(result)
		printResultYAML(result)
	}
}

func stripEnvelope(result interface{}) interface{} {
	vr := reflect.ValueOf(result)
	if vr.Kind() == reflect.Ptr {
		vr = vr.Elem()
	}

	if vr.Kind() == reflect.Struct {
		payloadValue := vr.FieldByName("Payload")
		if payloadValue.IsValid() {
			resultsValue := reflect.Indirect(payloadValue).FieldByName("Results")
			if resultsValue.IsValid() {
				return resultsValue.Interface()
			}
			return payloadValue.Interface()
		}
	}

	return result
}

func indent(lines []string, slice bool) []string {
	ret := []string{}

	first := true
	for _, line := range lines {
		if slice && first {
			first = false
			ret = append(ret, fmt.Sprintf("- %s", line))
		} else {
			ret = append(ret, fmt.Sprintf("  %s	", line))
		}
	}

	return ret
}

func printResultHumanValue(r reflect.Value) []string {
	vr := reflect.Indirect(r)
	ret := []string{}

	switch vr.Kind() {
	case reflect.Struct:
		for i := 0; i < vr.NumField(); i++ {
			f := vr.Field(i)
			t := vr.Type().Field(i)
			if t.PkgPath == "" { // empty PkgPath indicates exported field
				newLines := printResultHumanValue(f)
				if !t.Anonymous {
					if len(newLines) <= 1 {
						ret = append(ret, fmt.Sprintf("%s: %s", t.Name, strings.Join(newLines, "")))
					} else {
						ret = append(ret, fmt.Sprintf("%s: ", t.Name))
						ret = append(ret, indent(newLines, false)...)
					}
				} else {
					ret = append(ret, newLines...)
				}
			}
		}
	case reflect.Slice:
		for i := 0; i < vr.Len(); i++ {
			newLines := printResultHumanValue(vr.Index(i))

			ret = append(ret, indent(newLines, true)...)
		}
	default:
		ret = []string{fmt.Sprintf("%s", vr.String())}
	}

	return ret
}

func printResultHuman(result interface{}) {

	for _, line := range printResultHumanValue(reflect.ValueOf(result)) {
		fmt.Println(line)
	}
}

func applyJSONPathFilter(result interface{}, jsonpathQuery string) (interface{}, error) {
	buf, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("JSON marshalling error: %v", err)
	}

	var js interface{}
	err = json.Unmarshal(buf, &js)
	if err != nil {
		return nil, fmt.Errorf("JSON unmarshalling error: %v", err)
	}

	if jsonpathQuery == "" {
		return js, nil
	}

	jp, err := jsonpath.Get(jsonpathQuery, js)
	if err != nil {
		return nil, fmt.Errorf("JSONPath parsing error: %v", err)
	}
	return jp, nil

}

func printResultYAML(result interface{}) {
	out, err := yaml.Marshal(result)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	fmt.Println(string(out))
}

func printResultJSON(result interface{}) {
	out, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	fmt.Println(string(out))
}
