package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"encoding/json"

	"github.com/PaesslerAG/jsonpath"
	"github.com/bndr/gotabulate"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

func logHTTPResponse(httpResponse *http.Response) {
	if httpResponse == nil {
		return
	}

	log.Printf("HTTP Response: %d %v", httpResponse.StatusCode, httpResponse.Status)
}

func resultHandler(result interface{}, httpResponse *http.Response, err error) {
	if err != nil || verbose {
		logHTTPResponse(httpResponse)
	}

	if err != nil {
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
		printResultDefault(result)
	}
}

var uninterestingAttributes = []string{
	//From MoBaseMo
	"AccountMoid",
	"ClassId",
	"CreateTime",
	"DomainGroupMoid",
	"ModTime",
	"ObjectType",
	"Owners",
	"SharedScope",
	"Tags",
	"VersionContext",
	"Ancestors",
	"Parent",
	"PermissionResources",
	"DisplayNames",

	//From PolicyAbstractPolicy
	"Description", // include/exclude this? TBD?

	//Others, to be validated
	"ApplianceAccount",
	"Organization",
	"Profiles",
}

func removeUninterestingAttributes(mo *map[string]interface{}) {
	for _, attr := range uninterestingAttributes {
		delete(*mo, attr)
	}
}

func collapseReferences(mo *map[string]interface{}) {
	for k, v := range *mo {
		if in, ok := v.(map[string]interface{}); ok {
			if classID, ok := in["ClassId"]; ok && classID == "mo.MoRef" {
				(*mo)[k] = fmt.Sprintf("MoRef[%v/%v]", in["ObjectType"], in["Moid"])
			}
		}
	}
}

// simplifyResult prepares alters the structure of returned data to be appropriate for default(human) output
//  Specifically, this means unwrapping returned lists (if of form result[XX]["Results"] -> []interface{},
//  just keep the []interface) and removing boilerplate attributes from the MOs
func simplifyResult(result interface{}) interface{} {
	var ret interface{}

	ret = result

	// see if input is of form result[XX]["Results"] -> []interface{},
	//  if so, just keep the []interface{}
	if m, ok := result.(map[string]interface{}); ok {
		// if the input is a map of strings, go through each item to see if it is also a map:
		for _, v := range m {
			if m2, ok := v.(map[string]interface{}); ok {
				if r, ok := m2["Results"]; ok {
					if res, ok := r.([]interface{}); ok {
						ret = res
					}
				}
			}
		}
	}

	// remove uninteresting attributes from individual mo
	if m, ok := ret.(map[string]interface{}); ok {
		removeUninterestingAttributes(&m)
		collapseReferences(&m)
	}

	// remove uninteresting attributes from list of mos
	if li, ok := ret.([]interface{}); ok {
		for _, i := range li {
			if m, ok := i.(map[string]interface{}); ok {
				removeUninterestingAttributes(&m)
				collapseReferences(&m)
			}
		}

	}

	return ret
}

func stringify(in interface{}) string {
	if in, ok := in.(bool); ok {
		if in {
			return "True"
		}

		return "False"
	}

	if in, ok := in.([]interface{}); ok {
		l := []string{}
		for _, v := range in {
			l = append(l, stringify(v))
		}
		return strings.Join(l, ", ")
	}

	return fmt.Sprintf("%v", in)
}

func prepareResultTable(in interface{}) ([][]string, []string) {
	outHeaders := []string{}
	outData := [][]string{}

	// If in is a map, make it a 1 length slice
	if _, ok := in.(map[string]interface{}); ok {
		in = []interface{}{
			in,
		}
	}

	if inList, ok := in.([]interface{}); ok {
		if len(inList) < 1 {
			// Empty list, return empty
			return outData, outHeaders
		}

		// Get the headers from the first element
		firstRow := inList[0]
		if rowMap, ok := firstRow.(map[string]interface{}); ok {
			for k := range rowMap {
				outHeaders = append(outHeaders, k)
			}
		}

		for _, row := range inList {
			if rowMap, ok := row.(map[string]interface{}); ok {
				newRow := []string{}
				for _, k := range outHeaders {
					newRow = append(newRow, stringify(rowMap[k]))
				}
				outData = append(outData, newRow)
			}
		}

	}

	return outData, outHeaders
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

const defaultOutputMaxColumns int = 8

func printResultDefault(result interface{}) {
	result = simplifyResult(result)

	tableData, tableHeaders := prepareResultTable(result)

	// Pretty rough but if the output will be very wide fall back to YAML formatting the output
	if len(tableHeaders) > defaultOutputMaxColumns {
		log.Println("Too many columns for table format, falling back to vertical output. NOTE: this is not valid YAML; use --output yaml to get valid YAML.")
		printResultYAML(result)
		return
	}

	// If the result is just 1 item also fall back to YAML but we
	if len(tableData) == 1 {
		log.Println("Single result, falling back to vertical output. NOTE: this is not valid YAML; use --output yaml to get valid YAML.")
		printResultYAML(result)
		return
	}

	t := gotabulate.Create(tableData)
	t.SetHeaders(tableHeaders)

	fmt.Println(t.Render("simple"))
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
	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	fmt.Println(string(out))
}
