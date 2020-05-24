# StorageFlexFlashControllerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerState** | Pointer to **string** |  | [optional] [readonly] 
**FfControllerId** | Pointer to **string** |  | [optional] [readonly] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**FlexFlashControllerProps** | Pointer to [**[]StorageFlexFlashControllerPropsRelationship**](storage.FlexFlashControllerProps.Relationship.md) | An array of relationships to storageFlexFlashControllerProps resources. | [optional] [readonly] 
**FlexFlashPhysicalDrives** | Pointer to [**[]StorageFlexFlashPhysicalDriveRelationship**](storage.FlexFlashPhysicalDrive.Relationship.md) | An array of relationships to storageFlexFlashPhysicalDrive resources. | [optional] [readonly] 
**FlexFlashVirtualDrives** | Pointer to [**[]StorageFlexFlashVirtualDriveRelationship**](storage.FlexFlashVirtualDrive.Relationship.md) | An array of relationships to storageFlexFlashVirtualDrive resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewStorageFlexFlashControllerAllOf

`func NewStorageFlexFlashControllerAllOf() *StorageFlexFlashControllerAllOf`

NewStorageFlexFlashControllerAllOf instantiates a new StorageFlexFlashControllerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexFlashControllerAllOfWithDefaults

`func NewStorageFlexFlashControllerAllOfWithDefaults() *StorageFlexFlashControllerAllOf`

NewStorageFlexFlashControllerAllOfWithDefaults instantiates a new StorageFlexFlashControllerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerState

`func (o *StorageFlexFlashControllerAllOf) GetControllerState() string`

GetControllerState returns the ControllerState field if non-nil, zero value otherwise.

### GetControllerStateOk

`func (o *StorageFlexFlashControllerAllOf) GetControllerStateOk() (*string, bool)`

GetControllerStateOk returns a tuple with the ControllerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerState

`func (o *StorageFlexFlashControllerAllOf) SetControllerState(v string)`

SetControllerState sets ControllerState field to given value.

### HasControllerState

`func (o *StorageFlexFlashControllerAllOf) HasControllerState() bool`

HasControllerState returns a boolean if a field has been set.

### GetFfControllerId

`func (o *StorageFlexFlashControllerAllOf) GetFfControllerId() string`

GetFfControllerId returns the FfControllerId field if non-nil, zero value otherwise.

### GetFfControllerIdOk

`func (o *StorageFlexFlashControllerAllOf) GetFfControllerIdOk() (*string, bool)`

GetFfControllerIdOk returns a tuple with the FfControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFfControllerId

`func (o *StorageFlexFlashControllerAllOf) SetFfControllerId(v string)`

SetFfControllerId sets FfControllerId field to given value.

### HasFfControllerId

`func (o *StorageFlexFlashControllerAllOf) HasFfControllerId() bool`

HasFfControllerId returns a boolean if a field has been set.

### GetComputeBoard

`func (o *StorageFlexFlashControllerAllOf) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *StorageFlexFlashControllerAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *StorageFlexFlashControllerAllOf) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *StorageFlexFlashControllerAllOf) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetFlexFlashControllerProps

`func (o *StorageFlexFlashControllerAllOf) GetFlexFlashControllerProps() []StorageFlexFlashControllerPropsRelationship`

GetFlexFlashControllerProps returns the FlexFlashControllerProps field if non-nil, zero value otherwise.

### GetFlexFlashControllerPropsOk

`func (o *StorageFlexFlashControllerAllOf) GetFlexFlashControllerPropsOk() (*[]StorageFlexFlashControllerPropsRelationship, bool)`

GetFlexFlashControllerPropsOk returns a tuple with the FlexFlashControllerProps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlashControllerProps

`func (o *StorageFlexFlashControllerAllOf) SetFlexFlashControllerProps(v []StorageFlexFlashControllerPropsRelationship)`

SetFlexFlashControllerProps sets FlexFlashControllerProps field to given value.

### HasFlexFlashControllerProps

`func (o *StorageFlexFlashControllerAllOf) HasFlexFlashControllerProps() bool`

HasFlexFlashControllerProps returns a boolean if a field has been set.

### GetFlexFlashPhysicalDrives

`func (o *StorageFlexFlashControllerAllOf) GetFlexFlashPhysicalDrives() []StorageFlexFlashPhysicalDriveRelationship`

GetFlexFlashPhysicalDrives returns the FlexFlashPhysicalDrives field if non-nil, zero value otherwise.

### GetFlexFlashPhysicalDrivesOk

`func (o *StorageFlexFlashControllerAllOf) GetFlexFlashPhysicalDrivesOk() (*[]StorageFlexFlashPhysicalDriveRelationship, bool)`

GetFlexFlashPhysicalDrivesOk returns a tuple with the FlexFlashPhysicalDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlashPhysicalDrives

`func (o *StorageFlexFlashControllerAllOf) SetFlexFlashPhysicalDrives(v []StorageFlexFlashPhysicalDriveRelationship)`

SetFlexFlashPhysicalDrives sets FlexFlashPhysicalDrives field to given value.

### HasFlexFlashPhysicalDrives

`func (o *StorageFlexFlashControllerAllOf) HasFlexFlashPhysicalDrives() bool`

HasFlexFlashPhysicalDrives returns a boolean if a field has been set.

### GetFlexFlashVirtualDrives

`func (o *StorageFlexFlashControllerAllOf) GetFlexFlashVirtualDrives() []StorageFlexFlashVirtualDriveRelationship`

GetFlexFlashVirtualDrives returns the FlexFlashVirtualDrives field if non-nil, zero value otherwise.

### GetFlexFlashVirtualDrivesOk

`func (o *StorageFlexFlashControllerAllOf) GetFlexFlashVirtualDrivesOk() (*[]StorageFlexFlashVirtualDriveRelationship, bool)`

GetFlexFlashVirtualDrivesOk returns a tuple with the FlexFlashVirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlashVirtualDrives

`func (o *StorageFlexFlashControllerAllOf) SetFlexFlashVirtualDrives(v []StorageFlexFlashVirtualDriveRelationship)`

SetFlexFlashVirtualDrives sets FlexFlashVirtualDrives field to given value.

### HasFlexFlashVirtualDrives

`func (o *StorageFlexFlashControllerAllOf) HasFlexFlashVirtualDrives() bool`

HasFlexFlashVirtualDrives returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexFlashControllerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexFlashControllerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexFlashControllerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexFlashControllerAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


