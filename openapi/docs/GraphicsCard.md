# GraphicsCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardId** | Pointer to **int64** | It shows the id of graphics card. | [optional] [readonly] 
**DeviceId** | Pointer to **int64** | It shows the device id of grphics card. | [optional] [readonly] 
**ExpanderSlot** | Pointer to **string** | It shows the expander slot inforamtion for the card. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | It shows current firmware version of graphics card. | [optional] [readonly] 
**Mode** | Pointer to **string** | It shows the current mode of graphics card. | [optional] [readonly] 
**NumGpus** | Pointer to **string** | It shows number of controllers under each card. | [optional] 
**OperState** | Pointer to **string** | It shows the current operational state of graphics card. | [optional] [readonly] 
**PciAddress** | Pointer to **string** | It shows the pci address of graphics card. | [optional] [readonly] 
**PciAddressList** | Pointer to **string** | This list contains the pci address of all controllers for corresponding card. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | It shows the pci slot name for grapchics card. | [optional] [readonly] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**GraphicsControllers** | Pointer to [**[]GraphicsControllerRelationship**](graphics.Controller.Relationship.md) | An array of relationships to graphicsController resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewGraphicsCard

`func NewGraphicsCard() *GraphicsCard`

NewGraphicsCard instantiates a new GraphicsCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphicsCardWithDefaults

`func NewGraphicsCardWithDefaults() *GraphicsCard`

NewGraphicsCardWithDefaults instantiates a new GraphicsCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardId

`func (o *GraphicsCard) GetCardId() int64`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *GraphicsCard) GetCardIdOk() (*int64, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *GraphicsCard) SetCardId(v int64)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *GraphicsCard) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetDeviceId

`func (o *GraphicsCard) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GraphicsCard) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GraphicsCard) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GraphicsCard) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetExpanderSlot

`func (o *GraphicsCard) GetExpanderSlot() string`

GetExpanderSlot returns the ExpanderSlot field if non-nil, zero value otherwise.

### GetExpanderSlotOk

`func (o *GraphicsCard) GetExpanderSlotOk() (*string, bool)`

GetExpanderSlotOk returns a tuple with the ExpanderSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanderSlot

`func (o *GraphicsCard) SetExpanderSlot(v string)`

SetExpanderSlot sets ExpanderSlot field to given value.

### HasExpanderSlot

`func (o *GraphicsCard) HasExpanderSlot() bool`

HasExpanderSlot returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *GraphicsCard) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *GraphicsCard) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *GraphicsCard) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *GraphicsCard) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetMode

`func (o *GraphicsCard) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GraphicsCard) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GraphicsCard) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GraphicsCard) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNumGpus

`func (o *GraphicsCard) GetNumGpus() string`

GetNumGpus returns the NumGpus field if non-nil, zero value otherwise.

### GetNumGpusOk

`func (o *GraphicsCard) GetNumGpusOk() (*string, bool)`

GetNumGpusOk returns a tuple with the NumGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumGpus

`func (o *GraphicsCard) SetNumGpus(v string)`

SetNumGpus sets NumGpus field to given value.

### HasNumGpus

`func (o *GraphicsCard) HasNumGpus() bool`

HasNumGpus returns a boolean if a field has been set.

### GetOperState

`func (o *GraphicsCard) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *GraphicsCard) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *GraphicsCard) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *GraphicsCard) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPciAddress

`func (o *GraphicsCard) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *GraphicsCard) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *GraphicsCard) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *GraphicsCard) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.

### GetPciAddressList

`func (o *GraphicsCard) GetPciAddressList() string`

GetPciAddressList returns the PciAddressList field if non-nil, zero value otherwise.

### GetPciAddressListOk

`func (o *GraphicsCard) GetPciAddressListOk() (*string, bool)`

GetPciAddressListOk returns a tuple with the PciAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddressList

`func (o *GraphicsCard) SetPciAddressList(v string)`

SetPciAddressList sets PciAddressList field to given value.

### HasPciAddressList

`func (o *GraphicsCard) HasPciAddressList() bool`

HasPciAddressList returns a boolean if a field has been set.

### GetPciSlot

`func (o *GraphicsCard) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *GraphicsCard) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *GraphicsCard) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *GraphicsCard) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetComputeBoard

`func (o *GraphicsCard) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *GraphicsCard) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *GraphicsCard) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *GraphicsCard) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetGraphicsControllers

`func (o *GraphicsCard) GetGraphicsControllers() []GraphicsControllerRelationship`

GetGraphicsControllers returns the GraphicsControllers field if non-nil, zero value otherwise.

### GetGraphicsControllersOk

`func (o *GraphicsCard) GetGraphicsControllersOk() (*[]GraphicsControllerRelationship, bool)`

GetGraphicsControllersOk returns a tuple with the GraphicsControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsControllers

`func (o *GraphicsCard) SetGraphicsControllers(v []GraphicsControllerRelationship)`

SetGraphicsControllers sets GraphicsControllers field to given value.

### HasGraphicsControllers

`func (o *GraphicsCard) HasGraphicsControllers() bool`

HasGraphicsControllers returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *GraphicsCard) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *GraphicsCard) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *GraphicsCard) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *GraphicsCard) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


