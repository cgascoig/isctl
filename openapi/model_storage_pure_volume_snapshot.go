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

// StoragePureVolumeSnapshot Volume snapshot entity in Pure protection group. Volume snapshots are created either on-demand or using scheduler. Snapshots are immutable and it cannot be connected to hosts or host groups, and therefore the data it contains cannot be read or written.
type StoragePureVolumeSnapshot struct {
	StorageSnapshot `yaml:"StorageSnapshot,inline"`
	// Unique serial number of the snapshot allocated by the storage array.
	Serial           *string                              `json:"Serial,omitempty" yaml:"Serial,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewStoragePureVolumeSnapshot instantiates a new StoragePureVolumeSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureVolumeSnapshot() *StoragePureVolumeSnapshot {
	this := StoragePureVolumeSnapshot{}
	return &this
}

// NewStoragePureVolumeSnapshotWithDefaults instantiates a new StoragePureVolumeSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureVolumeSnapshotWithDefaults() *StoragePureVolumeSnapshot {
	this := StoragePureVolumeSnapshot{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *StoragePureVolumeSnapshot) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureVolumeSnapshot) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *StoragePureVolumeSnapshot) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *StoragePureVolumeSnapshot) SetSerial(v string) {
	o.Serial = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureVolumeSnapshot) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureVolumeSnapshot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureVolumeSnapshot) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureVolumeSnapshot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePureVolumeSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageSnapshot, errStorageSnapshot := json.Marshal(o.StorageSnapshot)
	if errStorageSnapshot != nil {
		return []byte{}, errStorageSnapshot
	}
	errStorageSnapshot = json.Unmarshal([]byte(serializedStorageSnapshot), &toSerialize)
	if errStorageSnapshot != nil {
		return []byte{}, errStorageSnapshot
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableStoragePureVolumeSnapshot struct {
	value *StoragePureVolumeSnapshot
	isSet bool
}

func (v NullableStoragePureVolumeSnapshot) Get() *StoragePureVolumeSnapshot {
	return v.value
}

func (v *NullableStoragePureVolumeSnapshot) Set(val *StoragePureVolumeSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureVolumeSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureVolumeSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureVolumeSnapshot(val *StoragePureVolumeSnapshot) *NullableStoragePureVolumeSnapshot {
	return &NullableStoragePureVolumeSnapshot{value: val, isSet: true}
}

func (v NullableStoragePureVolumeSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureVolumeSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
