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

// StoragePurePort Port entity in Pure FlashArray.
type StoragePurePort struct {
	StorageBasePhysicalPort
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the port to which this port has failed over.
	Failover *string `json:"Failover,omitempty"`
	// Ip address of iSCSI portal configured on the port.
	Portal               *string                              `json:"Portal,omitempty"`
	Array                *StoragePureArrayRelationship        `json:"Array,omitempty"`
	Controller           *StoragePureControllerRelationship   `json:"Controller,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePurePort StoragePurePort

// NewStoragePurePort instantiates a new StoragePurePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePurePort(classId string, objectType string) *StoragePurePort {
	this := StoragePurePort{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Unknown"
	this.Status = &status
	var type_ string = "FC"
	this.Type = &type_
	return &this
}

// NewStoragePurePortWithDefaults instantiates a new StoragePurePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePurePortWithDefaults() *StoragePurePort {
	this := StoragePurePort{}
	var classId string = "storage.PurePort"
	this.ClassId = classId
	var objectType string = "storage.PurePort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePurePort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePurePort) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StoragePurePort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePurePort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFailover returns the Failover field value if set, zero value otherwise.
func (o *StoragePurePort) GetFailover() string {
	if o == nil || o.Failover == nil {
		var ret string
		return ret
	}
	return *o.Failover
}

// GetFailoverOk returns a tuple with the Failover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetFailoverOk() (*string, bool) {
	if o == nil || o.Failover == nil {
		return nil, false
	}
	return o.Failover, true
}

// HasFailover returns a boolean if a field has been set.
func (o *StoragePurePort) HasFailover() bool {
	if o != nil && o.Failover != nil {
		return true
	}

	return false
}

// SetFailover gets a reference to the given string and assigns it to the Failover field.
func (o *StoragePurePort) SetFailover(v string) {
	o.Failover = &v
}

// GetPortal returns the Portal field value if set, zero value otherwise.
func (o *StoragePurePort) GetPortal() string {
	if o == nil || o.Portal == nil {
		var ret string
		return ret
	}
	return *o.Portal
}

// GetPortalOk returns a tuple with the Portal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetPortalOk() (*string, bool) {
	if o == nil || o.Portal == nil {
		return nil, false
	}
	return o.Portal, true
}

// HasPortal returns a boolean if a field has been set.
func (o *StoragePurePort) HasPortal() bool {
	if o != nil && o.Portal != nil {
		return true
	}

	return false
}

// SetPortal gets a reference to the given string and assigns it to the Portal field.
func (o *StoragePurePort) SetPortal(v string) {
	o.Portal = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePurePort) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePurePort) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePurePort) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *StoragePurePort) GetController() StoragePureControllerRelationship {
	if o == nil || o.Controller == nil {
		var ret StoragePureControllerRelationship
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetControllerOk() (*StoragePureControllerRelationship, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *StoragePurePort) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given StoragePureControllerRelationship and assigns it to the Controller field.
func (o *StoragePurePort) SetController(v StoragePureControllerRelationship) {
	o.Controller = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePurePort) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePurePort) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePurePort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePurePort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBasePhysicalPort, errStorageBasePhysicalPort := json.Marshal(o.StorageBasePhysicalPort)
	if errStorageBasePhysicalPort != nil {
		return []byte{}, errStorageBasePhysicalPort
	}
	errStorageBasePhysicalPort = json.Unmarshal([]byte(serializedStorageBasePhysicalPort), &toSerialize)
	if errStorageBasePhysicalPort != nil {
		return []byte{}, errStorageBasePhysicalPort
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Failover != nil {
		toSerialize["Failover"] = o.Failover
	}
	if o.Portal != nil {
		toSerialize["Portal"] = o.Portal
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.Controller != nil {
		toSerialize["Controller"] = o.Controller
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePurePort) UnmarshalJSON(bytes []byte) (err error) {
	type StoragePurePortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the port to which this port has failed over.
		Failover *string `json:"Failover,omitempty"`
		// Ip address of iSCSI portal configured on the port.
		Portal           *string                              `json:"Portal,omitempty"`
		Array            *StoragePureArrayRelationship        `json:"Array,omitempty"`
		Controller       *StoragePureControllerRelationship   `json:"Controller,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStoragePurePortWithoutEmbeddedStruct := StoragePurePortWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStoragePurePortWithoutEmbeddedStruct)
	if err == nil {
		varStoragePurePort := _StoragePurePort{}
		varStoragePurePort.ClassId = varStoragePurePortWithoutEmbeddedStruct.ClassId
		varStoragePurePort.ObjectType = varStoragePurePortWithoutEmbeddedStruct.ObjectType
		varStoragePurePort.Failover = varStoragePurePortWithoutEmbeddedStruct.Failover
		varStoragePurePort.Portal = varStoragePurePortWithoutEmbeddedStruct.Portal
		varStoragePurePort.Array = varStoragePurePortWithoutEmbeddedStruct.Array
		varStoragePurePort.Controller = varStoragePurePortWithoutEmbeddedStruct.Controller
		varStoragePurePort.RegisteredDevice = varStoragePurePortWithoutEmbeddedStruct.RegisteredDevice
		*o = StoragePurePort(varStoragePurePort)
	} else {
		return err
	}

	varStoragePurePort := _StoragePurePort{}

	err = json.Unmarshal(bytes, &varStoragePurePort)
	if err == nil {
		o.StorageBasePhysicalPort = varStoragePurePort.StorageBasePhysicalPort
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Failover")
		delete(additionalProperties, "Portal")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "Controller")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBasePhysicalPort := reflect.ValueOf(o.StorageBasePhysicalPort)
		for i := 0; i < reflectStorageBasePhysicalPort.Type().NumField(); i++ {
			t := reflectStorageBasePhysicalPort.Type().Field(i)

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

type NullableStoragePurePort struct {
	value *StoragePurePort
	isSet bool
}

func (v NullableStoragePurePort) Get() *StoragePurePort {
	return v.value
}

func (v *NullableStoragePurePort) Set(val *StoragePurePort) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePurePort) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePurePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePurePort(val *StoragePurePort) *NullableStoragePurePort {
	return &NullableStoragePurePort{value: val, isSet: true}
}

func (v NullableStoragePurePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePurePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
