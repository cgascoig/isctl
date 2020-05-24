# ResourceMembershipHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this resource membership holder. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Memberships** | Pointer to [**[]ResourceMembershipRelationship**](resource.Membership.Relationship.md) | An array of relationships to resourceMembership resources. | [optional] [readonly] 

## Methods

### NewResourceMembershipHolder

`func NewResourceMembershipHolder() *ResourceMembershipHolder`

NewResourceMembershipHolder instantiates a new ResourceMembershipHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMembershipHolderWithDefaults

`func NewResourceMembershipHolderWithDefaults() *ResourceMembershipHolder`

NewResourceMembershipHolderWithDefaults instantiates a new ResourceMembershipHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourceMembershipHolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceMembershipHolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceMembershipHolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceMembershipHolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *ResourceMembershipHolder) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ResourceMembershipHolder) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ResourceMembershipHolder) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ResourceMembershipHolder) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetMemberships

`func (o *ResourceMembershipHolder) GetMemberships() []ResourceMembershipRelationship`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *ResourceMembershipHolder) GetMembershipsOk() (*[]ResourceMembershipRelationship, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *ResourceMembershipHolder) SetMemberships(v []ResourceMembershipRelationship)`

SetMemberships sets Memberships field to given value.

### HasMemberships

`func (o *ResourceMembershipHolder) HasMemberships() bool`

HasMemberships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


