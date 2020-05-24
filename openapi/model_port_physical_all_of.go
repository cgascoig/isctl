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

// PortPhysicalAllOf Definition of the list of properties defined in 'port.Physical', excluding properties defined in parent classes.
type PortPhysicalAllOf struct {
	OperState *string `json:"OperState,omitempty" yaml:"OperState,omitempty"`
	Role      *string `json:"Role,omitempty" yaml:"Role,omitempty"`
}

// NewPortPhysicalAllOf instantiates a new PortPhysicalAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortPhysicalAllOf() *PortPhysicalAllOf {
	this := PortPhysicalAllOf{}
	return &this
}

// NewPortPhysicalAllOfWithDefaults instantiates a new PortPhysicalAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortPhysicalAllOfWithDefaults() *PortPhysicalAllOf {
	this := PortPhysicalAllOf{}
	return &this
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *PortPhysicalAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortPhysicalAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *PortPhysicalAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *PortPhysicalAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PortPhysicalAllOf) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortPhysicalAllOf) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PortPhysicalAllOf) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *PortPhysicalAllOf) SetRole(v string) {
	o.Role = &v
}

func (o PortPhysicalAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.Role != nil {
		toSerialize["Role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullablePortPhysicalAllOf struct {
	value *PortPhysicalAllOf
	isSet bool
}

func (v NullablePortPhysicalAllOf) Get() *PortPhysicalAllOf {
	return v.value
}

func (v *NullablePortPhysicalAllOf) Set(val *PortPhysicalAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePortPhysicalAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePortPhysicalAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortPhysicalAllOf(val *PortPhysicalAllOf) *NullablePortPhysicalAllOf {
	return &NullablePortPhysicalAllOf{value: val, isSet: true}
}

func (v NullablePortPhysicalAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortPhysicalAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
