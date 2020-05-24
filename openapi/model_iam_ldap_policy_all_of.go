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

// IamLdapPolicyAllOf Definition of the list of properties defined in 'iam.LdapPolicy', excluding properties defined in parent classes.
type IamLdapPolicyAllOf struct {
	BaseProperties *IamLdapBaseProperties `json:"BaseProperties,omitempty" yaml:"BaseProperties,omitempty"`
	DnsParameters  *IamLdapDnsParameters  `json:"DnsParameters,omitempty" yaml:"DnsParameters,omitempty"`
	// Enables DNS to access LDAP servers.
	EnableDns *bool `json:"EnableDns,omitempty" yaml:"EnableDns,omitempty"`
	// LDAP server performs authentication.
	Enabled *bool `json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	// Search precedence between local user database and LDAP user database.
	UserSearchPrecedence *string                 `json:"UserSearchPrecedence,omitempty" yaml:"UserSearchPrecedence,omitempty"`
	Var0Idp              *IamIdpRelationship     `json:"_0_Idp,omitempty" yaml:"_0_Idp,omitempty"`
	ApplianceAccount     *IamAccountRelationship `json:"ApplianceAccount,omitempty" yaml:"ApplianceAccount,omitempty"`
	// An array of relationships to iamLdapGroup resources.
	Groups       *[]IamLdapGroupRelationship           `json:"Groups,omitempty" yaml:"Groups,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty" yaml:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles *[]PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty" yaml:"Profiles,omitempty"`
	// An array of relationships to iamLdapProvider resources.
	Providers *[]IamLdapProviderRelationship `json:"Providers,omitempty" yaml:"Providers,omitempty"`
}

// NewIamLdapPolicyAllOf instantiates a new IamLdapPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamLdapPolicyAllOf() *IamLdapPolicyAllOf {
	this := IamLdapPolicyAllOf{}
	var userSearchPrecedence string = "LocalUserDb"
	this.UserSearchPrecedence = &userSearchPrecedence
	return &this
}

// NewIamLdapPolicyAllOfWithDefaults instantiates a new IamLdapPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamLdapPolicyAllOfWithDefaults() *IamLdapPolicyAllOf {
	this := IamLdapPolicyAllOf{}
	var userSearchPrecedence string = "LocalUserDb"
	this.UserSearchPrecedence = &userSearchPrecedence
	return &this
}

// GetBaseProperties returns the BaseProperties field value if set, zero value otherwise.
func (o *IamLdapPolicyAllOf) GetBaseProperties() IamLdapBaseProperties {
	if o == nil || o.BaseProperties == nil {
		var ret IamLdapBaseProperties
		return ret
	}
	return *o.BaseProperties
}

// GetBasePropertiesOk returns a tuple with the BaseProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapPolicyAllOf) GetBasePropertiesOk() (*IamLdapBaseProperties, bool) {
	if o == nil || o.BaseProperties == nil {
		return nil, false
	}
	return o.BaseProperties, true
}

// HasBaseProperties returns a boolean if a field has been set.
func (o *IamLdapPolicyAllOf) HasBaseProperties() bool {
	if o != nil && o.BaseProperties != nil {
		return true
	}

	return false
}

// SetBaseProperties gets a reference to the given IamLdapBaseProperties and assigns it to the BaseProperties field.
func (o *IamLdapPolicyAllOf) SetBaseProperties(v IamLdapBaseProperties) {
	o.BaseProperties = &v
}

// GetDnsParameters returns the DnsParameters field value if set, zero value otherwise.
func (o *IamLdapPolicyAllOf) GetDnsParameters() IamLdapDnsParameters {
	if o == nil || o.DnsParameters == nil {
		var ret IamLdapDnsParameters
		return ret
	}
	return *o.DnsParameters
}

// GetDnsParametersOk returns a tuple with the DnsParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapPolicyAllOf) GetDnsParametersOk() (*IamLdapDnsParameters, bool) {
	if o == nil || o.DnsParameters == nil {
		return nil, false
	}
	return o.DnsParameters, true
}

// HasDnsParameters returns a boolean if a field has been set.
func (o *IamLdapPolicyAllOf) HasDnsParameters() bool {
	if o != nil && o.DnsParameters != nil {
		return true
	}

	return false
}

// SetDnsParameters gets a reference to the given IamLdapDnsParameters and assigns it to the DnsParameters field.
func (o *IamLdapPolicyAllOf) SetDnsParameters(v IamLdapDnsParameters) {
	o.DnsParameters = &v
}

// GetEnableDns returns the EnableDns field value if set, zero value otherwise.
func (o *IamLdapPolicyAllOf) GetEnableDns() bool {
	if o == nil || o.EnableDns == nil {
		var ret bool
		return ret
	}
	return *o.EnableDns
}

// GetEnableDnsOk returns a tuple with the EnableDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapPolicyAllOf) GetEnableDnsOk() (*bool, bool) {
	if o == nil || o.EnableDns == nil {
		return nil, false
	}
	return o.EnableDns, true
}

// HasEnableDns returns a boolean if a field has been set.
func (o *IamLdapPolicyAllOf) HasEnableDns() bool {
	if o != nil && o.EnableDns != nil {
		return true
	}

	return false
}

// SetEnableDns gets a reference to the given bool and assigns it to the EnableDns field.
func (o *IamLdapPolicyAllOf) SetEnableDns(v bool) {
	o.EnableDns = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IamLdapPolicyAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapPolicyAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IamLdapPolicyAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IamLdapPolicyAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUserSearchPrecedence returns the UserSearchPrecedence field value if set, zero value otherwise.
func (o *IamLdapPolicyAllOf) GetUserSearchPrecedence() string {
	if o == nil || o.UserSearchPrecedence == nil {
		var ret string
		return ret
	}
	return *o.UserSearchPrecedence
}

// GetUserSearchPrecedenceOk returns a tuple with the UserSearchPrecedence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapPolicyAllOf) GetUserSearchPrecedenceOk() (*string, bool) {
	if o == nil || o.UserSearchPrecedence == nil {
		return nil, false
	}
	return o.UserSearchPrecedence, true
}

// HasUserSearchPrecedence returns a boolean if a field has been set.
func (o *IamLdapPolicyAllOf) HasUserSearchPrecedence() bool {
	if o != nil && o.UserSearchPrecedence != nil {
		return true
	}

	return false
}

// SetUserSearchPrecedence gets a reference to the given string and assigns it to the UserSearchPrecedence field.
func (o *IamLdapPolicyAllOf) SetUserSearchPrecedence(v string) {
	o.UserSearchPrecedence = &v
}

// GetVar0Idp returns the Var0Idp field value if set, zero value otherwise.
func (o *IamLdapPolicyAllOf) GetVar0Idp() IamIdpRelationship {
	if o == nil || o.Var0Idp == nil {
		var ret IamIdpRelationship
		return ret
	}
	return *o.Var0Idp
}

// GetVar0IdpOk returns a tuple with the Var0Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapPolicyAllOf) GetVar0IdpOk() (*IamIdpRelationship, bool) {
	if o == nil || o.Var0Idp == nil {
		return nil, false
	}
	return o.Var0Idp, true
}

// HasVar0Idp returns a boolean if a field has been set.
func (o *IamLdapPolicyAllOf) HasVar0Idp() bool {
	if o != nil && o.Var0Idp != nil {
		return true
	}

	return false
}

// SetVar0Idp gets a reference to the given IamIdpRelationship and assigns it to the Var0Idp field.
func (o *IamLdapPolicyAllOf) SetVar0Idp(v IamIdpRelationship) {
	o.Var0Idp = &v
}

// GetApplianceAccount returns the ApplianceAccount field value if set, zero value otherwise.
func (o *IamLdapPolicyAllOf) GetApplianceAccount() IamAccountRelationship {
	if o == nil || o.ApplianceAccount == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.ApplianceAccount
}

// GetApplianceAccountOk returns a tuple with the ApplianceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapPolicyAllOf) GetApplianceAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.ApplianceAccount == nil {
		return nil, false
	}
	return o.ApplianceAccount, true
}

// HasApplianceAccount returns a boolean if a field has been set.
func (o *IamLdapPolicyAllOf) HasApplianceAccount() bool {
	if o != nil && o.ApplianceAccount != nil {
		return true
	}

	return false
}

// SetApplianceAccount gets a reference to the given IamAccountRelationship and assigns it to the ApplianceAccount field.
func (o *IamLdapPolicyAllOf) SetApplianceAccount(v IamAccountRelationship) {
	o.ApplianceAccount = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *IamLdapPolicyAllOf) GetGroups() []IamLdapGroupRelationship {
	if o == nil || o.Groups == nil {
		var ret []IamLdapGroupRelationship
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapPolicyAllOf) GetGroupsOk() (*[]IamLdapGroupRelationship, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *IamLdapPolicyAllOf) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []IamLdapGroupRelationship and assigns it to the Groups field.
func (o *IamLdapPolicyAllOf) SetGroups(v []IamLdapGroupRelationship) {
	o.Groups = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *IamLdapPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *IamLdapPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *IamLdapPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *IamLdapPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil || o.Profiles == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return *o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *IamLdapPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *IamLdapPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *IamLdapPolicyAllOf) GetProviders() []IamLdapProviderRelationship {
	if o == nil || o.Providers == nil {
		var ret []IamLdapProviderRelationship
		return ret
	}
	return *o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapPolicyAllOf) GetProvidersOk() (*[]IamLdapProviderRelationship, bool) {
	if o == nil || o.Providers == nil {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *IamLdapPolicyAllOf) HasProviders() bool {
	if o != nil && o.Providers != nil {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []IamLdapProviderRelationship and assigns it to the Providers field.
func (o *IamLdapPolicyAllOf) SetProviders(v []IamLdapProviderRelationship) {
	o.Providers = &v
}

func (o IamLdapPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseProperties != nil {
		toSerialize["BaseProperties"] = o.BaseProperties
	}
	if o.DnsParameters != nil {
		toSerialize["DnsParameters"] = o.DnsParameters
	}
	if o.EnableDns != nil {
		toSerialize["EnableDns"] = o.EnableDns
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.UserSearchPrecedence != nil {
		toSerialize["UserSearchPrecedence"] = o.UserSearchPrecedence
	}
	if o.Var0Idp != nil {
		toSerialize["_0_Idp"] = o.Var0Idp
	}
	if o.ApplianceAccount != nil {
		toSerialize["ApplianceAccount"] = o.ApplianceAccount
	}
	if o.Groups != nil {
		toSerialize["Groups"] = o.Groups
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}
	if o.Providers != nil {
		toSerialize["Providers"] = o.Providers
	}
	return json.Marshal(toSerialize)
}

type NullableIamLdapPolicyAllOf struct {
	value *IamLdapPolicyAllOf
	isSet bool
}

func (v NullableIamLdapPolicyAllOf) Get() *IamLdapPolicyAllOf {
	return v.value
}

func (v *NullableIamLdapPolicyAllOf) Set(val *IamLdapPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamLdapPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamLdapPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamLdapPolicyAllOf(val *IamLdapPolicyAllOf) *NullableIamLdapPolicyAllOf {
	return &NullableIamLdapPolicyAllOf{value: val, isSet: true}
}

func (v NullableIamLdapPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamLdapPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
