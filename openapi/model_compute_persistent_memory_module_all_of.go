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

// ComputePersistentMemoryModuleAllOf Definition of the list of properties defined in 'compute.PersistentMemoryModule', excluding properties defined in parent classes.
type ComputePersistentMemoryModuleAllOf struct {
	// Socket ID of the Persistent Memory Module on the server.
	SocketId *string `json:"SocketId,omitempty" yaml:"SocketId,omitempty"`
	// Socket Memory ID of the Persistent Memory Module on the server.
	SocketMemoryId *string `json:"SocketMemoryId,omitempty" yaml:"SocketMemoryId,omitempty"`
}

// NewComputePersistentMemoryModuleAllOf instantiates a new ComputePersistentMemoryModuleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputePersistentMemoryModuleAllOf() *ComputePersistentMemoryModuleAllOf {
	this := ComputePersistentMemoryModuleAllOf{}
	return &this
}

// NewComputePersistentMemoryModuleAllOfWithDefaults instantiates a new ComputePersistentMemoryModuleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputePersistentMemoryModuleAllOfWithDefaults() *ComputePersistentMemoryModuleAllOf {
	this := ComputePersistentMemoryModuleAllOf{}
	return &this
}

// GetSocketId returns the SocketId field value if set, zero value otherwise.
func (o *ComputePersistentMemoryModuleAllOf) GetSocketId() string {
	if o == nil || o.SocketId == nil {
		var ret string
		return ret
	}
	return *o.SocketId
}

// GetSocketIdOk returns a tuple with the SocketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryModuleAllOf) GetSocketIdOk() (*string, bool) {
	if o == nil || o.SocketId == nil {
		return nil, false
	}
	return o.SocketId, true
}

// HasSocketId returns a boolean if a field has been set.
func (o *ComputePersistentMemoryModuleAllOf) HasSocketId() bool {
	if o != nil && o.SocketId != nil {
		return true
	}

	return false
}

// SetSocketId gets a reference to the given string and assigns it to the SocketId field.
func (o *ComputePersistentMemoryModuleAllOf) SetSocketId(v string) {
	o.SocketId = &v
}

// GetSocketMemoryId returns the SocketMemoryId field value if set, zero value otherwise.
func (o *ComputePersistentMemoryModuleAllOf) GetSocketMemoryId() string {
	if o == nil || o.SocketMemoryId == nil {
		var ret string
		return ret
	}
	return *o.SocketMemoryId
}

// GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryModuleAllOf) GetSocketMemoryIdOk() (*string, bool) {
	if o == nil || o.SocketMemoryId == nil {
		return nil, false
	}
	return o.SocketMemoryId, true
}

// HasSocketMemoryId returns a boolean if a field has been set.
func (o *ComputePersistentMemoryModuleAllOf) HasSocketMemoryId() bool {
	if o != nil && o.SocketMemoryId != nil {
		return true
	}

	return false
}

// SetSocketMemoryId gets a reference to the given string and assigns it to the SocketMemoryId field.
func (o *ComputePersistentMemoryModuleAllOf) SetSocketMemoryId(v string) {
	o.SocketMemoryId = &v
}

func (o ComputePersistentMemoryModuleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SocketId != nil {
		toSerialize["SocketId"] = o.SocketId
	}
	if o.SocketMemoryId != nil {
		toSerialize["SocketMemoryId"] = o.SocketMemoryId
	}
	return json.Marshal(toSerialize)
}

type NullableComputePersistentMemoryModuleAllOf struct {
	value *ComputePersistentMemoryModuleAllOf
	isSet bool
}

func (v NullableComputePersistentMemoryModuleAllOf) Get() *ComputePersistentMemoryModuleAllOf {
	return v.value
}

func (v *NullableComputePersistentMemoryModuleAllOf) Set(val *ComputePersistentMemoryModuleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputePersistentMemoryModuleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputePersistentMemoryModuleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputePersistentMemoryModuleAllOf(val *ComputePersistentMemoryModuleAllOf) *NullableComputePersistentMemoryModuleAllOf {
	return &NullableComputePersistentMemoryModuleAllOf{value: val, isSet: true}
}

func (v NullableComputePersistentMemoryModuleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputePersistentMemoryModuleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
