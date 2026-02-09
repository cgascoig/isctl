package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/icza/dyno"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v3"

	"github.com/cgascoig/isctl/pkg/gen"
	"github.com/cgascoig/isctl/pkg/oapi"
	"github.com/cgascoig/isctl/pkg/util"
)

type editConfig struct {
	client *util.IsctlClient
}

func newCmdEdit(client *util.IsctlClient) *cobra.Command {
	log.Trace("Running edit cmd generator")
	config := editConfig{
		client: client,
	}

	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit an Intersight resource in your default editor",
		Long: `Edit an Intersight resource using your preferred text editor.

The resource will be fetched, filtered to show only editable properties,
and opened in your editor. After saving and closing the editor, any changes
will be applied via an update operation.

The editor is determined by:
1. $EDITOR environment variable
2. $VISUAL environment variable  
3. Platform default (vi on Unix, notepad on Windows)

Examples:
  # Edit an NTP policy by name
  isctl edit ntp policy name my-ntp-policy

  # Edit an NTP policy by Moid
  isctl edit ntp policy moid 1234567890abcdef12345678`,
	}

	// Generate subcommands from the CLI tree for resources that have update operations
	cliTree := oapi.GenerateCliTree()
	config.addEditSubcommands(cmd, cliTree, []string{})

	return cmd
}

func init() {
	auxCommandsGenerators = append(auxCommandsGenerators, newCmdEdit)
}

// addEditSubcommands finds the "update" branch of the CLI tree and mirrors its structure
// under the edit command.
func (config *editConfig) addEditSubcommands(parent *cobra.Command, cliTree *oapi.CliItem, path []string) {
	if cliTree.Children == nil {
		return
	}

	// Find the "update" command at the top level - this is the tree we want to mirror
	updateItem, hasUpdate := cliTree.Children["update"]
	if !hasUpdate || updateItem.Children == nil {
		log.Trace("No 'update' command found in CLI tree")
		return
	}

	// Recursively build edit subcommands that mirror the update tree structure
	config.mirrorUpdateTree(parent, updateItem, path)
}

// mirrorUpdateTree recursively creates edit subcommands that mirror the update command tree
func (config *editConfig) mirrorUpdateTree(parent *cobra.Command, item *oapi.CliItem, path []string) {
	if item.Children == nil {
		return
	}

	for token, child := range item.Children {
		// If this has an Operation with a moid/name parameter, create the leaf command
		if child.Parameter == "moid" || child.Parameter == "name" {
			paramCmd := config.createEditParamCommand(child, path)
			if paramCmd != nil {
				parent.AddCommand(paramCmd)
			}
		} else {
			// This is a grouping node, create a subcommand and recurse
			subCmd := &cobra.Command{
				Use:   token,
				Short: child.Help,
			}
			config.mirrorUpdateTree(subCmd, child, append(path, token))
			// Only add if it has subcommands
			if len(subCmd.Commands()) > 0 {
				parent.AddCommand(subCmd)
			}
		}
	}
}

// createEditParamCommand creates a command for editing by moid or name
func (config *editConfig) createEditParamCommand(cliItem *oapi.CliItem, path []string) *cobra.Command {
	if cliItem.Operation == nil {
		return nil
	}

	// ReturnClassID() returns schema reference like "#/components/schemas/ntp.Policy"
	// Convert to class ID like "ntp.Policy" for meta lookups
	classID := oapi.SchemaNameToClassId(cliItem.Operation.ReturnClassID())
	paramType := cliItem.Parameter // "moid" or "name"

	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%s <%s>", paramType, paramType),
		Short: fmt.Sprintf("Edit by %s", paramType),
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := config.runEdit(classID, paramType, args[0])
			if err != nil {
				log.Fatalf("Error editing resource: %v", err)
			}
		},
	}

	return cmd
}

// runEdit executes the edit workflow
func (config *editConfig) runEdit(classID, paramType, paramValue string) error {
	// Step 1: Fetch the object
	log.Debugf("Fetching %s with %s=%s", classID, paramType, paramValue)

	getOperation := gen.GetGetOperationForClassID(classID)
	if getOperation == nil {
		return fmt.Errorf("no get operation found for class %s", classID)
	}

	var queryParams map[string]string
	var args []string

	if paramType == "name" {
		queryParams = map[string]string{"filter": fmt.Sprintf("Name eq '%s'", paramValue)}
	} else {
		args = []string{paramValue}
	}

	res, err := getOperation.Execute(config.client, args, queryParams)
	if err != nil {
		return fmt.Errorf("error fetching object: %w", err)
	}

	// Get the Moid for later update
	moid, ok := util.GetMoid(res)
	if !ok {
		return fmt.Errorf("object not found or multiple objects match")
	}

	// Convert result to map
	mo, err := resultToMap(res)
	if err != nil {
		return fmt.Errorf("error processing result: %w", err)
	}

	// Step 2: Filter to writable properties and convert to YAML
	originalYAML, err := oapi.FormatEditableYAML(mo, classID)
	if err != nil {
		return fmt.Errorf("error formatting YAML: %w", err)
	}

	// Step 3: Write to temp file
	tmpFile, err := os.CreateTemp("", "isctl-edit-*.yaml")
	if err != nil {
		return fmt.Errorf("error creating temp file: %w", err)
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)

	if _, err := tmpFile.Write(originalYAML); err != nil {
		tmpFile.Close()
		return fmt.Errorf("error writing temp file: %w", err)
	}
	tmpFile.Close()

	// Step 4: Open editor
	editor := getEditor()
	log.Debugf("Opening editor: %s %s", editor, tmpPath)

	editorCmd := exec.Command(editor, tmpPath)
	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr

	if err := editorCmd.Run(); err != nil {
		return fmt.Errorf("error running editor: %w", err)
	}

	// Step 5: Read modified file
	modifiedYAML, err := os.ReadFile(tmpPath)
	if err != nil {
		return fmt.Errorf("error reading modified file: %w", err)
	}

	// Step 6: Check if content changed
	if bytes.Equal(originalYAML, modifiedYAML) {
		fmt.Println("Edit cancelled, no changes made.")
		return nil
	}

	// Step 7: Parse modified YAML
	var modifiedMO map[string]any
	if err := yaml.Unmarshal(modifiedYAML, &modifiedMO); err != nil {
		return fmt.Errorf("error parsing modified YAML: %w", err)
	}

	// Step 8: Perform update
	log.Debugf("Updating %s with Moid %s", classID, moid)

	updateOperation := gen.GetUpdateOperationForClassID(classID)
	if updateOperation == nil {
		return fmt.Errorf("no update operation found for class %s", classID)
	}

	if err := updateOperation.SetBodyParams(config.client, modifiedMO); err != nil {
		return fmt.Errorf("error setting update body: %w", err)
	}

	_, err = updateOperation.Execute(config.client, []string{moid}, nil)
	if err != nil {
		return fmt.Errorf("error executing update: %w", err)
	}

	fmt.Println("Edit applied successfully.")
	return nil
}

// getEditor returns the editor command to use
func getEditor() string {
	if editor := os.Getenv("EDITOR"); editor != "" {
		return editor
	}
	if editor := os.Getenv("VISUAL"); editor != "" {
		return editor
	}
	if runtime.GOOS == "windows" {
		return "notepad"
	}
	return "vi"
}

// resultToMap converts an API result to a map[string]any
func resultToMap(result any) (map[string]any, error) {
	// Handle list results (from filter queries)
	if results, err := dyno.GetSlice(result, "Results"); err == nil {
		if len(results) == 0 {
			return nil, fmt.Errorf("no results found")
		}
		if len(results) > 1 {
			return nil, fmt.Errorf("multiple results found, expected exactly one")
		}
		if mo, ok := results[0].(map[string]any); ok {
			return mo, nil
		}
		return nil, fmt.Errorf("unexpected result type")
	}

	// Handle single object result
	if mo, ok := result.(map[string]any); ok {
		return mo, nil
	}

	return nil, fmt.Errorf("unexpected result type: %T", result)
}
