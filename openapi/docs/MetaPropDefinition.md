# MetaPropDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiAccess** | Pointer to **string** | API access control for the property. Examples are NoAccess, ReadOnly, CreateOnly etc. | [optional] [readonly] [default to "NoAccess"]
**Name** | Pointer to **string** | The name of the property. | [optional] [readonly] 
**OpSecurity** | Pointer to **string** | The data-at-rest security protection applied to this property when a Managed Object is persisted. For example, Encrypted or Cleartext. | [optional] [readonly] [default to "ClearText"]
**SearchWeight** | Pointer to **float32** | Enables the property to be searchable from global search. A value of 0 means this property is not globally searchable. | [optional] [readonly] 

## Methods

### NewMetaPropDefinition

`func NewMetaPropDefinition() *MetaPropDefinition`

NewMetaPropDefinition instantiates a new MetaPropDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaPropDefinitionWithDefaults

`func NewMetaPropDefinitionWithDefaults() *MetaPropDefinition`

NewMetaPropDefinitionWithDefaults instantiates a new MetaPropDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiAccess

`func (o *MetaPropDefinition) GetApiAccess() string`

GetApiAccess returns the ApiAccess field if non-nil, zero value otherwise.

### GetApiAccessOk

`func (o *MetaPropDefinition) GetApiAccessOk() (*string, bool)`

GetApiAccessOk returns a tuple with the ApiAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccess

`func (o *MetaPropDefinition) SetApiAccess(v string)`

SetApiAccess sets ApiAccess field to given value.

### HasApiAccess

`func (o *MetaPropDefinition) HasApiAccess() bool`

HasApiAccess returns a boolean if a field has been set.

### GetName

`func (o *MetaPropDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaPropDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaPropDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaPropDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpSecurity

`func (o *MetaPropDefinition) GetOpSecurity() string`

GetOpSecurity returns the OpSecurity field if non-nil, zero value otherwise.

### GetOpSecurityOk

`func (o *MetaPropDefinition) GetOpSecurityOk() (*string, bool)`

GetOpSecurityOk returns a tuple with the OpSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpSecurity

`func (o *MetaPropDefinition) SetOpSecurity(v string)`

SetOpSecurity sets OpSecurity field to given value.

### HasOpSecurity

`func (o *MetaPropDefinition) HasOpSecurity() bool`

HasOpSecurity returns a boolean if a field has been set.

### GetSearchWeight

`func (o *MetaPropDefinition) GetSearchWeight() float32`

GetSearchWeight returns the SearchWeight field if non-nil, zero value otherwise.

### GetSearchWeightOk

`func (o *MetaPropDefinition) GetSearchWeightOk() (*float32, bool)`

GetSearchWeightOk returns a tuple with the SearchWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchWeight

`func (o *MetaPropDefinition) SetSearchWeight(v float32)`

SetSearchWeight sets SearchWeight field to given value.

### HasSearchWeight

`func (o *MetaPropDefinition) HasSearchWeight() bool`

HasSearchWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


