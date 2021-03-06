/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-11T18:30:19Z.
 *
 * API version: 1.0.9-3252
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// AssetDeviceConfiguration The configuration of a device connector. Configuration properties may be changed by a Intersight user or by a device administrator using the connector's API exposed through the platforms management interface.
type AssetDeviceConfiguration struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Specifies whether configuration through the platforms local management interface has been disabled, with only configuration through the Intersight service enabled.
	LocalConfigurationLocked *bool `json:"LocalConfigurationLocked,omitempty" yaml:"LocalConfigurationLocked,omitempty"`
	// The log level of the device connector service.
	LogLevel *string                              `json:"LogLevel,omitempty" yaml:"LogLevel,omitempty"`
	Device   *AssetDeviceRegistrationRelationship `json:"Device,omitempty" yaml:"Device,omitempty"`
}

// NewAssetDeviceConfiguration instantiates a new AssetDeviceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceConfiguration(classId string, objectType string) *AssetDeviceConfiguration {
	this := AssetDeviceConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetDeviceConfigurationWithDefaults instantiates a new AssetDeviceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceConfigurationWithDefaults() *AssetDeviceConfiguration {
	this := AssetDeviceConfiguration{}
	var classId string = "asset.DeviceConfiguration"
	this.ClassId = classId
	var objectType string = "asset.DeviceConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetDeviceConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetDeviceConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetDeviceConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetDeviceConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLocalConfigurationLocked returns the LocalConfigurationLocked field value if set, zero value otherwise.
func (o *AssetDeviceConfiguration) GetLocalConfigurationLocked() bool {
	if o == nil || o.LocalConfigurationLocked == nil {
		var ret bool
		return ret
	}
	return *o.LocalConfigurationLocked
}

// GetLocalConfigurationLockedOk returns a tuple with the LocalConfigurationLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConfiguration) GetLocalConfigurationLockedOk() (*bool, bool) {
	if o == nil || o.LocalConfigurationLocked == nil {
		return nil, false
	}
	return o.LocalConfigurationLocked, true
}

// HasLocalConfigurationLocked returns a boolean if a field has been set.
func (o *AssetDeviceConfiguration) HasLocalConfigurationLocked() bool {
	if o != nil && o.LocalConfigurationLocked != nil {
		return true
	}

	return false
}

// SetLocalConfigurationLocked gets a reference to the given bool and assigns it to the LocalConfigurationLocked field.
func (o *AssetDeviceConfiguration) SetLocalConfigurationLocked(v bool) {
	o.LocalConfigurationLocked = &v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *AssetDeviceConfiguration) GetLogLevel() string {
	if o == nil || o.LogLevel == nil {
		var ret string
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConfiguration) GetLogLevelOk() (*string, bool) {
	if o == nil || o.LogLevel == nil {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *AssetDeviceConfiguration) HasLogLevel() bool {
	if o != nil && o.LogLevel != nil {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given string and assigns it to the LogLevel field.
func (o *AssetDeviceConfiguration) SetLogLevel(v string) {
	o.LogLevel = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *AssetDeviceConfiguration) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConfiguration) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *AssetDeviceConfiguration) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *AssetDeviceConfiguration) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

func (o AssetDeviceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LocalConfigurationLocked != nil {
		toSerialize["LocalConfigurationLocked"] = o.LocalConfigurationLocked
	}
	if o.LogLevel != nil {
		toSerialize["LogLevel"] = o.LogLevel
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	return json.Marshal(toSerialize)
}

type NullableAssetDeviceConfiguration struct {
	value *AssetDeviceConfiguration
	isSet bool
}

func (v NullableAssetDeviceConfiguration) Get() *AssetDeviceConfiguration {
	return v.value
}

func (v *NullableAssetDeviceConfiguration) Set(val *AssetDeviceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceConfiguration(val *AssetDeviceConfiguration) *NullableAssetDeviceConfiguration {
	return &NullableAssetDeviceConfiguration{value: val, isSet: true}
}

func (v NullableAssetDeviceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
