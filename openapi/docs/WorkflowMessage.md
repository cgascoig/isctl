# WorkflowMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **interface{}** | An i18n message that can be translated in multiple languages to support internationalization. type: string | [optional] 
**Severity** | Pointer to **string** | The severity of the Task or Workflow message warning/error/info etc. | [optional] [default to "Info"]

## Methods

### NewWorkflowMessage

`func NewWorkflowMessage() *WorkflowMessage`

NewWorkflowMessage instantiates a new WorkflowMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMessageWithDefaults

`func NewWorkflowMessageWithDefaults() *WorkflowMessage`

NewWorkflowMessageWithDefaults instantiates a new WorkflowMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *WorkflowMessage) GetMessage() interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkflowMessage) GetMessageOk() (*interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkflowMessage) SetMessage(v interface{})`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WorkflowMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *WorkflowMessage) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *WorkflowMessage) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSeverity

`func (o *WorkflowMessage) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *WorkflowMessage) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *WorkflowMessage) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *WorkflowMessage) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


