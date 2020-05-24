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

// NiaapiDcnmLatestMaintainedRelease This contains the latest maintained release information for each release on DCNM.
type NiaapiDcnmLatestMaintainedRelease struct {
	NiaapiMaintainedRelease `yaml:"NiaapiMaintainedRelease,inline"`
}

// NewNiaapiDcnmLatestMaintainedRelease instantiates a new NiaapiDcnmLatestMaintainedRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiDcnmLatestMaintainedRelease() *NiaapiDcnmLatestMaintainedRelease {
	this := NiaapiDcnmLatestMaintainedRelease{}
	return &this
}

// NewNiaapiDcnmLatestMaintainedReleaseWithDefaults instantiates a new NiaapiDcnmLatestMaintainedRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiDcnmLatestMaintainedReleaseWithDefaults() *NiaapiDcnmLatestMaintainedRelease {
	this := NiaapiDcnmLatestMaintainedRelease{}
	return &this
}

func (o NiaapiDcnmLatestMaintainedRelease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedNiaapiMaintainedRelease, errNiaapiMaintainedRelease := json.Marshal(o.NiaapiMaintainedRelease)
	if errNiaapiMaintainedRelease != nil {
		return []byte{}, errNiaapiMaintainedRelease
	}
	errNiaapiMaintainedRelease = json.Unmarshal([]byte(serializedNiaapiMaintainedRelease), &toSerialize)
	if errNiaapiMaintainedRelease != nil {
		return []byte{}, errNiaapiMaintainedRelease
	}
	return json.Marshal(toSerialize)
}

type NullableNiaapiDcnmLatestMaintainedRelease struct {
	value *NiaapiDcnmLatestMaintainedRelease
	isSet bool
}

func (v NullableNiaapiDcnmLatestMaintainedRelease) Get() *NiaapiDcnmLatestMaintainedRelease {
	return v.value
}

func (v *NullableNiaapiDcnmLatestMaintainedRelease) Set(val *NiaapiDcnmLatestMaintainedRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiDcnmLatestMaintainedRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiDcnmLatestMaintainedRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiDcnmLatestMaintainedRelease(val *NiaapiDcnmLatestMaintainedRelease) *NullableNiaapiDcnmLatestMaintainedRelease {
	return &NullableNiaapiDcnmLatestMaintainedRelease{value: val, isSet: true}
}

func (v NullableNiaapiDcnmLatestMaintainedRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiDcnmLatestMaintainedRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
