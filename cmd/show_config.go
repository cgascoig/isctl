package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cgascoig/isctl/pkg/util"
)

func newCmdShowConfig(client *util.IsctlClient) *cobra.Command {
	log.Trace("Running show-configuration cmd generator")
	cmd := &cobra.Command{
		Use:    "show-configuration",
		Run:    func(cmd *cobra.Command, args []string) { runCmdShowConfig(client, cmd, args) },
		Short:  "Show the resolved configuration",
		Hidden: true,
	}

	return cmd
}

func runCmdShowConfig(client *util.IsctlClient, cmd *cobra.Command, args []string) {
	fmt.Printf("%s: %#v\n", CKIntersightApiKeyId, client.IntersightConfig.KeyID)
	fmt.Printf("%s: %#v\n", CKIntersightSecretKey, client.IntersightConfig.KeyData)
	fmt.Printf("%s: %#v\n", CKIntersightFqdn, client.IntersightConfig.Host)
	fmt.Printf("%s: %#v\n", CKIntersightInsecure, gK.Bool(CKIntersightInsecure))
}

func init() {
	auxCommandsGenerators = append(auxCommandsGenerators, newCmdShowConfig)
}
