# PciSwitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | It shows the device id of the switch. | [optional] [readonly] 
**Health** | Pointer to **string** | It shows the composite health of the switch. | [optional] [readonly] 
**NumOfAdaptors** | Pointer to **string** | It shows the number of gpus and pci adapters connected the switch. | [optional] [readonly] 
**PciAddress** | Pointer to **string** | It shows shows the PCI address of switch. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | It shows the PCI slot name for switch. | [optional] [readonly] 
**ProductName** | Pointer to **string** | It shows the model information for the switch. | [optional] [readonly] 
**ProductRevision** | Pointer to **string** | It shows the revision for the product. | [optional] [readonly] 
**SubDeviceId** | Pointer to **string** | It shows the sub device id of the switch. | [optional] [readonly] 
**SubVendorId** | Pointer to **string** | It shows the sub vendor id of the switch. | [optional] [readonly] 
**Temperature** | Pointer to **string** | It shows the current temperature of the switch. | [optional] [readonly] 
**Type** | Pointer to **string** | It shows the type inforamtion of switch. | [optional] 
**VendorId** | Pointer to **string** | It shows the vendor id of the switch. | [optional] [readonly] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**Links** | Pointer to [**[]PciLinkRelationship**](pci.Link.Relationship.md) | An array of relationships to pciLink resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewPciSwitch

`func NewPciSwitch() *PciSwitch`

NewPciSwitch instantiates a new PciSwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPciSwitchWithDefaults

`func NewPciSwitchWithDefaults() *PciSwitch`

NewPciSwitchWithDefaults instantiates a new PciSwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *PciSwitch) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *PciSwitch) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *PciSwitch) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *PciSwitch) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetHealth

`func (o *PciSwitch) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *PciSwitch) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *PciSwitch) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *PciSwitch) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetNumOfAdaptors

`func (o *PciSwitch) GetNumOfAdaptors() string`

GetNumOfAdaptors returns the NumOfAdaptors field if non-nil, zero value otherwise.

### GetNumOfAdaptorsOk

`func (o *PciSwitch) GetNumOfAdaptorsOk() (*string, bool)`

GetNumOfAdaptorsOk returns a tuple with the NumOfAdaptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfAdaptors

`func (o *PciSwitch) SetNumOfAdaptors(v string)`

SetNumOfAdaptors sets NumOfAdaptors field to given value.

### HasNumOfAdaptors

`func (o *PciSwitch) HasNumOfAdaptors() bool`

HasNumOfAdaptors returns a boolean if a field has been set.

### GetPciAddress

`func (o *PciSwitch) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *PciSwitch) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *PciSwitch) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *PciSwitch) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.

### GetPciSlot

`func (o *PciSwitch) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *PciSwitch) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *PciSwitch) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *PciSwitch) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetProductName

`func (o *PciSwitch) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *PciSwitch) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *PciSwitch) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *PciSwitch) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductRevision

`func (o *PciSwitch) GetProductRevision() string`

GetProductRevision returns the ProductRevision field if non-nil, zero value otherwise.

### GetProductRevisionOk

`func (o *PciSwitch) GetProductRevisionOk() (*string, bool)`

GetProductRevisionOk returns a tuple with the ProductRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductRevision

`func (o *PciSwitch) SetProductRevision(v string)`

SetProductRevision sets ProductRevision field to given value.

### HasProductRevision

`func (o *PciSwitch) HasProductRevision() bool`

HasProductRevision returns a boolean if a field has been set.

### GetSubDeviceId

`func (o *PciSwitch) GetSubDeviceId() string`

GetSubDeviceId returns the SubDeviceId field if non-nil, zero value otherwise.

### GetSubDeviceIdOk

`func (o *PciSwitch) GetSubDeviceIdOk() (*string, bool)`

GetSubDeviceIdOk returns a tuple with the SubDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDeviceId

`func (o *PciSwitch) SetSubDeviceId(v string)`

SetSubDeviceId sets SubDeviceId field to given value.

### HasSubDeviceId

`func (o *PciSwitch) HasSubDeviceId() bool`

HasSubDeviceId returns a boolean if a field has been set.

### GetSubVendorId

`func (o *PciSwitch) GetSubVendorId() string`

GetSubVendorId returns the SubVendorId field if non-nil, zero value otherwise.

### GetSubVendorIdOk

`func (o *PciSwitch) GetSubVendorIdOk() (*string, bool)`

GetSubVendorIdOk returns a tuple with the SubVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubVendorId

`func (o *PciSwitch) SetSubVendorId(v string)`

SetSubVendorId sets SubVendorId field to given value.

### HasSubVendorId

`func (o *PciSwitch) HasSubVendorId() bool`

HasSubVendorId returns a boolean if a field has been set.

### GetTemperature

`func (o *PciSwitch) GetTemperature() string`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *PciSwitch) GetTemperatureOk() (*string, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *PciSwitch) SetTemperature(v string)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *PciSwitch) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetType

`func (o *PciSwitch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PciSwitch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PciSwitch) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PciSwitch) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendorId

`func (o *PciSwitch) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *PciSwitch) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *PciSwitch) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *PciSwitch) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetComputeBoard

`func (o *PciSwitch) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *PciSwitch) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *PciSwitch) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *PciSwitch) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetLinks

`func (o *PciSwitch) GetLinks() []PciLinkRelationship`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PciSwitch) GetLinksOk() (*[]PciLinkRelationship, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PciSwitch) SetLinks(v []PciLinkRelationship)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PciSwitch) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *PciSwitch) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PciSwitch) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PciSwitch) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PciSwitch) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


