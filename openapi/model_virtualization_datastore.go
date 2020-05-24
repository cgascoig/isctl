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

// VirtualizationDatastore Common attributes of a datastore allocated to a hypervisor. Serves as a base class for all concrete datastore types. A datastore is assigned to a datacenter and virtual machines use it for storage. The datastore could be backed by NFS, VMFS, etc.
type VirtualizationDatastore struct {
	VirtualizationSourceDevice `yaml:"VirtualizationSourceDevice,inline"`
	Capacity                   *VirtualizationStorageCapacity `json:"Capacity,omitempty" yaml:"Capacity,omitempty"`
	// Number of hosts attached to or supported-by this datastore.
	HostCount *int64 `json:"HostCount,omitempty" yaml:"HostCount,omitempty"`
	// The internally generated identity of this datastore. This entity is not manipulated by users. It aids in uniquely identifying the datastore object. For VMware, this is a MOR (managed object reference).
	Identity *string `json:"Identity,omitempty" yaml:"Identity,omitempty"`
	// Name of this datastore supplied by user. It is not the identity of the datastore. The name is subject to user manipulations.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// A string indicating the type of the datastore (VMFS, NFS, etc).
	Type *string `json:"Type,omitempty" yaml:"Type,omitempty"`
	// Number of virtual machines relying on (using) this datastore.
	VmCount *int64 `json:"VmCount,omitempty" yaml:"VmCount,omitempty"`
}

// NewVirtualizationDatastore instantiates a new VirtualizationDatastore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationDatastore() *VirtualizationDatastore {
	this := VirtualizationDatastore{}
	var type_ string = "Unknown"
	this.Type = &type_
	return &this
}

// NewVirtualizationDatastoreWithDefaults instantiates a new VirtualizationDatastore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationDatastoreWithDefaults() *VirtualizationDatastore {
	this := VirtualizationDatastore{}
	var type_ string = "Unknown"
	this.Type = &type_
	return &this
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *VirtualizationDatastore) GetCapacity() VirtualizationStorageCapacity {
	if o == nil || o.Capacity == nil {
		var ret VirtualizationStorageCapacity
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationDatastore) GetCapacityOk() (*VirtualizationStorageCapacity, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *VirtualizationDatastore) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given VirtualizationStorageCapacity and assigns it to the Capacity field.
func (o *VirtualizationDatastore) SetCapacity(v VirtualizationStorageCapacity) {
	o.Capacity = &v
}

// GetHostCount returns the HostCount field value if set, zero value otherwise.
func (o *VirtualizationDatastore) GetHostCount() int64 {
	if o == nil || o.HostCount == nil {
		var ret int64
		return ret
	}
	return *o.HostCount
}

// GetHostCountOk returns a tuple with the HostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationDatastore) GetHostCountOk() (*int64, bool) {
	if o == nil || o.HostCount == nil {
		return nil, false
	}
	return o.HostCount, true
}

// HasHostCount returns a boolean if a field has been set.
func (o *VirtualizationDatastore) HasHostCount() bool {
	if o != nil && o.HostCount != nil {
		return true
	}

	return false
}

// SetHostCount gets a reference to the given int64 and assigns it to the HostCount field.
func (o *VirtualizationDatastore) SetHostCount(v int64) {
	o.HostCount = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationDatastore) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationDatastore) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationDatastore) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationDatastore) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationDatastore) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationDatastore) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationDatastore) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationDatastore) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VirtualizationDatastore) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationDatastore) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VirtualizationDatastore) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VirtualizationDatastore) SetType(v string) {
	o.Type = &v
}

// GetVmCount returns the VmCount field value if set, zero value otherwise.
func (o *VirtualizationDatastore) GetVmCount() int64 {
	if o == nil || o.VmCount == nil {
		var ret int64
		return ret
	}
	return *o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationDatastore) GetVmCountOk() (*int64, bool) {
	if o == nil || o.VmCount == nil {
		return nil, false
	}
	return o.VmCount, true
}

// HasVmCount returns a boolean if a field has been set.
func (o *VirtualizationDatastore) HasVmCount() bool {
	if o != nil && o.VmCount != nil {
		return true
	}

	return false
}

// SetVmCount gets a reference to the given int64 and assigns it to the VmCount field.
func (o *VirtualizationDatastore) SetVmCount(v int64) {
	o.VmCount = &v
}

func (o VirtualizationDatastore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationSourceDevice, errVirtualizationSourceDevice := json.Marshal(o.VirtualizationSourceDevice)
	if errVirtualizationSourceDevice != nil {
		return []byte{}, errVirtualizationSourceDevice
	}
	errVirtualizationSourceDevice = json.Unmarshal([]byte(serializedVirtualizationSourceDevice), &toSerialize)
	if errVirtualizationSourceDevice != nil {
		return []byte{}, errVirtualizationSourceDevice
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.HostCount != nil {
		toSerialize["HostCount"] = o.HostCount
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VmCount != nil {
		toSerialize["VmCount"] = o.VmCount
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualizationDatastore struct {
	value *VirtualizationDatastore
	isSet bool
}

func (v NullableVirtualizationDatastore) Get() *VirtualizationDatastore {
	return v.value
}

func (v *NullableVirtualizationDatastore) Set(val *VirtualizationDatastore) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationDatastore) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationDatastore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationDatastore(val *VirtualizationDatastore) *NullableVirtualizationDatastore {
	return &NullableVirtualizationDatastore{value: val, isSet: true}
}

func (v NullableVirtualizationDatastore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationDatastore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
