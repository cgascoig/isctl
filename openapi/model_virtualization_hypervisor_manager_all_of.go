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

// VirtualizationHypervisorManagerAllOf Definition of the list of properties defined in 'virtualization.HypervisorManager', excluding properties defined in parent classes.
type VirtualizationHypervisorManagerAllOf struct {
	// Identity of the hypervisor (not manipulated by user). It could be a UUID too. For example, c917093f-5443-4748-bc09-eec72ded7608.
	Identity *string `json:"Identity,omitempty" yaml:"Identity,omitempty"`
	// The user provided name for the hypervisor manager. For example, vCenterIreland. Usually, this name is subject to manipulation by the user. It is not the identity of the hypervisor.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// Release version of the Hypervisor Manger (VMware vCenter Server 6.0.0 build-4541947).
	Version          *string                              `json:"Version,omitempty" yaml:"Version,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewVirtualizationHypervisorManagerAllOf instantiates a new VirtualizationHypervisorManagerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationHypervisorManagerAllOf() *VirtualizationHypervisorManagerAllOf {
	this := VirtualizationHypervisorManagerAllOf{}
	return &this
}

// NewVirtualizationHypervisorManagerAllOfWithDefaults instantiates a new VirtualizationHypervisorManagerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationHypervisorManagerAllOfWithDefaults() *VirtualizationHypervisorManagerAllOf {
	this := VirtualizationHypervisorManagerAllOf{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationHypervisorManagerAllOf) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHypervisorManagerAllOf) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationHypervisorManagerAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationHypervisorManagerAllOf) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationHypervisorManagerAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHypervisorManagerAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationHypervisorManagerAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationHypervisorManagerAllOf) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VirtualizationHypervisorManagerAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHypervisorManagerAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VirtualizationHypervisorManagerAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *VirtualizationHypervisorManagerAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *VirtualizationHypervisorManagerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHypervisorManagerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *VirtualizationHypervisorManagerAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *VirtualizationHypervisorManagerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o VirtualizationHypervisorManagerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualizationHypervisorManagerAllOf struct {
	value *VirtualizationHypervisorManagerAllOf
	isSet bool
}

func (v NullableVirtualizationHypervisorManagerAllOf) Get() *VirtualizationHypervisorManagerAllOf {
	return v.value
}

func (v *NullableVirtualizationHypervisorManagerAllOf) Set(val *VirtualizationHypervisorManagerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationHypervisorManagerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationHypervisorManagerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationHypervisorManagerAllOf(val *VirtualizationHypervisorManagerAllOf) *NullableVirtualizationHypervisorManagerAllOf {
	return &NullableVirtualizationHypervisorManagerAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationHypervisorManagerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationHypervisorManagerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
