# TelemetryDruidPostAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The post-aggregator type. | 
**Name** | Pointer to **string** | Output name for the post-aggregator. | [optional] 
**FieldName** | Pointer to **string** | The name field value of the hyperUnique aggregator. | [optional] 
**Value** | Pointer to **float64** | The numerical value. | [optional] 
**Fields** | Pointer to **string** | the fields that are used to compute the greatest or least value. | [optional] 

## Methods

### NewTelemetryDruidPostAggregator

`func NewTelemetryDruidPostAggregator(type_ string, ) *TelemetryDruidPostAggregator`

NewTelemetryDruidPostAggregator instantiates a new TelemetryDruidPostAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidPostAggregatorWithDefaults

`func NewTelemetryDruidPostAggregatorWithDefaults() *TelemetryDruidPostAggregator`

NewTelemetryDruidPostAggregatorWithDefaults instantiates a new TelemetryDruidPostAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidPostAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidPostAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidPostAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidPostAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidPostAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidPostAggregator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryDruidPostAggregator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFieldName

`func (o *TelemetryDruidPostAggregator) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *TelemetryDruidPostAggregator) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *TelemetryDruidPostAggregator) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *TelemetryDruidPostAggregator) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetValue

`func (o *TelemetryDruidPostAggregator) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TelemetryDruidPostAggregator) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TelemetryDruidPostAggregator) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *TelemetryDruidPostAggregator) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetFields

`func (o *TelemetryDruidPostAggregator) GetFields() string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TelemetryDruidPostAggregator) GetFieldsOk() (*string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TelemetryDruidPostAggregator) SetFields(v string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TelemetryDruidPostAggregator) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


