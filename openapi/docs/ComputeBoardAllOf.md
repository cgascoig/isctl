# ComputeBoardAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoardId** | Pointer to **int64** |  | [optional] [readonly] 
**CpuTypeController** | Pointer to **string** |  | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | Current power state of the mother board of the server. | [optional] [readonly] 
**Presence** | Pointer to **string** | Identifies the presence of the mother board of the server. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**EquipmentTpms** | Pointer to [**[]EquipmentTpmRelationship**](equipment.Tpm.Relationship.md) | An array of relationships to equipmentTpm resources. | [optional] [readonly] 
**GraphicsCards** | Pointer to [**[]GraphicsCardRelationship**](graphics.Card.Relationship.md) | An array of relationships to graphicsCard resources. | [optional] [readonly] 
**MemoryArrays** | Pointer to [**[]MemoryArrayRelationship**](memory.Array.Relationship.md) | An array of relationships to memoryArray resources. | [optional] [readonly] 
**PciCoprocessorCards** | Pointer to [**[]PciCoprocessorCardRelationship**](pci.CoprocessorCard.Relationship.md) | An array of relationships to pciCoprocessorCard resources. | [optional] [readonly] 
**PciSwitch** | Pointer to [**[]PciSwitchRelationship**](pci.Switch.Relationship.md) | An array of relationships to pciSwitch resources. | [optional] [readonly] 
**PersistentMemoryConfiguration** | Pointer to [**MemoryPersistentMemoryConfigurationRelationship**](memory.PersistentMemoryConfiguration.Relationship.md) |  | [optional] 
**Processors** | Pointer to [**[]ProcessorUnitRelationship**](processor.Unit.Relationship.md) | An array of relationships to processorUnit resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**SecurityUnits** | Pointer to [**[]SecurityUnitRelationship**](security.Unit.Relationship.md) | An array of relationships to securityUnit resources. | [optional] [readonly] 
**StorageControllers** | Pointer to [**[]StorageControllerRelationship**](storage.Controller.Relationship.md) | An array of relationships to storageController resources. | [optional] [readonly] 
**StorageFlexFlashControllers** | Pointer to [**[]StorageFlexFlashControllerRelationship**](storage.FlexFlashController.Relationship.md) | An array of relationships to storageFlexFlashController resources. | [optional] [readonly] 
**StorageFlexUtilControllers** | Pointer to [**[]StorageFlexUtilControllerRelationship**](storage.FlexUtilController.Relationship.md) | An array of relationships to storageFlexUtilController resources. | [optional] [readonly] 

## Methods

### NewComputeBoardAllOf

`func NewComputeBoardAllOf() *ComputeBoardAllOf`

NewComputeBoardAllOf instantiates a new ComputeBoardAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeBoardAllOfWithDefaults

`func NewComputeBoardAllOfWithDefaults() *ComputeBoardAllOf`

NewComputeBoardAllOfWithDefaults instantiates a new ComputeBoardAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoardId

`func (o *ComputeBoardAllOf) GetBoardId() int64`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *ComputeBoardAllOf) GetBoardIdOk() (*int64, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *ComputeBoardAllOf) SetBoardId(v int64)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *ComputeBoardAllOf) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetCpuTypeController

`func (o *ComputeBoardAllOf) GetCpuTypeController() string`

GetCpuTypeController returns the CpuTypeController field if non-nil, zero value otherwise.

### GetCpuTypeControllerOk

`func (o *ComputeBoardAllOf) GetCpuTypeControllerOk() (*string, bool)`

GetCpuTypeControllerOk returns a tuple with the CpuTypeController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTypeController

`func (o *ComputeBoardAllOf) SetCpuTypeController(v string)`

SetCpuTypeController sets CpuTypeController field to given value.

### HasCpuTypeController

`func (o *ComputeBoardAllOf) HasCpuTypeController() bool`

HasCpuTypeController returns a boolean if a field has been set.

### GetOperPowerState

`func (o *ComputeBoardAllOf) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *ComputeBoardAllOf) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *ComputeBoardAllOf) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *ComputeBoardAllOf) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetPresence

`func (o *ComputeBoardAllOf) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *ComputeBoardAllOf) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *ComputeBoardAllOf) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *ComputeBoardAllOf) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetComputeBlade

`func (o *ComputeBoardAllOf) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *ComputeBoardAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *ComputeBoardAllOf) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *ComputeBoardAllOf) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *ComputeBoardAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *ComputeBoardAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *ComputeBoardAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *ComputeBoardAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetEquipmentTpms

`func (o *ComputeBoardAllOf) GetEquipmentTpms() []EquipmentTpmRelationship`

GetEquipmentTpms returns the EquipmentTpms field if non-nil, zero value otherwise.

### GetEquipmentTpmsOk

`func (o *ComputeBoardAllOf) GetEquipmentTpmsOk() (*[]EquipmentTpmRelationship, bool)`

GetEquipmentTpmsOk returns a tuple with the EquipmentTpms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentTpms

`func (o *ComputeBoardAllOf) SetEquipmentTpms(v []EquipmentTpmRelationship)`

SetEquipmentTpms sets EquipmentTpms field to given value.

### HasEquipmentTpms

`func (o *ComputeBoardAllOf) HasEquipmentTpms() bool`

HasEquipmentTpms returns a boolean if a field has been set.

### GetGraphicsCards

`func (o *ComputeBoardAllOf) GetGraphicsCards() []GraphicsCardRelationship`

GetGraphicsCards returns the GraphicsCards field if non-nil, zero value otherwise.

### GetGraphicsCardsOk

`func (o *ComputeBoardAllOf) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool)`

GetGraphicsCardsOk returns a tuple with the GraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCards

`func (o *ComputeBoardAllOf) SetGraphicsCards(v []GraphicsCardRelationship)`

SetGraphicsCards sets GraphicsCards field to given value.

### HasGraphicsCards

`func (o *ComputeBoardAllOf) HasGraphicsCards() bool`

HasGraphicsCards returns a boolean if a field has been set.

### GetMemoryArrays

`func (o *ComputeBoardAllOf) GetMemoryArrays() []MemoryArrayRelationship`

GetMemoryArrays returns the MemoryArrays field if non-nil, zero value otherwise.

### GetMemoryArraysOk

`func (o *ComputeBoardAllOf) GetMemoryArraysOk() (*[]MemoryArrayRelationship, bool)`

GetMemoryArraysOk returns a tuple with the MemoryArrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryArrays

`func (o *ComputeBoardAllOf) SetMemoryArrays(v []MemoryArrayRelationship)`

SetMemoryArrays sets MemoryArrays field to given value.

### HasMemoryArrays

`func (o *ComputeBoardAllOf) HasMemoryArrays() bool`

HasMemoryArrays returns a boolean if a field has been set.

### GetPciCoprocessorCards

`func (o *ComputeBoardAllOf) GetPciCoprocessorCards() []PciCoprocessorCardRelationship`

GetPciCoprocessorCards returns the PciCoprocessorCards field if non-nil, zero value otherwise.

### GetPciCoprocessorCardsOk

`func (o *ComputeBoardAllOf) GetPciCoprocessorCardsOk() (*[]PciCoprocessorCardRelationship, bool)`

GetPciCoprocessorCardsOk returns a tuple with the PciCoprocessorCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciCoprocessorCards

`func (o *ComputeBoardAllOf) SetPciCoprocessorCards(v []PciCoprocessorCardRelationship)`

SetPciCoprocessorCards sets PciCoprocessorCards field to given value.

### HasPciCoprocessorCards

`func (o *ComputeBoardAllOf) HasPciCoprocessorCards() bool`

HasPciCoprocessorCards returns a boolean if a field has been set.

### GetPciSwitch

`func (o *ComputeBoardAllOf) GetPciSwitch() []PciSwitchRelationship`

GetPciSwitch returns the PciSwitch field if non-nil, zero value otherwise.

### GetPciSwitchOk

`func (o *ComputeBoardAllOf) GetPciSwitchOk() (*[]PciSwitchRelationship, bool)`

GetPciSwitchOk returns a tuple with the PciSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSwitch

`func (o *ComputeBoardAllOf) SetPciSwitch(v []PciSwitchRelationship)`

SetPciSwitch sets PciSwitch field to given value.

### HasPciSwitch

`func (o *ComputeBoardAllOf) HasPciSwitch() bool`

HasPciSwitch returns a boolean if a field has been set.

### GetPersistentMemoryConfiguration

`func (o *ComputeBoardAllOf) GetPersistentMemoryConfiguration() MemoryPersistentMemoryConfigurationRelationship`

GetPersistentMemoryConfiguration returns the PersistentMemoryConfiguration field if non-nil, zero value otherwise.

### GetPersistentMemoryConfigurationOk

`func (o *ComputeBoardAllOf) GetPersistentMemoryConfigurationOk() (*MemoryPersistentMemoryConfigurationRelationship, bool)`

GetPersistentMemoryConfigurationOk returns a tuple with the PersistentMemoryConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryConfiguration

`func (o *ComputeBoardAllOf) SetPersistentMemoryConfiguration(v MemoryPersistentMemoryConfigurationRelationship)`

SetPersistentMemoryConfiguration sets PersistentMemoryConfiguration field to given value.

### HasPersistentMemoryConfiguration

`func (o *ComputeBoardAllOf) HasPersistentMemoryConfiguration() bool`

HasPersistentMemoryConfiguration returns a boolean if a field has been set.

### GetProcessors

`func (o *ComputeBoardAllOf) GetProcessors() []ProcessorUnitRelationship`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *ComputeBoardAllOf) GetProcessorsOk() (*[]ProcessorUnitRelationship, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *ComputeBoardAllOf) SetProcessors(v []ProcessorUnitRelationship)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *ComputeBoardAllOf) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ComputeBoardAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ComputeBoardAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ComputeBoardAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ComputeBoardAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSecurityUnits

`func (o *ComputeBoardAllOf) GetSecurityUnits() []SecurityUnitRelationship`

GetSecurityUnits returns the SecurityUnits field if non-nil, zero value otherwise.

### GetSecurityUnitsOk

`func (o *ComputeBoardAllOf) GetSecurityUnitsOk() (*[]SecurityUnitRelationship, bool)`

GetSecurityUnitsOk returns a tuple with the SecurityUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityUnits

`func (o *ComputeBoardAllOf) SetSecurityUnits(v []SecurityUnitRelationship)`

SetSecurityUnits sets SecurityUnits field to given value.

### HasSecurityUnits

`func (o *ComputeBoardAllOf) HasSecurityUnits() bool`

HasSecurityUnits returns a boolean if a field has been set.

### GetStorageControllers

`func (o *ComputeBoardAllOf) GetStorageControllers() []StorageControllerRelationship`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ComputeBoardAllOf) GetStorageControllersOk() (*[]StorageControllerRelationship, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ComputeBoardAllOf) SetStorageControllers(v []StorageControllerRelationship)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ComputeBoardAllOf) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetStorageFlexFlashControllers

`func (o *ComputeBoardAllOf) GetStorageFlexFlashControllers() []StorageFlexFlashControllerRelationship`

GetStorageFlexFlashControllers returns the StorageFlexFlashControllers field if non-nil, zero value otherwise.

### GetStorageFlexFlashControllersOk

`func (o *ComputeBoardAllOf) GetStorageFlexFlashControllersOk() (*[]StorageFlexFlashControllerRelationship, bool)`

GetStorageFlexFlashControllersOk returns a tuple with the StorageFlexFlashControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexFlashControllers

`func (o *ComputeBoardAllOf) SetStorageFlexFlashControllers(v []StorageFlexFlashControllerRelationship)`

SetStorageFlexFlashControllers sets StorageFlexFlashControllers field to given value.

### HasStorageFlexFlashControllers

`func (o *ComputeBoardAllOf) HasStorageFlexFlashControllers() bool`

HasStorageFlexFlashControllers returns a boolean if a field has been set.

### GetStorageFlexUtilControllers

`func (o *ComputeBoardAllOf) GetStorageFlexUtilControllers() []StorageFlexUtilControllerRelationship`

GetStorageFlexUtilControllers returns the StorageFlexUtilControllers field if non-nil, zero value otherwise.

### GetStorageFlexUtilControllersOk

`func (o *ComputeBoardAllOf) GetStorageFlexUtilControllersOk() (*[]StorageFlexUtilControllerRelationship, bool)`

GetStorageFlexUtilControllersOk returns a tuple with the StorageFlexUtilControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexUtilControllers

`func (o *ComputeBoardAllOf) SetStorageFlexUtilControllers(v []StorageFlexUtilControllerRelationship)`

SetStorageFlexUtilControllers sets StorageFlexUtilControllers field to given value.

### HasStorageFlexUtilControllers

`func (o *ComputeBoardAllOf) HasStorageFlexUtilControllers() bool`

HasStorageFlexUtilControllers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


