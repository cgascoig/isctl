# WorkflowApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AssetTargetMoid** | Pointer to **string** | Asset target defines the remote target endpoints which are either managed by Intersight or their service APIs are invoked from Intersight. Generic API executor service Jasmine can invoke these remote APIs via its executors. Asset targets can be accessed directly for cloud targets or via an associated Intersight Assist. Prerequisite to use asset targets is to persist the target details. Asset target MoRef can be given as task input of type TargetDataType. Fusion determines and populates the target context with the Assist if associated with. It is set internally at the API level. In case of an associated assist, it is used by Assist to fetch the target details and form the API request to send to endpoints. In case of cloud asset targets, Jasmine fetched the target details from DB, forms the API request and sends it to the target. | [optional] [readonly] 
**Body** | Pointer to **string** | The optional request body that is sent as part of this API request. The request body can contain a golang template that can be populated with task input parameters and previous API output parameters. | [optional] 
**ContentType** | Pointer to **string** | Intersight Orchestrator, with the support of response parser specification, can extract the values from API responses and map them to task output parameters. The value extraction is supported for response content types XML, JSON and Text. The type of the content that gets passed as payload and response in this API. The supported values are json, xml, text. | [optional] 
**Description** | Pointer to **string** | A description that task designer can add to individual API requests that explain  what the API call is about. | [optional] 
**ErrorContentType** | Pointer to **string** | Intersight Orchestrator, with the support of response parser specification, can extract the values from API responses and map them to task output parameters. The value extraction is supported for response content types XML, JSON and Text. Optional input to specify the content type in case of error API response. This should be used if the content type of error response is different from that of the success response. If not specified, contentType input value is used to parse the error response. | [optional] 
**Label** | Pointer to **string** | A user friendly label that task designers have given to the batch API request. | [optional] 
**Name** | Pointer to **string** | A reference name for this API request within the batch API request. This name shall be used to map the API output parameters to subsequent API input parameters within a batch API task. | [optional] 
**Outcomes** | Pointer to **interface{}** | All the possible outcomes of this API are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as &#39;true&#39;. This is an optional property and if not specified the task will be marked as success. | [optional] 
**ResponseSpec** | Pointer to **interface{}** | The optional grammar specification for parsing the response to extract the required values. The specification should have extraction specification specified for all the API output parameters. | [optional] 
**SkipOnCondition** | Pointer to **string** | The skip expression, if provided, allows the batch API executor to skip the api execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed. | [optional] 
**StartDelay** | Pointer to **int64** | The delay in seconds after which the API needs to be executed. By default, the given API is executed immediately. Specifying a start delay adds to the delay to execution. Start Delay is not supported for the first API in the Batch and cumulative delay of all the APIs in the Batch should not exceed the task time out. | [optional] 
**Timeout** | Pointer to **int64** | The duration in seconds by which the API response is expected from the API target. If the end point does not respond for the API request within this timeout duration, the task will be marked as failed. | [optional] 

## Methods

### NewWorkflowApi

`func NewWorkflowApi(classId string, objectType string, ) *WorkflowApi`

NewWorkflowApi instantiates a new WorkflowApi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowApiWithDefaults

`func NewWorkflowApiWithDefaults() *WorkflowApi`

NewWorkflowApiWithDefaults instantiates a new WorkflowApi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowApi) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowApi) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowApi) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowApi) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowApi) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowApi) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAssetTargetMoid

`func (o *WorkflowApi) GetAssetTargetMoid() string`

GetAssetTargetMoid returns the AssetTargetMoid field if non-nil, zero value otherwise.

### GetAssetTargetMoidOk

`func (o *WorkflowApi) GetAssetTargetMoidOk() (*string, bool)`

GetAssetTargetMoidOk returns a tuple with the AssetTargetMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTargetMoid

`func (o *WorkflowApi) SetAssetTargetMoid(v string)`

SetAssetTargetMoid sets AssetTargetMoid field to given value.

### HasAssetTargetMoid

`func (o *WorkflowApi) HasAssetTargetMoid() bool`

HasAssetTargetMoid returns a boolean if a field has been set.

### GetBody

`func (o *WorkflowApi) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WorkflowApi) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WorkflowApi) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *WorkflowApi) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetContentType

`func (o *WorkflowApi) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *WorkflowApi) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *WorkflowApi) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *WorkflowApi) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowApi) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowApi) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowApi) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowApi) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrorContentType

`func (o *WorkflowApi) GetErrorContentType() string`

GetErrorContentType returns the ErrorContentType field if non-nil, zero value otherwise.

### GetErrorContentTypeOk

`func (o *WorkflowApi) GetErrorContentTypeOk() (*string, bool)`

GetErrorContentTypeOk returns a tuple with the ErrorContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorContentType

`func (o *WorkflowApi) SetErrorContentType(v string)`

SetErrorContentType sets ErrorContentType field to given value.

### HasErrorContentType

`func (o *WorkflowApi) HasErrorContentType() bool`

HasErrorContentType returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowApi) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowApi) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowApi) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowApi) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowApi) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowApi) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowApi) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowApi) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutcomes

`func (o *WorkflowApi) GetOutcomes() interface{}`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *WorkflowApi) GetOutcomesOk() (*interface{}, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *WorkflowApi) SetOutcomes(v interface{})`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *WorkflowApi) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomesNil

`func (o *WorkflowApi) SetOutcomesNil(b bool)`

 SetOutcomesNil sets the value for Outcomes to be an explicit nil

### UnsetOutcomes
`func (o *WorkflowApi) UnsetOutcomes()`

UnsetOutcomes ensures that no value is present for Outcomes, not even an explicit nil
### GetResponseSpec

`func (o *WorkflowApi) GetResponseSpec() interface{}`

GetResponseSpec returns the ResponseSpec field if non-nil, zero value otherwise.

### GetResponseSpecOk

`func (o *WorkflowApi) GetResponseSpecOk() (*interface{}, bool)`

GetResponseSpecOk returns a tuple with the ResponseSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSpec

`func (o *WorkflowApi) SetResponseSpec(v interface{})`

SetResponseSpec sets ResponseSpec field to given value.

### HasResponseSpec

`func (o *WorkflowApi) HasResponseSpec() bool`

HasResponseSpec returns a boolean if a field has been set.

### SetResponseSpecNil

`func (o *WorkflowApi) SetResponseSpecNil(b bool)`

 SetResponseSpecNil sets the value for ResponseSpec to be an explicit nil

### UnsetResponseSpec
`func (o *WorkflowApi) UnsetResponseSpec()`

UnsetResponseSpec ensures that no value is present for ResponseSpec, not even an explicit nil
### GetSkipOnCondition

`func (o *WorkflowApi) GetSkipOnCondition() string`

GetSkipOnCondition returns the SkipOnCondition field if non-nil, zero value otherwise.

### GetSkipOnConditionOk

`func (o *WorkflowApi) GetSkipOnConditionOk() (*string, bool)`

GetSkipOnConditionOk returns a tuple with the SkipOnCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnCondition

`func (o *WorkflowApi) SetSkipOnCondition(v string)`

SetSkipOnCondition sets SkipOnCondition field to given value.

### HasSkipOnCondition

`func (o *WorkflowApi) HasSkipOnCondition() bool`

HasSkipOnCondition returns a boolean if a field has been set.

### GetStartDelay

`func (o *WorkflowApi) GetStartDelay() int64`

GetStartDelay returns the StartDelay field if non-nil, zero value otherwise.

### GetStartDelayOk

`func (o *WorkflowApi) GetStartDelayOk() (*int64, bool)`

GetStartDelayOk returns a tuple with the StartDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDelay

`func (o *WorkflowApi) SetStartDelay(v int64)`

SetStartDelay sets StartDelay field to given value.

### HasStartDelay

`func (o *WorkflowApi) HasStartDelay() bool`

HasStartDelay returns a boolean if a field has been set.

### GetTimeout

`func (o *WorkflowApi) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *WorkflowApi) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *WorkflowApi) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *WorkflowApi) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


