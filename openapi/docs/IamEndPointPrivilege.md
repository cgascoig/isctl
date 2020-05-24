# IamEndPointPrivilege

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The functionality of this privilege. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the end point privilege. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the end point like Cisco UCS Fabric Interconnect or Cisco IMC. | [optional] [readonly] [default to ""]
**System** | Pointer to [**IamSystemRelationship**](iam.System.Relationship.md) |  | [optional] 

## Methods

### NewIamEndPointPrivilege

`func NewIamEndPointPrivilege() *IamEndPointPrivilege`

NewIamEndPointPrivilege instantiates a new IamEndPointPrivilege object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointPrivilegeWithDefaults

`func NewIamEndPointPrivilegeWithDefaults() *IamEndPointPrivilege`

NewIamEndPointPrivilegeWithDefaults instantiates a new IamEndPointPrivilege object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IamEndPointPrivilege) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamEndPointPrivilege) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamEndPointPrivilege) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamEndPointPrivilege) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *IamEndPointPrivilege) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamEndPointPrivilege) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamEndPointPrivilege) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamEndPointPrivilege) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IamEndPointPrivilege) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamEndPointPrivilege) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamEndPointPrivilege) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamEndPointPrivilege) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSystem

`func (o *IamEndPointPrivilege) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IamEndPointPrivilege) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IamEndPointPrivilege) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IamEndPointPrivilege) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


