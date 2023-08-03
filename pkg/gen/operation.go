package gen

import (
	"net/http"

	"github.com/cgascoig/isctl/pkg/util"
)

// Operation is a generic interface for Intersight Operations
type Operation interface {
	SetBodyParams(*util.IsctlClient, map[string]interface{}) error

	GetReferencedClasses(params map[string]interface{}) ([]string, error)

	Execute(*util.IsctlClient, []string, map[string]string) (interface{}, *http.Response, error)
}
