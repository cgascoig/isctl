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

// FabricSwitchClusterProfile This specifies the configuration policies for a cluster of switches.
type FabricSwitchClusterProfile struct {
	PolicyAbstractProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                      `json:"ObjectType"`
	ConfigContext NullablePolicyConfigContext `json:"ConfigContext,omitempty"`
	// Number of switch profiles that are part of this cluster profile.
	SwitchProfilesCount *int64                                `json:"SwitchProfilesCount,omitempty"`
	Organization        *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to fabricSwitchProfile resources.
	SwitchProfiles       []FabricSwitchProfileRelationship `json:"SwitchProfiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricSwitchClusterProfile FabricSwitchClusterProfile

// NewFabricSwitchClusterProfile instantiates a new FabricSwitchClusterProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSwitchClusterProfile(classId string, objectType string) *FabricSwitchClusterProfile {
	this := FabricSwitchClusterProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	return &this
}

// NewFabricSwitchClusterProfileWithDefaults instantiates a new FabricSwitchClusterProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSwitchClusterProfileWithDefaults() *FabricSwitchClusterProfile {
	this := FabricSwitchClusterProfile{}
	var classId string = "fabric.SwitchClusterProfile"
	this.ClassId = classId
	var objectType string = "fabric.SwitchClusterProfile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricSwitchClusterProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricSwitchClusterProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricSwitchClusterProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricSwitchClusterProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricSwitchClusterProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricSwitchClusterProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigContext returns the ConfigContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchClusterProfile) GetConfigContext() PolicyConfigContext {
	if o == nil || o.ConfigContext.Get() == nil {
		var ret PolicyConfigContext
		return ret
	}
	return *o.ConfigContext.Get()
}

// GetConfigContextOk returns a tuple with the ConfigContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchClusterProfile) GetConfigContextOk() (*PolicyConfigContext, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigContext.Get(), o.ConfigContext.IsSet()
}

// HasConfigContext returns a boolean if a field has been set.
func (o *FabricSwitchClusterProfile) HasConfigContext() bool {
	if o != nil && o.ConfigContext.IsSet() {
		return true
	}

	return false
}

// SetConfigContext gets a reference to the given NullablePolicyConfigContext and assigns it to the ConfigContext field.
func (o *FabricSwitchClusterProfile) SetConfigContext(v PolicyConfigContext) {
	o.ConfigContext.Set(&v)
}

// SetConfigContextNil sets the value for ConfigContext to be an explicit nil
func (o *FabricSwitchClusterProfile) SetConfigContextNil() {
	o.ConfigContext.Set(nil)
}

// UnsetConfigContext ensures that no value is present for ConfigContext, not even an explicit nil
func (o *FabricSwitchClusterProfile) UnsetConfigContext() {
	o.ConfigContext.Unset()
}

// GetSwitchProfilesCount returns the SwitchProfilesCount field value if set, zero value otherwise.
func (o *FabricSwitchClusterProfile) GetSwitchProfilesCount() int64 {
	if o == nil || o.SwitchProfilesCount == nil {
		var ret int64
		return ret
	}
	return *o.SwitchProfilesCount
}

// GetSwitchProfilesCountOk returns a tuple with the SwitchProfilesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchClusterProfile) GetSwitchProfilesCountOk() (*int64, bool) {
	if o == nil || o.SwitchProfilesCount == nil {
		return nil, false
	}
	return o.SwitchProfilesCount, true
}

// HasSwitchProfilesCount returns a boolean if a field has been set.
func (o *FabricSwitchClusterProfile) HasSwitchProfilesCount() bool {
	if o != nil && o.SwitchProfilesCount != nil {
		return true
	}

	return false
}

// SetSwitchProfilesCount gets a reference to the given int64 and assigns it to the SwitchProfilesCount field.
func (o *FabricSwitchClusterProfile) SetSwitchProfilesCount(v int64) {
	o.SwitchProfilesCount = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricSwitchClusterProfile) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchClusterProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricSwitchClusterProfile) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricSwitchClusterProfile) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetSwitchProfiles returns the SwitchProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchClusterProfile) GetSwitchProfiles() []FabricSwitchProfileRelationship {
	if o == nil {
		var ret []FabricSwitchProfileRelationship
		return ret
	}
	return o.SwitchProfiles
}

// GetSwitchProfilesOk returns a tuple with the SwitchProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchClusterProfile) GetSwitchProfilesOk() (*[]FabricSwitchProfileRelationship, bool) {
	if o == nil || o.SwitchProfiles == nil {
		return nil, false
	}
	return &o.SwitchProfiles, true
}

// HasSwitchProfiles returns a boolean if a field has been set.
func (o *FabricSwitchClusterProfile) HasSwitchProfiles() bool {
	if o != nil && o.SwitchProfiles != nil {
		return true
	}

	return false
}

// SetSwitchProfiles gets a reference to the given []FabricSwitchProfileRelationship and assigns it to the SwitchProfiles field.
func (o *FabricSwitchClusterProfile) SetSwitchProfiles(v []FabricSwitchProfileRelationship) {
	o.SwitchProfiles = v
}

func (o FabricSwitchClusterProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractProfile, errPolicyAbstractProfile := json.Marshal(o.PolicyAbstractProfile)
	if errPolicyAbstractProfile != nil {
		return []byte{}, errPolicyAbstractProfile
	}
	errPolicyAbstractProfile = json.Unmarshal([]byte(serializedPolicyAbstractProfile), &toSerialize)
	if errPolicyAbstractProfile != nil {
		return []byte{}, errPolicyAbstractProfile
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigContext.IsSet() {
		toSerialize["ConfigContext"] = o.ConfigContext.Get()
	}
	if o.SwitchProfilesCount != nil {
		toSerialize["SwitchProfilesCount"] = o.SwitchProfilesCount
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.SwitchProfiles != nil {
		toSerialize["SwitchProfiles"] = o.SwitchProfiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricSwitchClusterProfile) UnmarshalJSON(bytes []byte) (err error) {
	type FabricSwitchClusterProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType    string                      `json:"ObjectType"`
		ConfigContext NullablePolicyConfigContext `json:"ConfigContext,omitempty"`
		// Number of switch profiles that are part of this cluster profile.
		SwitchProfilesCount *int64                                `json:"SwitchProfilesCount,omitempty"`
		Organization        *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to fabricSwitchProfile resources.
		SwitchProfiles []FabricSwitchProfileRelationship `json:"SwitchProfiles,omitempty"`
	}

	varFabricSwitchClusterProfileWithoutEmbeddedStruct := FabricSwitchClusterProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricSwitchClusterProfileWithoutEmbeddedStruct)
	if err == nil {
		varFabricSwitchClusterProfile := _FabricSwitchClusterProfile{}
		varFabricSwitchClusterProfile.ClassId = varFabricSwitchClusterProfileWithoutEmbeddedStruct.ClassId
		varFabricSwitchClusterProfile.ObjectType = varFabricSwitchClusterProfileWithoutEmbeddedStruct.ObjectType
		varFabricSwitchClusterProfile.ConfigContext = varFabricSwitchClusterProfileWithoutEmbeddedStruct.ConfigContext
		varFabricSwitchClusterProfile.SwitchProfilesCount = varFabricSwitchClusterProfileWithoutEmbeddedStruct.SwitchProfilesCount
		varFabricSwitchClusterProfile.Organization = varFabricSwitchClusterProfileWithoutEmbeddedStruct.Organization
		varFabricSwitchClusterProfile.SwitchProfiles = varFabricSwitchClusterProfileWithoutEmbeddedStruct.SwitchProfiles
		*o = FabricSwitchClusterProfile(varFabricSwitchClusterProfile)
	} else {
		return err
	}

	varFabricSwitchClusterProfile := _FabricSwitchClusterProfile{}

	err = json.Unmarshal(bytes, &varFabricSwitchClusterProfile)
	if err == nil {
		o.PolicyAbstractProfile = varFabricSwitchClusterProfile.PolicyAbstractProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigContext")
		delete(additionalProperties, "SwitchProfilesCount")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "SwitchProfiles")

		// remove fields from embedded structs
		reflectPolicyAbstractProfile := reflect.ValueOf(o.PolicyAbstractProfile)
		for i := 0; i < reflectPolicyAbstractProfile.Type().NumField(); i++ {
			t := reflectPolicyAbstractProfile.Type().Field(i)

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

type NullableFabricSwitchClusterProfile struct {
	value *FabricSwitchClusterProfile
	isSet bool
}

func (v NullableFabricSwitchClusterProfile) Get() *FabricSwitchClusterProfile {
	return v.value
}

func (v *NullableFabricSwitchClusterProfile) Set(val *FabricSwitchClusterProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSwitchClusterProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSwitchClusterProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSwitchClusterProfile(val *FabricSwitchClusterProfile) *NullableFabricSwitchClusterProfile {
	return &NullableFabricSwitchClusterProfile{value: val, isSet: true}
}

func (v NullableFabricSwitchClusterProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSwitchClusterProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
