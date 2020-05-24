# IamSessionLimitsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **string** | A discriminator value to disambiguate the schema of a HTTP GET response body. | 
**Count** | Pointer to **int32** | The total number of &#39;iam.SessionLimits&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]IamSessionLimits**](iam.SessionLimits.md) | The array of &#39;iam.SessionLimits&#39; resources matching the request. | [optional] 

## Methods

### NewIamSessionLimitsResponse

`func NewIamSessionLimitsResponse(objectType string, ) *IamSessionLimitsResponse`

NewIamSessionLimitsResponse instantiates a new IamSessionLimitsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSessionLimitsResponseWithDefaults

`func NewIamSessionLimitsResponseWithDefaults() *IamSessionLimitsResponse`

NewIamSessionLimitsResponseWithDefaults instantiates a new IamSessionLimitsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *IamSessionLimitsResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamSessionLimitsResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamSessionLimitsResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *IamSessionLimitsResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IamSessionLimitsResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IamSessionLimitsResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IamSessionLimitsResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *IamSessionLimitsResponse) GetResults() []IamSessionLimits`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IamSessionLimitsResponse) GetResultsOk() (*[]IamSessionLimits, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IamSessionLimitsResponse) SetResults(v []IamSessionLimits)`

SetResults sets Results field to given value.

### HasResults

`func (o *IamSessionLimitsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


