package main

import (
	"os"
	"path/filepath"
	"regexp"
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
	flag "github.com/spf13/pflag"
)

var (
	// This is the global koanf instance
	gK = koanf.New(".")
)

const (
	// New config keys
	CKIntersightApiKeyId          = "intersight_api_key_id"
	CKIntersightSecretKey         = "intersight_secret_key"
	CKIntersightFqdn              = "intersight_fqdn"
	CKOutputFormat                = "output"
	CKIntersightInsecure          = "intersight_insecure"
	CKIntersightProxy             = "intersight_proxy"
	CKIntersightClientId          = "intersight_client_id"
	CKIntersightClientSecret      = "intersight_client_secret"
	CKIntersightTokenUrl          = "intersight_token_url"
	CKIntersightTokenCachePath    = "intersight_token_cache_path"
	CKIntersightDisableTokenCache = "intersight_disable_token_cache"
	CKReadableMoRefs              = "readable_morefs"
	CKIncludeEmptyFields          = "include_empty_fields"

	// New cmd flags
	FlagIntersightApiKeyId          = "intersight-api-key-id"
	FlagIntersightSecretKey         = "intersight-secret-key"
	FlagIntersightFqdn              = "intersight-fqdn"
	FlagIntersightProxy             = "intersight-proxy"
	FlagIntersightClientId          = "intersight-client-id"
	FlagIntersightClientSecret      = "intersight-client-secret"
	FlagIntersightTokenUrl          = "intersight-token-url"
	FlagIntersightTokenCachePath    = "intersight-token-cache-path"
	FlagIntersightDisableTokenCache = "intersight-disable-token-cache"
	FlagReadableMoRefs              = "readable-morefs"
	FlagIncludeEmptyFields          = "include-empty-fields"

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
	if _, err := os.Stat(configFile); err == nil {
		if err := gK.Load(file.Provider(configFile), yaml.Parser(), koanf.WithMergeFunc(aliasReplacingMerge)); err != nil {
			log.Fatalf("error loading config: %v", err)
		}
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
	} else if c := os.Getenv("ISCONFIG"); c != "" {
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
		CKIntersightFqdn:              "intersight.com",
		CKOutputFormat:                "default",
		CKIntersightApiKeyId:          "",
		CKIntersightSecretKey:         "",
		CKIntersightInsecure:          false,
		CKIntersightProxy:             "",
		CKIntersightClientId:          "",
		CKIntersightClientSecret:      "",
		CKIntersightTokenUrl:          "",
		CKIntersightTokenCachePath:    "",
		CKIntersightDisableTokenCache: false,
		CKReadableMoRefs:              false,
		CKIncludeEmptyFields:          false,
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
		"INTERSIGHT_API_KEY_ID":          CKIntersightApiKeyId,
		"intersight_api_key_id":          CKIntersightApiKeyId,
		"INTERSIGHT_SECRET_KEY":          CKIntersightSecretKey,
		"intersight_secret_key":          CKIntersightSecretKey,
		"INTERSIGHT_FQDN":                CKIntersightFqdn,
		"intersight_fqdn":                CKIntersightFqdn,
		"INTERSIGHT_PROXY":               CKIntersightProxy,
		"intersight_proxy":               CKIntersightProxy,
		"INTERSIGHT_CLIENT_ID":           CKIntersightClientId,
		"intersight_client_id":           CKIntersightClientId,
		"INTERSIGHT_CLIENT_SECRET":       CKIntersightClientSecret,
		"intersight_client_secret":       CKIntersightClientSecret,
		"INTERSIGHT_TOKEN_URL":           CKIntersightTokenUrl,
		"intersight_token_url":           CKIntersightTokenUrl,
		"INTERSIGHT_TOKEN_CACHE_PATH":    CKIntersightTokenCachePath,
		"intersight_token_cache_path":    CKIntersightTokenCachePath,
		"INTERSIGHT_DISABLE_TOKEN_CACHE": CKIntersightDisableTokenCache,
		"intersight_disable_token_cache": CKIntersightDisableTokenCache,
	}

	if key, ok := aliases[s]; ok {
		return key
	}
	return ""
}

func loadFlags(flags *flag.FlagSet) {
	aliases := map[string]string{
		CKKeyID:                         CKIntersightApiKeyId,
		CKKeyFile:                       CKIntersightSecretKey,
		CKServer:                        CKIntersightFqdn,
		CKInsecure:                      CKIntersightInsecure,
		FlagIntersightApiKeyId:          CKIntersightApiKeyId,
		FlagIntersightSecretKey:         CKIntersightSecretKey,
		FlagIntersightFqdn:              CKIntersightFqdn,
		FlagIntersightProxy:             CKIntersightProxy,
		FlagIntersightClientId:          CKIntersightClientId,
		FlagIntersightClientSecret:      CKIntersightClientSecret,
		FlagIntersightTokenUrl:          CKIntersightTokenUrl,
		FlagIntersightTokenCachePath:    CKIntersightTokenCachePath,
		FlagIntersightDisableTokenCache: CKIntersightDisableTokenCache,
		FlagReadableMoRefs:              CKReadableMoRefs,
		FlagIncludeEmptyFields:          CKIncludeEmptyFields,
	}

	err := gK.Load(posflag.ProviderWithFlag(flags, ".", gK, func(f *flag.Flag) (string, any) {
		if new, ok := aliases[f.Name]; ok {
			return new, posflag.FlagVal(flags, f)
		}
		return f.Name, posflag.FlagVal(flags, f)
	}), nil)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
}

func isKeyData(s string) bool {
	re := regexp.MustCompile(`(?s)^\s*-----BEGIN[A-Z ]*KEY-----.*-----END[A-Z ]*KEY-----\s*$`)
	return re.MatchString(s)
}

// getDefaultTokenCachePath returns the default token cache path following XDG Base Directory spec.
func getDefaultTokenCachePath() string {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		home, _ := homedir.Dir()
		configHome = filepath.Join(home, ".config")
	}
	return filepath.Join(configHome, "isctl", "token-cache.json")
}
