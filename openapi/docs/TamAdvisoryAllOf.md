# TamAdvisoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Brief description of the advisory details. | [optional] 
**Name** | Pointer to **string** | A user defined name for the Intersight Advisory. | [optional] 
**Severity** | Pointer to [**TamSeverity**](tam.Severity.md) |  | [optional] 
**State** | Pointer to **string** | Current state of the advisory. Indicates if the user is interested in getting updates for the advisory. | [optional] [default to "active"]

## Methods

### NewTamAdvisoryAllOf

`func NewTamAdvisoryAllOf() *TamAdvisoryAllOf`

NewTamAdvisoryAllOf instantiates a new TamAdvisoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamAdvisoryAllOfWithDefaults

`func NewTamAdvisoryAllOfWithDefaults() *TamAdvisoryAllOf`

NewTamAdvisoryAllOfWithDefaults instantiates a new TamAdvisoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TamAdvisoryAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TamAdvisoryAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TamAdvisoryAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TamAdvisoryAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *TamAdvisoryAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TamAdvisoryAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TamAdvisoryAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TamAdvisoryAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *TamAdvisoryAllOf) GetSeverity() TamSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *TamAdvisoryAllOf) GetSeverityOk() (*TamSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *TamAdvisoryAllOf) SetSeverity(v TamSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *TamAdvisoryAllOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetState

`func (o *TamAdvisoryAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TamAdvisoryAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TamAdvisoryAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TamAdvisoryAllOf) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


