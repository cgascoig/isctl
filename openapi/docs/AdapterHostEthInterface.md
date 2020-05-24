# AdapterHostEthInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminState** | Pointer to **string** |  | [optional] [readonly] 
**EpDn** | Pointer to **string** |  | [optional] [readonly] 
**HostEthInterfaceId** | Pointer to **int64** |  | [optional] [readonly] 
**InterfaceType** | Pointer to **string** |  | [optional] [readonly] 
**MacAddress** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**OperState** | Pointer to **string** |  | [optional] [readonly] 
**Operability** | Pointer to **string** |  | [optional] [readonly] 
**OriginalMacAddress** | Pointer to **string** |  | [optional] [readonly] 
**PciAddr** | Pointer to **string** |  | [optional] [readonly] 
**PeerDn** | Pointer to **string** |  | [optional] [readonly] 
**VirtualizationPreference** | Pointer to **string** |  | [optional] [readonly] 
**VnicDn** | Pointer to **string** |  | [optional] [readonly] 
**AdapterUnit** | Pointer to [**AdapterUnitRelationship**](adapter.Unit.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAdapterHostEthInterface

`func NewAdapterHostEthInterface() *AdapterHostEthInterface`

NewAdapterHostEthInterface instantiates a new AdapterHostEthInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterHostEthInterfaceWithDefaults

`func NewAdapterHostEthInterfaceWithDefaults() *AdapterHostEthInterface`

NewAdapterHostEthInterfaceWithDefaults instantiates a new AdapterHostEthInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminState

`func (o *AdapterHostEthInterface) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *AdapterHostEthInterface) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *AdapterHostEthInterface) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *AdapterHostEthInterface) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetEpDn

`func (o *AdapterHostEthInterface) GetEpDn() string`

GetEpDn returns the EpDn field if non-nil, zero value otherwise.

### GetEpDnOk

`func (o *AdapterHostEthInterface) GetEpDnOk() (*string, bool)`

GetEpDnOk returns a tuple with the EpDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpDn

`func (o *AdapterHostEthInterface) SetEpDn(v string)`

SetEpDn sets EpDn field to given value.

### HasEpDn

`func (o *AdapterHostEthInterface) HasEpDn() bool`

HasEpDn returns a boolean if a field has been set.

### GetHostEthInterfaceId

`func (o *AdapterHostEthInterface) GetHostEthInterfaceId() int64`

GetHostEthInterfaceId returns the HostEthInterfaceId field if non-nil, zero value otherwise.

### GetHostEthInterfaceIdOk

`func (o *AdapterHostEthInterface) GetHostEthInterfaceIdOk() (*int64, bool)`

GetHostEthInterfaceIdOk returns a tuple with the HostEthInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostEthInterfaceId

`func (o *AdapterHostEthInterface) SetHostEthInterfaceId(v int64)`

SetHostEthInterfaceId sets HostEthInterfaceId field to given value.

### HasHostEthInterfaceId

`func (o *AdapterHostEthInterface) HasHostEthInterfaceId() bool`

HasHostEthInterfaceId returns a boolean if a field has been set.

### GetInterfaceType

`func (o *AdapterHostEthInterface) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *AdapterHostEthInterface) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *AdapterHostEthInterface) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *AdapterHostEthInterface) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetMacAddress

`func (o *AdapterHostEthInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *AdapterHostEthInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *AdapterHostEthInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *AdapterHostEthInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetName

`func (o *AdapterHostEthInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdapterHostEthInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdapterHostEthInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdapterHostEthInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *AdapterHostEthInterface) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *AdapterHostEthInterface) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *AdapterHostEthInterface) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *AdapterHostEthInterface) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *AdapterHostEthInterface) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *AdapterHostEthInterface) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *AdapterHostEthInterface) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *AdapterHostEthInterface) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetOriginalMacAddress

`func (o *AdapterHostEthInterface) GetOriginalMacAddress() string`

GetOriginalMacAddress returns the OriginalMacAddress field if non-nil, zero value otherwise.

### GetOriginalMacAddressOk

`func (o *AdapterHostEthInterface) GetOriginalMacAddressOk() (*string, bool)`

GetOriginalMacAddressOk returns a tuple with the OriginalMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMacAddress

`func (o *AdapterHostEthInterface) SetOriginalMacAddress(v string)`

SetOriginalMacAddress sets OriginalMacAddress field to given value.

### HasOriginalMacAddress

`func (o *AdapterHostEthInterface) HasOriginalMacAddress() bool`

HasOriginalMacAddress returns a boolean if a field has been set.

### GetPciAddr

`func (o *AdapterHostEthInterface) GetPciAddr() string`

GetPciAddr returns the PciAddr field if non-nil, zero value otherwise.

### GetPciAddrOk

`func (o *AdapterHostEthInterface) GetPciAddrOk() (*string, bool)`

GetPciAddrOk returns a tuple with the PciAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddr

`func (o *AdapterHostEthInterface) SetPciAddr(v string)`

SetPciAddr sets PciAddr field to given value.

### HasPciAddr

`func (o *AdapterHostEthInterface) HasPciAddr() bool`

HasPciAddr returns a boolean if a field has been set.

### GetPeerDn

`func (o *AdapterHostEthInterface) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *AdapterHostEthInterface) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *AdapterHostEthInterface) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *AdapterHostEthInterface) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetVirtualizationPreference

`func (o *AdapterHostEthInterface) GetVirtualizationPreference() string`

GetVirtualizationPreference returns the VirtualizationPreference field if non-nil, zero value otherwise.

### GetVirtualizationPreferenceOk

`func (o *AdapterHostEthInterface) GetVirtualizationPreferenceOk() (*string, bool)`

GetVirtualizationPreferenceOk returns a tuple with the VirtualizationPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationPreference

`func (o *AdapterHostEthInterface) SetVirtualizationPreference(v string)`

SetVirtualizationPreference sets VirtualizationPreference field to given value.

### HasVirtualizationPreference

`func (o *AdapterHostEthInterface) HasVirtualizationPreference() bool`

HasVirtualizationPreference returns a boolean if a field has been set.

### GetVnicDn

`func (o *AdapterHostEthInterface) GetVnicDn() string`

GetVnicDn returns the VnicDn field if non-nil, zero value otherwise.

### GetVnicDnOk

`func (o *AdapterHostEthInterface) GetVnicDnOk() (*string, bool)`

GetVnicDnOk returns a tuple with the VnicDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnicDn

`func (o *AdapterHostEthInterface) SetVnicDn(v string)`

SetVnicDn sets VnicDn field to given value.

### HasVnicDn

`func (o *AdapterHostEthInterface) HasVnicDn() bool`

HasVnicDn returns a boolean if a field has been set.

### GetAdapterUnit

`func (o *AdapterHostEthInterface) GetAdapterUnit() AdapterUnitRelationship`

GetAdapterUnit returns the AdapterUnit field if non-nil, zero value otherwise.

### GetAdapterUnitOk

`func (o *AdapterHostEthInterface) GetAdapterUnitOk() (*AdapterUnitRelationship, bool)`

GetAdapterUnitOk returns a tuple with the AdapterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnit

`func (o *AdapterHostEthInterface) SetAdapterUnit(v AdapterUnitRelationship)`

SetAdapterUnit sets AdapterUnit field to given value.

### HasAdapterUnit

`func (o *AdapterHostEthInterface) HasAdapterUnit() bool`

HasAdapterUnit returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AdapterHostEthInterface) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AdapterHostEthInterface) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AdapterHostEthInterface) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AdapterHostEthInterface) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


