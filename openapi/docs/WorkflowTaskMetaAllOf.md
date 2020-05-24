# WorkflowTaskMetaAllOf

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

### NewWorkflowTaskMetaAllOf

`func NewWorkflowTaskMetaAllOf() *WorkflowTaskMetaAllOf`

NewWorkflowTaskMetaAllOf instantiates a new WorkflowTaskMetaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskMetaAllOfWithDefaults

`func NewWorkflowTaskMetaAllOfWithDefaults() *WorkflowTaskMetaAllOf`

NewWorkflowTaskMetaAllOfWithDefaults instantiates a new WorkflowTaskMetaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *WorkflowTaskMetaAllOf) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *WorkflowTaskMetaAllOf) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *WorkflowTaskMetaAllOf) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *WorkflowTaskMetaAllOf) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowTaskMetaAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowTaskMetaAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowTaskMetaAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowTaskMetaAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputKeys

`func (o *WorkflowTaskMetaAllOf) GetInputKeys() []string`

GetInputKeys returns the InputKeys field if non-nil, zero value otherwise.

### GetInputKeysOk

`func (o *WorkflowTaskMetaAllOf) GetInputKeysOk() (*[]string, bool)`

GetInputKeysOk returns a tuple with the InputKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputKeys

`func (o *WorkflowTaskMetaAllOf) SetInputKeys(v []string)`

SetInputKeys sets InputKeys field to given value.

### HasInputKeys

`func (o *WorkflowTaskMetaAllOf) HasInputKeys() bool`

HasInputKeys returns a boolean if a field has been set.

### GetInternal

`func (o *WorkflowTaskMetaAllOf) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *WorkflowTaskMetaAllOf) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *WorkflowTaskMetaAllOf) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *WorkflowTaskMetaAllOf) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTaskMetaAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTaskMetaAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTaskMetaAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTaskMetaAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputKeys

`func (o *WorkflowTaskMetaAllOf) GetOutputKeys() []string`

GetOutputKeys returns the OutputKeys field if non-nil, zero value otherwise.

### GetOutputKeysOk

`func (o *WorkflowTaskMetaAllOf) GetOutputKeysOk() (*[]string, bool)`

GetOutputKeysOk returns a tuple with the OutputKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputKeys

`func (o *WorkflowTaskMetaAllOf) SetOutputKeys(v []string)`

SetOutputKeys sets OutputKeys field to given value.

### HasOutputKeys

`func (o *WorkflowTaskMetaAllOf) HasOutputKeys() bool`

HasOutputKeys returns a boolean if a field has been set.

### GetResponseTimeoutSec

`func (o *WorkflowTaskMetaAllOf) GetResponseTimeoutSec() int64`

GetResponseTimeoutSec returns the ResponseTimeoutSec field if non-nil, zero value otherwise.

### GetResponseTimeoutSecOk

`func (o *WorkflowTaskMetaAllOf) GetResponseTimeoutSecOk() (*int64, bool)`

GetResponseTimeoutSecOk returns a tuple with the ResponseTimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeoutSec

`func (o *WorkflowTaskMetaAllOf) SetResponseTimeoutSec(v int64)`

SetResponseTimeoutSec sets ResponseTimeoutSec field to given value.

### HasResponseTimeoutSec

`func (o *WorkflowTaskMetaAllOf) HasResponseTimeoutSec() bool`

HasResponseTimeoutSec returns a boolean if a field has been set.

### GetRetryCount

`func (o *WorkflowTaskMetaAllOf) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *WorkflowTaskMetaAllOf) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *WorkflowTaskMetaAllOf) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *WorkflowTaskMetaAllOf) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySec

`func (o *WorkflowTaskMetaAllOf) GetRetryDelaySec() int64`

GetRetryDelaySec returns the RetryDelaySec field if non-nil, zero value otherwise.

### GetRetryDelaySecOk

`func (o *WorkflowTaskMetaAllOf) GetRetryDelaySecOk() (*int64, bool)`

GetRetryDelaySecOk returns a tuple with the RetryDelaySec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySec

`func (o *WorkflowTaskMetaAllOf) SetRetryDelaySec(v int64)`

SetRetryDelaySec sets RetryDelaySec field to given value.

### HasRetryDelaySec

`func (o *WorkflowTaskMetaAllOf) HasRetryDelaySec() bool`

HasRetryDelaySec returns a boolean if a field has been set.

### GetRetryLogic

`func (o *WorkflowTaskMetaAllOf) GetRetryLogic() string`

GetRetryLogic returns the RetryLogic field if non-nil, zero value otherwise.

### GetRetryLogicOk

`func (o *WorkflowTaskMetaAllOf) GetRetryLogicOk() (*string, bool)`

GetRetryLogicOk returns a tuple with the RetryLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryLogic

`func (o *WorkflowTaskMetaAllOf) SetRetryLogic(v string)`

SetRetryLogic sets RetryLogic field to given value.

### HasRetryLogic

`func (o *WorkflowTaskMetaAllOf) HasRetryLogic() bool`

HasRetryLogic returns a boolean if a field has been set.

### GetSrc

`func (o *WorkflowTaskMetaAllOf) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *WorkflowTaskMetaAllOf) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *WorkflowTaskMetaAllOf) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *WorkflowTaskMetaAllOf) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetTimeoutPolicy

`func (o *WorkflowTaskMetaAllOf) GetTimeoutPolicy() string`

GetTimeoutPolicy returns the TimeoutPolicy field if non-nil, zero value otherwise.

### GetTimeoutPolicyOk

`func (o *WorkflowTaskMetaAllOf) GetTimeoutPolicyOk() (*string, bool)`

GetTimeoutPolicyOk returns a tuple with the TimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutPolicy

`func (o *WorkflowTaskMetaAllOf) SetTimeoutPolicy(v string)`

SetTimeoutPolicy sets TimeoutPolicy field to given value.

### HasTimeoutPolicy

`func (o *WorkflowTaskMetaAllOf) HasTimeoutPolicy() bool`

HasTimeoutPolicy returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *WorkflowTaskMetaAllOf) GetTimeoutSec() int64`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *WorkflowTaskMetaAllOf) GetTimeoutSecOk() (*int64, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *WorkflowTaskMetaAllOf) SetTimeoutSec(v int64)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *WorkflowTaskMetaAllOf) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


