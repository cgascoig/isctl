# IamEndPointPrivilegeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The functionality of this privilege. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the end point privilege. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the end point like Cisco UCS Fabric Interconnect or Cisco IMC. | [optional] [readonly] [default to ""]
**System** | Pointer to [**IamSystemRelationship**](iam.System.Relationship.md) |  | [optional] 

## Methods

### NewIamEndPointPrivilegeAllOf

`func NewIamEndPointPrivilegeAllOf() *IamEndPointPrivilegeAllOf`

NewIamEndPointPrivilegeAllOf instantiates a new IamEndPointPrivilegeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointPrivilegeAllOfWithDefaults

`func NewIamEndPointPrivilegeAllOfWithDefaults() *IamEndPointPrivilegeAllOf`

NewIamEndPointPrivilegeAllOfWithDefaults instantiates a new IamEndPointPrivilegeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IamEndPointPrivilegeAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamEndPointPrivilegeAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamEndPointPrivilegeAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamEndPointPrivilegeAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *IamEndPointPrivilegeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamEndPointPrivilegeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamEndPointPrivilegeAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamEndPointPrivilegeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IamEndPointPrivilegeAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamEndPointPrivilegeAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamEndPointPrivilegeAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamEndPointPrivilegeAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSystem

`func (o *IamEndPointPrivilegeAllOf) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IamEndPointPrivilegeAllOf) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IamEndPointPrivilegeAllOf) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IamEndPointPrivilegeAllOf) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


