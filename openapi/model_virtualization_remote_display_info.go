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

// VirtualizationRemoteDisplayInfo Details of the remote display settings.
type VirtualizationRemoteDisplayInfo struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// The password used for remote access. It is stored base64 encoded.
	RemoteDisplayPassword *string `json:"RemoteDisplayPassword,omitempty" yaml:"RemoteDisplayPassword,omitempty"`
	// The access key for the remote display, potentially a long string.
	RemoteDisplayVncKey *string `json:"RemoteDisplayVncKey,omitempty" yaml:"RemoteDisplayVncKey,omitempty"`
	// Applies only when remoteDisplayvnc is enabled.
	RemoteDisplayVncPort *int64 `json:"RemoteDisplayVncPort,omitempty" yaml:"RemoteDisplayVncPort,omitempty"`
}

// NewVirtualizationRemoteDisplayInfo instantiates a new VirtualizationRemoteDisplayInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationRemoteDisplayInfo() *VirtualizationRemoteDisplayInfo {
	this := VirtualizationRemoteDisplayInfo{}
	return &this
}

// NewVirtualizationRemoteDisplayInfoWithDefaults instantiates a new VirtualizationRemoteDisplayInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationRemoteDisplayInfoWithDefaults() *VirtualizationRemoteDisplayInfo {
	this := VirtualizationRemoteDisplayInfo{}
	return &this
}

// GetRemoteDisplayPassword returns the RemoteDisplayPassword field value if set, zero value otherwise.
func (o *VirtualizationRemoteDisplayInfo) GetRemoteDisplayPassword() string {
	if o == nil || o.RemoteDisplayPassword == nil {
		var ret string
		return ret
	}
	return *o.RemoteDisplayPassword
}

// GetRemoteDisplayPasswordOk returns a tuple with the RemoteDisplayPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationRemoteDisplayInfo) GetRemoteDisplayPasswordOk() (*string, bool) {
	if o == nil || o.RemoteDisplayPassword == nil {
		return nil, false
	}
	return o.RemoteDisplayPassword, true
}

// HasRemoteDisplayPassword returns a boolean if a field has been set.
func (o *VirtualizationRemoteDisplayInfo) HasRemoteDisplayPassword() bool {
	if o != nil && o.RemoteDisplayPassword != nil {
		return true
	}

	return false
}

// SetRemoteDisplayPassword gets a reference to the given string and assigns it to the RemoteDisplayPassword field.
func (o *VirtualizationRemoteDisplayInfo) SetRemoteDisplayPassword(v string) {
	o.RemoteDisplayPassword = &v
}

// GetRemoteDisplayVncKey returns the RemoteDisplayVncKey field value if set, zero value otherwise.
func (o *VirtualizationRemoteDisplayInfo) GetRemoteDisplayVncKey() string {
	if o == nil || o.RemoteDisplayVncKey == nil {
		var ret string
		return ret
	}
	return *o.RemoteDisplayVncKey
}

// GetRemoteDisplayVncKeyOk returns a tuple with the RemoteDisplayVncKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationRemoteDisplayInfo) GetRemoteDisplayVncKeyOk() (*string, bool) {
	if o == nil || o.RemoteDisplayVncKey == nil {
		return nil, false
	}
	return o.RemoteDisplayVncKey, true
}

// HasRemoteDisplayVncKey returns a boolean if a field has been set.
func (o *VirtualizationRemoteDisplayInfo) HasRemoteDisplayVncKey() bool {
	if o != nil && o.RemoteDisplayVncKey != nil {
		return true
	}

	return false
}

// SetRemoteDisplayVncKey gets a reference to the given string and assigns it to the RemoteDisplayVncKey field.
func (o *VirtualizationRemoteDisplayInfo) SetRemoteDisplayVncKey(v string) {
	o.RemoteDisplayVncKey = &v
}

// GetRemoteDisplayVncPort returns the RemoteDisplayVncPort field value if set, zero value otherwise.
func (o *VirtualizationRemoteDisplayInfo) GetRemoteDisplayVncPort() int64 {
	if o == nil || o.RemoteDisplayVncPort == nil {
		var ret int64
		return ret
	}
	return *o.RemoteDisplayVncPort
}

// GetRemoteDisplayVncPortOk returns a tuple with the RemoteDisplayVncPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationRemoteDisplayInfo) GetRemoteDisplayVncPortOk() (*int64, bool) {
	if o == nil || o.RemoteDisplayVncPort == nil {
		return nil, false
	}
	return o.RemoteDisplayVncPort, true
}

// HasRemoteDisplayVncPort returns a boolean if a field has been set.
func (o *VirtualizationRemoteDisplayInfo) HasRemoteDisplayVncPort() bool {
	if o != nil && o.RemoteDisplayVncPort != nil {
		return true
	}

	return false
}

// SetRemoteDisplayVncPort gets a reference to the given int64 and assigns it to the RemoteDisplayVncPort field.
func (o *VirtualizationRemoteDisplayInfo) SetRemoteDisplayVncPort(v int64) {
	o.RemoteDisplayVncPort = &v
}

func (o VirtualizationRemoteDisplayInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.RemoteDisplayPassword != nil {
		toSerialize["RemoteDisplayPassword"] = o.RemoteDisplayPassword
	}
	if o.RemoteDisplayVncKey != nil {
		toSerialize["RemoteDisplayVncKey"] = o.RemoteDisplayVncKey
	}
	if o.RemoteDisplayVncPort != nil {
		toSerialize["RemoteDisplayVncPort"] = o.RemoteDisplayVncPort
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualizationRemoteDisplayInfo struct {
	value *VirtualizationRemoteDisplayInfo
	isSet bool
}

func (v NullableVirtualizationRemoteDisplayInfo) Get() *VirtualizationRemoteDisplayInfo {
	return v.value
}

func (v *NullableVirtualizationRemoteDisplayInfo) Set(val *VirtualizationRemoteDisplayInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationRemoteDisplayInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationRemoteDisplayInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationRemoteDisplayInfo(val *VirtualizationRemoteDisplayInfo) *NullableVirtualizationRemoteDisplayInfo {
	return &NullableVirtualizationRemoteDisplayInfo{value: val, isSet: true}
}

func (v NullableVirtualizationRemoteDisplayInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationRemoteDisplayInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
