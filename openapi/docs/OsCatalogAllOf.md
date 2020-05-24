# OsCatalogAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The catalog name. There will be one for system and one for each user account. | [optional] 
**ConfigurationFiles** | Pointer to [**[]OsConfigurationFileRelationship**](os.ConfigurationFile.Relationship.md) | An array of relationships to osConfigurationFile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewOsCatalogAllOf

`func NewOsCatalogAllOf() *OsCatalogAllOf`

NewOsCatalogAllOf instantiates a new OsCatalogAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsCatalogAllOfWithDefaults

`func NewOsCatalogAllOfWithDefaults() *OsCatalogAllOf`

NewOsCatalogAllOfWithDefaults instantiates a new OsCatalogAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OsCatalogAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsCatalogAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsCatalogAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsCatalogAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfigurationFiles

`func (o *OsCatalogAllOf) GetConfigurationFiles() []OsConfigurationFileRelationship`

GetConfigurationFiles returns the ConfigurationFiles field if non-nil, zero value otherwise.

### GetConfigurationFilesOk

`func (o *OsCatalogAllOf) GetConfigurationFilesOk() (*[]OsConfigurationFileRelationship, bool)`

GetConfigurationFilesOk returns a tuple with the ConfigurationFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFiles

`func (o *OsCatalogAllOf) SetConfigurationFiles(v []OsConfigurationFileRelationship)`

SetConfigurationFiles sets ConfigurationFiles field to given value.

### HasConfigurationFiles

`func (o *OsCatalogAllOf) HasConfigurationFiles() bool`

HasConfigurationFiles returns a boolean if a field has been set.

### GetOrganization

`func (o *OsCatalogAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OsCatalogAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OsCatalogAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OsCatalogAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


