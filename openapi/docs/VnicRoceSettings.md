# VnicRoceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If enabled sets RDMA over Converged Ethernet (RoCE) on this virtual interface. | [optional] 
**MemoryRegions** | Pointer to **int64** | The number of memory regions per adapter. Recommended value &#x3D; integer power of 2. | [optional] 
**QueuePairs** | Pointer to **int64** | The number of queue pairs per adapter. Recommended value &#x3D; integer power of 2. | [optional] 
**ResourceGroups** | Pointer to **int64** | The number of resource groups per adapter. Recommended value &#x3D; be an integer power of 2 greater than or equal to the number of CPU cores on the system for optimum performance. | [optional] 

## Methods

### NewVnicRoceSettings

`func NewVnicRoceSettings() *VnicRoceSettings`

NewVnicRoceSettings instantiates a new VnicRoceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicRoceSettingsWithDefaults

`func NewVnicRoceSettingsWithDefaults() *VnicRoceSettings`

NewVnicRoceSettingsWithDefaults instantiates a new VnicRoceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *VnicRoceSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VnicRoceSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VnicRoceSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VnicRoceSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMemoryRegions

`func (o *VnicRoceSettings) GetMemoryRegions() int64`

GetMemoryRegions returns the MemoryRegions field if non-nil, zero value otherwise.

### GetMemoryRegionsOk

`func (o *VnicRoceSettings) GetMemoryRegionsOk() (*int64, bool)`

GetMemoryRegionsOk returns a tuple with the MemoryRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRegions

`func (o *VnicRoceSettings) SetMemoryRegions(v int64)`

SetMemoryRegions sets MemoryRegions field to given value.

### HasMemoryRegions

`func (o *VnicRoceSettings) HasMemoryRegions() bool`

HasMemoryRegions returns a boolean if a field has been set.

### GetQueuePairs

`func (o *VnicRoceSettings) GetQueuePairs() int64`

GetQueuePairs returns the QueuePairs field if non-nil, zero value otherwise.

### GetQueuePairsOk

`func (o *VnicRoceSettings) GetQueuePairsOk() (*int64, bool)`

GetQueuePairsOk returns a tuple with the QueuePairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePairs

`func (o *VnicRoceSettings) SetQueuePairs(v int64)`

SetQueuePairs sets QueuePairs field to given value.

### HasQueuePairs

`func (o *VnicRoceSettings) HasQueuePairs() bool`

HasQueuePairs returns a boolean if a field has been set.

### GetResourceGroups

`func (o *VnicRoceSettings) GetResourceGroups() int64`

GetResourceGroups returns the ResourceGroups field if non-nil, zero value otherwise.

### GetResourceGroupsOk

`func (o *VnicRoceSettings) GetResourceGroupsOk() (*int64, bool)`

GetResourceGroupsOk returns a tuple with the ResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroups

`func (o *VnicRoceSettings) SetResourceGroups(v int64)`

SetResourceGroups sets ResourceGroups field to given value.

### HasResourceGroups

`func (o *VnicRoceSettings) HasResourceGroups() bool`

HasResourceGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


