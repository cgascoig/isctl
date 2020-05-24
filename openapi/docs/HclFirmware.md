# HclFirmware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriverName** | Pointer to **string** | Protocol for which the driver is provided. E.g.  enic, fnic, lsi_mr3. | [optional] 
**DriverVersion** | Pointer to **string** | Version of the Driver supported. | [optional] 
**ErrorCode** | Pointer to **string** | Error code for the support status. | [optional] [readonly] [default to "Success"]
**FirmwareVersion** | Pointer to **string** | Firmware version of the product/adapter supported. | [optional] 
**Id** | Pointer to **string** | Identifier of the firmware. | [optional] 
**LatestDriver** | Pointer to **bool** | True if the driver is latest recommended driver. | [optional] [readonly] 
**LatestFirmware** | Pointer to **bool** | True if the firmware is latest recommended firmware. | [optional] [readonly] 

## Methods

### NewHclFirmware

`func NewHclFirmware() *HclFirmware`

NewHclFirmware instantiates a new HclFirmware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclFirmwareWithDefaults

`func NewHclFirmwareWithDefaults() *HclFirmware`

NewHclFirmwareWithDefaults instantiates a new HclFirmware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriverName

`func (o *HclFirmware) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *HclFirmware) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *HclFirmware) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *HclFirmware) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### GetDriverVersion

`func (o *HclFirmware) GetDriverVersion() string`

GetDriverVersion returns the DriverVersion field if non-nil, zero value otherwise.

### GetDriverVersionOk

`func (o *HclFirmware) GetDriverVersionOk() (*string, bool)`

GetDriverVersionOk returns a tuple with the DriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverVersion

`func (o *HclFirmware) SetDriverVersion(v string)`

SetDriverVersion sets DriverVersion field to given value.

### HasDriverVersion

`func (o *HclFirmware) HasDriverVersion() bool`

HasDriverVersion returns a boolean if a field has been set.

### GetErrorCode

`func (o *HclFirmware) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *HclFirmware) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *HclFirmware) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *HclFirmware) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *HclFirmware) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *HclFirmware) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *HclFirmware) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *HclFirmware) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetId

`func (o *HclFirmware) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HclFirmware) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HclFirmware) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HclFirmware) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatestDriver

`func (o *HclFirmware) GetLatestDriver() bool`

GetLatestDriver returns the LatestDriver field if non-nil, zero value otherwise.

### GetLatestDriverOk

`func (o *HclFirmware) GetLatestDriverOk() (*bool, bool)`

GetLatestDriverOk returns a tuple with the LatestDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDriver

`func (o *HclFirmware) SetLatestDriver(v bool)`

SetLatestDriver sets LatestDriver field to given value.

### HasLatestDriver

`func (o *HclFirmware) HasLatestDriver() bool`

HasLatestDriver returns a boolean if a field has been set.

### GetLatestFirmware

`func (o *HclFirmware) GetLatestFirmware() bool`

GetLatestFirmware returns the LatestFirmware field if non-nil, zero value otherwise.

### GetLatestFirmwareOk

`func (o *HclFirmware) GetLatestFirmwareOk() (*bool, bool)`

GetLatestFirmwareOk returns a tuple with the LatestFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestFirmware

`func (o *HclFirmware) SetLatestFirmware(v bool)`

SetLatestFirmware sets LatestFirmware field to given value.

### HasLatestFirmware

`func (o *HclFirmware) HasLatestFirmware() bool`

HasLatestFirmware returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


