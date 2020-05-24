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
	"time"
)

// IamAppRegistrationAllOf Definition of the list of properties defined in 'iam.AppRegistration', excluding properties defined in parent classes.
type IamAppRegistrationAllOf struct {
	// A unique identifier for the OAuth2 client application. The client ID is auto-generated when the AppRegistration object is created.
	ClientId *string `json:"ClientId,omitempty" yaml:"ClientId,omitempty"`
	// App Registration name specified by user.
	ClientName *string `json:"ClientName,omitempty" yaml:"ClientName,omitempty"`
	// The OAuth2 client secret. The value of this property is generated when grantType includes 'client-credentials'. Otherwise, no client-secret is generated.
	ClientSecret *string `json:"ClientSecret,omitempty" yaml:"ClientSecret,omitempty"`
	// The type of the OAuth2 client (public or confidential), as specified in https://tools.ietf.org/html/rfc6749#section-2.1.
	ClientType *string `json:"ClientType,omitempty" yaml:"ClientType,omitempty"`
	// Description of the application.
	Description  *string   `json:"Description,omitempty" yaml:"Description,omitempty"`
	GrantTypes   *[]string `json:"GrantTypes,omitempty" yaml:"GrantTypes,omitempty"`
	RedirectUris *[]string `json:"RedirectUris,omitempty" yaml:"RedirectUris,omitempty"`
	// Set value to true to renew the client-secret. Applicable to client_credentials grant type.
	RenewClientSecret *bool     `json:"RenewClientSecret,omitempty" yaml:"RenewClientSecret,omitempty"`
	ResponseTypes     *[]string `json:"ResponseTypes,omitempty" yaml:"ResponseTypes,omitempty"`
	// Used to perform revocation for tokens of AppRegistration. Updated only internally is case Revoke property come from UI with value true. On each request with OAuth2 access token the CreationTime of the OAuth2 token will be compared to RevokationTimestamp of the corresponding App Registration.
	RevocationTimestamp *time.Time `json:"RevocationTimestamp,omitempty" yaml:"RevocationTimestamp,omitempty"`
	// Used to trigger update the revocationTimestamp value. If UI sent updating request with the Revoke value is true, then update RevocationTimestamp.
	Revoke  *bool                   `json:"Revoke,omitempty" yaml:"Revoke,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty" yaml:"Account,omitempty"`
	// An array of relationships to iamOAuthToken resources.
	OauthTokens *[]IamOAuthTokenRelationship `json:"OauthTokens,omitempty" yaml:"OauthTokens,omitempty"`
	Permission  *IamPermissionRelationship   `json:"Permission,omitempty" yaml:"Permission,omitempty"`
	// An array of relationships to iamRole resources.
	Roles *[]IamRoleRelationship `json:"Roles,omitempty" yaml:"Roles,omitempty"`
	User  *IamUserRelationship   `json:"User,omitempty" yaml:"User,omitempty"`
}

// NewIamAppRegistrationAllOf instantiates a new IamAppRegistrationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamAppRegistrationAllOf() *IamAppRegistrationAllOf {
	this := IamAppRegistrationAllOf{}
	var clientType string = "public"
	this.ClientType = &clientType
	return &this
}

// NewIamAppRegistrationAllOfWithDefaults instantiates a new IamAppRegistrationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamAppRegistrationAllOfWithDefaults() *IamAppRegistrationAllOf {
	this := IamAppRegistrationAllOf{}
	var clientType string = "public"
	this.ClientType = &clientType
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IamAppRegistrationAllOf) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *IamAppRegistrationAllOf) SetClientName(v string) {
	o.ClientName = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *IamAppRegistrationAllOf) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetClientType returns the ClientType field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetClientType() string {
	if o == nil || o.ClientType == nil {
		var ret string
		return ret
	}
	return *o.ClientType
}

// GetClientTypeOk returns a tuple with the ClientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetClientTypeOk() (*string, bool) {
	if o == nil || o.ClientType == nil {
		return nil, false
	}
	return o.ClientType, true
}

// HasClientType returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasClientType() bool {
	if o != nil && o.ClientType != nil {
		return true
	}

	return false
}

// SetClientType gets a reference to the given string and assigns it to the ClientType field.
func (o *IamAppRegistrationAllOf) SetClientType(v string) {
	o.ClientType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamAppRegistrationAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetGrantTypes returns the GrantTypes field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetGrantTypes() []string {
	if o == nil || o.GrantTypes == nil {
		var ret []string
		return ret
	}
	return *o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetGrantTypesOk() (*[]string, bool) {
	if o == nil || o.GrantTypes == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// HasGrantTypes returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasGrantTypes() bool {
	if o != nil && o.GrantTypes != nil {
		return true
	}

	return false
}

// SetGrantTypes gets a reference to the given []string and assigns it to the GrantTypes field.
func (o *IamAppRegistrationAllOf) SetGrantTypes(v []string) {
	o.GrantTypes = &v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetRedirectUris() []string {
	if o == nil || o.RedirectUris == nil {
		var ret []string
		return ret
	}
	return *o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetRedirectUrisOk() (*[]string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *IamAppRegistrationAllOf) SetRedirectUris(v []string) {
	o.RedirectUris = &v
}

// GetRenewClientSecret returns the RenewClientSecret field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetRenewClientSecret() bool {
	if o == nil || o.RenewClientSecret == nil {
		var ret bool
		return ret
	}
	return *o.RenewClientSecret
}

// GetRenewClientSecretOk returns a tuple with the RenewClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetRenewClientSecretOk() (*bool, bool) {
	if o == nil || o.RenewClientSecret == nil {
		return nil, false
	}
	return o.RenewClientSecret, true
}

// HasRenewClientSecret returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasRenewClientSecret() bool {
	if o != nil && o.RenewClientSecret != nil {
		return true
	}

	return false
}

// SetRenewClientSecret gets a reference to the given bool and assigns it to the RenewClientSecret field.
func (o *IamAppRegistrationAllOf) SetRenewClientSecret(v bool) {
	o.RenewClientSecret = &v
}

// GetResponseTypes returns the ResponseTypes field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetResponseTypes() []string {
	if o == nil || o.ResponseTypes == nil {
		var ret []string
		return ret
	}
	return *o.ResponseTypes
}

// GetResponseTypesOk returns a tuple with the ResponseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetResponseTypesOk() (*[]string, bool) {
	if o == nil || o.ResponseTypes == nil {
		return nil, false
	}
	return o.ResponseTypes, true
}

// HasResponseTypes returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasResponseTypes() bool {
	if o != nil && o.ResponseTypes != nil {
		return true
	}

	return false
}

// SetResponseTypes gets a reference to the given []string and assigns it to the ResponseTypes field.
func (o *IamAppRegistrationAllOf) SetResponseTypes(v []string) {
	o.ResponseTypes = &v
}

// GetRevocationTimestamp returns the RevocationTimestamp field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetRevocationTimestamp() time.Time {
	if o == nil || o.RevocationTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.RevocationTimestamp
}

// GetRevocationTimestampOk returns a tuple with the RevocationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetRevocationTimestampOk() (*time.Time, bool) {
	if o == nil || o.RevocationTimestamp == nil {
		return nil, false
	}
	return o.RevocationTimestamp, true
}

// HasRevocationTimestamp returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasRevocationTimestamp() bool {
	if o != nil && o.RevocationTimestamp != nil {
		return true
	}

	return false
}

// SetRevocationTimestamp gets a reference to the given time.Time and assigns it to the RevocationTimestamp field.
func (o *IamAppRegistrationAllOf) SetRevocationTimestamp(v time.Time) {
	o.RevocationTimestamp = &v
}

// GetRevoke returns the Revoke field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetRevoke() bool {
	if o == nil || o.Revoke == nil {
		var ret bool
		return ret
	}
	return *o.Revoke
}

// GetRevokeOk returns a tuple with the Revoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetRevokeOk() (*bool, bool) {
	if o == nil || o.Revoke == nil {
		return nil, false
	}
	return o.Revoke, true
}

// HasRevoke returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasRevoke() bool {
	if o != nil && o.Revoke != nil {
		return true
	}

	return false
}

// SetRevoke gets a reference to the given bool and assigns it to the Revoke field.
func (o *IamAppRegistrationAllOf) SetRevoke(v bool) {
	o.Revoke = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamAppRegistrationAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetOauthTokens returns the OauthTokens field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetOauthTokens() []IamOAuthTokenRelationship {
	if o == nil || o.OauthTokens == nil {
		var ret []IamOAuthTokenRelationship
		return ret
	}
	return *o.OauthTokens
}

// GetOauthTokensOk returns a tuple with the OauthTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetOauthTokensOk() (*[]IamOAuthTokenRelationship, bool) {
	if o == nil || o.OauthTokens == nil {
		return nil, false
	}
	return o.OauthTokens, true
}

// HasOauthTokens returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasOauthTokens() bool {
	if o != nil && o.OauthTokens != nil {
		return true
	}

	return false
}

// SetOauthTokens gets a reference to the given []IamOAuthTokenRelationship and assigns it to the OauthTokens field.
func (o *IamAppRegistrationAllOf) SetOauthTokens(v []IamOAuthTokenRelationship) {
	o.OauthTokens = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetPermission() IamPermissionRelationship {
	if o == nil || o.Permission == nil {
		var ret IamPermissionRelationship
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetPermissionOk() (*IamPermissionRelationship, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given IamPermissionRelationship and assigns it to the Permission field.
func (o *IamAppRegistrationAllOf) SetPermission(v IamPermissionRelationship) {
	o.Permission = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetRoles() []IamRoleRelationship {
	if o == nil || o.Roles == nil {
		var ret []IamRoleRelationship
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetRolesOk() (*[]IamRoleRelationship, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []IamRoleRelationship and assigns it to the Roles field.
func (o *IamAppRegistrationAllOf) SetRoles(v []IamRoleRelationship) {
	o.Roles = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IamAppRegistrationAllOf) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistrationAllOf) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IamAppRegistrationAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *IamAppRegistrationAllOf) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o IamAppRegistrationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["ClientId"] = o.ClientId
	}
	if o.ClientName != nil {
		toSerialize["ClientName"] = o.ClientName
	}
	if o.ClientSecret != nil {
		toSerialize["ClientSecret"] = o.ClientSecret
	}
	if o.ClientType != nil {
		toSerialize["ClientType"] = o.ClientType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.GrantTypes != nil {
		toSerialize["GrantTypes"] = o.GrantTypes
	}
	if o.RedirectUris != nil {
		toSerialize["RedirectUris"] = o.RedirectUris
	}
	if o.RenewClientSecret != nil {
		toSerialize["RenewClientSecret"] = o.RenewClientSecret
	}
	if o.ResponseTypes != nil {
		toSerialize["ResponseTypes"] = o.ResponseTypes
	}
	if o.RevocationTimestamp != nil {
		toSerialize["RevocationTimestamp"] = o.RevocationTimestamp
	}
	if o.Revoke != nil {
		toSerialize["Revoke"] = o.Revoke
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.OauthTokens != nil {
		toSerialize["OauthTokens"] = o.OauthTokens
	}
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}
	if o.Roles != nil {
		toSerialize["Roles"] = o.Roles
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableIamAppRegistrationAllOf struct {
	value *IamAppRegistrationAllOf
	isSet bool
}

func (v NullableIamAppRegistrationAllOf) Get() *IamAppRegistrationAllOf {
	return v.value
}

func (v *NullableIamAppRegistrationAllOf) Set(val *IamAppRegistrationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamAppRegistrationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamAppRegistrationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamAppRegistrationAllOf(val *IamAppRegistrationAllOf) *NullableIamAppRegistrationAllOf {
	return &NullableIamAppRegistrationAllOf{value: val, isSet: true}
}

func (v NullableIamAppRegistrationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamAppRegistrationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
