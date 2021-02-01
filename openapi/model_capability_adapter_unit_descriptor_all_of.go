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

// CapabilityAdapterUnitDescriptorAllOf Definition of the list of properties defined in 'capability.AdapterUnitDescriptor', excluding properties defined in parent classes.
type CapabilityAdapterUnitDescriptorAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Order in which the ports are connected; sequential or cyclic.
	ConnectivityOrder *string `json:"ConnectivityOrder,omitempty" yaml:"ConnectivityOrder,omitempty"`
	// The port speed for ethernet ports in Mbps.
	EthernetPortSpeed *int64 `json:"EthernetPortSpeed,omitempty" yaml:"EthernetPortSpeed,omitempty"`
	// The port speed for fibre channel ports in Mbps.
	FibreChannelPortSpeed *int64 `json:"FibreChannelPortSpeed,omitempty" yaml:"FibreChannelPortSpeed,omitempty"`
	// Number of Dce Ports for the adaptor.
	NumDcePorts *int64 `json:"NumDcePorts,omitempty" yaml:"NumDcePorts,omitempty"`
	// Prom card type for the adaptor.
	PromCardType *string `json:"PromCardType,omitempty" yaml:"PromCardType,omitempty"`
}

// NewCapabilityAdapterUnitDescriptorAllOf instantiates a new CapabilityAdapterUnitDescriptorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityAdapterUnitDescriptorAllOf(classId string, objectType string) *CapabilityAdapterUnitDescriptorAllOf {
	this := CapabilityAdapterUnitDescriptorAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityAdapterUnitDescriptorAllOfWithDefaults instantiates a new CapabilityAdapterUnitDescriptorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityAdapterUnitDescriptorAllOfWithDefaults() *CapabilityAdapterUnitDescriptorAllOf {
	this := CapabilityAdapterUnitDescriptorAllOf{}
	var classId string = "capability.AdapterUnitDescriptor"
	this.ClassId = classId
	var objectType string = "capability.AdapterUnitDescriptor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityAdapterUnitDescriptorAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityAdapterUnitDescriptorAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityAdapterUnitDescriptorAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityAdapterUnitDescriptorAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectivityOrder returns the ConnectivityOrder field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetConnectivityOrder() string {
	if o == nil || o.ConnectivityOrder == nil {
		var ret string
		return ret
	}
	return *o.ConnectivityOrder
}

// GetConnectivityOrderOk returns a tuple with the ConnectivityOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetConnectivityOrderOk() (*string, bool) {
	if o == nil || o.ConnectivityOrder == nil {
		return nil, false
	}
	return o.ConnectivityOrder, true
}

// HasConnectivityOrder returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasConnectivityOrder() bool {
	if o != nil && o.ConnectivityOrder != nil {
		return true
	}

	return false
}

// SetConnectivityOrder gets a reference to the given string and assigns it to the ConnectivityOrder field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetConnectivityOrder(v string) {
	o.ConnectivityOrder = &v
}

// GetEthernetPortSpeed returns the EthernetPortSpeed field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetEthernetPortSpeed() int64 {
	if o == nil || o.EthernetPortSpeed == nil {
		var ret int64
		return ret
	}
	return *o.EthernetPortSpeed
}

// GetEthernetPortSpeedOk returns a tuple with the EthernetPortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetEthernetPortSpeedOk() (*int64, bool) {
	if o == nil || o.EthernetPortSpeed == nil {
		return nil, false
	}
	return o.EthernetPortSpeed, true
}

// HasEthernetPortSpeed returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasEthernetPortSpeed() bool {
	if o != nil && o.EthernetPortSpeed != nil {
		return true
	}

	return false
}

// SetEthernetPortSpeed gets a reference to the given int64 and assigns it to the EthernetPortSpeed field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetEthernetPortSpeed(v int64) {
	o.EthernetPortSpeed = &v
}

// GetFibreChannelPortSpeed returns the FibreChannelPortSpeed field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetFibreChannelPortSpeed() int64 {
	if o == nil || o.FibreChannelPortSpeed == nil {
		var ret int64
		return ret
	}
	return *o.FibreChannelPortSpeed
}

// GetFibreChannelPortSpeedOk returns a tuple with the FibreChannelPortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetFibreChannelPortSpeedOk() (*int64, bool) {
	if o == nil || o.FibreChannelPortSpeed == nil {
		return nil, false
	}
	return o.FibreChannelPortSpeed, true
}

// HasFibreChannelPortSpeed returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasFibreChannelPortSpeed() bool {
	if o != nil && o.FibreChannelPortSpeed != nil {
		return true
	}

	return false
}

// SetFibreChannelPortSpeed gets a reference to the given int64 and assigns it to the FibreChannelPortSpeed field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetFibreChannelPortSpeed(v int64) {
	o.FibreChannelPortSpeed = &v
}

// GetNumDcePorts returns the NumDcePorts field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetNumDcePorts() int64 {
	if o == nil || o.NumDcePorts == nil {
		var ret int64
		return ret
	}
	return *o.NumDcePorts
}

// GetNumDcePortsOk returns a tuple with the NumDcePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetNumDcePortsOk() (*int64, bool) {
	if o == nil || o.NumDcePorts == nil {
		return nil, false
	}
	return o.NumDcePorts, true
}

// HasNumDcePorts returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasNumDcePorts() bool {
	if o != nil && o.NumDcePorts != nil {
		return true
	}

	return false
}

// SetNumDcePorts gets a reference to the given int64 and assigns it to the NumDcePorts field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetNumDcePorts(v int64) {
	o.NumDcePorts = &v
}

// GetPromCardType returns the PromCardType field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetPromCardType() string {
	if o == nil || o.PromCardType == nil {
		var ret string
		return ret
	}
	return *o.PromCardType
}

// GetPromCardTypeOk returns a tuple with the PromCardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetPromCardTypeOk() (*string, bool) {
	if o == nil || o.PromCardType == nil {
		return nil, false
	}
	return o.PromCardType, true
}

// HasPromCardType returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasPromCardType() bool {
	if o != nil && o.PromCardType != nil {
		return true
	}

	return false
}

// SetPromCardType gets a reference to the given string and assigns it to the PromCardType field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetPromCardType(v string) {
	o.PromCardType = &v
}

func (o CapabilityAdapterUnitDescriptorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConnectivityOrder != nil {
		toSerialize["ConnectivityOrder"] = o.ConnectivityOrder
	}
	if o.EthernetPortSpeed != nil {
		toSerialize["EthernetPortSpeed"] = o.EthernetPortSpeed
	}
	if o.FibreChannelPortSpeed != nil {
		toSerialize["FibreChannelPortSpeed"] = o.FibreChannelPortSpeed
	}
	if o.NumDcePorts != nil {
		toSerialize["NumDcePorts"] = o.NumDcePorts
	}
	if o.PromCardType != nil {
		toSerialize["PromCardType"] = o.PromCardType
	}
	return json.Marshal(toSerialize)
}

type NullableCapabilityAdapterUnitDescriptorAllOf struct {
	value *CapabilityAdapterUnitDescriptorAllOf
	isSet bool
}

func (v NullableCapabilityAdapterUnitDescriptorAllOf) Get() *CapabilityAdapterUnitDescriptorAllOf {
	return v.value
}

func (v *NullableCapabilityAdapterUnitDescriptorAllOf) Set(val *CapabilityAdapterUnitDescriptorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityAdapterUnitDescriptorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityAdapterUnitDescriptorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityAdapterUnitDescriptorAllOf(val *CapabilityAdapterUnitDescriptorAllOf) *NullableCapabilityAdapterUnitDescriptorAllOf {
	return &NullableCapabilityAdapterUnitDescriptorAllOf{value: val, isSet: true}
}

func (v NullableCapabilityAdapterUnitDescriptorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityAdapterUnitDescriptorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
