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

// AssetDeviceClaimAllOf Definition of the list of properties defined in 'asset.DeviceClaim', excluding properties defined in parent classes.
type AssetDeviceClaimAllOf struct {
	DeviceUpdates *[]AssetConnectionControlMessage `json:"DeviceUpdates,omitempty" yaml:"DeviceUpdates,omitempty"`
	// Obtained from the device connector management UI or API (REST endpoint '/connector/SecurityTokens').
	SecurityToken *string `json:"SecurityToken,omitempty" yaml:"SecurityToken,omitempty"`
	// Obtained from the device connector management UI or API (REST endpoint '/connector/DeviceIdentifiers').
	SerialNumber *string                              `json:"SerialNumber,omitempty" yaml:"SerialNumber,omitempty"`
	Account      *IamAccountRelationship              `json:"Account,omitempty" yaml:"Account,omitempty"`
	Device       *AssetDeviceRegistrationRelationship `json:"Device,omitempty" yaml:"Device,omitempty"`
}

// NewAssetDeviceClaimAllOf instantiates a new AssetDeviceClaimAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceClaimAllOf() *AssetDeviceClaimAllOf {
	this := AssetDeviceClaimAllOf{}
	return &this
}

// NewAssetDeviceClaimAllOfWithDefaults instantiates a new AssetDeviceClaimAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceClaimAllOfWithDefaults() *AssetDeviceClaimAllOf {
	this := AssetDeviceClaimAllOf{}
	return &this
}

// GetDeviceUpdates returns the DeviceUpdates field value if set, zero value otherwise.
func (o *AssetDeviceClaimAllOf) GetDeviceUpdates() []AssetConnectionControlMessage {
	if o == nil || o.DeviceUpdates == nil {
		var ret []AssetConnectionControlMessage
		return ret
	}
	return *o.DeviceUpdates
}

// GetDeviceUpdatesOk returns a tuple with the DeviceUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaimAllOf) GetDeviceUpdatesOk() (*[]AssetConnectionControlMessage, bool) {
	if o == nil || o.DeviceUpdates == nil {
		return nil, false
	}
	return o.DeviceUpdates, true
}

// HasDeviceUpdates returns a boolean if a field has been set.
func (o *AssetDeviceClaimAllOf) HasDeviceUpdates() bool {
	if o != nil && o.DeviceUpdates != nil {
		return true
	}

	return false
}

// SetDeviceUpdates gets a reference to the given []AssetConnectionControlMessage and assigns it to the DeviceUpdates field.
func (o *AssetDeviceClaimAllOf) SetDeviceUpdates(v []AssetConnectionControlMessage) {
	o.DeviceUpdates = &v
}

// GetSecurityToken returns the SecurityToken field value if set, zero value otherwise.
func (o *AssetDeviceClaimAllOf) GetSecurityToken() string {
	if o == nil || o.SecurityToken == nil {
		var ret string
		return ret
	}
	return *o.SecurityToken
}

// GetSecurityTokenOk returns a tuple with the SecurityToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaimAllOf) GetSecurityTokenOk() (*string, bool) {
	if o == nil || o.SecurityToken == nil {
		return nil, false
	}
	return o.SecurityToken, true
}

// HasSecurityToken returns a boolean if a field has been set.
func (o *AssetDeviceClaimAllOf) HasSecurityToken() bool {
	if o != nil && o.SecurityToken != nil {
		return true
	}

	return false
}

// SetSecurityToken gets a reference to the given string and assigns it to the SecurityToken field.
func (o *AssetDeviceClaimAllOf) SetSecurityToken(v string) {
	o.SecurityToken = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *AssetDeviceClaimAllOf) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaimAllOf) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *AssetDeviceClaimAllOf) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *AssetDeviceClaimAllOf) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AssetDeviceClaimAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaimAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AssetDeviceClaimAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *AssetDeviceClaimAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *AssetDeviceClaimAllOf) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaimAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *AssetDeviceClaimAllOf) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *AssetDeviceClaimAllOf) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

func (o AssetDeviceClaimAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceUpdates != nil {
		toSerialize["DeviceUpdates"] = o.DeviceUpdates
	}
	if o.SecurityToken != nil {
		toSerialize["SecurityToken"] = o.SecurityToken
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	return json.Marshal(toSerialize)
}

type NullableAssetDeviceClaimAllOf struct {
	value *AssetDeviceClaimAllOf
	isSet bool
}

func (v NullableAssetDeviceClaimAllOf) Get() *AssetDeviceClaimAllOf {
	return v.value
}

func (v *NullableAssetDeviceClaimAllOf) Set(val *AssetDeviceClaimAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceClaimAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceClaimAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceClaimAllOf(val *AssetDeviceClaimAllOf) *NullableAssetDeviceClaimAllOf {
	return &NullableAssetDeviceClaimAllOf{value: val, isSet: true}
}

func (v NullableAssetDeviceClaimAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceClaimAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
