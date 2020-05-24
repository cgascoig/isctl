package main

import (
	"fmt"
	"reflect"
	"strings"

	"encoding/json"

	"github.com/PaesslerAG/jsonpath"
	yaml "gopkg.in/yaml.v2"
)

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

func getResultObjects(result interface{}, jsonpathQuery *string) (interface{}, error) {
	buf, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("JSON marshalling error: %v", err)
	}

	var js interface{}
	err = json.Unmarshal(buf, &js)
	if err != nil {
		return nil, fmt.Errorf("JSON unmarshalling error: %v", err)
	}

	if jsonpathQuery == nil {
		return js, nil
	}

	jp, err := jsonpath.Get(*jsonpathQuery, js)
	if err != nil {
		return nil, fmt.Errorf("JSONPath parsing error: %v", err)
	}
	return jp, nil

}

func printResultYAML(result interface{}, jsonpathQuery *string) {
	res, err := getResultObjects(result, jsonpathQuery)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	out, err := yaml.Marshal(res)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	fmt.Println(string(out))
}

// func printResultYAML(result interface{}) {

// 	jp := jsonpath.New("test")

// 	err := jp.Parse("{$..}")
// 	if err != nil {
// 		fmt.Printf("JSONPath parsing error: %v", err)
// 		return
// 	}

// 	var buf bytes.Buffer

// 	err = jp.Execute(&buf, result)
// 	if err != nil {
// 		fmt.Printf("JSONPath execute error: %v", err)
// 		return
// 	}

// 	fmt.Println(buf.String())
// 	return

// 	// buf, err := json.Marshal(result)
// 	// if err != nil {
// 	// 	fmt.Printf("ERROR: %v", err)
// 	// 	return
// 	// }

// 	var js interface{}
// 	err = json.Unmarshal(buf.Bytes(), &js)
// 	if err != nil {
// 		fmt.Printf("JSON unmarshaling error: %v", err)
// 		return
// 	}

// 	y, err := yaml.Marshal(js)
// 	if err != nil {
// 		fmt.Printf("YAML marshaling error: %v", err)
// 		return
// 	}

// 	fmt.Println(string(y))
// }

func printResultJSON(result interface{}) {
	buf, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	fmt.Println(string(buf))
}
