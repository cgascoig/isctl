package main

import (
	"net/http"

	"github.com/cgascoig/isctl/openapi"
)

// Operation is a generic interface for Intersight Operations
type Operation interface {
	SetBodyParams(*openapi.APIClient, map[string]interface{}) error

	GetReferencedClasses(params map[string]interface{}) ([]string, error)

	Execute(*openapi.APIClient, []string, map[string]string) (interface{}, *http.Response, error)
}
