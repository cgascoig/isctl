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

// HyperflexNodeConfigPolicyAllOf Definition of the list of properties defined in 'hyperflex.NodeConfigPolicy', excluding properties defined in parent classes.
type HyperflexNodeConfigPolicyAllOf struct {
	DataIpRange *HyperflexIpAddrRange `json:"DataIpRange,omitempty" yaml:"DataIpRange,omitempty"`
	HxdpIpRange *HyperflexIpAddrRange `json:"HxdpIpRange,omitempty" yaml:"HxdpIpRange,omitempty"`
	MgmtIpRange *HyperflexIpAddrRange `json:"MgmtIpRange,omitempty" yaml:"MgmtIpRange,omitempty"`
	// The node name prefix that is used to automatically generate the default hostname for each server. A dash (-) will be appended to the prefix followed by the node number to form a hostname. This default naming scheme can be manually overridden in the node configuration. The maximum length of a prefix is 60, must only contain alphanumeric characters or dash (-), and must start with an alphanumeric character.
	NodeNamePrefix *string `json:"NodeNamePrefix,omitempty" yaml:"NodeNamePrefix,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles *[]HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty" yaml:"ClusterProfiles,omitempty"`
	Organization    *OrganizationOrganizationRelationship  `json:"Organization,omitempty" yaml:"Organization,omitempty"`
}

// NewHyperflexNodeConfigPolicyAllOf instantiates a new HyperflexNodeConfigPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexNodeConfigPolicyAllOf() *HyperflexNodeConfigPolicyAllOf {
	this := HyperflexNodeConfigPolicyAllOf{}
	return &this
}

// NewHyperflexNodeConfigPolicyAllOfWithDefaults instantiates a new HyperflexNodeConfigPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexNodeConfigPolicyAllOfWithDefaults() *HyperflexNodeConfigPolicyAllOf {
	this := HyperflexNodeConfigPolicyAllOf{}
	return &this
}

// GetDataIpRange returns the DataIpRange field value if set, zero value otherwise.
func (o *HyperflexNodeConfigPolicyAllOf) GetDataIpRange() HyperflexIpAddrRange {
	if o == nil || o.DataIpRange == nil {
		var ret HyperflexIpAddrRange
		return ret
	}
	return *o.DataIpRange
}

// GetDataIpRangeOk returns a tuple with the DataIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeConfigPolicyAllOf) GetDataIpRangeOk() (*HyperflexIpAddrRange, bool) {
	if o == nil || o.DataIpRange == nil {
		return nil, false
	}
	return o.DataIpRange, true
}

// HasDataIpRange returns a boolean if a field has been set.
func (o *HyperflexNodeConfigPolicyAllOf) HasDataIpRange() bool {
	if o != nil && o.DataIpRange != nil {
		return true
	}

	return false
}

// SetDataIpRange gets a reference to the given HyperflexIpAddrRange and assigns it to the DataIpRange field.
func (o *HyperflexNodeConfigPolicyAllOf) SetDataIpRange(v HyperflexIpAddrRange) {
	o.DataIpRange = &v
}

// GetHxdpIpRange returns the HxdpIpRange field value if set, zero value otherwise.
func (o *HyperflexNodeConfigPolicyAllOf) GetHxdpIpRange() HyperflexIpAddrRange {
	if o == nil || o.HxdpIpRange == nil {
		var ret HyperflexIpAddrRange
		return ret
	}
	return *o.HxdpIpRange
}

// GetHxdpIpRangeOk returns a tuple with the HxdpIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeConfigPolicyAllOf) GetHxdpIpRangeOk() (*HyperflexIpAddrRange, bool) {
	if o == nil || o.HxdpIpRange == nil {
		return nil, false
	}
	return o.HxdpIpRange, true
}

// HasHxdpIpRange returns a boolean if a field has been set.
func (o *HyperflexNodeConfigPolicyAllOf) HasHxdpIpRange() bool {
	if o != nil && o.HxdpIpRange != nil {
		return true
	}

	return false
}

// SetHxdpIpRange gets a reference to the given HyperflexIpAddrRange and assigns it to the HxdpIpRange field.
func (o *HyperflexNodeConfigPolicyAllOf) SetHxdpIpRange(v HyperflexIpAddrRange) {
	o.HxdpIpRange = &v
}

// GetMgmtIpRange returns the MgmtIpRange field value if set, zero value otherwise.
func (o *HyperflexNodeConfigPolicyAllOf) GetMgmtIpRange() HyperflexIpAddrRange {
	if o == nil || o.MgmtIpRange == nil {
		var ret HyperflexIpAddrRange
		return ret
	}
	return *o.MgmtIpRange
}

// GetMgmtIpRangeOk returns a tuple with the MgmtIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeConfigPolicyAllOf) GetMgmtIpRangeOk() (*HyperflexIpAddrRange, bool) {
	if o == nil || o.MgmtIpRange == nil {
		return nil, false
	}
	return o.MgmtIpRange, true
}

// HasMgmtIpRange returns a boolean if a field has been set.
func (o *HyperflexNodeConfigPolicyAllOf) HasMgmtIpRange() bool {
	if o != nil && o.MgmtIpRange != nil {
		return true
	}

	return false
}

// SetMgmtIpRange gets a reference to the given HyperflexIpAddrRange and assigns it to the MgmtIpRange field.
func (o *HyperflexNodeConfigPolicyAllOf) SetMgmtIpRange(v HyperflexIpAddrRange) {
	o.MgmtIpRange = &v
}

// GetNodeNamePrefix returns the NodeNamePrefix field value if set, zero value otherwise.
func (o *HyperflexNodeConfigPolicyAllOf) GetNodeNamePrefix() string {
	if o == nil || o.NodeNamePrefix == nil {
		var ret string
		return ret
	}
	return *o.NodeNamePrefix
}

// GetNodeNamePrefixOk returns a tuple with the NodeNamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeConfigPolicyAllOf) GetNodeNamePrefixOk() (*string, bool) {
	if o == nil || o.NodeNamePrefix == nil {
		return nil, false
	}
	return o.NodeNamePrefix, true
}

// HasNodeNamePrefix returns a boolean if a field has been set.
func (o *HyperflexNodeConfigPolicyAllOf) HasNodeNamePrefix() bool {
	if o != nil && o.NodeNamePrefix != nil {
		return true
	}

	return false
}

// SetNodeNamePrefix gets a reference to the given string and assigns it to the NodeNamePrefix field.
func (o *HyperflexNodeConfigPolicyAllOf) SetNodeNamePrefix(v string) {
	o.NodeNamePrefix = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise.
func (o *HyperflexNodeConfigPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil || o.ClusterProfiles == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return *o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeConfigPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexNodeConfigPolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexNodeConfigPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexNodeConfigPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeConfigPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexNodeConfigPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexNodeConfigPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexNodeConfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DataIpRange != nil {
		toSerialize["DataIpRange"] = o.DataIpRange
	}
	if o.HxdpIpRange != nil {
		toSerialize["HxdpIpRange"] = o.HxdpIpRange
	}
	if o.MgmtIpRange != nil {
		toSerialize["MgmtIpRange"] = o.MgmtIpRange
	}
	if o.NodeNamePrefix != nil {
		toSerialize["NodeNamePrefix"] = o.NodeNamePrefix
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexNodeConfigPolicyAllOf struct {
	value *HyperflexNodeConfigPolicyAllOf
	isSet bool
}

func (v NullableHyperflexNodeConfigPolicyAllOf) Get() *HyperflexNodeConfigPolicyAllOf {
	return v.value
}

func (v *NullableHyperflexNodeConfigPolicyAllOf) Set(val *HyperflexNodeConfigPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexNodeConfigPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexNodeConfigPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexNodeConfigPolicyAllOf(val *HyperflexNodeConfigPolicyAllOf) *NullableHyperflexNodeConfigPolicyAllOf {
	return &NullableHyperflexNodeConfigPolicyAllOf{value: val, isSet: true}
}

func (v NullableHyperflexNodeConfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexNodeConfigPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
