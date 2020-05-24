# VirtualizationVmCpuShareInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuLimit** | Pointer to **int64** | Upper limit on CPU allocation (MHz). | [optional] 
**CpuOverheadLimit** | Pointer to **int64** | Amount of CPU for virtualization overhead. | [optional] 
**CpuReservation** | Pointer to **int64** | Guaranteed minimum allocation of CPU resource (MHz). | [optional] 
**CpuShares** | Pointer to **int64** | Shows the relative importance of a VM. There is no unit for this value. It is a relative measure based on the settings for other resource pools. For more information, see VMware documentation. | [optional] 

## Methods

### NewVirtualizationVmCpuShareInfoAllOf

`func NewVirtualizationVmCpuShareInfoAllOf() *VirtualizationVmCpuShareInfoAllOf`

NewVirtualizationVmCpuShareInfoAllOf instantiates a new VirtualizationVmCpuShareInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmCpuShareInfoAllOfWithDefaults

`func NewVirtualizationVmCpuShareInfoAllOfWithDefaults() *VirtualizationVmCpuShareInfoAllOf`

NewVirtualizationVmCpuShareInfoAllOfWithDefaults instantiates a new VirtualizationVmCpuShareInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuLimit

`func (o *VirtualizationVmCpuShareInfoAllOf) GetCpuLimit() int64`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *VirtualizationVmCpuShareInfoAllOf) GetCpuLimitOk() (*int64, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *VirtualizationVmCpuShareInfoAllOf) SetCpuLimit(v int64)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *VirtualizationVmCpuShareInfoAllOf) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### GetCpuOverheadLimit

`func (o *VirtualizationVmCpuShareInfoAllOf) GetCpuOverheadLimit() int64`

GetCpuOverheadLimit returns the CpuOverheadLimit field if non-nil, zero value otherwise.

### GetCpuOverheadLimitOk

`func (o *VirtualizationVmCpuShareInfoAllOf) GetCpuOverheadLimitOk() (*int64, bool)`

GetCpuOverheadLimitOk returns a tuple with the CpuOverheadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOverheadLimit

`func (o *VirtualizationVmCpuShareInfoAllOf) SetCpuOverheadLimit(v int64)`

SetCpuOverheadLimit sets CpuOverheadLimit field to given value.

### HasCpuOverheadLimit

`func (o *VirtualizationVmCpuShareInfoAllOf) HasCpuOverheadLimit() bool`

HasCpuOverheadLimit returns a boolean if a field has been set.

### GetCpuReservation

`func (o *VirtualizationVmCpuShareInfoAllOf) GetCpuReservation() int64`

GetCpuReservation returns the CpuReservation field if non-nil, zero value otherwise.

### GetCpuReservationOk

`func (o *VirtualizationVmCpuShareInfoAllOf) GetCpuReservationOk() (*int64, bool)`

GetCpuReservationOk returns a tuple with the CpuReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuReservation

`func (o *VirtualizationVmCpuShareInfoAllOf) SetCpuReservation(v int64)`

SetCpuReservation sets CpuReservation field to given value.

### HasCpuReservation

`func (o *VirtualizationVmCpuShareInfoAllOf) HasCpuReservation() bool`

HasCpuReservation returns a boolean if a field has been set.

### GetCpuShares

`func (o *VirtualizationVmCpuShareInfoAllOf) GetCpuShares() int64`

GetCpuShares returns the CpuShares field if non-nil, zero value otherwise.

### GetCpuSharesOk

`func (o *VirtualizationVmCpuShareInfoAllOf) GetCpuSharesOk() (*int64, bool)`

GetCpuSharesOk returns a tuple with the CpuShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuShares

`func (o *VirtualizationVmCpuShareInfoAllOf) SetCpuShares(v int64)`

SetCpuShares sets CpuShares field to given value.

### HasCpuShares

`func (o *VirtualizationVmCpuShareInfoAllOf) HasCpuShares() bool`

HasCpuShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


