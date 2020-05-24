# StorageSnapshotAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to [**time.Time**](time.Time.md) | Exact date and time at which snapshot was created. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the snapshot which represents point-in-time copy of volume. | [optional] [readonly] 
**ProtectionGroupName** | Pointer to **string** | Name of the protection group to which the snapshot belongs. Value is empty, if the snapshot is created directly on volume. | [optional] [readonly] 
**Size** | Pointer to **int64** | Snapshot size represented in bytes. | [optional] [readonly] 
**Source** | Pointer to **string** | Source object on which the snapshot is created. It is the name of the originating volume. | [optional] [readonly] 
**ProtectionGroupSnapshot** | Pointer to [**StorageProtectionGroupSnapshotRelationship**](storage.ProtectionGroupSnapshot.Relationship.md) |  | [optional] 
**StorageArray** | Pointer to [**StorageGenericArrayRelationship**](storage.GenericArray.Relationship.md) |  | [optional] 
**Volume** | Pointer to [**StorageVolumeRelationship**](storage.Volume.Relationship.md) |  | [optional] 

## Methods

### NewStorageSnapshotAllOf

`func NewStorageSnapshotAllOf() *StorageSnapshotAllOf`

NewStorageSnapshotAllOf instantiates a new StorageSnapshotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSnapshotAllOfWithDefaults

`func NewStorageSnapshotAllOfWithDefaults() *StorageSnapshotAllOf`

NewStorageSnapshotAllOfWithDefaults instantiates a new StorageSnapshotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *StorageSnapshotAllOf) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StorageSnapshotAllOf) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StorageSnapshotAllOf) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *StorageSnapshotAllOf) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetName

`func (o *StorageSnapshotAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageSnapshotAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageSnapshotAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageSnapshotAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtectionGroupName

`func (o *StorageSnapshotAllOf) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *StorageSnapshotAllOf) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *StorageSnapshotAllOf) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *StorageSnapshotAllOf) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### GetSize

`func (o *StorageSnapshotAllOf) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageSnapshotAllOf) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageSnapshotAllOf) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageSnapshotAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSource

`func (o *StorageSnapshotAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StorageSnapshotAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StorageSnapshotAllOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StorageSnapshotAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetProtectionGroupSnapshot

`func (o *StorageSnapshotAllOf) GetProtectionGroupSnapshot() StorageProtectionGroupSnapshotRelationship`

GetProtectionGroupSnapshot returns the ProtectionGroupSnapshot field if non-nil, zero value otherwise.

### GetProtectionGroupSnapshotOk

`func (o *StorageSnapshotAllOf) GetProtectionGroupSnapshotOk() (*StorageProtectionGroupSnapshotRelationship, bool)`

GetProtectionGroupSnapshotOk returns a tuple with the ProtectionGroupSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupSnapshot

`func (o *StorageSnapshotAllOf) SetProtectionGroupSnapshot(v StorageProtectionGroupSnapshotRelationship)`

SetProtectionGroupSnapshot sets ProtectionGroupSnapshot field to given value.

### HasProtectionGroupSnapshot

`func (o *StorageSnapshotAllOf) HasProtectionGroupSnapshot() bool`

HasProtectionGroupSnapshot returns a boolean if a field has been set.

### GetStorageArray

`func (o *StorageSnapshotAllOf) GetStorageArray() StorageGenericArrayRelationship`

GetStorageArray returns the StorageArray field if non-nil, zero value otherwise.

### GetStorageArrayOk

`func (o *StorageSnapshotAllOf) GetStorageArrayOk() (*StorageGenericArrayRelationship, bool)`

GetStorageArrayOk returns a tuple with the StorageArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageArray

`func (o *StorageSnapshotAllOf) SetStorageArray(v StorageGenericArrayRelationship)`

SetStorageArray sets StorageArray field to given value.

### HasStorageArray

`func (o *StorageSnapshotAllOf) HasStorageArray() bool`

HasStorageArray returns a boolean if a field has been set.

### GetVolume

`func (o *StorageSnapshotAllOf) GetVolume() StorageVolumeRelationship`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StorageSnapshotAllOf) GetVolumeOk() (*StorageVolumeRelationship, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StorageSnapshotAllOf) SetVolume(v StorageVolumeRelationship)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StorageSnapshotAllOf) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


