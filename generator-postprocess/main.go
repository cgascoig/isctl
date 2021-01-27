package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"text/template"
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v2"
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

// Operation represents the YAML of one operation
type Operation struct {
	OperationID    string  `yaml:"operationId"`
	HTTPMethod     string  `yaml:"httpMethod"`
	ReturnType     string  `yaml:"returnType"`
	ReturnBaseType string  `yaml:"returnBaseType"`
	Summary        string  `yaml:"summary"`
	BaseName       string  `yaml:"baseName"`
	Params         []Param `yaml:"params"`
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
	return fmt.Sprintf("%s.%s", strings.ToLower(o.BaseName),m[1])
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

	opFileBytes, err := ioutil.ReadFile(operationsFilename)
	if err != nil {
		log.Fatalf("Error reading operations data: %v", err)
	}

	var opData OperationsFile
	err = yaml.Unmarshal(opFileBytes, &opData)
	if err != nil {
		log.Fatalf("Error unmarshalling data from operations file: %v", err)
	}

	cliTree := generateCliTree(&opData)

	data := struct {
		CliTree    *CliItem
		Operations []Operation
		Models     map[string]Model
	}{
		CliTree:    cliTree,
		Operations: opData.Operations,
		Models:     opData.Models,
	}

	generate("generator-postprocess/cli.go.tmpl", "cmd/cli.go", data)
	generate("generator-postprocess/types.go.tmpl", "cmd/types.go", data)
	generate("generator-postprocess/operations.go.tmpl", "cmd/operations.go", data)

	log.Printf("Finished post processing.")
}

func init() {
	flag.StringVar(&operationsFilename, "operations", "openapi/operations.yaml", "YAML file describing available operations")
	flag.StringVar(&templateFilename, "template", "commandline.gotmpl", "Template file to generate CLI code")
	flag.StringVar(&outputFilename, "output", "commandline.go", "File to output generated go code")
}
