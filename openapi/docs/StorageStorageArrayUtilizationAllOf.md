# StorageStorageArrayUtilizationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataReduction** | Pointer to **float32** | Ratio of mapped sectors within a volume versus the amount of physical space the data occupies after data compression and deduplication. The data reduction ratio does not include thin provisioning savings. For example, a data reduction ratio of 5.0 means that for every 5 MB the host writes to the array, 1 MB is stored on the array&#39;s flash modules. | [optional] [readonly] 
**Parity** | Pointer to **float32** | Percentage of data that is fully protected. The percentage value will drop below 100% if the data is not fully protected. | [optional] [readonly] 
**Provisioned** | Pointer to **int64** | Total provisioned storage capacity in Pure FlashArray, represented in bytes. | [optional] [readonly] 
**Shared** | Pointer to **int64** | Physical space occupied by deduplicated data, represented in bytes. The space is shared with other volumes and snapshots as a result of data deduplication. | [optional] [readonly] 
**Snapshot** | Pointer to **int64** | Physical space occupied by the snapshots, represented in bytes. | [optional] [readonly] 
**System** | Pointer to **int64** | Physical space occupied by internal array metadata, represented in bytes. | [optional] [readonly] 
**ThinProvisioned** | Pointer to **float32** | Percentage of volume sectors that do not contain host-written data because the hosts have not written data to them or the sectors have been explicitly trimmed. | [optional] [readonly] 
**TotalReduction** | Pointer to **float32** | Ratio of provisioned sectors within a volume versus the amount of physical space the data occupies after reduction via data compression and deduplication and with thin provisioning savings. Total reduction is data reduction with thin provisioning savings. For example, a total reduction ratio of 10.0 means that for every 10 MB of provisioned space, 1 MB is stored on the array&#39;s flash modules. | [optional] [readonly] 
**Volume** | Pointer to **int64** | Physical space occupied by volume data, excluding shared, array metadata and snapshots. Size is represented in bytes. | [optional] [readonly] 

## Methods

### NewStorageStorageArrayUtilizationAllOf

`func NewStorageStorageArrayUtilizationAllOf() *StorageStorageArrayUtilizationAllOf`

NewStorageStorageArrayUtilizationAllOf instantiates a new StorageStorageArrayUtilizationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageStorageArrayUtilizationAllOfWithDefaults

`func NewStorageStorageArrayUtilizationAllOfWithDefaults() *StorageStorageArrayUtilizationAllOf`

NewStorageStorageArrayUtilizationAllOfWithDefaults instantiates a new StorageStorageArrayUtilizationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataReduction

`func (o *StorageStorageArrayUtilizationAllOf) GetDataReduction() float32`

GetDataReduction returns the DataReduction field if non-nil, zero value otherwise.

### GetDataReductionOk

`func (o *StorageStorageArrayUtilizationAllOf) GetDataReductionOk() (*float32, bool)`

GetDataReductionOk returns a tuple with the DataReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReduction

`func (o *StorageStorageArrayUtilizationAllOf) SetDataReduction(v float32)`

SetDataReduction sets DataReduction field to given value.

### HasDataReduction

`func (o *StorageStorageArrayUtilizationAllOf) HasDataReduction() bool`

HasDataReduction returns a boolean if a field has been set.

### GetParity

`func (o *StorageStorageArrayUtilizationAllOf) GetParity() float32`

GetParity returns the Parity field if non-nil, zero value otherwise.

### GetParityOk

`func (o *StorageStorageArrayUtilizationAllOf) GetParityOk() (*float32, bool)`

GetParityOk returns a tuple with the Parity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParity

`func (o *StorageStorageArrayUtilizationAllOf) SetParity(v float32)`

SetParity sets Parity field to given value.

### HasParity

`func (o *StorageStorageArrayUtilizationAllOf) HasParity() bool`

HasParity returns a boolean if a field has been set.

### GetProvisioned

`func (o *StorageStorageArrayUtilizationAllOf) GetProvisioned() int64`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *StorageStorageArrayUtilizationAllOf) GetProvisionedOk() (*int64, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *StorageStorageArrayUtilizationAllOf) SetProvisioned(v int64)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *StorageStorageArrayUtilizationAllOf) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.

### GetShared

`func (o *StorageStorageArrayUtilizationAllOf) GetShared() int64`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *StorageStorageArrayUtilizationAllOf) GetSharedOk() (*int64, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *StorageStorageArrayUtilizationAllOf) SetShared(v int64)`

SetShared sets Shared field to given value.

### HasShared

`func (o *StorageStorageArrayUtilizationAllOf) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSnapshot

`func (o *StorageStorageArrayUtilizationAllOf) GetSnapshot() int64`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *StorageStorageArrayUtilizationAllOf) GetSnapshotOk() (*int64, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *StorageStorageArrayUtilizationAllOf) SetSnapshot(v int64)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *StorageStorageArrayUtilizationAllOf) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetSystem

`func (o *StorageStorageArrayUtilizationAllOf) GetSystem() int64`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *StorageStorageArrayUtilizationAllOf) GetSystemOk() (*int64, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *StorageStorageArrayUtilizationAllOf) SetSystem(v int64)`

SetSystem sets System field to given value.

### HasSystem

`func (o *StorageStorageArrayUtilizationAllOf) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetThinProvisioned

`func (o *StorageStorageArrayUtilizationAllOf) GetThinProvisioned() float32`

GetThinProvisioned returns the ThinProvisioned field if non-nil, zero value otherwise.

### GetThinProvisionedOk

`func (o *StorageStorageArrayUtilizationAllOf) GetThinProvisionedOk() (*float32, bool)`

GetThinProvisionedOk returns a tuple with the ThinProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioned

`func (o *StorageStorageArrayUtilizationAllOf) SetThinProvisioned(v float32)`

SetThinProvisioned sets ThinProvisioned field to given value.

### HasThinProvisioned

`func (o *StorageStorageArrayUtilizationAllOf) HasThinProvisioned() bool`

HasThinProvisioned returns a boolean if a field has been set.

### GetTotalReduction

`func (o *StorageStorageArrayUtilizationAllOf) GetTotalReduction() float32`

GetTotalReduction returns the TotalReduction field if non-nil, zero value otherwise.

### GetTotalReductionOk

`func (o *StorageStorageArrayUtilizationAllOf) GetTotalReductionOk() (*float32, bool)`

GetTotalReductionOk returns a tuple with the TotalReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReduction

`func (o *StorageStorageArrayUtilizationAllOf) SetTotalReduction(v float32)`

SetTotalReduction sets TotalReduction field to given value.

### HasTotalReduction

`func (o *StorageStorageArrayUtilizationAllOf) HasTotalReduction() bool`

HasTotalReduction returns a boolean if a field has been set.

### GetVolume

`func (o *StorageStorageArrayUtilizationAllOf) GetVolume() int64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StorageStorageArrayUtilizationAllOf) GetVolumeOk() (*int64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StorageStorageArrayUtilizationAllOf) SetVolume(v int64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StorageStorageArrayUtilizationAllOf) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


