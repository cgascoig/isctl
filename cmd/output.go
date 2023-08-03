package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"regexp"
	"sort"
	"strings"

	"encoding/json"

	"github.com/PaesslerAG/jsonpath"
	"github.com/bndr/gotabulate"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
	"github.com/cgascoig/isctl/pkg/util"
)

func logHTTPResponse(httpResponse *http.Response) {
	if httpResponse == nil {
		return
	}

	log.Debugf("HTTP Response: %d %v", httpResponse.StatusCode, httpResponse.Status)
}

func resultHandler(result interface{}, httpResponse *http.Response, err error, options ...util.ResultOpt) {
	var singleResult bool

	for _, opt := range options {
		if opt.SingleResult != nil {
			singleResult = *opt.SingleResult
		}
	}

	if err != nil || verbose {
		logHTTPResponse(httpResponse)
	}

	// if there is an error try to display something helpful
	if err != nil {
		if httpResponse != nil {
			body, err2 := ioutil.ReadAll(httpResponse.Body)
			if err2 != nil {
				log.Fatalf("ERROR: %v", err)
			} else {
				log.Fatalf("ERROR: %v: %s", err, string(body))
			}
		} else {
			log.Fatalf("ERROR: %v", err)
		}
	}

	if jsonPathFilter != "" {
		log.Warn("--jsonpath is deprecated and will be removed in the future. Please use -o jsonpath=<expr>. ")
	}
	result, err = applyJSONPathFilter(result, jsonPathFilter, singleResult)
	if err != nil {
		log.Fatalf("ERROR applying jsonPath filter: %v", err)
	}

	outputConfig := viper.GetString(keyOutputConfigKey)
	outputConfigParts := strings.SplitN(outputConfig, "=", 2)

	switch outputConfigParts[0] {
	case "yaml":
		printResultYAML(result)
	case "json":
		printResultJSON(result)
	case "custom-columns":
		if len(outputConfigParts) != 2 {
			log.Fatalf("custom-columns format specified but no custom columns given")
		}
		printResultCustomColumns(result, outputConfigParts[1])
	case "csv":
		if len(outputConfigParts) == 2 {
			printResultCSV(result, outputConfigParts[1])
		} else {
			printResultCSV(result, "")
		}
	case "jsonpath":
		if jsonPathFilter != "" {
			log.Warn("Using -o jsonpath and --jsonpath together may produce unpredictable results")
		}
		if len(outputConfigParts) == 2 {
			printResultJSONPath(result, outputConfigParts[1])
		} else {
			printResultJSONPath(result, "")
		}

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

func removeWrappers(result interface{}, singleResult bool) interface{} {
	var ret interface{}

	ret = result

	// see if input is of form result[XX]["Results"] -> []interface{} or result["Results"]->[]interface{},
	//  if so, just keep the []interface{}
	if m, ok := result.(map[string]interface{}); ok {
		if r, ok := m["Results"]; ok {
			// the input is of form result["Results"]->[]interface{}
			if res, ok := r.([]interface{}); ok {
				if singleResult && len(res) == 1 {
					ret = res[0]
				} else {
					ret = res
				}
			}
		} else {
			// if the input is a map of strings, go through each item to see if it is also a map:
			for _, v := range m {
				if m2, ok := v.(map[string]interface{}); ok {
					if r, ok := m2["Results"]; ok {
						if res, ok := r.([]interface{}); ok {
							if singleResult && len(res) == 1 {
								ret = res[0]
							} else {
								ret = res
							}
						}
					}
				}
			}
		}
	}

	return ret
}

func filterAttributes(result interface{}) interface{} {
	var ret interface{}

	ret = result

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
	switch v := in.(type) {
	case bool:
		if v {
			return "true"
		}

		return "false"
	case map[string]interface{}:
		key, hasKey := v["Key"]
		value, hasValue := v["Value"]
		if hasKey && hasValue && len(v) == 2 {
			return fmt.Sprintf("%v: %v", key, value)
		}
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

func prepareResultTable(in interface{}, sortHeaders bool) ([][]string, []string) {
	outHeaders := []string{}
	outData := [][]string{}

	// If in is not a slice, make it a 1 length slice
	if _, ok := in.([]interface{}); !ok {
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

		if sortHeaders {
			sort.Slice(outHeaders, func(i, j int) bool {
				if outHeaders[i] == "Name" {
					return true
				}
				if outHeaders[j] == "Name" {
					return false
				}
				if outHeaders[i] == "Moid" {
					return true
				}
				if outHeaders[j] == "Moid" {
					return false
				}
				return outHeaders[i] < outHeaders[j]
			})
		}

		for _, row := range inList {
			if rowMap, ok := row.(map[string]interface{}); ok {
				newRow := []string{}
				for _, k := range outHeaders {
					newRow = append(newRow, stringify(rowMap[k]))
				}
				outData = append(outData, newRow)
			} else {
				outData = append(outData, []string{stringify(row)})
			}
		}

	}

	return outData, outHeaders
}

func applyJSONPathFilter(result interface{}, jsonpathQuery string, singleResult bool) (interface{}, error) {
	buf, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("JSON marshalling error: %v", err)
	}

	var js interface{}
	err = json.Unmarshal(buf, &js)
	if err != nil {
		return nil, fmt.Errorf("JSON unmarshalling error: %v", err)
	}

	js = removeWrappers(js, singleResult)

	if jsonpathQuery == "" {
		return js, nil
	}

	jp, err := jsonpath.Get(jsonpathQuery, js)
	if err != nil {
		return nil, fmt.Errorf("JSONPath parsing error: %v", err)
	}
	return jp, nil

}

const defaultOutputMaxColumns int = 10

func printResultDefault(result interface{}) {
	result = filterAttributes(result)

	tableData, tableHeaders := prepareResultTable(result, true)

	if len(tableHeaders) == 0 {
		for _, row := range tableData {
			for _, v := range row {
				fmt.Printf("%v", v)
			}
			fmt.Printf("\n")
		}
		return
	}

	// Pretty rough but if the output will be very wide fall back to YAML formatting the output unless the output format is explicitly set to "table"
	outputFormat := strings.ToLower(viper.GetString(keyOutputConfigKey))
	if (len(tableHeaders) > defaultOutputMaxColumns) && outputFormat != "table" {
		log.Println("Too many columns for table format, falling back to vertical output. NOTE: this is not valid YAML; use --output yaml to get valid YAML.")
		printResultYAML(result)
		return
	}

	// If the result is just 1 item also fall back to YAML unless table is explicitly specified
	if len(tableData) == 1 && outputFormat != "table" {
		log.Println("Single result, falling back to vertical output. NOTE: this is not valid YAML; use --output yaml to get valid YAML.")
		printResultYAML(result)
		return
	}

	t := gotabulate.Create(tableData)
	t.SetHeaders(tableHeaders)
	t.SetDenseMode()

	fmt.Println(t.Render("simple"))
}

func printResultYAML(result interface{}) {
	out, err := yaml.Marshal(result)
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	fmt.Println(string(out))
}

func printResultJSON(result interface{}) {
	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	fmt.Println(string(out))
}

func relaxedJSONPathExpression(exp string) string {
	re := regexp.MustCompile(`^(\.|\[).+`)
	if re.MatchString(exp) {
		exp = fmt.Sprintf("$%s", exp)
	}

	return exp
}

func prepareResultTableCustomColumns(result interface{}, template string) ([][]string, []string) {
	// If in is not a slice, make it a 1 length slice
	if _, ok := result.([]interface{}); !ok {
		result = []interface{}{
			result,
		}
	}

	type columnSpec struct {
		Name string
		Spec string
	}

	columns := []columnSpec{}
	tableData := [][]string{}
	tableHeaders := []string{}

	if template == "" {
		// If there is no template, get all the column spec from the first object returned

		if inList, ok := result.([]interface{}); ok {
			// Get the headers from the first element
			firstRow := inList[0]
			if rowMap, ok := firstRow.(map[string]interface{}); ok {
				for k := range rowMap {
					columns = append(columns, columnSpec{Name: k, Spec: fmt.Sprintf("$.%s", k)})
				}
			}
		}

		sort.Slice(columns, func(i, j int) bool {
			if columns[i].Name == "Name" {
				return true
			}
			if columns[j].Name == "Name" {
				return false
			}
			if columns[i].Name == "Moid" {
				return true
			}
			if columns[j].Name == "Moid" {
				return false
			}
			return columns[i].Name < columns[j].Name
		})

		for _, col := range columns {
			tableHeaders = append(tableHeaders, col.Name)
		}
	} else {
		// create the columns spec from the template
		columnTmpls := strings.Split(template, ",")

		for _, columnTmpl := range columnTmpls {
			columnSpecSplit := strings.SplitN(columnTmpl, ":", 2)
			if len(columnSpecSplit) != 2 {
				log.Fatalf("unexpected custom-columns spec: %s, expected <header>:<json-path-expr>", columnTmpl)
			}
			columns = append(columns, columnSpec{Name: columnSpecSplit[0], Spec: relaxedJSONPathExpression(columnSpecSplit[1])})
			tableHeaders = append(tableHeaders, columnSpecSplit[0])
		}
	}

	if len(columns) < 1 {
		log.Fatal("no columns for output")
	}

	if inList, ok := result.([]interface{}); ok {
		for _, row := range inList {
			rowData := []string{}
			for _, column := range columns {
				coldata, err := jsonpath.Get(column.Spec, row)
				if err != nil {
					coldata = "<none>"
				}
				rowData = append(rowData, stringify(coldata))
			}
			tableData = append(tableData, rowData)
		}
	}

	return tableData, tableHeaders
}

func printResultCustomColumns(result interface{}, template string) {

	tableData, tableHeaders := prepareResultTableCustomColumns(result, template)
	t := gotabulate.Create(tableData)
	t.SetHeaders(tableHeaders)
	t.SetDenseMode()

	fmt.Println(t.Render("simple"))
}

func printCSVLine(fields []string) {
	qfields := []string{}
	for _, f := range fields {
		qfields = append(qfields, fmt.Sprintf("\"%s\"", f))
	}
	fmt.Println(strings.Join(qfields, ","))
}

func printResultCSV(result interface{}, template string) {
	tableData, tableHeaders := prepareResultTableCustomColumns(result, template)

	printCSVLine(tableHeaders)
	for _, row := range tableData {
		printCSVLine(row)
	}
}

func printResultJSONPath(result interface{}, template string) {
	jp, err := jsonpath.Get(relaxedJSONPathExpression(template), result)
	if err != nil {
		log.Fatalf("JSONPath parsing error: %v", err)
	}

	switch res := jp.(type) {
	case []interface{}:
		for _, v := range res {
			fmt.Printf("%v\n", v)
		}
	default:
		fmt.Printf("%v\n", res)
	}
}

type loggingTransport struct{}

func newLoggingTransport() http.RoundTripper {
	return &loggingTransport{}
}

func (t *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	outreq, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}
	log.Debugf("Sending API request:\n%s\n", string(outreq))

	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return res, err
	}

	outres, err := httputil.DumpResponse(res, true)
	if err != nil {
		return nil, err
	}

	log.Debugf("Received API response:\n%s\n", string(outres))
	return res, nil
}
