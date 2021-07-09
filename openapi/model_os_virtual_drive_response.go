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

// OsVirtualDriveResponse Virtual Drive target entry for the UI.
type OsVirtualDriveResponse struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Bootable field of the Virtual Drive target.
	Bootable *string `json:"Bootable,omitempty"`
	// Virtual Drive ID to be used as Install Target.
	Id *string `json:"Id,omitempty"`
	// The Virtual Drive Name to be used as Install Target.
	Name *string `json:"Name,omitempty"`
	// The Storage Controller associated to the virtual drive.
	StorageControllerSlotId *string `json:"StorageControllerSlotId,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _OsVirtualDriveResponse OsVirtualDriveResponse

// NewOsVirtualDriveResponse instantiates a new OsVirtualDriveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsVirtualDriveResponse(classId string, objectType string) *OsVirtualDriveResponse {
	this := OsVirtualDriveResponse{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsVirtualDriveResponseWithDefaults instantiates a new OsVirtualDriveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsVirtualDriveResponseWithDefaults() *OsVirtualDriveResponse {
	this := OsVirtualDriveResponse{}
	var classId string = "os.VirtualDriveResponse"
	this.ClassId = classId
	var objectType string = "os.VirtualDriveResponse"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsVirtualDriveResponse) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsVirtualDriveResponse) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsVirtualDriveResponse) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsVirtualDriveResponse) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsVirtualDriveResponse) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsVirtualDriveResponse) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *OsVirtualDriveResponse) GetBootable() string {
	if o == nil || o.Bootable == nil {
		var ret string
		return ret
	}
	return *o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsVirtualDriveResponse) GetBootableOk() (*string, bool) {
	if o == nil || o.Bootable == nil {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *OsVirtualDriveResponse) HasBootable() bool {
	if o != nil && o.Bootable != nil {
		return true
	}

	return false
}

// SetBootable gets a reference to the given string and assigns it to the Bootable field.
func (o *OsVirtualDriveResponse) SetBootable(v string) {
	o.Bootable = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OsVirtualDriveResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsVirtualDriveResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OsVirtualDriveResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OsVirtualDriveResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsVirtualDriveResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsVirtualDriveResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsVirtualDriveResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsVirtualDriveResponse) SetName(v string) {
	o.Name = &v
}

// GetStorageControllerSlotId returns the StorageControllerSlotId field value if set, zero value otherwise.
func (o *OsVirtualDriveResponse) GetStorageControllerSlotId() string {
	if o == nil || o.StorageControllerSlotId == nil {
		var ret string
		return ret
	}
	return *o.StorageControllerSlotId
}

// GetStorageControllerSlotIdOk returns a tuple with the StorageControllerSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsVirtualDriveResponse) GetStorageControllerSlotIdOk() (*string, bool) {
	if o == nil || o.StorageControllerSlotId == nil {
		return nil, false
	}
	return o.StorageControllerSlotId, true
}

// HasStorageControllerSlotId returns a boolean if a field has been set.
func (o *OsVirtualDriveResponse) HasStorageControllerSlotId() bool {
	if o != nil && o.StorageControllerSlotId != nil {
		return true
	}

	return false
}

// SetStorageControllerSlotId gets a reference to the given string and assigns it to the StorageControllerSlotId field.
func (o *OsVirtualDriveResponse) SetStorageControllerSlotId(v string) {
	o.StorageControllerSlotId = &v
}

func (o OsVirtualDriveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Bootable != nil {
		toSerialize["Bootable"] = o.Bootable
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.StorageControllerSlotId != nil {
		toSerialize["StorageControllerSlotId"] = o.StorageControllerSlotId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsVirtualDriveResponse) UnmarshalJSON(bytes []byte) (err error) {
	type OsVirtualDriveResponseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Bootable field of the Virtual Drive target.
		Bootable *string `json:"Bootable,omitempty"`
		// Virtual Drive ID to be used as Install Target.
		Id *string `json:"Id,omitempty"`
		// The Virtual Drive Name to be used as Install Target.
		Name *string `json:"Name,omitempty"`
		// The Storage Controller associated to the virtual drive.
		StorageControllerSlotId *string `json:"StorageControllerSlotId,omitempty"`
	}

	varOsVirtualDriveResponseWithoutEmbeddedStruct := OsVirtualDriveResponseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOsVirtualDriveResponseWithoutEmbeddedStruct)
	if err == nil {
		varOsVirtualDriveResponse := _OsVirtualDriveResponse{}
		varOsVirtualDriveResponse.ClassId = varOsVirtualDriveResponseWithoutEmbeddedStruct.ClassId
		varOsVirtualDriveResponse.ObjectType = varOsVirtualDriveResponseWithoutEmbeddedStruct.ObjectType
		varOsVirtualDriveResponse.Bootable = varOsVirtualDriveResponseWithoutEmbeddedStruct.Bootable
		varOsVirtualDriveResponse.Id = varOsVirtualDriveResponseWithoutEmbeddedStruct.Id
		varOsVirtualDriveResponse.Name = varOsVirtualDriveResponseWithoutEmbeddedStruct.Name
		varOsVirtualDriveResponse.StorageControllerSlotId = varOsVirtualDriveResponseWithoutEmbeddedStruct.StorageControllerSlotId
		*o = OsVirtualDriveResponse(varOsVirtualDriveResponse)
	} else {
		return err
	}

	varOsVirtualDriveResponse := _OsVirtualDriveResponse{}

	err = json.Unmarshal(bytes, &varOsVirtualDriveResponse)
	if err == nil {
		o.MoBaseComplexType = varOsVirtualDriveResponse.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bootable")
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "StorageControllerSlotId")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableOsVirtualDriveResponse struct {
	value *OsVirtualDriveResponse
	isSet bool
}

func (v NullableOsVirtualDriveResponse) Get() *OsVirtualDriveResponse {
	return v.value
}

func (v *NullableOsVirtualDriveResponse) Set(val *OsVirtualDriveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOsVirtualDriveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOsVirtualDriveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsVirtualDriveResponse(val *OsVirtualDriveResponse) *NullableOsVirtualDriveResponse {
	return &NullableOsVirtualDriveResponse{value: val, isSet: true}
}

func (v NullableOsVirtualDriveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsVirtualDriveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
