# VirtualizationHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuInfo** | Pointer to [**VirtualizationCpuInfo**](virtualization.CpuInfo.md) |  | [optional] 
**HardwareInfo** | Pointer to [**InfraHardwareInfo**](infra.HardwareInfo.md) |  | [optional] 
**HypervisorType** | Pointer to **string** | Identifies the broad type of the underlying hypervisor. | [optional] [default to "Unknown"]
**Identity** | Pointer to **string** | The internally generated identity of this host. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is an MOR (managed object reference). | [optional] 
**MaintenanceMode** | Pointer to **bool** | Is this host in maintenance mode. Set to true or false. | [optional] 
**MemoryCapacity** | Pointer to [**VirtualizationMemoryCapacity**](virtualization.MemoryCapacity.md) |  | [optional] 
**Model** | Pointer to **string** | Commercial model information about this hardware. | [optional] 
**Name** | Pointer to **string** | Name of this host supplied by user. It is not the identity of the host. The name is subject to user manipulations. | [optional] 
**ProcessorCapacity** | Pointer to [**VirtualizationComputeCapacity**](virtualization.ComputeCapacity.md) |  | [optional] 
**ProductInfo** | Pointer to [**VirtualizationProductInfo**](virtualization.ProductInfo.md) |  | [optional] 
**Serial** | Pointer to **string** | Serial number of this host (internally generated). | [optional] 
**Status** | Pointer to **string** | Host health status, as reported by the hypervisor platform. | [optional] [default to "Unknown"]
**UpTime** | Pointer to **string** | The uptime of the host, stored as Duration (from w3c). | [optional] 
**Uuid** | Pointer to **string** | Universally unique identity of this host (example b3d4483b-5560-9342-8309-b486c9236610). Internally generated. | [optional] 
**Vendor** | Pointer to **string** | Commercial vendor details of this hardware. | [optional] 

## Methods

### NewVirtualizationHost

`func NewVirtualizationHost() *VirtualizationHost`

NewVirtualizationHost instantiates a new VirtualizationHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationHostWithDefaults

`func NewVirtualizationHostWithDefaults() *VirtualizationHost`

NewVirtualizationHostWithDefaults instantiates a new VirtualizationHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuInfo

`func (o *VirtualizationHost) GetCpuInfo() VirtualizationCpuInfo`

GetCpuInfo returns the CpuInfo field if non-nil, zero value otherwise.

### GetCpuInfoOk

`func (o *VirtualizationHost) GetCpuInfoOk() (*VirtualizationCpuInfo, bool)`

GetCpuInfoOk returns a tuple with the CpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuInfo

`func (o *VirtualizationHost) SetCpuInfo(v VirtualizationCpuInfo)`

SetCpuInfo sets CpuInfo field to given value.

### HasCpuInfo

`func (o *VirtualizationHost) HasCpuInfo() bool`

HasCpuInfo returns a boolean if a field has been set.

### GetHardwareInfo

`func (o *VirtualizationHost) GetHardwareInfo() InfraHardwareInfo`

GetHardwareInfo returns the HardwareInfo field if non-nil, zero value otherwise.

### GetHardwareInfoOk

`func (o *VirtualizationHost) GetHardwareInfoOk() (*InfraHardwareInfo, bool)`

GetHardwareInfoOk returns a tuple with the HardwareInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareInfo

`func (o *VirtualizationHost) SetHardwareInfo(v InfraHardwareInfo)`

SetHardwareInfo sets HardwareInfo field to given value.

### HasHardwareInfo

`func (o *VirtualizationHost) HasHardwareInfo() bool`

HasHardwareInfo returns a boolean if a field has been set.

### GetHypervisorType

`func (o *VirtualizationHost) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *VirtualizationHost) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *VirtualizationHost) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *VirtualizationHost) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationHost) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationHost) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationHost) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationHost) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetMaintenanceMode

`func (o *VirtualizationHost) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *VirtualizationHost) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *VirtualizationHost) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *VirtualizationHost) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetMemoryCapacity

`func (o *VirtualizationHost) GetMemoryCapacity() VirtualizationMemoryCapacity`

GetMemoryCapacity returns the MemoryCapacity field if non-nil, zero value otherwise.

### GetMemoryCapacityOk

`func (o *VirtualizationHost) GetMemoryCapacityOk() (*VirtualizationMemoryCapacity, bool)`

GetMemoryCapacityOk returns a tuple with the MemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacity

`func (o *VirtualizationHost) SetMemoryCapacity(v VirtualizationMemoryCapacity)`

SetMemoryCapacity sets MemoryCapacity field to given value.

### HasMemoryCapacity

`func (o *VirtualizationHost) HasMemoryCapacity() bool`

HasMemoryCapacity returns a boolean if a field has been set.

### GetModel

`func (o *VirtualizationHost) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VirtualizationHost) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VirtualizationHost) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *VirtualizationHost) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationHost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationHost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessorCapacity

`func (o *VirtualizationHost) GetProcessorCapacity() VirtualizationComputeCapacity`

GetProcessorCapacity returns the ProcessorCapacity field if non-nil, zero value otherwise.

### GetProcessorCapacityOk

`func (o *VirtualizationHost) GetProcessorCapacityOk() (*VirtualizationComputeCapacity, bool)`

GetProcessorCapacityOk returns a tuple with the ProcessorCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCapacity

`func (o *VirtualizationHost) SetProcessorCapacity(v VirtualizationComputeCapacity)`

SetProcessorCapacity sets ProcessorCapacity field to given value.

### HasProcessorCapacity

`func (o *VirtualizationHost) HasProcessorCapacity() bool`

HasProcessorCapacity returns a boolean if a field has been set.

### GetProductInfo

`func (o *VirtualizationHost) GetProductInfo() VirtualizationProductInfo`

GetProductInfo returns the ProductInfo field if non-nil, zero value otherwise.

### GetProductInfoOk

`func (o *VirtualizationHost) GetProductInfoOk() (*VirtualizationProductInfo, bool)`

GetProductInfoOk returns a tuple with the ProductInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInfo

`func (o *VirtualizationHost) SetProductInfo(v VirtualizationProductInfo)`

SetProductInfo sets ProductInfo field to given value.

### HasProductInfo

`func (o *VirtualizationHost) HasProductInfo() bool`

HasProductInfo returns a boolean if a field has been set.

### GetSerial

`func (o *VirtualizationHost) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *VirtualizationHost) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *VirtualizationHost) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *VirtualizationHost) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationHost) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationHost) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationHost) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationHost) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpTime

`func (o *VirtualizationHost) GetUpTime() string`

GetUpTime returns the UpTime field if non-nil, zero value otherwise.

### GetUpTimeOk

`func (o *VirtualizationHost) GetUpTimeOk() (*string, bool)`

GetUpTimeOk returns a tuple with the UpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTime

`func (o *VirtualizationHost) SetUpTime(v string)`

SetUpTime sets UpTime field to given value.

### HasUpTime

`func (o *VirtualizationHost) HasUpTime() bool`

HasUpTime returns a boolean if a field has been set.

### GetUuid

`func (o *VirtualizationHost) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualizationHost) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualizationHost) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualizationHost) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendor

`func (o *VirtualizationHost) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VirtualizationHost) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VirtualizationHost) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *VirtualizationHost) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


