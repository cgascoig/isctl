package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"net/http/httptrace"
	"net/textproto"
	"os"
	"path/filepath"

	openapi "github.com/cgascoig/intersight-go-sdk/intersight"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile     string
	jsonPathFilter string
	verbose        bool
	httpTrace      bool

	authCtx context.Context

	auxCommandsGenerators []commandGenerator
)

type commandGenerator func(*openapi.APIClient) *cobra.Command

const (
	keyIDConfigKey   = "keyID"
	keyFileConfigKey = "keyFile"

	keyOutputConfigKey = "output"

	serverConfigKey = "server"
)

func main() {
	cobra.OnInitialize(initConfig)

	config := openapi.NewConfiguration()

	client := openapi.NewAPIClient(config)

	rootCmd := GetCommands(client, resultHandler)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.isctl.yaml)")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose logging")
	rootCmd.PersistentFlags().BoolVar(&httpTrace, "trace", false, "Enable HTTP tracing")
	rootCmd.PersistentFlags().MarkHidden("trace")

	rootCmd.PersistentFlags().String(keyIDConfigKey, "", "API Key ID")
	rootCmd.PersistentFlags().String(keyFileConfigKey, "", "API Private Key Filename")

	rootCmd.PersistentFlags().String(serverConfigKey, "intersight.com", "Intersight API Server Address (e.g.\"intersight.com\")")

	rootCmd.PersistentFlags().StringP(keyOutputConfigKey, "o", "default", "Output format. One of: default, yaml, json, table")
	rootCmd.PersistentFlags().StringVar(&jsonPathFilter, "jsonpath", "", "JSONPath filter to apply to the result (e.g. \"$.Name\")")

	viper.BindPFlag(keyIDConfigKey, rootCmd.PersistentFlags().Lookup(keyIDConfigKey))
	viper.BindPFlag(keyFileConfigKey, rootCmd.PersistentFlags().Lookup(keyFileConfigKey))
	viper.BindPFlag(keyOutputConfigKey, rootCmd.PersistentFlags().Lookup(keyOutputConfigKey))
	viper.BindPFlag(serverConfigKey, rootCmd.PersistentFlags().Lookup(serverConfigKey))

	configCmd := &cobra.Command{
		Use:               "configure",
		Run:               configure,
		Short:             "Configure the isctl command",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	}
	rootCmd.AddCommand(configCmd)
	// rootCmd.AddCommand(newCmdApply(client))
	for _, cmdGen := range auxCommandsGenerators {
		rootCmd.AddCommand(cmdGen(client))
	}
	rootCmd.PersistentPreRunE = validateFlags

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalln(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".isctl")
		configFile = filepath.Join(home, ".isctl.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Warn("Config file not found, creating empty default.")
			f, err := os.Create(configFile)
			if err != nil {
				log.Fatalf("ERROR while creating empty config file: %v", err)
			}
			f.Close()
		} else {
			// Config file was found but another error was produced
			log.Fatalf("ERROR Invalid config file: %v", err)
		}
	}
}

func configure(cmd *cobra.Command, args []string) {
	scanner := bufio.NewScanner(os.Stdin)

	// configure keyID
	fmt.Printf("%s is currently '%s'\n", keyIDConfigKey, viper.GetString(keyIDConfigKey))
	fmt.Printf("Enter new %s or press Enter to keep existing: ", keyIDConfigKey)
	scanner.Scan()
	if input := scanner.Text(); input != "" {
		viper.Set(keyIDConfigKey, input)
	}

	// configure key file name
	fmt.Printf("key filename is currently '%s'\n", viper.GetString(keyFileConfigKey))
	fmt.Printf("Enter new key file name or press Enter to keep existing: ")
	scanner.Scan()
	if input := scanner.Text(); input != "" {
		viper.Set(keyFileConfigKey, input)
	}

	// configure server
	fmt.Printf("Intersight API server is currently '%s'\n", viper.GetString(serverConfigKey))
	fmt.Printf("Enter new Intersight API server or press Enter to keep existing: ")
	scanner.Scan()
	if input := scanner.Text(); input != "" {
		viper.Set(serverConfigKey, input)
	}

	log.Println("Writing config file")
	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("Error occurred writing config file: %v", err)
	}
}

func validateFlags(cmd *cobra.Command, args []string) error {
	// Setup logging
	if verbose {
		log.SetLevel(log.DebugLevel)
		log.Debug("Logging level set to debug(verbose)")
	}

	var err error

	keyID := viper.GetString(keyIDConfigKey)
	if keyID == "" {
		return fmt.Errorf("%s is not set", keyIDConfigKey)
	}

	keyFile := viper.GetString(keyFileConfigKey)
	if keyFile == "" {
		return fmt.Errorf("%s is not set", keyFileConfigKey)
	}
	// try doing ~ expansion on the keyFile path
	if expandedKeyFile, err := homedir.Expand(keyFile); err == nil {
		keyFile = expandedKeyFile
	}

	authConfig := openapi.HttpSignatureAuth{
		KeyId:          keyID,
		PrivateKeyPath: keyFile,

		// Passphrase:           "my-passphrase",
		SigningScheme: openapi.HttpSigningSchemeRsaSha256,
		SignedHeaders: []string{
			openapi.HttpSignatureParameterRequestTarget, // The special (request-target) parameter expresses the HTTP request target.
			"Host",   // The Host request header specifies the domain name of the server, and optionally the TCP port number.
			"Date",   // The date and time at which the message was originated.
			"Digest", // A cryptographic digest of the request body.
		},
		SigningAlgorithm: openapi.HttpSigningAlgorithmRsaPKCS1v15,
	}

	authCtx, err = authConfig.ContextWithValue(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to create request context with authentication: %v", err)
	}

	authCtx = context.WithValue(authCtx, openapi.ContextServerVariables, map[string]string{
		"server": viper.GetString(serverConfigKey),
	})

	if httpTrace {
		trace := &httptrace.ClientTrace{
			GetConn: func(hostPort string) {
				log.Printf("Get Conn: %v\n", hostPort)
			},
			GotConn: func(connInfo httptrace.GotConnInfo) {
				log.Printf("Got Conn: %+v\n", connInfo)
			},

			PutIdleConn: func(err error) {
				log.Printf("PutIdleConn: err: %v\n", err)
			},

			GotFirstResponseByte: func() {
				log.Printf("GotFirstResponseByte\n")
			},

			Got100Continue: func() {
				log.Printf("Got100Continue\n")
			},

			Got1xxResponse: func(code int, header textproto.MIMEHeader) error {
				log.Printf("Got1xxResponse code: %v header %v\n", code, header)
				return nil
			},

			DNSStart: func(dsi httptrace.DNSStartInfo) {
				log.Printf("DNSStart %v\n", dsi)
			},

			DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
				log.Printf("DNS Info: %+v\n", dnsInfo)
			},

			ConnectStart: func(network, addr string) {
				log.Printf("ConnectStart network: %v, addr: %v\n", network, addr)
			},

			ConnectDone: func(network, addr string, err error) {
				log.Printf("ConnectDone network: %v addr %v err %v\n", network, addr, err)
			},

			TLSHandshakeStart: func() {
				log.Printf("TLSHandshakeStart\n")
			},

			TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
				log.Printf("TLSHandshakeDone connectionstate %v err %v\n", cs, err)
			},

			WroteHeaderField: func(key string, value []string) {
				log.Printf("WroteHeaderField key %v value %v\n", key, value)
			},

			WroteHeaders: func() {
				log.Printf("WroteHeaders\n")
			},

			Wait100Continue: func() {
				log.Printf("Wait100Continue\n")
			},

			WroteRequest: func(wri httptrace.WroteRequestInfo) {
				log.Printf("WroteRequest %v\n", wri)
			},
		}

		authCtx = httptrace.WithClientTrace(authCtx, trace)
	}

	return nil
}
