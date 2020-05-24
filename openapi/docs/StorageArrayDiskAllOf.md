# StorageArrayDiskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Disk name available in storage array. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Storage disk part number. | [optional] [readonly] 
**Protocol** | Pointer to **string** | Storage protocol used in disk for communication. Possible values are SAS, SATA and NVMe. | [optional] [readonly] [default to "Unknown"]
**Speed** | Pointer to **int64** | Disk speed for read or write operation measured in rpm. | [optional] [readonly] 
**Status** | Pointer to **string** | Storage disk health status. | [optional] [readonly] [default to "Unknown"]
**StorageUtilization** | Pointer to [**StorageCapacity**](storage.Capacity.md) |  | [optional] 
**Type** | Pointer to **string** | Storage disk type - it can be SSD, HDD, NVRAM. | [optional] [readonly] [default to "Unknown"]
**Version** | Pointer to **string** | Storage disk version number. | [optional] [readonly] 
**StorageArray** | Pointer to [**StorageGenericArrayRelationship**](storage.GenericArray.Relationship.md) |  | [optional] 

## Methods

### NewStorageArrayDiskAllOf

`func NewStorageArrayDiskAllOf() *StorageArrayDiskAllOf`

NewStorageArrayDiskAllOf instantiates a new StorageArrayDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageArrayDiskAllOfWithDefaults

`func NewStorageArrayDiskAllOfWithDefaults() *StorageArrayDiskAllOf`

NewStorageArrayDiskAllOfWithDefaults instantiates a new StorageArrayDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageArrayDiskAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageArrayDiskAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageArrayDiskAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageArrayDiskAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartNumber

`func (o *StorageArrayDiskAllOf) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *StorageArrayDiskAllOf) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *StorageArrayDiskAllOf) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *StorageArrayDiskAllOf) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetProtocol

`func (o *StorageArrayDiskAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StorageArrayDiskAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StorageArrayDiskAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *StorageArrayDiskAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSpeed

`func (o *StorageArrayDiskAllOf) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *StorageArrayDiskAllOf) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *StorageArrayDiskAllOf) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *StorageArrayDiskAllOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *StorageArrayDiskAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageArrayDiskAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageArrayDiskAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageArrayDiskAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StorageArrayDiskAllOf) GetStorageUtilization() StorageCapacity`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StorageArrayDiskAllOf) GetStorageUtilizationOk() (*StorageCapacity, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StorageArrayDiskAllOf) SetStorageUtilization(v StorageCapacity)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StorageArrayDiskAllOf) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### GetType

`func (o *StorageArrayDiskAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageArrayDiskAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageArrayDiskAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageArrayDiskAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *StorageArrayDiskAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StorageArrayDiskAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StorageArrayDiskAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StorageArrayDiskAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStorageArray

`func (o *StorageArrayDiskAllOf) GetStorageArray() StorageGenericArrayRelationship`

GetStorageArray returns the StorageArray field if non-nil, zero value otherwise.

### GetStorageArrayOk

`func (o *StorageArrayDiskAllOf) GetStorageArrayOk() (*StorageGenericArrayRelationship, bool)`

GetStorageArrayOk returns a tuple with the StorageArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageArray

`func (o *StorageArrayDiskAllOf) SetStorageArray(v StorageGenericArrayRelationship)`

SetStorageArray sets StorageArray field to given value.

### HasStorageArray

`func (o *StorageArrayDiskAllOf) HasStorageArray() bool`

HasStorageArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


