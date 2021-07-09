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

// StoragePureArrayUtilization Storage space consumption of Pure FlashArray.
type StoragePureArrayUtilization struct {
	StorageBaseCapacity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Ratio of mapped sectors within a volume versus the amount of physical space the data occupies after data compression and deduplication. The data reduction ratio does not include thin provisioning savings. For example, a data reduction ratio of 5.0 means that for every 5 MB the host writes to the array, 1 MB is stored on the array's flash modules.
	DataReduction *float32 `json:"DataReduction,omitempty"`
	// Percentage of data that is fully protected. The percentage value will drop below 100% if the data is not fully protected.
	Parity *float32 `json:"Parity,omitempty"`
	// Total provisioned storage capacity in Pure FlashArray, represented in bytes.
	Provisioned *int64 `json:"Provisioned,omitempty"`
	// Physical space occupied by deduplicated data, represented in bytes. The space is shared with other volumes and snapshots as a result of data deduplication.
	Shared *int64 `json:"Shared,omitempty"`
	// Physical space occupied by the snapshots, represented in bytes.
	Snapshot *int64 `json:"Snapshot,omitempty"`
	// Physical space occupied by internal array metadata, represented in bytes.
	System *int64 `json:"System,omitempty"`
	// Percentage of volume sectors that do not contain host-written data because the hosts have not written data to them or the sectors have been explicitly trimmed.
	ThinProvisioned *float32 `json:"ThinProvisioned,omitempty"`
	// Ratio of provisioned sectors within a volume versus the amount of physical space the data occupies after reduction via data compression and deduplication and with thin provisioning savings. Total reduction is data reduction with thin provisioning savings. For example, a total reduction ratio of 10.0 means that for every 10 MB of provisioned space, 1 MB is stored on the array's flash modules.
	TotalReduction *float32 `json:"TotalReduction,omitempty"`
	// Physical space occupied by volume data, excluding shared, array metadata and snapshots. Size is represented in bytes.
	Volume               *int64 `json:"Volume,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureArrayUtilization StoragePureArrayUtilization

// NewStoragePureArrayUtilization instantiates a new StoragePureArrayUtilization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureArrayUtilization(classId string, objectType string) *StoragePureArrayUtilization {
	this := StoragePureArrayUtilization{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureArrayUtilizationWithDefaults instantiates a new StoragePureArrayUtilization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureArrayUtilizationWithDefaults() *StoragePureArrayUtilization {
	this := StoragePureArrayUtilization{}
	var classId string = "storage.PureArrayUtilization"
	this.ClassId = classId
	var objectType string = "storage.PureArrayUtilization"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureArrayUtilization) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureArrayUtilization) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureArrayUtilization) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureArrayUtilization) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureArrayUtilization) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureArrayUtilization) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDataReduction returns the DataReduction field value if set, zero value otherwise.
func (o *StoragePureArrayUtilization) GetDataReduction() float32 {
	if o == nil || o.DataReduction == nil {
		var ret float32
		return ret
	}
	return *o.DataReduction
}

// GetDataReductionOk returns a tuple with the DataReduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureArrayUtilization) GetDataReductionOk() (*float32, bool) {
	if o == nil || o.DataReduction == nil {
		return nil, false
	}
	return o.DataReduction, true
}

// HasDataReduction returns a boolean if a field has been set.
func (o *StoragePureArrayUtilization) HasDataReduction() bool {
	if o != nil && o.DataReduction != nil {
		return true
	}

	return false
}

// SetDataReduction gets a reference to the given float32 and assigns it to the DataReduction field.
func (o *StoragePureArrayUtilization) SetDataReduction(v float32) {
	o.DataReduction = &v
}

// GetParity returns the Parity field value if set, zero value otherwise.
func (o *StoragePureArrayUtilization) GetParity() float32 {
	if o == nil || o.Parity == nil {
		var ret float32
		return ret
	}
	return *o.Parity
}

// GetParityOk returns a tuple with the Parity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureArrayUtilization) GetParityOk() (*float32, bool) {
	if o == nil || o.Parity == nil {
		return nil, false
	}
	return o.Parity, true
}

// HasParity returns a boolean if a field has been set.
func (o *StoragePureArrayUtilization) HasParity() bool {
	if o != nil && o.Parity != nil {
		return true
	}

	return false
}

// SetParity gets a reference to the given float32 and assigns it to the Parity field.
func (o *StoragePureArrayUtilization) SetParity(v float32) {
	o.Parity = &v
}

// GetProvisioned returns the Provisioned field value if set, zero value otherwise.
func (o *StoragePureArrayUtilization) GetProvisioned() int64 {
	if o == nil || o.Provisioned == nil {
		var ret int64
		return ret
	}
	return *o.Provisioned
}

// GetProvisionedOk returns a tuple with the Provisioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureArrayUtilization) GetProvisionedOk() (*int64, bool) {
	if o == nil || o.Provisioned == nil {
		return nil, false
	}
	return o.Provisioned, true
}

// HasProvisioned returns a boolean if a field has been set.
func (o *StoragePureArrayUtilization) HasProvisioned() bool {
	if o != nil && o.Provisioned != nil {
		return true
	}

	return false
}

// SetProvisioned gets a reference to the given int64 and assigns it to the Provisioned field.
func (o *StoragePureArrayUtilization) SetProvisioned(v int64) {
	o.Provisioned = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *StoragePureArrayUtilization) GetShared() int64 {
	if o == nil || o.Shared == nil {
		var ret int64
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureArrayUtilization) GetSharedOk() (*int64, bool) {
	if o == nil || o.Shared == nil {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *StoragePureArrayUtilization) HasShared() bool {
	if o != nil && o.Shared != nil {
		return true
	}

	return false
}

// SetShared gets a reference to the given int64 and assigns it to the Shared field.
func (o *StoragePureArrayUtilization) SetShared(v int64) {
	o.Shared = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *StoragePureArrayUtilization) GetSnapshot() int64 {
	if o == nil || o.Snapshot == nil {
		var ret int64
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureArrayUtilization) GetSnapshotOk() (*int64, bool) {
	if o == nil || o.Snapshot == nil {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *StoragePureArrayUtilization) HasSnapshot() bool {
	if o != nil && o.Snapshot != nil {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given int64 and assigns it to the Snapshot field.
func (o *StoragePureArrayUtilization) SetSnapshot(v int64) {
	o.Snapshot = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *StoragePureArrayUtilization) GetSystem() int64 {
	if o == nil || o.System == nil {
		var ret int64
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureArrayUtilization) GetSystemOk() (*int64, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *StoragePureArrayUtilization) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given int64 and assigns it to the System field.
func (o *StoragePureArrayUtilization) SetSystem(v int64) {
	o.System = &v
}

// GetThinProvisioned returns the ThinProvisioned field value if set, zero value otherwise.
func (o *StoragePureArrayUtilization) GetThinProvisioned() float32 {
	if o == nil || o.ThinProvisioned == nil {
		var ret float32
		return ret
	}
	return *o.ThinProvisioned
}

// GetThinProvisionedOk returns a tuple with the ThinProvisioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureArrayUtilization) GetThinProvisionedOk() (*float32, bool) {
	if o == nil || o.ThinProvisioned == nil {
		return nil, false
	}
	return o.ThinProvisioned, true
}

// HasThinProvisioned returns a boolean if a field has been set.
func (o *StoragePureArrayUtilization) HasThinProvisioned() bool {
	if o != nil && o.ThinProvisioned != nil {
		return true
	}

	return false
}

// SetThinProvisioned gets a reference to the given float32 and assigns it to the ThinProvisioned field.
func (o *StoragePureArrayUtilization) SetThinProvisioned(v float32) {
	o.ThinProvisioned = &v
}

// GetTotalReduction returns the TotalReduction field value if set, zero value otherwise.
func (o *StoragePureArrayUtilization) GetTotalReduction() float32 {
	if o == nil || o.TotalReduction == nil {
		var ret float32
		return ret
	}
	return *o.TotalReduction
}

// GetTotalReductionOk returns a tuple with the TotalReduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureArrayUtilization) GetTotalReductionOk() (*float32, bool) {
	if o == nil || o.TotalReduction == nil {
		return nil, false
	}
	return o.TotalReduction, true
}

// HasTotalReduction returns a boolean if a field has been set.
func (o *StoragePureArrayUtilization) HasTotalReduction() bool {
	if o != nil && o.TotalReduction != nil {
		return true
	}

	return false
}

// SetTotalReduction gets a reference to the given float32 and assigns it to the TotalReduction field.
func (o *StoragePureArrayUtilization) SetTotalReduction(v float32) {
	o.TotalReduction = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *StoragePureArrayUtilization) GetVolume() int64 {
	if o == nil || o.Volume == nil {
		var ret int64
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureArrayUtilization) GetVolumeOk() (*int64, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *StoragePureArrayUtilization) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given int64 and assigns it to the Volume field.
func (o *StoragePureArrayUtilization) SetVolume(v int64) {
	o.Volume = &v
}

func (o StoragePureArrayUtilization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseCapacity, errStorageBaseCapacity := json.Marshal(o.StorageBaseCapacity)
	if errStorageBaseCapacity != nil {
		return []byte{}, errStorageBaseCapacity
	}
	errStorageBaseCapacity = json.Unmarshal([]byte(serializedStorageBaseCapacity), &toSerialize)
	if errStorageBaseCapacity != nil {
		return []byte{}, errStorageBaseCapacity
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DataReduction != nil {
		toSerialize["DataReduction"] = o.DataReduction
	}
	if o.Parity != nil {
		toSerialize["Parity"] = o.Parity
	}
	if o.Provisioned != nil {
		toSerialize["Provisioned"] = o.Provisioned
	}
	if o.Shared != nil {
		toSerialize["Shared"] = o.Shared
	}
	if o.Snapshot != nil {
		toSerialize["Snapshot"] = o.Snapshot
	}
	if o.System != nil {
		toSerialize["System"] = o.System
	}
	if o.ThinProvisioned != nil {
		toSerialize["ThinProvisioned"] = o.ThinProvisioned
	}
	if o.TotalReduction != nil {
		toSerialize["TotalReduction"] = o.TotalReduction
	}
	if o.Volume != nil {
		toSerialize["Volume"] = o.Volume
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePureArrayUtilization) UnmarshalJSON(bytes []byte) (err error) {
	type StoragePureArrayUtilizationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Ratio of mapped sectors within a volume versus the amount of physical space the data occupies after data compression and deduplication. The data reduction ratio does not include thin provisioning savings. For example, a data reduction ratio of 5.0 means that for every 5 MB the host writes to the array, 1 MB is stored on the array's flash modules.
		DataReduction *float32 `json:"DataReduction,omitempty"`
		// Percentage of data that is fully protected. The percentage value will drop below 100% if the data is not fully protected.
		Parity *float32 `json:"Parity,omitempty"`
		// Total provisioned storage capacity in Pure FlashArray, represented in bytes.
		Provisioned *int64 `json:"Provisioned,omitempty"`
		// Physical space occupied by deduplicated data, represented in bytes. The space is shared with other volumes and snapshots as a result of data deduplication.
		Shared *int64 `json:"Shared,omitempty"`
		// Physical space occupied by the snapshots, represented in bytes.
		Snapshot *int64 `json:"Snapshot,omitempty"`
		// Physical space occupied by internal array metadata, represented in bytes.
		System *int64 `json:"System,omitempty"`
		// Percentage of volume sectors that do not contain host-written data because the hosts have not written data to them or the sectors have been explicitly trimmed.
		ThinProvisioned *float32 `json:"ThinProvisioned,omitempty"`
		// Ratio of provisioned sectors within a volume versus the amount of physical space the data occupies after reduction via data compression and deduplication and with thin provisioning savings. Total reduction is data reduction with thin provisioning savings. For example, a total reduction ratio of 10.0 means that for every 10 MB of provisioned space, 1 MB is stored on the array's flash modules.
		TotalReduction *float32 `json:"TotalReduction,omitempty"`
		// Physical space occupied by volume data, excluding shared, array metadata and snapshots. Size is represented in bytes.
		Volume *int64 `json:"Volume,omitempty"`
	}

	varStoragePureArrayUtilizationWithoutEmbeddedStruct := StoragePureArrayUtilizationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStoragePureArrayUtilizationWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureArrayUtilization := _StoragePureArrayUtilization{}
		varStoragePureArrayUtilization.ClassId = varStoragePureArrayUtilizationWithoutEmbeddedStruct.ClassId
		varStoragePureArrayUtilization.ObjectType = varStoragePureArrayUtilizationWithoutEmbeddedStruct.ObjectType
		varStoragePureArrayUtilization.DataReduction = varStoragePureArrayUtilizationWithoutEmbeddedStruct.DataReduction
		varStoragePureArrayUtilization.Parity = varStoragePureArrayUtilizationWithoutEmbeddedStruct.Parity
		varStoragePureArrayUtilization.Provisioned = varStoragePureArrayUtilizationWithoutEmbeddedStruct.Provisioned
		varStoragePureArrayUtilization.Shared = varStoragePureArrayUtilizationWithoutEmbeddedStruct.Shared
		varStoragePureArrayUtilization.Snapshot = varStoragePureArrayUtilizationWithoutEmbeddedStruct.Snapshot
		varStoragePureArrayUtilization.System = varStoragePureArrayUtilizationWithoutEmbeddedStruct.System
		varStoragePureArrayUtilization.ThinProvisioned = varStoragePureArrayUtilizationWithoutEmbeddedStruct.ThinProvisioned
		varStoragePureArrayUtilization.TotalReduction = varStoragePureArrayUtilizationWithoutEmbeddedStruct.TotalReduction
		varStoragePureArrayUtilization.Volume = varStoragePureArrayUtilizationWithoutEmbeddedStruct.Volume
		*o = StoragePureArrayUtilization(varStoragePureArrayUtilization)
	} else {
		return err
	}

	varStoragePureArrayUtilization := _StoragePureArrayUtilization{}

	err = json.Unmarshal(bytes, &varStoragePureArrayUtilization)
	if err == nil {
		o.StorageBaseCapacity = varStoragePureArrayUtilization.StorageBaseCapacity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DataReduction")
		delete(additionalProperties, "Parity")
		delete(additionalProperties, "Provisioned")
		delete(additionalProperties, "Shared")
		delete(additionalProperties, "Snapshot")
		delete(additionalProperties, "System")
		delete(additionalProperties, "ThinProvisioned")
		delete(additionalProperties, "TotalReduction")
		delete(additionalProperties, "Volume")

		// remove fields from embedded structs
		reflectStorageBaseCapacity := reflect.ValueOf(o.StorageBaseCapacity)
		for i := 0; i < reflectStorageBaseCapacity.Type().NumField(); i++ {
			t := reflectStorageBaseCapacity.Type().Field(i)

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

type NullableStoragePureArrayUtilization struct {
	value *StoragePureArrayUtilization
	isSet bool
}

func (v NullableStoragePureArrayUtilization) Get() *StoragePureArrayUtilization {
	return v.value
}

func (v *NullableStoragePureArrayUtilization) Set(val *StoragePureArrayUtilization) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureArrayUtilization) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureArrayUtilization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureArrayUtilization(val *StoragePureArrayUtilization) *NullableStoragePureArrayUtilization {
	return &NullableStoragePureArrayUtilization{value: val, isSet: true}
}

func (v NullableStoragePureArrayUtilization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureArrayUtilization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
