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

// WorkflowTargetDataType Data type to capture a target endpoint or device.
type WorkflowTargetDataType struct {
	WorkflowBaseDataType `yaml:"WorkflowBaseDataType,inline"`
	Properties           *[]WorkflowTargetProperty `json:"Properties,omitempty" yaml:"Properties,omitempty"`
}

// NewWorkflowTargetDataType instantiates a new WorkflowTargetDataType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTargetDataType() *WorkflowTargetDataType {
	this := WorkflowTargetDataType{}
	return &this
}

// NewWorkflowTargetDataTypeWithDefaults instantiates a new WorkflowTargetDataType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTargetDataTypeWithDefaults() *WorkflowTargetDataType {
	this := WorkflowTargetDataType{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *WorkflowTargetDataType) GetProperties() []WorkflowTargetProperty {
	if o == nil || o.Properties == nil {
		var ret []WorkflowTargetProperty
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetDataType) GetPropertiesOk() (*[]WorkflowTargetProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowTargetDataType) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []WorkflowTargetProperty and assigns it to the Properties field.
func (o *WorkflowTargetDataType) SetProperties(v []WorkflowTargetProperty) {
	o.Properties = &v
}

func (o WorkflowTargetDataType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowBaseDataType, errWorkflowBaseDataType := json.Marshal(o.WorkflowBaseDataType)
	if errWorkflowBaseDataType != nil {
		return []byte{}, errWorkflowBaseDataType
	}
	errWorkflowBaseDataType = json.Unmarshal([]byte(serializedWorkflowBaseDataType), &toSerialize)
	if errWorkflowBaseDataType != nil {
		return []byte{}, errWorkflowBaseDataType
	}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTargetDataType struct {
	value *WorkflowTargetDataType
	isSet bool
}

func (v NullableWorkflowTargetDataType) Get() *WorkflowTargetDataType {
	return v.value
}

func (v *NullableWorkflowTargetDataType) Set(val *WorkflowTargetDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTargetDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTargetDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTargetDataType(val *WorkflowTargetDataType) *NullableWorkflowTargetDataType {
	return &NullableWorkflowTargetDataType{value: val, isSet: true}
}

func (v NullableWorkflowTargetDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTargetDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
