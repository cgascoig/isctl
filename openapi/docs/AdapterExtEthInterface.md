# AdapterExtEthInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminState** | Pointer to **string** | The administrative state of the physical port. | [optional] [readonly] 
**EpDn** | Pointer to **string** |  | [optional] [readonly] 
**ExtEthInterfaceId** | Pointer to **string** |  | [optional] [readonly] 
**InterfaceType** | Pointer to **string** |  | [optional] [readonly] 
**MacAddress** | Pointer to **string** |  | [optional] [readonly] 
**OperState** | Pointer to **string** | The operational state of the physical port. | [optional] [readonly] 
**PeerDn** | Pointer to **string** |  | [optional] [readonly] 
**AdapterUnit** | Pointer to [**AdapterUnitRelationship**](adapter.Unit.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAdapterExtEthInterface

`func NewAdapterExtEthInterface() *AdapterExtEthInterface`

NewAdapterExtEthInterface instantiates a new AdapterExtEthInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterExtEthInterfaceWithDefaults

`func NewAdapterExtEthInterfaceWithDefaults() *AdapterExtEthInterface`

NewAdapterExtEthInterfaceWithDefaults instantiates a new AdapterExtEthInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminState

`func (o *AdapterExtEthInterface) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *AdapterExtEthInterface) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *AdapterExtEthInterface) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *AdapterExtEthInterface) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetEpDn

`func (o *AdapterExtEthInterface) GetEpDn() string`

GetEpDn returns the EpDn field if non-nil, zero value otherwise.

### GetEpDnOk

`func (o *AdapterExtEthInterface) GetEpDnOk() (*string, bool)`

GetEpDnOk returns a tuple with the EpDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpDn

`func (o *AdapterExtEthInterface) SetEpDn(v string)`

SetEpDn sets EpDn field to given value.

### HasEpDn

`func (o *AdapterExtEthInterface) HasEpDn() bool`

HasEpDn returns a boolean if a field has been set.

### GetExtEthInterfaceId

`func (o *AdapterExtEthInterface) GetExtEthInterfaceId() string`

GetExtEthInterfaceId returns the ExtEthInterfaceId field if non-nil, zero value otherwise.

### GetExtEthInterfaceIdOk

`func (o *AdapterExtEthInterface) GetExtEthInterfaceIdOk() (*string, bool)`

GetExtEthInterfaceIdOk returns a tuple with the ExtEthInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtEthInterfaceId

`func (o *AdapterExtEthInterface) SetExtEthInterfaceId(v string)`

SetExtEthInterfaceId sets ExtEthInterfaceId field to given value.

### HasExtEthInterfaceId

`func (o *AdapterExtEthInterface) HasExtEthInterfaceId() bool`

HasExtEthInterfaceId returns a boolean if a field has been set.

### GetInterfaceType

`func (o *AdapterExtEthInterface) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *AdapterExtEthInterface) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *AdapterExtEthInterface) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *AdapterExtEthInterface) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetMacAddress

`func (o *AdapterExtEthInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *AdapterExtEthInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *AdapterExtEthInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *AdapterExtEthInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetOperState

`func (o *AdapterExtEthInterface) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *AdapterExtEthInterface) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *AdapterExtEthInterface) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *AdapterExtEthInterface) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPeerDn

`func (o *AdapterExtEthInterface) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *AdapterExtEthInterface) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *AdapterExtEthInterface) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *AdapterExtEthInterface) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetAdapterUnit

`func (o *AdapterExtEthInterface) GetAdapterUnit() AdapterUnitRelationship`

GetAdapterUnit returns the AdapterUnit field if non-nil, zero value otherwise.

### GetAdapterUnitOk

`func (o *AdapterExtEthInterface) GetAdapterUnitOk() (*AdapterUnitRelationship, bool)`

GetAdapterUnitOk returns a tuple with the AdapterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnit

`func (o *AdapterExtEthInterface) SetAdapterUnit(v AdapterUnitRelationship)`

SetAdapterUnit sets AdapterUnit field to given value.

### HasAdapterUnit

`func (o *AdapterExtEthInterface) HasAdapterUnit() bool`

HasAdapterUnit returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AdapterExtEthInterface) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AdapterExtEthInterface) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AdapterExtEthInterface) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AdapterExtEthInterface) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


