package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime/pprof"

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
	traceEnvName      = "ISCTL_TRACE"
	cpuProfileEnvName = "ISCTL_CPU_PROFILE"
)

func main() {
	if os.Getenv(traceEnvName) != "" {
		log.SetLevel(log.TraceLevel)
	}

	log.Trace("isctl starting")

	if cpuProfile := os.Getenv(cpuProfileEnvName); cpuProfile != "" {
		f, err := os.Create(cpuProfile)
		if err != nil {
			log.Fatalf("creating cpu profile file: %v", err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	rootCmd := gen.GetCommands(client, resultHandler)
	rootCmd.Use = "isctl"

	log.Trace("Got generated commands")

	rootCmd.PersistentFlags().String("config", "", "config file (default is $HOME/.isctl.yaml)")

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose logging")

	// New config keys
	rootCmd.PersistentFlags().String(FlagIntersightApiKeyId, "", "Intersight API Key ID")
	rootCmd.PersistentFlags().String(FlagIntersightSecretKey, "", "Intersight Secret Key (filename)")
	rootCmd.PersistentFlags().String(FlagIntersightFqdn, "", "Intersight API FQDN (default intersight.com)")
	rootCmd.PersistentFlags().String(FlagIntersightProxy, "", "HTTP Proxy for Intersight API requests")
	rootCmd.PersistentFlags().String(FlagIntersightClientId, "", "Intersight Client ID")
	rootCmd.PersistentFlags().String(FlagIntersightClientSecret, "", "Intersight Client Secret")
	rootCmd.PersistentFlags().String(FlagIntersightTokenUrl, "", "Intersight Token URL")
	rootCmd.PersistentFlags().String(FlagIntersightTokenCachePath, "", "Path to cache OAuth tokens (default: $XDG_CONFIG_HOME/isctl/token-cache.json)")
	rootCmd.PersistentFlags().Bool(FlagIntersightDisableTokenCache, false, "Disable OAuth token caching")

	rootCmd.PersistentFlags().String(CKKeyID, "", "API Key ID [deprecated]")
	rootCmd.PersistentFlags().String(CKKeyFile, "", "API Private Key Filename [deprecated]")

	rootCmd.PersistentFlags().String(CKServer, "intersight.com", "Intersight API Server Address (e.g.\"intersight.com\") [deprecated]")
	rootCmd.PersistentFlags().BoolP(CKInsecure, "k", false, "Allow insecure server connections (disable SSL certificate validation)")

	rootCmd.PersistentFlags().StringP(CKOutputFormat, "o", "default", `Output format. One of default|yaml|json|table|jsonpath|custom-columns|csv|xlsx. Examples:
	Get Name attribute from all NTP policies: isctl get ntp policy -o jsonpath="[*].Name"
	Table with just Name and Enabled attributes: isctl get ntp policy -o custom-columns=NAME:.Name,ENABLED:.Enabled
	Output to CSV with default columns: isctl get ntp policy -o csv > out.csv
	Output to CSV with custom columns: isctl get ntp policy -o csv=NAME:.Name,ENABLED:.Enabled > out.csv
	Output to XLSX: isctl get ntp policy -o xlsx=out.xlsx
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

	fmt.Printf("Let's setup authentication. You can either configure %s and %s (for API Key authentication), or you can configure %s and %s (for OAuth authentication).\n", CKIntersightApiKeyId, CKIntersightSecretKey, CKIntersightClientId, CKIntersightClientSecret)

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

	// configure client ID
	fmt.Printf("%s is currently '%s'\n", CKIntersightClientId, gK.String(CKIntersightClientId))
	fmt.Printf("Enter new OAuth client ID or press Enter to keep existing: ")
	scanner.Scan()
	if input := scanner.Text(); input != "" {
		gK.Set(CKIntersightClientId, input)
	}

	// configure client secret
	fmt.Printf("%s is currently '%s'\n", CKIntersightClientSecret, gK.String(CKIntersightClientSecret))
	fmt.Printf("Enter new OAuth client secret or press Enter to keep existing: ")
	scanner.Scan()
	if input := scanner.Text(); input != "" {
		gK.Set(CKIntersightClientSecret, input)
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

	// Determine Authentication Method
	keyID := gK.String(CKIntersightApiKeyId)
	keyFile := gK.String(CKIntersightSecretKey)
	clientID := gK.String(CKIntersightClientId)
	clientSecret := gK.String(CKIntersightClientSecret)
	tokenUrl := gK.String(CKIntersightTokenUrl)

	var keyData string
	var useApiKey bool

	if keyID != "" {
		if keyFile == "" {
			return fmt.Errorf("%s is set but %s is missing", CKIntersightApiKeyId, CKIntersightSecretKey)
		}
		useApiKey = true
	} else if clientID != "" {
		if clientSecret == "" {
			return fmt.Errorf("%s is set but %s is missing", CKIntersightClientId, CKIntersightClientSecret)
		}
		useApiKey = false
	} else {
		return fmt.Errorf("either %s and %s OR %s and %s must be set", CKIntersightApiKeyId, CKIntersightSecretKey, CKIntersightClientId, CKIntersightClientSecret)
	}

	if useApiKey {
		if isKeyData(keyFile) {
			keyData = keyFile
		} else {
			// try doing ~ expansion on the keyFile path
			if expandedKeyFile, err := homedir.Expand(keyFile); err == nil {
				keyFile = expandedKeyFile
			}

			keyDataBytes, err := os.ReadFile(keyFile)
			if err != nil {
				return fmt.Errorf("unable to read key file: %v", err)
			}
			keyData = string(keyDataBytes)
		}
	}

	if gK.Bool(CKIntersightInsecure) {
		log.Trace("Disabled server certificate verification")
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	if gK.String(CKIntersightProxy) != "" {
		u, err := url.Parse(gK.String(CKIntersightProxy))
		if err != nil {
			log.Errorf("unable to parse proxy URL: %v", err)
		} else {
			http.DefaultTransport.(*http.Transport).Proxy = http.ProxyURL(u)
		}
	}

	// Determine token cache path
	tokenCachePath := gK.String(CKIntersightTokenCachePath)
	if tokenCachePath == "" && !gK.Bool(CKIntersightDisableTokenCache) {
		tokenCachePath = getDefaultTokenCachePath()
	}

	client.IntersightConfig = intersight.Config{
		KeyID:             keyID,
		KeyData:           keyData,
		ClientID:          clientID,
		ClientSecret:      clientSecret,
		TokenURL:          tokenUrl,
		BaseTransport:     httpTransport,
		Host:              gK.String(CKIntersightFqdn),
		TokenCachePath:    tokenCachePath,
		DisableTokenCache: gK.Bool(CKIntersightDisableTokenCache),
	}

	client.IntersightClient, err = intersight.NewClient(client.IntersightConfig)
	if err != nil {
		return fmt.Errorf("unable to setup Intersight API client: %v", err)
	}

	log.Trace("Finished validateFlags")

	return nil
}
