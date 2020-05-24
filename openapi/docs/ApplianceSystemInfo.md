# ApplianceSystemInfo

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

### NewApplianceSystemInfo

`func NewApplianceSystemInfo() *ApplianceSystemInfo`

NewApplianceSystemInfo instantiates a new ApplianceSystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceSystemInfoWithDefaults

`func NewApplianceSystemInfoWithDefaults() *ApplianceSystemInfo`

NewApplianceSystemInfoWithDefaults instantiates a new ApplianceSystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudConnStatus

`func (o *ApplianceSystemInfo) GetCloudConnStatus() string`

GetCloudConnStatus returns the CloudConnStatus field if non-nil, zero value otherwise.

### GetCloudConnStatusOk

`func (o *ApplianceSystemInfo) GetCloudConnStatusOk() (*string, bool)`

GetCloudConnStatusOk returns a tuple with the CloudConnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudConnStatus

`func (o *ApplianceSystemInfo) SetCloudConnStatus(v string)`

SetCloudConnStatus sets CloudConnStatus field to given value.

### HasCloudConnStatus

`func (o *ApplianceSystemInfo) HasCloudConnStatus() bool`

HasCloudConnStatus returns a boolean if a field has been set.

### GetDeploymentSize

`func (o *ApplianceSystemInfo) GetDeploymentSize() string`

GetDeploymentSize returns the DeploymentSize field if non-nil, zero value otherwise.

### GetDeploymentSizeOk

`func (o *ApplianceSystemInfo) GetDeploymentSizeOk() (*string, bool)`

GetDeploymentSizeOk returns a tuple with the DeploymentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSize

`func (o *ApplianceSystemInfo) SetDeploymentSize(v string)`

SetDeploymentSize sets DeploymentSize field to given value.

### HasDeploymentSize

`func (o *ApplianceSystemInfo) HasDeploymentSize() bool`

HasDeploymentSize returns a boolean if a field has been set.

### GetHostname

`func (o *ApplianceSystemInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApplianceSystemInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApplianceSystemInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApplianceSystemInfo) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetInitDone

`func (o *ApplianceSystemInfo) GetInitDone() bool`

GetInitDone returns the InitDone field if non-nil, zero value otherwise.

### GetInitDoneOk

`func (o *ApplianceSystemInfo) GetInitDoneOk() (*bool, bool)`

GetInitDoneOk returns a tuple with the InitDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitDone

`func (o *ApplianceSystemInfo) SetInitDone(v bool)`

SetInitDone sets InitDone field to given value.

### HasInitDone

`func (o *ApplianceSystemInfo) HasInitDone() bool`

HasInitDone returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *ApplianceSystemInfo) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ApplianceSystemInfo) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ApplianceSystemInfo) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ApplianceSystemInfo) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetSerialId

`func (o *ApplianceSystemInfo) GetSerialId() string`

GetSerialId returns the SerialId field if non-nil, zero value otherwise.

### GetSerialIdOk

`func (o *ApplianceSystemInfo) GetSerialIdOk() (*string, bool)`

GetSerialIdOk returns a tuple with the SerialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialId

`func (o *ApplianceSystemInfo) SetSerialId(v string)`

SetSerialId sets SerialId field to given value.

### HasSerialId

`func (o *ApplianceSystemInfo) HasSerialId() bool`

HasSerialId returns a boolean if a field has been set.

### GetTimeZone

`func (o *ApplianceSystemInfo) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ApplianceSystemInfo) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ApplianceSystemInfo) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *ApplianceSystemInfo) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetVersion

`func (o *ApplianceSystemInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplianceSystemInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplianceSystemInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplianceSystemInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


