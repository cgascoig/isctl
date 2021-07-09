/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-28T13:03:38Z.
 *
 * API version: 1.0.9-4267
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"reflect"
	"strings"
)

// HyperflexAppCatalog A catalog for managing application settings for HyperFlex cluster configuration service.
type HyperflexAppCatalog struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The catalog version used in HyperFlex cluster configuration service.
	Version              *string                                    `json:"Version,omitempty"`
	FeatureLimitExternal *HyperflexFeatureLimitExternalRelationship `json:"FeatureLimitExternal,omitempty"`
	FeatureLimitInternal *HyperflexFeatureLimitInternalRelationship `json:"FeatureLimitInternal,omitempty"`
	// An array of relationships to hyperflexHxdpVersion resources.
	HxdpVersions []HyperflexHxdpVersionRelationship `json:"HxdpVersions,omitempty"`
	// An array of relationships to hyperflexCapabilityInfo resources.
	HyperflexCapabilityInfos []HyperflexCapabilityInfoRelationship `json:"HyperflexCapabilityInfos,omitempty"`
	// An array of relationships to hclHyperflexSoftwareCompatibilityInfo resources.
	HyperflexSoftwareCompatibilityInfos []HclHyperflexSoftwareCompatibilityInfoRelationship `json:"HyperflexSoftwareCompatibilityInfos,omitempty"`
	ServerFirmwareVersion               *HyperflexServerFirmwareVersionRelationship         `json:"ServerFirmwareVersion,omitempty"`
	ServerModel                         *HyperflexServerModelRelationship                   `json:"ServerModel,omitempty"`
	// An array of relationships to hyperflexSoftwareDistributionEntry resources.
	SoftwareDistributions []HyperflexSoftwareDistributionEntryRelationship `json:"SoftwareDistributions,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _HyperflexAppCatalog HyperflexAppCatalog

// NewHyperflexAppCatalog instantiates a new HyperflexAppCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexAppCatalog(classId string, objectType string) *HyperflexAppCatalog {
	this := HyperflexAppCatalog{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexAppCatalogWithDefaults instantiates a new HyperflexAppCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexAppCatalogWithDefaults() *HyperflexAppCatalog {
	this := HyperflexAppCatalog{}
	var classId string = "hyperflex.AppCatalog"
	this.ClassId = classId
	var objectType string = "hyperflex.AppCatalog"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexAppCatalog) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalog) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexAppCatalog) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexAppCatalog) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalog) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexAppCatalog) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexAppCatalog) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalog) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexAppCatalog) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexAppCatalog) SetVersion(v string) {
	o.Version = &v
}

// GetFeatureLimitExternal returns the FeatureLimitExternal field value if set, zero value otherwise.
func (o *HyperflexAppCatalog) GetFeatureLimitExternal() HyperflexFeatureLimitExternalRelationship {
	if o == nil || o.FeatureLimitExternal == nil {
		var ret HyperflexFeatureLimitExternalRelationship
		return ret
	}
	return *o.FeatureLimitExternal
}

// GetFeatureLimitExternalOk returns a tuple with the FeatureLimitExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalog) GetFeatureLimitExternalOk() (*HyperflexFeatureLimitExternalRelationship, bool) {
	if o == nil || o.FeatureLimitExternal == nil {
		return nil, false
	}
	return o.FeatureLimitExternal, true
}

// HasFeatureLimitExternal returns a boolean if a field has been set.
func (o *HyperflexAppCatalog) HasFeatureLimitExternal() bool {
	if o != nil && o.FeatureLimitExternal != nil {
		return true
	}

	return false
}

// SetFeatureLimitExternal gets a reference to the given HyperflexFeatureLimitExternalRelationship and assigns it to the FeatureLimitExternal field.
func (o *HyperflexAppCatalog) SetFeatureLimitExternal(v HyperflexFeatureLimitExternalRelationship) {
	o.FeatureLimitExternal = &v
}

// GetFeatureLimitInternal returns the FeatureLimitInternal field value if set, zero value otherwise.
func (o *HyperflexAppCatalog) GetFeatureLimitInternal() HyperflexFeatureLimitInternalRelationship {
	if o == nil || o.FeatureLimitInternal == nil {
		var ret HyperflexFeatureLimitInternalRelationship
		return ret
	}
	return *o.FeatureLimitInternal
}

// GetFeatureLimitInternalOk returns a tuple with the FeatureLimitInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalog) GetFeatureLimitInternalOk() (*HyperflexFeatureLimitInternalRelationship, bool) {
	if o == nil || o.FeatureLimitInternal == nil {
		return nil, false
	}
	return o.FeatureLimitInternal, true
}

// HasFeatureLimitInternal returns a boolean if a field has been set.
func (o *HyperflexAppCatalog) HasFeatureLimitInternal() bool {
	if o != nil && o.FeatureLimitInternal != nil {
		return true
	}

	return false
}

// SetFeatureLimitInternal gets a reference to the given HyperflexFeatureLimitInternalRelationship and assigns it to the FeatureLimitInternal field.
func (o *HyperflexAppCatalog) SetFeatureLimitInternal(v HyperflexFeatureLimitInternalRelationship) {
	o.FeatureLimitInternal = &v
}

// GetHxdpVersions returns the HxdpVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexAppCatalog) GetHxdpVersions() []HyperflexHxdpVersionRelationship {
	if o == nil {
		var ret []HyperflexHxdpVersionRelationship
		return ret
	}
	return o.HxdpVersions
}

// GetHxdpVersionsOk returns a tuple with the HxdpVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexAppCatalog) GetHxdpVersionsOk() (*[]HyperflexHxdpVersionRelationship, bool) {
	if o == nil || o.HxdpVersions == nil {
		return nil, false
	}
	return &o.HxdpVersions, true
}

// HasHxdpVersions returns a boolean if a field has been set.
func (o *HyperflexAppCatalog) HasHxdpVersions() bool {
	if o != nil && o.HxdpVersions != nil {
		return true
	}

	return false
}

// SetHxdpVersions gets a reference to the given []HyperflexHxdpVersionRelationship and assigns it to the HxdpVersions field.
func (o *HyperflexAppCatalog) SetHxdpVersions(v []HyperflexHxdpVersionRelationship) {
	o.HxdpVersions = v
}

// GetHyperflexCapabilityInfos returns the HyperflexCapabilityInfos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexAppCatalog) GetHyperflexCapabilityInfos() []HyperflexCapabilityInfoRelationship {
	if o == nil {
		var ret []HyperflexCapabilityInfoRelationship
		return ret
	}
	return o.HyperflexCapabilityInfos
}

// GetHyperflexCapabilityInfosOk returns a tuple with the HyperflexCapabilityInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexAppCatalog) GetHyperflexCapabilityInfosOk() (*[]HyperflexCapabilityInfoRelationship, bool) {
	if o == nil || o.HyperflexCapabilityInfos == nil {
		return nil, false
	}
	return &o.HyperflexCapabilityInfos, true
}

// HasHyperflexCapabilityInfos returns a boolean if a field has been set.
func (o *HyperflexAppCatalog) HasHyperflexCapabilityInfos() bool {
	if o != nil && o.HyperflexCapabilityInfos != nil {
		return true
	}

	return false
}

// SetHyperflexCapabilityInfos gets a reference to the given []HyperflexCapabilityInfoRelationship and assigns it to the HyperflexCapabilityInfos field.
func (o *HyperflexAppCatalog) SetHyperflexCapabilityInfos(v []HyperflexCapabilityInfoRelationship) {
	o.HyperflexCapabilityInfos = v
}

// GetHyperflexSoftwareCompatibilityInfos returns the HyperflexSoftwareCompatibilityInfos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexAppCatalog) GetHyperflexSoftwareCompatibilityInfos() []HclHyperflexSoftwareCompatibilityInfoRelationship {
	if o == nil {
		var ret []HclHyperflexSoftwareCompatibilityInfoRelationship
		return ret
	}
	return o.HyperflexSoftwareCompatibilityInfos
}

// GetHyperflexSoftwareCompatibilityInfosOk returns a tuple with the HyperflexSoftwareCompatibilityInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexAppCatalog) GetHyperflexSoftwareCompatibilityInfosOk() (*[]HclHyperflexSoftwareCompatibilityInfoRelationship, bool) {
	if o == nil || o.HyperflexSoftwareCompatibilityInfos == nil {
		return nil, false
	}
	return &o.HyperflexSoftwareCompatibilityInfos, true
}

// HasHyperflexSoftwareCompatibilityInfos returns a boolean if a field has been set.
func (o *HyperflexAppCatalog) HasHyperflexSoftwareCompatibilityInfos() bool {
	if o != nil && o.HyperflexSoftwareCompatibilityInfos != nil {
		return true
	}

	return false
}

// SetHyperflexSoftwareCompatibilityInfos gets a reference to the given []HclHyperflexSoftwareCompatibilityInfoRelationship and assigns it to the HyperflexSoftwareCompatibilityInfos field.
func (o *HyperflexAppCatalog) SetHyperflexSoftwareCompatibilityInfos(v []HclHyperflexSoftwareCompatibilityInfoRelationship) {
	o.HyperflexSoftwareCompatibilityInfos = v
}

// GetServerFirmwareVersion returns the ServerFirmwareVersion field value if set, zero value otherwise.
func (o *HyperflexAppCatalog) GetServerFirmwareVersion() HyperflexServerFirmwareVersionRelationship {
	if o == nil || o.ServerFirmwareVersion == nil {
		var ret HyperflexServerFirmwareVersionRelationship
		return ret
	}
	return *o.ServerFirmwareVersion
}

// GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalog) GetServerFirmwareVersionOk() (*HyperflexServerFirmwareVersionRelationship, bool) {
	if o == nil || o.ServerFirmwareVersion == nil {
		return nil, false
	}
	return o.ServerFirmwareVersion, true
}

// HasServerFirmwareVersion returns a boolean if a field has been set.
func (o *HyperflexAppCatalog) HasServerFirmwareVersion() bool {
	if o != nil && o.ServerFirmwareVersion != nil {
		return true
	}

	return false
}

// SetServerFirmwareVersion gets a reference to the given HyperflexServerFirmwareVersionRelationship and assigns it to the ServerFirmwareVersion field.
func (o *HyperflexAppCatalog) SetServerFirmwareVersion(v HyperflexServerFirmwareVersionRelationship) {
	o.ServerFirmwareVersion = &v
}

// GetServerModel returns the ServerModel field value if set, zero value otherwise.
func (o *HyperflexAppCatalog) GetServerModel() HyperflexServerModelRelationship {
	if o == nil || o.ServerModel == nil {
		var ret HyperflexServerModelRelationship
		return ret
	}
	return *o.ServerModel
}

// GetServerModelOk returns a tuple with the ServerModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalog) GetServerModelOk() (*HyperflexServerModelRelationship, bool) {
	if o == nil || o.ServerModel == nil {
		return nil, false
	}
	return o.ServerModel, true
}

// HasServerModel returns a boolean if a field has been set.
func (o *HyperflexAppCatalog) HasServerModel() bool {
	if o != nil && o.ServerModel != nil {
		return true
	}

	return false
}

// SetServerModel gets a reference to the given HyperflexServerModelRelationship and assigns it to the ServerModel field.
func (o *HyperflexAppCatalog) SetServerModel(v HyperflexServerModelRelationship) {
	o.ServerModel = &v
}

// GetSoftwareDistributions returns the SoftwareDistributions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexAppCatalog) GetSoftwareDistributions() []HyperflexSoftwareDistributionEntryRelationship {
	if o == nil {
		var ret []HyperflexSoftwareDistributionEntryRelationship
		return ret
	}
	return o.SoftwareDistributions
}

// GetSoftwareDistributionsOk returns a tuple with the SoftwareDistributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexAppCatalog) GetSoftwareDistributionsOk() (*[]HyperflexSoftwareDistributionEntryRelationship, bool) {
	if o == nil || o.SoftwareDistributions == nil {
		return nil, false
	}
	return &o.SoftwareDistributions, true
}

// HasSoftwareDistributions returns a boolean if a field has been set.
func (o *HyperflexAppCatalog) HasSoftwareDistributions() bool {
	if o != nil && o.SoftwareDistributions != nil {
		return true
	}

	return false
}

// SetSoftwareDistributions gets a reference to the given []HyperflexSoftwareDistributionEntryRelationship and assigns it to the SoftwareDistributions field.
func (o *HyperflexAppCatalog) SetSoftwareDistributions(v []HyperflexSoftwareDistributionEntryRelationship) {
	o.SoftwareDistributions = v
}

func (o HyperflexAppCatalog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
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
	if o.SoftwareDistributions != nil {
		toSerialize["SoftwareDistributions"] = o.SoftwareDistributions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexAppCatalog) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexAppCatalogWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The catalog version used in HyperFlex cluster configuration service.
		Version              *string                                    `json:"Version,omitempty"`
		FeatureLimitExternal *HyperflexFeatureLimitExternalRelationship `json:"FeatureLimitExternal,omitempty"`
		FeatureLimitInternal *HyperflexFeatureLimitInternalRelationship `json:"FeatureLimitInternal,omitempty"`
		// An array of relationships to hyperflexHxdpVersion resources.
		HxdpVersions []HyperflexHxdpVersionRelationship `json:"HxdpVersions,omitempty"`
		// An array of relationships to hyperflexCapabilityInfo resources.
		HyperflexCapabilityInfos []HyperflexCapabilityInfoRelationship `json:"HyperflexCapabilityInfos,omitempty"`
		// An array of relationships to hclHyperflexSoftwareCompatibilityInfo resources.
		HyperflexSoftwareCompatibilityInfos []HclHyperflexSoftwareCompatibilityInfoRelationship `json:"HyperflexSoftwareCompatibilityInfos,omitempty"`
		ServerFirmwareVersion               *HyperflexServerFirmwareVersionRelationship         `json:"ServerFirmwareVersion,omitempty"`
		ServerModel                         *HyperflexServerModelRelationship                   `json:"ServerModel,omitempty"`
		// An array of relationships to hyperflexSoftwareDistributionEntry resources.
		SoftwareDistributions []HyperflexSoftwareDistributionEntryRelationship `json:"SoftwareDistributions,omitempty"`
	}

	varHyperflexAppCatalogWithoutEmbeddedStruct := HyperflexAppCatalogWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexAppCatalogWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexAppCatalog := _HyperflexAppCatalog{}
		varHyperflexAppCatalog.ClassId = varHyperflexAppCatalogWithoutEmbeddedStruct.ClassId
		varHyperflexAppCatalog.ObjectType = varHyperflexAppCatalogWithoutEmbeddedStruct.ObjectType
		varHyperflexAppCatalog.Version = varHyperflexAppCatalogWithoutEmbeddedStruct.Version
		varHyperflexAppCatalog.FeatureLimitExternal = varHyperflexAppCatalogWithoutEmbeddedStruct.FeatureLimitExternal
		varHyperflexAppCatalog.FeatureLimitInternal = varHyperflexAppCatalogWithoutEmbeddedStruct.FeatureLimitInternal
		varHyperflexAppCatalog.HxdpVersions = varHyperflexAppCatalogWithoutEmbeddedStruct.HxdpVersions
		varHyperflexAppCatalog.HyperflexCapabilityInfos = varHyperflexAppCatalogWithoutEmbeddedStruct.HyperflexCapabilityInfos
		varHyperflexAppCatalog.HyperflexSoftwareCompatibilityInfos = varHyperflexAppCatalogWithoutEmbeddedStruct.HyperflexSoftwareCompatibilityInfos
		varHyperflexAppCatalog.ServerFirmwareVersion = varHyperflexAppCatalogWithoutEmbeddedStruct.ServerFirmwareVersion
		varHyperflexAppCatalog.ServerModel = varHyperflexAppCatalogWithoutEmbeddedStruct.ServerModel
		varHyperflexAppCatalog.SoftwareDistributions = varHyperflexAppCatalogWithoutEmbeddedStruct.SoftwareDistributions
		*o = HyperflexAppCatalog(varHyperflexAppCatalog)
	} else {
		return err
	}

	varHyperflexAppCatalog := _HyperflexAppCatalog{}

	err = json.Unmarshal(bytes, &varHyperflexAppCatalog)
	if err == nil {
		o.MoBaseMo = varHyperflexAppCatalog.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "FeatureLimitExternal")
		delete(additionalProperties, "FeatureLimitInternal")
		delete(additionalProperties, "HxdpVersions")
		delete(additionalProperties, "HyperflexCapabilityInfos")
		delete(additionalProperties, "HyperflexSoftwareCompatibilityInfos")
		delete(additionalProperties, "ServerFirmwareVersion")
		delete(additionalProperties, "ServerModel")
		delete(additionalProperties, "SoftwareDistributions")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexAppCatalog struct {
	value *HyperflexAppCatalog
	isSet bool
}

func (v NullableHyperflexAppCatalog) Get() *HyperflexAppCatalog {
	return v.value
}

func (v *NullableHyperflexAppCatalog) Set(val *HyperflexAppCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAppCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAppCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAppCatalog(val *HyperflexAppCatalog) *NullableHyperflexAppCatalog {
	return &NullableHyperflexAppCatalog{value: val, isSet: true}
}

func (v NullableHyperflexAppCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAppCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
