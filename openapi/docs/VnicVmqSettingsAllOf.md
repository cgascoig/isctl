# VnicVmqSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.VmqSettings"]
**ObjectType** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.VmqSettings"]
**Enabled** | Pointer to **bool** | Enables VMQ feature on the virtual interface. | [optional] [default to false]
**MultiQueueSupport** | Pointer to **bool** | Enables Virtual Machine Multi-Queue feature on the virtual interface. VMMQ allows configuration of multiple I/O queues for a single VM and thus distributes traffic across multiple CPU cores in a VM. | [optional] [default to false]
**NumInterrupts** | Pointer to **int64** | The number of interrupt resources to be allocated. Recommended value is the number of CPU threads or logical processors available in the server. | [optional] [default to 16]
**NumSubVnics** | Pointer to **int64** | The number of sub vNICs to be created. | [optional] [default to 64]
**NumVmqs** | Pointer to **int64** | The number of hardware Virtual Machine Queues to be allocated. The number of VMQs per adapter must be one more than the maximum number of VM NICs. | [optional] [default to 4]
**VmmqAdapterPolicy** | Pointer to **string** | Ethernet Adapter policy to be associated with the Sub vNICs. The Transmit Queue and Receive Queue resource value of VMMQ adapter policy should be greater than or equal to the configured number of sub vNICs. | [optional] 

## Methods

### NewVnicVmqSettingsAllOf

`func NewVnicVmqSettingsAllOf(classId string, objectType string, ) *VnicVmqSettingsAllOf`

NewVnicVmqSettingsAllOf instantiates a new VnicVmqSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicVmqSettingsAllOfWithDefaults

`func NewVnicVmqSettingsAllOfWithDefaults() *VnicVmqSettingsAllOf`

NewVnicVmqSettingsAllOfWithDefaults instantiates a new VnicVmqSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicVmqSettingsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicVmqSettingsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicVmqSettingsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicVmqSettingsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicVmqSettingsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicVmqSettingsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *VnicVmqSettingsAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VnicVmqSettingsAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VnicVmqSettingsAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VnicVmqSettingsAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMultiQueueSupport

`func (o *VnicVmqSettingsAllOf) GetMultiQueueSupport() bool`

GetMultiQueueSupport returns the MultiQueueSupport field if non-nil, zero value otherwise.

### GetMultiQueueSupportOk

`func (o *VnicVmqSettingsAllOf) GetMultiQueueSupportOk() (*bool, bool)`

GetMultiQueueSupportOk returns a tuple with the MultiQueueSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiQueueSupport

`func (o *VnicVmqSettingsAllOf) SetMultiQueueSupport(v bool)`

SetMultiQueueSupport sets MultiQueueSupport field to given value.

### HasMultiQueueSupport

`func (o *VnicVmqSettingsAllOf) HasMultiQueueSupport() bool`

HasMultiQueueSupport returns a boolean if a field has been set.

### GetNumInterrupts

`func (o *VnicVmqSettingsAllOf) GetNumInterrupts() int64`

GetNumInterrupts returns the NumInterrupts field if non-nil, zero value otherwise.

### GetNumInterruptsOk

`func (o *VnicVmqSettingsAllOf) GetNumInterruptsOk() (*int64, bool)`

GetNumInterruptsOk returns a tuple with the NumInterrupts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInterrupts

`func (o *VnicVmqSettingsAllOf) SetNumInterrupts(v int64)`

SetNumInterrupts sets NumInterrupts field to given value.

### HasNumInterrupts

`func (o *VnicVmqSettingsAllOf) HasNumInterrupts() bool`

HasNumInterrupts returns a boolean if a field has been set.

### GetNumSubVnics

`func (o *VnicVmqSettingsAllOf) GetNumSubVnics() int64`

GetNumSubVnics returns the NumSubVnics field if non-nil, zero value otherwise.

### GetNumSubVnicsOk

`func (o *VnicVmqSettingsAllOf) GetNumSubVnicsOk() (*int64, bool)`

GetNumSubVnicsOk returns a tuple with the NumSubVnics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSubVnics

`func (o *VnicVmqSettingsAllOf) SetNumSubVnics(v int64)`

SetNumSubVnics sets NumSubVnics field to given value.

### HasNumSubVnics

`func (o *VnicVmqSettingsAllOf) HasNumSubVnics() bool`

HasNumSubVnics returns a boolean if a field has been set.

### GetNumVmqs

`func (o *VnicVmqSettingsAllOf) GetNumVmqs() int64`

GetNumVmqs returns the NumVmqs field if non-nil, zero value otherwise.

### GetNumVmqsOk

`func (o *VnicVmqSettingsAllOf) GetNumVmqsOk() (*int64, bool)`

GetNumVmqsOk returns a tuple with the NumVmqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVmqs

`func (o *VnicVmqSettingsAllOf) SetNumVmqs(v int64)`

SetNumVmqs sets NumVmqs field to given value.

### HasNumVmqs

`func (o *VnicVmqSettingsAllOf) HasNumVmqs() bool`

HasNumVmqs returns a boolean if a field has been set.

### GetVmmqAdapterPolicy

`func (o *VnicVmqSettingsAllOf) GetVmmqAdapterPolicy() string`

GetVmmqAdapterPolicy returns the VmmqAdapterPolicy field if non-nil, zero value otherwise.

### GetVmmqAdapterPolicyOk

`func (o *VnicVmqSettingsAllOf) GetVmmqAdapterPolicyOk() (*string, bool)`

GetVmmqAdapterPolicyOk returns a tuple with the VmmqAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmmqAdapterPolicy

`func (o *VnicVmqSettingsAllOf) SetVmmqAdapterPolicy(v string)`

SetVmmqAdapterPolicy sets VmmqAdapterPolicy field to given value.

### HasVmmqAdapterPolicy

`func (o *VnicVmqSettingsAllOf) HasVmmqAdapterPolicy() bool`

HasVmmqAdapterPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


