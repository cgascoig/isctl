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

// FabricVlan Configuration object for Virtual LAN.
type FabricVlan struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Used to determine whether this VLAN will be allowed on all uplink ports and PCs in this FI.
	AutoAllowOnUplinks *bool `json:"AutoAllowOnUplinks,omitempty" yaml:"AutoAllowOnUplinks,omitempty"`
	// Used to define whether this VLAN is to be classified as 'native' for traffic in this FI.
	IsNative *bool `json:"IsNative,omitempty" yaml:"IsNative,omitempty"`
	// The 'name' used to identify this VLAN.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// The identifier for this Virtual LAN.
	VlanId           *int64                              `json:"VlanId,omitempty" yaml:"VlanId,omitempty"`
	EthNetworkPolicy *FabricEthNetworkPolicyRelationship `json:"EthNetworkPolicy,omitempty" yaml:"EthNetworkPolicy,omitempty"`
	MulticastPolicy  *FabricMulticastPolicyRelationship  `json:"MulticastPolicy,omitempty" yaml:"MulticastPolicy,omitempty"`
}

// NewFabricVlan instantiates a new FabricVlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricVlan(classId string, objectType string) *FabricVlan {
	this := FabricVlan{}
	this.ClassId = classId
	this.ObjectType = objectType
	var autoAllowOnUplinks bool = true
	this.AutoAllowOnUplinks = &autoAllowOnUplinks
	return &this
}

// NewFabricVlanWithDefaults instantiates a new FabricVlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricVlanWithDefaults() *FabricVlan {
	this := FabricVlan{}
	var classId string = "fabric.Vlan"
	this.ClassId = classId
	var objectType string = "fabric.Vlan"
	this.ObjectType = objectType
	var autoAllowOnUplinks bool = true
	this.AutoAllowOnUplinks = &autoAllowOnUplinks
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricVlan) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricVlan) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricVlan) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricVlan) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAutoAllowOnUplinks returns the AutoAllowOnUplinks field value if set, zero value otherwise.
func (o *FabricVlan) GetAutoAllowOnUplinks() bool {
	if o == nil || o.AutoAllowOnUplinks == nil {
		var ret bool
		return ret
	}
	return *o.AutoAllowOnUplinks
}

// GetAutoAllowOnUplinksOk returns a tuple with the AutoAllowOnUplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetAutoAllowOnUplinksOk() (*bool, bool) {
	if o == nil || o.AutoAllowOnUplinks == nil {
		return nil, false
	}
	return o.AutoAllowOnUplinks, true
}

// HasAutoAllowOnUplinks returns a boolean if a field has been set.
func (o *FabricVlan) HasAutoAllowOnUplinks() bool {
	if o != nil && o.AutoAllowOnUplinks != nil {
		return true
	}

	return false
}

// SetAutoAllowOnUplinks gets a reference to the given bool and assigns it to the AutoAllowOnUplinks field.
func (o *FabricVlan) SetAutoAllowOnUplinks(v bool) {
	o.AutoAllowOnUplinks = &v
}

// GetIsNative returns the IsNative field value if set, zero value otherwise.
func (o *FabricVlan) GetIsNative() bool {
	if o == nil || o.IsNative == nil {
		var ret bool
		return ret
	}
	return *o.IsNative
}

// GetIsNativeOk returns a tuple with the IsNative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetIsNativeOk() (*bool, bool) {
	if o == nil || o.IsNative == nil {
		return nil, false
	}
	return o.IsNative, true
}

// HasIsNative returns a boolean if a field has been set.
func (o *FabricVlan) HasIsNative() bool {
	if o != nil && o.IsNative != nil {
		return true
	}

	return false
}

// SetIsNative gets a reference to the given bool and assigns it to the IsNative field.
func (o *FabricVlan) SetIsNative(v bool) {
	o.IsNative = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FabricVlan) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FabricVlan) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FabricVlan) SetName(v string) {
	o.Name = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *FabricVlan) GetVlanId() int64 {
	if o == nil || o.VlanId == nil {
		var ret int64
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetVlanIdOk() (*int64, bool) {
	if o == nil || o.VlanId == nil {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *FabricVlan) HasVlanId() bool {
	if o != nil && o.VlanId != nil {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int64 and assigns it to the VlanId field.
func (o *FabricVlan) SetVlanId(v int64) {
	o.VlanId = &v
}

// GetEthNetworkPolicy returns the EthNetworkPolicy field value if set, zero value otherwise.
func (o *FabricVlan) GetEthNetworkPolicy() FabricEthNetworkPolicyRelationship {
	if o == nil || o.EthNetworkPolicy == nil {
		var ret FabricEthNetworkPolicyRelationship
		return ret
	}
	return *o.EthNetworkPolicy
}

// GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetEthNetworkPolicyOk() (*FabricEthNetworkPolicyRelationship, bool) {
	if o == nil || o.EthNetworkPolicy == nil {
		return nil, false
	}
	return o.EthNetworkPolicy, true
}

// HasEthNetworkPolicy returns a boolean if a field has been set.
func (o *FabricVlan) HasEthNetworkPolicy() bool {
	if o != nil && o.EthNetworkPolicy != nil {
		return true
	}

	return false
}

// SetEthNetworkPolicy gets a reference to the given FabricEthNetworkPolicyRelationship and assigns it to the EthNetworkPolicy field.
func (o *FabricVlan) SetEthNetworkPolicy(v FabricEthNetworkPolicyRelationship) {
	o.EthNetworkPolicy = &v
}

// GetMulticastPolicy returns the MulticastPolicy field value if set, zero value otherwise.
func (o *FabricVlan) GetMulticastPolicy() FabricMulticastPolicyRelationship {
	if o == nil || o.MulticastPolicy == nil {
		var ret FabricMulticastPolicyRelationship
		return ret
	}
	return *o.MulticastPolicy
}

// GetMulticastPolicyOk returns a tuple with the MulticastPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetMulticastPolicyOk() (*FabricMulticastPolicyRelationship, bool) {
	if o == nil || o.MulticastPolicy == nil {
		return nil, false
	}
	return o.MulticastPolicy, true
}

// HasMulticastPolicy returns a boolean if a field has been set.
func (o *FabricVlan) HasMulticastPolicy() bool {
	if o != nil && o.MulticastPolicy != nil {
		return true
	}

	return false
}

// SetMulticastPolicy gets a reference to the given FabricMulticastPolicyRelationship and assigns it to the MulticastPolicy field.
func (o *FabricVlan) SetMulticastPolicy(v FabricMulticastPolicyRelationship) {
	o.MulticastPolicy = &v
}

func (o FabricVlan) MarshalJSON() ([]byte, error) {
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
	if o.AutoAllowOnUplinks != nil {
		toSerialize["AutoAllowOnUplinks"] = o.AutoAllowOnUplinks
	}
	if o.IsNative != nil {
		toSerialize["IsNative"] = o.IsNative
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.VlanId != nil {
		toSerialize["VlanId"] = o.VlanId
	}
	if o.EthNetworkPolicy != nil {
		toSerialize["EthNetworkPolicy"] = o.EthNetworkPolicy
	}
	if o.MulticastPolicy != nil {
		toSerialize["MulticastPolicy"] = o.MulticastPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableFabricVlan struct {
	value *FabricVlan
	isSet bool
}

func (v NullableFabricVlan) Get() *FabricVlan {
	return v.value
}

func (v *NullableFabricVlan) Set(val *FabricVlan) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricVlan) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricVlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricVlan(val *FabricVlan) *NullableFabricVlan {
	return &NullableFabricVlan{value: val, isSet: true}
}

func (v NullableFabricVlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricVlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
