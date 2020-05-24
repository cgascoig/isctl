# StorageReplicationBlackout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | The end time of day for replication blackout window. Example: 17:00:01 which is 17 hours, 0 minutes, 1 seconds. | [optional] [readonly] 
**Start** | Pointer to **string** | The start time of day when replication blackout is active. When replication blackout is active, the storage array temporarily disables replication. Example: 15:04:03.123 which is 15 hours, 4 minutes, 3 seconds and 123 milliseconds. The fractional seconds are written using the standard decimal notation which can be used for setting milliseconds and microseconds. | [optional] [readonly] 

## Methods

### NewStorageReplicationBlackout

`func NewStorageReplicationBlackout() *StorageReplicationBlackout`

NewStorageReplicationBlackout instantiates a new StorageReplicationBlackout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageReplicationBlackoutWithDefaults

`func NewStorageReplicationBlackoutWithDefaults() *StorageReplicationBlackout`

NewStorageReplicationBlackoutWithDefaults instantiates a new StorageReplicationBlackout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *StorageReplicationBlackout) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *StorageReplicationBlackout) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *StorageReplicationBlackout) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *StorageReplicationBlackout) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *StorageReplicationBlackout) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *StorageReplicationBlackout) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *StorageReplicationBlackout) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *StorageReplicationBlackout) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


