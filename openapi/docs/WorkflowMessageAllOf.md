# WorkflowMessageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **interface{}** | An i18n message that can be translated in multiple languages to support internationalization. type: string | [optional] 
**Severity** | Pointer to **string** | The severity of the Task or Workflow message warning/error/info etc. | [optional] [default to "Info"]

## Methods

### NewWorkflowMessageAllOf

`func NewWorkflowMessageAllOf() *WorkflowMessageAllOf`

NewWorkflowMessageAllOf instantiates a new WorkflowMessageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMessageAllOfWithDefaults

`func NewWorkflowMessageAllOfWithDefaults() *WorkflowMessageAllOf`

NewWorkflowMessageAllOfWithDefaults instantiates a new WorkflowMessageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *WorkflowMessageAllOf) GetMessage() interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkflowMessageAllOf) GetMessageOk() (*interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkflowMessageAllOf) SetMessage(v interface{})`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WorkflowMessageAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *WorkflowMessageAllOf) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *WorkflowMessageAllOf) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSeverity

`func (o *WorkflowMessageAllOf) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *WorkflowMessageAllOf) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *WorkflowMessageAllOf) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *WorkflowMessageAllOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


