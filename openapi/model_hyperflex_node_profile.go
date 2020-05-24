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

// HyperflexNodeProfile A configuration profile per node in the HyperFlex cluster. It defines node settings such as IP address configuration for hypervisor management network, storage data network, HyperFlex management network, and the assigned physical server.
type HyperflexNodeProfile struct {
	PolicyAbstractProfile `yaml:"PolicyAbstractProfile,inline"`
	// IP address for storage data network (Controller VM interface).
	HxdpDataIp *string `json:"HxdpDataIp,omitempty" yaml:"HxdpDataIp,omitempty"`
	// IP address for HyperFlex management network.
	HxdpMgmtIp *string `json:"HxdpMgmtIp,omitempty" yaml:"HxdpMgmtIp,omitempty"`
	// IP address for storage data network (Hypervisor interface).
	HypervisorDataIp *string `json:"HypervisorDataIp,omitempty" yaml:"HypervisorDataIp,omitempty"`
	// IP address for Hypervisor management network.
	HypervisorMgmtIp *string                              `json:"HypervisorMgmtIp,omitempty" yaml:"HypervisorMgmtIp,omitempty"`
	AssignedServer   *ComputeRackUnitRelationship         `json:"AssignedServer,omitempty" yaml:"AssignedServer,omitempty"`
	ClusterProfile   *HyperflexClusterProfileRelationship `json:"ClusterProfile,omitempty" yaml:"ClusterProfile,omitempty"`
}

// NewHyperflexNodeProfile instantiates a new HyperflexNodeProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexNodeProfile() *HyperflexNodeProfile {
	this := HyperflexNodeProfile{}
	return &this
}

// NewHyperflexNodeProfileWithDefaults instantiates a new HyperflexNodeProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexNodeProfileWithDefaults() *HyperflexNodeProfile {
	this := HyperflexNodeProfile{}
	return &this
}

// GetHxdpDataIp returns the HxdpDataIp field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetHxdpDataIp() string {
	if o == nil || o.HxdpDataIp == nil {
		var ret string
		return ret
	}
	return *o.HxdpDataIp
}

// GetHxdpDataIpOk returns a tuple with the HxdpDataIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetHxdpDataIpOk() (*string, bool) {
	if o == nil || o.HxdpDataIp == nil {
		return nil, false
	}
	return o.HxdpDataIp, true
}

// HasHxdpDataIp returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasHxdpDataIp() bool {
	if o != nil && o.HxdpDataIp != nil {
		return true
	}

	return false
}

// SetHxdpDataIp gets a reference to the given string and assigns it to the HxdpDataIp field.
func (o *HyperflexNodeProfile) SetHxdpDataIp(v string) {
	o.HxdpDataIp = &v
}

// GetHxdpMgmtIp returns the HxdpMgmtIp field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetHxdpMgmtIp() string {
	if o == nil || o.HxdpMgmtIp == nil {
		var ret string
		return ret
	}
	return *o.HxdpMgmtIp
}

// GetHxdpMgmtIpOk returns a tuple with the HxdpMgmtIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetHxdpMgmtIpOk() (*string, bool) {
	if o == nil || o.HxdpMgmtIp == nil {
		return nil, false
	}
	return o.HxdpMgmtIp, true
}

// HasHxdpMgmtIp returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasHxdpMgmtIp() bool {
	if o != nil && o.HxdpMgmtIp != nil {
		return true
	}

	return false
}

// SetHxdpMgmtIp gets a reference to the given string and assigns it to the HxdpMgmtIp field.
func (o *HyperflexNodeProfile) SetHxdpMgmtIp(v string) {
	o.HxdpMgmtIp = &v
}

// GetHypervisorDataIp returns the HypervisorDataIp field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetHypervisorDataIp() string {
	if o == nil || o.HypervisorDataIp == nil {
		var ret string
		return ret
	}
	return *o.HypervisorDataIp
}

// GetHypervisorDataIpOk returns a tuple with the HypervisorDataIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetHypervisorDataIpOk() (*string, bool) {
	if o == nil || o.HypervisorDataIp == nil {
		return nil, false
	}
	return o.HypervisorDataIp, true
}

// HasHypervisorDataIp returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasHypervisorDataIp() bool {
	if o != nil && o.HypervisorDataIp != nil {
		return true
	}

	return false
}

// SetHypervisorDataIp gets a reference to the given string and assigns it to the HypervisorDataIp field.
func (o *HyperflexNodeProfile) SetHypervisorDataIp(v string) {
	o.HypervisorDataIp = &v
}

// GetHypervisorMgmtIp returns the HypervisorMgmtIp field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetHypervisorMgmtIp() string {
	if o == nil || o.HypervisorMgmtIp == nil {
		var ret string
		return ret
	}
	return *o.HypervisorMgmtIp
}

// GetHypervisorMgmtIpOk returns a tuple with the HypervisorMgmtIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetHypervisorMgmtIpOk() (*string, bool) {
	if o == nil || o.HypervisorMgmtIp == nil {
		return nil, false
	}
	return o.HypervisorMgmtIp, true
}

// HasHypervisorMgmtIp returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasHypervisorMgmtIp() bool {
	if o != nil && o.HypervisorMgmtIp != nil {
		return true
	}

	return false
}

// SetHypervisorMgmtIp gets a reference to the given string and assigns it to the HypervisorMgmtIp field.
func (o *HyperflexNodeProfile) SetHypervisorMgmtIp(v string) {
	o.HypervisorMgmtIp = &v
}

// GetAssignedServer returns the AssignedServer field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetAssignedServer() ComputeRackUnitRelationship {
	if o == nil || o.AssignedServer == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.AssignedServer
}

// GetAssignedServerOk returns a tuple with the AssignedServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetAssignedServerOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.AssignedServer == nil {
		return nil, false
	}
	return o.AssignedServer, true
}

// HasAssignedServer returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasAssignedServer() bool {
	if o != nil && o.AssignedServer != nil {
		return true
	}

	return false
}

// SetAssignedServer gets a reference to the given ComputeRackUnitRelationship and assigns it to the AssignedServer field.
func (o *HyperflexNodeProfile) SetAssignedServer(v ComputeRackUnitRelationship) {
	o.AssignedServer = &v
}

// GetClusterProfile returns the ClusterProfile field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetClusterProfile() HyperflexClusterProfileRelationship {
	if o == nil || o.ClusterProfile == nil {
		var ret HyperflexClusterProfileRelationship
		return ret
	}
	return *o.ClusterProfile
}

// GetClusterProfileOk returns a tuple with the ClusterProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetClusterProfileOk() (*HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfile == nil {
		return nil, false
	}
	return o.ClusterProfile, true
}

// HasClusterProfile returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasClusterProfile() bool {
	if o != nil && o.ClusterProfile != nil {
		return true
	}

	return false
}

// SetClusterProfile gets a reference to the given HyperflexClusterProfileRelationship and assigns it to the ClusterProfile field.
func (o *HyperflexNodeProfile) SetClusterProfile(v HyperflexClusterProfileRelationship) {
	o.ClusterProfile = &v
}

func (o HyperflexNodeProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractProfile, errPolicyAbstractProfile := json.Marshal(o.PolicyAbstractProfile)
	if errPolicyAbstractProfile != nil {
		return []byte{}, errPolicyAbstractProfile
	}
	errPolicyAbstractProfile = json.Unmarshal([]byte(serializedPolicyAbstractProfile), &toSerialize)
	if errPolicyAbstractProfile != nil {
		return []byte{}, errPolicyAbstractProfile
	}
	if o.HxdpDataIp != nil {
		toSerialize["HxdpDataIp"] = o.HxdpDataIp
	}
	if o.HxdpMgmtIp != nil {
		toSerialize["HxdpMgmtIp"] = o.HxdpMgmtIp
	}
	if o.HypervisorDataIp != nil {
		toSerialize["HypervisorDataIp"] = o.HypervisorDataIp
	}
	if o.HypervisorMgmtIp != nil {
		toSerialize["HypervisorMgmtIp"] = o.HypervisorMgmtIp
	}
	if o.AssignedServer != nil {
		toSerialize["AssignedServer"] = o.AssignedServer
	}
	if o.ClusterProfile != nil {
		toSerialize["ClusterProfile"] = o.ClusterProfile
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexNodeProfile struct {
	value *HyperflexNodeProfile
	isSet bool
}

func (v NullableHyperflexNodeProfile) Get() *HyperflexNodeProfile {
	return v.value
}

func (v *NullableHyperflexNodeProfile) Set(val *HyperflexNodeProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexNodeProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexNodeProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexNodeProfile(val *HyperflexNodeProfile) *NullableHyperflexNodeProfile {
	return &NullableHyperflexNodeProfile{value: val, isSet: true}
}

func (v NullableHyperflexNodeProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexNodeProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
