package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"gopkg.in/yaml.v3"
)

// getVariables loads variables from multiple sources with the following precedence (highest to lowest):
// 1. Command line flags (--var)
// 2. Variable file (--var-file)
// 3. Environment variables (ISCTL_VAR_*)
func (config *applyConfig) getVariables() (map[string]interface{}, error) {
	vars := make(map[string]interface{})

	// 1. Environment variables starting with ISCTL_VAR_
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 && strings.HasPrefix(pair[0], "ISCTL_VAR_") {
			key := strings.TrimPrefix(pair[0], "ISCTL_VAR_")
			if key != "" {
				vars[key] = pair[1]
			}
		}
	}

	// 2. Variable file
	if config.varFile != "" {
		fileVars, err := loadVarFile(config.varFile)
		if err != nil {
			return nil, fmt.Errorf("failed to load variable file: %w", err)
		}
		for k, v := range fileVars {
			vars[k] = v
		}
	}

	// 3. Command line flags
	for _, v := range config.varFlags {
		pair := strings.SplitN(v, "=", 2)
		if len(pair) != 2 {
			return nil, fmt.Errorf("invalid variable format: %s (expected key=value)", v)
		}
		vars[pair[0]] = pair[1]
	}

	return vars, nil
}

func loadVarFile(filename string) (map[string]interface{}, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var vars map[string]interface{}
	err = yaml.Unmarshal(content, &vars)
	if err != nil {
		return nil, err
	}
	return vars, nil
}

// processTemplate executes the content as a Go template, providing the variables under the "Vars" key.
// It includes Sprig functions.
func processTemplate(content []byte, vars map[string]interface{}) ([]byte, error) {
	tmpl, err := template.New("manifest").Funcs(sprig.FuncMap()).Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	data := map[string]interface{}{
		"Vars": vars,
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.Bytes(), nil
}
