# ApplianceStatusCheckAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.StatusCheck"]
**ObjectType** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.StatusCheck"]
**Code** | Pointer to **string** | Unique identifier of the status check. | [optional] 
**Result** | Pointer to **string** | Result of this status check. * &#x60;OK&#x60; - Result of the check is OK. * &#x60;Warning&#x60; - Result of the check is Warning. * &#x60;Critical&#x60; - Result of the check is Critical. * &#x60;Info&#x60; - Result of the check is low. | [optional] [default to "OK"]

## Methods

### NewApplianceStatusCheckAllOf

`func NewApplianceStatusCheckAllOf(classId string, objectType string, ) *ApplianceStatusCheckAllOf`

NewApplianceStatusCheckAllOf instantiates a new ApplianceStatusCheckAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceStatusCheckAllOfWithDefaults

`func NewApplianceStatusCheckAllOfWithDefaults() *ApplianceStatusCheckAllOf`

NewApplianceStatusCheckAllOfWithDefaults instantiates a new ApplianceStatusCheckAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceStatusCheckAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceStatusCheckAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceStatusCheckAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceStatusCheckAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceStatusCheckAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceStatusCheckAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCode

`func (o *ApplianceStatusCheckAllOf) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApplianceStatusCheckAllOf) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApplianceStatusCheckAllOf) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApplianceStatusCheckAllOf) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetResult

`func (o *ApplianceStatusCheckAllOf) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ApplianceStatusCheckAllOf) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ApplianceStatusCheckAllOf) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ApplianceStatusCheckAllOf) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


