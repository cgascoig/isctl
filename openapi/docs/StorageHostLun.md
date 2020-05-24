# StorageHostLun

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

### NewStorageHostLun

`func NewStorageHostLun() *StorageHostLun`

NewStorageHostLun instantiates a new StorageHostLun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHostLunWithDefaults

`func NewStorageHostLunWithDefaults() *StorageHostLun`

NewStorageHostLunWithDefaults instantiates a new StorageHostLun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHlu

`func (o *StorageHostLun) GetHlu() int64`

GetHlu returns the Hlu field if non-nil, zero value otherwise.

### GetHluOk

`func (o *StorageHostLun) GetHluOk() (*int64, bool)`

GetHluOk returns a tuple with the Hlu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlu

`func (o *StorageHostLun) SetHlu(v int64)`

SetHlu sets Hlu field to given value.

### HasHlu

`func (o *StorageHostLun) HasHlu() bool`

HasHlu returns a boolean if a field has been set.

### GetHostName

`func (o *StorageHostLun) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *StorageHostLun) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *StorageHostLun) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *StorageHostLun) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetVolumeName

`func (o *StorageHostLun) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *StorageHostLun) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *StorageHostLun) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *StorageHostLun) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetHost

`func (o *StorageHostLun) GetHost() StorageHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StorageHostLun) GetHostOk() (*StorageHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StorageHostLun) SetHost(v StorageHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *StorageHostLun) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetStorageArray

`func (o *StorageHostLun) GetStorageArray() StorageGenericArrayRelationship`

GetStorageArray returns the StorageArray field if non-nil, zero value otherwise.

### GetStorageArrayOk

`func (o *StorageHostLun) GetStorageArrayOk() (*StorageGenericArrayRelationship, bool)`

GetStorageArrayOk returns a tuple with the StorageArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageArray

`func (o *StorageHostLun) SetStorageArray(v StorageGenericArrayRelationship)`

SetStorageArray sets StorageArray field to given value.

### HasStorageArray

`func (o *StorageHostLun) HasStorageArray() bool`

HasStorageArray returns a boolean if a field has been set.

### GetVolume

`func (o *StorageHostLun) GetVolume() StorageVolumeRelationship`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StorageHostLun) GetVolumeOk() (*StorageVolumeRelationship, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StorageHostLun) SetVolume(v StorageVolumeRelationship)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StorageHostLun) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


