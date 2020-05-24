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

// EquipmentSwitchCardAllOf Definition of the list of properties defined in 'equipment.SwitchCard', excluding properties defined in parent classes.
type EquipmentSwitchCardAllOf struct {
	Description    *string                     `json:"Description,omitempty" yaml:"Description,omitempty"`
	NumPorts       *int64                      `json:"NumPorts,omitempty" yaml:"NumPorts,omitempty"`
	Presence       *string                     `json:"Presence,omitempty" yaml:"Presence,omitempty"`
	SlotId         *int64                      `json:"SlotId,omitempty" yaml:"SlotId,omitempty"`
	State          *string                     `json:"State,omitempty" yaml:"State,omitempty"`
	NetworkElement *NetworkElementRelationship `json:"NetworkElement,omitempty" yaml:"NetworkElement,omitempty"`
	// An array of relationships to portGroup resources.
	PortGroups       *[]PortGroupRelationship             `json:"PortGroups,omitempty" yaml:"PortGroups,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewEquipmentSwitchCardAllOf instantiates a new EquipmentSwitchCardAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentSwitchCardAllOf() *EquipmentSwitchCardAllOf {
	this := EquipmentSwitchCardAllOf{}
	return &this
}

// NewEquipmentSwitchCardAllOfWithDefaults instantiates a new EquipmentSwitchCardAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentSwitchCardAllOfWithDefaults() *EquipmentSwitchCardAllOf {
	this := EquipmentSwitchCardAllOf{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EquipmentSwitchCardAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCardAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EquipmentSwitchCardAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EquipmentSwitchCardAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetNumPorts returns the NumPorts field value if set, zero value otherwise.
func (o *EquipmentSwitchCardAllOf) GetNumPorts() int64 {
	if o == nil || o.NumPorts == nil {
		var ret int64
		return ret
	}
	return *o.NumPorts
}

// GetNumPortsOk returns a tuple with the NumPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCardAllOf) GetNumPortsOk() (*int64, bool) {
	if o == nil || o.NumPorts == nil {
		return nil, false
	}
	return o.NumPorts, true
}

// HasNumPorts returns a boolean if a field has been set.
func (o *EquipmentSwitchCardAllOf) HasNumPorts() bool {
	if o != nil && o.NumPorts != nil {
		return true
	}

	return false
}

// SetNumPorts gets a reference to the given int64 and assigns it to the NumPorts field.
func (o *EquipmentSwitchCardAllOf) SetNumPorts(v int64) {
	o.NumPorts = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *EquipmentSwitchCardAllOf) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCardAllOf) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *EquipmentSwitchCardAllOf) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *EquipmentSwitchCardAllOf) SetPresence(v string) {
	o.Presence = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *EquipmentSwitchCardAllOf) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCardAllOf) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *EquipmentSwitchCardAllOf) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *EquipmentSwitchCardAllOf) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *EquipmentSwitchCardAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCardAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *EquipmentSwitchCardAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *EquipmentSwitchCardAllOf) SetState(v string) {
	o.State = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *EquipmentSwitchCardAllOf) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCardAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *EquipmentSwitchCardAllOf) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *EquipmentSwitchCardAllOf) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetPortGroups returns the PortGroups field value if set, zero value otherwise.
func (o *EquipmentSwitchCardAllOf) GetPortGroups() []PortGroupRelationship {
	if o == nil || o.PortGroups == nil {
		var ret []PortGroupRelationship
		return ret
	}
	return *o.PortGroups
}

// GetPortGroupsOk returns a tuple with the PortGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCardAllOf) GetPortGroupsOk() (*[]PortGroupRelationship, bool) {
	if o == nil || o.PortGroups == nil {
		return nil, false
	}
	return o.PortGroups, true
}

// HasPortGroups returns a boolean if a field has been set.
func (o *EquipmentSwitchCardAllOf) HasPortGroups() bool {
	if o != nil && o.PortGroups != nil {
		return true
	}

	return false
}

// SetPortGroups gets a reference to the given []PortGroupRelationship and assigns it to the PortGroups field.
func (o *EquipmentSwitchCardAllOf) SetPortGroups(v []PortGroupRelationship) {
	o.PortGroups = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentSwitchCardAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCardAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentSwitchCardAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentSwitchCardAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentSwitchCardAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.NumPorts != nil {
		toSerialize["NumPorts"] = o.NumPorts
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.PortGroups != nil {
		toSerialize["PortGroups"] = o.PortGroups
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableEquipmentSwitchCardAllOf struct {
	value *EquipmentSwitchCardAllOf
	isSet bool
}

func (v NullableEquipmentSwitchCardAllOf) Get() *EquipmentSwitchCardAllOf {
	return v.value
}

func (v *NullableEquipmentSwitchCardAllOf) Set(val *EquipmentSwitchCardAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentSwitchCardAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentSwitchCardAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentSwitchCardAllOf(val *EquipmentSwitchCardAllOf) *NullableEquipmentSwitchCardAllOf {
	return &NullableEquipmentSwitchCardAllOf{value: val, isSet: true}
}

func (v NullableEquipmentSwitchCardAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentSwitchCardAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
