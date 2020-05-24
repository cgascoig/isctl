# StorageReplicationScheduleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | Pointer to **string** | Replication frequency. It is an interval at which replication is set to trigger. Examples:     PT2H, Snapshot is generated every 2 hours.     P30D, Snapshot is scheduled for every 30 days.     PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds. | [optional] [readonly] 
**Name** | Pointer to **string** | Replication schedule name. | [optional] [readonly] 
**ReplicationTime** | Pointer to **string** | Preferred time of day at which to replicate the snapshots on target array. It is applicable only if the replication frequency is set for a day or more. Format: hh:mm:ss Example: 15:00:00, Replication is set for 3:00 PM. | [optional] [readonly] 
**RetentionTime** | Pointer to **string** | Duration to keep the replicated snapshots on the targets. Replicated snapshots are deleted from target array once the retention period has elapsed. Examples: P30D, Snapshots are available for 30 days. PT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds. | [optional] [readonly] 
**ProtectionGroup** | Pointer to [**StorageProtectionGroupRelationship**](storage.ProtectionGroup.Relationship.md) |  | [optional] 
**StorageArray** | Pointer to [**StorageGenericArrayRelationship**](storage.GenericArray.Relationship.md) |  | [optional] 

## Methods

### NewStorageReplicationScheduleAllOf

`func NewStorageReplicationScheduleAllOf() *StorageReplicationScheduleAllOf`

NewStorageReplicationScheduleAllOf instantiates a new StorageReplicationScheduleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageReplicationScheduleAllOfWithDefaults

`func NewStorageReplicationScheduleAllOfWithDefaults() *StorageReplicationScheduleAllOf`

NewStorageReplicationScheduleAllOfWithDefaults instantiates a new StorageReplicationScheduleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *StorageReplicationScheduleAllOf) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *StorageReplicationScheduleAllOf) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *StorageReplicationScheduleAllOf) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *StorageReplicationScheduleAllOf) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetName

`func (o *StorageReplicationScheduleAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageReplicationScheduleAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageReplicationScheduleAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageReplicationScheduleAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReplicationTime

`func (o *StorageReplicationScheduleAllOf) GetReplicationTime() string`

GetReplicationTime returns the ReplicationTime field if non-nil, zero value otherwise.

### GetReplicationTimeOk

`func (o *StorageReplicationScheduleAllOf) GetReplicationTimeOk() (*string, bool)`

GetReplicationTimeOk returns a tuple with the ReplicationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationTime

`func (o *StorageReplicationScheduleAllOf) SetReplicationTime(v string)`

SetReplicationTime sets ReplicationTime field to given value.

### HasReplicationTime

`func (o *StorageReplicationScheduleAllOf) HasReplicationTime() bool`

HasReplicationTime returns a boolean if a field has been set.

### GetRetentionTime

`func (o *StorageReplicationScheduleAllOf) GetRetentionTime() string`

GetRetentionTime returns the RetentionTime field if non-nil, zero value otherwise.

### GetRetentionTimeOk

`func (o *StorageReplicationScheduleAllOf) GetRetentionTimeOk() (*string, bool)`

GetRetentionTimeOk returns a tuple with the RetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionTime

`func (o *StorageReplicationScheduleAllOf) SetRetentionTime(v string)`

SetRetentionTime sets RetentionTime field to given value.

### HasRetentionTime

`func (o *StorageReplicationScheduleAllOf) HasRetentionTime() bool`

HasRetentionTime returns a boolean if a field has been set.

### GetProtectionGroup

`func (o *StorageReplicationScheduleAllOf) GetProtectionGroup() StorageProtectionGroupRelationship`

GetProtectionGroup returns the ProtectionGroup field if non-nil, zero value otherwise.

### GetProtectionGroupOk

`func (o *StorageReplicationScheduleAllOf) GetProtectionGroupOk() (*StorageProtectionGroupRelationship, bool)`

GetProtectionGroupOk returns a tuple with the ProtectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroup

`func (o *StorageReplicationScheduleAllOf) SetProtectionGroup(v StorageProtectionGroupRelationship)`

SetProtectionGroup sets ProtectionGroup field to given value.

### HasProtectionGroup

`func (o *StorageReplicationScheduleAllOf) HasProtectionGroup() bool`

HasProtectionGroup returns a boolean if a field has been set.

### GetStorageArray

`func (o *StorageReplicationScheduleAllOf) GetStorageArray() StorageGenericArrayRelationship`

GetStorageArray returns the StorageArray field if non-nil, zero value otherwise.

### GetStorageArrayOk

`func (o *StorageReplicationScheduleAllOf) GetStorageArrayOk() (*StorageGenericArrayRelationship, bool)`

GetStorageArrayOk returns a tuple with the StorageArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageArray

`func (o *StorageReplicationScheduleAllOf) SetStorageArray(v StorageGenericArrayRelationship)`

SetStorageArray sets StorageArray field to given value.

### HasStorageArray

`func (o *StorageReplicationScheduleAllOf) HasStorageArray() bool`

HasStorageArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


