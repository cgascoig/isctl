# MemoryArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrayId** | Pointer to **int64** | ID of the memory array on a server. | [optional] [readonly] 
**CpuId** | Pointer to **int64** | ID of the CPU that access this memory array. | [optional] [readonly] 
**CurrentCapacity** | Pointer to **string** | Current capacity of all the memory units on a server. | [optional] [readonly] 
**ErrorCorrection** | Pointer to **string** |  | [optional] [readonly] 
**MaxCapacity** | Pointer to **string** | Maximum capacity of all the memory units on a server. | [optional] [readonly] 
**MaxDevices** | Pointer to **string** |  | [optional] [readonly] 
**OperPowerState** | Pointer to **string** |  | [optional] [readonly] 
**Presence** | Pointer to **string** |  | [optional] [readonly] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**PersistentMemoryUnits** | Pointer to [**[]MemoryPersistentMemoryUnitRelationship**](memory.PersistentMemoryUnit.Relationship.md) | An array of relationships to memoryPersistentMemoryUnit resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Units** | Pointer to [**[]MemoryUnitRelationship**](memory.Unit.Relationship.md) | An array of relationships to memoryUnit resources. | [optional] [readonly] 

## Methods

### NewMemoryArray

`func NewMemoryArray() *MemoryArray`

NewMemoryArray instantiates a new MemoryArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryArrayWithDefaults

`func NewMemoryArrayWithDefaults() *MemoryArray`

NewMemoryArrayWithDefaults instantiates a new MemoryArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrayId

`func (o *MemoryArray) GetArrayId() int64`

GetArrayId returns the ArrayId field if non-nil, zero value otherwise.

### GetArrayIdOk

`func (o *MemoryArray) GetArrayIdOk() (*int64, bool)`

GetArrayIdOk returns a tuple with the ArrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayId

`func (o *MemoryArray) SetArrayId(v int64)`

SetArrayId sets ArrayId field to given value.

### HasArrayId

`func (o *MemoryArray) HasArrayId() bool`

HasArrayId returns a boolean if a field has been set.

### GetCpuId

`func (o *MemoryArray) GetCpuId() int64`

GetCpuId returns the CpuId field if non-nil, zero value otherwise.

### GetCpuIdOk

`func (o *MemoryArray) GetCpuIdOk() (*int64, bool)`

GetCpuIdOk returns a tuple with the CpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuId

`func (o *MemoryArray) SetCpuId(v int64)`

SetCpuId sets CpuId field to given value.

### HasCpuId

`func (o *MemoryArray) HasCpuId() bool`

HasCpuId returns a boolean if a field has been set.

### GetCurrentCapacity

`func (o *MemoryArray) GetCurrentCapacity() string`

GetCurrentCapacity returns the CurrentCapacity field if non-nil, zero value otherwise.

### GetCurrentCapacityOk

`func (o *MemoryArray) GetCurrentCapacityOk() (*string, bool)`

GetCurrentCapacityOk returns a tuple with the CurrentCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCapacity

`func (o *MemoryArray) SetCurrentCapacity(v string)`

SetCurrentCapacity sets CurrentCapacity field to given value.

### HasCurrentCapacity

`func (o *MemoryArray) HasCurrentCapacity() bool`

HasCurrentCapacity returns a boolean if a field has been set.

### GetErrorCorrection

`func (o *MemoryArray) GetErrorCorrection() string`

GetErrorCorrection returns the ErrorCorrection field if non-nil, zero value otherwise.

### GetErrorCorrectionOk

`func (o *MemoryArray) GetErrorCorrectionOk() (*string, bool)`

GetErrorCorrectionOk returns a tuple with the ErrorCorrection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCorrection

`func (o *MemoryArray) SetErrorCorrection(v string)`

SetErrorCorrection sets ErrorCorrection field to given value.

### HasErrorCorrection

`func (o *MemoryArray) HasErrorCorrection() bool`

HasErrorCorrection returns a boolean if a field has been set.

### GetMaxCapacity

`func (o *MemoryArray) GetMaxCapacity() string`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *MemoryArray) GetMaxCapacityOk() (*string, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *MemoryArray) SetMaxCapacity(v string)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *MemoryArray) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### GetMaxDevices

`func (o *MemoryArray) GetMaxDevices() string`

GetMaxDevices returns the MaxDevices field if non-nil, zero value otherwise.

### GetMaxDevicesOk

`func (o *MemoryArray) GetMaxDevicesOk() (*string, bool)`

GetMaxDevicesOk returns a tuple with the MaxDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDevices

`func (o *MemoryArray) SetMaxDevices(v string)`

SetMaxDevices sets MaxDevices field to given value.

### HasMaxDevices

`func (o *MemoryArray) HasMaxDevices() bool`

HasMaxDevices returns a boolean if a field has been set.

### GetOperPowerState

`func (o *MemoryArray) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *MemoryArray) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *MemoryArray) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *MemoryArray) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetPresence

`func (o *MemoryArray) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *MemoryArray) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *MemoryArray) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *MemoryArray) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetComputeBoard

`func (o *MemoryArray) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *MemoryArray) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *MemoryArray) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *MemoryArray) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetPersistentMemoryUnits

`func (o *MemoryArray) GetPersistentMemoryUnits() []MemoryPersistentMemoryUnitRelationship`

GetPersistentMemoryUnits returns the PersistentMemoryUnits field if non-nil, zero value otherwise.

### GetPersistentMemoryUnitsOk

`func (o *MemoryArray) GetPersistentMemoryUnitsOk() (*[]MemoryPersistentMemoryUnitRelationship, bool)`

GetPersistentMemoryUnitsOk returns a tuple with the PersistentMemoryUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryUnits

`func (o *MemoryArray) SetPersistentMemoryUnits(v []MemoryPersistentMemoryUnitRelationship)`

SetPersistentMemoryUnits sets PersistentMemoryUnits field to given value.

### HasPersistentMemoryUnits

`func (o *MemoryArray) HasPersistentMemoryUnits() bool`

HasPersistentMemoryUnits returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *MemoryArray) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryArray) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryArray) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryArray) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetUnits

`func (o *MemoryArray) GetUnits() []MemoryUnitRelationship`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *MemoryArray) GetUnitsOk() (*[]MemoryUnitRelationship, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *MemoryArray) SetUnits(v []MemoryUnitRelationship)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *MemoryArray) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


