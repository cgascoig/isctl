package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cgascoig/isctl/pkg/util"
)

var (
	commit  = "unknown"
	version = "unknown"
)

const ()

func newCmdVersion(client *util.IsctlClient) *cobra.Command {
	log.Trace("Running version cmd generator")
	return &cobra.Command{
		Use:               "version",
		Run:               runCmdVersion,
		Short:             "Print the version information",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	}
}

func runCmdVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("isctl version %s (Git commit %s)\n", version, commit)
}

func init() {
	auxCommandsGenerators = append(auxCommandsGenerators, newCmdVersion)
}
