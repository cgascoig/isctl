package main

import (
	"strings"

	"github.com/knadh/koanf/maps"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	flag "github.com/spf13/pflag"
)

var (
	// This is the global koanf instance
	gK = koanf.New(".")
)

const (
	// New config keys
	CKIntersightApiKeyId  = "intersight_api_key_id"
	CKIntersightSecretKey = "intersight_secret_key"
	CKIntersightFqdn      = "intersight_fqdn"
	CKOutputFormat        = "output"
	CKIntersightInsecure  = "intersight_insecure"

	// New cmd flags
	FlagIntersightApiKeyId  = "intersight-api-key-id"
	FlagIntersightSecretKey = "intersight-secret-key"
	FlagIntersightFqdn      = "intersight-fqdn"

	// Legacy / deprecated config keys
	CKKeyID    = "keyID"
	CKKeyFile  = "keyFile"
	CKServer   = "server"
	CKInsecure = "insecure"
)

func initConfig(flags *flag.FlagSet) {
	log.Trace("Starting initConfig")

	configFile := getConfigFilePath(flags)
	log.Tracef("Using config file path: %s", configFile)

	loadDefaults()
	log.Tracef("initConfig: config after loading defaults: %#v", gK.All())

	// Load from config file
	if err := gK.Load(file.Provider(configFile), yaml.Parser(), koanf.WithMergeFunc(aliasReplacingMerge)); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	log.Tracef("initConfig: config after loading config file: %#v", gK.All())

	// Load from environment variables
	gK.Load(env.Provider("INTERSIGHT_", ".", envMapper), nil)
	gK.Load(env.Provider("intersight_", ".", envMapper), nil)
	log.Tracef("initConfig: config after loading environment variables: %#v", gK.All())

	loadFlags(flags)
	log.Tracef("initConfig: config after loading command line flags: %#v", gK.All())

	log.Tracef("Finished initConfig")
}

func getConfigFilePath(flags *flag.FlagSet) string {
	var configFile string

	if c, err := flags.GetString("config"); err == nil && c != "" {
		configFile = c
	} else {
		configFile = "~/.isctl.yaml"
	}

	if expandedConfigFile, err := homedir.Expand(configFile); err == nil {
		configFile = expandedConfigFile
	}
	return configFile
}

func loadDefaults() {
	gK.Load(confmap.Provider(map[string]any{
		CKIntersightFqdn: "intersight.com",
		CKOutputFormat:   "default",
	}, "."), nil)
}

func aliasReplacingMerge(src, dest map[string]any) error {
	// map of alias old:new
	aliases := map[string]string{
		CKKeyID:    CKIntersightApiKeyId,
		CKKeyFile:  CKIntersightSecretKey,
		CKServer:   CKIntersightFqdn,
		CKInsecure: CKIntersightInsecure,
	}

	for old, new := range aliases {
		if val, ok := src[old]; ok {
			src[new] = val
			delete(src, old)
		}

		old = strings.ToLower(old)
		if val, ok := src[old]; ok {
			src[new] = val
			delete(src, old)
		}
	}

	maps.Merge(src, dest)

	return nil
}

func envMapper(s string) string {
	aliases := map[string]string{
		"INTERSIGHT_API_KEY_ID": CKIntersightApiKeyId,
		"intersight_api_key_id": CKIntersightApiKeyId,
		"INTERSIGHT_SECRET_KEY": CKIntersightSecretKey,
		"intersight_secret_key": CKIntersightSecretKey,
		"INTERSIGHT_FQDN":       CKIntersightFqdn,
		"intersight_fqdn":       CKIntersightFqdn,
	}

	if key, ok := aliases[s]; ok {
		return key
	}
	return ""
}

func loadFlags(flags *pflag.FlagSet) {
	aliases := map[string]string{
		CKKeyID:                 CKIntersightApiKeyId,
		CKKeyFile:               CKIntersightSecretKey,
		CKServer:                CKIntersightFqdn,
		CKInsecure:              CKIntersightInsecure,
		FlagIntersightApiKeyId:  CKIntersightApiKeyId,
		FlagIntersightSecretKey: CKIntersightSecretKey,
		FlagIntersightFqdn:      CKIntersightFqdn,
	}

	err := gK.Load(posflag.ProviderWithFlag(flags, ".", gK, func(f *pflag.Flag) (string, any) {
		if new, ok := aliases[f.Name]; ok {
			return new, posflag.FlagVal(flags, f)
		}
		return f.Name, posflag.FlagVal(flags, f)
	}), nil)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
}
