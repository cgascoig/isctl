# TelemetryDruidDataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of data source. | 
**DataSources** | Pointer to **[]string** | A list of data sources. | 
**Query** | Pointer to [**TelemetryDruidGroupByRequest**](telemetry.DruidGroupByRequest.md) |  | 

## Methods

### NewTelemetryDruidDataSource

`func NewTelemetryDruidDataSource(type_ string, dataSources []string, query TelemetryDruidGroupByRequest, ) *TelemetryDruidDataSource`

NewTelemetryDruidDataSource instantiates a new TelemetryDruidDataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidDataSourceWithDefaults

`func NewTelemetryDruidDataSourceWithDefaults() *TelemetryDruidDataSource`

NewTelemetryDruidDataSourceWithDefaults instantiates a new TelemetryDruidDataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidDataSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidDataSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidDataSource) SetType(v string)`

SetType sets Type field to given value.


### GetDataSources

`func (o *TelemetryDruidDataSource) GetDataSources() []string`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *TelemetryDruidDataSource) GetDataSourcesOk() (*[]string, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *TelemetryDruidDataSource) SetDataSources(v []string)`

SetDataSources sets DataSources field to given value.


### GetQuery

`func (o *TelemetryDruidDataSource) GetQuery() TelemetryDruidGroupByRequest`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TelemetryDruidDataSource) GetQueryOk() (*TelemetryDruidGroupByRequest, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TelemetryDruidDataSource) SetQuery(v TelemetryDruidGroupByRequest)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


