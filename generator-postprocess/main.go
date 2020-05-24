package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	yaml "gopkg.in/yaml.v2"
)

var (
	operationsFilename string
	templateFilename   string
	outputFilename     string
)

type OperationsFile struct {
	Operations []Operation `yaml:"operations"`
}

type Operation struct {
	OperationId string  `yaml:"operationId"`
	HttpMethod  string  `yaml:"httpMethod"`
	ReturnType  string  `yaml:"returnType"`
	Summary     string  `yaml:"summary"`
	BaseName    string  `yaml:"baseName"`
	Params      []Param `yaml:"params"`
}

type Param struct {
	ParamName    string `yaml:"paramName"`
	Description  string `yaml:"description"`
	DataType     string `yaml:"dataType"`
	Required     bool   `yaml:"required"`
	IsBodyParam  bool   `yaml:"isBodyParam"`
	IsPathParam  bool   `yaml:"isPathParam"`
	IsQueryParam bool   `yaml:"isQueryParam"`
}

func main() {
	flag.Parse()

	log.Printf("Generating '%s' using template file '%s' and operations data from '%s'", outputFilename, templateFilename, operationsFilename)

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

	tmpl, err := template.ParseFiles(templateFilename)
	if err != nil {
		log.Fatalf("Error loading template file: %v", err)
	}

	outfile, err := os.Create(outputFilename)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}

	err = tmpl.Execute(outfile, cliTree)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	log.Printf("Finished post processing.")
}

func init() {
	flag.StringVar(&operationsFilename, "operations", "openapi/operations.yaml", "YAML file describing available operations")
	flag.StringVar(&templateFilename, "template", "commandline.gotmpl", "Template file to generate CLI code")
	flag.StringVar(&outputFilename, "output", "commandline.go", "File to output generated go code")
}
