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

// VirtualizationVmDiskCommitInfoAllOf Definition of the list of properties defined in 'virtualization.VmDiskCommitInfo', excluding properties defined in parent classes.
type VirtualizationVmDiskCommitInfoAllOf struct {
	// Disk committed in bytes on this virtual machine (disk space used up).
	CommittedDisk *int64 `json:"CommittedDisk,omitempty" yaml:"CommittedDisk,omitempty"`
	// Total uncommitted disk space that is available for use (in bytes).
	UnCommittedDisk *int64 `json:"UnCommittedDisk,omitempty" yaml:"UnCommittedDisk,omitempty"`
	// Total unshared disk space (in bytes).
	UnsharedDisk *int64 `json:"UnsharedDisk,omitempty" yaml:"UnsharedDisk,omitempty"`
}

// NewVirtualizationVmDiskCommitInfoAllOf instantiates a new VirtualizationVmDiskCommitInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmDiskCommitInfoAllOf() *VirtualizationVmDiskCommitInfoAllOf {
	this := VirtualizationVmDiskCommitInfoAllOf{}
	return &this
}

// NewVirtualizationVmDiskCommitInfoAllOfWithDefaults instantiates a new VirtualizationVmDiskCommitInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmDiskCommitInfoAllOfWithDefaults() *VirtualizationVmDiskCommitInfoAllOf {
	this := VirtualizationVmDiskCommitInfoAllOf{}
	return &this
}

// GetCommittedDisk returns the CommittedDisk field value if set, zero value otherwise.
func (o *VirtualizationVmDiskCommitInfoAllOf) GetCommittedDisk() int64 {
	if o == nil || o.CommittedDisk == nil {
		var ret int64
		return ret
	}
	return *o.CommittedDisk
}

// GetCommittedDiskOk returns a tuple with the CommittedDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmDiskCommitInfoAllOf) GetCommittedDiskOk() (*int64, bool) {
	if o == nil || o.CommittedDisk == nil {
		return nil, false
	}
	return o.CommittedDisk, true
}

// HasCommittedDisk returns a boolean if a field has been set.
func (o *VirtualizationVmDiskCommitInfoAllOf) HasCommittedDisk() bool {
	if o != nil && o.CommittedDisk != nil {
		return true
	}

	return false
}

// SetCommittedDisk gets a reference to the given int64 and assigns it to the CommittedDisk field.
func (o *VirtualizationVmDiskCommitInfoAllOf) SetCommittedDisk(v int64) {
	o.CommittedDisk = &v
}

// GetUnCommittedDisk returns the UnCommittedDisk field value if set, zero value otherwise.
func (o *VirtualizationVmDiskCommitInfoAllOf) GetUnCommittedDisk() int64 {
	if o == nil || o.UnCommittedDisk == nil {
		var ret int64
		return ret
	}
	return *o.UnCommittedDisk
}

// GetUnCommittedDiskOk returns a tuple with the UnCommittedDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmDiskCommitInfoAllOf) GetUnCommittedDiskOk() (*int64, bool) {
	if o == nil || o.UnCommittedDisk == nil {
		return nil, false
	}
	return o.UnCommittedDisk, true
}

// HasUnCommittedDisk returns a boolean if a field has been set.
func (o *VirtualizationVmDiskCommitInfoAllOf) HasUnCommittedDisk() bool {
	if o != nil && o.UnCommittedDisk != nil {
		return true
	}

	return false
}

// SetUnCommittedDisk gets a reference to the given int64 and assigns it to the UnCommittedDisk field.
func (o *VirtualizationVmDiskCommitInfoAllOf) SetUnCommittedDisk(v int64) {
	o.UnCommittedDisk = &v
}

// GetUnsharedDisk returns the UnsharedDisk field value if set, zero value otherwise.
func (o *VirtualizationVmDiskCommitInfoAllOf) GetUnsharedDisk() int64 {
	if o == nil || o.UnsharedDisk == nil {
		var ret int64
		return ret
	}
	return *o.UnsharedDisk
}

// GetUnsharedDiskOk returns a tuple with the UnsharedDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmDiskCommitInfoAllOf) GetUnsharedDiskOk() (*int64, bool) {
	if o == nil || o.UnsharedDisk == nil {
		return nil, false
	}
	return o.UnsharedDisk, true
}

// HasUnsharedDisk returns a boolean if a field has been set.
func (o *VirtualizationVmDiskCommitInfoAllOf) HasUnsharedDisk() bool {
	if o != nil && o.UnsharedDisk != nil {
		return true
	}

	return false
}

// SetUnsharedDisk gets a reference to the given int64 and assigns it to the UnsharedDisk field.
func (o *VirtualizationVmDiskCommitInfoAllOf) SetUnsharedDisk(v int64) {
	o.UnsharedDisk = &v
}

func (o VirtualizationVmDiskCommitInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommittedDisk != nil {
		toSerialize["CommittedDisk"] = o.CommittedDisk
	}
	if o.UnCommittedDisk != nil {
		toSerialize["UnCommittedDisk"] = o.UnCommittedDisk
	}
	if o.UnsharedDisk != nil {
		toSerialize["UnsharedDisk"] = o.UnsharedDisk
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualizationVmDiskCommitInfoAllOf struct {
	value *VirtualizationVmDiskCommitInfoAllOf
	isSet bool
}

func (v NullableVirtualizationVmDiskCommitInfoAllOf) Get() *VirtualizationVmDiskCommitInfoAllOf {
	return v.value
}

func (v *NullableVirtualizationVmDiskCommitInfoAllOf) Set(val *VirtualizationVmDiskCommitInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmDiskCommitInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmDiskCommitInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmDiskCommitInfoAllOf(val *VirtualizationVmDiskCommitInfoAllOf) *NullableVirtualizationVmDiskCommitInfoAllOf {
	return &NullableVirtualizationVmDiskCommitInfoAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmDiskCommitInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmDiskCommitInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
