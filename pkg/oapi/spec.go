package oapi

import (
	_ "embed"
	"encoding/json"
	"log/slog"
)

//go:embed "intersight-openapi.json"
var specData []byte

var spec any

func lazyLoadSpec() map[string]any {
	if spec == nil {
		slog.Debug("begin unmarshalling JSON spec")
		err := json.Unmarshal(specData, &spec)
		if err != nil {
			slog.Error("error unmarshaling spec", "error", err)
		}
		slog.Debug("finished JSON unmarshal")
	}

	if spec, ok := spec.(map[string]any); ok {
		return spec
	}

	return nil
}
