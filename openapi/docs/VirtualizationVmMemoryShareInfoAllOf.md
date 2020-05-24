# VirtualizationVmMemoryShareInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemLimit** | Pointer to **int64** | Limit on the memory sharing imposed (in Mbytes). | [optional] 
**MemOverheadLimit** | Pointer to **int64** | Limit on memory overhead imposed (in Mbytes). | [optional] 
**MemReservation** | Pointer to **int64** | Similar to CPU reservations (Mbytes). | [optional] 
**MemShares** | Pointer to **int64** | Similar to CPU Shares but applicable to memory. There is no unit for this value. It is a relative measure based on the settings for other resource pools. | [optional] 

## Methods

### NewVirtualizationVmMemoryShareInfoAllOf

`func NewVirtualizationVmMemoryShareInfoAllOf() *VirtualizationVmMemoryShareInfoAllOf`

NewVirtualizationVmMemoryShareInfoAllOf instantiates a new VirtualizationVmMemoryShareInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmMemoryShareInfoAllOfWithDefaults

`func NewVirtualizationVmMemoryShareInfoAllOfWithDefaults() *VirtualizationVmMemoryShareInfoAllOf`

NewVirtualizationVmMemoryShareInfoAllOfWithDefaults instantiates a new VirtualizationVmMemoryShareInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemLimit

`func (o *VirtualizationVmMemoryShareInfoAllOf) GetMemLimit() int64`

GetMemLimit returns the MemLimit field if non-nil, zero value otherwise.

### GetMemLimitOk

`func (o *VirtualizationVmMemoryShareInfoAllOf) GetMemLimitOk() (*int64, bool)`

GetMemLimitOk returns a tuple with the MemLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemLimit

`func (o *VirtualizationVmMemoryShareInfoAllOf) SetMemLimit(v int64)`

SetMemLimit sets MemLimit field to given value.

### HasMemLimit

`func (o *VirtualizationVmMemoryShareInfoAllOf) HasMemLimit() bool`

HasMemLimit returns a boolean if a field has been set.

### GetMemOverheadLimit

`func (o *VirtualizationVmMemoryShareInfoAllOf) GetMemOverheadLimit() int64`

GetMemOverheadLimit returns the MemOverheadLimit field if non-nil, zero value otherwise.

### GetMemOverheadLimitOk

`func (o *VirtualizationVmMemoryShareInfoAllOf) GetMemOverheadLimitOk() (*int64, bool)`

GetMemOverheadLimitOk returns a tuple with the MemOverheadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemOverheadLimit

`func (o *VirtualizationVmMemoryShareInfoAllOf) SetMemOverheadLimit(v int64)`

SetMemOverheadLimit sets MemOverheadLimit field to given value.

### HasMemOverheadLimit

`func (o *VirtualizationVmMemoryShareInfoAllOf) HasMemOverheadLimit() bool`

HasMemOverheadLimit returns a boolean if a field has been set.

### GetMemReservation

`func (o *VirtualizationVmMemoryShareInfoAllOf) GetMemReservation() int64`

GetMemReservation returns the MemReservation field if non-nil, zero value otherwise.

### GetMemReservationOk

`func (o *VirtualizationVmMemoryShareInfoAllOf) GetMemReservationOk() (*int64, bool)`

GetMemReservationOk returns a tuple with the MemReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemReservation

`func (o *VirtualizationVmMemoryShareInfoAllOf) SetMemReservation(v int64)`

SetMemReservation sets MemReservation field to given value.

### HasMemReservation

`func (o *VirtualizationVmMemoryShareInfoAllOf) HasMemReservation() bool`

HasMemReservation returns a boolean if a field has been set.

### GetMemShares

`func (o *VirtualizationVmMemoryShareInfoAllOf) GetMemShares() int64`

GetMemShares returns the MemShares field if non-nil, zero value otherwise.

### GetMemSharesOk

`func (o *VirtualizationVmMemoryShareInfoAllOf) GetMemSharesOk() (*int64, bool)`

GetMemSharesOk returns a tuple with the MemShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemShares

`func (o *VirtualizationVmMemoryShareInfoAllOf) SetMemShares(v int64)`

SetMemShares sets MemShares field to given value.

### HasMemShares

`func (o *VirtualizationVmMemoryShareInfoAllOf) HasMemShares() bool`

HasMemShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


