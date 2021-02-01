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

// AssetTerraformIntegrationTerraformCloudOptions Cloud options for Terrraform Cloud platform type.
type AssetTerraformIntegrationTerraformCloudOptions struct {
	AssetServiceOptions `yaml:"AssetServiceOptions,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Default organization for Terraform Cloud platform type.
	DefaultTerraformOrganization *string `json:"DefaultTerraformOrganization,omitempty" yaml:"DefaultTerraformOrganization,omitempty"`
}

// NewAssetTerraformIntegrationTerraformCloudOptions instantiates a new AssetTerraformIntegrationTerraformCloudOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTerraformIntegrationTerraformCloudOptions(classId string, objectType string) *AssetTerraformIntegrationTerraformCloudOptions {
	this := AssetTerraformIntegrationTerraformCloudOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetTerraformIntegrationTerraformCloudOptionsWithDefaults instantiates a new AssetTerraformIntegrationTerraformCloudOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTerraformIntegrationTerraformCloudOptionsWithDefaults() *AssetTerraformIntegrationTerraformCloudOptions {
	this := AssetTerraformIntegrationTerraformCloudOptions{}
	var classId string = "asset.TerraformIntegrationTerraformCloudOptions"
	this.ClassId = classId
	var objectType string = "asset.TerraformIntegrationTerraformCloudOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetTerraformIntegrationTerraformCloudOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetTerraformIntegrationTerraformCloudOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetTerraformIntegrationTerraformCloudOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetTerraformIntegrationTerraformCloudOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetTerraformIntegrationTerraformCloudOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetTerraformIntegrationTerraformCloudOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultTerraformOrganization returns the DefaultTerraformOrganization field value if set, zero value otherwise.
func (o *AssetTerraformIntegrationTerraformCloudOptions) GetDefaultTerraformOrganization() string {
	if o == nil || o.DefaultTerraformOrganization == nil {
		var ret string
		return ret
	}
	return *o.DefaultTerraformOrganization
}

// GetDefaultTerraformOrganizationOk returns a tuple with the DefaultTerraformOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTerraformIntegrationTerraformCloudOptions) GetDefaultTerraformOrganizationOk() (*string, bool) {
	if o == nil || o.DefaultTerraformOrganization == nil {
		return nil, false
	}
	return o.DefaultTerraformOrganization, true
}

// HasDefaultTerraformOrganization returns a boolean if a field has been set.
func (o *AssetTerraformIntegrationTerraformCloudOptions) HasDefaultTerraformOrganization() bool {
	if o != nil && o.DefaultTerraformOrganization != nil {
		return true
	}

	return false
}

// SetDefaultTerraformOrganization gets a reference to the given string and assigns it to the DefaultTerraformOrganization field.
func (o *AssetTerraformIntegrationTerraformCloudOptions) SetDefaultTerraformOrganization(v string) {
	o.DefaultTerraformOrganization = &v
}

func (o AssetTerraformIntegrationTerraformCloudOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DefaultTerraformOrganization != nil {
		toSerialize["DefaultTerraformOrganization"] = o.DefaultTerraformOrganization
	}
	return json.Marshal(toSerialize)
}

type NullableAssetTerraformIntegrationTerraformCloudOptions struct {
	value *AssetTerraformIntegrationTerraformCloudOptions
	isSet bool
}

func (v NullableAssetTerraformIntegrationTerraformCloudOptions) Get() *AssetTerraformIntegrationTerraformCloudOptions {
	return v.value
}

func (v *NullableAssetTerraformIntegrationTerraformCloudOptions) Set(val *AssetTerraformIntegrationTerraformCloudOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTerraformIntegrationTerraformCloudOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTerraformIntegrationTerraformCloudOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTerraformIntegrationTerraformCloudOptions(val *AssetTerraformIntegrationTerraformCloudOptions) *NullableAssetTerraformIntegrationTerraformCloudOptions {
	return &NullableAssetTerraformIntegrationTerraformCloudOptions{value: val, isSet: true}
}

func (v NullableAssetTerraformIntegrationTerraformCloudOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTerraformIntegrationTerraformCloudOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
