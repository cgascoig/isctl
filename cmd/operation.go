package main

import (
	"net/http"

	openapi "github.com/cgascoig/intersight-go-sdk/intersight"
)

// Operation is a generic interface for Intersight Operations
type Operation interface {
	SetBodyParams(*openapi.APIClient, map[string]interface{}) error

	GetReferencedClasses(params map[string]interface{}) ([]string, error)

	Execute(*openapi.APIClient, []string, map[string]string) (interface{}, *http.Response, error)
}
