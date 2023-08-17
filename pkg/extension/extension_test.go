package extension

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndRunExtension(t *testing.T) {
	e := NewExtension()

	var err error

	var code []byte = []byte(
		`
def run():
	return [1,2]
`)

	err = e.SetCode(code)
	assert.Nil(t, err)
}

func TestInvalidCode(t *testing.T) {
	e := NewExtension()

	var err error

	var code []byte = []byte(
		`
	run1 = function() {
		return [1,2];
	}
`)

	err = e.SetCode(code)
	assert.NotNil(t, err)
}

func TestGetCommand(t *testing.T) {
	py := `
def cmd():
	return {
		"use": "iks",
		"short": "Manage IKS",
		"children": [
			{
				"use": "cluster",
				"children": [
					{
						"use": "create",
						"run": "cmd_iks_cluster_create",
					}
				]
			}
		]
	}
`
	e := NewExtension()

	var err error

	var code []byte = []byte(py)

	err = e.SetCode(code)
	assert.Nil(t, err)

	cmd, err := e.GetCommand()
	assert.Nil(t, err)

	assert.Equal(t, cmd.Use, "iks")
	assert.Equal(t, cmd.Short, "Manage IKS")

	// TODO: Fix test to check children

	// expectedCmd := &cobra.Command{
	// 	Use:   "iks",
	// 	Short: "Manage IKS",
	// }
	// clusterCmd := &cobra.Command{
	// 	Use: "cluster",
	// }
	// createCmd := &cobra.Command{
	// 	Use: "create",
	// 	Run: getRunFunction(e, "cmd_iks_cluster_create"),
	// }

	// clusterCmd.AddCommand(createCmd)
	// expectedCmd.AddCommand(clusterCmd)

	// assert.Equal(t, expectedCmd, cmd)
}
