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

// StoragePhysicalDiskUsage Has usage map between physical disks and virtual drives.
type StoragePhysicalDiskUsage struct {
	InventoryBase    `yaml:"InventoryBase,inline"`
	NumberOfBlocks   *string                              `json:"NumberOfBlocks,omitempty" yaml:"NumberOfBlocks,omitempty"`
	PhysicalDrive    *string                              `json:"PhysicalDrive,omitempty" yaml:"PhysicalDrive,omitempty"`
	Span             *string                              `json:"Span,omitempty" yaml:"Span,omitempty"`
	StartingBlock    *string                              `json:"StartingBlock,omitempty" yaml:"StartingBlock,omitempty"`
	State            *string                              `json:"State,omitempty" yaml:"State,omitempty"`
	VirtualDrive     *string                              `json:"VirtualDrive,omitempty" yaml:"VirtualDrive,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewStoragePhysicalDiskUsage instantiates a new StoragePhysicalDiskUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePhysicalDiskUsage() *StoragePhysicalDiskUsage {
	this := StoragePhysicalDiskUsage{}
	return &this
}

// NewStoragePhysicalDiskUsageWithDefaults instantiates a new StoragePhysicalDiskUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePhysicalDiskUsageWithDefaults() *StoragePhysicalDiskUsage {
	this := StoragePhysicalDiskUsage{}
	return &this
}

// GetNumberOfBlocks returns the NumberOfBlocks field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetNumberOfBlocks() string {
	if o == nil || o.NumberOfBlocks == nil {
		var ret string
		return ret
	}
	return *o.NumberOfBlocks
}

// GetNumberOfBlocksOk returns a tuple with the NumberOfBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetNumberOfBlocksOk() (*string, bool) {
	if o == nil || o.NumberOfBlocks == nil {
		return nil, false
	}
	return o.NumberOfBlocks, true
}

// HasNumberOfBlocks returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasNumberOfBlocks() bool {
	if o != nil && o.NumberOfBlocks != nil {
		return true
	}

	return false
}

// SetNumberOfBlocks gets a reference to the given string and assigns it to the NumberOfBlocks field.
func (o *StoragePhysicalDiskUsage) SetNumberOfBlocks(v string) {
	o.NumberOfBlocks = &v
}

// GetPhysicalDrive returns the PhysicalDrive field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetPhysicalDrive() string {
	if o == nil || o.PhysicalDrive == nil {
		var ret string
		return ret
	}
	return *o.PhysicalDrive
}

// GetPhysicalDriveOk returns a tuple with the PhysicalDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetPhysicalDriveOk() (*string, bool) {
	if o == nil || o.PhysicalDrive == nil {
		return nil, false
	}
	return o.PhysicalDrive, true
}

// HasPhysicalDrive returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasPhysicalDrive() bool {
	if o != nil && o.PhysicalDrive != nil {
		return true
	}

	return false
}

// SetPhysicalDrive gets a reference to the given string and assigns it to the PhysicalDrive field.
func (o *StoragePhysicalDiskUsage) SetPhysicalDrive(v string) {
	o.PhysicalDrive = &v
}

// GetSpan returns the Span field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetSpan() string {
	if o == nil || o.Span == nil {
		var ret string
		return ret
	}
	return *o.Span
}

// GetSpanOk returns a tuple with the Span field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetSpanOk() (*string, bool) {
	if o == nil || o.Span == nil {
		return nil, false
	}
	return o.Span, true
}

// HasSpan returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasSpan() bool {
	if o != nil && o.Span != nil {
		return true
	}

	return false
}

// SetSpan gets a reference to the given string and assigns it to the Span field.
func (o *StoragePhysicalDiskUsage) SetSpan(v string) {
	o.Span = &v
}

// GetStartingBlock returns the StartingBlock field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetStartingBlock() string {
	if o == nil || o.StartingBlock == nil {
		var ret string
		return ret
	}
	return *o.StartingBlock
}

// GetStartingBlockOk returns a tuple with the StartingBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetStartingBlockOk() (*string, bool) {
	if o == nil || o.StartingBlock == nil {
		return nil, false
	}
	return o.StartingBlock, true
}

// HasStartingBlock returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasStartingBlock() bool {
	if o != nil && o.StartingBlock != nil {
		return true
	}

	return false
}

// SetStartingBlock gets a reference to the given string and assigns it to the StartingBlock field.
func (o *StoragePhysicalDiskUsage) SetStartingBlock(v string) {
	o.StartingBlock = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StoragePhysicalDiskUsage) SetState(v string) {
	o.State = &v
}

// GetVirtualDrive returns the VirtualDrive field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetVirtualDrive() string {
	if o == nil || o.VirtualDrive == nil {
		var ret string
		return ret
	}
	return *o.VirtualDrive
}

// GetVirtualDriveOk returns a tuple with the VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetVirtualDriveOk() (*string, bool) {
	if o == nil || o.VirtualDrive == nil {
		return nil, false
	}
	return o.VirtualDrive, true
}

// HasVirtualDrive returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasVirtualDrive() bool {
	if o != nil && o.VirtualDrive != nil {
		return true
	}

	return false
}

// SetVirtualDrive gets a reference to the given string and assigns it to the VirtualDrive field.
func (o *StoragePhysicalDiskUsage) SetVirtualDrive(v string) {
	o.VirtualDrive = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePhysicalDiskUsage) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePhysicalDiskUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if o.NumberOfBlocks != nil {
		toSerialize["NumberOfBlocks"] = o.NumberOfBlocks
	}
	if o.PhysicalDrive != nil {
		toSerialize["PhysicalDrive"] = o.PhysicalDrive
	}
	if o.Span != nil {
		toSerialize["Span"] = o.Span
	}
	if o.StartingBlock != nil {
		toSerialize["StartingBlock"] = o.StartingBlock
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.VirtualDrive != nil {
		toSerialize["VirtualDrive"] = o.VirtualDrive
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableStoragePhysicalDiskUsage struct {
	value *StoragePhysicalDiskUsage
	isSet bool
}

func (v NullableStoragePhysicalDiskUsage) Get() *StoragePhysicalDiskUsage {
	return v.value
}

func (v *NullableStoragePhysicalDiskUsage) Set(val *StoragePhysicalDiskUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePhysicalDiskUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePhysicalDiskUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePhysicalDiskUsage(val *StoragePhysicalDiskUsage) *NullableStoragePhysicalDiskUsage {
	return &NullableStoragePhysicalDiskUsage{value: val, isSet: true}
}

func (v NullableStoragePhysicalDiskUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePhysicalDiskUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
