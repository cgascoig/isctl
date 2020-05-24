# TelemetryDruidStringFirstLastAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Output name for the minimum/maximum timestamp value. | 
**FieldName** | Pointer to **string** | Name of the metric column. | 
**MaxStringBytes** | Pointer to **int32** | null | [default to 1024]

## Methods

### NewTelemetryDruidStringFirstLastAggregator

`func NewTelemetryDruidStringFirstLastAggregator(name string, fieldName string, maxStringBytes int32, ) *TelemetryDruidStringFirstLastAggregator`

NewTelemetryDruidStringFirstLastAggregator instantiates a new TelemetryDruidStringFirstLastAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidStringFirstLastAggregatorWithDefaults

`func NewTelemetryDruidStringFirstLastAggregatorWithDefaults() *TelemetryDruidStringFirstLastAggregator`

NewTelemetryDruidStringFirstLastAggregatorWithDefaults instantiates a new TelemetryDruidStringFirstLastAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TelemetryDruidStringFirstLastAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidStringFirstLastAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidStringFirstLastAggregator) SetName(v string)`

SetName sets Name field to given value.


### GetFieldName

`func (o *TelemetryDruidStringFirstLastAggregator) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *TelemetryDruidStringFirstLastAggregator) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *TelemetryDruidStringFirstLastAggregator) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetMaxStringBytes

`func (o *TelemetryDruidStringFirstLastAggregator) GetMaxStringBytes() int32`

GetMaxStringBytes returns the MaxStringBytes field if non-nil, zero value otherwise.

### GetMaxStringBytesOk

`func (o *TelemetryDruidStringFirstLastAggregator) GetMaxStringBytesOk() (*int32, bool)`

GetMaxStringBytesOk returns a tuple with the MaxStringBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStringBytes

`func (o *TelemetryDruidStringFirstLastAggregator) SetMaxStringBytes(v int32)`

SetMaxStringBytes sets MaxStringBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


