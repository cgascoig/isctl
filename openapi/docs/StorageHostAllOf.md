# StorageHostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Short description about the host. | [optional] [readonly] 
**Initiators** | Pointer to [**[]StorageInitiator**](storage.Initiator.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the host in storage array. | [optional] [readonly] 
**OsType** | Pointer to **string** | Operating system running on the host. | [optional] [readonly] 
**StorageArray** | Pointer to [**StorageGenericArrayRelationship**](storage.GenericArray.Relationship.md) |  | [optional] 

## Methods

### NewStorageHostAllOf

`func NewStorageHostAllOf() *StorageHostAllOf`

NewStorageHostAllOf instantiates a new StorageHostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHostAllOfWithDefaults

`func NewStorageHostAllOfWithDefaults() *StorageHostAllOf`

NewStorageHostAllOfWithDefaults instantiates a new StorageHostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StorageHostAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageHostAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageHostAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageHostAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInitiators

`func (o *StorageHostAllOf) GetInitiators() []StorageInitiator`

GetInitiators returns the Initiators field if non-nil, zero value otherwise.

### GetInitiatorsOk

`func (o *StorageHostAllOf) GetInitiatorsOk() (*[]StorageInitiator, bool)`

GetInitiatorsOk returns a tuple with the Initiators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiators

`func (o *StorageHostAllOf) SetInitiators(v []StorageInitiator)`

SetInitiators sets Initiators field to given value.

### HasInitiators

`func (o *StorageHostAllOf) HasInitiators() bool`

HasInitiators returns a boolean if a field has been set.

### GetName

`func (o *StorageHostAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageHostAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageHostAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageHostAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsType

`func (o *StorageHostAllOf) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *StorageHostAllOf) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *StorageHostAllOf) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *StorageHostAllOf) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetStorageArray

`func (o *StorageHostAllOf) GetStorageArray() StorageGenericArrayRelationship`

GetStorageArray returns the StorageArray field if non-nil, zero value otherwise.

### GetStorageArrayOk

`func (o *StorageHostAllOf) GetStorageArrayOk() (*StorageGenericArrayRelationship, bool)`

GetStorageArrayOk returns a tuple with the StorageArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageArray

`func (o *StorageHostAllOf) SetStorageArray(v StorageGenericArrayRelationship)`

SetStorageArray sets StorageArray field to given value.

### HasStorageArray

`func (o *StorageHostAllOf) HasStorageArray() bool`

HasStorageArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


