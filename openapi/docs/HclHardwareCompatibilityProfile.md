# HclHardwareCompatibilityProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriverIsoUrl** | Pointer to **string** | Url for the ISO with the drivers supported for the server. | [optional] 
**ErrorCode** | Pointer to **string** | Error code indicating the compatibility status. | [optional] [readonly] [default to "Success"]
**Id** | Pointer to **string** | Identifier of the hardware compatibility profile. | [optional] 
**OsVendor** | Pointer to **string** | Vendor of the Operating System running on the server. | [optional] 
**OsVersion** | Pointer to **string** | Version of the Operating System running on the server. | [optional] 
**ProcessorModel** | Pointer to **string** | Model of the processor present in the server. | [optional] 
**Products** | Pointer to [**[]HclProduct**](hcl.Product.md) |  | [optional] 
**ServerModel** | Pointer to **string** | Model of the server as returned by UCSM/CIMC XML API. | [optional] 
**ServerRevision** | Pointer to **string** | Revision of the server model. | [optional] 
**UcsVersion** | Pointer to **string** | Version of the UCS software. | [optional] 
**VersionType** | Pointer to **string** | Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release. | [optional] [default to "UCSM"]

## Methods

### NewHclHardwareCompatibilityProfile

`func NewHclHardwareCompatibilityProfile() *HclHardwareCompatibilityProfile`

NewHclHardwareCompatibilityProfile instantiates a new HclHardwareCompatibilityProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclHardwareCompatibilityProfileWithDefaults

`func NewHclHardwareCompatibilityProfileWithDefaults() *HclHardwareCompatibilityProfile`

NewHclHardwareCompatibilityProfileWithDefaults instantiates a new HclHardwareCompatibilityProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriverIsoUrl

`func (o *HclHardwareCompatibilityProfile) GetDriverIsoUrl() string`

GetDriverIsoUrl returns the DriverIsoUrl field if non-nil, zero value otherwise.

### GetDriverIsoUrlOk

`func (o *HclHardwareCompatibilityProfile) GetDriverIsoUrlOk() (*string, bool)`

GetDriverIsoUrlOk returns a tuple with the DriverIsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverIsoUrl

`func (o *HclHardwareCompatibilityProfile) SetDriverIsoUrl(v string)`

SetDriverIsoUrl sets DriverIsoUrl field to given value.

### HasDriverIsoUrl

`func (o *HclHardwareCompatibilityProfile) HasDriverIsoUrl() bool`

HasDriverIsoUrl returns a boolean if a field has been set.

### GetErrorCode

`func (o *HclHardwareCompatibilityProfile) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *HclHardwareCompatibilityProfile) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *HclHardwareCompatibilityProfile) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *HclHardwareCompatibilityProfile) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetId

`func (o *HclHardwareCompatibilityProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HclHardwareCompatibilityProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HclHardwareCompatibilityProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HclHardwareCompatibilityProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOsVendor

`func (o *HclHardwareCompatibilityProfile) GetOsVendor() string`

GetOsVendor returns the OsVendor field if non-nil, zero value otherwise.

### GetOsVendorOk

`func (o *HclHardwareCompatibilityProfile) GetOsVendorOk() (*string, bool)`

GetOsVendorOk returns a tuple with the OsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVendor

`func (o *HclHardwareCompatibilityProfile) SetOsVendor(v string)`

SetOsVendor sets OsVendor field to given value.

### HasOsVendor

`func (o *HclHardwareCompatibilityProfile) HasOsVendor() bool`

HasOsVendor returns a boolean if a field has been set.

### GetOsVersion

`func (o *HclHardwareCompatibilityProfile) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *HclHardwareCompatibilityProfile) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *HclHardwareCompatibilityProfile) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *HclHardwareCompatibilityProfile) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetProcessorModel

`func (o *HclHardwareCompatibilityProfile) GetProcessorModel() string`

GetProcessorModel returns the ProcessorModel field if non-nil, zero value otherwise.

### GetProcessorModelOk

`func (o *HclHardwareCompatibilityProfile) GetProcessorModelOk() (*string, bool)`

GetProcessorModelOk returns a tuple with the ProcessorModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorModel

`func (o *HclHardwareCompatibilityProfile) SetProcessorModel(v string)`

SetProcessorModel sets ProcessorModel field to given value.

### HasProcessorModel

`func (o *HclHardwareCompatibilityProfile) HasProcessorModel() bool`

HasProcessorModel returns a boolean if a field has been set.

### GetProducts

`func (o *HclHardwareCompatibilityProfile) GetProducts() []HclProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *HclHardwareCompatibilityProfile) GetProductsOk() (*[]HclProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *HclHardwareCompatibilityProfile) SetProducts(v []HclProduct)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *HclHardwareCompatibilityProfile) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetServerModel

`func (o *HclHardwareCompatibilityProfile) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *HclHardwareCompatibilityProfile) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *HclHardwareCompatibilityProfile) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *HclHardwareCompatibilityProfile) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.

### GetServerRevision

`func (o *HclHardwareCompatibilityProfile) GetServerRevision() string`

GetServerRevision returns the ServerRevision field if non-nil, zero value otherwise.

### GetServerRevisionOk

`func (o *HclHardwareCompatibilityProfile) GetServerRevisionOk() (*string, bool)`

GetServerRevisionOk returns a tuple with the ServerRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRevision

`func (o *HclHardwareCompatibilityProfile) SetServerRevision(v string)`

SetServerRevision sets ServerRevision field to given value.

### HasServerRevision

`func (o *HclHardwareCompatibilityProfile) HasServerRevision() bool`

HasServerRevision returns a boolean if a field has been set.

### GetUcsVersion

`func (o *HclHardwareCompatibilityProfile) GetUcsVersion() string`

GetUcsVersion returns the UcsVersion field if non-nil, zero value otherwise.

### GetUcsVersionOk

`func (o *HclHardwareCompatibilityProfile) GetUcsVersionOk() (*string, bool)`

GetUcsVersionOk returns a tuple with the UcsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsVersion

`func (o *HclHardwareCompatibilityProfile) SetUcsVersion(v string)`

SetUcsVersion sets UcsVersion field to given value.

### HasUcsVersion

`func (o *HclHardwareCompatibilityProfile) HasUcsVersion() bool`

HasUcsVersion returns a boolean if a field has been set.

### GetVersionType

`func (o *HclHardwareCompatibilityProfile) GetVersionType() string`

GetVersionType returns the VersionType field if non-nil, zero value otherwise.

### GetVersionTypeOk

`func (o *HclHardwareCompatibilityProfile) GetVersionTypeOk() (*string, bool)`

GetVersionTypeOk returns a tuple with the VersionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionType

`func (o *HclHardwareCompatibilityProfile) SetVersionType(v string)`

SetVersionType sets VersionType field to given value.

### HasVersionType

`func (o *HclHardwareCompatibilityProfile) HasVersionType() bool`

HasVersionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


