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

// TamPsirtSeverityAllOf Definition of the list of properties defined in 'tam.PsirtSeverity', excluding properties defined in parent classes.
type TamPsirtSeverityAllOf struct {
	// Severity level associated with the security advisory.
	Level *string `json:"Level,omitempty" yaml:"Level,omitempty"`
}

// NewTamPsirtSeverityAllOf instantiates a new TamPsirtSeverityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamPsirtSeverityAllOf() *TamPsirtSeverityAllOf {
	this := TamPsirtSeverityAllOf{}
	var level string = "critical"
	this.Level = &level
	return &this
}

// NewTamPsirtSeverityAllOfWithDefaults instantiates a new TamPsirtSeverityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamPsirtSeverityAllOfWithDefaults() *TamPsirtSeverityAllOf {
	this := TamPsirtSeverityAllOf{}
	var level string = "critical"
	this.Level = &level
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *TamPsirtSeverityAllOf) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamPsirtSeverityAllOf) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *TamPsirtSeverityAllOf) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *TamPsirtSeverityAllOf) SetLevel(v string) {
	o.Level = &v
}

func (o TamPsirtSeverityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Level != nil {
		toSerialize["Level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableTamPsirtSeverityAllOf struct {
	value *TamPsirtSeverityAllOf
	isSet bool
}

func (v NullableTamPsirtSeverityAllOf) Get() *TamPsirtSeverityAllOf {
	return v.value
}

func (v *NullableTamPsirtSeverityAllOf) Set(val *TamPsirtSeverityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamPsirtSeverityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamPsirtSeverityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamPsirtSeverityAllOf(val *TamPsirtSeverityAllOf) *NullableTamPsirtSeverityAllOf {
	return &NullableTamPsirtSeverityAllOf{value: val, isSet: true}
}

func (v NullableTamPsirtSeverityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamPsirtSeverityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
