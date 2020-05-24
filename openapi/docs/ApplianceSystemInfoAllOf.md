# ApplianceSystemInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudConnStatus** | Pointer to **string** | Connection state of the Intersight Appliance to the Intersight. | [optional] [readonly] [default to ""]
**DeploymentSize** | Pointer to **string** | Current running deployment size for the Intersight Appliance cluster. Eg. small, medium, large etc. | [optional] [readonly] 
**Hostname** | Pointer to **string** | Publicly accessible FQDN or IP address of the Intersight Appliance. | [optional] [readonly] 
**InitDone** | Pointer to **bool** | Indicates that the setup initialization process has been completed. | [optional] [readonly] 
**OperationalStatus** | Pointer to **string** | Operational status of the Intersight Appliance cluster. | [optional] [readonly] [default to "Unknown"]
**SerialId** | Pointer to **string** | SerialId of the Intersight Appliance. SerialId is generated when the Intersight Appliance is setup. It is a unique UUID string, and serialId will not change for the life time of the Intersight Appliance. | [optional] [readonly] 
**TimeZone** | Pointer to **string** | Timezone of the Intersight Appliance. | [optional] [default to "Pacific/Niue"]
**Version** | Pointer to **string** | Current software version of the Intersight Appliance. | [optional] [readonly] 

## Methods

### NewApplianceSystemInfoAllOf

`func NewApplianceSystemInfoAllOf() *ApplianceSystemInfoAllOf`

NewApplianceSystemInfoAllOf instantiates a new ApplianceSystemInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceSystemInfoAllOfWithDefaults

`func NewApplianceSystemInfoAllOfWithDefaults() *ApplianceSystemInfoAllOf`

NewApplianceSystemInfoAllOfWithDefaults instantiates a new ApplianceSystemInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudConnStatus

`func (o *ApplianceSystemInfoAllOf) GetCloudConnStatus() string`

GetCloudConnStatus returns the CloudConnStatus field if non-nil, zero value otherwise.

### GetCloudConnStatusOk

`func (o *ApplianceSystemInfoAllOf) GetCloudConnStatusOk() (*string, bool)`

GetCloudConnStatusOk returns a tuple with the CloudConnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudConnStatus

`func (o *ApplianceSystemInfoAllOf) SetCloudConnStatus(v string)`

SetCloudConnStatus sets CloudConnStatus field to given value.

### HasCloudConnStatus

`func (o *ApplianceSystemInfoAllOf) HasCloudConnStatus() bool`

HasCloudConnStatus returns a boolean if a field has been set.

### GetDeploymentSize

`func (o *ApplianceSystemInfoAllOf) GetDeploymentSize() string`

GetDeploymentSize returns the DeploymentSize field if non-nil, zero value otherwise.

### GetDeploymentSizeOk

`func (o *ApplianceSystemInfoAllOf) GetDeploymentSizeOk() (*string, bool)`

GetDeploymentSizeOk returns a tuple with the DeploymentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSize

`func (o *ApplianceSystemInfoAllOf) SetDeploymentSize(v string)`

SetDeploymentSize sets DeploymentSize field to given value.

### HasDeploymentSize

`func (o *ApplianceSystemInfoAllOf) HasDeploymentSize() bool`

HasDeploymentSize returns a boolean if a field has been set.

### GetHostname

`func (o *ApplianceSystemInfoAllOf) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApplianceSystemInfoAllOf) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApplianceSystemInfoAllOf) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApplianceSystemInfoAllOf) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetInitDone

`func (o *ApplianceSystemInfoAllOf) GetInitDone() bool`

GetInitDone returns the InitDone field if non-nil, zero value otherwise.

### GetInitDoneOk

`func (o *ApplianceSystemInfoAllOf) GetInitDoneOk() (*bool, bool)`

GetInitDoneOk returns a tuple with the InitDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitDone

`func (o *ApplianceSystemInfoAllOf) SetInitDone(v bool)`

SetInitDone sets InitDone field to given value.

### HasInitDone

`func (o *ApplianceSystemInfoAllOf) HasInitDone() bool`

HasInitDone returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *ApplianceSystemInfoAllOf) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ApplianceSystemInfoAllOf) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ApplianceSystemInfoAllOf) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ApplianceSystemInfoAllOf) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetSerialId

`func (o *ApplianceSystemInfoAllOf) GetSerialId() string`

GetSerialId returns the SerialId field if non-nil, zero value otherwise.

### GetSerialIdOk

`func (o *ApplianceSystemInfoAllOf) GetSerialIdOk() (*string, bool)`

GetSerialIdOk returns a tuple with the SerialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialId

`func (o *ApplianceSystemInfoAllOf) SetSerialId(v string)`

SetSerialId sets SerialId field to given value.

### HasSerialId

`func (o *ApplianceSystemInfoAllOf) HasSerialId() bool`

HasSerialId returns a boolean if a field has been set.

### GetTimeZone

`func (o *ApplianceSystemInfoAllOf) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ApplianceSystemInfoAllOf) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ApplianceSystemInfoAllOf) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *ApplianceSystemInfoAllOf) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetVersion

`func (o *ApplianceSystemInfoAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplianceSystemInfoAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplianceSystemInfoAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplianceSystemInfoAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


