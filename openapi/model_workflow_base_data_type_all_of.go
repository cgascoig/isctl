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

// WorkflowBaseDataTypeAllOf Definition of the list of properties defined in 'workflow.BaseDataType', excluding properties defined in parent classes.
type WorkflowBaseDataTypeAllOf struct {
	Default *WorkflowDefaultValue `json:"Default,omitempty" yaml:"Default,omitempty"`
	// Provide a detailed description of the data type.
	Description *string `json:"Description,omitempty" yaml:"Description,omitempty"`
	// Descriptive label for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character.
	Label *string `json:"Label,omitempty" yaml:"Label,omitempty"`
	// Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// Specifies whether this parameter is required. The field is applicable for task and workflow.
	Required *bool `json:"Required,omitempty" yaml:"Required,omitempty"`
}

// NewWorkflowBaseDataTypeAllOf instantiates a new WorkflowBaseDataTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowBaseDataTypeAllOf() *WorkflowBaseDataTypeAllOf {
	this := WorkflowBaseDataTypeAllOf{}
	return &this
}

// NewWorkflowBaseDataTypeAllOfWithDefaults instantiates a new WorkflowBaseDataTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowBaseDataTypeAllOfWithDefaults() *WorkflowBaseDataTypeAllOf {
	this := WorkflowBaseDataTypeAllOf{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *WorkflowBaseDataTypeAllOf) GetDefault() WorkflowDefaultValue {
	if o == nil || o.Default == nil {
		var ret WorkflowDefaultValue
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBaseDataTypeAllOf) GetDefaultOk() (*WorkflowDefaultValue, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *WorkflowBaseDataTypeAllOf) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given WorkflowDefaultValue and assigns it to the Default field.
func (o *WorkflowBaseDataTypeAllOf) SetDefault(v WorkflowDefaultValue) {
	o.Default = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowBaseDataTypeAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBaseDataTypeAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowBaseDataTypeAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowBaseDataTypeAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowBaseDataTypeAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBaseDataTypeAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowBaseDataTypeAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowBaseDataTypeAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowBaseDataTypeAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBaseDataTypeAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowBaseDataTypeAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowBaseDataTypeAllOf) SetName(v string) {
	o.Name = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *WorkflowBaseDataTypeAllOf) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBaseDataTypeAllOf) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *WorkflowBaseDataTypeAllOf) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *WorkflowBaseDataTypeAllOf) SetRequired(v bool) {
	o.Required = &v
}

func (o WorkflowBaseDataTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Default != nil {
		toSerialize["Default"] = o.Default
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Required != nil {
		toSerialize["Required"] = o.Required
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowBaseDataTypeAllOf struct {
	value *WorkflowBaseDataTypeAllOf
	isSet bool
}

func (v NullableWorkflowBaseDataTypeAllOf) Get() *WorkflowBaseDataTypeAllOf {
	return v.value
}

func (v *NullableWorkflowBaseDataTypeAllOf) Set(val *WorkflowBaseDataTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowBaseDataTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowBaseDataTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowBaseDataTypeAllOf(val *WorkflowBaseDataTypeAllOf) *NullableWorkflowBaseDataTypeAllOf {
	return &NullableWorkflowBaseDataTypeAllOf{value: val, isSet: true}
}

func (v NullableWorkflowBaseDataTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowBaseDataTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
