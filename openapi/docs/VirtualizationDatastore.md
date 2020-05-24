# VirtualizationDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to [**VirtualizationStorageCapacity**](virtualization.StorageCapacity.md) |  | [optional] 
**HostCount** | Pointer to **int64** | Number of hosts attached to or supported-by this datastore. | [optional] 
**Identity** | Pointer to **string** | The internally generated identity of this datastore. This entity is not manipulated by users. It aids in uniquely identifying the datastore object. For VMware, this is a MOR (managed object reference). | [optional] 
**Name** | Pointer to **string** | Name of this datastore supplied by user. It is not the identity of the datastore. The name is subject to user manipulations. | [optional] 
**Type** | Pointer to **string** | A string indicating the type of the datastore (VMFS, NFS, etc). | [optional] [default to "Unknown"]
**VmCount** | Pointer to **int64** | Number of virtual machines relying on (using) this datastore. | [optional] 

## Methods

### NewVirtualizationDatastore

`func NewVirtualizationDatastore() *VirtualizationDatastore`

NewVirtualizationDatastore instantiates a new VirtualizationDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationDatastoreWithDefaults

`func NewVirtualizationDatastoreWithDefaults() *VirtualizationDatastore`

NewVirtualizationDatastoreWithDefaults instantiates a new VirtualizationDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *VirtualizationDatastore) GetCapacity() VirtualizationStorageCapacity`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationDatastore) GetCapacityOk() (*VirtualizationStorageCapacity, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationDatastore) SetCapacity(v VirtualizationStorageCapacity)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationDatastore) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetHostCount

`func (o *VirtualizationDatastore) GetHostCount() int64`

GetHostCount returns the HostCount field if non-nil, zero value otherwise.

### GetHostCountOk

`func (o *VirtualizationDatastore) GetHostCountOk() (*int64, bool)`

GetHostCountOk returns a tuple with the HostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCount

`func (o *VirtualizationDatastore) SetHostCount(v int64)`

SetHostCount sets HostCount field to given value.

### HasHostCount

`func (o *VirtualizationDatastore) HasHostCount() bool`

HasHostCount returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationDatastore) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationDatastore) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationDatastore) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationDatastore) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationDatastore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationDatastore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationDatastore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationDatastore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *VirtualizationDatastore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualizationDatastore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualizationDatastore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualizationDatastore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVmCount

`func (o *VirtualizationDatastore) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *VirtualizationDatastore) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *VirtualizationDatastore) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *VirtualizationDatastore) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


