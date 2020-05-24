# BootSdCardAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lun** | Pointer to **int64** | The Logical Unit Number (LUN) of the device. | [optional] 
**Subtype** | Pointer to **string** | The subtype for the selected device type. | [optional] [default to "None"]

## Methods

### NewBootSdCardAllOf

`func NewBootSdCardAllOf() *BootSdCardAllOf`

NewBootSdCardAllOf instantiates a new BootSdCardAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootSdCardAllOfWithDefaults

`func NewBootSdCardAllOfWithDefaults() *BootSdCardAllOf`

NewBootSdCardAllOfWithDefaults instantiates a new BootSdCardAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLun

`func (o *BootSdCardAllOf) GetLun() int64`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *BootSdCardAllOf) GetLunOk() (*int64, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *BootSdCardAllOf) SetLun(v int64)`

SetLun sets Lun field to given value.

### HasLun

`func (o *BootSdCardAllOf) HasLun() bool`

HasLun returns a boolean if a field has been set.

### GetSubtype

`func (o *BootSdCardAllOf) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *BootSdCardAllOf) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *BootSdCardAllOf) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *BootSdCardAllOf) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


