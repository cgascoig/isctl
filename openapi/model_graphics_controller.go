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

// GraphicsController Controller for a Graphics Card.
type GraphicsController struct {
	EquipmentBase `yaml:"EquipmentBase,inline"`
	// It shows the id of graphics controller.
	ControllerId *int64 `json:"ControllerId,omitempty" yaml:"ControllerId,omitempty"`
	// It shows the pci address of graphics controller.
	PciAddr *string `json:"PciAddr,omitempty" yaml:"PciAddr,omitempty"`
	// It shows the pci slot inforamtion of graphics controller.
	PciSlot          *string                              `json:"PciSlot,omitempty" yaml:"PciSlot,omitempty"`
	GraphicsCard     *GraphicsCardRelationship            `json:"GraphicsCard,omitempty" yaml:"GraphicsCard,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewGraphicsController instantiates a new GraphicsController object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphicsController() *GraphicsController {
	this := GraphicsController{}
	return &this
}

// NewGraphicsControllerWithDefaults instantiates a new GraphicsController object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphicsControllerWithDefaults() *GraphicsController {
	this := GraphicsController{}
	return &this
}

// GetControllerId returns the ControllerId field value if set, zero value otherwise.
func (o *GraphicsController) GetControllerId() int64 {
	if o == nil || o.ControllerId == nil {
		var ret int64
		return ret
	}
	return *o.ControllerId
}

// GetControllerIdOk returns a tuple with the ControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsController) GetControllerIdOk() (*int64, bool) {
	if o == nil || o.ControllerId == nil {
		return nil, false
	}
	return o.ControllerId, true
}

// HasControllerId returns a boolean if a field has been set.
func (o *GraphicsController) HasControllerId() bool {
	if o != nil && o.ControllerId != nil {
		return true
	}

	return false
}

// SetControllerId gets a reference to the given int64 and assigns it to the ControllerId field.
func (o *GraphicsController) SetControllerId(v int64) {
	o.ControllerId = &v
}

// GetPciAddr returns the PciAddr field value if set, zero value otherwise.
func (o *GraphicsController) GetPciAddr() string {
	if o == nil || o.PciAddr == nil {
		var ret string
		return ret
	}
	return *o.PciAddr
}

// GetPciAddrOk returns a tuple with the PciAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsController) GetPciAddrOk() (*string, bool) {
	if o == nil || o.PciAddr == nil {
		return nil, false
	}
	return o.PciAddr, true
}

// HasPciAddr returns a boolean if a field has been set.
func (o *GraphicsController) HasPciAddr() bool {
	if o != nil && o.PciAddr != nil {
		return true
	}

	return false
}

// SetPciAddr gets a reference to the given string and assigns it to the PciAddr field.
func (o *GraphicsController) SetPciAddr(v string) {
	o.PciAddr = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *GraphicsController) GetPciSlot() string {
	if o == nil || o.PciSlot == nil {
		var ret string
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsController) GetPciSlotOk() (*string, bool) {
	if o == nil || o.PciSlot == nil {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *GraphicsController) HasPciSlot() bool {
	if o != nil && o.PciSlot != nil {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given string and assigns it to the PciSlot field.
func (o *GraphicsController) SetPciSlot(v string) {
	o.PciSlot = &v
}

// GetGraphicsCard returns the GraphicsCard field value if set, zero value otherwise.
func (o *GraphicsController) GetGraphicsCard() GraphicsCardRelationship {
	if o == nil || o.GraphicsCard == nil {
		var ret GraphicsCardRelationship
		return ret
	}
	return *o.GraphicsCard
}

// GetGraphicsCardOk returns a tuple with the GraphicsCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsController) GetGraphicsCardOk() (*GraphicsCardRelationship, bool) {
	if o == nil || o.GraphicsCard == nil {
		return nil, false
	}
	return o.GraphicsCard, true
}

// HasGraphicsCard returns a boolean if a field has been set.
func (o *GraphicsController) HasGraphicsCard() bool {
	if o != nil && o.GraphicsCard != nil {
		return true
	}

	return false
}

// SetGraphicsCard gets a reference to the given GraphicsCardRelationship and assigns it to the GraphicsCard field.
func (o *GraphicsController) SetGraphicsCard(v GraphicsCardRelationship) {
	o.GraphicsCard = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *GraphicsController) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsController) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *GraphicsController) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *GraphicsController) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o GraphicsController) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if o.ControllerId != nil {
		toSerialize["ControllerId"] = o.ControllerId
	}
	if o.PciAddr != nil {
		toSerialize["PciAddr"] = o.PciAddr
	}
	if o.PciSlot != nil {
		toSerialize["PciSlot"] = o.PciSlot
	}
	if o.GraphicsCard != nil {
		toSerialize["GraphicsCard"] = o.GraphicsCard
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableGraphicsController struct {
	value *GraphicsController
	isSet bool
}

func (v NullableGraphicsController) Get() *GraphicsController {
	return v.value
}

func (v *NullableGraphicsController) Set(val *GraphicsController) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphicsController) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphicsController) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphicsController(val *GraphicsController) *NullableGraphicsController {
	return &NullableGraphicsController{value: val, isSet: true}
}

func (v NullableGraphicsController) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphicsController) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
