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

// WorkflowWorkflowProperties Properties for a workflow definition.
type WorkflowWorkflowProperties struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// When set to false the workflow is owned by the system and used for internal services. Such workflows cannot be directly used by external entities.
	ExternalMeta *bool `json:"ExternalMeta,omitempty" yaml:"ExternalMeta,omitempty"`
	// When true, this workflow can be retried if has not been modified for more than a period of 2 weeks.
	Retryable *bool `json:"Retryable,omitempty" yaml:"Retryable,omitempty"`
	// Supported status of the definition.
	SupportStatus *string `json:"SupportStatus,omitempty" yaml:"SupportStatus,omitempty"`
}

// NewWorkflowWorkflowProperties instantiates a new WorkflowWorkflowProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkflowProperties() *WorkflowWorkflowProperties {
	this := WorkflowWorkflowProperties{}
	var supportStatus string = "Supported"
	this.SupportStatus = &supportStatus
	return &this
}

// NewWorkflowWorkflowPropertiesWithDefaults instantiates a new WorkflowWorkflowProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkflowPropertiesWithDefaults() *WorkflowWorkflowProperties {
	this := WorkflowWorkflowProperties{}
	var supportStatus string = "Supported"
	this.SupportStatus = &supportStatus
	return &this
}

// GetExternalMeta returns the ExternalMeta field value if set, zero value otherwise.
func (o *WorkflowWorkflowProperties) GetExternalMeta() bool {
	if o == nil || o.ExternalMeta == nil {
		var ret bool
		return ret
	}
	return *o.ExternalMeta
}

// GetExternalMetaOk returns a tuple with the ExternalMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowProperties) GetExternalMetaOk() (*bool, bool) {
	if o == nil || o.ExternalMeta == nil {
		return nil, false
	}
	return o.ExternalMeta, true
}

// HasExternalMeta returns a boolean if a field has been set.
func (o *WorkflowWorkflowProperties) HasExternalMeta() bool {
	if o != nil && o.ExternalMeta != nil {
		return true
	}

	return false
}

// SetExternalMeta gets a reference to the given bool and assigns it to the ExternalMeta field.
func (o *WorkflowWorkflowProperties) SetExternalMeta(v bool) {
	o.ExternalMeta = &v
}

// GetRetryable returns the Retryable field value if set, zero value otherwise.
func (o *WorkflowWorkflowProperties) GetRetryable() bool {
	if o == nil || o.Retryable == nil {
		var ret bool
		return ret
	}
	return *o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowProperties) GetRetryableOk() (*bool, bool) {
	if o == nil || o.Retryable == nil {
		return nil, false
	}
	return o.Retryable, true
}

// HasRetryable returns a boolean if a field has been set.
func (o *WorkflowWorkflowProperties) HasRetryable() bool {
	if o != nil && o.Retryable != nil {
		return true
	}

	return false
}

// SetRetryable gets a reference to the given bool and assigns it to the Retryable field.
func (o *WorkflowWorkflowProperties) SetRetryable(v bool) {
	o.Retryable = &v
}

// GetSupportStatus returns the SupportStatus field value if set, zero value otherwise.
func (o *WorkflowWorkflowProperties) GetSupportStatus() string {
	if o == nil || o.SupportStatus == nil {
		var ret string
		return ret
	}
	return *o.SupportStatus
}

// GetSupportStatusOk returns a tuple with the SupportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowProperties) GetSupportStatusOk() (*string, bool) {
	if o == nil || o.SupportStatus == nil {
		return nil, false
	}
	return o.SupportStatus, true
}

// HasSupportStatus returns a boolean if a field has been set.
func (o *WorkflowWorkflowProperties) HasSupportStatus() bool {
	if o != nil && o.SupportStatus != nil {
		return true
	}

	return false
}

// SetSupportStatus gets a reference to the given string and assigns it to the SupportStatus field.
func (o *WorkflowWorkflowProperties) SetSupportStatus(v string) {
	o.SupportStatus = &v
}

func (o WorkflowWorkflowProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.ExternalMeta != nil {
		toSerialize["ExternalMeta"] = o.ExternalMeta
	}
	if o.Retryable != nil {
		toSerialize["Retryable"] = o.Retryable
	}
	if o.SupportStatus != nil {
		toSerialize["SupportStatus"] = o.SupportStatus
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowWorkflowProperties struct {
	value *WorkflowWorkflowProperties
	isSet bool
}

func (v NullableWorkflowWorkflowProperties) Get() *WorkflowWorkflowProperties {
	return v.value
}

func (v *NullableWorkflowWorkflowProperties) Set(val *WorkflowWorkflowProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkflowProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkflowProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkflowProperties(val *WorkflowWorkflowProperties) *NullableWorkflowWorkflowProperties {
	return &NullableWorkflowWorkflowProperties{value: val, isSet: true}
}

func (v NullableWorkflowWorkflowProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkflowProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
