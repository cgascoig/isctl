# AssetDeviceConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **string** | A discriminator value to disambiguate the schema of a HTTP GET response body. | 
**Count** | Pointer to **int32** | The total number of &#39;asset.DeviceConfiguration&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]AssetDeviceConfiguration**](asset.DeviceConfiguration.md) | The array of &#39;asset.DeviceConfiguration&#39; resources matching the request. | [optional] 

## Methods

### NewAssetDeviceConfigurationResponse

`func NewAssetDeviceConfigurationResponse(objectType string, ) *AssetDeviceConfigurationResponse`

NewAssetDeviceConfigurationResponse instantiates a new AssetDeviceConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceConfigurationResponseWithDefaults

`func NewAssetDeviceConfigurationResponseWithDefaults() *AssetDeviceConfigurationResponse`

NewAssetDeviceConfigurationResponseWithDefaults instantiates a new AssetDeviceConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *AssetDeviceConfigurationResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeviceConfigurationResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeviceConfigurationResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *AssetDeviceConfigurationResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AssetDeviceConfigurationResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AssetDeviceConfigurationResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AssetDeviceConfigurationResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *AssetDeviceConfigurationResponse) GetResults() []AssetDeviceConfiguration`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AssetDeviceConfigurationResponse) GetResultsOk() (*[]AssetDeviceConfiguration, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AssetDeviceConfigurationResponse) SetResults(v []AssetDeviceConfiguration)`

SetResults sets Results field to given value.

### HasResults

`func (o *AssetDeviceConfigurationResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


