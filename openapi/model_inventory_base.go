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

// InventoryBase Base class for all inventory MOs.
type InventoryBase struct {
	MoBaseMo   `yaml:"MoBaseMo,inline"`
	DeviceMoId *string `json:"DeviceMoId,omitempty" yaml:"DeviceMoId,omitempty"`
	// The Distinguished Name unambiguously identifies an object in the system.
	Dn *string `json:"Dn,omitempty" yaml:"Dn,omitempty"`
	// The Relative Name uniquely identifies an object within a given context.
	Rn *string `json:"Rn,omitempty" yaml:"Rn,omitempty"`
}

// NewInventoryBase instantiates a new InventoryBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryBase() *InventoryBase {
	this := InventoryBase{}
	return &this
}

// NewInventoryBaseWithDefaults instantiates a new InventoryBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryBaseWithDefaults() *InventoryBase {
	this := InventoryBase{}
	return &this
}

// GetDeviceMoId returns the DeviceMoId field value if set, zero value otherwise.
func (o *InventoryBase) GetDeviceMoId() string {
	if o == nil || o.DeviceMoId == nil {
		var ret string
		return ret
	}
	return *o.DeviceMoId
}

// GetDeviceMoIdOk returns a tuple with the DeviceMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryBase) GetDeviceMoIdOk() (*string, bool) {
	if o == nil || o.DeviceMoId == nil {
		return nil, false
	}
	return o.DeviceMoId, true
}

// HasDeviceMoId returns a boolean if a field has been set.
func (o *InventoryBase) HasDeviceMoId() bool {
	if o != nil && o.DeviceMoId != nil {
		return true
	}

	return false
}

// SetDeviceMoId gets a reference to the given string and assigns it to the DeviceMoId field.
func (o *InventoryBase) SetDeviceMoId(v string) {
	o.DeviceMoId = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *InventoryBase) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryBase) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *InventoryBase) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *InventoryBase) SetDn(v string) {
	o.Dn = &v
}

// GetRn returns the Rn field value if set, zero value otherwise.
func (o *InventoryBase) GetRn() string {
	if o == nil || o.Rn == nil {
		var ret string
		return ret
	}
	return *o.Rn
}

// GetRnOk returns a tuple with the Rn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryBase) GetRnOk() (*string, bool) {
	if o == nil || o.Rn == nil {
		return nil, false
	}
	return o.Rn, true
}

// HasRn returns a boolean if a field has been set.
func (o *InventoryBase) HasRn() bool {
	if o != nil && o.Rn != nil {
		return true
	}

	return false
}

// SetRn gets a reference to the given string and assigns it to the Rn field.
func (o *InventoryBase) SetRn(v string) {
	o.Rn = &v
}

func (o InventoryBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.DeviceMoId != nil {
		toSerialize["DeviceMoId"] = o.DeviceMoId
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.Rn != nil {
		toSerialize["Rn"] = o.Rn
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryBase struct {
	value *InventoryBase
	isSet bool
}

func (v NullableInventoryBase) Get() *InventoryBase {
	return v.value
}

func (v *NullableInventoryBase) Set(val *InventoryBase) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryBase) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryBase(val *InventoryBase) *NullableInventoryBase {
	return &NullableInventoryBase{value: val, isSet: true}
}

func (v NullableInventoryBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
