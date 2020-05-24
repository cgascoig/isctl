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

// HyperflexFeatureLimitEntryAllOf Definition of the list of properties defined in 'hyperflex.FeatureLimitEntry', excluding properties defined in parent classes.
type HyperflexFeatureLimitEntryAllOf struct {
	Constraint *HyperflexAppSettingConstraint `json:"Constraint,omitempty" yaml:"Constraint,omitempty"`
}

// NewHyperflexFeatureLimitEntryAllOf instantiates a new HyperflexFeatureLimitEntryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexFeatureLimitEntryAllOf() *HyperflexFeatureLimitEntryAllOf {
	this := HyperflexFeatureLimitEntryAllOf{}
	return &this
}

// NewHyperflexFeatureLimitEntryAllOfWithDefaults instantiates a new HyperflexFeatureLimitEntryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexFeatureLimitEntryAllOfWithDefaults() *HyperflexFeatureLimitEntryAllOf {
	this := HyperflexFeatureLimitEntryAllOf{}
	return &this
}

// GetConstraint returns the Constraint field value if set, zero value otherwise.
func (o *HyperflexFeatureLimitEntryAllOf) GetConstraint() HyperflexAppSettingConstraint {
	if o == nil || o.Constraint == nil {
		var ret HyperflexAppSettingConstraint
		return ret
	}
	return *o.Constraint
}

// GetConstraintOk returns a tuple with the Constraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexFeatureLimitEntryAllOf) GetConstraintOk() (*HyperflexAppSettingConstraint, bool) {
	if o == nil || o.Constraint == nil {
		return nil, false
	}
	return o.Constraint, true
}

// HasConstraint returns a boolean if a field has been set.
func (o *HyperflexFeatureLimitEntryAllOf) HasConstraint() bool {
	if o != nil && o.Constraint != nil {
		return true
	}

	return false
}

// SetConstraint gets a reference to the given HyperflexAppSettingConstraint and assigns it to the Constraint field.
func (o *HyperflexFeatureLimitEntryAllOf) SetConstraint(v HyperflexAppSettingConstraint) {
	o.Constraint = &v
}

func (o HyperflexFeatureLimitEntryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Constraint != nil {
		toSerialize["Constraint"] = o.Constraint
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexFeatureLimitEntryAllOf struct {
	value *HyperflexFeatureLimitEntryAllOf
	isSet bool
}

func (v NullableHyperflexFeatureLimitEntryAllOf) Get() *HyperflexFeatureLimitEntryAllOf {
	return v.value
}

func (v *NullableHyperflexFeatureLimitEntryAllOf) Set(val *HyperflexFeatureLimitEntryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexFeatureLimitEntryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexFeatureLimitEntryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexFeatureLimitEntryAllOf(val *HyperflexFeatureLimitEntryAllOf) *NullableHyperflexFeatureLimitEntryAllOf {
	return &NullableHyperflexFeatureLimitEntryAllOf{value: val, isSet: true}
}

func (v NullableHyperflexFeatureLimitEntryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexFeatureLimitEntryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
