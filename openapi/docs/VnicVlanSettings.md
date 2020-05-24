# VnicVlanSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVlan** | Pointer to **int64** | Default VLAN ID of the virtual interface. Setting the ID to 0 will not associate any default VLAN to the traffic on the virtual interface. | [optional] 
**Mode** | Pointer to **string** | Option to determine if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic. | [optional] [default to "ACCESS"]

## Methods

### NewVnicVlanSettings

`func NewVnicVlanSettings() *VnicVlanSettings`

NewVnicVlanSettings instantiates a new VnicVlanSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicVlanSettingsWithDefaults

`func NewVnicVlanSettingsWithDefaults() *VnicVlanSettings`

NewVnicVlanSettingsWithDefaults instantiates a new VnicVlanSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultVlan

`func (o *VnicVlanSettings) GetDefaultVlan() int64`

GetDefaultVlan returns the DefaultVlan field if non-nil, zero value otherwise.

### GetDefaultVlanOk

`func (o *VnicVlanSettings) GetDefaultVlanOk() (*int64, bool)`

GetDefaultVlanOk returns a tuple with the DefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlan

`func (o *VnicVlanSettings) SetDefaultVlan(v int64)`

SetDefaultVlan sets DefaultVlan field to given value.

### HasDefaultVlan

`func (o *VnicVlanSettings) HasDefaultVlan() bool`

HasDefaultVlan returns a boolean if a field has been set.

### GetMode

`func (o *VnicVlanSettings) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VnicVlanSettings) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VnicVlanSettings) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VnicVlanSettings) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


