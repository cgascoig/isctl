# VirtualizationVmCpuSocketInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoresPerSocket** | Pointer to **int64** | The number of core per CPU socket (value may be empty). | [optional] 
**NumCpus** | Pointer to **int64** | Number of CPUs allocated to this VM. | [optional] 
**NumSockets** | Pointer to **int64** | The number of CPU sockets allocated. | [optional] 

## Methods

### NewVirtualizationVmCpuSocketInfo

`func NewVirtualizationVmCpuSocketInfo() *VirtualizationVmCpuSocketInfo`

NewVirtualizationVmCpuSocketInfo instantiates a new VirtualizationVmCpuSocketInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmCpuSocketInfoWithDefaults

`func NewVirtualizationVmCpuSocketInfoWithDefaults() *VirtualizationVmCpuSocketInfo`

NewVirtualizationVmCpuSocketInfoWithDefaults instantiates a new VirtualizationVmCpuSocketInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoresPerSocket

`func (o *VirtualizationVmCpuSocketInfo) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *VirtualizationVmCpuSocketInfo) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *VirtualizationVmCpuSocketInfo) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *VirtualizationVmCpuSocketInfo) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetNumCpus

`func (o *VirtualizationVmCpuSocketInfo) GetNumCpus() int64`

GetNumCpus returns the NumCpus field if non-nil, zero value otherwise.

### GetNumCpusOk

`func (o *VirtualizationVmCpuSocketInfo) GetNumCpusOk() (*int64, bool)`

GetNumCpusOk returns a tuple with the NumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpus

`func (o *VirtualizationVmCpuSocketInfo) SetNumCpus(v int64)`

SetNumCpus sets NumCpus field to given value.

### HasNumCpus

`func (o *VirtualizationVmCpuSocketInfo) HasNumCpus() bool`

HasNumCpus returns a boolean if a field has been set.

### GetNumSockets

`func (o *VirtualizationVmCpuSocketInfo) GetNumSockets() int64`

GetNumSockets returns the NumSockets field if non-nil, zero value otherwise.

### GetNumSocketsOk

`func (o *VirtualizationVmCpuSocketInfo) GetNumSocketsOk() (*int64, bool)`

GetNumSocketsOk returns a tuple with the NumSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSockets

`func (o *VirtualizationVmCpuSocketInfo) SetNumSockets(v int64)`

SetNumSockets sets NumSockets field to given value.

### HasNumSockets

`func (o *VirtualizationVmCpuSocketInfo) HasNumSockets() bool`

HasNumSockets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


