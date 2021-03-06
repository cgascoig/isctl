# VnicFcInterruptSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.FcInterruptSettings"]
**ObjectType** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.FcInterruptSettings"]
**Mode** | Pointer to **string** | The preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option. * &#x60;MSIx&#x60; - Message Signaled Interrupt (MSI) mechanism with the optional extension (MSIx). MSIx is the recommended and default option. * &#x60;MSI&#x60; - Message Signaled Interrupt (MSI) mechanism that treats messages as interrupts. * &#x60;INTx&#x60; - Line-based interrupt (INTx) mechanism similar to the one used in Legacy systems. | [optional] [default to "MSIx"]

## Methods

### NewVnicFcInterruptSettings

`func NewVnicFcInterruptSettings(classId string, objectType string, ) *VnicFcInterruptSettings`

NewVnicFcInterruptSettings instantiates a new VnicFcInterruptSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcInterruptSettingsWithDefaults

`func NewVnicFcInterruptSettingsWithDefaults() *VnicFcInterruptSettings`

NewVnicFcInterruptSettingsWithDefaults instantiates a new VnicFcInterruptSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicFcInterruptSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcInterruptSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcInterruptSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicFcInterruptSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcInterruptSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcInterruptSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMode

`func (o *VnicFcInterruptSettings) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VnicFcInterruptSettings) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VnicFcInterruptSettings) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VnicFcInterruptSettings) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


