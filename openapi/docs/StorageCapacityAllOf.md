# StorageCapacityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **int64** | Total consumable storage capacity represented in bytes. System may reserve some space for internal purposes which is excluded from total capacity. | [optional] [readonly] 
**Free** | Pointer to **int64** | Unused space available for applications to consume, represented in bytes. | [optional] [readonly] 
**Total** | Pointer to **int64** | Total storage capacity, represented in bytes. It is set by the component manufacturer. | [optional] [readonly] 
**Used** | Pointer to **int64** | Used or consumed storage capacity, represented in bytes. | [optional] [readonly] 

## Methods

### NewStorageCapacityAllOf

`func NewStorageCapacityAllOf() *StorageCapacityAllOf`

NewStorageCapacityAllOf instantiates a new StorageCapacityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageCapacityAllOfWithDefaults

`func NewStorageCapacityAllOfWithDefaults() *StorageCapacityAllOf`

NewStorageCapacityAllOfWithDefaults instantiates a new StorageCapacityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *StorageCapacityAllOf) GetAvailable() int64`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *StorageCapacityAllOf) GetAvailableOk() (*int64, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *StorageCapacityAllOf) SetAvailable(v int64)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *StorageCapacityAllOf) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetFree

`func (o *StorageCapacityAllOf) GetFree() int64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *StorageCapacityAllOf) GetFreeOk() (*int64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *StorageCapacityAllOf) SetFree(v int64)`

SetFree sets Free field to given value.

### HasFree

`func (o *StorageCapacityAllOf) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetTotal

`func (o *StorageCapacityAllOf) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *StorageCapacityAllOf) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *StorageCapacityAllOf) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *StorageCapacityAllOf) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *StorageCapacityAllOf) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *StorageCapacityAllOf) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *StorageCapacityAllOf) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *StorageCapacityAllOf) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


