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

// TamIdentifiersAllOf Definition of the list of properties defined in 'tam.Identifiers', excluding properties defined in parent classes.
type TamIdentifiersAllOf struct {
	// Name of the filter paramter.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// Value of the filter paramter.
	Value *string `json:"Value,omitempty" yaml:"Value,omitempty"`
}

// NewTamIdentifiersAllOf instantiates a new TamIdentifiersAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamIdentifiersAllOf() *TamIdentifiersAllOf {
	this := TamIdentifiersAllOf{}
	return &this
}

// NewTamIdentifiersAllOfWithDefaults instantiates a new TamIdentifiersAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamIdentifiersAllOfWithDefaults() *TamIdentifiersAllOf {
	this := TamIdentifiersAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TamIdentifiersAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamIdentifiersAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TamIdentifiersAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TamIdentifiersAllOf) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TamIdentifiersAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamIdentifiersAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TamIdentifiersAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TamIdentifiersAllOf) SetValue(v string) {
	o.Value = &v
}

func (o TamIdentifiersAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableTamIdentifiersAllOf struct {
	value *TamIdentifiersAllOf
	isSet bool
}

func (v NullableTamIdentifiersAllOf) Get() *TamIdentifiersAllOf {
	return v.value
}

func (v *NullableTamIdentifiersAllOf) Set(val *TamIdentifiersAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamIdentifiersAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamIdentifiersAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamIdentifiersAllOf(val *TamIdentifiersAllOf) *NullableTamIdentifiersAllOf {
	return &NullableTamIdentifiersAllOf{value: val, isSet: true}
}

func (v NullableTamIdentifiersAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamIdentifiersAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
