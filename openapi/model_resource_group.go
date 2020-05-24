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

// ResourceGroup A group of REST resources, such as a group of compute.Blade MOs. A ResourceGroup can contain static members which are specified as a set of object references, or it can contain dynamic members, which are specified through OData query filters. A Resource can be part of multiple resource groups.
type ResourceGroup struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The name of this resource group.
	Name                    *string                            `json:"Name,omitempty" yaml:"Name,omitempty"`
	PerTypeCombinedSelector *[]ResourcePerTypeCombinedSelector `json:"PerTypeCombinedSelector,omitempty" yaml:"PerTypeCombinedSelector,omitempty"`
	// Qualifier shall be used to specify if we want to organize resources using multiple resource group or single For an account, resource groups can be of only one of the above types. (Both the types are mutually exclusive for an account.).
	Qualifier *string                 `json:"Qualifier,omitempty" yaml:"Qualifier,omitempty"`
	Selectors *[]ResourceSelector     `json:"Selectors,omitempty" yaml:"Selectors,omitempty"`
	Account   *IamAccountRelationship `json:"Account,omitempty" yaml:"Account,omitempty"`
	// An array of relationships to organizationOrganization resources.
	Organizations *[]OrganizationOrganizationRelationship `json:"Organizations,omitempty" yaml:"Organizations,omitempty"`
}

// NewResourceGroup instantiates a new ResourceGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceGroup() *ResourceGroup {
	this := ResourceGroup{}
	var qualifier string = "Allow-Selectors"
	this.Qualifier = &qualifier
	return &this
}

// NewResourceGroupWithDefaults instantiates a new ResourceGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceGroupWithDefaults() *ResourceGroup {
	this := ResourceGroup{}
	var qualifier string = "Allow-Selectors"
	this.Qualifier = &qualifier
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceGroup) SetName(v string) {
	o.Name = &v
}

// GetPerTypeCombinedSelector returns the PerTypeCombinedSelector field value if set, zero value otherwise.
func (o *ResourceGroup) GetPerTypeCombinedSelector() []ResourcePerTypeCombinedSelector {
	if o == nil || o.PerTypeCombinedSelector == nil {
		var ret []ResourcePerTypeCombinedSelector
		return ret
	}
	return *o.PerTypeCombinedSelector
}

// GetPerTypeCombinedSelectorOk returns a tuple with the PerTypeCombinedSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroup) GetPerTypeCombinedSelectorOk() (*[]ResourcePerTypeCombinedSelector, bool) {
	if o == nil || o.PerTypeCombinedSelector == nil {
		return nil, false
	}
	return o.PerTypeCombinedSelector, true
}

// HasPerTypeCombinedSelector returns a boolean if a field has been set.
func (o *ResourceGroup) HasPerTypeCombinedSelector() bool {
	if o != nil && o.PerTypeCombinedSelector != nil {
		return true
	}

	return false
}

// SetPerTypeCombinedSelector gets a reference to the given []ResourcePerTypeCombinedSelector and assigns it to the PerTypeCombinedSelector field.
func (o *ResourceGroup) SetPerTypeCombinedSelector(v []ResourcePerTypeCombinedSelector) {
	o.PerTypeCombinedSelector = &v
}

// GetQualifier returns the Qualifier field value if set, zero value otherwise.
func (o *ResourceGroup) GetQualifier() string {
	if o == nil || o.Qualifier == nil {
		var ret string
		return ret
	}
	return *o.Qualifier
}

// GetQualifierOk returns a tuple with the Qualifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroup) GetQualifierOk() (*string, bool) {
	if o == nil || o.Qualifier == nil {
		return nil, false
	}
	return o.Qualifier, true
}

// HasQualifier returns a boolean if a field has been set.
func (o *ResourceGroup) HasQualifier() bool {
	if o != nil && o.Qualifier != nil {
		return true
	}

	return false
}

// SetQualifier gets a reference to the given string and assigns it to the Qualifier field.
func (o *ResourceGroup) SetQualifier(v string) {
	o.Qualifier = &v
}

// GetSelectors returns the Selectors field value if set, zero value otherwise.
func (o *ResourceGroup) GetSelectors() []ResourceSelector {
	if o == nil || o.Selectors == nil {
		var ret []ResourceSelector
		return ret
	}
	return *o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroup) GetSelectorsOk() (*[]ResourceSelector, bool) {
	if o == nil || o.Selectors == nil {
		return nil, false
	}
	return o.Selectors, true
}

// HasSelectors returns a boolean if a field has been set.
func (o *ResourceGroup) HasSelectors() bool {
	if o != nil && o.Selectors != nil {
		return true
	}

	return false
}

// SetSelectors gets a reference to the given []ResourceSelector and assigns it to the Selectors field.
func (o *ResourceGroup) SetSelectors(v []ResourceSelector) {
	o.Selectors = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ResourceGroup) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroup) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ResourceGroup) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ResourceGroup) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *ResourceGroup) GetOrganizations() []OrganizationOrganizationRelationship {
	if o == nil || o.Organizations == nil {
		var ret []OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroup) GetOrganizationsOk() (*[]OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organizations == nil {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *ResourceGroup) HasOrganizations() bool {
	if o != nil && o.Organizations != nil {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []OrganizationOrganizationRelationship and assigns it to the Organizations field.
func (o *ResourceGroup) SetOrganizations(v []OrganizationOrganizationRelationship) {
	o.Organizations = &v
}

func (o ResourceGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PerTypeCombinedSelector != nil {
		toSerialize["PerTypeCombinedSelector"] = o.PerTypeCombinedSelector
	}
	if o.Qualifier != nil {
		toSerialize["Qualifier"] = o.Qualifier
	}
	if o.Selectors != nil {
		toSerialize["Selectors"] = o.Selectors
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Organizations != nil {
		toSerialize["Organizations"] = o.Organizations
	}
	return json.Marshal(toSerialize)
}

type NullableResourceGroup struct {
	value *ResourceGroup
	isSet bool
}

func (v NullableResourceGroup) Get() *ResourceGroup {
	return v.value
}

func (v *NullableResourceGroup) Set(val *ResourceGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceGroup(val *ResourceGroup) *NullableResourceGroup {
	return &NullableResourceGroup{value: val, isSet: true}
}

func (v NullableResourceGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
