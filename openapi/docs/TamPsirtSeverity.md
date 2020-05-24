# TamPsirtSeverity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** | Severity level associated with the security advisory. | [optional] [default to "critical"]

## Methods

### NewTamPsirtSeverity

`func NewTamPsirtSeverity() *TamPsirtSeverity`

NewTamPsirtSeverity instantiates a new TamPsirtSeverity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamPsirtSeverityWithDefaults

`func NewTamPsirtSeverityWithDefaults() *TamPsirtSeverity`

NewTamPsirtSeverityWithDefaults instantiates a new TamPsirtSeverity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *TamPsirtSeverity) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *TamPsirtSeverity) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *TamPsirtSeverity) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *TamPsirtSeverity) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


