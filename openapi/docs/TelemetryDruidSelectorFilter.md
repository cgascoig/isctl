# TelemetryDruidSelectorFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | Pointer to **string** | The name of a dimension. | 
**Value** | Pointer to **string** | The value of a dimension. | 

## Methods

### NewTelemetryDruidSelectorFilter

`func NewTelemetryDruidSelectorFilter(dimension string, value string, ) *TelemetryDruidSelectorFilter`

NewTelemetryDruidSelectorFilter instantiates a new TelemetryDruidSelectorFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidSelectorFilterWithDefaults

`func NewTelemetryDruidSelectorFilterWithDefaults() *TelemetryDruidSelectorFilter`

NewTelemetryDruidSelectorFilterWithDefaults instantiates a new TelemetryDruidSelectorFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *TelemetryDruidSelectorFilter) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidSelectorFilter) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidSelectorFilter) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetValue

`func (o *TelemetryDruidSelectorFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TelemetryDruidSelectorFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TelemetryDruidSelectorFilter) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


