# StorageHostLunAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hlu** | Pointer to **int64** | Logical unit number (LUN) by which hosts address specified volume. Hlu is a decimal representation of the LUN from the endpoint. | [optional] [readonly] 
**HostName** | Pointer to **string** | Name of the host associated with LUN. | [optional] [readonly] 
**VolumeName** | Pointer to **string** | Name of the storage volume associated with LUN. | [optional] [readonly] 
**Host** | Pointer to [**StorageHostRelationship**](storage.Host.Relationship.md) |  | [optional] 
**StorageArray** | Pointer to [**StorageGenericArrayRelationship**](storage.GenericArray.Relationship.md) |  | [optional] 
**Volume** | Pointer to [**StorageVolumeRelationship**](storage.Volume.Relationship.md) |  | [optional] 

## Methods

### NewStorageHostLunAllOf

`func NewStorageHostLunAllOf() *StorageHostLunAllOf`

NewStorageHostLunAllOf instantiates a new StorageHostLunAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHostLunAllOfWithDefaults

`func NewStorageHostLunAllOfWithDefaults() *StorageHostLunAllOf`

NewStorageHostLunAllOfWithDefaults instantiates a new StorageHostLunAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHlu

`func (o *StorageHostLunAllOf) GetHlu() int64`

GetHlu returns the Hlu field if non-nil, zero value otherwise.

### GetHluOk

`func (o *StorageHostLunAllOf) GetHluOk() (*int64, bool)`

GetHluOk returns a tuple with the Hlu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlu

`func (o *StorageHostLunAllOf) SetHlu(v int64)`

SetHlu sets Hlu field to given value.

### HasHlu

`func (o *StorageHostLunAllOf) HasHlu() bool`

HasHlu returns a boolean if a field has been set.

### GetHostName

`func (o *StorageHostLunAllOf) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *StorageHostLunAllOf) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *StorageHostLunAllOf) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *StorageHostLunAllOf) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetVolumeName

`func (o *StorageHostLunAllOf) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *StorageHostLunAllOf) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *StorageHostLunAllOf) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *StorageHostLunAllOf) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetHost

`func (o *StorageHostLunAllOf) GetHost() StorageHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StorageHostLunAllOf) GetHostOk() (*StorageHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StorageHostLunAllOf) SetHost(v StorageHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *StorageHostLunAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetStorageArray

`func (o *StorageHostLunAllOf) GetStorageArray() StorageGenericArrayRelationship`

GetStorageArray returns the StorageArray field if non-nil, zero value otherwise.

### GetStorageArrayOk

`func (o *StorageHostLunAllOf) GetStorageArrayOk() (*StorageGenericArrayRelationship, bool)`

GetStorageArrayOk returns a tuple with the StorageArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageArray

`func (o *StorageHostLunAllOf) SetStorageArray(v StorageGenericArrayRelationship)`

SetStorageArray sets StorageArray field to given value.

### HasStorageArray

`func (o *StorageHostLunAllOf) HasStorageArray() bool`

HasStorageArray returns a boolean if a field has been set.

### GetVolume

`func (o *StorageHostLunAllOf) GetVolume() StorageVolumeRelationship`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StorageHostLunAllOf) GetVolumeOk() (*StorageVolumeRelationship, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StorageHostLunAllOf) SetVolume(v StorageVolumeRelationship)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StorageHostLunAllOf) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


