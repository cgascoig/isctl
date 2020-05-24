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

// HyperflexAppCatalogAllOf Definition of the list of properties defined in 'hyperflex.AppCatalog', excluding properties defined in parent classes.
type HyperflexAppCatalogAllOf struct {
	// The catalog version used in HyperFlex cluster configuration service.
	Version              *string                                    `json:"Version,omitempty" yaml:"Version,omitempty"`
	FeatureLimitExternal *HyperflexFeatureLimitExternalRelationship `json:"FeatureLimitExternal,omitempty" yaml:"FeatureLimitExternal,omitempty"`
	FeatureLimitInternal *HyperflexFeatureLimitInternalRelationship `json:"FeatureLimitInternal,omitempty" yaml:"FeatureLimitInternal,omitempty"`
	// An array of relationships to hyperflexHxdpVersion resources.
	HxdpVersions *[]HyperflexHxdpVersionRelationship `json:"HxdpVersions,omitempty" yaml:"HxdpVersions,omitempty"`
	// An array of relationships to hyperflexCapabilityInfo resources.
	HyperflexCapabilityInfos *[]HyperflexCapabilityInfoRelationship `json:"HyperflexCapabilityInfos,omitempty" yaml:"HyperflexCapabilityInfos,omitempty"`
	// An array of relationships to hclHyperflexSoftwareCompatibilityInfo resources.
	HyperflexSoftwareCompatibilityInfos *[]HclHyperflexSoftwareCompatibilityInfoRelationship `json:"HyperflexSoftwareCompatibilityInfos,omitempty" yaml:"HyperflexSoftwareCompatibilityInfos,omitempty"`
	ServerFirmwareVersion               *HyperflexServerFirmwareVersionRelationship          `json:"ServerFirmwareVersion,omitempty" yaml:"ServerFirmwareVersion,omitempty"`
	ServerModel                         *HyperflexServerModelRelationship                    `json:"ServerModel,omitempty" yaml:"ServerModel,omitempty"`
}

// NewHyperflexAppCatalogAllOf instantiates a new HyperflexAppCatalogAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexAppCatalogAllOf() *HyperflexAppCatalogAllOf {
	this := HyperflexAppCatalogAllOf{}
	return &this
}

// NewHyperflexAppCatalogAllOfWithDefaults instantiates a new HyperflexAppCatalogAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexAppCatalogAllOfWithDefaults() *HyperflexAppCatalogAllOf {
	this := HyperflexAppCatalogAllOf{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexAppCatalogAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexAppCatalogAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetFeatureLimitExternal returns the FeatureLimitExternal field value if set, zero value otherwise.
func (o *HyperflexAppCatalogAllOf) GetFeatureLimitExternal() HyperflexFeatureLimitExternalRelationship {
	if o == nil || o.FeatureLimitExternal == nil {
		var ret HyperflexFeatureLimitExternalRelationship
		return ret
	}
	return *o.FeatureLimitExternal
}

// GetFeatureLimitExternalOk returns a tuple with the FeatureLimitExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetFeatureLimitExternalOk() (*HyperflexFeatureLimitExternalRelationship, bool) {
	if o == nil || o.FeatureLimitExternal == nil {
		return nil, false
	}
	return o.FeatureLimitExternal, true
}

// HasFeatureLimitExternal returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasFeatureLimitExternal() bool {
	if o != nil && o.FeatureLimitExternal != nil {
		return true
	}

	return false
}

// SetFeatureLimitExternal gets a reference to the given HyperflexFeatureLimitExternalRelationship and assigns it to the FeatureLimitExternal field.
func (o *HyperflexAppCatalogAllOf) SetFeatureLimitExternal(v HyperflexFeatureLimitExternalRelationship) {
	o.FeatureLimitExternal = &v
}

// GetFeatureLimitInternal returns the FeatureLimitInternal field value if set, zero value otherwise.
func (o *HyperflexAppCatalogAllOf) GetFeatureLimitInternal() HyperflexFeatureLimitInternalRelationship {
	if o == nil || o.FeatureLimitInternal == nil {
		var ret HyperflexFeatureLimitInternalRelationship
		return ret
	}
	return *o.FeatureLimitInternal
}

// GetFeatureLimitInternalOk returns a tuple with the FeatureLimitInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetFeatureLimitInternalOk() (*HyperflexFeatureLimitInternalRelationship, bool) {
	if o == nil || o.FeatureLimitInternal == nil {
		return nil, false
	}
	return o.FeatureLimitInternal, true
}

// HasFeatureLimitInternal returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasFeatureLimitInternal() bool {
	if o != nil && o.FeatureLimitInternal != nil {
		return true
	}

	return false
}

// SetFeatureLimitInternal gets a reference to the given HyperflexFeatureLimitInternalRelationship and assigns it to the FeatureLimitInternal field.
func (o *HyperflexAppCatalogAllOf) SetFeatureLimitInternal(v HyperflexFeatureLimitInternalRelationship) {
	o.FeatureLimitInternal = &v
}

// GetHxdpVersions returns the HxdpVersions field value if set, zero value otherwise.
func (o *HyperflexAppCatalogAllOf) GetHxdpVersions() []HyperflexHxdpVersionRelationship {
	if o == nil || o.HxdpVersions == nil {
		var ret []HyperflexHxdpVersionRelationship
		return ret
	}
	return *o.HxdpVersions
}

// GetHxdpVersionsOk returns a tuple with the HxdpVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetHxdpVersionsOk() (*[]HyperflexHxdpVersionRelationship, bool) {
	if o == nil || o.HxdpVersions == nil {
		return nil, false
	}
	return o.HxdpVersions, true
}

// HasHxdpVersions returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasHxdpVersions() bool {
	if o != nil && o.HxdpVersions != nil {
		return true
	}

	return false
}

// SetHxdpVersions gets a reference to the given []HyperflexHxdpVersionRelationship and assigns it to the HxdpVersions field.
func (o *HyperflexAppCatalogAllOf) SetHxdpVersions(v []HyperflexHxdpVersionRelationship) {
	o.HxdpVersions = &v
}

// GetHyperflexCapabilityInfos returns the HyperflexCapabilityInfos field value if set, zero value otherwise.
func (o *HyperflexAppCatalogAllOf) GetHyperflexCapabilityInfos() []HyperflexCapabilityInfoRelationship {
	if o == nil || o.HyperflexCapabilityInfos == nil {
		var ret []HyperflexCapabilityInfoRelationship
		return ret
	}
	return *o.HyperflexCapabilityInfos
}

// GetHyperflexCapabilityInfosOk returns a tuple with the HyperflexCapabilityInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetHyperflexCapabilityInfosOk() (*[]HyperflexCapabilityInfoRelationship, bool) {
	if o == nil || o.HyperflexCapabilityInfos == nil {
		return nil, false
	}
	return o.HyperflexCapabilityInfos, true
}

// HasHyperflexCapabilityInfos returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasHyperflexCapabilityInfos() bool {
	if o != nil && o.HyperflexCapabilityInfos != nil {
		return true
	}

	return false
}

// SetHyperflexCapabilityInfos gets a reference to the given []HyperflexCapabilityInfoRelationship and assigns it to the HyperflexCapabilityInfos field.
func (o *HyperflexAppCatalogAllOf) SetHyperflexCapabilityInfos(v []HyperflexCapabilityInfoRelationship) {
	o.HyperflexCapabilityInfos = &v
}

// GetHyperflexSoftwareCompatibilityInfos returns the HyperflexSoftwareCompatibilityInfos field value if set, zero value otherwise.
func (o *HyperflexAppCatalogAllOf) GetHyperflexSoftwareCompatibilityInfos() []HclHyperflexSoftwareCompatibilityInfoRelationship {
	if o == nil || o.HyperflexSoftwareCompatibilityInfos == nil {
		var ret []HclHyperflexSoftwareCompatibilityInfoRelationship
		return ret
	}
	return *o.HyperflexSoftwareCompatibilityInfos
}

// GetHyperflexSoftwareCompatibilityInfosOk returns a tuple with the HyperflexSoftwareCompatibilityInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetHyperflexSoftwareCompatibilityInfosOk() (*[]HclHyperflexSoftwareCompatibilityInfoRelationship, bool) {
	if o == nil || o.HyperflexSoftwareCompatibilityInfos == nil {
		return nil, false
	}
	return o.HyperflexSoftwareCompatibilityInfos, true
}

// HasHyperflexSoftwareCompatibilityInfos returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasHyperflexSoftwareCompatibilityInfos() bool {
	if o != nil && o.HyperflexSoftwareCompatibilityInfos != nil {
		return true
	}

	return false
}

// SetHyperflexSoftwareCompatibilityInfos gets a reference to the given []HclHyperflexSoftwareCompatibilityInfoRelationship and assigns it to the HyperflexSoftwareCompatibilityInfos field.
func (o *HyperflexAppCatalogAllOf) SetHyperflexSoftwareCompatibilityInfos(v []HclHyperflexSoftwareCompatibilityInfoRelationship) {
	o.HyperflexSoftwareCompatibilityInfos = &v
}

// GetServerFirmwareVersion returns the ServerFirmwareVersion field value if set, zero value otherwise.
func (o *HyperflexAppCatalogAllOf) GetServerFirmwareVersion() HyperflexServerFirmwareVersionRelationship {
	if o == nil || o.ServerFirmwareVersion == nil {
		var ret HyperflexServerFirmwareVersionRelationship
		return ret
	}
	return *o.ServerFirmwareVersion
}

// GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetServerFirmwareVersionOk() (*HyperflexServerFirmwareVersionRelationship, bool) {
	if o == nil || o.ServerFirmwareVersion == nil {
		return nil, false
	}
	return o.ServerFirmwareVersion, true
}

// HasServerFirmwareVersion returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasServerFirmwareVersion() bool {
	if o != nil && o.ServerFirmwareVersion != nil {
		return true
	}

	return false
}

// SetServerFirmwareVersion gets a reference to the given HyperflexServerFirmwareVersionRelationship and assigns it to the ServerFirmwareVersion field.
func (o *HyperflexAppCatalogAllOf) SetServerFirmwareVersion(v HyperflexServerFirmwareVersionRelationship) {
	o.ServerFirmwareVersion = &v
}

// GetServerModel returns the ServerModel field value if set, zero value otherwise.
func (o *HyperflexAppCatalogAllOf) GetServerModel() HyperflexServerModelRelationship {
	if o == nil || o.ServerModel == nil {
		var ret HyperflexServerModelRelationship
		return ret
	}
	return *o.ServerModel
}

// GetServerModelOk returns a tuple with the ServerModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetServerModelOk() (*HyperflexServerModelRelationship, bool) {
	if o == nil || o.ServerModel == nil {
		return nil, false
	}
	return o.ServerModel, true
}

// HasServerModel returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasServerModel() bool {
	if o != nil && o.ServerModel != nil {
		return true
	}

	return false
}

// SetServerModel gets a reference to the given HyperflexServerModelRelationship and assigns it to the ServerModel field.
func (o *HyperflexAppCatalogAllOf) SetServerModel(v HyperflexServerModelRelationship) {
	o.ServerModel = &v
}

func (o HyperflexAppCatalogAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.FeatureLimitExternal != nil {
		toSerialize["FeatureLimitExternal"] = o.FeatureLimitExternal
	}
	if o.FeatureLimitInternal != nil {
		toSerialize["FeatureLimitInternal"] = o.FeatureLimitInternal
	}
	if o.HxdpVersions != nil {
		toSerialize["HxdpVersions"] = o.HxdpVersions
	}
	if o.HyperflexCapabilityInfos != nil {
		toSerialize["HyperflexCapabilityInfos"] = o.HyperflexCapabilityInfos
	}
	if o.HyperflexSoftwareCompatibilityInfos != nil {
		toSerialize["HyperflexSoftwareCompatibilityInfos"] = o.HyperflexSoftwareCompatibilityInfos
	}
	if o.ServerFirmwareVersion != nil {
		toSerialize["ServerFirmwareVersion"] = o.ServerFirmwareVersion
	}
	if o.ServerModel != nil {
		toSerialize["ServerModel"] = o.ServerModel
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexAppCatalogAllOf struct {
	value *HyperflexAppCatalogAllOf
	isSet bool
}

func (v NullableHyperflexAppCatalogAllOf) Get() *HyperflexAppCatalogAllOf {
	return v.value
}

func (v *NullableHyperflexAppCatalogAllOf) Set(val *HyperflexAppCatalogAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAppCatalogAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAppCatalogAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAppCatalogAllOf(val *HyperflexAppCatalogAllOf) *NullableHyperflexAppCatalogAllOf {
	return &NullableHyperflexAppCatalogAllOf{value: val, isSet: true}
}

func (v NullableHyperflexAppCatalogAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAppCatalogAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
