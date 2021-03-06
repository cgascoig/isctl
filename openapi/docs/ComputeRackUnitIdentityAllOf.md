# ComputeRackUnitIdentityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.RackUnitIdentity"]
**ObjectType** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.RackUnitIdentity"]
**AdapterSerial** | Pointer to **string** | Serial Identifier of an adapter participating in SWM. | [optional] 

## Methods

### NewComputeRackUnitIdentityAllOf

`func NewComputeRackUnitIdentityAllOf(classId string, objectType string, ) *ComputeRackUnitIdentityAllOf`

NewComputeRackUnitIdentityAllOf instantiates a new ComputeRackUnitIdentityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeRackUnitIdentityAllOfWithDefaults

`func NewComputeRackUnitIdentityAllOfWithDefaults() *ComputeRackUnitIdentityAllOf`

NewComputeRackUnitIdentityAllOfWithDefaults instantiates a new ComputeRackUnitIdentityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeRackUnitIdentityAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeRackUnitIdentityAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeRackUnitIdentityAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeRackUnitIdentityAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeRackUnitIdentityAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeRackUnitIdentityAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapterSerial

`func (o *ComputeRackUnitIdentityAllOf) GetAdapterSerial() string`

GetAdapterSerial returns the AdapterSerial field if non-nil, zero value otherwise.

### GetAdapterSerialOk

`func (o *ComputeRackUnitIdentityAllOf) GetAdapterSerialOk() (*string, bool)`

GetAdapterSerialOk returns a tuple with the AdapterSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterSerial

`func (o *ComputeRackUnitIdentityAllOf) SetAdapterSerial(v string)`

SetAdapterSerial sets AdapterSerial field to given value.

### HasAdapterSerial

`func (o *ComputeRackUnitIdentityAllOf) HasAdapterSerial() bool`

HasAdapterSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


