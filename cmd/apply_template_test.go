package main

import (
	"os"
	"path/filepath"
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

	vars, err := config.getVariables([]string{})
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

func TestGetVariablesWithDirectory(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "isctl-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create isctl.vars.yaml in the directory
	varFileContent := []byte("DIR_VAR: dir-value\nOVERRIDE_ME: dir-value\nALSO_OVERRIDE_ME: dir-value")
	err = os.WriteFile(filepath.Join(tmpDir, "isctl.vars.yaml"), varFileContent, 0644)
	assert.NoError(t, err)

	// Set up environment variables
	os.Setenv("ISCTL_VAR_ENV_VAR", "env-value")
	os.Setenv("ISCTL_VAR_OVERRIDE_ME", "env-value")
	defer os.Unsetenv("ISCTL_VAR_ENV_VAR")
	defer os.Unsetenv("ISCTL_VAR_OVERRIDE_ME")

	config := &applyConfig{
		varFlags: []string{"FLAG_VAR=flag-value", "ALSO_OVERRIDE_ME=flag-value"},
	}

	// Pass the directory as one of the files to apply
	vars, err := config.getVariables([]string{tmpDir})
	assert.NoError(t, err)

	expected := map[string]interface{}{
		"ENV_VAR":          "env-value",
		"DIR_VAR":          "dir-value",
		"FLAG_VAR":         "flag-value",
		"OVERRIDE_ME":      "env-value",  // Env overrides Directory
		"ALSO_OVERRIDE_ME": "flag-value", // Flag overrides Directory
	}

	for k, v := range expected {
		assert.Equal(t, v, vars[k], "Variable %s mismatch", k)
	}
}

func TestGetVariablesWithDirectoryCollision(t *testing.T) {
	// Create two temporary directories
	tmpDir1, err := os.MkdirTemp("", "isctl-test-1")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir1)

	tmpDir2, err := os.MkdirTemp("", "isctl-test-2")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir2)

	// Create isctl.vars.yaml in both directories defining the same variable
	varFileContent1 := []byte("COLLISION_VAR: value1")
	err = os.WriteFile(filepath.Join(tmpDir1, "isctl.vars.yaml"), varFileContent1, 0644)
	assert.NoError(t, err)

	varFileContent2 := []byte("COLLISION_VAR: value2")
	err = os.WriteFile(filepath.Join(tmpDir2, "isctl.vars.yaml"), varFileContent2, 0644)
	assert.NoError(t, err)

	config := &applyConfig{}

	// Pass both directories
	_, err = config.getVariables([]string{tmpDir1, tmpDir2})

	// Expect an error due to collision
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "variable 'COLLISION_VAR' is defined in multiple directory variable files")
}

func TestGetVariablesWithDirectoryNotRegularFile(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "isctl-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create a directory named isctl.vars.yaml (not a regular file)
	err = os.Mkdir(filepath.Join(tmpDir, "isctl.vars.yaml"), 0755)
	assert.NoError(t, err)

	config := &applyConfig{}

	// Pass the directory
	// Should not error, but should not load any variables
	vars, err := config.getVariables([]string{tmpDir})
	assert.NoError(t, err)
	assert.Empty(t, vars)
}
