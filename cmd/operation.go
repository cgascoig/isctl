package main

import (
	"net/http"

	"github.com/cgascoig/isctl/openapi"
)

type Operation interface {
	SetBodyParams(*openapi.APIClient, map[string]interface{}) error

	Execute(*openapi.APIClient, []string, map[string]string) (interface{}, *http.Response, error)
}
