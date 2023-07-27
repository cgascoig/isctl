package main

import (
	"net/http"
)

// Operation is a generic interface for Intersight Operations
type Operation interface {
	SetBodyParams(*isctlClient, map[string]interface{}) error

	GetReferencedClasses(params map[string]interface{}) ([]string, error)

	Execute(*isctlClient, []string, map[string]string) (interface{}, *http.Response, error)
}
