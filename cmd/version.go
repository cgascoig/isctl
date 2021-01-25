package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	commit  = "unknown"
	version = "unknown"
)

const ()

func newCmdVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Run:   runCmdVersion,
		Short: "Print the version information",
	}
}

func runCmdVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("isctl version %s (Git commit %s)\n", version, commit)
}

func init() {
	auxCommands = append(auxCommands, newCmdVersion())
}
