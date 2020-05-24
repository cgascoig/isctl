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

// PolicyConfigResultContext The context for a validation/config result such as the related entity's name, type, MOID etc.
type PolicyConfigResultContext struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// The data of the object present in config result context.
	EntityData *map[string]interface{} `json:"EntityData,omitempty" yaml:"EntityData,omitempty"`
	// The Moid of the object present in config result context.
	EntityMoid *string `json:"EntityMoid,omitempty" yaml:"EntityMoid,omitempty"`
	// The name of the object present in config result context.
	EntityName *string `json:"EntityName,omitempty" yaml:"EntityName,omitempty"`
	// The type of the object present in config result context.
	EntityType *string `json:"EntityType,omitempty" yaml:"EntityType,omitempty"`
}

// NewPolicyConfigResultContext instantiates a new PolicyConfigResultContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyConfigResultContext() *PolicyConfigResultContext {
	this := PolicyConfigResultContext{}
	return &this
}

// NewPolicyConfigResultContextWithDefaults instantiates a new PolicyConfigResultContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyConfigResultContextWithDefaults() *PolicyConfigResultContext {
	this := PolicyConfigResultContext{}
	return &this
}

// GetEntityData returns the EntityData field value if set, zero value otherwise.
func (o *PolicyConfigResultContext) GetEntityData() map[string]interface{} {
	if o == nil || o.EntityData == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.EntityData
}

// GetEntityDataOk returns a tuple with the EntityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContext) GetEntityDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.EntityData == nil {
		return nil, false
	}
	return o.EntityData, true
}

// HasEntityData returns a boolean if a field has been set.
func (o *PolicyConfigResultContext) HasEntityData() bool {
	if o != nil && o.EntityData != nil {
		return true
	}

	return false
}

// SetEntityData gets a reference to the given map[string]interface{} and assigns it to the EntityData field.
func (o *PolicyConfigResultContext) SetEntityData(v map[string]interface{}) {
	o.EntityData = &v
}

// GetEntityMoid returns the EntityMoid field value if set, zero value otherwise.
func (o *PolicyConfigResultContext) GetEntityMoid() string {
	if o == nil || o.EntityMoid == nil {
		var ret string
		return ret
	}
	return *o.EntityMoid
}

// GetEntityMoidOk returns a tuple with the EntityMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContext) GetEntityMoidOk() (*string, bool) {
	if o == nil || o.EntityMoid == nil {
		return nil, false
	}
	return o.EntityMoid, true
}

// HasEntityMoid returns a boolean if a field has been set.
func (o *PolicyConfigResultContext) HasEntityMoid() bool {
	if o != nil && o.EntityMoid != nil {
		return true
	}

	return false
}

// SetEntityMoid gets a reference to the given string and assigns it to the EntityMoid field.
func (o *PolicyConfigResultContext) SetEntityMoid(v string) {
	o.EntityMoid = &v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise.
func (o *PolicyConfigResultContext) GetEntityName() string {
	if o == nil || o.EntityName == nil {
		var ret string
		return ret
	}
	return *o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContext) GetEntityNameOk() (*string, bool) {
	if o == nil || o.EntityName == nil {
		return nil, false
	}
	return o.EntityName, true
}

// HasEntityName returns a boolean if a field has been set.
func (o *PolicyConfigResultContext) HasEntityName() bool {
	if o != nil && o.EntityName != nil {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given string and assigns it to the EntityName field.
func (o *PolicyConfigResultContext) SetEntityName(v string) {
	o.EntityName = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *PolicyConfigResultContext) GetEntityType() string {
	if o == nil || o.EntityType == nil {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContext) GetEntityTypeOk() (*string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *PolicyConfigResultContext) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *PolicyConfigResultContext) SetEntityType(v string) {
	o.EntityType = &v
}

func (o PolicyConfigResultContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.EntityData != nil {
		toSerialize["EntityData"] = o.EntityData
	}
	if o.EntityMoid != nil {
		toSerialize["EntityMoid"] = o.EntityMoid
	}
	if o.EntityName != nil {
		toSerialize["EntityName"] = o.EntityName
	}
	if o.EntityType != nil {
		toSerialize["EntityType"] = o.EntityType
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyConfigResultContext struct {
	value *PolicyConfigResultContext
	isSet bool
}

func (v NullablePolicyConfigResultContext) Get() *PolicyConfigResultContext {
	return v.value
}

func (v *NullablePolicyConfigResultContext) Set(val *PolicyConfigResultContext) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyConfigResultContext) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyConfigResultContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyConfigResultContext(val *PolicyConfigResultContext) *NullablePolicyConfigResultContext {
	return &NullablePolicyConfigResultContext{value: val, isSet: true}
}

func (v NullablePolicyConfigResultContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyConfigResultContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
