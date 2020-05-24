# OsWindowsParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Edition** | Pointer to **string** | Lists all the editions supported for Windows Server installation. | [optional] [default to "Standard"]

## Methods

### NewOsWindowsParameters

`func NewOsWindowsParameters() *OsWindowsParameters`

NewOsWindowsParameters instantiates a new OsWindowsParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsWindowsParametersWithDefaults

`func NewOsWindowsParametersWithDefaults() *OsWindowsParameters`

NewOsWindowsParametersWithDefaults instantiates a new OsWindowsParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdition

`func (o *OsWindowsParameters) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *OsWindowsParameters) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *OsWindowsParameters) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *OsWindowsParameters) HasEdition() bool`

HasEdition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


