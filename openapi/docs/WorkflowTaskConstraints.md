# WorkflowTaskConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetDataType** | Pointer to **map[string]interface{}** | List of property constraints that helps to narrow down task implementations based on target device input. | [optional] 

## Methods

### NewWorkflowTaskConstraints

`func NewWorkflowTaskConstraints() *WorkflowTaskConstraints`

NewWorkflowTaskConstraints instantiates a new WorkflowTaskConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskConstraintsWithDefaults

`func NewWorkflowTaskConstraintsWithDefaults() *WorkflowTaskConstraints`

NewWorkflowTaskConstraintsWithDefaults instantiates a new WorkflowTaskConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetDataType

`func (o *WorkflowTaskConstraints) GetTargetDataType() map[string]interface{}`

GetTargetDataType returns the TargetDataType field if non-nil, zero value otherwise.

### GetTargetDataTypeOk

`func (o *WorkflowTaskConstraints) GetTargetDataTypeOk() (*map[string]interface{}, bool)`

GetTargetDataTypeOk returns a tuple with the TargetDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDataType

`func (o *WorkflowTaskConstraints) SetTargetDataType(v map[string]interface{})`

SetTargetDataType sets TargetDataType field to given value.

### HasTargetDataType

`func (o *WorkflowTaskConstraints) HasTargetDataType() bool`

HasTargetDataType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


