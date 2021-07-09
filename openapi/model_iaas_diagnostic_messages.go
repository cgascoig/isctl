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

// IaasDiagnosticMessages Gets diagnostics messages from UCSD.
type IaasDiagnosticMessages struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Message category of the alerts.
	Category *string `json:"Category,omitempty"`
	// Unique ID of UCS Director getting registerd with Intersight.
	Guid *string `json:"Guid,omitempty"`
	// Message target type of the alerts.
	Item *string `json:"Item,omitempty"`
	// Last checked time of the alerts.
	LastChecked *string `json:"LastChecked,omitempty"`
	// Detailed info about the alert.
	Message *string `json:"Message,omitempty"`
	// Recommended fix for the alert.
	Recommendation *string `json:"Recommendation,omitempty"`
	// Status of the given alert.
	Status *string `json:"Status,omitempty"`
	// Status Id of the given alert.
	StatusId             *string                              `json:"StatusId,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasDiagnosticMessages IaasDiagnosticMessages

// NewIaasDiagnosticMessages instantiates a new IaasDiagnosticMessages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasDiagnosticMessages(classId string, objectType string) *IaasDiagnosticMessages {
	this := IaasDiagnosticMessages{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasDiagnosticMessagesWithDefaults instantiates a new IaasDiagnosticMessages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasDiagnosticMessagesWithDefaults() *IaasDiagnosticMessages {
	this := IaasDiagnosticMessages{}
	var classId string = "iaas.DiagnosticMessages"
	this.ClassId = classId
	var objectType string = "iaas.DiagnosticMessages"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasDiagnosticMessages) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasDiagnosticMessages) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasDiagnosticMessages) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IaasDiagnosticMessages) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasDiagnosticMessages) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasDiagnosticMessages) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *IaasDiagnosticMessages) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDiagnosticMessages) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *IaasDiagnosticMessages) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *IaasDiagnosticMessages) SetCategory(v string) {
	o.Category = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *IaasDiagnosticMessages) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDiagnosticMessages) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *IaasDiagnosticMessages) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *IaasDiagnosticMessages) SetGuid(v string) {
	o.Guid = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *IaasDiagnosticMessages) GetItem() string {
	if o == nil || o.Item == nil {
		var ret string
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDiagnosticMessages) GetItemOk() (*string, bool) {
	if o == nil || o.Item == nil {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *IaasDiagnosticMessages) HasItem() bool {
	if o != nil && o.Item != nil {
		return true
	}

	return false
}

// SetItem gets a reference to the given string and assigns it to the Item field.
func (o *IaasDiagnosticMessages) SetItem(v string) {
	o.Item = &v
}

// GetLastChecked returns the LastChecked field value if set, zero value otherwise.
func (o *IaasDiagnosticMessages) GetLastChecked() string {
	if o == nil || o.LastChecked == nil {
		var ret string
		return ret
	}
	return *o.LastChecked
}

// GetLastCheckedOk returns a tuple with the LastChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDiagnosticMessages) GetLastCheckedOk() (*string, bool) {
	if o == nil || o.LastChecked == nil {
		return nil, false
	}
	return o.LastChecked, true
}

// HasLastChecked returns a boolean if a field has been set.
func (o *IaasDiagnosticMessages) HasLastChecked() bool {
	if o != nil && o.LastChecked != nil {
		return true
	}

	return false
}

// SetLastChecked gets a reference to the given string and assigns it to the LastChecked field.
func (o *IaasDiagnosticMessages) SetLastChecked(v string) {
	o.LastChecked = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *IaasDiagnosticMessages) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDiagnosticMessages) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *IaasDiagnosticMessages) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *IaasDiagnosticMessages) SetMessage(v string) {
	o.Message = &v
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise.
func (o *IaasDiagnosticMessages) GetRecommendation() string {
	if o == nil || o.Recommendation == nil {
		var ret string
		return ret
	}
	return *o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDiagnosticMessages) GetRecommendationOk() (*string, bool) {
	if o == nil || o.Recommendation == nil {
		return nil, false
	}
	return o.Recommendation, true
}

// HasRecommendation returns a boolean if a field has been set.
func (o *IaasDiagnosticMessages) HasRecommendation() bool {
	if o != nil && o.Recommendation != nil {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given string and assigns it to the Recommendation field.
func (o *IaasDiagnosticMessages) SetRecommendation(v string) {
	o.Recommendation = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IaasDiagnosticMessages) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDiagnosticMessages) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IaasDiagnosticMessages) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IaasDiagnosticMessages) SetStatus(v string) {
	o.Status = &v
}

// GetStatusId returns the StatusId field value if set, zero value otherwise.
func (o *IaasDiagnosticMessages) GetStatusId() string {
	if o == nil || o.StatusId == nil {
		var ret string
		return ret
	}
	return *o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDiagnosticMessages) GetStatusIdOk() (*string, bool) {
	if o == nil || o.StatusId == nil {
		return nil, false
	}
	return o.StatusId, true
}

// HasStatusId returns a boolean if a field has been set.
func (o *IaasDiagnosticMessages) HasStatusId() bool {
	if o != nil && o.StatusId != nil {
		return true
	}

	return false
}

// SetStatusId gets a reference to the given string and assigns it to the StatusId field.
func (o *IaasDiagnosticMessages) SetStatusId(v string) {
	o.StatusId = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *IaasDiagnosticMessages) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDiagnosticMessages) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *IaasDiagnosticMessages) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *IaasDiagnosticMessages) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o IaasDiagnosticMessages) MarshalJSON() ([]byte, error) {
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
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.Guid != nil {
		toSerialize["Guid"] = o.Guid
	}
	if o.Item != nil {
		toSerialize["Item"] = o.Item
	}
	if o.LastChecked != nil {
		toSerialize["LastChecked"] = o.LastChecked
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Recommendation != nil {
		toSerialize["Recommendation"] = o.Recommendation
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusId != nil {
		toSerialize["StatusId"] = o.StatusId
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IaasDiagnosticMessages) UnmarshalJSON(bytes []byte) (err error) {
	type IaasDiagnosticMessagesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Message category of the alerts.
		Category *string `json:"Category,omitempty"`
		// Unique ID of UCS Director getting registerd with Intersight.
		Guid *string `json:"Guid,omitempty"`
		// Message target type of the alerts.
		Item *string `json:"Item,omitempty"`
		// Last checked time of the alerts.
		LastChecked *string `json:"LastChecked,omitempty"`
		// Detailed info about the alert.
		Message *string `json:"Message,omitempty"`
		// Recommended fix for the alert.
		Recommendation *string `json:"Recommendation,omitempty"`
		// Status of the given alert.
		Status *string `json:"Status,omitempty"`
		// Status Id of the given alert.
		StatusId         *string                              `json:"StatusId,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varIaasDiagnosticMessagesWithoutEmbeddedStruct := IaasDiagnosticMessagesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIaasDiagnosticMessagesWithoutEmbeddedStruct)
	if err == nil {
		varIaasDiagnosticMessages := _IaasDiagnosticMessages{}
		varIaasDiagnosticMessages.ClassId = varIaasDiagnosticMessagesWithoutEmbeddedStruct.ClassId
		varIaasDiagnosticMessages.ObjectType = varIaasDiagnosticMessagesWithoutEmbeddedStruct.ObjectType
		varIaasDiagnosticMessages.Category = varIaasDiagnosticMessagesWithoutEmbeddedStruct.Category
		varIaasDiagnosticMessages.Guid = varIaasDiagnosticMessagesWithoutEmbeddedStruct.Guid
		varIaasDiagnosticMessages.Item = varIaasDiagnosticMessagesWithoutEmbeddedStruct.Item
		varIaasDiagnosticMessages.LastChecked = varIaasDiagnosticMessagesWithoutEmbeddedStruct.LastChecked
		varIaasDiagnosticMessages.Message = varIaasDiagnosticMessagesWithoutEmbeddedStruct.Message
		varIaasDiagnosticMessages.Recommendation = varIaasDiagnosticMessagesWithoutEmbeddedStruct.Recommendation
		varIaasDiagnosticMessages.Status = varIaasDiagnosticMessagesWithoutEmbeddedStruct.Status
		varIaasDiagnosticMessages.StatusId = varIaasDiagnosticMessagesWithoutEmbeddedStruct.StatusId
		varIaasDiagnosticMessages.RegisteredDevice = varIaasDiagnosticMessagesWithoutEmbeddedStruct.RegisteredDevice
		*o = IaasDiagnosticMessages(varIaasDiagnosticMessages)
	} else {
		return err
	}

	varIaasDiagnosticMessages := _IaasDiagnosticMessages{}

	err = json.Unmarshal(bytes, &varIaasDiagnosticMessages)
	if err == nil {
		o.MoBaseMo = varIaasDiagnosticMessages.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "Guid")
		delete(additionalProperties, "Item")
		delete(additionalProperties, "LastChecked")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Recommendation")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusId")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableIaasDiagnosticMessages struct {
	value *IaasDiagnosticMessages
	isSet bool
}

func (v NullableIaasDiagnosticMessages) Get() *IaasDiagnosticMessages {
	return v.value
}

func (v *NullableIaasDiagnosticMessages) Set(val *IaasDiagnosticMessages) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasDiagnosticMessages) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasDiagnosticMessages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasDiagnosticMessages(val *IaasDiagnosticMessages) *NullableIaasDiagnosticMessages {
	return &NullableIaasDiagnosticMessages{value: val, isSet: true}
}

func (v NullableIaasDiagnosticMessages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasDiagnosticMessages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
