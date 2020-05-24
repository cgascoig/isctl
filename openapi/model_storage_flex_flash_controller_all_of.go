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

// StorageFlexFlashControllerAllOf Definition of the list of properties defined in 'storage.FlexFlashController', excluding properties defined in parent classes.
type StorageFlexFlashControllerAllOf struct {
	ControllerState *string                   `json:"ControllerState,omitempty" yaml:"ControllerState,omitempty"`
	FfControllerId  *string                   `json:"FfControllerId,omitempty" yaml:"FfControllerId,omitempty"`
	ComputeBoard    *ComputeBoardRelationship `json:"ComputeBoard,omitempty" yaml:"ComputeBoard,omitempty"`
	// An array of relationships to storageFlexFlashControllerProps resources.
	FlexFlashControllerProps *[]StorageFlexFlashControllerPropsRelationship `json:"FlexFlashControllerProps,omitempty" yaml:"FlexFlashControllerProps,omitempty"`
	// An array of relationships to storageFlexFlashPhysicalDrive resources.
	FlexFlashPhysicalDrives *[]StorageFlexFlashPhysicalDriveRelationship `json:"FlexFlashPhysicalDrives,omitempty" yaml:"FlexFlashPhysicalDrives,omitempty"`
	// An array of relationships to storageFlexFlashVirtualDrive resources.
	FlexFlashVirtualDrives *[]StorageFlexFlashVirtualDriveRelationship `json:"FlexFlashVirtualDrives,omitempty" yaml:"FlexFlashVirtualDrives,omitempty"`
	RegisteredDevice       *AssetDeviceRegistrationRelationship        `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewStorageFlexFlashControllerAllOf instantiates a new StorageFlexFlashControllerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageFlexFlashControllerAllOf() *StorageFlexFlashControllerAllOf {
	this := StorageFlexFlashControllerAllOf{}
	return &this
}

// NewStorageFlexFlashControllerAllOfWithDefaults instantiates a new StorageFlexFlashControllerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageFlexFlashControllerAllOfWithDefaults() *StorageFlexFlashControllerAllOf {
	this := StorageFlexFlashControllerAllOf{}
	return &this
}

// GetControllerState returns the ControllerState field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerAllOf) GetControllerState() string {
	if o == nil || o.ControllerState == nil {
		var ret string
		return ret
	}
	return *o.ControllerState
}

// GetControllerStateOk returns a tuple with the ControllerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerAllOf) GetControllerStateOk() (*string, bool) {
	if o == nil || o.ControllerState == nil {
		return nil, false
	}
	return o.ControllerState, true
}

// HasControllerState returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasControllerState() bool {
	if o != nil && o.ControllerState != nil {
		return true
	}

	return false
}

// SetControllerState gets a reference to the given string and assigns it to the ControllerState field.
func (o *StorageFlexFlashControllerAllOf) SetControllerState(v string) {
	o.ControllerState = &v
}

// GetFfControllerId returns the FfControllerId field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerAllOf) GetFfControllerId() string {
	if o == nil || o.FfControllerId == nil {
		var ret string
		return ret
	}
	return *o.FfControllerId
}

// GetFfControllerIdOk returns a tuple with the FfControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerAllOf) GetFfControllerIdOk() (*string, bool) {
	if o == nil || o.FfControllerId == nil {
		return nil, false
	}
	return o.FfControllerId, true
}

// HasFfControllerId returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasFfControllerId() bool {
	if o != nil && o.FfControllerId != nil {
		return true
	}

	return false
}

// SetFfControllerId gets a reference to the given string and assigns it to the FfControllerId field.
func (o *StorageFlexFlashControllerAllOf) SetFfControllerId(v string) {
	o.FfControllerId = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerAllOf) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || o.ComputeBoard == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.ComputeBoard == nil {
		return nil, false
	}
	return o.ComputeBoard, true
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard != nil {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given ComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *StorageFlexFlashControllerAllOf) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard = &v
}

// GetFlexFlashControllerProps returns the FlexFlashControllerProps field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerAllOf) GetFlexFlashControllerProps() []StorageFlexFlashControllerPropsRelationship {
	if o == nil || o.FlexFlashControllerProps == nil {
		var ret []StorageFlexFlashControllerPropsRelationship
		return ret
	}
	return *o.FlexFlashControllerProps
}

// GetFlexFlashControllerPropsOk returns a tuple with the FlexFlashControllerProps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerAllOf) GetFlexFlashControllerPropsOk() (*[]StorageFlexFlashControllerPropsRelationship, bool) {
	if o == nil || o.FlexFlashControllerProps == nil {
		return nil, false
	}
	return o.FlexFlashControllerProps, true
}

// HasFlexFlashControllerProps returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasFlexFlashControllerProps() bool {
	if o != nil && o.FlexFlashControllerProps != nil {
		return true
	}

	return false
}

// SetFlexFlashControllerProps gets a reference to the given []StorageFlexFlashControllerPropsRelationship and assigns it to the FlexFlashControllerProps field.
func (o *StorageFlexFlashControllerAllOf) SetFlexFlashControllerProps(v []StorageFlexFlashControllerPropsRelationship) {
	o.FlexFlashControllerProps = &v
}

// GetFlexFlashPhysicalDrives returns the FlexFlashPhysicalDrives field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerAllOf) GetFlexFlashPhysicalDrives() []StorageFlexFlashPhysicalDriveRelationship {
	if o == nil || o.FlexFlashPhysicalDrives == nil {
		var ret []StorageFlexFlashPhysicalDriveRelationship
		return ret
	}
	return *o.FlexFlashPhysicalDrives
}

// GetFlexFlashPhysicalDrivesOk returns a tuple with the FlexFlashPhysicalDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerAllOf) GetFlexFlashPhysicalDrivesOk() (*[]StorageFlexFlashPhysicalDriveRelationship, bool) {
	if o == nil || o.FlexFlashPhysicalDrives == nil {
		return nil, false
	}
	return o.FlexFlashPhysicalDrives, true
}

// HasFlexFlashPhysicalDrives returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasFlexFlashPhysicalDrives() bool {
	if o != nil && o.FlexFlashPhysicalDrives != nil {
		return true
	}

	return false
}

// SetFlexFlashPhysicalDrives gets a reference to the given []StorageFlexFlashPhysicalDriveRelationship and assigns it to the FlexFlashPhysicalDrives field.
func (o *StorageFlexFlashControllerAllOf) SetFlexFlashPhysicalDrives(v []StorageFlexFlashPhysicalDriveRelationship) {
	o.FlexFlashPhysicalDrives = &v
}

// GetFlexFlashVirtualDrives returns the FlexFlashVirtualDrives field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerAllOf) GetFlexFlashVirtualDrives() []StorageFlexFlashVirtualDriveRelationship {
	if o == nil || o.FlexFlashVirtualDrives == nil {
		var ret []StorageFlexFlashVirtualDriveRelationship
		return ret
	}
	return *o.FlexFlashVirtualDrives
}

// GetFlexFlashVirtualDrivesOk returns a tuple with the FlexFlashVirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerAllOf) GetFlexFlashVirtualDrivesOk() (*[]StorageFlexFlashVirtualDriveRelationship, bool) {
	if o == nil || o.FlexFlashVirtualDrives == nil {
		return nil, false
	}
	return o.FlexFlashVirtualDrives, true
}

// HasFlexFlashVirtualDrives returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasFlexFlashVirtualDrives() bool {
	if o != nil && o.FlexFlashVirtualDrives != nil {
		return true
	}

	return false
}

// SetFlexFlashVirtualDrives gets a reference to the given []StorageFlexFlashVirtualDriveRelationship and assigns it to the FlexFlashVirtualDrives field.
func (o *StorageFlexFlashControllerAllOf) SetFlexFlashVirtualDrives(v []StorageFlexFlashVirtualDriveRelationship) {
	o.FlexFlashVirtualDrives = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageFlexFlashControllerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageFlexFlashControllerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ControllerState != nil {
		toSerialize["ControllerState"] = o.ControllerState
	}
	if o.FfControllerId != nil {
		toSerialize["FfControllerId"] = o.FfControllerId
	}
	if o.ComputeBoard != nil {
		toSerialize["ComputeBoard"] = o.ComputeBoard
	}
	if o.FlexFlashControllerProps != nil {
		toSerialize["FlexFlashControllerProps"] = o.FlexFlashControllerProps
	}
	if o.FlexFlashPhysicalDrives != nil {
		toSerialize["FlexFlashPhysicalDrives"] = o.FlexFlashPhysicalDrives
	}
	if o.FlexFlashVirtualDrives != nil {
		toSerialize["FlexFlashVirtualDrives"] = o.FlexFlashVirtualDrives
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableStorageFlexFlashControllerAllOf struct {
	value *StorageFlexFlashControllerAllOf
	isSet bool
}

func (v NullableStorageFlexFlashControllerAllOf) Get() *StorageFlexFlashControllerAllOf {
	return v.value
}

func (v *NullableStorageFlexFlashControllerAllOf) Set(val *StorageFlexFlashControllerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFlexFlashControllerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFlexFlashControllerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFlexFlashControllerAllOf(val *StorageFlexFlashControllerAllOf) *NullableStorageFlexFlashControllerAllOf {
	return &NullableStorageFlexFlashControllerAllOf{value: val, isSet: true}
}

func (v NullableStorageFlexFlashControllerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFlexFlashControllerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
