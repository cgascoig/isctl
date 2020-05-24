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

// HyperflexExtIscsiStoragePolicy A policy specifying external storage connectivity information via Fabric attached FCoE storage.
type HyperflexExtIscsiStoragePolicy struct {
	PolicyAbstractPolicy `yaml:"PolicyAbstractPolicy,inline"`
	// Enable or disable external FCoE storage configuration.
	AdminState  *bool               `json:"AdminState,omitempty" yaml:"AdminState,omitempty"`
	ExtaTraffic *HyperflexNamedVlan `json:"ExtaTraffic,omitempty" yaml:"ExtaTraffic,omitempty"`
	ExtbTraffic *HyperflexNamedVlan `json:"ExtbTraffic,omitempty" yaml:"ExtbTraffic,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles *[]HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty" yaml:"ClusterProfiles,omitempty"`
	Organization    *OrganizationOrganizationRelationship  `json:"Organization,omitempty" yaml:"Organization,omitempty"`
}

// NewHyperflexExtIscsiStoragePolicy instantiates a new HyperflexExtIscsiStoragePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexExtIscsiStoragePolicy() *HyperflexExtIscsiStoragePolicy {
	this := HyperflexExtIscsiStoragePolicy{}
	return &this
}

// NewHyperflexExtIscsiStoragePolicyWithDefaults instantiates a new HyperflexExtIscsiStoragePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexExtIscsiStoragePolicyWithDefaults() *HyperflexExtIscsiStoragePolicy {
	this := HyperflexExtIscsiStoragePolicy{}
	return &this
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *HyperflexExtIscsiStoragePolicy) GetAdminState() bool {
	if o == nil || o.AdminState == nil {
		var ret bool
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexExtIscsiStoragePolicy) GetAdminStateOk() (*bool, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *HyperflexExtIscsiStoragePolicy) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given bool and assigns it to the AdminState field.
func (o *HyperflexExtIscsiStoragePolicy) SetAdminState(v bool) {
	o.AdminState = &v
}

// GetExtaTraffic returns the ExtaTraffic field value if set, zero value otherwise.
func (o *HyperflexExtIscsiStoragePolicy) GetExtaTraffic() HyperflexNamedVlan {
	if o == nil || o.ExtaTraffic == nil {
		var ret HyperflexNamedVlan
		return ret
	}
	return *o.ExtaTraffic
}

// GetExtaTrafficOk returns a tuple with the ExtaTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexExtIscsiStoragePolicy) GetExtaTrafficOk() (*HyperflexNamedVlan, bool) {
	if o == nil || o.ExtaTraffic == nil {
		return nil, false
	}
	return o.ExtaTraffic, true
}

// HasExtaTraffic returns a boolean if a field has been set.
func (o *HyperflexExtIscsiStoragePolicy) HasExtaTraffic() bool {
	if o != nil && o.ExtaTraffic != nil {
		return true
	}

	return false
}

// SetExtaTraffic gets a reference to the given HyperflexNamedVlan and assigns it to the ExtaTraffic field.
func (o *HyperflexExtIscsiStoragePolicy) SetExtaTraffic(v HyperflexNamedVlan) {
	o.ExtaTraffic = &v
}

// GetExtbTraffic returns the ExtbTraffic field value if set, zero value otherwise.
func (o *HyperflexExtIscsiStoragePolicy) GetExtbTraffic() HyperflexNamedVlan {
	if o == nil || o.ExtbTraffic == nil {
		var ret HyperflexNamedVlan
		return ret
	}
	return *o.ExtbTraffic
}

// GetExtbTrafficOk returns a tuple with the ExtbTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexExtIscsiStoragePolicy) GetExtbTrafficOk() (*HyperflexNamedVlan, bool) {
	if o == nil || o.ExtbTraffic == nil {
		return nil, false
	}
	return o.ExtbTraffic, true
}

// HasExtbTraffic returns a boolean if a field has been set.
func (o *HyperflexExtIscsiStoragePolicy) HasExtbTraffic() bool {
	if o != nil && o.ExtbTraffic != nil {
		return true
	}

	return false
}

// SetExtbTraffic gets a reference to the given HyperflexNamedVlan and assigns it to the ExtbTraffic field.
func (o *HyperflexExtIscsiStoragePolicy) SetExtbTraffic(v HyperflexNamedVlan) {
	o.ExtbTraffic = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise.
func (o *HyperflexExtIscsiStoragePolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil || o.ClusterProfiles == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return *o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexExtIscsiStoragePolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexExtIscsiStoragePolicy) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexExtIscsiStoragePolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexExtIscsiStoragePolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexExtIscsiStoragePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexExtIscsiStoragePolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexExtIscsiStoragePolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexExtIscsiStoragePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.ExtaTraffic != nil {
		toSerialize["ExtaTraffic"] = o.ExtaTraffic
	}
	if o.ExtbTraffic != nil {
		toSerialize["ExtbTraffic"] = o.ExtbTraffic
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexExtIscsiStoragePolicy struct {
	value *HyperflexExtIscsiStoragePolicy
	isSet bool
}

func (v NullableHyperflexExtIscsiStoragePolicy) Get() *HyperflexExtIscsiStoragePolicy {
	return v.value
}

func (v *NullableHyperflexExtIscsiStoragePolicy) Set(val *HyperflexExtIscsiStoragePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexExtIscsiStoragePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexExtIscsiStoragePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexExtIscsiStoragePolicy(val *HyperflexExtIscsiStoragePolicy) *NullableHyperflexExtIscsiStoragePolicy {
	return &NullableHyperflexExtIscsiStoragePolicy{value: val, isSet: true}
}

func (v NullableHyperflexExtIscsiStoragePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexExtIscsiStoragePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
