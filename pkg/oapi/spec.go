package oapi

import (
	_ "embed"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

//go:embed "intersight-openapi.json"
var specData []byte

var spec any

func lazyLoadSpec() map[string]any {
	if spec == nil {
		log.Trace("begin unmarshalling JSON spec")
		err := json.Unmarshal(specData, &spec)
		if err != nil {
			log.Errorf("error unmarshaling spec: %v", err)
		}
		log.Trace("finished JSON unmarshal")
	}

	if spec, ok := spec.(map[string]any); ok {
		return spec
	}

	return nil
}
