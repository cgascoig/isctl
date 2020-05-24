package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cgascoig/isctl/openapi"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
	authCtx    context.Context
)

const (
	keyIDConfigKey   = "keyID"
	keyFileConfigKey = "keyFile"
)

func main() {
	cobra.OnInitialize(initConfig)

	config := openapi.NewConfiguration()

	client := openapi.NewAPIClient(config)

	rootCmd := GetCommands(client, resultHandler)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.isctl.yaml)")

	rootCmd.PersistentFlags().String(keyIDConfigKey, "", "API Key ID")
	rootCmd.PersistentFlags().String(keyFileConfigKey, "", "API Private Key Filename")
	rootCmd.PersistentFlags().StringP("output", "o", "default", "Output format. One of: default, yaml, json, jsonpath=...")
	viper.BindPFlag(keyIDConfigKey, rootCmd.PersistentFlags().Lookup(keyIDConfigKey))
	viper.BindPFlag(keyFileConfigKey, rootCmd.PersistentFlags().Lookup(keyFileConfigKey))
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))

	configCmd := &cobra.Command{
		Use: "configure",
		Run: configure,
	}
	rootCmd.AddCommand(configCmd)
	rootCmd.PersistentPreRunE = validateFlags

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func resultHandler(result interface{}, httpResponse *http.Response, err error) {
	if err != nil {
		log.Printf("HTTP Response status text: %v", httpResponse.Status)
		log.Fatalf("ERROR: %v\n", err)
	}

	// result = stripEnvelope(result)
	printResultHuman(result)
	// printResultYAML(result)
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
			log.Println("Config file not found, creating empty default.")
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

	log.Println("Writing config file")
	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("Error occurred writing config file: %v", err)
	}
}

func validateFlags(cmd *cobra.Command, args []string) error {
	// Skip validation for config command
	if cmd.Use == "configure" {
		return nil
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

	return nil
}
