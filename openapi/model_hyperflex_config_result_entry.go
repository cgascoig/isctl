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

// HyperflexConfigResultEntry An entry that describes the result of a cluster validation or deployment operation.
type HyperflexConfigResultEntry struct {
	PolicyAbstractConfigResultEntry `yaml:"PolicyAbstractConfigResultEntry,inline"`
	ConfigResult                    *HyperflexConfigResultRelationship `json:"ConfigResult,omitempty" yaml:"ConfigResult,omitempty"`
}

// NewHyperflexConfigResultEntry instantiates a new HyperflexConfigResultEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexConfigResultEntry() *HyperflexConfigResultEntry {
	this := HyperflexConfigResultEntry{}
	return &this
}

// NewHyperflexConfigResultEntryWithDefaults instantiates a new HyperflexConfigResultEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexConfigResultEntryWithDefaults() *HyperflexConfigResultEntry {
	this := HyperflexConfigResultEntry{}
	return &this
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *HyperflexConfigResultEntry) GetConfigResult() HyperflexConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret HyperflexConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexConfigResultEntry) GetConfigResultOk() (*HyperflexConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *HyperflexConfigResultEntry) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given HyperflexConfigResultRelationship and assigns it to the ConfigResult field.
func (o *HyperflexConfigResultEntry) SetConfigResult(v HyperflexConfigResultRelationship) {
	o.ConfigResult = &v
}

func (o HyperflexConfigResultEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigResultEntry, errPolicyAbstractConfigResultEntry := json.Marshal(o.PolicyAbstractConfigResultEntry)
	if errPolicyAbstractConfigResultEntry != nil {
		return []byte{}, errPolicyAbstractConfigResultEntry
	}
	errPolicyAbstractConfigResultEntry = json.Unmarshal([]byte(serializedPolicyAbstractConfigResultEntry), &toSerialize)
	if errPolicyAbstractConfigResultEntry != nil {
		return []byte{}, errPolicyAbstractConfigResultEntry
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexConfigResultEntry struct {
	value *HyperflexConfigResultEntry
	isSet bool
}

func (v NullableHyperflexConfigResultEntry) Get() *HyperflexConfigResultEntry {
	return v.value
}

func (v *NullableHyperflexConfigResultEntry) Set(val *HyperflexConfigResultEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexConfigResultEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexConfigResultEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexConfigResultEntry(val *HyperflexConfigResultEntry) *NullableHyperflexConfigResultEntry {
	return &NullableHyperflexConfigResultEntry{value: val, isSet: true}
}

func (v NullableHyperflexConfigResultEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexConfigResultEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
