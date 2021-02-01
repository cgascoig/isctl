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

// VirtualizationBaseVmConfiguration Specify ESXi virtual machine custom specification.
type VirtualizationBaseVmConfiguration struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
}

// NewVirtualizationBaseVmConfiguration instantiates a new VirtualizationBaseVmConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseVmConfiguration() *VirtualizationBaseVmConfiguration {
	this := VirtualizationBaseVmConfiguration{}
	return &this
}

// NewVirtualizationBaseVmConfigurationWithDefaults instantiates a new VirtualizationBaseVmConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseVmConfigurationWithDefaults() *VirtualizationBaseVmConfiguration {
	this := VirtualizationBaseVmConfiguration{}
	return &this
}

func (o VirtualizationBaseVmConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualizationBaseVmConfiguration struct {
	value *VirtualizationBaseVmConfiguration
	isSet bool
}

func (v NullableVirtualizationBaseVmConfiguration) Get() *VirtualizationBaseVmConfiguration {
	return v.value
}

func (v *NullableVirtualizationBaseVmConfiguration) Set(val *VirtualizationBaseVmConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseVmConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseVmConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseVmConfiguration(val *VirtualizationBaseVmConfiguration) *NullableVirtualizationBaseVmConfiguration {
	return &NullableVirtualizationBaseVmConfiguration{value: val, isSet: true}
}

func (v NullableVirtualizationBaseVmConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseVmConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
