# HclFirmwareAllOf

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

### NewHclFirmwareAllOf

`func NewHclFirmwareAllOf() *HclFirmwareAllOf`

NewHclFirmwareAllOf instantiates a new HclFirmwareAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclFirmwareAllOfWithDefaults

`func NewHclFirmwareAllOfWithDefaults() *HclFirmwareAllOf`

NewHclFirmwareAllOfWithDefaults instantiates a new HclFirmwareAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriverName

`func (o *HclFirmwareAllOf) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *HclFirmwareAllOf) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *HclFirmwareAllOf) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *HclFirmwareAllOf) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### GetDriverVersion

`func (o *HclFirmwareAllOf) GetDriverVersion() string`

GetDriverVersion returns the DriverVersion field if non-nil, zero value otherwise.

### GetDriverVersionOk

`func (o *HclFirmwareAllOf) GetDriverVersionOk() (*string, bool)`

GetDriverVersionOk returns a tuple with the DriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverVersion

`func (o *HclFirmwareAllOf) SetDriverVersion(v string)`

SetDriverVersion sets DriverVersion field to given value.

### HasDriverVersion

`func (o *HclFirmwareAllOf) HasDriverVersion() bool`

HasDriverVersion returns a boolean if a field has been set.

### GetErrorCode

`func (o *HclFirmwareAllOf) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *HclFirmwareAllOf) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *HclFirmwareAllOf) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *HclFirmwareAllOf) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *HclFirmwareAllOf) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *HclFirmwareAllOf) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *HclFirmwareAllOf) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *HclFirmwareAllOf) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetId

`func (o *HclFirmwareAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HclFirmwareAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HclFirmwareAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HclFirmwareAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatestDriver

`func (o *HclFirmwareAllOf) GetLatestDriver() bool`

GetLatestDriver returns the LatestDriver field if non-nil, zero value otherwise.

### GetLatestDriverOk

`func (o *HclFirmwareAllOf) GetLatestDriverOk() (*bool, bool)`

GetLatestDriverOk returns a tuple with the LatestDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDriver

`func (o *HclFirmwareAllOf) SetLatestDriver(v bool)`

SetLatestDriver sets LatestDriver field to given value.

### HasLatestDriver

`func (o *HclFirmwareAllOf) HasLatestDriver() bool`

HasLatestDriver returns a boolean if a field has been set.

### GetLatestFirmware

`func (o *HclFirmwareAllOf) GetLatestFirmware() bool`

GetLatestFirmware returns the LatestFirmware field if non-nil, zero value otherwise.

### GetLatestFirmwareOk

`func (o *HclFirmwareAllOf) GetLatestFirmwareOk() (*bool, bool)`

GetLatestFirmwareOk returns a tuple with the LatestFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestFirmware

`func (o *HclFirmwareAllOf) SetLatestFirmware(v bool)`

SetLatestFirmware sets LatestFirmware field to given value.

### HasLatestFirmware

`func (o *HclFirmwareAllOf) HasLatestFirmware() bool`

HasLatestFirmware returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


