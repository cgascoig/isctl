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

// WorkflowDefaultValue Captures default vales for a data type.
type WorkflowDefaultValue struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// Override the default value provided for the data type. When true, allow the user to enter value for the data type.
	Override *bool `json:"Override,omitempty" yaml:"Override,omitempty"`
	// Default value for the data type. If default value was provided and the input was required the default value will be used as the input.
	Value *map[string]interface{} `json:"Value,omitempty" yaml:"Value,omitempty"`
}

// NewWorkflowDefaultValue instantiates a new WorkflowDefaultValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowDefaultValue() *WorkflowDefaultValue {
	this := WorkflowDefaultValue{}
	return &this
}

// NewWorkflowDefaultValueWithDefaults instantiates a new WorkflowDefaultValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowDefaultValueWithDefaults() *WorkflowDefaultValue {
	this := WorkflowDefaultValue{}
	return &this
}

// GetOverride returns the Override field value if set, zero value otherwise.
func (o *WorkflowDefaultValue) GetOverride() bool {
	if o == nil || o.Override == nil {
		var ret bool
		return ret
	}
	return *o.Override
}

// GetOverrideOk returns a tuple with the Override field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDefaultValue) GetOverrideOk() (*bool, bool) {
	if o == nil || o.Override == nil {
		return nil, false
	}
	return o.Override, true
}

// HasOverride returns a boolean if a field has been set.
func (o *WorkflowDefaultValue) HasOverride() bool {
	if o != nil && o.Override != nil {
		return true
	}

	return false
}

// SetOverride gets a reference to the given bool and assigns it to the Override field.
func (o *WorkflowDefaultValue) SetOverride(v bool) {
	o.Override = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WorkflowDefaultValue) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDefaultValue) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WorkflowDefaultValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *WorkflowDefaultValue) SetValue(v map[string]interface{}) {
	o.Value = &v
}

func (o WorkflowDefaultValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Override != nil {
		toSerialize["Override"] = o.Override
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowDefaultValue struct {
	value *WorkflowDefaultValue
	isSet bool
}

func (v NullableWorkflowDefaultValue) Get() *WorkflowDefaultValue {
	return v.value
}

func (v *NullableWorkflowDefaultValue) Set(val *WorkflowDefaultValue) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowDefaultValue) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowDefaultValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowDefaultValue(val *WorkflowDefaultValue) *NullableWorkflowDefaultValue {
	return &NullableWorkflowDefaultValue{value: val, isSet: true}
}

func (v NullableWorkflowDefaultValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowDefaultValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
