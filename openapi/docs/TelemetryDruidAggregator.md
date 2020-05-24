# TelemetryDruidAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The aggregator type. | 
**Name** | Pointer to **string** | Output name for the minimum/maximum timestamp value. | 
**FieldName** | Pointer to **string** | Name of the metric column. | 
**MaxStringBytes** | Pointer to **int32** | null | [default to 1024]

## Methods

### NewTelemetryDruidAggregator

`func NewTelemetryDruidAggregator(type_ string, name string, fieldName string, maxStringBytes int32, ) *TelemetryDruidAggregator`

NewTelemetryDruidAggregator instantiates a new TelemetryDruidAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidAggregatorWithDefaults

`func NewTelemetryDruidAggregatorWithDefaults() *TelemetryDruidAggregator`

NewTelemetryDruidAggregatorWithDefaults instantiates a new TelemetryDruidAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidAggregator) SetName(v string)`

SetName sets Name field to given value.


### GetFieldName

`func (o *TelemetryDruidAggregator) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *TelemetryDruidAggregator) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *TelemetryDruidAggregator) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetMaxStringBytes

`func (o *TelemetryDruidAggregator) GetMaxStringBytes() int32`

GetMaxStringBytes returns the MaxStringBytes field if non-nil, zero value otherwise.

### GetMaxStringBytesOk

`func (o *TelemetryDruidAggregator) GetMaxStringBytesOk() (*int32, bool)`

GetMaxStringBytesOk returns a tuple with the MaxStringBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStringBytes

`func (o *TelemetryDruidAggregator) SetMaxStringBytes(v int32)`

SetMaxStringBytes sets MaxStringBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


