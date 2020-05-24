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

// StorageEnclosureDiskSlotEp Physical Disk slots on the enclosure.
type StorageEnclosureDiskSlotEp struct {
	InventoryBase `yaml:"InventoryBase,inline"`
	// This field identifies the zoning configuration applied to  this enclosure slot.
	DrivePath *string `json:"DrivePath,omitempty" yaml:"DrivePath,omitempty"`
	// This field identifies the health of the disk inserted in the slot.
	Health *string `json:"Health,omitempty" yaml:"Health,omitempty"`
	// This field identifies the disk is present in the enclosure slot.
	Presence *string `json:"Presence,omitempty" yaml:"Presence,omitempty"`
	// This field represents the slot Id in the storage enclosure.
	Slot             *string                              `json:"Slot,omitempty" yaml:"Slot,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
	StorageEnclosure *StorageEnclosureRelationship        `json:"StorageEnclosure,omitempty" yaml:"StorageEnclosure,omitempty"`
}

// NewStorageEnclosureDiskSlotEp instantiates a new StorageEnclosureDiskSlotEp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageEnclosureDiskSlotEp() *StorageEnclosureDiskSlotEp {
	this := StorageEnclosureDiskSlotEp{}
	return &this
}

// NewStorageEnclosureDiskSlotEpWithDefaults instantiates a new StorageEnclosureDiskSlotEp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageEnclosureDiskSlotEpWithDefaults() *StorageEnclosureDiskSlotEp {
	this := StorageEnclosureDiskSlotEp{}
	return &this
}

// GetDrivePath returns the DrivePath field value if set, zero value otherwise.
func (o *StorageEnclosureDiskSlotEp) GetDrivePath() string {
	if o == nil || o.DrivePath == nil {
		var ret string
		return ret
	}
	return *o.DrivePath
}

// GetDrivePathOk returns a tuple with the DrivePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEp) GetDrivePathOk() (*string, bool) {
	if o == nil || o.DrivePath == nil {
		return nil, false
	}
	return o.DrivePath, true
}

// HasDrivePath returns a boolean if a field has been set.
func (o *StorageEnclosureDiskSlotEp) HasDrivePath() bool {
	if o != nil && o.DrivePath != nil {
		return true
	}

	return false
}

// SetDrivePath gets a reference to the given string and assigns it to the DrivePath field.
func (o *StorageEnclosureDiskSlotEp) SetDrivePath(v string) {
	o.DrivePath = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *StorageEnclosureDiskSlotEp) GetHealth() string {
	if o == nil || o.Health == nil {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEp) GetHealthOk() (*string, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *StorageEnclosureDiskSlotEp) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *StorageEnclosureDiskSlotEp) SetHealth(v string) {
	o.Health = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *StorageEnclosureDiskSlotEp) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEp) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *StorageEnclosureDiskSlotEp) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *StorageEnclosureDiskSlotEp) SetPresence(v string) {
	o.Presence = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *StorageEnclosureDiskSlotEp) GetSlot() string {
	if o == nil || o.Slot == nil {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEp) GetSlotOk() (*string, bool) {
	if o == nil || o.Slot == nil {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *StorageEnclosureDiskSlotEp) HasSlot() bool {
	if o != nil && o.Slot != nil {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *StorageEnclosureDiskSlotEp) SetSlot(v string) {
	o.Slot = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageEnclosureDiskSlotEp) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEp) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageEnclosureDiskSlotEp) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageEnclosureDiskSlotEp) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageEnclosure returns the StorageEnclosure field value if set, zero value otherwise.
func (o *StorageEnclosureDiskSlotEp) GetStorageEnclosure() StorageEnclosureRelationship {
	if o == nil || o.StorageEnclosure == nil {
		var ret StorageEnclosureRelationship
		return ret
	}
	return *o.StorageEnclosure
}

// GetStorageEnclosureOk returns a tuple with the StorageEnclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEp) GetStorageEnclosureOk() (*StorageEnclosureRelationship, bool) {
	if o == nil || o.StorageEnclosure == nil {
		return nil, false
	}
	return o.StorageEnclosure, true
}

// HasStorageEnclosure returns a boolean if a field has been set.
func (o *StorageEnclosureDiskSlotEp) HasStorageEnclosure() bool {
	if o != nil && o.StorageEnclosure != nil {
		return true
	}

	return false
}

// SetStorageEnclosure gets a reference to the given StorageEnclosureRelationship and assigns it to the StorageEnclosure field.
func (o *StorageEnclosureDiskSlotEp) SetStorageEnclosure(v StorageEnclosureRelationship) {
	o.StorageEnclosure = &v
}

func (o StorageEnclosureDiskSlotEp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if o.DrivePath != nil {
		toSerialize["DrivePath"] = o.DrivePath
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.Slot != nil {
		toSerialize["Slot"] = o.Slot
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageEnclosure != nil {
		toSerialize["StorageEnclosure"] = o.StorageEnclosure
	}
	return json.Marshal(toSerialize)
}

type NullableStorageEnclosureDiskSlotEp struct {
	value *StorageEnclosureDiskSlotEp
	isSet bool
}

func (v NullableStorageEnclosureDiskSlotEp) Get() *StorageEnclosureDiskSlotEp {
	return v.value
}

func (v *NullableStorageEnclosureDiskSlotEp) Set(val *StorageEnclosureDiskSlotEp) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageEnclosureDiskSlotEp) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageEnclosureDiskSlotEp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageEnclosureDiskSlotEp(val *StorageEnclosureDiskSlotEp) *NullableStorageEnclosureDiskSlotEp {
	return &NullableStorageEnclosureDiskSlotEp{value: val, isSet: true}
}

func (v NullableStorageEnclosureDiskSlotEp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageEnclosureDiskSlotEp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
