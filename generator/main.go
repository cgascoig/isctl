package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

var (
	operationsFilename string
	templateFilename   string
	outputFilename     string
)

// OperationsFile represents the YAML operations file
type OperationsFile struct {
	Operations []Operation      `yaml:"operations"`
	Models     map[string]Model `yaml:"models"`
}

type MoRef struct {
	DataType string
	Path     []string
}

func (mor *MoRef) IsList() bool {
	r := regexp.MustCompile(`^\[\]`)
	return r.MatchString(mor.DataType)
}

// Operation represents the YAML of one operation
type Operation struct {
	Path           string  `yaml:"path"`
	OperationID    string  `yaml:"operationId"`
	HTTPMethod     string  `yaml:"httpMethod"`
	ReturnType     string  `yaml:"returnType"`
	ReturnBaseType string  `yaml:"returnBaseType"`
	Summary        string  `yaml:"summary"`
	BaseName       string  `yaml:"baseName"`
	Params         []Param `yaml:"params"`

	MoRefs []MoRef
}

func (o *Operation) IsListOperation() bool {
	r := regexp.MustCompile(`List$`)
	return r.MatchString(o.OperationID)
}

func (o *Operation) IsUpdateOperation() bool {
	r := regexp.MustCompile(`^Update`)
	return r.MatchString(o.OperationID)
}

func (o *Operation) IsCreateOperation() bool {
	r := regexp.MustCompile(`^Create`)
	return r.MatchString(o.OperationID)
}

func (o *Operation) IsDeleteOperation() bool {
	r := regexp.MustCompile(`^Delete`)
	return r.MatchString(o.OperationID)
}

func (o *Operation) DeleteOperationDataType() string {
	re := fmt.Sprintf(`^Delete%s(\w+)$`, o.BaseName)
	r := regexp.MustCompile(re)
	m := r.FindStringSubmatch(o.OperationID)
	if len(m) != 2 {
		return ""
	}
	return fmt.Sprintf("%s.%s", strings.ToLower(o.BaseName), m[1])
}

func (o *Operation) ReturnClassID() string {
	regExpClassID := regexp.MustCompile(`\.`)
	if regExpClassID.MatchString(o.ReturnBaseType) {
		// ReturnBaseType is not empty and contains a "." - so it must be a classID
		return o.ReturnBaseType
	}

	if regexp.MustCompile(`^Delete`).MatchString(o.OperationID) {
		return o.DeleteOperationDataType()
	}

	re := fmt.Sprintf(`^%s(\w+)$`, o.BaseName)
	r := regexp.MustCompile(re)
	m := r.FindStringSubmatch(o.ReturnBaseType)
	if len(m) == 2 {
		r = regexp.MustCompile(`^(\w+)Response$`)
		m2 := r.FindStringSubmatch(m[1])
		if len(m2) == 2 {
			return fmt.Sprintf("%s.%s.Response", strings.ToLower(o.BaseName), m2[1])
		} else {
			return fmt.Sprintf("%s.%s", strings.ToLower(o.BaseName), m[1])
		}
	} else {
		return ""
	}

}

func (o *Operation) OperationClassID() string {
	if o.ReturnBaseType != "" {
		r := regexp.MustCompile(`^(\w+\.\w+)\.Response$`)
		m := r.FindStringSubmatch(o.ReturnBaseType)
		if m == nil || len(m) != 2 {
			return o.ReturnBaseType
		}
		return m[1]
	}
	re := fmt.Sprintf(`^\w+%s(\w+)$`, o.BaseName)
	r := regexp.MustCompile(re)
	m := r.FindStringSubmatch(o.OperationID)
	if len(m) != 2 {
		return ""
	}
	return fmt.Sprintf("%s.%s", strings.ToLower(o.BaseName), m[1])
}

// Param represents the YAML of one parameter
type Param struct {
	ParamName    string `yaml:"paramName"`
	Description  string `yaml:"description"`
	DataType     string `yaml:"dataType"`
	Required     bool   `yaml:"required"`
	IsBodyParam  bool   `yaml:"isBodyParam"`
	IsPathParam  bool   `yaml:"isPathParam"`
	IsQueryParam bool   `yaml:"isQueryParam"`
}

// Model represents the YAML of one model from the OpenAPI spec
type Model struct {
	Parents []string `yaml:"parents"`
	Vars    []Var    `yaml:"vars"`
}

// Var represents the YAML of one member var of the OpenAPI model
type Var struct {
	Name     string `yaml:"name"`
	DataType string `yaml:"dataType"`
	Required bool   `yaml:"required"`
	ReadOnly bool   `yaml:"readOnly"`
	Nullable bool   `yaml:"nullable"`
	Default  string `yaml:"default"`
	Ignore   bool   // Used in cli.gotmpl to skip vars
}

func (v *Var) IsList() bool {
	validTypeRegExp := regexp.MustCompile(`^\[\][a-zA-Z0-9]+$`)
	return validTypeRegExp.MatchString(v.DataType) && v.DataType != "[]int64" && v.DataType != "[]float32" && v.DataType != "[]int32"
}

func (v *Var) ListElementType() string {
	validTypeRegExp := regexp.MustCompile(`^\[\]([a-zA-Z0-9]+)$`)
	m := validTypeRegExp.FindStringSubmatch(v.DataType)
	return m[1]
}

func (v *Var) ValidGenericType() bool {
	validTypeRegExp := regexp.MustCompile(`^(\[\])?([a-zA-Z0-9]+)?([a-zA-Z0-9]+)$`)
	return validTypeRegExp.MatchString(v.DataType)
}

func (v *Var) NormalisedType() string {
	var t string
	isList := false
	if v.IsList() {
		isList = true
		t = v.ListElementType()
	} else {
		t = v.DataType
	}

	// if type is already in pkg.type, do not change, otherwise change type to openapi.type
	qualifiedTypeRegExp := regexp.MustCompile(`^([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)$`)
	if !qualifiedTypeRegExp.MatchString(t) {
		t = fmt.Sprintf("openapi.%s", t)
	}

	if isList {
		return fmt.Sprintf("[]%s", t)
	} else {
		return t
	}
}

type TemplateData struct {
	CliTree    *CliItem
	Operations []Operation
	Models     map[string]Model
}

func (t TemplateData) GetTypeToIntersightClassIDMapping() map[string]string {
	ret := map[string]string{}

	for modelName, model := range t.Models {
		ret[modelName] = getModelClassID(t.Models, model)
	}

	return ret
}

func getModelClassID(models map[string]Model, model Model) string {
	for _, v := range model.Vars {
		if v.Name == "ClassId" {
			return v.Default
		}
	}

	for _, p := range model.Parents {
		m := models[p]
		classID := getModelClassID(models, m)
		if classID != "" && classID != "<nil>" {
			return classID
		}
	}

	return ""
}

func generate(templateFilename string, outputFilename string, data interface{}) {
	log.Printf("Generating '%s' using template file '%s'", outputFilename, templateFilename)

	tmpl, err := template.ParseFiles(templateFilename)
	if err != nil {
		log.Fatalf("Error loading template file: %v", err)
	}

	outfile, err := os.Create(outputFilename)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}

	err = tmpl.Execute(outfile, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func main() {
	flag.Parse()

	log.Printf("Using operations data from '%s'", operationsFilename)

	opData, err := ReadOpenAPIFile(operationsFilename)
	if err != nil {
		log.Fatalf("Error reading OpenAPI spec: %v", err)
	}

	// Sort some stuff so the generated code doesn't change when generated multiple times
	for _, m := range opData.Models {
		sort.Slice(m.Vars, func(i, j int) bool { return m.Vars[i].Name < m.Vars[j].Name })
	}
	sort.Slice(opData.Operations, func(i, j int) bool { return opData.Operations[i].OperationID < opData.Operations[j].OperationID })

	cliTree := generateCliTree(opData)

	data := TemplateData{
		CliTree:    cliTree,
		Operations: opData.Operations,
		Models:     opData.Models,
	}

	generate("generator/cli.go.tmpl", "pkg/gen/cli.go", data)
	generate("generator/types.go.tmpl", "pkg/gen/types.go", data)
	generate("generator/operations.go.tmpl", "pkg/gen/operations.go", data)

	log.Printf("Finished post processing.")
}

func init() {
	flag.StringVar(&operationsFilename, "operations", "spec/openapi.yaml", "OpenAPI v3 spec")
	flag.StringVar(&templateFilename, "template", "commandline.gotmpl", "Template file to generate CLI code")
	flag.StringVar(&outputFilename, "output", "commandline.go", "File to output generated go code")
}
