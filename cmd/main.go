package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/knadh/koanf/parsers/yaml"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cgascoig/intersight-simple-go/intersight"
	"github.com/cgascoig/isctl/pkg/gen"
	"github.com/cgascoig/isctl/pkg/util"
)

var (
	client         = &util.IsctlClient{}
	jsonPathFilter string

	auxCommandsGenerators []commandGenerator
)

type commandGenerator func(*util.IsctlClient) *cobra.Command

const (
	traceEnvName = "ISCTL_TRACE"
)

func main() {
	if os.Getenv(traceEnvName) != "" {
		log.SetLevel(log.TraceLevel)
	}

	log.Trace("isctl starting")

	rootCmd := gen.GetCommands(client, resultHandler)
	rootCmd.Use = "isctl"

	log.Trace("Got generated commands")

	rootCmd.PersistentFlags().String("config", "", "config file (default is $HOME/.isctl.yaml)")

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose logging")

	// New config keys
	rootCmd.PersistentFlags().String(FlagIntersightApiKeyId, "", "Intersight API Key ID")
	rootCmd.PersistentFlags().String(FlagIntersightSecretKey, "", "Intersight Secret Key (filename)")
	rootCmd.PersistentFlags().String(FlagIntersightFqdn, "", "Intersight API FQDN (default intersight.com)")

	rootCmd.PersistentFlags().String(CKKeyID, "", "API Key ID [deprecated]")
	rootCmd.PersistentFlags().String(CKKeyFile, "", "API Private Key Filename [deprecated]")

	rootCmd.PersistentFlags().String(CKServer, "intersight.com", "Intersight API Server Address (e.g.\"intersight.com\") [deprecated]")
	rootCmd.PersistentFlags().BoolP(CKInsecure, "k", false, "Allow insecure server connections (disable SSL certificate validation)")

	rootCmd.PersistentFlags().StringP(CKOutputFormat, "o", "default", `Output format. One of default|yaml|json|table|jsonpath|custom-columns|csv. Examples:
	Get Name attribute from all NTP policies: isctl get ntp policy -o jsonpath="[*].Name"
	Table with just Name and Enabled attributes: isctl get ntp policy -o custom-columns=NAME:.Name,ENABLED:.Enabled
	See [https://isctl.netlify.app/1-basic-queries/#output-customisation]`)
	rootCmd.PersistentFlags().StringVar(&jsonPathFilter, "jsonpath", "", "JSONPath filter to apply to the result (e.g. \"$.Name\")")

	cobra.OnInitialize(func() {
		initConfig(rootCmd.PersistentFlags())
	})

	configCmd := &cobra.Command{
		Use:               "configure",
		Run:               configure,
		Short:             "Configure the isctl command",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	}
	rootCmd.AddCommand(configCmd)

	log.Trace("Running auxCommandGenerators")
	for _, cmdGen := range auxCommandsGenerators {
		rootCmd.AddCommand(cmdGen(client))
	}
	log.Trace("Finished auxCommandGenerators")

	log.Trace("Loading extensions")
	for _, cmd := range loadExtensions(client) {
		log.Tracef("Loaded extension command %s", cmd.Use)
		rootCmd.AddCommand(cmd)
	}
	log.Trace("Finished loading extensions")

	rootCmd.PersistentPreRunE = validateFlags

	log.Trace("CLI building complete")
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func configure(cmd *cobra.Command, args []string) {
	log.Trace("Starting configure")
	scanner := bufio.NewScanner(os.Stdin)

	// configure keyID
	fmt.Printf("%s is currently '%s'\n", CKIntersightApiKeyId, gK.String(CKIntersightApiKeyId))
	fmt.Printf("Enter new %s or press Enter to keep existing: ", CKIntersightApiKeyId)
	scanner.Scan()
	if input := scanner.Text(); input != "" {
		gK.Set(CKIntersightApiKeyId, input)
	}

	// configure key file name
	fmt.Printf("%s is currently '%s'\n", CKIntersightSecretKey, gK.String(CKIntersightSecretKey))
	fmt.Printf("Enter new key file name or press Enter to keep existing: ")
	scanner.Scan()
	if input := scanner.Text(); input != "" {
		gK.Set(CKIntersightSecretKey, input)
	}

	// configure server
	fmt.Printf("%s is currently '%s'\n", CKIntersightFqdn, gK.String(CKIntersightFqdn))
	fmt.Printf("Enter new Intersight FQDN or press Enter to keep existing: ")
	scanner.Scan()
	if input := scanner.Text(); input != "" {
		gK.Set(CKIntersightFqdn, input)
	}

	log.Println("Writing config file")
	b, err := gK.Marshal(yaml.Parser())
	if err != nil {
		log.Fatalf("Error occurred writing config file: %v", err)
	}
	err = os.WriteFile(getConfigFilePath(cmd.PersistentFlags()), b, 0666)
	if err != nil {
		log.Fatalf("Error occurred writing config file: %v", err)
	}
	log.Trace("Finished configure")
}

func validateFlags(cmd *cobra.Command, args []string) error {
	log.Trace("Starting validateFlags")

	var httpTransport http.RoundTripper = http.DefaultTransport

	// Setup logging
	if gK.Bool("verbose") && os.Getenv(traceEnvName) == "" {
		log.SetLevel(log.DebugLevel)
		log.Debug("Logging level set to debug(verbose)")
	}

	if logLevel := log.GetLevel(); logLevel == log.DebugLevel || logLevel == log.TraceLevel {
		httpTransport = newLoggingTransport()
	}

	var err error

	keyID := gK.String(CKIntersightApiKeyId)
	if keyID == "" {
		return fmt.Errorf("%s is not set", CKIntersightApiKeyId)
	}

	keyFile := gK.String(CKIntersightSecretKey)
	if keyFile == "" {
		return fmt.Errorf("%s is not set", CKIntersightSecretKey)
	}

	var keyData string

	if isKeyData(keyFile) {
		keyData = keyFile
	} else {
		// try doing ~ expansion on the keyFile path
		if expandedKeyFile, err := homedir.Expand(keyFile); err == nil {
			keyFile = expandedKeyFile
		}

		keyDataBytes, err := os.ReadFile(keyFile)
		if err != nil {
			return fmt.Errorf("Unable to key file: %v", err)
		}
		keyData = string(keyDataBytes)
	}

	if gK.Bool(CKIntersightInsecure) {
		log.Trace("Disabled server certificate verification")
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	client.IntersightConfig = intersight.Config{
		KeyID:         keyID,
		KeyData:       keyData,
		BaseTransport: httpTransport,
		Host:          gK.String(CKIntersightFqdn),
	}

	client.IntersightClient, err = intersight.NewClient(client.IntersightConfig)
	if err != nil {
		return fmt.Errorf("Unable to setup Intersight API client: %v", err)
	}

	log.Trace("Finished validateFlags")

	return nil
}
