/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-04-17T15:33:06-07:00.
 *
 * API version: 1.0.9-1628
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

// ServerConfigResultEntryResponse - The response body of a HTTP GET request for the 'server.ConfigResultEntry' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'server.ConfigResultEntry' resources.
type ServerConfigResultEntryResponse struct {
	MoDocumentCount             *MoDocumentCount
	ServerConfigResultEntryList *ServerConfigResultEntryList
}

// MoDocumentCountAsServerConfigResultEntryResponse is a convenience function that returns MoDocumentCount wrapped in ServerConfigResultEntryResponse
func MoDocumentCountAsServerConfigResultEntryResponse(v *MoDocumentCount) ServerConfigResultEntryResponse {
	return ServerConfigResultEntryResponse{MoDocumentCount: v}
}

// ServerConfigResultEntryListAsServerConfigResultEntryResponse is a convenience function that returns ServerConfigResultEntryList wrapped in ServerConfigResultEntryResponse
func ServerConfigResultEntryListAsServerConfigResultEntryResponse(v *ServerConfigResultEntryList) ServerConfigResultEntryResponse {
	return ServerConfigResultEntryResponse{ServerConfigResultEntryList: v}
}

// Unmarshl JSON data into one of the pointers in the struct
func (dst *ServerConfigResultEntryResponse) UnmarshalJSON(data []byte) error {
	var err error

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		return err
	}
	if v, ok := unmarshaled["ObjectType"]; ok {
		switch v {
		case "mo.DocumentCount":
			var result *MoDocumentCount = &MoDocumentCount{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.MoDocumentCount = result
			return nil
		case "server.ConfigResultEntry.List":
			var result *ServerConfigResultEntryList = &ServerConfigResultEntryList{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.ServerConfigResultEntryList = result
			return nil
		default:
			return fmt.Errorf("No oneOf model has 'ObjectType' equal to %s", v)
		}
	} else {
		return fmt.Errorf("Discriminator property 'ObjectType' not found in unmarshaled payload: %+v", unmarshaled)
	}

}

// Marshl data from the first non-nil pointers in the struct to JSON
func (src *ServerConfigResultEntryResponse) MarshalJSON() ([]byte, error) {
	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.ServerConfigResultEntryList != nil {
		return json.Marshal(&src.ServerConfigResultEntryList)
	}

	return nil, nil // no data in oneOf schemas
}

// Marshl data from the first non-nil pointers in the struct to YAML
func (src *ServerConfigResultEntryResponse) MarshalYAML() ([]byte, error) {
	if src.MoDocumentCount != nil {
		return yaml.Marshal(&src.MoDocumentCount)
	}

	if src.ServerConfigResultEntryList != nil {
		return yaml.Marshal(&src.ServerConfigResultEntryList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServerConfigResultEntryResponse) GetActualInstance() interface{} {
	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.ServerConfigResultEntryList != nil {
		return obj.ServerConfigResultEntryList
	}

	// all schemas are nil
	return nil
}
