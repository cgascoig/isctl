package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVariables(t *testing.T) {
	// Set up environment variables
	os.Setenv("ISCTL_VAR_ENV_VAR", "env-value")
	os.Setenv("ISCTL_VAR_OVERRIDE_ME", "env-value")
	defer os.Unsetenv("ISCTL_VAR_ENV_VAR")
	defer os.Unsetenv("ISCTL_VAR_OVERRIDE_ME")

	// Create a temporary variable file
	varFileContent := []byte("FILE_VAR: file-value\nOVERRIDE_ME: file-value\nALSO_OVERRIDE_ME: file-value")
	tmpFile, err := os.CreateTemp("", "vars.yaml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.Write(varFileContent)
	assert.NoError(t, err)
	tmpFile.Close()

	config := &applyConfig{
		varFlags: []string{"FLAG_VAR=flag-value", "ALSO_OVERRIDE_ME=flag-value"},
		varFile:  tmpFile.Name(),
	}

	vars, err := config.getVariables()
	assert.NoError(t, err)

	expected := map[string]interface{}{
		"ENV_VAR":          "env-value",
		"FILE_VAR":         "file-value",
		"FLAG_VAR":         "flag-value",
		"OVERRIDE_ME":      "file-value", // File overrides Env
		"ALSO_OVERRIDE_ME": "flag-value", // Flag overrides File
	}

	for k, v := range expected {
		assert.Equal(t, v, vars[k], "Variable %s mismatch", k)
	}
}

func TestProcessTemplate(t *testing.T) {
	vars := map[string]interface{}{
		"Name": "world",
	}

	tests := []struct {
		name    string
		content string
		want    string
		wantErr bool
	}{
		{
			name:    "Simple variable",
			content: "Hello {{ .Vars.Name }}",
			want:    "Hello world",
			wantErr: false,
		},
		{
			name:    "Sprig function (upper)",
			content: "Hello {{ .Vars.Name | upper }}",
			want:    "Hello WORLD",
			wantErr: false,
		},
		{
			name:    "Sprig function (default)",
			content: "Hello {{ .Vars.Missing | default \"nobody\" }}",
			want:    "Hello nobody",
			wantErr: false,
		},
		{
			name:    "Invalid template",
			content: "Hello {{ .Vars.Name }",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := processTemplate([]byte(tt.content), vars)
			if (err != nil) != tt.wantErr {
				t.Errorf("processTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.Equal(t, tt.want, string(got))
			}
		})
	}
}
