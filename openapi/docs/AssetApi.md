# \AssetApi

All URIs are relative to *https://intersight.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssetDeviceClaim**](AssetApi.md#CreateAssetDeviceClaim) | **Post** /api/v1/asset/DeviceClaims | Create a &#39;asset.DeviceClaim&#39; resource.
[**CreateAssetTarget**](AssetApi.md#CreateAssetTarget) | **Post** /api/v1/asset/Targets | Create a &#39;asset.Target&#39; resource.
[**DeleteAssetDeviceClaim**](AssetApi.md#DeleteAssetDeviceClaim) | **Delete** /api/v1/asset/DeviceClaims/{Moid} | Delete a &#39;asset.DeviceClaim&#39; resource.
[**DeleteAssetDeviceRegistration**](AssetApi.md#DeleteAssetDeviceRegistration) | **Delete** /api/v1/asset/DeviceRegistrations/{Moid} | Deletes the resource representing the device connector. All associated REST resources will be deleted. In particular, inventory and operational data associated with this device will be deleted.
[**DeleteAssetTarget**](AssetApi.md#DeleteAssetTarget) | **Delete** /api/v1/asset/Targets/{Moid} | Delete a &#39;asset.Target&#39; resource.
[**GetAssetClusterMemberByMoid**](AssetApi.md#GetAssetClusterMemberByMoid) | **Get** /api/v1/asset/ClusterMembers/{Moid} | Read a &#39;asset.ClusterMember&#39; resource.
[**GetAssetClusterMemberList**](AssetApi.md#GetAssetClusterMemberList) | **Get** /api/v1/asset/ClusterMembers | Read a &#39;asset.ClusterMember&#39; resource.
[**GetAssetDeviceConfigurationByMoid**](AssetApi.md#GetAssetDeviceConfigurationByMoid) | **Get** /api/v1/asset/DeviceConfigurations/{Moid} | Read a &#39;asset.DeviceConfiguration&#39; resource.
[**GetAssetDeviceConfigurationList**](AssetApi.md#GetAssetDeviceConfigurationList) | **Get** /api/v1/asset/DeviceConfigurations | Read a &#39;asset.DeviceConfiguration&#39; resource.
[**GetAssetDeviceConnectorManagerByMoid**](AssetApi.md#GetAssetDeviceConnectorManagerByMoid) | **Get** /api/v1/asset/DeviceConnectorManagers/{Moid} | Read a &#39;asset.DeviceConnectorManager&#39; resource.
[**GetAssetDeviceConnectorManagerList**](AssetApi.md#GetAssetDeviceConnectorManagerList) | **Get** /api/v1/asset/DeviceConnectorManagers | Read a &#39;asset.DeviceConnectorManager&#39; resource.
[**GetAssetDeviceContractInformationByMoid**](AssetApi.md#GetAssetDeviceContractInformationByMoid) | **Get** /api/v1/asset/DeviceContractInformations/{Moid} | Read a &#39;asset.DeviceContractInformation&#39; resource.
[**GetAssetDeviceContractInformationList**](AssetApi.md#GetAssetDeviceContractInformationList) | **Get** /api/v1/asset/DeviceContractInformations | Read a &#39;asset.DeviceContractInformation&#39; resource.
[**GetAssetDeviceRegistrationByMoid**](AssetApi.md#GetAssetDeviceRegistrationByMoid) | **Get** /api/v1/asset/DeviceRegistrations/{Moid} | Read a &#39;asset.DeviceRegistration&#39; resource.
[**GetAssetDeviceRegistrationList**](AssetApi.md#GetAssetDeviceRegistrationList) | **Get** /api/v1/asset/DeviceRegistrations | Read a &#39;asset.DeviceRegistration&#39; resource.
[**GetAssetSubscriptionDeviceContractInformationByMoid**](AssetApi.md#GetAssetSubscriptionDeviceContractInformationByMoid) | **Get** /api/v1/asset/SubscriptionDeviceContractInformations/{Moid} | Read a &#39;asset.SubscriptionDeviceContractInformation&#39; resource.
[**GetAssetSubscriptionDeviceContractInformationList**](AssetApi.md#GetAssetSubscriptionDeviceContractInformationList) | **Get** /api/v1/asset/SubscriptionDeviceContractInformations | Read a &#39;asset.SubscriptionDeviceContractInformation&#39; resource.
[**GetAssetTargetByMoid**](AssetApi.md#GetAssetTargetByMoid) | **Get** /api/v1/asset/Targets/{Moid} | Read a &#39;asset.Target&#39; resource.
[**GetAssetTargetList**](AssetApi.md#GetAssetTargetList) | **Get** /api/v1/asset/Targets | Read a &#39;asset.Target&#39; resource.
[**PatchAssetDeviceConfiguration**](AssetApi.md#PatchAssetDeviceConfiguration) | **Patch** /api/v1/asset/DeviceConfigurations/{Moid} | Update a &#39;asset.DeviceConfiguration&#39; resource.
[**PatchAssetDeviceContractInformation**](AssetApi.md#PatchAssetDeviceContractInformation) | **Patch** /api/v1/asset/DeviceContractInformations/{Moid} | Update a &#39;asset.DeviceContractInformation&#39; resource.
[**PatchAssetDeviceRegistration**](AssetApi.md#PatchAssetDeviceRegistration) | **Patch** /api/v1/asset/DeviceRegistrations/{Moid} | Updates the resource representing the device connector. For example, this can be used to annotate the device connector resource with user-specified tags.
[**PatchAssetTarget**](AssetApi.md#PatchAssetTarget) | **Patch** /api/v1/asset/Targets/{Moid} | Update a &#39;asset.Target&#39; resource.
[**UpdateAssetDeviceConfiguration**](AssetApi.md#UpdateAssetDeviceConfiguration) | **Post** /api/v1/asset/DeviceConfigurations/{Moid} | Update a &#39;asset.DeviceConfiguration&#39; resource.
[**UpdateAssetDeviceContractInformation**](AssetApi.md#UpdateAssetDeviceContractInformation) | **Post** /api/v1/asset/DeviceContractInformations/{Moid} | Update a &#39;asset.DeviceContractInformation&#39; resource.
[**UpdateAssetDeviceRegistration**](AssetApi.md#UpdateAssetDeviceRegistration) | **Post** /api/v1/asset/DeviceRegistrations/{Moid} | Updates the resource representing the device connector. For example, this can be used to annotate the device connector resource with user-specified tags.
[**UpdateAssetTarget**](AssetApi.md#UpdateAssetTarget) | **Post** /api/v1/asset/Targets/{Moid} | Update a &#39;asset.Target&#39; resource.



## CreateAssetDeviceClaim

> AssetDeviceClaim CreateAssetDeviceClaim(ctx).AssetDeviceClaim(assetDeviceClaim).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'asset.DeviceClaim' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetDeviceClaim := openapiclient.asset.DeviceClaim{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{openapiclient.mo.Tag{Key: "Key_example", Value: "Value_example"}), VersionContext: openapiclient.mo.VersionContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", InterestedMos: []MoMoRef{openapiclient.mo.MoRef{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example"}), RefMo: openapiclient.mo.MoRef{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example"}, Timestamp: "TODO", Version: "Version_example", VersionType: "VersionType_example"}, Ancestors: []MoBaseMoRelationship{openapiclient.mo.BaseMo.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{openapiclient.mo.Tag{Key: "Key_example", Value: "Value_example"}), VersionContext: openapiclient.mo.VersionContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", InterestedMos: []MoMoRef{), RefMo: , Timestamp: "TODO", Version: "Version_example", VersionType: "VersionType_example"}, Ancestors: []MoBaseMoRelationship{openapiclient.mo.BaseMo.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }}), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }}), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceUpdates: []AssetConnectionControlMessage{openapiclient.asset.ConnectionControlMessage{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Account: "Account_example", ConnectorVersion: "ConnectorVersion_example", DeviceId: "DeviceId_example", DomainGroup: "DomainGroup_example", Evict: false, Leadership: "Leadership_example", NewIdentity: "NewIdentity_example", Partition: int64(123)}), SecurityToken: "SecurityToken_example", SerialNumber: "SerialNumber_example", Account: openapiclient.iam.Account.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Status: "Status_example", AppRegistrations: []IamAppRegistrationRelationship{openapiclient.iam.AppRegistration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientId: "ClientId_example", ClientName: "ClientName_example", ClientSecret: "ClientSecret_example", ClientType: "ClientType_example", Description: "Description_example", GrantTypes: []string{"GrantTypes_example"), RedirectUris: []string{"RedirectUris_example"), RenewClientSecret: false, ResponseTypes: []string{"ResponseTypes_example"), RevocationTimestamp: "TODO", Revoke: false, Account: openapiclient.iam.Account.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Status: "Status_example", AppRegistrations: []IamAppRegistrationRelationship{openapiclient.iam.AppRegistration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientId: "ClientId_example", ClientName: "ClientName_example", ClientSecret: "ClientSecret_example", ClientType: "ClientType_example", Description: "Description_example", GrantTypes: []string{"GrantTypes_example"), RedirectUris: []string{"RedirectUris_example"), RenewClientSecret: false, ResponseTypes: []string{"ResponseTypes_example"), RevocationTimestamp: "TODO", Revoke: false, Account: , OauthTokens: []IamOAuthTokenRelationship{openapiclient.iam.OAuthToken.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccessExpirationTime: "TODO", ClientId: "ClientId_example", ClientIpAddress: "ClientIpAddress_example", ClientName: "ClientName_example", ExpirationTime: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", TokenId: "TokenId_example", UserMeta: openapiclient.iam.ClientMeta{ClassId: "ClassId_example", ObjectType: "ObjectType_example", DeviceModel: "DeviceModel_example", UserAgent: "UserAgent_example"}, AppRegistration: , Permission: openapiclient.iam.Permission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , EndPointRoles: []IamEndPointRoleRelationship{openapiclient.iam.EndPointRole.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", RoleType: "RoleType_example", Type: "Type_example", Account: , EndPointPrivileges: []IamEndPointPrivilegeRelationship{openapiclient.iam.EndPointPrivilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", System: openapiclient.iam.System.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointPrivileges: []IamEndPointPrivilegeRelationship{openapiclient.iam.EndPointPrivilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", System: openapiclient.iam.System.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointPrivileges: []IamEndPointPrivilegeRelationship{), EndPointRoles: []IamEndPointRoleRelationship{openapiclient.iam.EndPointRole.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", RoleType: "RoleType_example", Type: "Type_example", Account: , EndPointPrivileges: []IamEndPointPrivilegeRelationship{), System: }), Idp: openapiclient.iam.Idp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", EnableSingleLogout: false, IdpEntityId: "IdpEntityId_example", Metadata: "Metadata_example", Name: "Name_example", Type: "Type_example", Account: , LdapPolicy: openapiclient.iam.LdapPolicy.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", BaseProperties: openapiclient.iam.LdapBaseProperties{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Attribute: "Attribute_example", BaseDn: "BaseDn_example", BindDn: "BindDn_example", BindMethod: "BindMethod_example", Domain: "Domain_example", EnableEncryption: false, EnableGroupAuthorization: false, Filter: "Filter_example", GroupAttribute: "GroupAttribute_example", IsPasswordSet: false, NestedGroupSearchDepth: int64(123), Password: "Password_example", Timeout: int64(123)}, DnsParameters: openapiclient.iam.LdapDnsParameters{ClassId: "ClassId_example", ObjectType: "ObjectType_example", SearchDomain: "SearchDomain_example", SearchForest: "SearchForest_example", Source: "Source_example"}, EnableDns: false, Enabled: false, UserSearchPrecedence: "UserSearchPrecedence_example", Var0Idp: openapiclient.iam.Idp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", EnableSingleLogout: false, IdpEntityId: "IdpEntityId_example", Metadata: "Metadata_example", Name: "Name_example", Type: "Type_example", Account: , LdapPolicy: openapiclient.iam.LdapPolicy.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", BaseProperties: openapiclient.iam.LdapBaseProperties{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Attribute: "Attribute_example", BaseDn: "BaseDn_example", BindDn: "BindDn_example", BindMethod: "BindMethod_example", Domain: "Domain_example", EnableEncryption: false, EnableGroupAuthorization: false, Filter: "Filter_example", GroupAttribute: "GroupAttribute_example", IsPasswordSet: false, NestedGroupSearchDepth: int64(123), Password: "Password_example", Timeout: int64(123)}, DnsParameters: openapiclient.iam.LdapDnsParameters{ClassId: "ClassId_example", ObjectType: "ObjectType_example", SearchDomain: "SearchDomain_example", SearchForest: "SearchForest_example", Source: "Source_example"}, EnableDns: false, Enabled: false, UserSearchPrecedence: "UserSearchPrecedence_example", Var0Idp: , ApplianceAccount: , Groups: []IamLdapGroupRelationship{openapiclient.iam.LdapGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Domain: "Domain_example", Name: "Name_example", EndPointRole: []IamEndPointRoleRelationship{), LdapPolicy: }), Organization: openapiclient.organization.Organization.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , ResourceGroups: []ResourceGroupRelationship{openapiclient.resource.Group.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", PerTypeCombinedSelector: []ResourcePerTypeCombinedSelector{openapiclient.resource.PerTypeCombinedSelector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", CombinedSelector: "CombinedSelector_example", EmptyFilter: false, SelectorObjectType: "SelectorObjectType_example"}), Qualifier: "Qualifier_example", Selectors: []ResourceSelector{openapiclient.resource.Selector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Selector: "Selector_example"}), Account: , Organizations: []OrganizationOrganizationRelationship{openapiclient.organization.Organization.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , ResourceGroups: []ResourceGroupRelationship{openapiclient.resource.Group.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", PerTypeCombinedSelector: []ResourcePerTypeCombinedSelector{openapiclient.resource.PerTypeCombinedSelector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", CombinedSelector: "CombinedSelector_example", EmptyFilter: false, SelectorObjectType: "SelectorObjectType_example"}), Qualifier: "Qualifier_example", Selectors: []ResourceSelector{openapiclient.resource.Selector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Selector: "Selector_example"}), Account: , Organizations: []OrganizationOrganizationRelationship{)})})})}, Profiles: []PolicyAbstractConfigProfileRelationship{openapiclient.policy.AbstractConfigProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: openapiclient.policy.AbstractProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: openapiclient.policy.AbstractProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: }}, Action: "Action_example", ConfigContext: openapiclient.policy.ConfigContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", ConfigState: "ConfigState_example", ControlAction: "ControlAction_example", ErrorState: "ErrorState_example", OperState: "OperState_example"}}), Providers: []IamLdapProviderRelationship{openapiclient.iam.LdapProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Port: int64(123), Server: "Server_example", LdapPolicy: })}, System: , UserPreferences: []IamUserPreferenceRelationship{openapiclient.iam.UserPreference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Preference: 123, UserUniqueIdentifier: "UserUniqueIdentifier_example", Idp: , IdpReference: openapiclient.iam.IdpReference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", MultiFactorAuthentication: false, Name: "Name_example", Account: , Idp: , UserPreferences: []IamUserPreferenceRelationship{openapiclient.iam.UserPreference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Preference: 123, UserUniqueIdentifier: "UserUniqueIdentifier_example", Idp: , IdpReference: openapiclient.iam.IdpReference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", MultiFactorAuthentication: false, Name: "Name_example", Account: , Idp: , UserPreferences: []IamUserPreferenceRelationship{), Usergroups: []IamUserGroupRelationship{openapiclient.iam.UserGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Idp: , Idpreference: , Permissions: []IamPermissionRelationship{openapiclient.iam.Permission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , EndPointRoles: []IamEndPointRoleRelationship{), PrivilegeSets: []IamPrivilegeSetRelationship{openapiclient.iam.PrivilegeSet.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , AssociatedPrivilegeSets: []IamPrivilegeSetRelationship{openapiclient.iam.PrivilegeSet.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , AssociatedPrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{openapiclient.iam.Privilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HostnamePrefix: "HostnamePrefix_example", Method: "Method_example", Name: "Name_example", RestPath: "RestPath_example", UrlPrefix: "UrlPrefix_example", Account: , System: }), System: }), Privileges: []IamPrivilegeRelationship{openapiclient.iam.Privilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HostnamePrefix: "HostnamePrefix_example", Method: "Method_example", Name: "Name_example", RestPath: "RestPath_example", UrlPrefix: "UrlPrefix_example", Account: , System: }), System: }), ResourceRoles: []IamResourceRolesRelationship{openapiclient.iam.ResourceRoles.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointRoles: []IamEndPointRoleRelationship{), Permission: , PrivilegeSets: []IamPrivilegeSetRelationship{), Resource: , Roles: []IamRoleRelationship{openapiclient.iam.Role.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , PrivilegeSets: []IamPrivilegeSetRelationship{), System: })}), Roles: []IamRoleRelationship{openapiclient.iam.Role.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , PrivilegeSets: []IamPrivilegeSetRelationship{), System: }), SessionLimits: openapiclient.iam.SessionLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, IdleTimeOut: int64(123), MaximumLimit: int64(123), PerUserLimit: int64(123), SessionTimeOut: int64(123), Account: , Permission: }, UserGroups: []IamUserGroupRelationship{openapiclient.iam.UserGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Idp: , Idpreference: , Permissions: []IamPermissionRelationship{), Qualifier: openapiclient.iam.Qualifier.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Value: []string{"Value_example"), Usergroup: }, Users: []IamUserRelationship{openapiclient.iam.User.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientIpAddress: "ClientIpAddress_example", Email: "Email_example", FirstName: "FirstName_example", LastLoginTime: "TODO", LastName: "LastName_example", Name: "Name_example", UserIdOrEmail: "UserIdOrEmail_example", UserType: "UserType_example", UserUniqueIdentifier: "UserUniqueIdentifier_example", ApiKeys: []IamApiKeyRelationship{openapiclient.iam.ApiKey.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HashAlgorithm: "HashAlgorithm_example", KeySpec: openapiclient.pkix.KeyGenerationSpec{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Name: "Name_example"}, PrivateKey: "PrivateKey_example", Purpose: "Purpose_example", SigningAlgorithm: "SigningAlgorithm_example", Permission: , User: openapiclient.iam.User.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientIpAddress: "ClientIpAddress_example", Email: "Email_example", FirstName: "FirstName_example", LastLoginTime: "TODO", LastName: "LastName_example", Name: "Name_example", UserIdOrEmail: "UserIdOrEmail_example", UserType: "UserType_example", UserUniqueIdentifier: "UserUniqueIdentifier_example", ApiKeys: []IamApiKeyRelationship{openapiclient.iam.ApiKey.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HashAlgorithm: "HashAlgorithm_example", KeySpec: openapiclient.pkix.KeyGenerationSpec{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Name: "Name_example"}, PrivateKey: "PrivateKey_example", Purpose: "Purpose_example", SigningAlgorithm: "SigningAlgorithm_example", Permission: , User: }), AppRegistrations: []IamAppRegistrationRelationship{), Idp: , Idpreference: , LocalUserPassword: openapiclient.iam.LocalUserPassword.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, CurrentPassword: "CurrentPassword_example", IsCurrentPasswordSet: false, IsNewPasswordSet: false, NewPassword: "NewPassword_example", Password: 123, User: }, OauthTokens: []IamOAuthTokenRelationship{openapiclient.iam.OAuthToken.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccessExpirationTime: "TODO", ClientId: "ClientId_example", ClientIpAddress: "ClientIpAddress_example", ClientName: "ClientName_example", ExpirationTime: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", TokenId: "TokenId_example", UserMeta: openapiclient.iam.ClientMeta{ClassId: "ClassId_example", ObjectType: "ObjectType_example", DeviceModel: "DeviceModel_example", UserAgent: "UserAgent_example"}, AppRegistration: , Permission: , User: }), Permissions: []IamPermissionRelationship{), Sessions: []IamSessionRelationship{openapiclient.iam.Session.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccountPermissions: []IamAccountPermissions{openapiclient.iam.AccountPermissions{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountIdentifier: "AccountIdentifier_example", AccountName: "AccountName_example", AccountStatus: "AccountStatus_example", Permissions: []IamPermissionReference{openapiclient.iam.PermissionReference{ClassId: "ClassId_example", ObjectType: "ObjectType_example", PermissionIdentifier: "PermissionIdentifier_example", PermissionName: "PermissionName_example"})}), ClientIpAddress: "ClientIpAddress_example", Expiration: "TODO", IdleTimeExpiration: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", Permission: , User: })}}), AppRegistrations: []IamAppRegistrationRelationship{), Idp: , Idpreference: , LocalUserPassword: openapiclient.iam.LocalUserPassword.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, CurrentPassword: "CurrentPassword_example", IsCurrentPasswordSet: false, IsNewPasswordSet: false, NewPassword: "NewPassword_example", Password: 123, User: }, OauthTokens: []IamOAuthTokenRelationship{), Permissions: []IamPermissionRelationship{), Sessions: []IamSessionRelationship{openapiclient.iam.Session.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccountPermissions: []IamAccountPermissions{openapiclient.iam.AccountPermissions{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountIdentifier: "AccountIdentifier_example", AccountName: "AccountName_example", AccountStatus: "AccountStatus_example", Permissions: []IamPermissionReference{openapiclient.iam.PermissionReference{ClassId: "ClassId_example", ObjectType: "ObjectType_example", PermissionIdentifier: "PermissionIdentifier_example", PermissionName: "PermissionName_example"})}), ClientIpAddress: "ClientIpAddress_example", Expiration: "TODO", IdleTimeExpiration: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", Permission: , User: })})}), Users: []IamUserRelationship{)}), Qualifier: openapiclient.iam.Qualifier.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Value: []string{"Value_example"), Usergroup: }, Users: []IamUserRelationship{)}), Users: []IamUserRelationship{)}}), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}}), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, ApplianceAccount: , Groups: []IamLdapGroupRelationship{openapiclient.iam.LdapGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Domain: "Domain_example", Name: "Name_example", EndPointRole: []IamEndPointRoleRelationship{), LdapPolicy: }), Organization: , Profiles: []PolicyAbstractConfigProfileRelationship{openapiclient.policy.AbstractConfigProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: , Action: "Action_example", ConfigContext: openapiclient.policy.ConfigContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", ConfigState: "ConfigState_example", ControlAction: "ControlAction_example", ErrorState: "ErrorState_example", OperState: "OperState_example"}}), Providers: []IamLdapProviderRelationship{openapiclient.iam.LdapProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Port: int64(123), Server: "Server_example", LdapPolicy: })}, System: , UserPreferences: []IamUserPreferenceRelationship{), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), Roles: []IamRoleRelationship{), ServiceProvider: openapiclient.iam.ServiceProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EntityId: "EntityId_example", Metadata: "Metadata_example", Name: "Name_example", System: }}}), EndPointRoles: []IamEndPointRoleRelationship{), Idp: , PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), Roles: []IamRoleRelationship{), ServiceProvider: openapiclient.iam.ServiceProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EntityId: "EntityId_example", Metadata: "Metadata_example", Name: "Name_example", System: }}}), System: }), PrivilegeSets: []IamPrivilegeSetRelationship{), ResourceRoles: []IamResourceRolesRelationship{openapiclient.iam.ResourceRoles.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointRoles: []IamEndPointRoleRelationship{), Permission: , PrivilegeSets: []IamPrivilegeSetRelationship{), Resource: , Roles: []IamRoleRelationship{)}), Roles: []IamRoleRelationship{), SessionLimits: openapiclient.iam.SessionLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, IdleTimeOut: int64(123), MaximumLimit: int64(123), PerUserLimit: int64(123), SessionTimeOut: int64(123), Account: , Permission: }, UserGroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, User: }), Permission: , Roles: []IamRoleRelationship{), User: }), DomainGroups: []IamDomainGroupRelationship{openapiclient.iam.DomainGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Partition1: int64(123), Partition2: int64(123), Partition3: int64(123), PartitionKey: "PartitionKey_example", Usage: int64(123), Account: }), EndPointRoles: []IamEndPointRoleRelationship{), Idpreferences: []IamIdpReferenceRelationship{), Idps: []IamIdpRelationship{), Permissions: []IamPermissionRelationship{), PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), ResourceLimits: openapiclient.iam.ResourceLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PerAccountUserLimit: int64(123), Account: }, Roles: []IamRoleRelationship{), SecurityHolder: openapiclient.iam.SecurityHolder.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Account: , IpRulesConfiguration: openapiclient.iam.IpAccessManagement.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Enable: false, LastRecoveryTime: "TODO", Holder: openapiclient.iam.SecurityHolder.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Account: , IpRulesConfiguration: openapiclient.iam.IpAccessManagement.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Enable: false, LastRecoveryTime: "TODO", Holder: , IpAddresses: []IamIpAddressRelationship{openapiclient.iam.IpAddress.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Address: "Address_example", Description: "Description_example", IpAccessManagement: })}, ResourcePermissions: []IamResourcePermissionRelationship{openapiclient.iam.ResourcePermission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PermissionRoles: []IamPermissionToRoles{openapiclient.iam.PermissionToRoles{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Permission: openapiclient.cmrf.CmRf{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example"}, Roles: []CmrfCmRf{openapiclient.cmrf.CmRf{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example"})}), TargetApp: "TargetApp_example", Holder: , Resource: })}, IpAddresses: []IamIpAddressRelationship{openapiclient.iam.IpAddress.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Address: "Address_example", Description: "Description_example", IpAccessManagement: })}, ResourcePermissions: []IamResourcePermissionRelationship{openapiclient.iam.ResourcePermission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PermissionRoles: []IamPermissionToRoles{openapiclient.iam.PermissionToRoles{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Permission: , Roles: []CmrfCmRf{)}), TargetApp: "TargetApp_example", Holder: , Resource: })}, SessionLimits: }, OauthTokens: []IamOAuthTokenRelationship{), Permission: , Roles: []IamRoleRelationship{), User: }), DomainGroups: []IamDomainGroupRelationship{openapiclient.iam.DomainGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Partition1: int64(123), Partition2: int64(123), Partition3: int64(123), PartitionKey: "PartitionKey_example", Usage: int64(123), Account: }), EndPointRoles: []IamEndPointRoleRelationship{), Idpreferences: []IamIdpReferenceRelationship{), Idps: []IamIdpRelationship{), Permissions: []IamPermissionRelationship{), PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), ResourceLimits: openapiclient.iam.ResourceLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PerAccountUserLimit: int64(123), Account: }, Roles: []IamRoleRelationship{), SecurityHolder: , SessionLimits: }, Device: openapiclient.asset.DeviceRegistration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ApiVersion: int64(123), AppPartitionNumber: int64(123), ConnectionId: "ConnectionId_example", ConnectionReason: "ConnectionReason_example", ConnectionStatus: "ConnectionStatus_example", ConnectionStatusLastChangeTime: "TODO", ConnectorVersion: "ConnectorVersion_example", DeviceExternalIpAddress: "DeviceExternalIpAddress_example", ProxyApp: "ProxyApp_example", AccessKeyId: "AccessKeyId_example", ClaimedByUserName: "ClaimedByUserName_example", ClaimedTime: "TODO", DeviceHostname: []string{"DeviceHostname_example"), DeviceIpAddress: []string{"DeviceIpAddress_example"), ExecutionMode: "ExecutionMode_example", ParentSignature: openapiclient.asset.ParentConnectionSignature{ClassId: "ClassId_example", ObjectType: "ObjectType_example", DeviceId: "DeviceId_example", NodeId: "NodeId_example", Signature: 123, TimeStamp: "TODO"}, Pid: []string{"Pid_example"), PlatformType: "PlatformType_example", PublicAccessKey: "PublicAccessKey_example", ReadOnly: false, Serial: []string{"Serial_example"), Vendor: "Vendor_example", Account: , ClaimedByUser: , ClusterMembers: []AssetClusterMemberRelationship{openapiclient.asset.ClusterMember.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ApiVersion: int64(123), AppPartitionNumber: int64(123), ConnectionId: "ConnectionId_example", ConnectionReason: "ConnectionReason_example", ConnectionStatus: "ConnectionStatus_example", ConnectionStatusLastChangeTime: "TODO", ConnectorVersion: "ConnectorVersion_example", DeviceExternalIpAddress: "DeviceExternalIpAddress_example", ProxyApp: "ProxyApp_example", Leadership: "Leadership_example", LockedLeader: false, MemberIdentity: "MemberIdentity_example", ParentClusterMemberIdentity: "ParentClusterMemberIdentity_example", Sudi: openapiclient.asset.SudiInfo{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Pid: "Pid_example", SerialNumber: "SerialNumber_example", Signature: "Signature_example", Status: "Status_example", SudiCertificate: openapiclient.x509.Certificate{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Issuer: openapiclient.pkix.DistinguishedName{ClassId: "ClassId_example", ObjectType: "ObjectType_example", CommonName: "CommonName_example", Country: []string{"Country_example"), Locality: []string{"Locality_example"), Organization: []string{"Organization_example"), OrganizationalUnit: []string{"OrganizationalUnit_example"), State: []string{"State_example")}, NotAfter: "TODO", NotBefore: "TODO", PemCertificate: "PemCertificate_example", Sha256Fingerprint: "Sha256Fingerprint_example", SignatureAlgorithm: "SignatureAlgorithm_example", Subject: openapiclient.pkix.DistinguishedName{ClassId: "ClassId_example", ObjectType: "ObjectType_example", CommonName: "CommonName_example", Country: []string{"Country_example"), Locality: []string{"Locality_example"), Organization: []string{"Organization_example"), OrganizationalUnit: []string{"OrganizationalUnit_example"), State: []string{"State_example")}}}, Device: openapiclient.asset.DeviceRegistration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ApiVersion: int64(123), AppPartitionNumber: int64(123), ConnectionId: "ConnectionId_example", ConnectionReason: "ConnectionReason_example", ConnectionStatus: "ConnectionStatus_example", ConnectionStatusLastChangeTime: "TODO", ConnectorVersion: "ConnectorVersion_example", DeviceExternalIpAddress: "DeviceExternalIpAddress_example", ProxyApp: "ProxyApp_example", AccessKeyId: "AccessKeyId_example", ClaimedByUserName: "ClaimedByUserName_example", ClaimedTime: "TODO", DeviceHostname: []string{"DeviceHostname_example"), DeviceIpAddress: []string{"DeviceIpAddress_example"), ExecutionMode: "ExecutionMode_example", ParentSignature: openapiclient.asset.ParentConnectionSignature{ClassId: "ClassId_example", ObjectType: "ObjectType_example", DeviceId: "DeviceId_example", NodeId: "NodeId_example", Signature: 123, TimeStamp: "TODO"}, Pid: []string{"Pid_example"), PlatformType: "PlatformType_example", PublicAccessKey: "PublicAccessKey_example", ReadOnly: false, Serial: []string{"Serial_example"), Vendor: "Vendor_example", Account: , ClaimedByUser: , ClusterMembers: []AssetClusterMemberRelationship{openapiclient.asset.ClusterMember.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ApiVersion: int64(123), AppPartitionNumber: int64(123), ConnectionId: "ConnectionId_example", ConnectionReason: "ConnectionReason_example", ConnectionStatus: "ConnectionStatus_example", ConnectionStatusLastChangeTime: "TODO", ConnectorVersion: "ConnectorVersion_example", DeviceExternalIpAddress: "DeviceExternalIpAddress_example", ProxyApp: "ProxyApp_example", Leadership: "Leadership_example", LockedLeader: false, MemberIdentity: "MemberIdentity_example", ParentClusterMemberIdentity: "ParentClusterMemberIdentity_example", Sudi: openapiclient.asset.SudiInfo{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Pid: "Pid_example", SerialNumber: "SerialNumber_example", Signature: "Signature_example", Status: "Status_example", SudiCertificate: openapiclient.x509.Certificate{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Issuer: , NotAfter: "TODO", NotBefore: "TODO", PemCertificate: "PemCertificate_example", Sha256Fingerprint: "Sha256Fingerprint_example", SignatureAlgorithm: "SignatureAlgorithm_example", Subject: }}, Device: }), DeviceClaim: openapiclient.asset.DeviceClaim.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceUpdates: []AssetConnectionControlMessage{openapiclient.asset.ConnectionControlMessage{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Account: "Account_example", ConnectorVersion: "ConnectorVersion_example", DeviceId: "DeviceId_example", DomainGroup: "DomainGroup_example", Evict: false, Leadership: "Leadership_example", NewIdentity: "NewIdentity_example", Partition: int64(123)}), SecurityToken: "SecurityToken_example", SerialNumber: "SerialNumber_example", Account: , Device: }, DeviceConfiguration: openapiclient.asset.DeviceConfiguration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, LocalConfigurationLocked: false, LogLevel: "LogLevel_example", Device: }, DomainGroup: , ParentConnection: }}), DeviceClaim: openapiclient.asset.DeviceClaim.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceUpdates: []AssetConnectionControlMessage{), SecurityToken: "SecurityToken_example", SerialNumber: "SerialNumber_example", Account: , Device: }, DeviceConfiguration: openapiclient.asset.DeviceConfiguration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, LocalConfigurationLocked: false, LogLevel: "LogLevel_example", Device: }, DomainGroup: , ParentConnection: }} // AssetDeviceClaim | The 'asset.DeviceClaim' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.CreateAssetDeviceClaim(context.Background(), assetDeviceClaim).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.CreateAssetDeviceClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssetDeviceClaim`: AssetDeviceClaim
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.CreateAssetDeviceClaim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetDeviceClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetDeviceClaim** | [**AssetDeviceClaim**](AssetDeviceClaim.md) | The &#39;asset.DeviceClaim&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**AssetDeviceClaim**](asset.DeviceClaim.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetTarget

> AssetTarget CreateAssetTarget(ctx).AssetTarget(assetTarget).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'asset.Target' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetTarget := openapiclient.asset.Target{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClaimedByUserName: "ClaimedByUserName_example", Connections: []AssetConnection{openapiclient.asset.Connection{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Credential: openapiclient.asset.Credential{ClassId: "ClassId_example", ObjectType: "ObjectType_example"}}), ConnectorVersion: "ConnectorVersion_example", ExternalIpAddress: "ExternalIpAddress_example", IpAddress: []string{"IpAddress_example"), Name: "Name_example", ProductId: []string{"ProductId_example"), ReadOnly: false, Services: []AssetService{openapiclient.asset.Service{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Options: openapiclient.asset.ServiceOptions{ClassId: "ClassId_example", ObjectType: "ObjectType_example"}}), Status: "Status_example", StatusErrorReason: "StatusErrorReason_example", TargetId: []string{"TargetId_example"), TargetType: "TargetType_example", Account: , Assist: openapiclient.asset.Target.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClaimedByUserName: "ClaimedByUserName_example", Connections: []AssetConnection{openapiclient.asset.Connection{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Credential: openapiclient.asset.Credential{ClassId: "ClassId_example", ObjectType: "ObjectType_example"}}), ConnectorVersion: "ConnectorVersion_example", ExternalIpAddress: "ExternalIpAddress_example", IpAddress: []string{"IpAddress_example"), Name: "Name_example", ProductId: []string{"ProductId_example"), ReadOnly: false, Services: []AssetService{openapiclient.asset.Service{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Options: openapiclient.asset.ServiceOptions{ClassId: "ClassId_example", ObjectType: "ObjectType_example"}}), Status: "Status_example", StatusErrorReason: "StatusErrorReason_example", TargetId: []string{"TargetId_example"), TargetType: "TargetType_example", Account: , Assist: openapiclient.asset.Target.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClaimedByUserName: "ClaimedByUserName_example", Connections: []AssetConnection{), ConnectorVersion: "ConnectorVersion_example", ExternalIpAddress: "ExternalIpAddress_example", IpAddress: []string{"IpAddress_example"), Name: "Name_example", ProductId: []string{"ProductId_example"), ReadOnly: false, Services: []AssetService{), Status: "Status_example", StatusErrorReason: "StatusErrorReason_example", TargetId: []string{"TargetId_example"), TargetType: "TargetType_example", Account: , Assist: , RegisteredDevice: }, RegisteredDevice: }, RegisteredDevice: } // AssetTarget | The 'asset.Target' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.CreateAssetTarget(context.Background(), assetTarget).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.CreateAssetTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssetTarget`: AssetTarget
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.CreateAssetTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetTarget** | [**AssetTarget**](AssetTarget.md) | The &#39;asset.Target&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**AssetTarget**](asset.Target.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssetDeviceClaim

> DeleteAssetDeviceClaim(ctx, moid).Execute()

Delete a 'asset.DeviceClaim' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.DeleteAssetDeviceClaim(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.DeleteAssetDeviceClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssetDeviceClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssetDeviceRegistration

> DeleteAssetDeviceRegistration(ctx, moid).Execute()

Deletes the resource representing the device connector. All associated REST resources will be deleted. In particular, inventory and operational data associated with this device will be deleted.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.DeleteAssetDeviceRegistration(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.DeleteAssetDeviceRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssetDeviceRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssetTarget

> DeleteAssetTarget(ctx, moid).Execute()

Delete a 'asset.Target' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.DeleteAssetTarget(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.DeleteAssetTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssetTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetClusterMemberByMoid

> AssetClusterMember GetAssetClusterMemberByMoid(ctx, moid).Execute()

Read a 'asset.ClusterMember' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.GetAssetClusterMemberByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetClusterMemberByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetClusterMemberByMoid`: AssetClusterMember
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetClusterMemberByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetClusterMemberByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetClusterMember**](asset.ClusterMember.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetClusterMemberList

> AssetClusterMemberResponse GetAssetClusterMemberList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'asset.ClusterMember' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.GetAssetClusterMemberList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetClusterMemberList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetClusterMemberList`: AssetClusterMemberResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetClusterMemberList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetClusterMemberListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**AssetClusterMemberResponse**](asset.ClusterMember.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetDeviceConfigurationByMoid

> AssetDeviceConfiguration GetAssetDeviceConfigurationByMoid(ctx, moid).Execute()

Read a 'asset.DeviceConfiguration' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.GetAssetDeviceConfigurationByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetDeviceConfigurationByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetDeviceConfigurationByMoid`: AssetDeviceConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetDeviceConfigurationByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDeviceConfigurationByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetDeviceConfiguration**](asset.DeviceConfiguration.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetDeviceConfigurationList

> AssetDeviceConfigurationResponse GetAssetDeviceConfigurationList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'asset.DeviceConfiguration' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.GetAssetDeviceConfigurationList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetDeviceConfigurationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetDeviceConfigurationList`: AssetDeviceConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetDeviceConfigurationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDeviceConfigurationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**AssetDeviceConfigurationResponse**](asset.DeviceConfiguration.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetDeviceConnectorManagerByMoid

> AssetDeviceConnectorManager GetAssetDeviceConnectorManagerByMoid(ctx, moid).Execute()

Read a 'asset.DeviceConnectorManager' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.GetAssetDeviceConnectorManagerByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetDeviceConnectorManagerByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetDeviceConnectorManagerByMoid`: AssetDeviceConnectorManager
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetDeviceConnectorManagerByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDeviceConnectorManagerByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetDeviceConnectorManager**](asset.DeviceConnectorManager.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetDeviceConnectorManagerList

> AssetDeviceConnectorManagerResponse GetAssetDeviceConnectorManagerList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'asset.DeviceConnectorManager' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.GetAssetDeviceConnectorManagerList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetDeviceConnectorManagerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetDeviceConnectorManagerList`: AssetDeviceConnectorManagerResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetDeviceConnectorManagerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDeviceConnectorManagerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**AssetDeviceConnectorManagerResponse**](asset.DeviceConnectorManager.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetDeviceContractInformationByMoid

> AssetDeviceContractInformation GetAssetDeviceContractInformationByMoid(ctx, moid).Execute()

Read a 'asset.DeviceContractInformation' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.GetAssetDeviceContractInformationByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetDeviceContractInformationByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetDeviceContractInformationByMoid`: AssetDeviceContractInformation
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetDeviceContractInformationByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDeviceContractInformationByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetDeviceContractInformation**](asset.DeviceContractInformation.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetDeviceContractInformationList

> AssetDeviceContractInformationResponse GetAssetDeviceContractInformationList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'asset.DeviceContractInformation' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.GetAssetDeviceContractInformationList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetDeviceContractInformationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetDeviceContractInformationList`: AssetDeviceContractInformationResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetDeviceContractInformationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDeviceContractInformationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**AssetDeviceContractInformationResponse**](asset.DeviceContractInformation.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetDeviceRegistrationByMoid

> AssetDeviceRegistration GetAssetDeviceRegistrationByMoid(ctx, moid).Execute()

Read a 'asset.DeviceRegistration' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.GetAssetDeviceRegistrationByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetDeviceRegistrationByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetDeviceRegistrationByMoid`: AssetDeviceRegistration
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetDeviceRegistrationByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDeviceRegistrationByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetDeviceRegistration**](asset.DeviceRegistration.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetDeviceRegistrationList

> AssetDeviceRegistrationResponse GetAssetDeviceRegistrationList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'asset.DeviceRegistration' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.GetAssetDeviceRegistrationList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetDeviceRegistrationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetDeviceRegistrationList`: AssetDeviceRegistrationResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetDeviceRegistrationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDeviceRegistrationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**AssetDeviceRegistrationResponse**](asset.DeviceRegistration.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetSubscriptionDeviceContractInformationByMoid

> AssetSubscriptionDeviceContractInformation GetAssetSubscriptionDeviceContractInformationByMoid(ctx, moid).Execute()

Read a 'asset.SubscriptionDeviceContractInformation' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.GetAssetSubscriptionDeviceContractInformationByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetSubscriptionDeviceContractInformationByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetSubscriptionDeviceContractInformationByMoid`: AssetSubscriptionDeviceContractInformation
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetSubscriptionDeviceContractInformationByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetSubscriptionDeviceContractInformationByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetSubscriptionDeviceContractInformation**](asset.SubscriptionDeviceContractInformation.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetSubscriptionDeviceContractInformationList

> AssetSubscriptionDeviceContractInformationResponse GetAssetSubscriptionDeviceContractInformationList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'asset.SubscriptionDeviceContractInformation' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.GetAssetSubscriptionDeviceContractInformationList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetSubscriptionDeviceContractInformationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetSubscriptionDeviceContractInformationList`: AssetSubscriptionDeviceContractInformationResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetSubscriptionDeviceContractInformationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetSubscriptionDeviceContractInformationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**AssetSubscriptionDeviceContractInformationResponse**](asset.SubscriptionDeviceContractInformation.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetTargetByMoid

> AssetTarget GetAssetTargetByMoid(ctx, moid).Execute()

Read a 'asset.Target' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.GetAssetTargetByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetTargetByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetTargetByMoid`: AssetTarget
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetTargetByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetTargetByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetTarget**](asset.Target.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetTargetList

> AssetTargetResponse GetAssetTargetList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'asset.Target' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.GetAssetTargetList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetTargetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetTargetList`: AssetTargetResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetTargetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetTargetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**AssetTargetResponse**](asset.Target.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAssetDeviceConfiguration

> AssetDeviceConfiguration PatchAssetDeviceConfiguration(ctx, moid).AssetDeviceConfiguration(assetDeviceConfiguration).IfMatch(ifMatch).Execute()

Update a 'asset.DeviceConfiguration' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    assetDeviceConfiguration := openapiclient.asset.DeviceConfiguration{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, LocalConfigurationLocked: false, LogLevel: "LogLevel_example", Device: } // AssetDeviceConfiguration | The 'asset.DeviceConfiguration' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.PatchAssetDeviceConfiguration(context.Background(), moid, assetDeviceConfiguration).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.PatchAssetDeviceConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAssetDeviceConfiguration`: AssetDeviceConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.PatchAssetDeviceConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAssetDeviceConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetDeviceConfiguration** | [**AssetDeviceConfiguration**](AssetDeviceConfiguration.md) | The &#39;asset.DeviceConfiguration&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**AssetDeviceConfiguration**](asset.DeviceConfiguration.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAssetDeviceContractInformation

> AssetDeviceContractInformation PatchAssetDeviceContractInformation(ctx, moid).AssetDeviceContractInformation(assetDeviceContractInformation).IfMatch(ifMatch).Execute()

Update a 'asset.DeviceContractInformation' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    assetDeviceContractInformation := openapiclient.asset.DeviceContractInformation{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Contract: openapiclient.asset.ContractInformation{ClassId: "ClassId_example", ObjectType: "ObjectType_example", BillTo: openapiclient.asset.AddressInformation{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Address1: "Address1_example", Address2: "Address2_example", Address3: "Address3_example", City: "City_example", Country: "Country_example", County: "County_example", Location: "Location_example", Name: "Name_example", PostalCode: "PostalCode_example", Province: "Province_example", State: "State_example"}, BillToGlobalUltimate: openapiclient.asset.GlobalUltimate{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Id: "Id_example", Name: "Name_example"}, ContractNumber: "ContractNumber_example", LineStatus: "LineStatus_example"}, ContractStatus: "ContractStatus_example", CoveredProductLineEndDate: "CoveredProductLineEndDate_example", DeviceId: "DeviceId_example", DeviceType: "DeviceType_example", EndCustomer: openapiclient.asset.CustomerInformation{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Address: openapiclient.asset.AddressInformation{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Address1: "Address1_example", Address2: "Address2_example", Address3: "Address3_example", City: "City_example", Country: "Country_example", County: "County_example", Location: "Location_example", Name: "Name_example", PostalCode: "PostalCode_example", Province: "Province_example", State: "State_example"}, Id: "Id_example", Name: "Name_example"}, EndUserGlobalUltimate: openapiclient.asset.GlobalUltimate{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Id: "Id_example", Name: "Name_example"}, IsValid: false, ItemType: "ItemType_example", MaintenancePurchaseOrderNumber: "MaintenancePurchaseOrderNumber_example", MaintenanceSalesOrderNumber: "MaintenanceSalesOrderNumber_example", PlatformType: "PlatformType_example", Product: openapiclient.asset.ProductInformation{ClassId: "ClassId_example", ObjectType: "ObjectType_example", BillTo: , Description: "Description_example", Family: "Family_example", Group: "Group_example", Number: "Number_example", ShipTo: , SubType: "SubType_example"}, PurchaseOrderNumber: "PurchaseOrderNumber_example", ResellerGlobalUltimate: , SalesOrderNumber: "SalesOrderNumber_example", ServiceDescription: "ServiceDescription_example", ServiceEndDate: "TODO", ServiceLevel: "ServiceLevel_example", ServiceSku: "ServiceSku_example", ServiceStartDate: "TODO", StateContract: "StateContract_example", WarrantyEndDate: "WarrantyEndDate_example", WarrantyType: "WarrantyType_example", RegisteredDevice: } // AssetDeviceContractInformation | The 'asset.DeviceContractInformation' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.PatchAssetDeviceContractInformation(context.Background(), moid, assetDeviceContractInformation).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.PatchAssetDeviceContractInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAssetDeviceContractInformation`: AssetDeviceContractInformation
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.PatchAssetDeviceContractInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAssetDeviceContractInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetDeviceContractInformation** | [**AssetDeviceContractInformation**](AssetDeviceContractInformation.md) | The &#39;asset.DeviceContractInformation&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**AssetDeviceContractInformation**](asset.DeviceContractInformation.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAssetDeviceRegistration

> AssetDeviceRegistration PatchAssetDeviceRegistration(ctx, moid).AssetDeviceRegistration(assetDeviceRegistration).IfMatch(ifMatch).Execute()

Updates the resource representing the device connector. For example, this can be used to annotate the device connector resource with user-specified tags.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    assetDeviceRegistration := openapiclient.asset.DeviceRegistration{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ApiVersion: int64(123), AppPartitionNumber: int64(123), ConnectionId: "ConnectionId_example", ConnectionReason: "ConnectionReason_example", ConnectionStatus: "ConnectionStatus_example", ConnectionStatusLastChangeTime: "TODO", ConnectorVersion: "ConnectorVersion_example", DeviceExternalIpAddress: "DeviceExternalIpAddress_example", ProxyApp: "ProxyApp_example", AccessKeyId: "AccessKeyId_example", ClaimedByUserName: "ClaimedByUserName_example", ClaimedTime: "TODO", DeviceHostname: []string{"DeviceHostname_example"), DeviceIpAddress: []string{"DeviceIpAddress_example"), ExecutionMode: "ExecutionMode_example", ParentSignature: , Pid: []string{"Pid_example"), PlatformType: "PlatformType_example", PublicAccessKey: "PublicAccessKey_example", ReadOnly: false, Serial: []string{"Serial_example"), Vendor: "Vendor_example", Account: , ClaimedByUser: , ClusterMembers: []AssetClusterMemberRelationship{), DeviceClaim: , DeviceConfiguration: , DomainGroup: , ParentConnection: } // AssetDeviceRegistration | The 'asset.DeviceRegistration' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.PatchAssetDeviceRegistration(context.Background(), moid, assetDeviceRegistration).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.PatchAssetDeviceRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAssetDeviceRegistration`: AssetDeviceRegistration
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.PatchAssetDeviceRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAssetDeviceRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetDeviceRegistration** | [**AssetDeviceRegistration**](AssetDeviceRegistration.md) | The &#39;asset.DeviceRegistration&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**AssetDeviceRegistration**](asset.DeviceRegistration.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAssetTarget

> AssetTarget PatchAssetTarget(ctx, moid).AssetTarget(assetTarget).IfMatch(ifMatch).Execute()

Update a 'asset.Target' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    assetTarget := openapiclient.asset.Target{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClaimedByUserName: "ClaimedByUserName_example", Connections: []AssetConnection{), ConnectorVersion: "ConnectorVersion_example", ExternalIpAddress: "ExternalIpAddress_example", IpAddress: []string{"IpAddress_example"), Name: "Name_example", ProductId: []string{"ProductId_example"), ReadOnly: false, Services: []AssetService{), Status: "Status_example", StatusErrorReason: "StatusErrorReason_example", TargetId: []string{"TargetId_example"), TargetType: "TargetType_example", Account: , Assist: , RegisteredDevice: } // AssetTarget | The 'asset.Target' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.PatchAssetTarget(context.Background(), moid, assetTarget).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.PatchAssetTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAssetTarget`: AssetTarget
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.PatchAssetTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAssetTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetTarget** | [**AssetTarget**](AssetTarget.md) | The &#39;asset.Target&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**AssetTarget**](asset.Target.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetDeviceConfiguration

> AssetDeviceConfiguration UpdateAssetDeviceConfiguration(ctx, moid).AssetDeviceConfiguration(assetDeviceConfiguration).IfMatch(ifMatch).Execute()

Update a 'asset.DeviceConfiguration' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    assetDeviceConfiguration := openapiclient.asset.DeviceConfiguration{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, LocalConfigurationLocked: false, LogLevel: "LogLevel_example", Device: } // AssetDeviceConfiguration | The 'asset.DeviceConfiguration' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.UpdateAssetDeviceConfiguration(context.Background(), moid, assetDeviceConfiguration).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.UpdateAssetDeviceConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetDeviceConfiguration`: AssetDeviceConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.UpdateAssetDeviceConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetDeviceConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetDeviceConfiguration** | [**AssetDeviceConfiguration**](AssetDeviceConfiguration.md) | The &#39;asset.DeviceConfiguration&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**AssetDeviceConfiguration**](asset.DeviceConfiguration.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetDeviceContractInformation

> AssetDeviceContractInformation UpdateAssetDeviceContractInformation(ctx, moid).AssetDeviceContractInformation(assetDeviceContractInformation).IfMatch(ifMatch).Execute()

Update a 'asset.DeviceContractInformation' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    assetDeviceContractInformation := openapiclient.asset.DeviceContractInformation{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Contract: openapiclient.asset.ContractInformation{ClassId: "ClassId_example", ObjectType: "ObjectType_example", BillTo: , BillToGlobalUltimate: , ContractNumber: "ContractNumber_example", LineStatus: "LineStatus_example"}, ContractStatus: "ContractStatus_example", CoveredProductLineEndDate: "CoveredProductLineEndDate_example", DeviceId: "DeviceId_example", DeviceType: "DeviceType_example", EndCustomer: openapiclient.asset.CustomerInformation{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Address: , Id: "Id_example", Name: "Name_example"}, EndUserGlobalUltimate: , IsValid: false, ItemType: "ItemType_example", MaintenancePurchaseOrderNumber: "MaintenancePurchaseOrderNumber_example", MaintenanceSalesOrderNumber: "MaintenanceSalesOrderNumber_example", PlatformType: "PlatformType_example", Product: openapiclient.asset.ProductInformation{ClassId: "ClassId_example", ObjectType: "ObjectType_example", BillTo: , Description: "Description_example", Family: "Family_example", Group: "Group_example", Number: "Number_example", ShipTo: , SubType: "SubType_example"}, PurchaseOrderNumber: "PurchaseOrderNumber_example", ResellerGlobalUltimate: , SalesOrderNumber: "SalesOrderNumber_example", ServiceDescription: "ServiceDescription_example", ServiceEndDate: "TODO", ServiceLevel: "ServiceLevel_example", ServiceSku: "ServiceSku_example", ServiceStartDate: "TODO", StateContract: "StateContract_example", WarrantyEndDate: "WarrantyEndDate_example", WarrantyType: "WarrantyType_example", RegisteredDevice: } // AssetDeviceContractInformation | The 'asset.DeviceContractInformation' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.UpdateAssetDeviceContractInformation(context.Background(), moid, assetDeviceContractInformation).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.UpdateAssetDeviceContractInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetDeviceContractInformation`: AssetDeviceContractInformation
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.UpdateAssetDeviceContractInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetDeviceContractInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetDeviceContractInformation** | [**AssetDeviceContractInformation**](AssetDeviceContractInformation.md) | The &#39;asset.DeviceContractInformation&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**AssetDeviceContractInformation**](asset.DeviceContractInformation.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetDeviceRegistration

> AssetDeviceRegistration UpdateAssetDeviceRegistration(ctx, moid).AssetDeviceRegistration(assetDeviceRegistration).IfMatch(ifMatch).Execute()

Updates the resource representing the device connector. For example, this can be used to annotate the device connector resource with user-specified tags.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    assetDeviceRegistration := openapiclient.asset.DeviceRegistration{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ApiVersion: int64(123), AppPartitionNumber: int64(123), ConnectionId: "ConnectionId_example", ConnectionReason: "ConnectionReason_example", ConnectionStatus: "ConnectionStatus_example", ConnectionStatusLastChangeTime: "TODO", ConnectorVersion: "ConnectorVersion_example", DeviceExternalIpAddress: "DeviceExternalIpAddress_example", ProxyApp: "ProxyApp_example", AccessKeyId: "AccessKeyId_example", ClaimedByUserName: "ClaimedByUserName_example", ClaimedTime: "TODO", DeviceHostname: []string{"DeviceHostname_example"), DeviceIpAddress: []string{"DeviceIpAddress_example"), ExecutionMode: "ExecutionMode_example", ParentSignature: , Pid: []string{"Pid_example"), PlatformType: "PlatformType_example", PublicAccessKey: "PublicAccessKey_example", ReadOnly: false, Serial: []string{"Serial_example"), Vendor: "Vendor_example", Account: , ClaimedByUser: , ClusterMembers: []AssetClusterMemberRelationship{), DeviceClaim: , DeviceConfiguration: , DomainGroup: , ParentConnection: } // AssetDeviceRegistration | The 'asset.DeviceRegistration' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.UpdateAssetDeviceRegistration(context.Background(), moid, assetDeviceRegistration).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.UpdateAssetDeviceRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetDeviceRegistration`: AssetDeviceRegistration
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.UpdateAssetDeviceRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetDeviceRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetDeviceRegistration** | [**AssetDeviceRegistration**](AssetDeviceRegistration.md) | The &#39;asset.DeviceRegistration&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**AssetDeviceRegistration**](asset.DeviceRegistration.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetTarget

> AssetTarget UpdateAssetTarget(ctx, moid).AssetTarget(assetTarget).IfMatch(ifMatch).Execute()

Update a 'asset.Target' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    assetTarget :=  // AssetTarget | The 'asset.Target' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetApi.UpdateAssetTarget(context.Background(), moid, assetTarget).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.UpdateAssetTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetTarget`: AssetTarget
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.UpdateAssetTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetTarget** | [**AssetTarget**](AssetTarget.md) | The &#39;asset.Target&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**AssetTarget**](asset.Target.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

