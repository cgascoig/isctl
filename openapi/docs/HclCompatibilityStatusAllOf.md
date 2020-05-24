# HclCompatibilityStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileList** | Pointer to [**[]HclHardwareCompatibilityProfile**](hcl.HardwareCompatibilityProfile.md) |  | [optional] 
**RequestType** | Pointer to **string** | Type of the request to be served. | [optional] [default to "FillSupportedVersions"]

## Methods

### NewHclCompatibilityStatusAllOf

`func NewHclCompatibilityStatusAllOf() *HclCompatibilityStatusAllOf`

NewHclCompatibilityStatusAllOf instantiates a new HclCompatibilityStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclCompatibilityStatusAllOfWithDefaults

`func NewHclCompatibilityStatusAllOfWithDefaults() *HclCompatibilityStatusAllOf`

NewHclCompatibilityStatusAllOfWithDefaults instantiates a new HclCompatibilityStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileList

`func (o *HclCompatibilityStatusAllOf) GetProfileList() []HclHardwareCompatibilityProfile`

GetProfileList returns the ProfileList field if non-nil, zero value otherwise.

### GetProfileListOk

`func (o *HclCompatibilityStatusAllOf) GetProfileListOk() (*[]HclHardwareCompatibilityProfile, bool)`

GetProfileListOk returns a tuple with the ProfileList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileList

`func (o *HclCompatibilityStatusAllOf) SetProfileList(v []HclHardwareCompatibilityProfile)`

SetProfileList sets ProfileList field to given value.

### HasProfileList

`func (o *HclCompatibilityStatusAllOf) HasProfileList() bool`

HasProfileList returns a boolean if a field has been set.

### GetRequestType

`func (o *HclCompatibilityStatusAllOf) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *HclCompatibilityStatusAllOf) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *HclCompatibilityStatusAllOf) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *HclCompatibilityStatusAllOf) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


