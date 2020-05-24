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

// ContentComplexType If the given API/device response is a collection of items, each item and its properties can be modeled as a complex type. The types are uniquely named within the grammar and provides the list of parameters to be extracted from each item. Name of the complex type can be used as the type of parameter that represents the complex value.
type ContentComplexType struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// The unique name of this complex type within the grammar specification.
	Name       *string                 `json:"Name,omitempty" yaml:"Name,omitempty"`
	Parameters *[]ContentBaseParameter `json:"Parameters,omitempty" yaml:"Parameters,omitempty"`
}

// NewContentComplexType instantiates a new ContentComplexType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentComplexType() *ContentComplexType {
	this := ContentComplexType{}
	return &this
}

// NewContentComplexTypeWithDefaults instantiates a new ContentComplexType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentComplexTypeWithDefaults() *ContentComplexType {
	this := ContentComplexType{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContentComplexType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentComplexType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContentComplexType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContentComplexType) SetName(v string) {
	o.Name = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ContentComplexType) GetParameters() []ContentBaseParameter {
	if o == nil || o.Parameters == nil {
		var ret []ContentBaseParameter
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentComplexType) GetParametersOk() (*[]ContentBaseParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ContentComplexType) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ContentBaseParameter and assigns it to the Parameters field.
func (o *ContentComplexType) SetParameters(v []ContentBaseParameter) {
	o.Parameters = &v
}

func (o ContentComplexType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Parameters != nil {
		toSerialize["Parameters"] = o.Parameters
	}
	return json.Marshal(toSerialize)
}

type NullableContentComplexType struct {
	value *ContentComplexType
	isSet bool
}

func (v NullableContentComplexType) Get() *ContentComplexType {
	return v.value
}

func (v *NullableContentComplexType) Set(val *ContentComplexType) {
	v.value = val
	v.isSet = true
}

func (v NullableContentComplexType) IsSet() bool {
	return v.isSet
}

func (v *NullableContentComplexType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentComplexType(val *ContentComplexType) *NullableContentComplexType {
	return &NullableContentComplexType{value: val, isSet: true}
}

func (v NullableContentComplexType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentComplexType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
