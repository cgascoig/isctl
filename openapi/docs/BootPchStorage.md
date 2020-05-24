# BootPchStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lun** | Pointer to **int64** | The Logical Unit Number (LUN) of the device. It is the Virtual Drive number for Cisco UCS C-Series Servers. Virtual Drive number is found in storage inventory. | [optional] 

## Methods

### NewBootPchStorage

`func NewBootPchStorage() *BootPchStorage`

NewBootPchStorage instantiates a new BootPchStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootPchStorageWithDefaults

`func NewBootPchStorageWithDefaults() *BootPchStorage`

NewBootPchStorageWithDefaults instantiates a new BootPchStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLun

`func (o *BootPchStorage) GetLun() int64`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *BootPchStorage) GetLunOk() (*int64, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *BootPchStorage) SetLun(v int64)`

SetLun sets Lun field to given value.

### HasLun

`func (o *BootPchStorage) HasLun() bool`

HasLun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


