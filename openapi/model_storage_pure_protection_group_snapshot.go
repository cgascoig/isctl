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

// StoragePureProtectionGroupSnapshot Protection group snapshot entity in Pure protection group.
type StoragePureProtectionGroupSnapshot struct {
	StorageProtectionGroupSnapshot `yaml:"StorageProtectionGroupSnapshot,inline"`
	RegisteredDevice               *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewStoragePureProtectionGroupSnapshot instantiates a new StoragePureProtectionGroupSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureProtectionGroupSnapshot() *StoragePureProtectionGroupSnapshot {
	this := StoragePureProtectionGroupSnapshot{}
	return &this
}

// NewStoragePureProtectionGroupSnapshotWithDefaults instantiates a new StoragePureProtectionGroupSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureProtectionGroupSnapshotWithDefaults() *StoragePureProtectionGroupSnapshot {
	this := StoragePureProtectionGroupSnapshot{}
	return &this
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureProtectionGroupSnapshot) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureProtectionGroupSnapshot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureProtectionGroupSnapshot) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureProtectionGroupSnapshot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePureProtectionGroupSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageProtectionGroupSnapshot, errStorageProtectionGroupSnapshot := json.Marshal(o.StorageProtectionGroupSnapshot)
	if errStorageProtectionGroupSnapshot != nil {
		return []byte{}, errStorageProtectionGroupSnapshot
	}
	errStorageProtectionGroupSnapshot = json.Unmarshal([]byte(serializedStorageProtectionGroupSnapshot), &toSerialize)
	if errStorageProtectionGroupSnapshot != nil {
		return []byte{}, errStorageProtectionGroupSnapshot
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableStoragePureProtectionGroupSnapshot struct {
	value *StoragePureProtectionGroupSnapshot
	isSet bool
}

func (v NullableStoragePureProtectionGroupSnapshot) Get() *StoragePureProtectionGroupSnapshot {
	return v.value
}

func (v *NullableStoragePureProtectionGroupSnapshot) Set(val *StoragePureProtectionGroupSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureProtectionGroupSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureProtectionGroupSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureProtectionGroupSnapshot(val *StoragePureProtectionGroupSnapshot) *NullableStoragePureProtectionGroupSnapshot {
	return &NullableStoragePureProtectionGroupSnapshot{value: val, isSet: true}
}

func (v NullableStoragePureProtectionGroupSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureProtectionGroupSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
