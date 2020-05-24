# WorkflowApiAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | The optional request body that is sent as part of this API request. The request body can contain a golang template that can be populated with task input parameters and previous API output parameters. | [optional] 
**ContentType** | Pointer to **string** | Intersight Orchestrator, with the support of response parser specification, can extract the values from API responses and map them to task output parameters. The value extraction is supported for response content types XML and JSON. The type of the content that gets passed as payload and response in this API. | [optional] [default to "json"]
**Name** | Pointer to **string** | A reference name for this API request within the batch API request. This name shall be used to map the API output parameters to subsequent API input parameters within a batch API task. | [optional] 
**Outcomes** | Pointer to **map[string]interface{}** | All the possible outcomes of this API are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as &#39;true&#39;. This is an optional property and if not specified the task will be marked as success. | [optional] 
**ResponseSpec** | Pointer to [**ContentGrammar**](content.Grammar.md) |  | [optional] 
**SkipOnCondition** | Pointer to **string** | The skip expression, if provided, allows the batch API executor to skip the api execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed. | [optional] 
**Timeout** | Pointer to **int64** | The duration in seconds by which the API response is expected from the API target. If the end point does not respond for the API request within this timeout duration, the task will be marked as failed. | [optional] 

## Methods

### NewWorkflowApiAllOf

`func NewWorkflowApiAllOf() *WorkflowApiAllOf`

NewWorkflowApiAllOf instantiates a new WorkflowApiAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowApiAllOfWithDefaults

`func NewWorkflowApiAllOfWithDefaults() *WorkflowApiAllOf`

NewWorkflowApiAllOfWithDefaults instantiates a new WorkflowApiAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *WorkflowApiAllOf) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WorkflowApiAllOf) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WorkflowApiAllOf) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *WorkflowApiAllOf) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetContentType

`func (o *WorkflowApiAllOf) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *WorkflowApiAllOf) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *WorkflowApiAllOf) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *WorkflowApiAllOf) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetName

`func (o *WorkflowApiAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowApiAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowApiAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowApiAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutcomes

`func (o *WorkflowApiAllOf) GetOutcomes() map[string]interface{}`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *WorkflowApiAllOf) GetOutcomesOk() (*map[string]interface{}, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *WorkflowApiAllOf) SetOutcomes(v map[string]interface{})`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *WorkflowApiAllOf) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### GetResponseSpec

`func (o *WorkflowApiAllOf) GetResponseSpec() ContentGrammar`

GetResponseSpec returns the ResponseSpec field if non-nil, zero value otherwise.

### GetResponseSpecOk

`func (o *WorkflowApiAllOf) GetResponseSpecOk() (*ContentGrammar, bool)`

GetResponseSpecOk returns a tuple with the ResponseSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSpec

`func (o *WorkflowApiAllOf) SetResponseSpec(v ContentGrammar)`

SetResponseSpec sets ResponseSpec field to given value.

### HasResponseSpec

`func (o *WorkflowApiAllOf) HasResponseSpec() bool`

HasResponseSpec returns a boolean if a field has been set.

### GetSkipOnCondition

`func (o *WorkflowApiAllOf) GetSkipOnCondition() string`

GetSkipOnCondition returns the SkipOnCondition field if non-nil, zero value otherwise.

### GetSkipOnConditionOk

`func (o *WorkflowApiAllOf) GetSkipOnConditionOk() (*string, bool)`

GetSkipOnConditionOk returns a tuple with the SkipOnCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnCondition

`func (o *WorkflowApiAllOf) SetSkipOnCondition(v string)`

SetSkipOnCondition sets SkipOnCondition field to given value.

### HasSkipOnCondition

`func (o *WorkflowApiAllOf) HasSkipOnCondition() bool`

HasSkipOnCondition returns a boolean if a field has been set.

### GetTimeout

`func (o *WorkflowApiAllOf) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *WorkflowApiAllOf) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *WorkflowApiAllOf) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *WorkflowApiAllOf) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


