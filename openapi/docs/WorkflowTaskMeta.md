# WorkflowTaskMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to **string** | A task execution type to indicate if it is a system task. | [optional] 
**Description** | Pointer to **string** | A description of the task. | [optional] 
**InputKeys** | Pointer to **[]string** |  | [optional] 
**Internal** | Pointer to **bool** | Denotes whether or not this is an internal task.  Internal tasks will be hidden from the UI within a workflow. | [optional] 
**Name** | Pointer to **string** | A task name that should be unique in Conductor DB. | [optional] 
**OutputKeys** | Pointer to **[]string** |  | [optional] 
**ResponseTimeoutSec** | Pointer to **int64** | The worker respnose timeout value. | [optional] 
**RetryCount** | Pointer to **int64** | A number of reties for this task. | [optional] 
**RetryDelaySec** | Pointer to **int64** | The time on which the retry will be delayed. | [optional] 
**RetryLogic** | Pointer to **string** | A logic which defines the way to handle retry (FIXED, EXPONENTIAL_BACKOFF). | [optional] 
**Src** | Pointer to **string** | A service owns the task metadata. | [optional] 
**TimeoutPolicy** | Pointer to **string** | A policy which defines the way to handle timeout (RETRY, TIME_OUT_WF, ALERT_ONLY). | [optional] 
**TimeoutSec** | Pointer to **int64** | A timeout value for the task in seconds. | [optional] 

## Methods

### NewWorkflowTaskMeta

`func NewWorkflowTaskMeta() *WorkflowTaskMeta`

NewWorkflowTaskMeta instantiates a new WorkflowTaskMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskMetaWithDefaults

`func NewWorkflowTaskMetaWithDefaults() *WorkflowTaskMeta`

NewWorkflowTaskMetaWithDefaults instantiates a new WorkflowTaskMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *WorkflowTaskMeta) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *WorkflowTaskMeta) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *WorkflowTaskMeta) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *WorkflowTaskMeta) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowTaskMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowTaskMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowTaskMeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowTaskMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputKeys

`func (o *WorkflowTaskMeta) GetInputKeys() []string`

GetInputKeys returns the InputKeys field if non-nil, zero value otherwise.

### GetInputKeysOk

`func (o *WorkflowTaskMeta) GetInputKeysOk() (*[]string, bool)`

GetInputKeysOk returns a tuple with the InputKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputKeys

`func (o *WorkflowTaskMeta) SetInputKeys(v []string)`

SetInputKeys sets InputKeys field to given value.

### HasInputKeys

`func (o *WorkflowTaskMeta) HasInputKeys() bool`

HasInputKeys returns a boolean if a field has been set.

### GetInternal

`func (o *WorkflowTaskMeta) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *WorkflowTaskMeta) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *WorkflowTaskMeta) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *WorkflowTaskMeta) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTaskMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTaskMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTaskMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTaskMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputKeys

`func (o *WorkflowTaskMeta) GetOutputKeys() []string`

GetOutputKeys returns the OutputKeys field if non-nil, zero value otherwise.

### GetOutputKeysOk

`func (o *WorkflowTaskMeta) GetOutputKeysOk() (*[]string, bool)`

GetOutputKeysOk returns a tuple with the OutputKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputKeys

`func (o *WorkflowTaskMeta) SetOutputKeys(v []string)`

SetOutputKeys sets OutputKeys field to given value.

### HasOutputKeys

`func (o *WorkflowTaskMeta) HasOutputKeys() bool`

HasOutputKeys returns a boolean if a field has been set.

### GetResponseTimeoutSec

`func (o *WorkflowTaskMeta) GetResponseTimeoutSec() int64`

GetResponseTimeoutSec returns the ResponseTimeoutSec field if non-nil, zero value otherwise.

### GetResponseTimeoutSecOk

`func (o *WorkflowTaskMeta) GetResponseTimeoutSecOk() (*int64, bool)`

GetResponseTimeoutSecOk returns a tuple with the ResponseTimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeoutSec

`func (o *WorkflowTaskMeta) SetResponseTimeoutSec(v int64)`

SetResponseTimeoutSec sets ResponseTimeoutSec field to given value.

### HasResponseTimeoutSec

`func (o *WorkflowTaskMeta) HasResponseTimeoutSec() bool`

HasResponseTimeoutSec returns a boolean if a field has been set.

### GetRetryCount

`func (o *WorkflowTaskMeta) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *WorkflowTaskMeta) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *WorkflowTaskMeta) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *WorkflowTaskMeta) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySec

`func (o *WorkflowTaskMeta) GetRetryDelaySec() int64`

GetRetryDelaySec returns the RetryDelaySec field if non-nil, zero value otherwise.

### GetRetryDelaySecOk

`func (o *WorkflowTaskMeta) GetRetryDelaySecOk() (*int64, bool)`

GetRetryDelaySecOk returns a tuple with the RetryDelaySec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySec

`func (o *WorkflowTaskMeta) SetRetryDelaySec(v int64)`

SetRetryDelaySec sets RetryDelaySec field to given value.

### HasRetryDelaySec

`func (o *WorkflowTaskMeta) HasRetryDelaySec() bool`

HasRetryDelaySec returns a boolean if a field has been set.

### GetRetryLogic

`func (o *WorkflowTaskMeta) GetRetryLogic() string`

GetRetryLogic returns the RetryLogic field if non-nil, zero value otherwise.

### GetRetryLogicOk

`func (o *WorkflowTaskMeta) GetRetryLogicOk() (*string, bool)`

GetRetryLogicOk returns a tuple with the RetryLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryLogic

`func (o *WorkflowTaskMeta) SetRetryLogic(v string)`

SetRetryLogic sets RetryLogic field to given value.

### HasRetryLogic

`func (o *WorkflowTaskMeta) HasRetryLogic() bool`

HasRetryLogic returns a boolean if a field has been set.

### GetSrc

`func (o *WorkflowTaskMeta) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *WorkflowTaskMeta) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *WorkflowTaskMeta) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *WorkflowTaskMeta) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetTimeoutPolicy

`func (o *WorkflowTaskMeta) GetTimeoutPolicy() string`

GetTimeoutPolicy returns the TimeoutPolicy field if non-nil, zero value otherwise.

### GetTimeoutPolicyOk

`func (o *WorkflowTaskMeta) GetTimeoutPolicyOk() (*string, bool)`

GetTimeoutPolicyOk returns a tuple with the TimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutPolicy

`func (o *WorkflowTaskMeta) SetTimeoutPolicy(v string)`

SetTimeoutPolicy sets TimeoutPolicy field to given value.

### HasTimeoutPolicy

`func (o *WorkflowTaskMeta) HasTimeoutPolicy() bool`

HasTimeoutPolicy returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *WorkflowTaskMeta) GetTimeoutSec() int64`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *WorkflowTaskMeta) GetTimeoutSecOk() (*int64, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *WorkflowTaskMeta) SetTimeoutSec(v int64)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *WorkflowTaskMeta) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


