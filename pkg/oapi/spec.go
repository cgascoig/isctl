package oapi

import (
	_ "embed"
	"encoding/json"
	"sync"

	log "github.com/sirupsen/logrus"
)

//go:embed "intersight-openapi.json"
var specData []byte

var (
	spec     map[string]any
	specOnce sync.Once
)

func lazyLoadSpec() map[string]any {
	specOnce.Do(func() {
		log.Trace("begin unmarshalling JSON spec")
		if err := json.Unmarshal(specData, &spec); err != nil {
			log.Fatalf("Error loading spec: %v", err)
		}
		log.Trace("finished JSON unmarshal")
	})
	return spec
}
