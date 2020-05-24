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

// AssetDeviceConnectorManager Information pertaining to a Registered Intersight Assist Appliance Device in Intersight.
type AssetDeviceConnectorManager struct {
	MoBaseMo         `yaml:"MoBaseMo,inline"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewAssetDeviceConnectorManager instantiates a new AssetDeviceConnectorManager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceConnectorManager() *AssetDeviceConnectorManager {
	this := AssetDeviceConnectorManager{}
	return &this
}

// NewAssetDeviceConnectorManagerWithDefaults instantiates a new AssetDeviceConnectorManager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceConnectorManagerWithDefaults() *AssetDeviceConnectorManager {
	this := AssetDeviceConnectorManager{}
	return &this
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *AssetDeviceConnectorManager) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConnectorManager) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *AssetDeviceConnectorManager) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *AssetDeviceConnectorManager) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o AssetDeviceConnectorManager) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableAssetDeviceConnectorManager struct {
	value *AssetDeviceConnectorManager
	isSet bool
}

func (v NullableAssetDeviceConnectorManager) Get() *AssetDeviceConnectorManager {
	return v.value
}

func (v *NullableAssetDeviceConnectorManager) Set(val *AssetDeviceConnectorManager) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceConnectorManager) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceConnectorManager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceConnectorManager(val *AssetDeviceConnectorManager) *NullableAssetDeviceConnectorManager {
	return &NullableAssetDeviceConnectorManager{value: val, isSet: true}
}

func (v NullableAssetDeviceConnectorManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceConnectorManager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
