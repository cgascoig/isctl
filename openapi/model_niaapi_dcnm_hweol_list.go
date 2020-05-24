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
)

// NiaapiDcnmHweolList This resource list is returned as a response to a HTTP GET request that does not include a specific resource identifier.
type NiaapiDcnmHweolList struct {
	MoBaseResponse `yaml:"MoBaseResponse,inline"`
	// The total number of 'niaapi.DcnmHweol' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
	Count *int32 `json:"Count,omitempty" yaml:"Count,omitempty"`
	// The array of 'niaapi.DcnmHweol' resources matching the request.
	Results *[]NiaapiDcnmHweol `json:"Results,omitempty" yaml:"Results,omitempty"`
}

// NewNiaapiDcnmHweolList instantiates a new NiaapiDcnmHweolList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiDcnmHweolList() *NiaapiDcnmHweolList {
	this := NiaapiDcnmHweolList{}
	return &this
}

// NewNiaapiDcnmHweolListWithDefaults instantiates a new NiaapiDcnmHweolList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiDcnmHweolListWithDefaults() *NiaapiDcnmHweolList {
	this := NiaapiDcnmHweolList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *NiaapiDcnmHweolList) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiDcnmHweolList) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *NiaapiDcnmHweolList) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *NiaapiDcnmHweolList) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *NiaapiDcnmHweolList) GetResults() []NiaapiDcnmHweol {
	if o == nil || o.Results == nil {
		var ret []NiaapiDcnmHweol
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiDcnmHweolList) GetResultsOk() (*[]NiaapiDcnmHweol, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *NiaapiDcnmHweolList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []NiaapiDcnmHweol and assigns it to the Results field.
func (o *NiaapiDcnmHweolList) SetResults(v []NiaapiDcnmHweol) {
	o.Results = &v
}

func (o NiaapiDcnmHweolList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseResponse, errMoBaseResponse := json.Marshal(o.MoBaseResponse)
	if errMoBaseResponse != nil {
		return []byte{}, errMoBaseResponse
	}
	errMoBaseResponse = json.Unmarshal([]byte(serializedMoBaseResponse), &toSerialize)
	if errMoBaseResponse != nil {
		return []byte{}, errMoBaseResponse
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableNiaapiDcnmHweolList struct {
	value *NiaapiDcnmHweolList
	isSet bool
}

func (v NullableNiaapiDcnmHweolList) Get() *NiaapiDcnmHweolList {
	return v.value
}

func (v *NullableNiaapiDcnmHweolList) Set(val *NiaapiDcnmHweolList) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiDcnmHweolList) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiDcnmHweolList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiDcnmHweolList(val *NiaapiDcnmHweolList) *NullableNiaapiDcnmHweolList {
	return &NullableNiaapiDcnmHweolList{value: val, isSet: true}
}

func (v NullableNiaapiDcnmHweolList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiDcnmHweolList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
