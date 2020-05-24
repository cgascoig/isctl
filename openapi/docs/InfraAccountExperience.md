# InfraAccountExperience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to [**[]InfraFeatureDefinition**](infra.FeatureDefinition.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewInfraAccountExperience

`func NewInfraAccountExperience() *InfraAccountExperience`

NewInfraAccountExperience instantiates a new InfraAccountExperience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraAccountExperienceWithDefaults

`func NewInfraAccountExperienceWithDefaults() *InfraAccountExperience`

NewInfraAccountExperienceWithDefaults instantiates a new InfraAccountExperience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *InfraAccountExperience) GetFeatures() []InfraFeatureDefinition`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *InfraAccountExperience) GetFeaturesOk() (*[]InfraFeatureDefinition, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *InfraAccountExperience) SetFeatures(v []InfraFeatureDefinition)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *InfraAccountExperience) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetAccount

`func (o *InfraAccountExperience) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InfraAccountExperience) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InfraAccountExperience) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *InfraAccountExperience) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


