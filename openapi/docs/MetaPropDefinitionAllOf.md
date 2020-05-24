# MetaPropDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiAccess** | Pointer to **string** | API access control for the property. Examples are NoAccess, ReadOnly, CreateOnly etc. | [optional] [readonly] [default to "NoAccess"]
**Name** | Pointer to **string** | The name of the property. | [optional] [readonly] 
**OpSecurity** | Pointer to **string** | The data-at-rest security protection applied to this property when a Managed Object is persisted. For example, Encrypted or Cleartext. | [optional] [readonly] [default to "ClearText"]
**SearchWeight** | Pointer to **float32** | Enables the property to be searchable from global search. A value of 0 means this property is not globally searchable. | [optional] [readonly] 

## Methods

### NewMetaPropDefinitionAllOf

`func NewMetaPropDefinitionAllOf() *MetaPropDefinitionAllOf`

NewMetaPropDefinitionAllOf instantiates a new MetaPropDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaPropDefinitionAllOfWithDefaults

`func NewMetaPropDefinitionAllOfWithDefaults() *MetaPropDefinitionAllOf`

NewMetaPropDefinitionAllOfWithDefaults instantiates a new MetaPropDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiAccess

`func (o *MetaPropDefinitionAllOf) GetApiAccess() string`

GetApiAccess returns the ApiAccess field if non-nil, zero value otherwise.

### GetApiAccessOk

`func (o *MetaPropDefinitionAllOf) GetApiAccessOk() (*string, bool)`

GetApiAccessOk returns a tuple with the ApiAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccess

`func (o *MetaPropDefinitionAllOf) SetApiAccess(v string)`

SetApiAccess sets ApiAccess field to given value.

### HasApiAccess

`func (o *MetaPropDefinitionAllOf) HasApiAccess() bool`

HasApiAccess returns a boolean if a field has been set.

### GetName

`func (o *MetaPropDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaPropDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaPropDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaPropDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpSecurity

`func (o *MetaPropDefinitionAllOf) GetOpSecurity() string`

GetOpSecurity returns the OpSecurity field if non-nil, zero value otherwise.

### GetOpSecurityOk

`func (o *MetaPropDefinitionAllOf) GetOpSecurityOk() (*string, bool)`

GetOpSecurityOk returns a tuple with the OpSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpSecurity

`func (o *MetaPropDefinitionAllOf) SetOpSecurity(v string)`

SetOpSecurity sets OpSecurity field to given value.

### HasOpSecurity

`func (o *MetaPropDefinitionAllOf) HasOpSecurity() bool`

HasOpSecurity returns a boolean if a field has been set.

### GetSearchWeight

`func (o *MetaPropDefinitionAllOf) GetSearchWeight() float32`

GetSearchWeight returns the SearchWeight field if non-nil, zero value otherwise.

### GetSearchWeightOk

`func (o *MetaPropDefinitionAllOf) GetSearchWeightOk() (*float32, bool)`

GetSearchWeightOk returns a tuple with the SearchWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchWeight

`func (o *MetaPropDefinitionAllOf) SetSearchWeight(v float32)`

SetSearchWeight sets SearchWeight field to given value.

### HasSearchWeight

`func (o *MetaPropDefinitionAllOf) HasSearchWeight() bool`

HasSearchWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


