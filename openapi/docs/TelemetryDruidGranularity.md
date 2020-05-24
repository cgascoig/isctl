# TelemetryDruidGranularity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | the type of granularity. | 
**Period** | Pointer to **string** | The period in ISO 8601 format. Examples are P2W, P3M, PT1H30M, PT0.750S. | 
**TimeZone** | Pointer to **string** | An optional value specifying the time zone. Standard [IANA time zones](http://joda-time.sourceforge.net/timezones.html) are supported. | [optional] 

## Methods

### NewTelemetryDruidGranularity

`func NewTelemetryDruidGranularity(type_ string, period string, ) *TelemetryDruidGranularity`

NewTelemetryDruidGranularity instantiates a new TelemetryDruidGranularity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidGranularityWithDefaults

`func NewTelemetryDruidGranularityWithDefaults() *TelemetryDruidGranularity`

NewTelemetryDruidGranularityWithDefaults instantiates a new TelemetryDruidGranularity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidGranularity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidGranularity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidGranularity) SetType(v string)`

SetType sets Type field to given value.


### GetPeriod

`func (o *TelemetryDruidGranularity) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TelemetryDruidGranularity) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TelemetryDruidGranularity) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetTimeZone

`func (o *TelemetryDruidGranularity) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *TelemetryDruidGranularity) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *TelemetryDruidGranularity) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *TelemetryDruidGranularity) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


