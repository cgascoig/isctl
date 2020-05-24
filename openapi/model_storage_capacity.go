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

// StorageCapacity Storage capacity information which includes, total capacity, available capacity, used capacity and free capacity.
type StorageCapacity struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// Total consumable storage capacity represented in bytes. System may reserve some space for internal purposes which is excluded from total capacity.
	Available *int64 `json:"Available,omitempty" yaml:"Available,omitempty"`
	// Unused space available for applications to consume, represented in bytes.
	Free *int64 `json:"Free,omitempty" yaml:"Free,omitempty"`
	// Total storage capacity, represented in bytes. It is set by the component manufacturer.
	Total *int64 `json:"Total,omitempty" yaml:"Total,omitempty"`
	// Used or consumed storage capacity, represented in bytes.
	Used *int64 `json:"Used,omitempty" yaml:"Used,omitempty"`
}

// NewStorageCapacity instantiates a new StorageCapacity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageCapacity() *StorageCapacity {
	this := StorageCapacity{}
	return &this
}

// NewStorageCapacityWithDefaults instantiates a new StorageCapacity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageCapacityWithDefaults() *StorageCapacity {
	this := StorageCapacity{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *StorageCapacity) GetAvailable() int64 {
	if o == nil || o.Available == nil {
		var ret int64
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageCapacity) GetAvailableOk() (*int64, bool) {
	if o == nil || o.Available == nil {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *StorageCapacity) HasAvailable() bool {
	if o != nil && o.Available != nil {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given int64 and assigns it to the Available field.
func (o *StorageCapacity) SetAvailable(v int64) {
	o.Available = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *StorageCapacity) GetFree() int64 {
	if o == nil || o.Free == nil {
		var ret int64
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageCapacity) GetFreeOk() (*int64, bool) {
	if o == nil || o.Free == nil {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *StorageCapacity) HasFree() bool {
	if o != nil && o.Free != nil {
		return true
	}

	return false
}

// SetFree gets a reference to the given int64 and assigns it to the Free field.
func (o *StorageCapacity) SetFree(v int64) {
	o.Free = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *StorageCapacity) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageCapacity) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *StorageCapacity) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *StorageCapacity) SetTotal(v int64) {
	o.Total = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *StorageCapacity) GetUsed() int64 {
	if o == nil || o.Used == nil {
		var ret int64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageCapacity) GetUsedOk() (*int64, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *StorageCapacity) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int64 and assigns it to the Used field.
func (o *StorageCapacity) SetUsed(v int64) {
	o.Used = &v
}

func (o StorageCapacity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Available != nil {
		toSerialize["Available"] = o.Available
	}
	if o.Free != nil {
		toSerialize["Free"] = o.Free
	}
	if o.Total != nil {
		toSerialize["Total"] = o.Total
	}
	if o.Used != nil {
		toSerialize["Used"] = o.Used
	}
	return json.Marshal(toSerialize)
}

type NullableStorageCapacity struct {
	value *StorageCapacity
	isSet bool
}

func (v NullableStorageCapacity) Get() *StorageCapacity {
	return v.value
}

func (v *NullableStorageCapacity) Set(val *StorageCapacity) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageCapacity) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageCapacity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageCapacity(val *StorageCapacity) *NullableStorageCapacity {
	return &NullableStorageCapacity{value: val, isSet: true}
}

func (v NullableStorageCapacity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageCapacity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
