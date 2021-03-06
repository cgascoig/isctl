# AssetUsernamePasswordCredentialAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.UsernamePasswordCredential"]
**ObjectType** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.UsernamePasswordCredential"]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | The password used to authenticate with a managed target. | [optional] 
**Username** | Pointer to **string** | The username used to authenticate with a managed target. | [optional] 

## Methods

### NewAssetUsernamePasswordCredentialAllOf

`func NewAssetUsernamePasswordCredentialAllOf(classId string, objectType string, ) *AssetUsernamePasswordCredentialAllOf`

NewAssetUsernamePasswordCredentialAllOf instantiates a new AssetUsernamePasswordCredentialAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetUsernamePasswordCredentialAllOfWithDefaults

`func NewAssetUsernamePasswordCredentialAllOfWithDefaults() *AssetUsernamePasswordCredentialAllOf`

NewAssetUsernamePasswordCredentialAllOfWithDefaults instantiates a new AssetUsernamePasswordCredentialAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetUsernamePasswordCredentialAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetUsernamePasswordCredentialAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetUsernamePasswordCredentialAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetUsernamePasswordCredentialAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetUsernamePasswordCredentialAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetUsernamePasswordCredentialAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPasswordSet

`func (o *AssetUsernamePasswordCredentialAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *AssetUsernamePasswordCredentialAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *AssetUsernamePasswordCredentialAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *AssetUsernamePasswordCredentialAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *AssetUsernamePasswordCredentialAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AssetUsernamePasswordCredentialAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AssetUsernamePasswordCredentialAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AssetUsernamePasswordCredentialAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *AssetUsernamePasswordCredentialAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AssetUsernamePasswordCredentialAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AssetUsernamePasswordCredentialAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AssetUsernamePasswordCredentialAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


