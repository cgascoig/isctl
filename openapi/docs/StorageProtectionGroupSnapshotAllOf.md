# StorageProtectionGroupSnapshotAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to [**time.Time**](time.Time.md) | Protection group snapshot creation time. | [optional] [readonly] 
**Name** | Pointer to **string** | Protection group snapshot name which represents point-in-time copy of all members in protection group. | [optional] [readonly] 
**Size** | Pointer to **int64** | Snapshot size represented in bytes. It is a cumulative size of all snapshots in a set. | [optional] [readonly] 
**Source** | Pointer to **string** | Source protection group name on which the snapshot is created. | [optional] [readonly] 
**ProtectionGroup** | Pointer to [**StorageProtectionGroupRelationship**](storage.ProtectionGroup.Relationship.md) |  | [optional] 
**StorageArray** | Pointer to [**StorageGenericArrayRelationship**](storage.GenericArray.Relationship.md) |  | [optional] 

## Methods

### NewStorageProtectionGroupSnapshotAllOf

`func NewStorageProtectionGroupSnapshotAllOf() *StorageProtectionGroupSnapshotAllOf`

NewStorageProtectionGroupSnapshotAllOf instantiates a new StorageProtectionGroupSnapshotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageProtectionGroupSnapshotAllOfWithDefaults

`func NewStorageProtectionGroupSnapshotAllOfWithDefaults() *StorageProtectionGroupSnapshotAllOf`

NewStorageProtectionGroupSnapshotAllOfWithDefaults instantiates a new StorageProtectionGroupSnapshotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *StorageProtectionGroupSnapshotAllOf) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StorageProtectionGroupSnapshotAllOf) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StorageProtectionGroupSnapshotAllOf) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *StorageProtectionGroupSnapshotAllOf) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetName

`func (o *StorageProtectionGroupSnapshotAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageProtectionGroupSnapshotAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageProtectionGroupSnapshotAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageProtectionGroupSnapshotAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *StorageProtectionGroupSnapshotAllOf) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageProtectionGroupSnapshotAllOf) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageProtectionGroupSnapshotAllOf) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageProtectionGroupSnapshotAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSource

`func (o *StorageProtectionGroupSnapshotAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StorageProtectionGroupSnapshotAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StorageProtectionGroupSnapshotAllOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StorageProtectionGroupSnapshotAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetProtectionGroup

`func (o *StorageProtectionGroupSnapshotAllOf) GetProtectionGroup() StorageProtectionGroupRelationship`

GetProtectionGroup returns the ProtectionGroup field if non-nil, zero value otherwise.

### GetProtectionGroupOk

`func (o *StorageProtectionGroupSnapshotAllOf) GetProtectionGroupOk() (*StorageProtectionGroupRelationship, bool)`

GetProtectionGroupOk returns a tuple with the ProtectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroup

`func (o *StorageProtectionGroupSnapshotAllOf) SetProtectionGroup(v StorageProtectionGroupRelationship)`

SetProtectionGroup sets ProtectionGroup field to given value.

### HasProtectionGroup

`func (o *StorageProtectionGroupSnapshotAllOf) HasProtectionGroup() bool`

HasProtectionGroup returns a boolean if a field has been set.

### GetStorageArray

`func (o *StorageProtectionGroupSnapshotAllOf) GetStorageArray() StorageGenericArrayRelationship`

GetStorageArray returns the StorageArray field if non-nil, zero value otherwise.

### GetStorageArrayOk

`func (o *StorageProtectionGroupSnapshotAllOf) GetStorageArrayOk() (*StorageGenericArrayRelationship, bool)`

GetStorageArrayOk returns a tuple with the StorageArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageArray

`func (o *StorageProtectionGroupSnapshotAllOf) SetStorageArray(v StorageGenericArrayRelationship)`

SetStorageArray sets StorageArray field to given value.

### HasStorageArray

`func (o *StorageProtectionGroupSnapshotAllOf) HasStorageArray() bool`

HasStorageArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


