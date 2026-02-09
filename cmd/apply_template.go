package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"gopkg.in/yaml.v3"
)

// getVariables loads variables from multiple sources with the following precedence (highest to lowest):
// 1. Command line flags (--var)
// 2. Variable file (--var-file)
// 3. Environment variables (ISCTL_VAR_*)
// 4. Directory variables (isctl.vars.yaml/yml in directories passed to -f)
func (config *applyConfig) getVariables(applyFilenames []string) (map[string]interface{}, error) {
	vars := make(map[string]interface{})

	// Directory variables (lowest precedence, so loaded first)

	// We need to track where variables came from to detect collisions
	dirVarSources := make(map[string]string)

	for _, path := range applyFilenames {
		info, err := os.Stat(path)
		if err != nil {
			continue // Skip invalid paths, they will be handled later
		}
		if info.IsDir() {
			// Check for isctl.vars.yaml or isctl.vars.yml
			for _, ext := range []string{".yaml", ".yml"} {
				varFileName := "isctl.vars" + ext
				varFile := filepath.Join(path, varFileName)

				fileInfo, err := os.Stat(varFile)
				// Check if file exists and is a regular file
				if err == nil && fileInfo.Mode().IsRegular() {
					fileVars, err := loadVarFile(varFile)
					if err != nil {
						return nil, fmt.Errorf("failed to load directory variable file %s: %w", varFile, err)
					}

					for k, v := range fileVars {
						// Check for collisions with other directory variable files
						if source, exists := dirVarSources[k]; exists {
							return nil, fmt.Errorf("variable '%s' is defined in multiple directory variable files: %s and %s", k, source, varFile)
						}

						vars[k] = v
						dirVarSources[k] = varFile
					}
				}
			}
		}
	}

	// Environment variables starting with ISCTL_VAR_
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 && strings.HasPrefix(pair[0], "ISCTL_VAR_") {
			key := strings.TrimPrefix(pair[0], "ISCTL_VAR_")
			if key != "" {
				vars[key] = pair[1]
			}
		}
	}

	// Variable file
	if config.varFile != "" {
		fileVars, err := loadVarFile(config.varFile)
		if err != nil {
			return nil, fmt.Errorf("failed to load variable file: %w", err)
		}
		for k, v := range fileVars {
			vars[k] = v
		}
	}

	// Command line flags
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
