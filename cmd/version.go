package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	gitCommit     = "unknown"
	versionString = "unknown"
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
	fmt.Printf("isctl version %s (Git commit %s)\n", versionString, gitCommit)
}
