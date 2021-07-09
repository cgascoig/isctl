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
)

// IamEndPointUserAllOf Definition of the list of properties defined in 'iam.EndPointUser', excluding properties defined in parent classes.
type IamEndPointUserAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the user to be created on the endpoint. It can be any string that adheres to the following constraints. It can have alphanumeric characters, dots, underscores and hyphen. It cannot be more than 16 characters.
	Name *string `json:"Name,omitempty"`
	// An array of relationships to iamEndPointUserRole resources.
	EndPointUserRole     []IamEndPointUserRoleRelationship     `json:"EndPointUserRole,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamEndPointUserAllOf IamEndPointUserAllOf

// NewIamEndPointUserAllOf instantiates a new IamEndPointUserAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamEndPointUserAllOf(classId string, objectType string) *IamEndPointUserAllOf {
	this := IamEndPointUserAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamEndPointUserAllOfWithDefaults instantiates a new IamEndPointUserAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamEndPointUserAllOfWithDefaults() *IamEndPointUserAllOf {
	this := IamEndPointUserAllOf{}
	var classId string = "iam.EndPointUser"
	this.ClassId = classId
	var objectType string = "iam.EndPointUser"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamEndPointUserAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamEndPointUserAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamEndPointUserAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamEndPointUserAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamEndPointUserAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamEndPointUserAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamEndPointUserAllOf) SetName(v string) {
	o.Name = &v
}

// GetEndPointUserRole returns the EndPointUserRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserAllOf) GetEndPointUserRole() []IamEndPointUserRoleRelationship {
	if o == nil {
		var ret []IamEndPointUserRoleRelationship
		return ret
	}
	return o.EndPointUserRole
}

// GetEndPointUserRoleOk returns a tuple with the EndPointUserRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserAllOf) GetEndPointUserRoleOk() (*[]IamEndPointUserRoleRelationship, bool) {
	if o == nil || o.EndPointUserRole == nil {
		return nil, false
	}
	return &o.EndPointUserRole, true
}

// HasEndPointUserRole returns a boolean if a field has been set.
func (o *IamEndPointUserAllOf) HasEndPointUserRole() bool {
	if o != nil && o.EndPointUserRole != nil {
		return true
	}

	return false
}

// SetEndPointUserRole gets a reference to the given []IamEndPointUserRoleRelationship and assigns it to the EndPointUserRole field.
func (o *IamEndPointUserAllOf) SetEndPointUserRole(v []IamEndPointUserRoleRelationship) {
	o.EndPointUserRole = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *IamEndPointUserAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *IamEndPointUserAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *IamEndPointUserAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o IamEndPointUserAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.EndPointUserRole != nil {
		toSerialize["EndPointUserRole"] = o.EndPointUserRole
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamEndPointUserAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamEndPointUserAllOf := _IamEndPointUserAllOf{}

	if err = json.Unmarshal(bytes, &varIamEndPointUserAllOf); err == nil {
		*o = IamEndPointUserAllOf(varIamEndPointUserAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "EndPointUserRole")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamEndPointUserAllOf struct {
	value *IamEndPointUserAllOf
	isSet bool
}

func (v NullableIamEndPointUserAllOf) Get() *IamEndPointUserAllOf {
	return v.value
}

func (v *NullableIamEndPointUserAllOf) Set(val *IamEndPointUserAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointUserAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointUserAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointUserAllOf(val *IamEndPointUserAllOf) *NullableIamEndPointUserAllOf {
	return &NullableIamEndPointUserAllOf{value: val, isSet: true}
}

func (v NullableIamEndPointUserAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointUserAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
