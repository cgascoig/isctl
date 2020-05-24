# StorageInitiatorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iqn** | Pointer to **string** | IQN (iSCSI qualified name). Can be up to 255 characters long and has the format iqn.yyyy-mm.naming-authority:unique name. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the initiator represented in the storage array. | [optional] [readonly] 
**Type** | Pointer to **string** | Initiator type, it can be FC or iSCSI. | [optional] [readonly] [default to "FC"]
**Wwn** | Pointer to **string** | World wide name, 128 bit address represented in hexadecimal notation. For example, 51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01. | [optional] [readonly] 

## Methods

### NewStorageInitiatorAllOf

`func NewStorageInitiatorAllOf() *StorageInitiatorAllOf`

NewStorageInitiatorAllOf instantiates a new StorageInitiatorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageInitiatorAllOfWithDefaults

`func NewStorageInitiatorAllOfWithDefaults() *StorageInitiatorAllOf`

NewStorageInitiatorAllOfWithDefaults instantiates a new StorageInitiatorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIqn

`func (o *StorageInitiatorAllOf) GetIqn() string`

GetIqn returns the Iqn field if non-nil, zero value otherwise.

### GetIqnOk

`func (o *StorageInitiatorAllOf) GetIqnOk() (*string, bool)`

GetIqnOk returns a tuple with the Iqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqn

`func (o *StorageInitiatorAllOf) SetIqn(v string)`

SetIqn sets Iqn field to given value.

### HasIqn

`func (o *StorageInitiatorAllOf) HasIqn() bool`

HasIqn returns a boolean if a field has been set.

### GetName

`func (o *StorageInitiatorAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageInitiatorAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageInitiatorAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageInitiatorAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *StorageInitiatorAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageInitiatorAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageInitiatorAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageInitiatorAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWwn

`func (o *StorageInitiatorAllOf) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *StorageInitiatorAllOf) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *StorageInitiatorAllOf) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *StorageInitiatorAllOf) HasWwn() bool`

HasWwn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


