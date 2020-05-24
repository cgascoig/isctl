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

// SnmpUser Complex type for a User based security model, for communication between an agent and manager. Applicable only for SNMPv3.
type SnmpUser struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// Authorization password for the user.
	AuthPassword *string `json:"AuthPassword,omitempty" yaml:"AuthPassword,omitempty"`
	// Authorization protocol for authenticating the user.
	AuthType *string `json:"AuthType,omitempty" yaml:"AuthType,omitempty"`
	// Indicates whether the value of the 'authPassword' property has been set.
	IsAuthPasswordSet *bool `json:"IsAuthPasswordSet,omitempty" yaml:"IsAuthPasswordSet,omitempty"`
	// Indicates whether the value of the 'privacyPassword' property has been set.
	IsPrivacyPasswordSet *bool `json:"IsPrivacyPasswordSet,omitempty" yaml:"IsPrivacyPasswordSet,omitempty"`
	// SNMP username. Must have a minimum of 1 and and a maximum of 31 characters.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// Privacy password for the user.
	PrivacyPassword *string `json:"PrivacyPassword,omitempty" yaml:"PrivacyPassword,omitempty"`
	// Privacy protocol for the user.
	PrivacyType *string `json:"PrivacyType,omitempty" yaml:"PrivacyType,omitempty"`
	// Security mechanism used for communication between agent and manager.
	SecurityLevel *string `json:"SecurityLevel,omitempty" yaml:"SecurityLevel,omitempty"`
}

// NewSnmpUser instantiates a new SnmpUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpUser() *SnmpUser {
	this := SnmpUser{}
	var authType string = "NA"
	this.AuthType = &authType
	var privacyType string = "NA"
	this.PrivacyType = &privacyType
	var securityLevel string = "AuthPriv"
	this.SecurityLevel = &securityLevel
	return &this
}

// NewSnmpUserWithDefaults instantiates a new SnmpUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpUserWithDefaults() *SnmpUser {
	this := SnmpUser{}
	var authType string = "NA"
	this.AuthType = &authType
	var privacyType string = "NA"
	this.PrivacyType = &privacyType
	var securityLevel string = "AuthPriv"
	this.SecurityLevel = &securityLevel
	return &this
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *SnmpUser) GetAuthPassword() string {
	if o == nil || o.AuthPassword == nil {
		var ret string
		return ret
	}
	return *o.AuthPassword
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUser) GetAuthPasswordOk() (*string, bool) {
	if o == nil || o.AuthPassword == nil {
		return nil, false
	}
	return o.AuthPassword, true
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *SnmpUser) HasAuthPassword() bool {
	if o != nil && o.AuthPassword != nil {
		return true
	}

	return false
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *SnmpUser) SetAuthPassword(v string) {
	o.AuthPassword = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *SnmpUser) GetAuthType() string {
	if o == nil || o.AuthType == nil {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUser) GetAuthTypeOk() (*string, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *SnmpUser) HasAuthType() bool {
	if o != nil && o.AuthType != nil {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *SnmpUser) SetAuthType(v string) {
	o.AuthType = &v
}

// GetIsAuthPasswordSet returns the IsAuthPasswordSet field value if set, zero value otherwise.
func (o *SnmpUser) GetIsAuthPasswordSet() bool {
	if o == nil || o.IsAuthPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsAuthPasswordSet
}

// GetIsAuthPasswordSetOk returns a tuple with the IsAuthPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUser) GetIsAuthPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsAuthPasswordSet == nil {
		return nil, false
	}
	return o.IsAuthPasswordSet, true
}

// HasIsAuthPasswordSet returns a boolean if a field has been set.
func (o *SnmpUser) HasIsAuthPasswordSet() bool {
	if o != nil && o.IsAuthPasswordSet != nil {
		return true
	}

	return false
}

// SetIsAuthPasswordSet gets a reference to the given bool and assigns it to the IsAuthPasswordSet field.
func (o *SnmpUser) SetIsAuthPasswordSet(v bool) {
	o.IsAuthPasswordSet = &v
}

// GetIsPrivacyPasswordSet returns the IsPrivacyPasswordSet field value if set, zero value otherwise.
func (o *SnmpUser) GetIsPrivacyPasswordSet() bool {
	if o == nil || o.IsPrivacyPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivacyPasswordSet
}

// GetIsPrivacyPasswordSetOk returns a tuple with the IsPrivacyPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUser) GetIsPrivacyPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPrivacyPasswordSet == nil {
		return nil, false
	}
	return o.IsPrivacyPasswordSet, true
}

// HasIsPrivacyPasswordSet returns a boolean if a field has been set.
func (o *SnmpUser) HasIsPrivacyPasswordSet() bool {
	if o != nil && o.IsPrivacyPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPrivacyPasswordSet gets a reference to the given bool and assigns it to the IsPrivacyPasswordSet field.
func (o *SnmpUser) SetIsPrivacyPasswordSet(v bool) {
	o.IsPrivacyPasswordSet = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SnmpUser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUser) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SnmpUser) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SnmpUser) SetName(v string) {
	o.Name = &v
}

// GetPrivacyPassword returns the PrivacyPassword field value if set, zero value otherwise.
func (o *SnmpUser) GetPrivacyPassword() string {
	if o == nil || o.PrivacyPassword == nil {
		var ret string
		return ret
	}
	return *o.PrivacyPassword
}

// GetPrivacyPasswordOk returns a tuple with the PrivacyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUser) GetPrivacyPasswordOk() (*string, bool) {
	if o == nil || o.PrivacyPassword == nil {
		return nil, false
	}
	return o.PrivacyPassword, true
}

// HasPrivacyPassword returns a boolean if a field has been set.
func (o *SnmpUser) HasPrivacyPassword() bool {
	if o != nil && o.PrivacyPassword != nil {
		return true
	}

	return false
}

// SetPrivacyPassword gets a reference to the given string and assigns it to the PrivacyPassword field.
func (o *SnmpUser) SetPrivacyPassword(v string) {
	o.PrivacyPassword = &v
}

// GetPrivacyType returns the PrivacyType field value if set, zero value otherwise.
func (o *SnmpUser) GetPrivacyType() string {
	if o == nil || o.PrivacyType == nil {
		var ret string
		return ret
	}
	return *o.PrivacyType
}

// GetPrivacyTypeOk returns a tuple with the PrivacyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUser) GetPrivacyTypeOk() (*string, bool) {
	if o == nil || o.PrivacyType == nil {
		return nil, false
	}
	return o.PrivacyType, true
}

// HasPrivacyType returns a boolean if a field has been set.
func (o *SnmpUser) HasPrivacyType() bool {
	if o != nil && o.PrivacyType != nil {
		return true
	}

	return false
}

// SetPrivacyType gets a reference to the given string and assigns it to the PrivacyType field.
func (o *SnmpUser) SetPrivacyType(v string) {
	o.PrivacyType = &v
}

// GetSecurityLevel returns the SecurityLevel field value if set, zero value otherwise.
func (o *SnmpUser) GetSecurityLevel() string {
	if o == nil || o.SecurityLevel == nil {
		var ret string
		return ret
	}
	return *o.SecurityLevel
}

// GetSecurityLevelOk returns a tuple with the SecurityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUser) GetSecurityLevelOk() (*string, bool) {
	if o == nil || o.SecurityLevel == nil {
		return nil, false
	}
	return o.SecurityLevel, true
}

// HasSecurityLevel returns a boolean if a field has been set.
func (o *SnmpUser) HasSecurityLevel() bool {
	if o != nil && o.SecurityLevel != nil {
		return true
	}

	return false
}

// SetSecurityLevel gets a reference to the given string and assigns it to the SecurityLevel field.
func (o *SnmpUser) SetSecurityLevel(v string) {
	o.SecurityLevel = &v
}

func (o SnmpUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.AuthPassword != nil {
		toSerialize["AuthPassword"] = o.AuthPassword
	}
	if o.AuthType != nil {
		toSerialize["AuthType"] = o.AuthType
	}
	if o.IsAuthPasswordSet != nil {
		toSerialize["IsAuthPasswordSet"] = o.IsAuthPasswordSet
	}
	if o.IsPrivacyPasswordSet != nil {
		toSerialize["IsPrivacyPasswordSet"] = o.IsPrivacyPasswordSet
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PrivacyPassword != nil {
		toSerialize["PrivacyPassword"] = o.PrivacyPassword
	}
	if o.PrivacyType != nil {
		toSerialize["PrivacyType"] = o.PrivacyType
	}
	if o.SecurityLevel != nil {
		toSerialize["SecurityLevel"] = o.SecurityLevel
	}
	return json.Marshal(toSerialize)
}

type NullableSnmpUser struct {
	value *SnmpUser
	isSet bool
}

func (v NullableSnmpUser) Get() *SnmpUser {
	return v.value
}

func (v *NullableSnmpUser) Set(val *SnmpUser) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpUser) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpUser(val *SnmpUser) *NullableSnmpUser {
	return &NullableSnmpUser{value: val, isSet: true}
}

func (v NullableSnmpUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
