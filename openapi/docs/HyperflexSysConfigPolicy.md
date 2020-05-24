# HyperflexSysConfigPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsDomainName** | Pointer to **string** | The DNS Search Domain Name. This setting applies to HyperFlex Data Platform 3.0 or later only. | [optional] 
**DnsServers** | Pointer to **[]string** |  | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**Timezone** | Pointer to **string** | The timezone of the HyperFlex cluster&#39;s system clock. | [optional] [default to "Pacific/Niue"]
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexSysConfigPolicy

`func NewHyperflexSysConfigPolicy() *HyperflexSysConfigPolicy`

NewHyperflexSysConfigPolicy instantiates a new HyperflexSysConfigPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSysConfigPolicyWithDefaults

`func NewHyperflexSysConfigPolicyWithDefaults() *HyperflexSysConfigPolicy`

NewHyperflexSysConfigPolicyWithDefaults instantiates a new HyperflexSysConfigPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsDomainName

`func (o *HyperflexSysConfigPolicy) GetDnsDomainName() string`

GetDnsDomainName returns the DnsDomainName field if non-nil, zero value otherwise.

### GetDnsDomainNameOk

`func (o *HyperflexSysConfigPolicy) GetDnsDomainNameOk() (*string, bool)`

GetDnsDomainNameOk returns a tuple with the DnsDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomainName

`func (o *HyperflexSysConfigPolicy) SetDnsDomainName(v string)`

SetDnsDomainName sets DnsDomainName field to given value.

### HasDnsDomainName

`func (o *HyperflexSysConfigPolicy) HasDnsDomainName() bool`

HasDnsDomainName returns a boolean if a field has been set.

### GetDnsServers

`func (o *HyperflexSysConfigPolicy) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *HyperflexSysConfigPolicy) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *HyperflexSysConfigPolicy) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *HyperflexSysConfigPolicy) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetNtpServers

`func (o *HyperflexSysConfigPolicy) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *HyperflexSysConfigPolicy) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *HyperflexSysConfigPolicy) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *HyperflexSysConfigPolicy) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetTimezone

`func (o *HyperflexSysConfigPolicy) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *HyperflexSysConfigPolicy) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *HyperflexSysConfigPolicy) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *HyperflexSysConfigPolicy) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexSysConfigPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexSysConfigPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexSysConfigPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexSysConfigPolicy) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### GetOrganization

`func (o *HyperflexSysConfigPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexSysConfigPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexSysConfigPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexSysConfigPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


