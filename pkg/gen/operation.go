package gen

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/cgascoig/isctl/pkg/oapi"
	"github.com/cgascoig/isctl/pkg/util"
	log "github.com/sirupsen/logrus"
)

// // Operation is a generic interface for Intersight Operations
// type Operation interface {
// 	SetBodyParams(*util.IsctlClient, map[string]interface{}) error

// 	GetReferencedClasses(params map[string]interface{}) ([]string, error)

// 	Execute(*util.IsctlClient, []string, map[string]string) (interface{}, error)
// }

type Operation struct {
	classID string
	method  string

	operation *oapi.Operation

	body any
}

func newOperationFromOperation(operation *oapi.Operation) *Operation {
	if operation == nil {
		return nil
	}
	return &Operation{
		operation: operation,
		classID:   oapi.SchemaNameToClassId(operation.OperationClassID()),
		method:    operation.HTTPMethod,
	}
}

func resolveReferences(client *util.IsctlClient, body any) (any, error) {
	log.Debugf("Attempting to resolve references for %#v", body)
	switch body := body.(type) {
	case map[string]any:
		for k, v := range body {
			new, err := resolveReferences(client, v)
			if err != nil {
				return nil, err
			}
			body[k] = new
		}
		return body, nil
	case []any:
		for i, v := range body {
			new, err := resolveReferences(client, v)
			if err != nil {
				return nil, err
			}
			body[i] = new
		}
		return body, nil
	case *oapi.MoRef:
		newBody, err := GetMoMoRef(client, body)
		if err != nil {
			return nil, err
		}
		return newBody, nil
	}

	return body, nil
}

func (o *Operation) SetBodyParams(client *util.IsctlClient, params map[string]any) error {
	if len(params) == 0 {
		if strings.EqualFold("get", o.method) {
			o.body = nil
		} else {
			o.body = map[string]any{
				"ClassId":    o.classID,
				"ObjectType": o.classID,
			}
		}

		return nil
	}

	oapi.CanonicaliseMoRefs(&params, o.classID)

	newBody, err := resolveReferences(client, params)
	if err != nil {
		return err
	}
	o.body = newBody

	log.Tracef("SetBodyParams: after relationship resolving: %#v", o.body)

	return nil
}

func (o *Operation) Execute(client *util.IsctlClient, args []string, queryParams map[string]string) (any, error) {
	if o.operation == nil {
		return nil, fmt.Errorf("cannot execute operation - operation not prepared yet")
	}
	js := []byte("")
	if o.body != nil {
		var err error
		js, err = json.Marshal(o.body)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal request body")
		}
	}

	path, err := ReplaceArgs(o.operation.Path, args)
	if err != nil {
		return nil, err
	}
	queryString := EncodeQueryParams(queryParams)
	if queryString != "" {
		path = fmt.Sprintf("%s?%s", path, queryString)
	}
	res, err := client.IntersightClient.Call(o.operation.HTTPMethod, path, js)
	return res, err

}

func GetOperationForClassID(method, classID string) *Operation {
	return newOperationFromOperation(oapi.FindOperation(method, classID))
	// return newOperation(method, classID)
}

func GetGetOperationForClassID(classID string) *Operation {
	return GetOperationForClassID("get", classID)
}

func GetDeleteOperationForClassID(classID string) *Operation {
	return GetOperationForClassID("delete", classID)
}

func GetCreateOperationForClassID(classID string) *Operation {
	return GetOperationForClassID("post", classID)
}

func GetUpdateOperationForClassID(classID string) *Operation {
	return GetOperationForClassID("patch", classID)
}

func GetOperationForRelationship(relationship string) *Operation {
	classID := getClassIDFromRelationship(relationship)

	return GetGetOperationForClassID(classID)
}

func getClassIDFromRelationship(relationship string) string {
	r := regexp.MustCompile(`\.Relationship$`)
	return r.ReplaceAllString(relationship, "")
}

func GetReferencedClasses(params any) []string {
	switch params := params.(type) {
	case map[string]any:
		ret := []string{}
		for _, v := range params {
			ret = append(ret, GetReferencedClasses(v)...)
		}
		return ret
	case []any:
		ret := []string{}
		for _, v := range params {
			ret = append(ret, GetReferencedClasses(v)...)
		}
		return ret
	case *oapi.MoRef:
		return []string{getClassIDFromRelationship(params.RelationshipType)}
	}

	return []string{}
}
