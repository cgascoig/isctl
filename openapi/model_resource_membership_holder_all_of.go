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

// ResourceMembershipHolderAllOf Definition of the list of properties defined in 'resource.MembershipHolder', excluding properties defined in parent classes.
type ResourceMembershipHolderAllOf struct {
	// The name of this resource membership holder.
	Name    *string                 `json:"Name,omitempty" yaml:"Name,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty" yaml:"Account,omitempty"`
	// An array of relationships to resourceMembership resources.
	Memberships *[]ResourceMembershipRelationship `json:"Memberships,omitempty" yaml:"Memberships,omitempty"`
}

// NewResourceMembershipHolderAllOf instantiates a new ResourceMembershipHolderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceMembershipHolderAllOf() *ResourceMembershipHolderAllOf {
	this := ResourceMembershipHolderAllOf{}
	return &this
}

// NewResourceMembershipHolderAllOfWithDefaults instantiates a new ResourceMembershipHolderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceMembershipHolderAllOfWithDefaults() *ResourceMembershipHolderAllOf {
	this := ResourceMembershipHolderAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceMembershipHolderAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceMembershipHolderAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceMembershipHolderAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceMembershipHolderAllOf) SetName(v string) {
	o.Name = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ResourceMembershipHolderAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceMembershipHolderAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ResourceMembershipHolderAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ResourceMembershipHolderAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetMemberships returns the Memberships field value if set, zero value otherwise.
func (o *ResourceMembershipHolderAllOf) GetMemberships() []ResourceMembershipRelationship {
	if o == nil || o.Memberships == nil {
		var ret []ResourceMembershipRelationship
		return ret
	}
	return *o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceMembershipHolderAllOf) GetMembershipsOk() (*[]ResourceMembershipRelationship, bool) {
	if o == nil || o.Memberships == nil {
		return nil, false
	}
	return o.Memberships, true
}

// HasMemberships returns a boolean if a field has been set.
func (o *ResourceMembershipHolderAllOf) HasMemberships() bool {
	if o != nil && o.Memberships != nil {
		return true
	}

	return false
}

// SetMemberships gets a reference to the given []ResourceMembershipRelationship and assigns it to the Memberships field.
func (o *ResourceMembershipHolderAllOf) SetMemberships(v []ResourceMembershipRelationship) {
	o.Memberships = &v
}

func (o ResourceMembershipHolderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Memberships != nil {
		toSerialize["Memberships"] = o.Memberships
	}
	return json.Marshal(toSerialize)
}

type NullableResourceMembershipHolderAllOf struct {
	value *ResourceMembershipHolderAllOf
	isSet bool
}

func (v NullableResourceMembershipHolderAllOf) Get() *ResourceMembershipHolderAllOf {
	return v.value
}

func (v *NullableResourceMembershipHolderAllOf) Set(val *ResourceMembershipHolderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceMembershipHolderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceMembershipHolderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceMembershipHolderAllOf(val *ResourceMembershipHolderAllOf) *NullableResourceMembershipHolderAllOf {
	return &NullableResourceMembershipHolderAllOf{value: val, isSet: true}
}

func (v NullableResourceMembershipHolderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceMembershipHolderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
