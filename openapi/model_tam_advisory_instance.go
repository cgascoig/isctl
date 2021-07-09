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
	"time"
)

// TamAdvisoryInstance Instance of an Intersight advisory applicable for an Intersight managed object. An advisory instance is created when a given advisory is found applicable for an Intersight managed object. An advisory instance is retained for some time even after being cleared for historical purposes. A 'cleared' advisory instance is deleted after the retention time is elaspsed.
type TamAdvisoryInstance struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Moid of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases.
	AffectedObjectMoid *string `json:"AffectedObjectMoid,omitempty"`
	// Object type of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases.
	AffectedObjectType *string `json:"AffectedObjectType,omitempty"`
	// Timestamp when a state change was observed on this advisory instnace.
	LastStateChangeTime *time.Time `json:"LastStateChangeTime,omitempty"`
	// Timestamp when this advisory was last evaluated.
	LastVerifiedTime *time.Time `json:"LastVerifiedTime,omitempty"`
	// Current state of the advisory instance (Active/Cleared/Unknown etc.). * `unknown` - Intersight is unable to determine if the Advisory instance is applicable for the affected managed object. * `active` - Advisory instance is currently active and applicable for the affected managed object. * `cleared` - Advisory instance is no longer applicable for the affected managed object.
	State                *string                              `json:"State,omitempty"`
	Advisory             *TamBaseAdvisoryRelationship         `json:"Advisory,omitempty"`
	AffectedObject       *MoBaseMoRelationship                `json:"AffectedObject,omitempty"`
	DeviceRegistration   *AssetDeviceRegistrationRelationship `json:"DeviceRegistration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamAdvisoryInstance TamAdvisoryInstance

// NewTamAdvisoryInstance instantiates a new TamAdvisoryInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamAdvisoryInstance(classId string, objectType string) *TamAdvisoryInstance {
	this := TamAdvisoryInstance{}
	this.ClassId = classId
	this.ObjectType = objectType
	var state string = "unknown"
	this.State = &state
	return &this
}

// NewTamAdvisoryInstanceWithDefaults instantiates a new TamAdvisoryInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamAdvisoryInstanceWithDefaults() *TamAdvisoryInstance {
	this := TamAdvisoryInstance{}
	var classId string = "tam.AdvisoryInstance"
	this.ClassId = classId
	var objectType string = "tam.AdvisoryInstance"
	this.ObjectType = objectType
	var state string = "unknown"
	this.State = &state
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamAdvisoryInstance) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstance) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamAdvisoryInstance) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamAdvisoryInstance) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstance) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamAdvisoryInstance) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAffectedObjectMoid returns the AffectedObjectMoid field value if set, zero value otherwise.
func (o *TamAdvisoryInstance) GetAffectedObjectMoid() string {
	if o == nil || o.AffectedObjectMoid == nil {
		var ret string
		return ret
	}
	return *o.AffectedObjectMoid
}

// GetAffectedObjectMoidOk returns a tuple with the AffectedObjectMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstance) GetAffectedObjectMoidOk() (*string, bool) {
	if o == nil || o.AffectedObjectMoid == nil {
		return nil, false
	}
	return o.AffectedObjectMoid, true
}

// HasAffectedObjectMoid returns a boolean if a field has been set.
func (o *TamAdvisoryInstance) HasAffectedObjectMoid() bool {
	if o != nil && o.AffectedObjectMoid != nil {
		return true
	}

	return false
}

// SetAffectedObjectMoid gets a reference to the given string and assigns it to the AffectedObjectMoid field.
func (o *TamAdvisoryInstance) SetAffectedObjectMoid(v string) {
	o.AffectedObjectMoid = &v
}

// GetAffectedObjectType returns the AffectedObjectType field value if set, zero value otherwise.
func (o *TamAdvisoryInstance) GetAffectedObjectType() string {
	if o == nil || o.AffectedObjectType == nil {
		var ret string
		return ret
	}
	return *o.AffectedObjectType
}

// GetAffectedObjectTypeOk returns a tuple with the AffectedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstance) GetAffectedObjectTypeOk() (*string, bool) {
	if o == nil || o.AffectedObjectType == nil {
		return nil, false
	}
	return o.AffectedObjectType, true
}

// HasAffectedObjectType returns a boolean if a field has been set.
func (o *TamAdvisoryInstance) HasAffectedObjectType() bool {
	if o != nil && o.AffectedObjectType != nil {
		return true
	}

	return false
}

// SetAffectedObjectType gets a reference to the given string and assigns it to the AffectedObjectType field.
func (o *TamAdvisoryInstance) SetAffectedObjectType(v string) {
	o.AffectedObjectType = &v
}

// GetLastStateChangeTime returns the LastStateChangeTime field value if set, zero value otherwise.
func (o *TamAdvisoryInstance) GetLastStateChangeTime() time.Time {
	if o == nil || o.LastStateChangeTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastStateChangeTime
}

// GetLastStateChangeTimeOk returns a tuple with the LastStateChangeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstance) GetLastStateChangeTimeOk() (*time.Time, bool) {
	if o == nil || o.LastStateChangeTime == nil {
		return nil, false
	}
	return o.LastStateChangeTime, true
}

// HasLastStateChangeTime returns a boolean if a field has been set.
func (o *TamAdvisoryInstance) HasLastStateChangeTime() bool {
	if o != nil && o.LastStateChangeTime != nil {
		return true
	}

	return false
}

// SetLastStateChangeTime gets a reference to the given time.Time and assigns it to the LastStateChangeTime field.
func (o *TamAdvisoryInstance) SetLastStateChangeTime(v time.Time) {
	o.LastStateChangeTime = &v
}

// GetLastVerifiedTime returns the LastVerifiedTime field value if set, zero value otherwise.
func (o *TamAdvisoryInstance) GetLastVerifiedTime() time.Time {
	if o == nil || o.LastVerifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastVerifiedTime
}

// GetLastVerifiedTimeOk returns a tuple with the LastVerifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstance) GetLastVerifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastVerifiedTime == nil {
		return nil, false
	}
	return o.LastVerifiedTime, true
}

// HasLastVerifiedTime returns a boolean if a field has been set.
func (o *TamAdvisoryInstance) HasLastVerifiedTime() bool {
	if o != nil && o.LastVerifiedTime != nil {
		return true
	}

	return false
}

// SetLastVerifiedTime gets a reference to the given time.Time and assigns it to the LastVerifiedTime field.
func (o *TamAdvisoryInstance) SetLastVerifiedTime(v time.Time) {
	o.LastVerifiedTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TamAdvisoryInstance) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstance) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TamAdvisoryInstance) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TamAdvisoryInstance) SetState(v string) {
	o.State = &v
}

// GetAdvisory returns the Advisory field value if set, zero value otherwise.
func (o *TamAdvisoryInstance) GetAdvisory() TamBaseAdvisoryRelationship {
	if o == nil || o.Advisory == nil {
		var ret TamBaseAdvisoryRelationship
		return ret
	}
	return *o.Advisory
}

// GetAdvisoryOk returns a tuple with the Advisory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstance) GetAdvisoryOk() (*TamBaseAdvisoryRelationship, bool) {
	if o == nil || o.Advisory == nil {
		return nil, false
	}
	return o.Advisory, true
}

// HasAdvisory returns a boolean if a field has been set.
func (o *TamAdvisoryInstance) HasAdvisory() bool {
	if o != nil && o.Advisory != nil {
		return true
	}

	return false
}

// SetAdvisory gets a reference to the given TamBaseAdvisoryRelationship and assigns it to the Advisory field.
func (o *TamAdvisoryInstance) SetAdvisory(v TamBaseAdvisoryRelationship) {
	o.Advisory = &v
}

// GetAffectedObject returns the AffectedObject field value if set, zero value otherwise.
func (o *TamAdvisoryInstance) GetAffectedObject() MoBaseMoRelationship {
	if o == nil || o.AffectedObject == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AffectedObject
}

// GetAffectedObjectOk returns a tuple with the AffectedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstance) GetAffectedObjectOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AffectedObject == nil {
		return nil, false
	}
	return o.AffectedObject, true
}

// HasAffectedObject returns a boolean if a field has been set.
func (o *TamAdvisoryInstance) HasAffectedObject() bool {
	if o != nil && o.AffectedObject != nil {
		return true
	}

	return false
}

// SetAffectedObject gets a reference to the given MoBaseMoRelationship and assigns it to the AffectedObject field.
func (o *TamAdvisoryInstance) SetAffectedObject(v MoBaseMoRelationship) {
	o.AffectedObject = &v
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise.
func (o *TamAdvisoryInstance) GetDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceRegistration
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstance) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceRegistration == nil {
		return nil, false
	}
	return o.DeviceRegistration, true
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *TamAdvisoryInstance) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration != nil {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceRegistration field.
func (o *TamAdvisoryInstance) SetDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.DeviceRegistration = &v
}

func (o TamAdvisoryInstance) MarshalJSON() ([]byte, error) {
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
	if o.AffectedObjectMoid != nil {
		toSerialize["AffectedObjectMoid"] = o.AffectedObjectMoid
	}
	if o.AffectedObjectType != nil {
		toSerialize["AffectedObjectType"] = o.AffectedObjectType
	}
	if o.LastStateChangeTime != nil {
		toSerialize["LastStateChangeTime"] = o.LastStateChangeTime
	}
	if o.LastVerifiedTime != nil {
		toSerialize["LastVerifiedTime"] = o.LastVerifiedTime
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Advisory != nil {
		toSerialize["Advisory"] = o.Advisory
	}
	if o.AffectedObject != nil {
		toSerialize["AffectedObject"] = o.AffectedObject
	}
	if o.DeviceRegistration != nil {
		toSerialize["DeviceRegistration"] = o.DeviceRegistration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamAdvisoryInstance) UnmarshalJSON(bytes []byte) (err error) {
	type TamAdvisoryInstanceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Moid of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases.
		AffectedObjectMoid *string `json:"AffectedObjectMoid,omitempty"`
		// Object type of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases.
		AffectedObjectType *string `json:"AffectedObjectType,omitempty"`
		// Timestamp when a state change was observed on this advisory instnace.
		LastStateChangeTime *time.Time `json:"LastStateChangeTime,omitempty"`
		// Timestamp when this advisory was last evaluated.
		LastVerifiedTime *time.Time `json:"LastVerifiedTime,omitempty"`
		// Current state of the advisory instance (Active/Cleared/Unknown etc.). * `unknown` - Intersight is unable to determine if the Advisory instance is applicable for the affected managed object. * `active` - Advisory instance is currently active and applicable for the affected managed object. * `cleared` - Advisory instance is no longer applicable for the affected managed object.
		State              *string                              `json:"State,omitempty"`
		Advisory           *TamBaseAdvisoryRelationship         `json:"Advisory,omitempty"`
		AffectedObject     *MoBaseMoRelationship                `json:"AffectedObject,omitempty"`
		DeviceRegistration *AssetDeviceRegistrationRelationship `json:"DeviceRegistration,omitempty"`
	}

	varTamAdvisoryInstanceWithoutEmbeddedStruct := TamAdvisoryInstanceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTamAdvisoryInstanceWithoutEmbeddedStruct)
	if err == nil {
		varTamAdvisoryInstance := _TamAdvisoryInstance{}
		varTamAdvisoryInstance.ClassId = varTamAdvisoryInstanceWithoutEmbeddedStruct.ClassId
		varTamAdvisoryInstance.ObjectType = varTamAdvisoryInstanceWithoutEmbeddedStruct.ObjectType
		varTamAdvisoryInstance.AffectedObjectMoid = varTamAdvisoryInstanceWithoutEmbeddedStruct.AffectedObjectMoid
		varTamAdvisoryInstance.AffectedObjectType = varTamAdvisoryInstanceWithoutEmbeddedStruct.AffectedObjectType
		varTamAdvisoryInstance.LastStateChangeTime = varTamAdvisoryInstanceWithoutEmbeddedStruct.LastStateChangeTime
		varTamAdvisoryInstance.LastVerifiedTime = varTamAdvisoryInstanceWithoutEmbeddedStruct.LastVerifiedTime
		varTamAdvisoryInstance.State = varTamAdvisoryInstanceWithoutEmbeddedStruct.State
		varTamAdvisoryInstance.Advisory = varTamAdvisoryInstanceWithoutEmbeddedStruct.Advisory
		varTamAdvisoryInstance.AffectedObject = varTamAdvisoryInstanceWithoutEmbeddedStruct.AffectedObject
		varTamAdvisoryInstance.DeviceRegistration = varTamAdvisoryInstanceWithoutEmbeddedStruct.DeviceRegistration
		*o = TamAdvisoryInstance(varTamAdvisoryInstance)
	} else {
		return err
	}

	varTamAdvisoryInstance := _TamAdvisoryInstance{}

	err = json.Unmarshal(bytes, &varTamAdvisoryInstance)
	if err == nil {
		o.MoBaseMo = varTamAdvisoryInstance.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AffectedObjectMoid")
		delete(additionalProperties, "AffectedObjectType")
		delete(additionalProperties, "LastStateChangeTime")
		delete(additionalProperties, "LastVerifiedTime")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Advisory")
		delete(additionalProperties, "AffectedObject")
		delete(additionalProperties, "DeviceRegistration")

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

type NullableTamAdvisoryInstance struct {
	value *TamAdvisoryInstance
	isSet bool
}

func (v NullableTamAdvisoryInstance) Get() *TamAdvisoryInstance {
	return v.value
}

func (v *NullableTamAdvisoryInstance) Set(val *TamAdvisoryInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableTamAdvisoryInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableTamAdvisoryInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamAdvisoryInstance(val *TamAdvisoryInstance) *NullableTamAdvisoryInstance {
	return &NullableTamAdvisoryInstance{value: val, isSet: true}
}

func (v NullableTamAdvisoryInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamAdvisoryInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
