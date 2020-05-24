# StoragePureReplicationScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **string** | A discriminator value to disambiguate the schema of a HTTP GET response body. | 
**Count** | Pointer to **int32** | The total number of &#39;storage.PureReplicationSchedule&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]StoragePureReplicationSchedule**](storage.PureReplicationSchedule.md) | The array of &#39;storage.PureReplicationSchedule&#39; resources matching the request. | [optional] 

## Methods

### NewStoragePureReplicationScheduleResponse

`func NewStoragePureReplicationScheduleResponse(objectType string, ) *StoragePureReplicationScheduleResponse`

NewStoragePureReplicationScheduleResponse instantiates a new StoragePureReplicationScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureReplicationScheduleResponseWithDefaults

`func NewStoragePureReplicationScheduleResponseWithDefaults() *StoragePureReplicationScheduleResponse`

NewStoragePureReplicationScheduleResponseWithDefaults instantiates a new StoragePureReplicationScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *StoragePureReplicationScheduleResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureReplicationScheduleResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureReplicationScheduleResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *StoragePureReplicationScheduleResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StoragePureReplicationScheduleResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StoragePureReplicationScheduleResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *StoragePureReplicationScheduleResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *StoragePureReplicationScheduleResponse) GetResults() []StoragePureReplicationSchedule`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *StoragePureReplicationScheduleResponse) GetResultsOk() (*[]StoragePureReplicationSchedule, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *StoragePureReplicationScheduleResponse) SetResults(v []StoragePureReplicationSchedule)`

SetResults sets Results field to given value.

### HasResults

`func (o *StoragePureReplicationScheduleResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


