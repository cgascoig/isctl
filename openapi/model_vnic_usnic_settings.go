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

// VnicUsnicSettings User Space NIC Settings that enable low-latency and higher throughput by bypassing the kernel layer when sending/receiving packets.
type VnicUsnicSettings struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// Class of Service to be used for traffic on the usNIC.
	Cos *int64 `json:"Cos,omitempty" yaml:"Cos,omitempty"`
	// Number of usNIC interfaces to be created.
	Count *int64 `json:"Count,omitempty" yaml:"Count,omitempty"`
	// Ethernet Adapter policy to be associated with the usNICs.
	UsnicAdapterPolicy *string `json:"UsnicAdapterPolicy,omitempty" yaml:"UsnicAdapterPolicy,omitempty"`
}

// NewVnicUsnicSettings instantiates a new VnicUsnicSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicUsnicSettings() *VnicUsnicSettings {
	this := VnicUsnicSettings{}
	return &this
}

// NewVnicUsnicSettingsWithDefaults instantiates a new VnicUsnicSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicUsnicSettingsWithDefaults() *VnicUsnicSettings {
	this := VnicUsnicSettings{}
	return &this
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *VnicUsnicSettings) GetCos() int64 {
	if o == nil || o.Cos == nil {
		var ret int64
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicUsnicSettings) GetCosOk() (*int64, bool) {
	if o == nil || o.Cos == nil {
		return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *VnicUsnicSettings) HasCos() bool {
	if o != nil && o.Cos != nil {
		return true
	}

	return false
}

// SetCos gets a reference to the given int64 and assigns it to the Cos field.
func (o *VnicUsnicSettings) SetCos(v int64) {
	o.Cos = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *VnicUsnicSettings) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicUsnicSettings) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *VnicUsnicSettings) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *VnicUsnicSettings) SetCount(v int64) {
	o.Count = &v
}

// GetUsnicAdapterPolicy returns the UsnicAdapterPolicy field value if set, zero value otherwise.
func (o *VnicUsnicSettings) GetUsnicAdapterPolicy() string {
	if o == nil || o.UsnicAdapterPolicy == nil {
		var ret string
		return ret
	}
	return *o.UsnicAdapterPolicy
}

// GetUsnicAdapterPolicyOk returns a tuple with the UsnicAdapterPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicUsnicSettings) GetUsnicAdapterPolicyOk() (*string, bool) {
	if o == nil || o.UsnicAdapterPolicy == nil {
		return nil, false
	}
	return o.UsnicAdapterPolicy, true
}

// HasUsnicAdapterPolicy returns a boolean if a field has been set.
func (o *VnicUsnicSettings) HasUsnicAdapterPolicy() bool {
	if o != nil && o.UsnicAdapterPolicy != nil {
		return true
	}

	return false
}

// SetUsnicAdapterPolicy gets a reference to the given string and assigns it to the UsnicAdapterPolicy field.
func (o *VnicUsnicSettings) SetUsnicAdapterPolicy(v string) {
	o.UsnicAdapterPolicy = &v
}

func (o VnicUsnicSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Cos != nil {
		toSerialize["Cos"] = o.Cos
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.UsnicAdapterPolicy != nil {
		toSerialize["UsnicAdapterPolicy"] = o.UsnicAdapterPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableVnicUsnicSettings struct {
	value *VnicUsnicSettings
	isSet bool
}

func (v NullableVnicUsnicSettings) Get() *VnicUsnicSettings {
	return v.value
}

func (v *NullableVnicUsnicSettings) Set(val *VnicUsnicSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicUsnicSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicUsnicSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicUsnicSettings(val *VnicUsnicSettings) *NullableVnicUsnicSettings {
	return &NullableVnicUsnicSettings{value: val, isSet: true}
}

func (v NullableVnicUsnicSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicUsnicSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
