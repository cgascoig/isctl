/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-11T18:30:19Z.
 *
 * API version: 1.0.9-3252
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// HyperflexNamedVsanAllOf Definition of the list of properties defined in 'hyperflex.NamedVsan', excluding properties defined in parent classes.
type HyperflexNamedVsanAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The name of the VSAN. The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// The ID of the named VSAN. The ID can be any number between 1 and 4093, inclusive.
	VsanId *int64 `json:"VsanId,omitempty" yaml:"VsanId,omitempty"`
}

// NewHyperflexNamedVsanAllOf instantiates a new HyperflexNamedVsanAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexNamedVsanAllOf(classId string, objectType string) *HyperflexNamedVsanAllOf {
	this := HyperflexNamedVsanAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexNamedVsanAllOfWithDefaults instantiates a new HyperflexNamedVsanAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexNamedVsanAllOfWithDefaults() *HyperflexNamedVsanAllOf {
	this := HyperflexNamedVsanAllOf{}
	var classId string = "hyperflex.NamedVsan"
	this.ClassId = classId
	var objectType string = "hyperflex.NamedVsan"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexNamedVsanAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexNamedVsanAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexNamedVsanAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexNamedVsanAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexNamedVsanAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexNamedVsanAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexNamedVsanAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNamedVsanAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexNamedVsanAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexNamedVsanAllOf) SetName(v string) {
	o.Name = &v
}

// GetVsanId returns the VsanId field value if set, zero value otherwise.
func (o *HyperflexNamedVsanAllOf) GetVsanId() int64 {
	if o == nil || o.VsanId == nil {
		var ret int64
		return ret
	}
	return *o.VsanId
}

// GetVsanIdOk returns a tuple with the VsanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNamedVsanAllOf) GetVsanIdOk() (*int64, bool) {
	if o == nil || o.VsanId == nil {
		return nil, false
	}
	return o.VsanId, true
}

// HasVsanId returns a boolean if a field has been set.
func (o *HyperflexNamedVsanAllOf) HasVsanId() bool {
	if o != nil && o.VsanId != nil {
		return true
	}

	return false
}

// SetVsanId gets a reference to the given int64 and assigns it to the VsanId field.
func (o *HyperflexNamedVsanAllOf) SetVsanId(v int64) {
	o.VsanId = &v
}

func (o HyperflexNamedVsanAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.VsanId != nil {
		toSerialize["VsanId"] = o.VsanId
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexNamedVsanAllOf struct {
	value *HyperflexNamedVsanAllOf
	isSet bool
}

func (v NullableHyperflexNamedVsanAllOf) Get() *HyperflexNamedVsanAllOf {
	return v.value
}

func (v *NullableHyperflexNamedVsanAllOf) Set(val *HyperflexNamedVsanAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexNamedVsanAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexNamedVsanAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexNamedVsanAllOf(val *HyperflexNamedVsanAllOf) *NullableHyperflexNamedVsanAllOf {
	return &NullableHyperflexNamedVsanAllOf{value: val, isSet: true}
}

func (v NullableHyperflexNamedVsanAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexNamedVsanAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
