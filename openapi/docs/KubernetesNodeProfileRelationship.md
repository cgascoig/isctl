# KubernetesNodeProfileRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mo.MoRef"]
**ObjectType** | Pointer to **string** | The fully-qualified name of the remote type referred by this relationship. | 
**Moid** | Pointer to **string** | The Moid of the referenced REST resource. | [optional] [readonly] 
**Selector** | Pointer to **string** | An OData $filter expression which describes the REST resource to be referenced. This field may be set instead of &#39;moid&#39; by clients. 1. If &#39;moid&#39; is set this field is ignored. 1. If &#39;selector&#39; is set and &#39;moid&#39; is empty/absent from the request, Intersight determines the Moid of the resource matching the filter expression and populates it in the MoRef that is part of the object instance being inserted/updated to fulfill the REST request. An error is returned if the filter matches zero or more than one REST resource. An example filter string is: Serial eq &#39;3AA8B7T11&#39;. | [optional] [readonly] 
**Link** | Pointer to **string** | A URL to an instance of the &#39;mo.MoRef&#39; class. | [optional] 

## Methods

### NewKubernetesNodeProfileRelationship

`func NewKubernetesNodeProfileRelationship(classId string, objectType string, ) *KubernetesNodeProfileRelationship`

NewKubernetesNodeProfileRelationship instantiates a new KubernetesNodeProfileRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeProfileRelationshipWithDefaults

`func NewKubernetesNodeProfileRelationshipWithDefaults() *KubernetesNodeProfileRelationship`

NewKubernetesNodeProfileRelationshipWithDefaults instantiates a new KubernetesNodeProfileRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNodeProfileRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNodeProfileRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNodeProfileRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNodeProfileRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNodeProfileRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNodeProfileRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMoid

`func (o *KubernetesNodeProfileRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *KubernetesNodeProfileRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *KubernetesNodeProfileRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *KubernetesNodeProfileRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetSelector

`func (o *KubernetesNodeProfileRelationship) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *KubernetesNodeProfileRelationship) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *KubernetesNodeProfileRelationship) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *KubernetesNodeProfileRelationship) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetLink

`func (o *KubernetesNodeProfileRelationship) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *KubernetesNodeProfileRelationship) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *KubernetesNodeProfileRelationship) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *KubernetesNodeProfileRelationship) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


