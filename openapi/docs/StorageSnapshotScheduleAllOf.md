# StorageSnapshotScheduleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | Pointer to **string** | Snapshot frequency. It is an interval at which snapshot is set to trigger on source array. Examples:     PT2H Snapshot is generated every 2 hours.     P4D, Snapshot is scheduled for every 4 days.     PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the snapshot schedule. | [optional] 
**RetentionTime** | Pointer to **string** | Duration to keep the snapshots on the source array. Once this period expires, system deletes the snapshot automatically from the source array. Examples: P200D,  200 days. PT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds. | [optional] [readonly] 
**SnapshotTime** | Pointer to **string** | Preferred time of the day to capture the snapshot. It is applicable only if the frequency is set for a day or more. Format: hh:mm:ss Example: 08:30:00, Snapshot is set for 08:30 AM. | [optional] [readonly] 
**ProtectionGroup** | Pointer to [**StorageProtectionGroupRelationship**](storage.ProtectionGroup.Relationship.md) |  | [optional] 
**StorageArray** | Pointer to [**StorageGenericArrayRelationship**](storage.GenericArray.Relationship.md) |  | [optional] 

## Methods

### NewStorageSnapshotScheduleAllOf

`func NewStorageSnapshotScheduleAllOf() *StorageSnapshotScheduleAllOf`

NewStorageSnapshotScheduleAllOf instantiates a new StorageSnapshotScheduleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSnapshotScheduleAllOfWithDefaults

`func NewStorageSnapshotScheduleAllOfWithDefaults() *StorageSnapshotScheduleAllOf`

NewStorageSnapshotScheduleAllOfWithDefaults instantiates a new StorageSnapshotScheduleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *StorageSnapshotScheduleAllOf) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *StorageSnapshotScheduleAllOf) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *StorageSnapshotScheduleAllOf) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *StorageSnapshotScheduleAllOf) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetName

`func (o *StorageSnapshotScheduleAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageSnapshotScheduleAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageSnapshotScheduleAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageSnapshotScheduleAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetentionTime

`func (o *StorageSnapshotScheduleAllOf) GetRetentionTime() string`

GetRetentionTime returns the RetentionTime field if non-nil, zero value otherwise.

### GetRetentionTimeOk

`func (o *StorageSnapshotScheduleAllOf) GetRetentionTimeOk() (*string, bool)`

GetRetentionTimeOk returns a tuple with the RetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionTime

`func (o *StorageSnapshotScheduleAllOf) SetRetentionTime(v string)`

SetRetentionTime sets RetentionTime field to given value.

### HasRetentionTime

`func (o *StorageSnapshotScheduleAllOf) HasRetentionTime() bool`

HasRetentionTime returns a boolean if a field has been set.

### GetSnapshotTime

`func (o *StorageSnapshotScheduleAllOf) GetSnapshotTime() string`

GetSnapshotTime returns the SnapshotTime field if non-nil, zero value otherwise.

### GetSnapshotTimeOk

`func (o *StorageSnapshotScheduleAllOf) GetSnapshotTimeOk() (*string, bool)`

GetSnapshotTimeOk returns a tuple with the SnapshotTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTime

`func (o *StorageSnapshotScheduleAllOf) SetSnapshotTime(v string)`

SetSnapshotTime sets SnapshotTime field to given value.

### HasSnapshotTime

`func (o *StorageSnapshotScheduleAllOf) HasSnapshotTime() bool`

HasSnapshotTime returns a boolean if a field has been set.

### GetProtectionGroup

`func (o *StorageSnapshotScheduleAllOf) GetProtectionGroup() StorageProtectionGroupRelationship`

GetProtectionGroup returns the ProtectionGroup field if non-nil, zero value otherwise.

### GetProtectionGroupOk

`func (o *StorageSnapshotScheduleAllOf) GetProtectionGroupOk() (*StorageProtectionGroupRelationship, bool)`

GetProtectionGroupOk returns a tuple with the ProtectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroup

`func (o *StorageSnapshotScheduleAllOf) SetProtectionGroup(v StorageProtectionGroupRelationship)`

SetProtectionGroup sets ProtectionGroup field to given value.

### HasProtectionGroup

`func (o *StorageSnapshotScheduleAllOf) HasProtectionGroup() bool`

HasProtectionGroup returns a boolean if a field has been set.

### GetStorageArray

`func (o *StorageSnapshotScheduleAllOf) GetStorageArray() StorageGenericArrayRelationship`

GetStorageArray returns the StorageArray field if non-nil, zero value otherwise.

### GetStorageArrayOk

`func (o *StorageSnapshotScheduleAllOf) GetStorageArrayOk() (*StorageGenericArrayRelationship, bool)`

GetStorageArrayOk returns a tuple with the StorageArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageArray

`func (o *StorageSnapshotScheduleAllOf) SetStorageArray(v StorageGenericArrayRelationship)`

SetStorageArray sets StorageArray field to given value.

### HasStorageArray

`func (o *StorageSnapshotScheduleAllOf) HasStorageArray() bool`

HasStorageArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


