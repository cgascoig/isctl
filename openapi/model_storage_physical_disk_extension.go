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

// StoragePhysicalDiskExtension Information of disks as reported by controller. In certain cases like S-series servers, disk information will be reported by controller separately and this represents such information.
type StoragePhysicalDiskExtension struct {
	EquipmentBase `yaml:"EquipmentBase,inline"`
	// It shows whether disk is bootable or not.
	Bootable *string `json:"Bootable,omitempty" yaml:"Bootable,omitempty"`
	// It shows the Physical drive Dn.
	DiskDn *string `json:"DiskDn,omitempty" yaml:"DiskDn,omitempty"`
	// It shows storage Enclosure slotId.
	DiskId *int64 `json:"DiskId,omitempty" yaml:"DiskId,omitempty"`
	// It shows the current drive state of disk.
	DiskState *string `json:"DiskState,omitempty" yaml:"DiskState,omitempty"`
	// It shows the current drive state of disk.
	Health            *string                              `json:"Health,omitempty" yaml:"Health,omitempty"`
	PhysicalDisk      *StoragePhysicalDiskRelationship     `json:"PhysicalDisk,omitempty" yaml:"PhysicalDisk,omitempty"`
	RegisteredDevice  *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
	StorageController *StorageControllerRelationship       `json:"StorageController,omitempty" yaml:"StorageController,omitempty"`
}

// NewStoragePhysicalDiskExtension instantiates a new StoragePhysicalDiskExtension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePhysicalDiskExtension() *StoragePhysicalDiskExtension {
	this := StoragePhysicalDiskExtension{}
	return &this
}

// NewStoragePhysicalDiskExtensionWithDefaults instantiates a new StoragePhysicalDiskExtension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePhysicalDiskExtensionWithDefaults() *StoragePhysicalDiskExtension {
	this := StoragePhysicalDiskExtension{}
	return &this
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtension) GetBootable() string {
	if o == nil || o.Bootable == nil {
		var ret string
		return ret
	}
	return *o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtension) GetBootableOk() (*string, bool) {
	if o == nil || o.Bootable == nil {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtension) HasBootable() bool {
	if o != nil && o.Bootable != nil {
		return true
	}

	return false
}

// SetBootable gets a reference to the given string and assigns it to the Bootable field.
func (o *StoragePhysicalDiskExtension) SetBootable(v string) {
	o.Bootable = &v
}

// GetDiskDn returns the DiskDn field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtension) GetDiskDn() string {
	if o == nil || o.DiskDn == nil {
		var ret string
		return ret
	}
	return *o.DiskDn
}

// GetDiskDnOk returns a tuple with the DiskDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtension) GetDiskDnOk() (*string, bool) {
	if o == nil || o.DiskDn == nil {
		return nil, false
	}
	return o.DiskDn, true
}

// HasDiskDn returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtension) HasDiskDn() bool {
	if o != nil && o.DiskDn != nil {
		return true
	}

	return false
}

// SetDiskDn gets a reference to the given string and assigns it to the DiskDn field.
func (o *StoragePhysicalDiskExtension) SetDiskDn(v string) {
	o.DiskDn = &v
}

// GetDiskId returns the DiskId field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtension) GetDiskId() int64 {
	if o == nil || o.DiskId == nil {
		var ret int64
		return ret
	}
	return *o.DiskId
}

// GetDiskIdOk returns a tuple with the DiskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtension) GetDiskIdOk() (*int64, bool) {
	if o == nil || o.DiskId == nil {
		return nil, false
	}
	return o.DiskId, true
}

// HasDiskId returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtension) HasDiskId() bool {
	if o != nil && o.DiskId != nil {
		return true
	}

	return false
}

// SetDiskId gets a reference to the given int64 and assigns it to the DiskId field.
func (o *StoragePhysicalDiskExtension) SetDiskId(v int64) {
	o.DiskId = &v
}

// GetDiskState returns the DiskState field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtension) GetDiskState() string {
	if o == nil || o.DiskState == nil {
		var ret string
		return ret
	}
	return *o.DiskState
}

// GetDiskStateOk returns a tuple with the DiskState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtension) GetDiskStateOk() (*string, bool) {
	if o == nil || o.DiskState == nil {
		return nil, false
	}
	return o.DiskState, true
}

// HasDiskState returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtension) HasDiskState() bool {
	if o != nil && o.DiskState != nil {
		return true
	}

	return false
}

// SetDiskState gets a reference to the given string and assigns it to the DiskState field.
func (o *StoragePhysicalDiskExtension) SetDiskState(v string) {
	o.DiskState = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtension) GetHealth() string {
	if o == nil || o.Health == nil {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtension) GetHealthOk() (*string, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtension) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *StoragePhysicalDiskExtension) SetHealth(v string) {
	o.Health = &v
}

// GetPhysicalDisk returns the PhysicalDisk field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtension) GetPhysicalDisk() StoragePhysicalDiskRelationship {
	if o == nil || o.PhysicalDisk == nil {
		var ret StoragePhysicalDiskRelationship
		return ret
	}
	return *o.PhysicalDisk
}

// GetPhysicalDiskOk returns a tuple with the PhysicalDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtension) GetPhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool) {
	if o == nil || o.PhysicalDisk == nil {
		return nil, false
	}
	return o.PhysicalDisk, true
}

// HasPhysicalDisk returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtension) HasPhysicalDisk() bool {
	if o != nil && o.PhysicalDisk != nil {
		return true
	}

	return false
}

// SetPhysicalDisk gets a reference to the given StoragePhysicalDiskRelationship and assigns it to the PhysicalDisk field.
func (o *StoragePhysicalDiskExtension) SetPhysicalDisk(v StoragePhysicalDiskRelationship) {
	o.PhysicalDisk = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtension) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtension) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtension) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePhysicalDiskExtension) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageController returns the StorageController field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtension) GetStorageController() StorageControllerRelationship {
	if o == nil || o.StorageController == nil {
		var ret StorageControllerRelationship
		return ret
	}
	return *o.StorageController
}

// GetStorageControllerOk returns a tuple with the StorageController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtension) GetStorageControllerOk() (*StorageControllerRelationship, bool) {
	if o == nil || o.StorageController == nil {
		return nil, false
	}
	return o.StorageController, true
}

// HasStorageController returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtension) HasStorageController() bool {
	if o != nil && o.StorageController != nil {
		return true
	}

	return false
}

// SetStorageController gets a reference to the given StorageControllerRelationship and assigns it to the StorageController field.
func (o *StoragePhysicalDiskExtension) SetStorageController(v StorageControllerRelationship) {
	o.StorageController = &v
}

func (o StoragePhysicalDiskExtension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if o.Bootable != nil {
		toSerialize["Bootable"] = o.Bootable
	}
	if o.DiskDn != nil {
		toSerialize["DiskDn"] = o.DiskDn
	}
	if o.DiskId != nil {
		toSerialize["DiskId"] = o.DiskId
	}
	if o.DiskState != nil {
		toSerialize["DiskState"] = o.DiskState
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.PhysicalDisk != nil {
		toSerialize["PhysicalDisk"] = o.PhysicalDisk
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageController != nil {
		toSerialize["StorageController"] = o.StorageController
	}
	return json.Marshal(toSerialize)
}

type NullableStoragePhysicalDiskExtension struct {
	value *StoragePhysicalDiskExtension
	isSet bool
}

func (v NullableStoragePhysicalDiskExtension) Get() *StoragePhysicalDiskExtension {
	return v.value
}

func (v *NullableStoragePhysicalDiskExtension) Set(val *StoragePhysicalDiskExtension) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePhysicalDiskExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePhysicalDiskExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePhysicalDiskExtension(val *StoragePhysicalDiskExtension) *NullableStoragePhysicalDiskExtension {
	return &NullableStoragePhysicalDiskExtension{value: val, isSet: true}
}

func (v NullableStoragePhysicalDiskExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePhysicalDiskExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
