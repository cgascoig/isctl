# \MemoryApi

All URIs are relative to *https://intersight.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMemoryPersistentMemoryPolicy**](MemoryApi.md#CreateMemoryPersistentMemoryPolicy) | **Post** /memory/PersistentMemoryPolicies | Create a &#39;memory.PersistentMemoryPolicy&#39; resource.
[**DeleteMemoryPersistentMemoryPolicy**](MemoryApi.md#DeleteMemoryPersistentMemoryPolicy) | **Delete** /memory/PersistentMemoryPolicies/{Moid} | Delete a &#39;memory.PersistentMemoryPolicy&#39; resource.
[**GetMemoryArrayByMoid**](MemoryApi.md#GetMemoryArrayByMoid) | **Get** /memory/Arrays/{Moid} | Read a &#39;memory.Array&#39; resource.
[**GetMemoryArrayList**](MemoryApi.md#GetMemoryArrayList) | **Get** /memory/Arrays | Read a &#39;memory.Array&#39; resource.
[**GetMemoryPersistentMemoryConfigResultByMoid**](MemoryApi.md#GetMemoryPersistentMemoryConfigResultByMoid) | **Get** /memory/PersistentMemoryConfigResults/{Moid} | Read a &#39;memory.PersistentMemoryConfigResult&#39; resource.
[**GetMemoryPersistentMemoryConfigResultList**](MemoryApi.md#GetMemoryPersistentMemoryConfigResultList) | **Get** /memory/PersistentMemoryConfigResults | Read a &#39;memory.PersistentMemoryConfigResult&#39; resource.
[**GetMemoryPersistentMemoryConfigurationByMoid**](MemoryApi.md#GetMemoryPersistentMemoryConfigurationByMoid) | **Get** /memory/PersistentMemoryConfigurations/{Moid} | Read a &#39;memory.PersistentMemoryConfiguration&#39; resource.
[**GetMemoryPersistentMemoryConfigurationList**](MemoryApi.md#GetMemoryPersistentMemoryConfigurationList) | **Get** /memory/PersistentMemoryConfigurations | Read a &#39;memory.PersistentMemoryConfiguration&#39; resource.
[**GetMemoryPersistentMemoryNamespaceByMoid**](MemoryApi.md#GetMemoryPersistentMemoryNamespaceByMoid) | **Get** /memory/PersistentMemoryNamespaces/{Moid} | Read a &#39;memory.PersistentMemoryNamespace&#39; resource.
[**GetMemoryPersistentMemoryNamespaceConfigResultByMoid**](MemoryApi.md#GetMemoryPersistentMemoryNamespaceConfigResultByMoid) | **Get** /memory/PersistentMemoryNamespaceConfigResults/{Moid} | Read a &#39;memory.PersistentMemoryNamespaceConfigResult&#39; resource.
[**GetMemoryPersistentMemoryNamespaceConfigResultList**](MemoryApi.md#GetMemoryPersistentMemoryNamespaceConfigResultList) | **Get** /memory/PersistentMemoryNamespaceConfigResults | Read a &#39;memory.PersistentMemoryNamespaceConfigResult&#39; resource.
[**GetMemoryPersistentMemoryNamespaceList**](MemoryApi.md#GetMemoryPersistentMemoryNamespaceList) | **Get** /memory/PersistentMemoryNamespaces | Read a &#39;memory.PersistentMemoryNamespace&#39; resource.
[**GetMemoryPersistentMemoryPolicyByMoid**](MemoryApi.md#GetMemoryPersistentMemoryPolicyByMoid) | **Get** /memory/PersistentMemoryPolicies/{Moid} | Read a &#39;memory.PersistentMemoryPolicy&#39; resource.
[**GetMemoryPersistentMemoryPolicyList**](MemoryApi.md#GetMemoryPersistentMemoryPolicyList) | **Get** /memory/PersistentMemoryPolicies | Read a &#39;memory.PersistentMemoryPolicy&#39; resource.
[**GetMemoryPersistentMemoryRegionByMoid**](MemoryApi.md#GetMemoryPersistentMemoryRegionByMoid) | **Get** /memory/PersistentMemoryRegions/{Moid} | Read a &#39;memory.PersistentMemoryRegion&#39; resource.
[**GetMemoryPersistentMemoryRegionList**](MemoryApi.md#GetMemoryPersistentMemoryRegionList) | **Get** /memory/PersistentMemoryRegions | Read a &#39;memory.PersistentMemoryRegion&#39; resource.
[**GetMemoryPersistentMemoryUnitByMoid**](MemoryApi.md#GetMemoryPersistentMemoryUnitByMoid) | **Get** /memory/PersistentMemoryUnits/{Moid} | Read a &#39;memory.PersistentMemoryUnit&#39; resource.
[**GetMemoryPersistentMemoryUnitList**](MemoryApi.md#GetMemoryPersistentMemoryUnitList) | **Get** /memory/PersistentMemoryUnits | Read a &#39;memory.PersistentMemoryUnit&#39; resource.
[**GetMemoryUnitByMoid**](MemoryApi.md#GetMemoryUnitByMoid) | **Get** /memory/Units/{Moid} | Read a &#39;memory.Unit&#39; resource.
[**GetMemoryUnitList**](MemoryApi.md#GetMemoryUnitList) | **Get** /memory/Units | Read a &#39;memory.Unit&#39; resource.
[**PatchMemoryArray**](MemoryApi.md#PatchMemoryArray) | **Patch** /memory/Arrays/{Moid} | Update a &#39;memory.Array&#39; resource.
[**PatchMemoryPersistentMemoryConfigResult**](MemoryApi.md#PatchMemoryPersistentMemoryConfigResult) | **Patch** /memory/PersistentMemoryConfigResults/{Moid} | Update a &#39;memory.PersistentMemoryConfigResult&#39; resource.
[**PatchMemoryPersistentMemoryConfiguration**](MemoryApi.md#PatchMemoryPersistentMemoryConfiguration) | **Patch** /memory/PersistentMemoryConfigurations/{Moid} | Update a &#39;memory.PersistentMemoryConfiguration&#39; resource.
[**PatchMemoryPersistentMemoryNamespace**](MemoryApi.md#PatchMemoryPersistentMemoryNamespace) | **Patch** /memory/PersistentMemoryNamespaces/{Moid} | Update a &#39;memory.PersistentMemoryNamespace&#39; resource.
[**PatchMemoryPersistentMemoryNamespaceConfigResult**](MemoryApi.md#PatchMemoryPersistentMemoryNamespaceConfigResult) | **Patch** /memory/PersistentMemoryNamespaceConfigResults/{Moid} | Update a &#39;memory.PersistentMemoryNamespaceConfigResult&#39; resource.
[**PatchMemoryPersistentMemoryPolicy**](MemoryApi.md#PatchMemoryPersistentMemoryPolicy) | **Patch** /memory/PersistentMemoryPolicies/{Moid} | Update a &#39;memory.PersistentMemoryPolicy&#39; resource.
[**PatchMemoryPersistentMemoryRegion**](MemoryApi.md#PatchMemoryPersistentMemoryRegion) | **Patch** /memory/PersistentMemoryRegions/{Moid} | Update a &#39;memory.PersistentMemoryRegion&#39; resource.
[**PatchMemoryPersistentMemoryUnit**](MemoryApi.md#PatchMemoryPersistentMemoryUnit) | **Patch** /memory/PersistentMemoryUnits/{Moid} | Update a &#39;memory.PersistentMemoryUnit&#39; resource.
[**PatchMemoryUnit**](MemoryApi.md#PatchMemoryUnit) | **Patch** /memory/Units/{Moid} | Update a &#39;memory.Unit&#39; resource.
[**UpdateMemoryArray**](MemoryApi.md#UpdateMemoryArray) | **Post** /memory/Arrays/{Moid} | Update a &#39;memory.Array&#39; resource.
[**UpdateMemoryPersistentMemoryConfigResult**](MemoryApi.md#UpdateMemoryPersistentMemoryConfigResult) | **Post** /memory/PersistentMemoryConfigResults/{Moid} | Update a &#39;memory.PersistentMemoryConfigResult&#39; resource.
[**UpdateMemoryPersistentMemoryConfiguration**](MemoryApi.md#UpdateMemoryPersistentMemoryConfiguration) | **Post** /memory/PersistentMemoryConfigurations/{Moid} | Update a &#39;memory.PersistentMemoryConfiguration&#39; resource.
[**UpdateMemoryPersistentMemoryNamespace**](MemoryApi.md#UpdateMemoryPersistentMemoryNamespace) | **Post** /memory/PersistentMemoryNamespaces/{Moid} | Update a &#39;memory.PersistentMemoryNamespace&#39; resource.
[**UpdateMemoryPersistentMemoryNamespaceConfigResult**](MemoryApi.md#UpdateMemoryPersistentMemoryNamespaceConfigResult) | **Post** /memory/PersistentMemoryNamespaceConfigResults/{Moid} | Update a &#39;memory.PersistentMemoryNamespaceConfigResult&#39; resource.
[**UpdateMemoryPersistentMemoryPolicy**](MemoryApi.md#UpdateMemoryPersistentMemoryPolicy) | **Post** /memory/PersistentMemoryPolicies/{Moid} | Update a &#39;memory.PersistentMemoryPolicy&#39; resource.
[**UpdateMemoryPersistentMemoryRegion**](MemoryApi.md#UpdateMemoryPersistentMemoryRegion) | **Post** /memory/PersistentMemoryRegions/{Moid} | Update a &#39;memory.PersistentMemoryRegion&#39; resource.
[**UpdateMemoryPersistentMemoryUnit**](MemoryApi.md#UpdateMemoryPersistentMemoryUnit) | **Post** /memory/PersistentMemoryUnits/{Moid} | Update a &#39;memory.PersistentMemoryUnit&#39; resource.
[**UpdateMemoryUnit**](MemoryApi.md#UpdateMemoryUnit) | **Post** /memory/Units/{Moid} | Update a &#39;memory.Unit&#39; resource.



## CreateMemoryPersistentMemoryPolicy

> MemoryPersistentMemoryPolicy CreateMemoryPersistentMemoryPolicy(ctx).MemoryPersistentMemoryPolicy(memoryPersistentMemoryPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'memory.PersistentMemoryPolicy' resource.

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
    memoryPersistentMemoryPolicy := openapiclient.memory.PersistentMemoryPolicy{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{openapiclient.mo.Tag{Key: "Key_example", Value: "Value_example"}), VersionContext: openapiclient.mo.VersionContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", InterestedMos: []MoMoRef{openapiclient.mo.MoRef{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example"}), RefMo: openapiclient.mo.MoRef{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example"}, Timestamp: "TODO", Version: "Version_example", VersionType: "VersionType_example"}, Ancestors: []MoBaseMoRelationship{openapiclient.mo.BaseMo.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{openapiclient.mo.Tag{Key: "Key_example", Value: "Value_example"}), VersionContext: openapiclient.mo.VersionContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", InterestedMos: []MoMoRef{), RefMo: , Timestamp: "TODO", Version: "Version_example", VersionType: "VersionType_example"}, Ancestors: []MoBaseMoRelationship{openapiclient.mo.BaseMo.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }}), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }}), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Goals: []MemoryPersistentMemoryGoal{openapiclient.memory.PersistentMemoryGoal{ClassId: "ClassId_example", ObjectType: "ObjectType_example", MemoryModePercentage: 123, PersistentMemoryType: "PersistentMemoryType_example", SocketId: "SocketId_example"}), LocalSecurity: openapiclient.memory.PersistentMemoryLocalSecurity{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Enabled: false, IsSecurePassphraseSet: false, SecurePassphrase: "SecurePassphrase_example"}, LogicalNamespaces: []MemoryPersistentMemoryLogicalNamespace{openapiclient.memory.PersistentMemoryLogicalNamespace{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Capacity: 123, Mode: "Mode_example", Name: "Name_example", SocketId: 123, SocketMemoryId: "SocketMemoryId_example"}), ManagementMode: "ManagementMode_example", RetainNamespaces: false, Organization: openapiclient.organization.Organization.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: openapiclient.iam.Account.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Status: "Status_example", AppRegistrations: []IamAppRegistrationRelationship{openapiclient.iam.AppRegistration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientId: "ClientId_example", ClientName: "ClientName_example", ClientSecret: "ClientSecret_example", ClientType: "ClientType_example", Description: "Description_example", GrantTypes: []string{"GrantTypes_example"), RedirectUris: []string{"RedirectUris_example"), RenewClientSecret: false, ResponseTypes: []string{"ResponseTypes_example"), RevocationTimestamp: "TODO", Revoke: false, Account: openapiclient.iam.Account.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Status: "Status_example", AppRegistrations: []IamAppRegistrationRelationship{openapiclient.iam.AppRegistration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientId: "ClientId_example", ClientName: "ClientName_example", ClientSecret: "ClientSecret_example", ClientType: "ClientType_example", Description: "Description_example", GrantTypes: []string{"GrantTypes_example"), RedirectUris: []string{"RedirectUris_example"), RenewClientSecret: false, ResponseTypes: []string{"ResponseTypes_example"), RevocationTimestamp: "TODO", Revoke: false, Account: , OauthTokens: []IamOAuthTokenRelationship{openapiclient.iam.OAuthToken.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccessExpirationTime: "TODO", ClientId: "ClientId_example", ClientIpAddress: "ClientIpAddress_example", ClientName: "ClientName_example", ExpirationTime: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", TokenId: "TokenId_example", UserMeta: openapiclient.iam.ClientMeta{ClassId: "ClassId_example", ObjectType: "ObjectType_example", DeviceModel: "DeviceModel_example", UserAgent: "UserAgent_example"}, AppRegistration: , Permission: openapiclient.iam.Permission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , EndPointRoles: []IamEndPointRoleRelationship{openapiclient.iam.EndPointRole.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", RoleType: "RoleType_example", Type: "Type_example", Account: , EndPointPrivileges: []IamEndPointPrivilegeRelationship{openapiclient.iam.EndPointPrivilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", System: openapiclient.iam.System.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointPrivileges: []IamEndPointPrivilegeRelationship{openapiclient.iam.EndPointPrivilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", System: openapiclient.iam.System.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointPrivileges: []IamEndPointPrivilegeRelationship{), EndPointRoles: []IamEndPointRoleRelationship{openapiclient.iam.EndPointRole.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", RoleType: "RoleType_example", Type: "Type_example", Account: , EndPointPrivileges: []IamEndPointPrivilegeRelationship{), System: }), Idp: openapiclient.iam.Idp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", Metadata: "Metadata_example", Name: "Name_example", Type: "Type_example", Account: , LdapPolicy: openapiclient.iam.LdapPolicy.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", BaseProperties: openapiclient.iam.LdapBaseProperties{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Attribute: "Attribute_example", BaseDn: "BaseDn_example", BindDn: "BindDn_example", BindMethod: "BindMethod_example", Domain: "Domain_example", EnableEncryption: false, EnableGroupAuthorization: false, Filter: "Filter_example", GroupAttribute: "GroupAttribute_example", IsPasswordSet: false, NestedGroupSearchDepth: int64(123), Password: "Password_example", Timeout: int64(123)}, DnsParameters: openapiclient.iam.LdapDnsParameters{ClassId: "ClassId_example", ObjectType: "ObjectType_example", SearchDomain: "SearchDomain_example", SearchForest: "SearchForest_example", Source: "Source_example"}, EnableDns: false, Enabled: false, UserSearchPrecedence: "UserSearchPrecedence_example", Var0Idp: openapiclient.iam.Idp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", Metadata: "Metadata_example", Name: "Name_example", Type: "Type_example", Account: , LdapPolicy: openapiclient.iam.LdapPolicy.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", BaseProperties: openapiclient.iam.LdapBaseProperties{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Attribute: "Attribute_example", BaseDn: "BaseDn_example", BindDn: "BindDn_example", BindMethod: "BindMethod_example", Domain: "Domain_example", EnableEncryption: false, EnableGroupAuthorization: false, Filter: "Filter_example", GroupAttribute: "GroupAttribute_example", IsPasswordSet: false, NestedGroupSearchDepth: int64(123), Password: "Password_example", Timeout: int64(123)}, DnsParameters: openapiclient.iam.LdapDnsParameters{ClassId: "ClassId_example", ObjectType: "ObjectType_example", SearchDomain: "SearchDomain_example", SearchForest: "SearchForest_example", Source: "Source_example"}, EnableDns: false, Enabled: false, UserSearchPrecedence: "UserSearchPrecedence_example", Var0Idp: , ApplianceAccount: , Groups: []IamLdapGroupRelationship{openapiclient.iam.LdapGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Domain: "Domain_example", Name: "Name_example", EndPointRole: []IamEndPointRoleRelationship{), LdapPolicy: }), Organization: openapiclient.organization.Organization.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , ResourceGroups: []ResourceGroupRelationship{openapiclient.resource.Group.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", PerTypeCombinedSelector: []ResourcePerTypeCombinedSelector{openapiclient.resource.PerTypeCombinedSelector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", CombinedSelector: "CombinedSelector_example", EmptyFilter: false, SelectorObjectType: "SelectorObjectType_example"}), Qualifier: "Qualifier_example", Selectors: []ResourceSelector{openapiclient.resource.Selector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Selector: "Selector_example"}), Account: , Organizations: []OrganizationOrganizationRelationship{)})}, Profiles: []PolicyAbstractConfigProfileRelationship{openapiclient.policy.AbstractConfigProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: openapiclient.policy.AbstractProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: openapiclient.policy.AbstractProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: }}, Action: "Action_example", ConfigContext: openapiclient.policy.ConfigContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", ConfigState: "ConfigState_example", ControlAction: "ControlAction_example", ErrorState: "ErrorState_example", OperState: "OperState_example"}}), Providers: []IamLdapProviderRelationship{openapiclient.iam.LdapProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Port: int64(123), Server: "Server_example", LdapPolicy: })}, System: , UserPreferences: []IamUserPreferenceRelationship{openapiclient.iam.UserPreference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Preference: 123, Idp: , IdpReference: openapiclient.iam.IdpReference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", MultiFactorAuthentication: false, Name: "Name_example", Account: , Idp: , UserPreferences: []IamUserPreferenceRelationship{openapiclient.iam.UserPreference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Preference: 123, Idp: , IdpReference: openapiclient.iam.IdpReference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", MultiFactorAuthentication: false, Name: "Name_example", Account: , Idp: , UserPreferences: []IamUserPreferenceRelationship{), Usergroups: []IamUserGroupRelationship{openapiclient.iam.UserGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Idp: , Idpreference: , Permissions: []IamPermissionRelationship{openapiclient.iam.Permission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , EndPointRoles: []IamEndPointRoleRelationship{), ResourceRoles: []IamResourceRolesRelationship{openapiclient.iam.ResourceRoles.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointRoles: []IamEndPointRoleRelationship{), Permission: , Resource: , Roles: []IamRoleRelationship{openapiclient.iam.Role.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , PrivilegeSets: []IamPrivilegeSetRelationship{openapiclient.iam.PrivilegeSet.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , AssociatedPrivilegeSets: []IamPrivilegeSetRelationship{openapiclient.iam.PrivilegeSet.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , AssociatedPrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{openapiclient.iam.Privilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HostnamePrefix: "HostnamePrefix_example", Method: "Method_example", Name: "Name_example", RestPath: "RestPath_example", UrlPrefix: "UrlPrefix_example", Account: , System: }), System: }), Privileges: []IamPrivilegeRelationship{openapiclient.iam.Privilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HostnamePrefix: "HostnamePrefix_example", Method: "Method_example", Name: "Name_example", RestPath: "RestPath_example", UrlPrefix: "UrlPrefix_example", Account: , System: }), System: }), System: })}), Roles: []IamRoleRelationship{openapiclient.iam.Role.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , PrivilegeSets: []IamPrivilegeSetRelationship{), System: }), SessionLimits: openapiclient.iam.SessionLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, IdleTimeOut: int64(123), MaximumLimit: int64(123), PerUserLimit: int64(123), SessionTimeOut: int64(123), Account: , Permission: }, UserGroups: []IamUserGroupRelationship{openapiclient.iam.UserGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Idp: , Idpreference: , Permissions: []IamPermissionRelationship{), Qualifier: openapiclient.iam.Qualifier.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Value: []string{"Value_example"), Usergroup: }, Users: []IamUserRelationship{openapiclient.iam.User.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientIpAddress: "ClientIpAddress_example", Email: "Email_example", FirstName: "FirstName_example", LastLoginTime: "TODO", LastName: "LastName_example", Name: "Name_example", UserIdOrEmail: "UserIdOrEmail_example", UserType: "UserType_example", ApiKeys: []IamApiKeyRelationship{openapiclient.iam.ApiKey.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HashAlgorithm: "HashAlgorithm_example", KeySpec: openapiclient.pkix.KeyGenerationSpec{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Name: "Name_example"}, PrivateKey: "PrivateKey_example", Purpose: "Purpose_example", SigningAlgorithm: "SigningAlgorithm_example", Permission: , User: openapiclient.iam.User.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientIpAddress: "ClientIpAddress_example", Email: "Email_example", FirstName: "FirstName_example", LastLoginTime: "TODO", LastName: "LastName_example", Name: "Name_example", UserIdOrEmail: "UserIdOrEmail_example", UserType: "UserType_example", ApiKeys: []IamApiKeyRelationship{openapiclient.iam.ApiKey.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HashAlgorithm: "HashAlgorithm_example", KeySpec: openapiclient.pkix.KeyGenerationSpec{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Name: "Name_example"}, PrivateKey: "PrivateKey_example", Purpose: "Purpose_example", SigningAlgorithm: "SigningAlgorithm_example", Permission: , User: }), AppRegistrations: []IamAppRegistrationRelationship{), Idp: , Idpreference: , LocalUserPassword: openapiclient.iam.LocalUserPassword.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, CurrentPassword: "CurrentPassword_example", NewPassword: "NewPassword_example", Password: 123, User: }, OauthTokens: []IamOAuthTokenRelationship{openapiclient.iam.OAuthToken.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccessExpirationTime: "TODO", ClientId: "ClientId_example", ClientIpAddress: "ClientIpAddress_example", ClientName: "ClientName_example", ExpirationTime: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", TokenId: "TokenId_example", UserMeta: openapiclient.iam.ClientMeta{ClassId: "ClassId_example", ObjectType: "ObjectType_example", DeviceModel: "DeviceModel_example", UserAgent: "UserAgent_example"}, AppRegistration: , Permission: , User: }), Permissions: []IamPermissionRelationship{), Sessions: []IamSessionRelationship{openapiclient.iam.Session.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccountPermissions: []IamAccountPermissions{openapiclient.iam.AccountPermissions{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountIdentifier: "AccountIdentifier_example", AccountName: "AccountName_example", AccountStatus: "AccountStatus_example", Permissions: []IamPermissionReference{openapiclient.iam.PermissionReference{ClassId: "ClassId_example", ObjectType: "ObjectType_example", PermissionIdentifier: "PermissionIdentifier_example", PermissionName: "PermissionName_example"})}), ClientIpAddress: "ClientIpAddress_example", Expiration: "TODO", IdleTimeExpiration: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", Permission: , User: })}}), AppRegistrations: []IamAppRegistrationRelationship{), Idp: , Idpreference: , LocalUserPassword: openapiclient.iam.LocalUserPassword.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, CurrentPassword: "CurrentPassword_example", NewPassword: "NewPassword_example", Password: 123, User: }, OauthTokens: []IamOAuthTokenRelationship{), Permissions: []IamPermissionRelationship{), Sessions: []IamSessionRelationship{openapiclient.iam.Session.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccountPermissions: []IamAccountPermissions{openapiclient.iam.AccountPermissions{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountIdentifier: "AccountIdentifier_example", AccountName: "AccountName_example", AccountStatus: "AccountStatus_example", Permissions: []IamPermissionReference{openapiclient.iam.PermissionReference{ClassId: "ClassId_example", ObjectType: "ObjectType_example", PermissionIdentifier: "PermissionIdentifier_example", PermissionName: "PermissionName_example"})}), ClientIpAddress: "ClientIpAddress_example", Expiration: "TODO", IdleTimeExpiration: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", Permission: , User: })})}), Users: []IamUserRelationship{)}), Qualifier: openapiclient.iam.Qualifier.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Value: []string{"Value_example"), Usergroup: }, Users: []IamUserRelationship{)}), Users: []IamUserRelationship{)}}), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}}), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, ApplianceAccount: , Groups: []IamLdapGroupRelationship{openapiclient.iam.LdapGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Domain: "Domain_example", Name: "Name_example", EndPointRole: []IamEndPointRoleRelationship{), LdapPolicy: }), Organization: , Profiles: []PolicyAbstractConfigProfileRelationship{openapiclient.policy.AbstractConfigProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: , Action: "Action_example", ConfigContext: openapiclient.policy.ConfigContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", ConfigState: "ConfigState_example", ControlAction: "ControlAction_example", ErrorState: "ErrorState_example", OperState: "OperState_example"}}), Providers: []IamLdapProviderRelationship{openapiclient.iam.LdapProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Port: int64(123), Server: "Server_example", LdapPolicy: })}, System: , UserPreferences: []IamUserPreferenceRelationship{), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), Roles: []IamRoleRelationship{), ServiceProvider: openapiclient.iam.ServiceProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EntityId: "EntityId_example", Metadata: "Metadata_example", Name: "Name_example", System: }}}), EndPointRoles: []IamEndPointRoleRelationship{), Idp: , PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), Roles: []IamRoleRelationship{), ServiceProvider: openapiclient.iam.ServiceProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EntityId: "EntityId_example", Metadata: "Metadata_example", Name: "Name_example", System: }}}), System: }), ResourceRoles: []IamResourceRolesRelationship{openapiclient.iam.ResourceRoles.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointRoles: []IamEndPointRoleRelationship{), Permission: , Resource: , Roles: []IamRoleRelationship{)}), Roles: []IamRoleRelationship{), SessionLimits: openapiclient.iam.SessionLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, IdleTimeOut: int64(123), MaximumLimit: int64(123), PerUserLimit: int64(123), SessionTimeOut: int64(123), Account: , Permission: }, UserGroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, User: }), Permission: , Roles: []IamRoleRelationship{), User: }), DomainGroups: []IamDomainGroupRelationship{openapiclient.iam.DomainGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Partition1: int64(123), Partition2: int64(123), Partition3: int64(123), PartitionKey: "PartitionKey_example", Usage: int64(123), Account: }), EndPointRoles: []IamEndPointRoleRelationship{), Idpreferences: []IamIdpReferenceRelationship{), Idps: []IamIdpRelationship{), Permissions: []IamPermissionRelationship{), PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), ResourceLimits: openapiclient.iam.ResourceLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PerAccountUserLimit: int64(123), Account: }, Roles: []IamRoleRelationship{), SecurityHolder: openapiclient.iam.SecurityHolder.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Account: , ResourcePermissions: []IamResourcePermissionRelationship{openapiclient.iam.ResourcePermission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PermissionRoles: []IamPermissionToRoles{openapiclient.iam.PermissionToRoles{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Permission: openapiclient.cmrf.CmRf{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example"}, Roles: []CmrfCmRf{openapiclient.cmrf.CmRf{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example"})}), TargetApp: "TargetApp_example", Holder: openapiclient.iam.SecurityHolder.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Account: , ResourcePermissions: []IamResourcePermissionRelationship{openapiclient.iam.ResourcePermission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PermissionRoles: []IamPermissionToRoles{openapiclient.iam.PermissionToRoles{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Permission: , Roles: []CmrfCmRf{)}), TargetApp: "TargetApp_example", Holder: , Resource: })}, Resource: })}, SessionLimits: }, OauthTokens: []IamOAuthTokenRelationship{), Permission: , Roles: []IamRoleRelationship{), User: }), DomainGroups: []IamDomainGroupRelationship{openapiclient.iam.DomainGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Partition1: int64(123), Partition2: int64(123), Partition3: int64(123), PartitionKey: "PartitionKey_example", Usage: int64(123), Account: }), EndPointRoles: []IamEndPointRoleRelationship{), Idpreferences: []IamIdpReferenceRelationship{), Idps: []IamIdpRelationship{), Permissions: []IamPermissionRelationship{), PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), ResourceLimits: openapiclient.iam.ResourceLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PerAccountUserLimit: int64(123), Account: }, Roles: []IamRoleRelationship{), SecurityHolder: , SessionLimits: }, ResourceGroups: []ResourceGroupRelationship{openapiclient.resource.Group.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", PerTypeCombinedSelector: []ResourcePerTypeCombinedSelector{openapiclient.resource.PerTypeCombinedSelector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", CombinedSelector: "CombinedSelector_example", EmptyFilter: false, SelectorObjectType: "SelectorObjectType_example"}), Qualifier: "Qualifier_example", Selectors: []ResourceSelector{openapiclient.resource.Selector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Selector: "Selector_example"}), Account: , Organizations: []OrganizationOrganizationRelationship{)})}, Profiles: []PolicyAbstractConfigProfileRelationship{)} // MemoryPersistentMemoryPolicy | The 'memory.PersistentMemoryPolicy' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.CreateMemoryPersistentMemoryPolicy(context.Background(), memoryPersistentMemoryPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.CreateMemoryPersistentMemoryPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMemoryPersistentMemoryPolicy`: MemoryPersistentMemoryPolicy
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.CreateMemoryPersistentMemoryPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMemoryPersistentMemoryPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **memoryPersistentMemoryPolicy** | [**MemoryPersistentMemoryPolicy**](MemoryPersistentMemoryPolicy.md) | The &#39;memory.PersistentMemoryPolicy&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**MemoryPersistentMemoryPolicy**](memory.PersistentMemoryPolicy.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMemoryPersistentMemoryPolicy

> DeleteMemoryPersistentMemoryPolicy(ctx, moid).Execute()

Delete a 'memory.PersistentMemoryPolicy' resource.

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
    resp, r, err := api_client.MemoryApi.DeleteMemoryPersistentMemoryPolicy(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.DeleteMemoryPersistentMemoryPolicy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMemoryPersistentMemoryPolicyRequest struct via the builder pattern


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


## GetMemoryArrayByMoid

> MemoryArray GetMemoryArrayByMoid(ctx, moid).Execute()

Read a 'memory.Array' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryArrayByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryArrayByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryArrayByMoid`: MemoryArray
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryArrayByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryArrayByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MemoryArray**](memory.Array.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryArrayList

> MemoryArrayResponse GetMemoryArrayList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'memory.Array' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryArrayList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryArrayList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryArrayList`: MemoryArrayResponse
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryArrayList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryArrayListRequest struct via the builder pattern


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

[**MemoryArrayResponse**](memory.Array.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryPersistentMemoryConfigResultByMoid

> MemoryPersistentMemoryConfigResult GetMemoryPersistentMemoryConfigResultByMoid(ctx, moid).Execute()

Read a 'memory.PersistentMemoryConfigResult' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryPersistentMemoryConfigResultByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryPersistentMemoryConfigResultByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryPersistentMemoryConfigResultByMoid`: MemoryPersistentMemoryConfigResult
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryPersistentMemoryConfigResultByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryPersistentMemoryConfigResultByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MemoryPersistentMemoryConfigResult**](memory.PersistentMemoryConfigResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryPersistentMemoryConfigResultList

> MemoryPersistentMemoryConfigResultResponse GetMemoryPersistentMemoryConfigResultList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'memory.PersistentMemoryConfigResult' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryPersistentMemoryConfigResultList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryPersistentMemoryConfigResultList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryPersistentMemoryConfigResultList`: MemoryPersistentMemoryConfigResultResponse
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryPersistentMemoryConfigResultList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryPersistentMemoryConfigResultListRequest struct via the builder pattern


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

[**MemoryPersistentMemoryConfigResultResponse**](memory.PersistentMemoryConfigResult.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryPersistentMemoryConfigurationByMoid

> MemoryPersistentMemoryConfiguration GetMemoryPersistentMemoryConfigurationByMoid(ctx, moid).Execute()

Read a 'memory.PersistentMemoryConfiguration' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryPersistentMemoryConfigurationByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryPersistentMemoryConfigurationByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryPersistentMemoryConfigurationByMoid`: MemoryPersistentMemoryConfiguration
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryPersistentMemoryConfigurationByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryPersistentMemoryConfigurationByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MemoryPersistentMemoryConfiguration**](memory.PersistentMemoryConfiguration.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryPersistentMemoryConfigurationList

> MemoryPersistentMemoryConfigurationResponse GetMemoryPersistentMemoryConfigurationList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'memory.PersistentMemoryConfiguration' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryPersistentMemoryConfigurationList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryPersistentMemoryConfigurationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryPersistentMemoryConfigurationList`: MemoryPersistentMemoryConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryPersistentMemoryConfigurationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryPersistentMemoryConfigurationListRequest struct via the builder pattern


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

[**MemoryPersistentMemoryConfigurationResponse**](memory.PersistentMemoryConfiguration.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryPersistentMemoryNamespaceByMoid

> MemoryPersistentMemoryNamespace GetMemoryPersistentMemoryNamespaceByMoid(ctx, moid).Execute()

Read a 'memory.PersistentMemoryNamespace' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryPersistentMemoryNamespaceByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryPersistentMemoryNamespaceByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryPersistentMemoryNamespaceByMoid`: MemoryPersistentMemoryNamespace
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryPersistentMemoryNamespaceByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryPersistentMemoryNamespaceByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MemoryPersistentMemoryNamespace**](memory.PersistentMemoryNamespace.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryPersistentMemoryNamespaceConfigResultByMoid

> MemoryPersistentMemoryNamespaceConfigResult GetMemoryPersistentMemoryNamespaceConfigResultByMoid(ctx, moid).Execute()

Read a 'memory.PersistentMemoryNamespaceConfigResult' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryPersistentMemoryNamespaceConfigResultByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryPersistentMemoryNamespaceConfigResultByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryPersistentMemoryNamespaceConfigResultByMoid`: MemoryPersistentMemoryNamespaceConfigResult
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryPersistentMemoryNamespaceConfigResultByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryPersistentMemoryNamespaceConfigResultByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MemoryPersistentMemoryNamespaceConfigResult**](memory.PersistentMemoryNamespaceConfigResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryPersistentMemoryNamespaceConfigResultList

> MemoryPersistentMemoryNamespaceConfigResultResponse GetMemoryPersistentMemoryNamespaceConfigResultList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'memory.PersistentMemoryNamespaceConfigResult' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryPersistentMemoryNamespaceConfigResultList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryPersistentMemoryNamespaceConfigResultList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryPersistentMemoryNamespaceConfigResultList`: MemoryPersistentMemoryNamespaceConfigResultResponse
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryPersistentMemoryNamespaceConfigResultList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryPersistentMemoryNamespaceConfigResultListRequest struct via the builder pattern


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

[**MemoryPersistentMemoryNamespaceConfigResultResponse**](memory.PersistentMemoryNamespaceConfigResult.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryPersistentMemoryNamespaceList

> MemoryPersistentMemoryNamespaceResponse GetMemoryPersistentMemoryNamespaceList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'memory.PersistentMemoryNamespace' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryPersistentMemoryNamespaceList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryPersistentMemoryNamespaceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryPersistentMemoryNamespaceList`: MemoryPersistentMemoryNamespaceResponse
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryPersistentMemoryNamespaceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryPersistentMemoryNamespaceListRequest struct via the builder pattern


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

[**MemoryPersistentMemoryNamespaceResponse**](memory.PersistentMemoryNamespace.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryPersistentMemoryPolicyByMoid

> MemoryPersistentMemoryPolicy GetMemoryPersistentMemoryPolicyByMoid(ctx, moid).Execute()

Read a 'memory.PersistentMemoryPolicy' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryPersistentMemoryPolicyByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryPersistentMemoryPolicyByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryPersistentMemoryPolicyByMoid`: MemoryPersistentMemoryPolicy
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryPersistentMemoryPolicyByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryPersistentMemoryPolicyByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MemoryPersistentMemoryPolicy**](memory.PersistentMemoryPolicy.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryPersistentMemoryPolicyList

> MemoryPersistentMemoryPolicyResponse GetMemoryPersistentMemoryPolicyList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'memory.PersistentMemoryPolicy' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryPersistentMemoryPolicyList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryPersistentMemoryPolicyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryPersistentMemoryPolicyList`: MemoryPersistentMemoryPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryPersistentMemoryPolicyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryPersistentMemoryPolicyListRequest struct via the builder pattern


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

[**MemoryPersistentMemoryPolicyResponse**](memory.PersistentMemoryPolicy.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryPersistentMemoryRegionByMoid

> MemoryPersistentMemoryRegion GetMemoryPersistentMemoryRegionByMoid(ctx, moid).Execute()

Read a 'memory.PersistentMemoryRegion' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryPersistentMemoryRegionByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryPersistentMemoryRegionByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryPersistentMemoryRegionByMoid`: MemoryPersistentMemoryRegion
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryPersistentMemoryRegionByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryPersistentMemoryRegionByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MemoryPersistentMemoryRegion**](memory.PersistentMemoryRegion.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryPersistentMemoryRegionList

> MemoryPersistentMemoryRegionResponse GetMemoryPersistentMemoryRegionList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'memory.PersistentMemoryRegion' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryPersistentMemoryRegionList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryPersistentMemoryRegionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryPersistentMemoryRegionList`: MemoryPersistentMemoryRegionResponse
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryPersistentMemoryRegionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryPersistentMemoryRegionListRequest struct via the builder pattern


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

[**MemoryPersistentMemoryRegionResponse**](memory.PersistentMemoryRegion.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryPersistentMemoryUnitByMoid

> MemoryPersistentMemoryUnit GetMemoryPersistentMemoryUnitByMoid(ctx, moid).Execute()

Read a 'memory.PersistentMemoryUnit' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryPersistentMemoryUnitByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryPersistentMemoryUnitByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryPersistentMemoryUnitByMoid`: MemoryPersistentMemoryUnit
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryPersistentMemoryUnitByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryPersistentMemoryUnitByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MemoryPersistentMemoryUnit**](memory.PersistentMemoryUnit.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryPersistentMemoryUnitList

> MemoryPersistentMemoryUnitResponse GetMemoryPersistentMemoryUnitList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'memory.PersistentMemoryUnit' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryPersistentMemoryUnitList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryPersistentMemoryUnitList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryPersistentMemoryUnitList`: MemoryPersistentMemoryUnitResponse
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryPersistentMemoryUnitList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryPersistentMemoryUnitListRequest struct via the builder pattern


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

[**MemoryPersistentMemoryUnitResponse**](memory.PersistentMemoryUnit.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryUnitByMoid

> MemoryUnit GetMemoryUnitByMoid(ctx, moid).Execute()

Read a 'memory.Unit' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryUnitByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryUnitByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryUnitByMoid`: MemoryUnit
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryUnitByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryUnitByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MemoryUnit**](memory.Unit.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoryUnitList

> MemoryUnitResponse GetMemoryUnitList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'memory.Unit' resource.

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
    resp, r, err := api_client.MemoryApi.GetMemoryUnitList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.GetMemoryUnitList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemoryUnitList`: MemoryUnitResponse
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.GetMemoryUnitList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoryUnitListRequest struct via the builder pattern


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

[**MemoryUnitResponse**](memory.Unit.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMemoryArray

> MemoryArray PatchMemoryArray(ctx, moid).MemoryArray(memoryArray).IfMatch(ifMatch).Execute()

Update a 'memory.Array' resource.

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
    memoryArray := openapiclient.memory.Array{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ArrayId: int64(123), CpuId: int64(123), CurrentCapacity: "CurrentCapacity_example", ErrorCorrection: "ErrorCorrection_example", MaxCapacity: "MaxCapacity_example", MaxDevices: "MaxDevices_example", OperPowerState: "OperPowerState_example", Presence: "Presence_example", ComputeBoard: openapiclient.compute.Board.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", BoardId: int64(123), CpuTypeController: "CpuTypeController_example", OperPowerState: "OperPowerState_example", Presence: "Presence_example", ComputeBlade: openapiclient.compute.Blade.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminPowerState: "AdminPowerState_example", AssetTag: "AssetTag_example", AvailableMemory: int64(123), FaultSummary: int64(123), KvmIpAddresses: []ComputeIpAddress{openapiclient.compute.IpAddress{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Address: "Address_example", Category: "Category_example", DefaultGateway: "DefaultGateway_example", Dn: "Dn_example", HttpPort: int64(123), HttpsPort: int64(123), KvmPort: int64(123), Name: "Name_example", Subnet: "Subnet_example", Type: "Type_example"}), MemorySpeed: "MemorySpeed_example", MgmtIpAddress: "MgmtIpAddress_example", NumAdaptors: int64(123), NumCpuCores: int64(123), NumCpuCoresEnabled: int64(123), NumCpus: int64(123), NumEthHostInterfaces: int64(123), NumFcHostInterfaces: int64(123), NumThreads: int64(123), OperPowerState: "OperPowerState_example", OperState: "OperState_example", Operability: "Operability_example", PlatformType: "PlatformType_example", Presence: "Presence_example", ServiceProfile: "ServiceProfile_example", TotalMemory: int64(123), UserLabel: "UserLabel_example", Uuid: "Uuid_example", ChassisId: "ChassisId_example", ScaledMode: "ScaledMode_example", SlotId: int64(123), Adapters: []AdapterUnitRelationship{openapiclient.adapter.Unit.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdapterId: "AdapterId_example", BaseMacAddress: "BaseMacAddress_example", Integrated: "Integrated_example", OperState: "OperState_example", Operability: "Operability_example", PartNumber: "PartNumber_example", PciSlot: "PciSlot_example", Power: "Power_example", Presence: "Presence_example", Thermal: "Thermal_example", Vid: "Vid_example", ComputeBlade: openapiclient.compute.Blade.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminPowerState: "AdminPowerState_example", AssetTag: "AssetTag_example", AvailableMemory: int64(123), FaultSummary: int64(123), KvmIpAddresses: []ComputeIpAddress{openapiclient.compute.IpAddress{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Address: "Address_example", Category: "Category_example", DefaultGateway: "DefaultGateway_example", Dn: "Dn_example", HttpPort: int64(123), HttpsPort: int64(123), KvmPort: int64(123), Name: "Name_example", Subnet: "Subnet_example", Type: "Type_example"}), MemorySpeed: "MemorySpeed_example", MgmtIpAddress: "MgmtIpAddress_example", NumAdaptors: int64(123), NumCpuCores: int64(123), NumCpuCoresEnabled: int64(123), NumCpus: int64(123), NumEthHostInterfaces: int64(123), NumFcHostInterfaces: int64(123), NumThreads: int64(123), OperPowerState: "OperPowerState_example", OperState: "OperState_example", Operability: "Operability_example", PlatformType: "PlatformType_example", Presence: "Presence_example", ServiceProfile: "ServiceProfile_example", TotalMemory: int64(123), UserLabel: "UserLabel_example", Uuid: "Uuid_example", ChassisId: "ChassisId_example", ScaledMode: "ScaledMode_example", SlotId: int64(123), Adapters: []AdapterUnitRelationship{openapiclient.adapter.Unit.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdapterId: "AdapterId_example", BaseMacAddress: "BaseMacAddress_example", Integrated: "Integrated_example", OperState: "OperState_example", Operability: "Operability_example", PartNumber: "PartNumber_example", PciSlot: "PciSlot_example", Power: "Power_example", Presence: "Presence_example", Thermal: "Thermal_example", Vid: "Vid_example", ComputeBlade: , ComputeRackUnit: openapiclient.compute.RackUnit.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminPowerState: "AdminPowerState_example", AssetTag: "AssetTag_example", AvailableMemory: int64(123), FaultSummary: int64(123), KvmIpAddresses: []ComputeIpAddress{), MemorySpeed: "MemorySpeed_example", MgmtIpAddress: "MgmtIpAddress_example", NumAdaptors: int64(123), NumCpuCores: int64(123), NumCpuCoresEnabled: int64(123), NumCpus: int64(123), NumEthHostInterfaces: int64(123), NumFcHostInterfaces: int64(123), NumThreads: int64(123), OperPowerState: "OperPowerState_example", OperState: "OperState_example", Operability: "Operability_example", PlatformType: "PlatformType_example", Presence: "Presence_example", ServiceProfile: "ServiceProfile_example", TotalMemory: int64(123), UserLabel: "UserLabel_example", Uuid: "Uuid_example", ServerId: int64(123), Adapters: []AdapterUnitRelationship{), BiosBootmode: openapiclient.bios.BootMode.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", ActualBootMode: "ActualBootMode_example", ComputeRackUnit: openapiclient.compute.RackUnit.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminPowerState: "AdminPowerState_example", AssetTag: "AssetTag_example", AvailableMemory: int64(123), FaultSummary: int64(123), KvmIpAddresses: []ComputeIpAddress{), MemorySpeed: "MemorySpeed_example", MgmtIpAddress: "MgmtIpAddress_example", NumAdaptors: int64(123), NumCpuCores: int64(123), NumCpuCoresEnabled: int64(123), NumCpus: int64(123), NumEthHostInterfaces: int64(123), NumFcHostInterfaces: int64(123), NumThreads: int64(123), OperPowerState: "OperPowerState_example", OperState: "OperState_example", Operability: "Operability_example", PlatformType: "PlatformType_example", Presence: "Presence_example", ServiceProfile: "ServiceProfile_example", TotalMemory: int64(123), UserLabel: "UserLabel_example", Uuid: "Uuid_example", ServerId: int64(123), Adapters: []AdapterUnitRelationship{), BiosBootmode: openapiclient.bios.BootMode.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", ActualBootMode: "ActualBootMode_example", ComputeRackUnit: , RegisteredDevice: openapiclient.asset.DeviceRegistration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ApiVersion: int64(123), AppPartitionNumber: int64(123), ConnectionId: "ConnectionId_example", ConnectionReason: "ConnectionReason_example", ConnectionStatus: "ConnectionStatus_example", ConnectionStatusLastChangeTime: "TODO", ConnectorVersion: "ConnectorVersion_example", DeviceExternalIpAddress: "DeviceExternalIpAddress_example", ProxyApp: "ProxyApp_example", AccessKeyId: "AccessKeyId_example", ClaimedByUserName: "ClaimedByUserName_example", ClaimedTime: "TODO", DeviceHostname: []string{"DeviceHostname_example"), DeviceIpAddress: []string{"DeviceIpAddress_example"), ExecutionMode: "ExecutionMode_example", ParentSignature: openapiclient.asset.ParentConnectionSignature{ClassId: "ClassId_example", ObjectType: "ObjectType_example", DeviceId: "DeviceId_example", NodeId: "NodeId_example", Signature: 123, TimeStamp: "TODO"}, Pid: []string{"Pid_example"), PlatformType: "PlatformType_example", PublicAccessKey: "PublicAccessKey_example", ReadOnly: false, Serial: []string{"Serial_example"), Vendor: "Vendor_example", Account: , ClaimedByUser: , ClusterMembers: []AssetClusterMemberRelationship{openapiclient.asset.ClusterMember.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ApiVersion: int64(123), AppPartitionNumber: int64(123), ConnectionId: "ConnectionId_example", ConnectionReason: "ConnectionReason_example", ConnectionStatus: "ConnectionStatus_example", ConnectionStatusLastChangeTime: "TODO", ConnectorVersion: "ConnectorVersion_example", DeviceExternalIpAddress: "DeviceExternalIpAddress_example", ProxyApp: "ProxyApp_example", Leadership: "Leadership_example", LockedLeader: false, MemberIdentity: "MemberIdentity_example", ParentClusterMemberIdentity: "ParentClusterMemberIdentity_example", Sudi: openapiclient.asset.SudiInfo{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Pid: "Pid_example", SerialNumber: "SerialNumber_example", Signature: "Signature_example", Status: "Status_example", SudiCertificate: openapiclient.x509.Certificate{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Issuer: openapiclient.pkix.DistinguishedName{ClassId: "ClassId_example", ObjectType: "ObjectType_example", CommonName: "CommonName_example", Country: []string{"Country_example"), Locality: []string{"Locality_example"), Organization: []string{"Organization_example"), OrganizationalUnit: []string{"OrganizationalUnit_example"), State: []string{"State_example")}, NotAfter: "TODO", NotBefore: "TODO", PemCertificate: "PemCertificate_example", Sha256Fingerprint: "Sha256Fingerprint_example", SignatureAlgorithm: "SignatureAlgorithm_example", Subject: openapiclient.pkix.DistinguishedName{ClassId: "ClassId_example", ObjectType: "ObjectType_example", CommonName: "CommonName_example", Country: []string{"Country_example"), Locality: []string{"Locality_example"), Organization: []string{"Organization_example"), OrganizationalUnit: []string{"OrganizationalUnit_example"), State: []string{"State_example")}}}, Device: openapiclient.asset.DeviceRegistration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ApiVersion: int64(123), AppPartitionNumber: int64(123), ConnectionId: "ConnectionId_example", ConnectionReason: "ConnectionReason_example", ConnectionStatus: "ConnectionStatus_example", ConnectionStatusLastChangeTime: "TODO", ConnectorVersion: "ConnectorVersion_example", DeviceExternalIpAddress: "DeviceExternalIpAddress_example", ProxyApp: "ProxyApp_example", AccessKeyId: "AccessKeyId_example", ClaimedByUserName: "ClaimedByUserName_example", ClaimedTime: "TODO", DeviceHostname: []string{"DeviceHostname_example"), DeviceIpAddress: []string{"DeviceIpAddress_example"), ExecutionMode: "ExecutionMode_example", ParentSignature: openapiclient.asset.ParentConnectionSignature{ClassId: "ClassId_example", ObjectType: "ObjectType_example", DeviceId: "DeviceId_example", NodeId: "NodeId_example", Signature: 123, TimeStamp: "TODO"}, Pid: []string{"Pid_example"), PlatformType: "PlatformType_example", PublicAccessKey: "PublicAccessKey_example", ReadOnly: false, Serial: []string{"Serial_example"), Vendor: "Vendor_example", Account: , ClaimedByUser: , ClusterMembers: []AssetClusterMemberRelationship{openapiclient.asset.ClusterMember.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ApiVersion: int64(123), AppPartitionNumber: int64(123), ConnectionId: "ConnectionId_example", ConnectionReason: "ConnectionReason_example", ConnectionStatus: "ConnectionStatus_example", ConnectionStatusLastChangeTime: "TODO", ConnectorVersion: "ConnectorVersion_example", DeviceExternalIpAddress: "DeviceExternalIpAddress_example", ProxyApp: "ProxyApp_example", Leadership: "Leadership_example", LockedLeader: false, MemberIdentity: "MemberIdentity_example", ParentClusterMemberIdentity: "ParentClusterMemberIdentity_example", Sudi: openapiclient.asset.SudiInfo{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Pid: "Pid_example", SerialNumber: "SerialNumber_example", Signature: "Signature_example", Status: "Status_example", SudiCertificate: openapiclient.x509.Certificate{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Issuer: , NotAfter: "TODO", NotBefore: "TODO", PemCertificate: "PemCertificate_example", Sha256Fingerprint: "Sha256Fingerprint_example", SignatureAlgorithm: "SignatureAlgorithm_example", Subject: }}, Device: }), DeviceClaim: openapiclient.asset.DeviceClaim.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceUpdates: []AssetConnectionControlMessage{openapiclient.asset.ConnectionControlMessage{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Account: "Account_example", ConnectorVersion: "ConnectorVersion_example", DeviceId: "DeviceId_example", DomainGroup: "DomainGroup_example", Evict: false, Leadership: "Leadership_example", NewIdentity: "NewIdentity_example", Partition: int64(123)}), SecurityToken: "SecurityToken_example", SerialNumber: "SerialNumber_example", Account: , Device: }, DeviceConfiguration: openapiclient.asset.DeviceConfiguration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, LocalConfigurationLocked: false, LogLevel: "LogLevel_example", Device: }, DomainGroup: , ParentConnection: }}), DeviceClaim: openapiclient.asset.DeviceClaim.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceUpdates: []AssetConnectionControlMessage{openapiclient.asset.ConnectionControlMessage{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Account: "Account_example", ConnectorVersion: "ConnectorVersion_example", DeviceId: "DeviceId_example", DomainGroup: "DomainGroup_example", Evict: false, Leadership: "Leadership_example", NewIdentity: "NewIdentity_example", Partition: int64(123)}), SecurityToken: "SecurityToken_example", SerialNumber: "SerialNumber_example", Account: , Device: }, DeviceConfiguration: openapiclient.asset.DeviceConfiguration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, LocalConfigurationLocked: false, LogLevel: "LogLevel_example", Device: }, DomainGroup: , ParentConnection: }}, Biosunits: []BiosUnitRelationship{openapiclient.bios.Unit.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", InitSeq: "InitSeq_example", InitTs: "InitTs_example", ComputeBlade: , ComputeRackUnit: , RegisteredDevice: , RunningFirmware: []FirmwareRunningFirmwareRelationship{openapiclient.firmware.RunningFirmware.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Component: "Component_example", PackageVersion: "PackageVersion_example", Type: "Type_example", Version: "Version_example", BiosUnit: openapiclient.bios.Unit.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", InitSeq: "InitSeq_example", InitTs: "InitTs_example", ComputeBlade: , ComputeRackUnit: , RegisteredDevice: , RunningFirmware: []FirmwareRunningFirmwareRelationship{openapiclient.firmware.RunningFirmware.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Component: "Component_example", PackageVersion: "PackageVersion_example", Type: "Type_example", Version: "Version_example", BiosUnit: , ManagementController: openapiclient.management.Controller.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", AdapterUnit: , ComputeBlade: , ComputeRackUnit: , ManagementInterfaces: []ManagementInterfaceRelationship{openapiclient.management.Interface.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Gateway: "Gateway_example", HostName: "HostName_example", IpAddress: "IpAddress_example", Ipv4Address: "Ipv4Address_example", Ipv4Gateway: "Ipv4Gateway_example", Ipv4Mask: "Ipv4Mask_example", Ipv6Address: "Ipv6Address_example", Ipv6Gateway: "Ipv6Gateway_example", Ipv6Prefix: int64(123), MacAddress: "MacAddress_example", Mask: "Mask_example", SwitchId: "SwitchId_example", UemConnStatus: "UemConnStatus_example", VirtualHostName: "VirtualHostName_example", ManagementController: openapiclient.management.Controller.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", AdapterUnit: , ComputeBlade: , ComputeRackUnit: , ManagementInterfaces: []ManagementInterfaceRelationship{openapiclient.management.Interface.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Gateway: "Gateway_example", HostName: "HostName_example", IpAddress: "IpAddress_example", Ipv4Address: "Ipv4Address_example", Ipv4Gateway: "Ipv4Gateway_example", Ipv4Mask: "Ipv4Mask_example", Ipv6Address: "Ipv6Address_example", Ipv6Gateway: "Ipv6Gateway_example", Ipv6Prefix: int64(123), MacAddress: "MacAddress_example", Mask: "Mask_example", SwitchId: "SwitchId_example", UemConnStatus: "UemConnStatus_example", VirtualHostName: "VirtualHostName_example", ManagementController: , RegisteredDevice: }), NetworkElement: openapiclient.network.Element.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminInbandInterfaceState: "AdminInbandInterfaceState_example", FaultSummary: int64(123), InbandIpAddress: "InbandIpAddress_example", InbandIpGateway: "InbandIpGateway_example", InbandIpMask: "InbandIpMask_example", InbandVlan: int64(123), OutOfBandIpAddress: "OutOfBandIpAddress_example", OutOfBandIpGateway: "OutOfBandIpGateway_example", OutOfBandIpMask: "OutOfBandIpMask_example", OutOfBandIpv4Address: "OutOfBandIpv4Address_example", OutOfBandIpv4Gateway: "OutOfBandIpv4Gateway_example", OutOfBandIpv4Mask: "OutOfBandIpv4Mask_example", OutOfBandIpv6Address: "OutOfBandIpv6Address_example", OutOfBandIpv6Gateway: "OutOfBandIpv6Gateway_example", OutOfBandIpv6Prefix: "OutOfBandIpv6Prefix_example", OutOfBandMac: "OutOfBandMac_example", SwitchId: "SwitchId_example", Cards: []EquipmentSwitchCardRelationship{openapiclient.equipment.SwitchCard.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", Description: "Description_example", NumPorts: int64(123), Presence: "Presence_example", SlotId: int64(123), State: "State_example", NetworkElement: openapiclient.network.Element.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminInbandInterfaceState: "AdminInbandInterfaceState_example", FaultSummary: int64(123), InbandIpAddress: "InbandIpAddress_example", InbandIpGateway: "InbandIpGateway_example", InbandIpMask: "InbandIpMask_example", InbandVlan: int64(123), OutOfBandIpAddress: "OutOfBandIpAddress_example", OutOfBandIpGateway: "OutOfBandIpGateway_example", OutOfBandIpMask: "OutOfBandIpMask_example", OutOfBandIpv4Address: "OutOfBandIpv4Address_example", OutOfBandIpv4Gateway: "OutOfBandIpv4Gateway_example", OutOfBandIpv4Mask: "OutOfBandIpv4Mask_example", OutOfBandIpv6Address: "OutOfBandIpv6Address_example", OutOfBandIpv6Gateway: "OutOfBandIpv6Gateway_example", OutOfBandIpv6Prefix: "OutOfBandIpv6Prefix_example", OutOfBandMac: "OutOfBandMac_example", SwitchId: "SwitchId_example", Cards: []EquipmentSwitchCardRelationship{openapiclient.equipment.SwitchCard.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", Description: "Description_example", NumPorts: int64(123), Presence: "Presence_example", SlotId: int64(123), State: "State_example", NetworkElement: , PortGroups: []PortGroupRelationship{openapiclient.port.Group.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Transport: "Transport_example", EquipmentSharedIoModule: openapiclient.equipment.SharedIoModule.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ConfigState: "ConfigState_example", Discovery: "Discovery_example", MacOfSharedIomAside: "MacOfSharedIomAside_example", MacOfSharedIomBside: "MacOfSharedIomBside_example", OperState: "OperState_example", PartNumber: "PartNumber_example", Reachability: "Reachability_example", UsrLbl: "UsrLbl_example", Vid: "Vid_example", EquipmentSystemIoController: openapiclient.equipment.SystemIoController.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ChassisId: "ChassisId_example", ConnectionPath: "ConnectionPath_example", ConnectionStatus: "ConnectionStatus_example", Description: "Description_example", ManagingInstance: "ManagingInstance_example", OperState: "OperState_example", PartNumber: "PartNumber_example", Pid: "Pid_example", SystemIoControllerId: int64(123), EquipmentChassis: openapiclient.equipment.Chassis.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ChassisId: int64(123), ConnectionPath: "ConnectionPath_example", ConnectionStatus: "ConnectionStatus_example", Description: "Description_example", FaultSummary: int64(123), Name: "Name_example", OperState: "OperState_example", PartNumber: "PartNumber_example", Pid: "Pid_example", PlatformType: "PlatformType_example", ProductName: "ProductName_example", Sku: "Sku_example", Vid: "Vid_example", Blades: []ComputeBladeRelationship{), Fanmodules: []EquipmentFanModuleRelationship{openapiclient.equipment.FanModule.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", Description: "Description_example", ModuleId: int64(123), OperState: "OperState_example", PartNumber: "PartNumber_example", Pid: "Pid_example", Presence: "Presence_example", Sku: "Sku_example", TrayId: int64(123), Vid: "Vid_example", ComputeRackUnit: , EquipmentChassis: openapiclient.equipment.Chassis.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ChassisId: int64(123), ConnectionPath: "ConnectionPath_example", ConnectionStatus: "ConnectionStatus_example", Description: "Description_example", FaultSummary: int64(123), Name: "Name_example", OperState: "OperState_example", PartNumber: "PartNumber_example", Pid: "Pid_example", PlatformType: "PlatformType_example", ProductName: "ProductName_example", Sku: "Sku_example", Vid: "Vid_example", Blades: []ComputeBladeRelationship{), Fanmodules: []EquipmentFanModuleRelationship{openapiclient.equipment.FanModule.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", Description: "Description_example", ModuleId: int64(123), OperState: "OperState_example", PartNumber: "PartNumber_example", Pid: "Pid_example", Presence: "Presence_example", Sku: "Sku_example", TrayId: int64(123), Vid: "Vid_example", ComputeRackUnit: , EquipmentChassis: , EquipmentRackEnclosure: openapiclient.equipment.RackEnclosure.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", EnclosureId: int64(123), Fanmodules: []EquipmentFanModuleRelationship{), Psus: []EquipmentPsuRelationship{openapiclient.equipment.Psu.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", Description: "Description_example", OperState: "OperState_example", PartNumber: "PartNumber_example", Pid: "Pid_example", Presence: "Presence_example", PsuFwVersion: "PsuFwVersion_example", PsuId: int64(123), PsuInputSrc: "PsuInputSrc_example", PsuType: "PsuType_example", PsuWattage: "PsuWattage_example", Sku: "Sku_example", Vid: "Vid_example", Voltage: "Voltage_example", ComputeRackUnit: , EquipmentChassis: , EquipmentRackEnclosure: openapiclient.equipment.RackEnclosure.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", EnclosureId: int64(123), Fanmodules: []EquipmentFanModuleRelationship{), Psus: []EquipmentPsuRelationship{openapiclient.equipment.Psu.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", Description: "Description_example", OperState: "OperState_example", PartNumber: "PartNumber_example", Pid: "Pid_example", Presence: "Presence_example", PsuFwVersion: "PsuFwVersion_example", PsuId: int64(123), PsuInputSrc: "PsuInputSrc_example", PsuType: "PsuType_example", PsuWattage: "PsuWattage_example", Sku: "Sku_example", Vid: "Vid_example", Voltage: "Voltage_example", ComputeRackUnit: , EquipmentChassis: , EquipmentRackEnclosure: , NetworkElement: , RegisteredDevice: }), RegisteredDevice: , Slots: []EquipmentRackEnclosureSlotRelationship{openapiclient.equipment.RackEnclosureSlot.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", RackId: int64(123), RackUnitDn: "RackUnitDn_example", EquipmentRackEnclosure: , RackUnit: , RegisteredDevice: })}, NetworkElement: , RegisteredDevice: }), RegisteredDevice: , Slots: []EquipmentRackEnclosureSlotRelationship{openapiclient.equipment.RackEnclosureSlot.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", RackId: int64(123), RackUnitDn: "RackUnitDn_example", EquipmentRackEnclosure: , RackUnit: , RegisteredDevice: })}, Fans: []EquipmentFanRelationship{openapiclient.equipment.Fan.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", Description: "Description_example", FanId: int64(123), FanModuleId: int64(123), ModuleId: int64(123), OperState: "OperState_example", PartNumber: "PartNumber_example", Pid: "Pid_example", Presence: "Presence_example", Sku: "Sku_example", TrayId: int64(123), Vid: "Vid_example", EquipmentFanModule: , RegisteredDevice: }), NetworkElement: , RegisteredDevice: }), Ioms: []EquipmentIoCardRelationship{openapiclient.equipment.IoCard.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", OperState: "OperState_example", EquipmentChassis: , RegisteredDevice: }), Psus: []EquipmentPsuRelationship{), RegisteredDevice: , Sasexpanders: []StorageSasExpanderRelationship{openapiclient.storage.SasExpander.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ExpanderId: int64(123), Name: "Name_example", OperState: "OperState_example", Operability: "Operability_example", Presence: "Presence_example", SasAddress: "SasAddress_example", Controller: , EquipmentChassis: , RegisteredDevice: }), Siocs: []EquipmentSystemIoControllerRelationship{openapiclient.equipment.SystemIoController.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ChassisId: "ChassisId_example", ConnectionPath: "ConnectionPath_example", ConnectionStatus: "ConnectionStatus_example", Description: "Description_example", ManagingInstance: "ManagingInstance_example", OperState: "OperState_example", PartNumber: "PartNumber_example", Pid: "Pid_example", SystemIoControllerId: int64(123), EquipmentChassis: , RegisteredDevice: , SharedIoModule: openapiclient.equipment.SharedIoModule.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ConfigState: "ConfigState_example", Discovery: "Discovery_example", MacOfSharedIomAside: "MacOfSharedIomAside_example", MacOfSharedIomBside: "MacOfSharedIomBside_example", OperState: "OperState_example", PartNumber: "PartNumber_example", Reachability: "Reachability_example", UsrLbl: "UsrLbl_example", Vid: "Vid_example", EquipmentSystemIoController: , PortGroups: []PortGroupRelationship{openapiclient.port.Group.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Transport: "Transport_example", EquipmentSharedIoModule: , EquipmentSwitchCard: , EthernetPorts: []EtherPhysicalPortRelationship{openapiclient.ether.PhysicalPort.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", OperState: "OperState_example", Role: "Role_example", MacAddress: "MacAddress_example", PeerDn: "PeerDn_example", TransceiverType: "TransceiverType_example", PortGroup: , PortSubGroup: openapiclient.port.SubGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Transport: "Transport_example", EthernetPorts: []EtherPhysicalPortRelationship{openapiclient.ether.PhysicalPort.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", OperState: "OperState_example", Role: "Role_example", MacAddress: "MacAddress_example", PeerDn: "PeerDn_example", TransceiverType: "TransceiverType_example", PortGroup: , PortSubGroup: openapiclient.port.SubGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Transport: "Transport_example", EthernetPorts: []EtherPhysicalPortRelationship{), PortGroup: , RegisteredDevice: }, RegisteredDevice: }), PortGroup: , RegisteredDevice: }, RegisteredDevice: }), FcPorts: []FcPhysicalPortRelationship{openapiclient.fc.PhysicalPort.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", OperState: "OperState_example", Role: "Role_example", PeerDn: "PeerDn_example", TransceiverType: "TransceiverType_example", Wwn: "Wwn_example", PortGroup: , RegisteredDevice: }), RegisteredDevice: , SubGroups: []PortSubGroupRelationship{)}), RegisteredDevice: }}), StorageEnclosures: []StorageEnclosureRelationship{openapiclient.storage.Enclosure.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ChassisId: int64(123), Description: "Description_example", EnclosureId: int64(123), NumSlots: int64(123), Presence: "Presence_example", ServerId: int64(123), Type: "Type_example", ComputeBlade: , ComputeRackUnit: , EnclosureDiskSlots: []StorageEnclosureDiskSlotEpRelationship{openapiclient.storage.EnclosureDiskSlotEp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", DrivePath: "DrivePath_example", Health: "Health_example", Presence: "Presence_example", Slot: "Slot_example", RegisteredDevice: , StorageEnclosure: openapiclient.storage.Enclosure.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ChassisId: int64(123), Description: "Description_example", EnclosureId: int64(123), NumSlots: int64(123), Presence: "Presence_example", ServerId: int64(123), Type: "Type_example", ComputeBlade: , ComputeRackUnit: , EnclosureDiskSlots: []StorageEnclosureDiskSlotEpRelationship{openapiclient.storage.EnclosureDiskSlotEp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", DrivePath: "DrivePath_example", Health: "Health_example", Presence: "Presence_example", Slot: "Slot_example", RegisteredDevice: , StorageEnclosure: }), EnclosureDisks: []StorageEnclosureDiskRelationship{openapiclient.storage.EnclosureDisk.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", BlockSize: "BlockSize_example", DiskId: "DiskId_example", DiskState: "DiskState_example", Health: "Health_example", NumBlocks: "NumBlocks_example", Pid: "Pid_example", SasAddress1: "SasAddress1_example", SasAddress2: "SasAddress2_example", Size: "Size_example", PhysicalDisk: openapiclient.storage.PhysicalDisk.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", BlockSize: "BlockSize_example", Bootable: "Bootable_example", ConfigurationCheckpoint: "ConfigurationCheckpoint_example", ConfigurationState: "ConfigurationState_example", DiscoveredPath: "DiscoveredPath_example", DiskId: "DiskId_example", DiskState: "DiskState_example", DriveFirmware: "DriveFirmware_example", DriveState: "DriveState_example", FdeCapable: "FdeCapable_example", LinkSpeed: "LinkSpeed_example", LinkState: "LinkState_example", NumBlocks: "NumBlocks_example", OperPowerState: "OperPowerState_example", OperQualifierReason: "OperQualifierReason_example", Operability: "Operability_example", PhysicalBlockSize: "PhysicalBlockSize_example", Pid: "Pid_example", Presence: "Presence_example", Protocol: "Protocol_example", RawSize: "RawSize_example", Secured: "Secured_example", Size: "Size_example", Thermal: "Thermal_example", Type: "Type_example", VariantType: "VariantType_example", LocatorLed: openapiclient.equipment.LocatorLed.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Color: "Color_example", OperState: "OperState_example", ComputeBlade: , ComputeRackUnit: , RegisteredDevice: , StoragePhysicalDisk: openapiclient.storage.PhysicalDisk.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", BlockSize: "BlockSize_example", Bootable: "Bootable_example", ConfigurationCheckpoint: "ConfigurationCheckpoint_example", ConfigurationState: "ConfigurationState_example", DiscoveredPath: "DiscoveredPath_example", DiskId: "DiskId_example", DiskState: "DiskState_example", DriveFirmware: "DriveFirmware_example", DriveState: "DriveState_example", FdeCapable: "FdeCapable_example", LinkSpeed: "LinkSpeed_example", LinkState: "LinkState_example", NumBlocks: "NumBlocks_example", OperPowerState: "OperPowerState_example", OperQualifierReason: "OperQualifierReason_example", Operability: "Operability_example", PhysicalBlockSize: "PhysicalBlockSize_example", Pid: "Pid_example", Presence: "Presence_example", Protocol: "Protocol_example", RawSize: "RawSize_example", Secured: "Secured_example", Size: "Size_example", Thermal: "Thermal_example", Type: "Type_example", VariantType: "VariantType_example", LocatorLed: openapiclient.equipment.LocatorLed.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Color: "Color_example", OperState: "OperState_example", ComputeBlade: , ComputeRackUnit: , RegisteredDevice: , StoragePhysicalDisk: }, PhysicalDiskExtensions: []StoragePhysicalDiskExtensionRelationship{openapiclient.storage.PhysicalDiskExtension.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", Bootable: "Bootable_example", DiskDn: "DiskDn_example", DiskId: int64(123), DiskState: "DiskState_example", Health: "Health_example", PhysicalDisk: , RegisteredDevice: , StorageController: openapiclient.storage.Controller.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ControllerFlags: "ControllerFlags_example", ControllerId: "ControllerId_example", ControllerStatus: "ControllerStatus_example", HwRevision: "HwRevision_example", OobInterfaceSupported: "OobInterfaceSupported_example", OperState: "OperState_example", Operability: "Operability_example", PciAddr: "PciAddr_example", PciSlot: "PciSlot_example", Presence: "Presence_example", RaidSupport: "RaidSupport_example", RebuildRate: "RebuildRate_example", SelfEncryptEnabled: "SelfEncryptEnabled_example", Type: "Type_example", ComputeBoard: openapiclient.compute.Board.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", BoardId: int64(123), CpuTypeController: "CpuTypeController_example", OperPowerState: "OperPowerState_example", Presence: "Presence_example", ComputeBlade: , ComputeRackUnit: , EquipmentTpms: []EquipmentTpmRelationship{openapiclient.equipment.Tpm.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ActivationStatus: "ActivationStatus_example", AdminState: "AdminState_example", Ownership: "Ownership_example", Presence: "Presence_example", TpmId: int64(123), Version: "Version_example", ComputeBoard: , RegisteredDevice: }), GraphicsCards: []GraphicsCardRelationship{openapiclient.graphics.Card.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", CardId: int64(123), DeviceId: int64(123), ExpanderSlot: "ExpanderSlot_example", FirmwareVersion: "FirmwareVersion_example", Mode: "Mode_example", NumGpus: "NumGpus_example", OperState: "OperState_example", PciAddress: "PciAddress_example", PciAddressList: "PciAddressList_example", PciSlot: "PciSlot_example", ComputeBoard: , GraphicsControllers: []GraphicsControllerRelationship{openapiclient.graphics.Controller.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ControllerId: int64(123), PciAddr: "PciAddr_example", PciSlot: "PciSlot_example", GraphicsCard: openapiclient.graphics.Card.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", CardId: int64(123), DeviceId: int64(123), ExpanderSlot: "ExpanderSlot_example", FirmwareVersion: "FirmwareVersion_example", Mode: "Mode_example", NumGpus: "NumGpus_example", OperState: "OperState_example", PciAddress: "PciAddress_example", PciAddressList: "PciAddressList_example", PciSlot: "PciSlot_example", ComputeBoard: , GraphicsControllers: []GraphicsControllerRelationship{openapiclient.graphics.Controller.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ControllerId: int64(123), PciAddr: "PciAddr_example", PciSlot: "PciSlot_example", GraphicsCard: , RegisteredDevice: }), RegisteredDevice: }, RegisteredDevice: }), RegisteredDevice: }), MemoryArrays: []MemoryArrayRelationship{openapiclient.memory.Array.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ArrayId: int64(123), CpuId: int64(123), CurrentCapacity: "CurrentCapacity_example", ErrorCorrection: "ErrorCorrection_example", MaxCapacity: "MaxCapacity_example", MaxDevices: "MaxDevices_example", OperPowerState: "OperPowerState_example", Presence: "Presence_example", ComputeBoard: , PersistentMemoryUnits: []MemoryPersistentMemoryUnitRelationship{openapiclient.memory.PersistentMemoryUnit.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", ArrayId: int64(123), Bank: int64(123), Capacity: "Capacity_example", Clock: "Clock_example", FormFactor: "FormFactor_example", Latency: "Latency_example", Location: "Location_example", OperPowerState: "OperPowerState_example", OperState: "OperState_example", Operability: "Operability_example", Presence: "Presence_example", Set: int64(123), Speed: "Speed_example", Thermal: "Thermal_example", Type: "Type_example", Visibility: "Visibility_example", Width: "Width_example", AppDirectCapacity: "AppDirectCapacity_example", CountStatus: "CountStatus_example", FirmwareVersion: "FirmwareVersion_example", FrozenStatus: "FrozenStatus_example", HealthState: "HealthState_example", LockStatus: "LockStatus_example", MemoryCapacity: "MemoryCapacity_example", MemoryId: int64(123), PersistentMemoryCapacity: "PersistentMemoryCapacity_example", ReservedCapacity: "ReservedCapacity_example", SecurityStatus: "SecurityStatus_example", SocketId: "SocketId_example", SocketMemoryId: "SocketMemoryId_example", TotalCapacity: "TotalCapacity_example", Uid: "Uid_example", MemoryArray: openapiclient.memory.Array.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ArrayId: int64(123), CpuId: int64(123), CurrentCapacity: "CurrentCapacity_example", ErrorCorrection: "ErrorCorrection_example", MaxCapacity: "MaxCapacity_example", MaxDevices: "MaxDevices_example", OperPowerState: "OperPowerState_example", Presence: "Presence_example", ComputeBoard: , PersistentMemoryUnits: []MemoryPersistentMemoryUnitRelationship{openapiclient.memory.PersistentMemoryUnit.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", ArrayId: int64(123), Bank: int64(123), Capacity: "Capacity_example", Clock: "Clock_example", FormFactor: "FormFactor_example", Latency: "Latency_example", Location: "Location_example", OperPowerState: "OperPowerState_example", OperState: "OperState_example", Operability: "Operability_example", Presence: "Presence_example", Set: int64(123), Speed: "Speed_example", Thermal: "Thermal_example", Type: "Type_example", Visibility: "Visibility_example", Width: "Width_example", AppDirectCapacity: "AppDirectCapacity_example", CountStatus: "CountStatus_example", FirmwareVersion: "FirmwareVersion_example", FrozenStatus: "FrozenStatus_example", HealthState: "HealthState_example", LockStatus: "LockStatus_example", MemoryCapacity: "MemoryCapacity_example", MemoryId: int64(123), PersistentMemoryCapacity: "PersistentMemoryCapacity_example", ReservedCapacity: "ReservedCapacity_example", SecurityStatus: "SecurityStatus_example", SocketId: "SocketId_example", SocketMemoryId: "SocketMemoryId_example", TotalCapacity: "TotalCapacity_example", Uid: "Uid_example", MemoryArray: , RegisteredDevice: }), RegisteredDevice: , Units: []MemoryUnitRelationship{openapiclient.memory.Unit.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", ArrayId: int64(123), Bank: int64(123), Capacity: "Capacity_example", Clock: "Clock_example", FormFactor: "FormFactor_example", Latency: "Latency_example", Location: "Location_example", OperPowerState: "OperPowerState_example", OperState: "OperState_example", Operability: "Operability_example", Presence: "Presence_example", Set: int64(123), Speed: "Speed_example", Thermal: "Thermal_example", Type: "Type_example", Visibility: "Visibility_example", Width: "Width_example", MemoryId: int64(123), MemoryArray: , RegisteredDevice: })}, RegisteredDevice: }), RegisteredDevice: , Units: []MemoryUnitRelationship{openapiclient.memory.Unit.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", ArrayId: int64(123), Bank: int64(123), Capacity: "Capacity_example", Clock: "Clock_example", FormFactor: "FormFactor_example", Latency: "Latency_example", Location: "Location_example", OperPowerState: "OperPowerState_example", OperState: "OperState_example", Operability: "Operability_example", Presence: "Presence_example", Set: int64(123), Speed: "Speed_example", Thermal: "Thermal_example", Type: "Type_example", Visibility: "Visibility_example", Width: "Width_example", MemoryId: int64(123), MemoryArray: , RegisteredDevice: })}), PciCoprocessorCards: []PciCoprocessorCardRelationship{openapiclient.pci.CoprocessorCard.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", CardId: int64(123), PciSlot: "PciSlot_example", ComputeBoard: , RegisteredDevice: }), PciSwitch: []PciSwitchRelationship{openapiclient.pci.Switch.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", DeviceId: "DeviceId_example", Health: "Health_example", NumOfAdaptors: "NumOfAdaptors_example", PciAddress: "PciAddress_example", PciSlot: "PciSlot_example", ProductName: "ProductName_example", ProductRevision: "ProductRevision_example", SubDeviceId: "SubDeviceId_example", SubVendorId: "SubVendorId_example", Temperature: "Temperature_example", Type: "Type_example", VendorId: "VendorId_example", ComputeBoard: , Links: []PciLinkRelationship{openapiclient.pci.Link.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", Adapter: "Adapter_example", LinkSpeed: "LinkSpeed_example", LinkStatus: "LinkStatus_example", LinkWidth: "LinkWidth_example", PciSlot: "PciSlot_example", SlotStatus: "SlotStatus_example", PciSwitch: openapiclient.pci.Switch.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", DeviceId: "DeviceId_example", Health: "Health_example", NumOfAdaptors: "NumOfAdaptors_example", PciAddress: "PciAddress_example", PciSlot: "PciSlot_example", ProductName: "ProductName_example", ProductRevision: "ProductRevision_example", SubDeviceId: "SubDeviceId_example", SubVendorId: "SubVendorId_example", Temperature: "Temperature_example", Type: "Type_example", VendorId: "VendorId_example", ComputeBoard: , Links: []PciLinkRelationship{openapiclient.pci.Link.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", Adapter: "Adapter_example", LinkSpeed: "LinkSpeed_example", LinkStatus: "LinkStatus_example", LinkWidth: "LinkWidth_example", PciSlot: "PciSlot_example", SlotStatus: "SlotStatus_example", PciSwitch: , RegisteredDevice: }), RegisteredDevice: }, RegisteredDevice: }), RegisteredDevice: }), PersistentMemoryConfiguration: openapiclient.memory.PersistentMemoryConfiguration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", MemoryCapacity: "MemoryCapacity_example", NumOfModules: "NumOfModules_example", NumOfRegions: "NumOfRegions_example", PersistentMemoryCapacity: "PersistentMemoryCapacity_example", ReservedCapacity: "ReservedCapacity_example", SecurityState: "SecurityState_example", TotalCapacity: "TotalCapacity_example", ComputeBoard: , PersistentMemoryConfigResult: openapiclient.memory.PersistentMemoryConfigResult.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", ConfigErrorDesc: "ConfigErrorDesc_example", ConfigResult: "ConfigResult_example", ConfigSequenceNo: int64(123), ConfigState: "ConfigState_example", MemoryPersistentMemoryConfiguration: openapiclient.memory.PersistentMemoryConfiguration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", MemoryCapacity: "MemoryCapacity_example", NumOfModules: "NumOfModules_example", NumOfRegions: "NumOfRegions_example", PersistentMemoryCapacity: "PersistentMemoryCapacity_example", ReservedCapacity: "ReservedCapacity_example", SecurityState: "SecurityState_example", TotalCapacity: "TotalCapacity_example", ComputeBoard: , PersistentMemoryConfigResult: openapiclient.memory.PersistentMemoryConfigResult.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", ConfigErrorDesc: "ConfigErrorDesc_example", ConfigResult: "ConfigResult_example", ConfigSequenceNo: int64(123), ConfigState: "ConfigState_example", MemoryPersistentMemoryConfiguration: , PersistentMemoryNamespaceConfigResults: []MemoryPersistentMemoryNamespaceConfigResultRelationship{openapiclient.memory.PersistentMemoryNamespaceConfigResult.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", ConfigStatus: "ConfigStatus_example", Name: "Name_example", SocketId: "SocketId_example", SocketMemoryId: "SocketMemoryId_example", MemoryPersistentMemoryConfigResult: , RegisteredDevice: }), RegisteredDevice: }, PersistentMemoryRegions: []MemoryPersistentMemoryRegionRelationship{openapiclient.memory.PersistentMemoryRegion.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", FreeCapacity: "FreeCapacity_example", HealthState: "HealthState_example", InterleavedSetId: "InterleavedSetId_example", LocaterIds: "LocaterIds_example", PersistentMemoryType: "PersistentMemoryType_example", RegionId: "RegionId_example", SocketId: "SocketId_example", SocketMemoryId: "SocketMemoryId_example", TotalCapacity: "TotalCapacity_example", MemoryPersistentMemoryConfiguration: , PersistentMemoryNamespaces: []MemoryPersistentMemoryNamespaceRelationship{openapiclient.memory.PersistentMemoryNamespace.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Capacity: "Capacity_example", HealthState: "HealthState_example", LabelVersion: "LabelVersion_example", Mode: "Mode_example", Name: "Name_example", Uuid: "Uuid_example", MemoryPersistentMemoryRegion: openapiclient.memory.PersistentMemoryRegion.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", FreeCapacity: "FreeCapacity_example", HealthState: "HealthState_example", InterleavedSetId: "InterleavedSetId_example", LocaterIds: "LocaterIds_example", PersistentMemoryType: "PersistentMemoryType_example", RegionId: "RegionId_example", SocketId: "SocketId_example", SocketMemoryId: "SocketMemoryId_example", TotalCapacity: "TotalCapacity_example", MemoryPersistentMemoryConfiguration: , PersistentMemoryNamespaces: []MemoryPersistentMemoryNamespaceRelationship{openapiclient.memory.PersistentMemoryNamespace.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Capacity: "Capacity_example", HealthState: "HealthState_example", LabelVersion: "LabelVersion_example", Mode: "Mode_example", Name: "Name_example", Uuid: "Uuid_example", MemoryPersistentMemoryRegion: , RegisteredDevice: }), RegisteredDevice: }, RegisteredDevice: }), RegisteredDevice: }), RegisteredDevice: }, PersistentMemoryNamespaceConfigResults: []MemoryPersistentMemoryNamespaceConfigResultRelationship{openapiclient.memory.PersistentMemoryNamespaceConfigResult.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", ConfigStatus: "ConfigStatus_example", Name: "Name_example", SocketId: "SocketId_example", SocketMemoryId: "SocketMemoryId_example", MemoryPersistentMemoryConfigResult: , RegisteredDevice: }), RegisteredDevice: }, PersistentMemoryRegions: []MemoryPersistentMemoryRegionRelationship{), RegisteredDevice: }, Processors: []ProcessorUnitRelationship{openapiclient.processor.Unit.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", Architecture: "Architecture_example", NumCores: int64(123), NumCoresEnabled: "NumCoresEnabled_example", NumThreads: "NumThreads_example", OperPowerState: "OperPowerState_example", OperState: "OperState_example", Operability: "Operability_example", Presence: "Presence_example", ProcessorId: int64(123), SocketDesignation: "SocketDesignation_example", Speed: 123, Stepping: "Stepping_example", Thermal: "Thermal_example", ComputeBoard: , RegisteredDevice: }), RegisteredDevice: , SecurityUnits: []SecurityUnitRelationship{openapiclient.security.Unit.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", OperState: "OperState_example", Operability: "Operability_example", PartNumber: "PartNumber_example", PciSlot: "PciSlot_example", Power: "Power_example", Presence: "Presence_example", Thermal: "Thermal_example", UnitId: int64(123), Vid: "Vid_example", Voltage: "Voltage_example", ComputeBoard: , RegisteredDevice: }), StorageControllers: []StorageControllerRelationship{openapiclient.storage.Controller.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ControllerFlags: "ControllerFlags_example", ControllerId: "ControllerId_example", ControllerStatus: "ControllerStatus_example", HwRevision: "HwRevision_example", OobInterfaceSupported: "OobInterfaceSupported_example", OperState: "OperState_example", Operability: "Operability_example", PciAddr: "PciAddr_example", PciSlot: "PciSlot_example", Presence: "Presence_example", RaidSupport: "RaidSupport_example", RebuildRate: "RebuildRate_example", SelfEncryptEnabled: "SelfEncryptEnabled_example", Type: "Type_example", ComputeBoard: , PhysicalDiskExtensions: []StoragePhysicalDiskExtensionRelationship{openapiclient.storage.PhysicalDiskExtension.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", Bootable: "Bootable_example", DiskDn: "DiskDn_example", DiskId: int64(123), DiskState: "DiskState_example", Health: "Health_example", PhysicalDisk: , RegisteredDevice: , StorageController: }), PhysicalDisks: []StoragePhysicalDiskRelationship{), RegisteredDevice: , RunningFirmware: []FirmwareRunningFirmwareRelationship{), VirtualDriveExtensions: []StorageVirtualDriveExtensionRelationship{openapiclient.storage.VirtualDriveExtension.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Bootable: "Bootable_example", ContainerId: int64(123), DriveState: "DriveState_example", Name: "Name_example", OperDeviceId: "OperDeviceId_example", Uuid: "Uuid_example", VendorUuid: "VendorUuid_example", VirtualDriveDn: "VirtualDriveDn_example", VirtualDriveId: "VirtualDriveId_example", RegisteredDevice: , StorageController: , VirtualDrive: openapiclient.storage.VirtualDrive.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AccessPolicy: "AccessPolicy_example", ActualWriteCachePolicy: "ActualWriteCachePolicy_example", AvailableSize: "AvailableSize_example", BlockSize: "BlockSize_example", Bootable: "Bootable_example", ConfigState: "ConfigState_example", ConfiguredWriteCachePolicy: "ConfiguredWriteCachePolicy_example", ConnectionProtocol: "ConnectionProtocol_example", DriveCache: "DriveCache_example", DriveSecurity: "DriveSecurity_example", DriveState: "DriveState_example", IoPolicy: "IoPolicy_example", Name: "Name_example", NumBlocks: "NumBlocks_example", OperState: "OperState_example", Operability: "Operability_example", PhysicalBlockSize: "PhysicalBlockSize_example", Presence: "Presence_example", ReadPolicy: "ReadPolicy_example", SecurityFlags: "SecurityFlags_example", Size: "Size_example", StripSize: "StripSize_example", Type: "Type_example", Uuid: "Uuid_example", VendorUuid: "VendorUuid_example", VirtualDriveId: "VirtualDriveId_example", PhysicalDiskUsages: []StoragePhysicalDiskUsageRelationship{openapiclient.storage.PhysicalDiskUsage.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", NumberOfBlocks: "NumberOfBlocks_example", PhysicalDrive: "PhysicalDrive_example", Span: "Span_example", StartingBlock: "StartingBlock_example", State: "State_example", VirtualDrive: "VirtualDrive_example", RegisteredDevice: }), RegisteredDevice: , StorageController: , VdMemberEps: []StorageVdMemberEpRelationship{openapiclient.storage.VdMemberEp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", OperQualifierReason: "OperQualifierReason_example", Presence: "Presence_example", Role: "Role_example", SpanId: "SpanId_example", VdMemberEpId: int64(123), RegisteredDevice: , StorageVirtualDrive: openapiclient.storage.VirtualDrive.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AccessPolicy: "AccessPolicy_example", ActualWriteCachePolicy: "ActualWriteCachePolicy_example", AvailableSize: "AvailableSize_example", BlockSize: "BlockSize_example", Bootable: "Bootable_example", ConfigState: "ConfigState_example", ConfiguredWriteCachePolicy: "ConfiguredWriteCachePolicy_example", ConnectionProtocol: "ConnectionProtocol_example", DriveCache: "DriveCache_example", DriveSecurity: "DriveSecurity_example", DriveState: "DriveState_example", IoPolicy: "IoPolicy_example", Name: "Name_example", NumBlocks: "NumBlocks_example", OperState: "OperState_example", Operability: "Operability_example", PhysicalBlockSize: "PhysicalBlockSize_example", Presence: "Presence_example", ReadPolicy: "ReadPolicy_example", SecurityFlags: "SecurityFlags_example", Size: "Size_example", StripSize: "StripSize_example", Type: "Type_example", Uuid: "Uuid_example", VendorUuid: "VendorUuid_example", VirtualDriveId: "VirtualDriveId_example", PhysicalDiskUsages: []StoragePhysicalDiskUsageRelationship{openapiclient.storage.PhysicalDiskUsage.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", NumberOfBlocks: "NumberOfBlocks_example", PhysicalDrive: "PhysicalDrive_example", Span: "Span_example", StartingBlock: "StartingBlock_example", State: "State_example", VirtualDrive: "VirtualDrive_example", RegisteredDevice: }), RegisteredDevice: , StorageController: , VdMemberEps: []StorageVdMemberEpRelationship{openapiclient.storage.VdMemberEp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", OperQualifierReason: "OperQualifierReason_example", Presence: "Presence_example", Role: "Role_example", SpanId: "SpanId_example", VdMemberEpId: int64(123), RegisteredDevice: , StorageVirtualDrive: }), VirtualDriveExtension: openapiclient.storage.VirtualDriveExtension.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Bootable: "Bootable_example", ContainerId: int64(123), DriveState: "DriveState_example", Name: "Name_example", OperDeviceId: "OperDeviceId_example", Uuid: "Uuid_example", VendorUuid: "VendorUuid_example", VirtualDriveDn: "VirtualDriveDn_example", VirtualDriveId: "VirtualDriveId_example", RegisteredDevice: , StorageController: , VirtualDrive: }}}), VirtualDriveExtension: }}), VirtualDrives: []StorageVirtualDriveRelationship{)}), StorageFlexFlashControllers: []StorageFlexFlashControllerRelationship{openapiclient.storage.FlexFlashController.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ControllerState: "ControllerState_example", FfControllerId: "FfControllerId_example", ComputeBoard: , FlexFlashControllerProps: []StorageFlexFlashControllerPropsRelationship{openapiclient.storage.FlexFlashControllerProps.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", CardsManageable: "CardsManageable_example", ConfiguredMode: "ConfiguredMode_example", ControllerName: "ControllerName_example", ControllerStatus: "ControllerStatus_example", FwVersion: "FwVersion_example", InternalState: "InternalState_example", OperatingMode: "OperatingMode_example", PhysicalDriveCount: "PhysicalDriveCount_example", ProductName: "ProductName_example", StartupFwVersion: "StartupFwVersion_example", VirtualDriveCount: "VirtualDriveCount_example", RegisteredDevice: , StorageFlexFlashController: openapiclient.storage.FlexFlashController.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ControllerState: "ControllerState_example", FfControllerId: "FfControllerId_example", ComputeBoard: , FlexFlashControllerProps: []StorageFlexFlashControllerPropsRelationship{openapiclient.storage.FlexFlashControllerProps.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", CardsManageable: "CardsManageable_example", ConfiguredMode: "ConfiguredMode_example", ControllerName: "ControllerName_example", ControllerStatus: "ControllerStatus_example", FwVersion: "FwVersion_example", InternalState: "InternalState_example", OperatingMode: "OperatingMode_example", PhysicalDriveCount: "PhysicalDriveCount_example", ProductName: "ProductName_example", StartupFwVersion: "StartupFwVersion_example", VirtualDriveCount: "VirtualDriveCount_example", RegisteredDevice: , StorageFlexFlashController: }), FlexFlashPhysicalDrives: []StorageFlexFlashPhysicalDriveRelationship{openapiclient.storage.FlexFlashPhysicalDrive.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", CardStatus: "CardStatus_example", CardType: "CardType_example", OemId: "OemId_example", PdStatus: "PdStatus_example", RegisteredDevice: , StorageFlexFlashController: }), FlexFlashVirtualDrives: []StorageFlexFlashVirtualDriveRelationship{openapiclient.storage.FlexFlashVirtualDrive.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", DriveScope: "DriveScope_example", DriveStatus: "DriveStatus_example", PartitionId: "PartitionId_example", ResidentImage: "ResidentImage_example", Size: "Size_example", VirtualDrive: "VirtualDrive_example", RegisteredDevice: , StorageFlexFlashController: }), RegisteredDevice: }}), FlexFlashPhysicalDrives: []StorageFlexFlashPhysicalDriveRelationship{openapiclient.storage.FlexFlashPhysicalDrive.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", CardStatus: "CardStatus_example", CardType: "CardType_example", OemId: "OemId_example", PdStatus: "PdStatus_example", RegisteredDevice: , StorageFlexFlashController: }), FlexFlashVirtualDrives: []StorageFlexFlashVirtualDriveRelationship{openapiclient.storage.FlexFlashVirtualDrive.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", DriveScope: "DriveScope_example", DriveStatus: "DriveStatus_example", PartitionId: "PartitionId_example", ResidentImage: "ResidentImage_example", Size: "Size_example", VirtualDrive: "VirtualDrive_example", RegisteredDevice: , StorageFlexFlashController: }), RegisteredDevice: }), StorageFlexUtilControllers: []StorageFlexUtilControllerRelationship{openapiclient.storage.FlexUtilController.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", ControllerName: "ControllerName_example", ControllerStatus: "ControllerStatus_example", FfControllerId: "FfControllerId_example", InternalState: "InternalState_example", ComputeBoard: , FlexUtilPhysicalDrives: []StorageFlexUtilPhysicalDriveRelationship{openapiclient.storage.FlexUtilPhysicalDrive.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", BlockSize: "BlockSize_example", Capacity: "Capacity_example", Controller: "Controller_example", DrivesEnabled: "DrivesEnabled_example", Health: "Health_example", ManufacturerDate: "ManufacturerDate_example", ManufacturerId: "ManufacturerId_example", OemId: "OemId_example", PartitionCount: "PartitionCount_example", PdStatus: "PdStatus_example", PhysicalDrive: "PhysicalDrive_example", ProductName: "ProductName_example", ProductRevision: "ProductRevision_example", ReadErrorCount: "ReadErrorCount_example", ReadErrorThreshold: "ReadErrorThreshold_example", WriteEnabled: "WriteEnabled_example", WriteErrorCount: "WriteErrorCount_example", WriteErrorThreshold: "WriteErrorThreshold_example", RegisteredDevice: , StorageFlexUtilController: openapiclient.storage.FlexUtilController.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", ControllerName: "ControllerName_example", ControllerStatus: "ControllerStatus_example", FfControllerId: "FfControllerId_example", InternalState: "InternalState_example", ComputeBoard: , FlexUtilPhysicalDrives: []StorageFlexUtilPhysicalDriveRelationship{openapiclient.storage.FlexUtilPhysicalDrive.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", BlockSize: "BlockSize_example", Capacity: "Capacity_example", Controller: "Controller_example", DrivesEnabled: "DrivesEnabled_example", Health: "Health_example", ManufacturerDate: "ManufacturerDate_example", ManufacturerId: "ManufacturerId_example", OemId: "OemId_example", PartitionCount: "PartitionCount_example", PdStatus: "PdStatus_example", PhysicalDrive: "PhysicalDrive_example", ProductName: "ProductName_example", ProductRevision: "ProductRevision_example", ReadErrorCount: "ReadErrorCount_example", ReadErrorThreshold: "ReadErrorThreshold_example", WriteEnabled: "WriteEnabled_example", WriteErrorCount: "WriteErrorCount_example", WriteErrorThreshold: "WriteErrorThreshold_example", RegisteredDevice: , StorageFlexUtilController: }), FlexUtilVirtualDrives: []StorageFlexUtilVirtualDriveRelationship{openapiclient.storage.FlexUtilVirtualDrive.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", DriveStatus: "DriveStatus_example", DriveType: "DriveType_example", PartitionId: "PartitionId_example", PartitionName: "PartitionName_example", ResidentImage: "ResidentImage_example", Size: "Size_example", VirtualDrive: "VirtualDrive_example", RegisteredDevice: , StorageFlexUtilController: }), RegisteredDevice: }}), FlexUtilVirtualDrives: []StorageFlexUtilVirtualDriveRelationship{openapiclient.storage.FlexUtilVirtualDrive.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", DriveStatus: "DriveStatus_example", DriveType: "DriveType_example", PartitionId: "PartitionId_example", PartitionName: "PartitionName_example", ResidentImage: "ResidentImage_example", Size: "Size_example", VirtualDrive: "VirtualDrive_example", RegisteredDevice: , StorageFlexUtilController: }), RegisteredDevice: })}, PhysicalDiskExtensions: []StoragePhysicalDiskExtensionRelationship{), PhysicalDisks: []StoragePhysicalDiskRelationship{), RegisteredDevice: , RunningFirmware: []FirmwareRunningFirmwareRelationship{), VirtualDriveExtensions: []StorageVirtualDriveExtensionRelationship{), VirtualDrives: []StorageVirtualDriveRelationship{)}}), RegisteredDevice: , RunningFirmware: []FirmwareRunningFirmwareRelationship{), SasPorts: []StorageSasPortRelationship{openapiclient.storage.SasPort.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Address: "Address_example", DiskId: int64(123), EndPointId: int64(123), LinkDescription: "LinkDescription_example", LinkSpeed: "LinkSpeed_example", RegisteredDevice: , StoragePhysicalDisk: }), StorageController: , StorageEnclosure: }}, PhysicalDiskExtensions: []StoragePhysicalDiskExtensionRelationship{), RegisteredDevice: , RunningFirmware: []FirmwareRunningFirmwareRelationship{), SasPorts: []StorageSasPortRelationship{openapiclient.storage.SasPort.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Address: "Address_example", DiskId: int64(123), EndPointId: int64(123), LinkDescription: "LinkDescription_example", LinkSpeed: "LinkSpeed_example", RegisteredDevice: , StoragePhysicalDisk: }), StorageController: , StorageEnclosure: }, RegisteredDevice: , StorageEnclosure: }), EquipmentChassis: , PhysicalDisks: []StoragePhysicalDiskRelationship{), RegisteredDevice: }}), EnclosureDisks: []StorageEnclosureDiskRelationship{openapiclient.storage.EnclosureDisk.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", BlockSize: "BlockSize_example", DiskId: "DiskId_example", DiskState: "DiskState_example", Health: "Health_example", NumBlocks: "NumBlocks_example", Pid: "Pid_example", SasAddress1: "SasAddress1_example", SasAddress2: "SasAddress2_example", Size: "Size_example", PhysicalDisk: , RegisteredDevice: , StorageEnclosure: }), EquipmentChassis: , PhysicalDisks: []StoragePhysicalDiskRelationship{), RegisteredDevice: })}, EquipmentRackEnclosure: , Fans: []EquipmentFanRelationship{openapiclient.equipment.Fan.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", Description: "Description_example", FanId: int64(123), FanModuleId: int64(123), ModuleId: int64(123), OperState: "OperState_example", PartNumber: "PartNumber_example", Pid: "Pid_example", Presence: "Presence_example", Sku: "Sku_example", TrayId: int64(123), Vid: "Vid_example", EquipmentFanModule: , RegisteredDevice: }), NetworkElement: , RegisteredDevice: }), Ioms: []EquipmentIoCardRelationship{openapiclient.equipment.IoCard.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", OperState: "OperState_example", EquipmentChassis: , RegisteredDevice: }), Psus: []EquipmentPsuRelationship{), RegisteredDevice: , Sasexpanders: []StorageSasExpanderRelationship{openapiclient.storage.SasExpander.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ExpanderId: int64(123), Name: "Name_example", OperState: "OperState_example", Operability: "Operability_example", Presence: "Presence_example", SasAddress: "SasAddress_example", Controller: , EquipmentChassis: , RegisteredDevice: }), Siocs: []EquipmentSystemIoControllerRelationship{), StorageEnclosures: []StorageEnclosureRelationship{)}, RegisteredDevice: , SharedIoModule: }, PortGroups: []PortGroupRelationship{), RegisteredDevice: }, EquipmentSwitchCard: , EthernetPorts: []EtherPhysicalPortRelationship{), FcPorts: []FcPhysicalPortRelationship{openapiclient.fc.PhysicalPort.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", OperState: "OperState_example", Role: "Role_example", PeerDn: "PeerDn_example", TransceiverType: "TransceiverType_example", Wwn: "Wwn_example", PortGroup: , RegisteredDevice: }), RegisteredDevice: , SubGroups: []PortSubGroupRelationship{)}), RegisteredDevice: }), Fanmodules: []EquipmentFanModuleRelationship{), ManagementContoller: , ManagementEntity: openapiclient.management.Entity.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", EntityId: "EntityId_example", Leadership: "Leadership_example", NetworkElement: , RegisteredDevice: }, Psus: []EquipmentPsuRelationship{), RegisteredDevice: , TopSystem: openapiclient.top.System.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Ipv4Address: "Ipv4Address_example", Ipv6Address: "Ipv6Address_example", Mode: "Mode_example", Name: "Name_example", TimeZone: "TimeZone_example", ComputeBlades: []ComputeBladeRelationship{), ComputeRackUnits: []ComputeRackUnitRelationship{), ManagementController: , NetworkElements: []NetworkElementRelationship{), RegisteredDevice: }, UcsmRunningFirmware: }, PortGroups: []PortGroupRelationship{), RegisteredDevice: }), Fanmodules: []EquipmentFanModuleRelationship{), ManagementContoller: , ManagementEntity: openapiclient.management.Entity.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", EntityId: "EntityId_example", Leadership: "Leadership_example", NetworkElement: , RegisteredDevice: }, Psus: []EquipmentPsuRelationship{), RegisteredDevice: , TopSystem: openapiclient.top.System.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Ipv4Address: "Ipv4Address_example", Ipv6Address: "Ipv6Address_example", Mode: "Mode_example", Name: "Name_example", TimeZone: "TimeZone_example", ComputeBlades: []ComputeBladeRelationship{), ComputeRackUnits: []ComputeRackUnitRelationship{), ManagementController: , NetworkElements: []NetworkElementRelationship{), RegisteredDevice: }, UcsmRunningFirmware: }, RegisteredDevice: , RunningFirmware: []FirmwareRunningFirmwareRelationship{), StorageSasExpander: , TopSystem: }, RegisteredDevice: }), NetworkElement: , RegisteredDevice: , RunningFirmware: []FirmwareRunningFirmwareRelationship{), StorageSasExpander: , TopSystem: }, NetworkElements: []NetworkElementRelationship{), RegisteredDevice: , StorageController: , StoragePhysicalDisk: })}, ManagementController: , NetworkElements: []NetworkElementRelationship{), RegisteredDevice: , StorageController: , StoragePhysicalDisk: })}), Bmc: , Board: , BootDeviceBootmode: openapiclient.boot.DeviceBootMode.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", ConfiguredBootMode: "ConfiguredBootMode_example", ComputeRackUnit: , RegisteredDevice: }, Fanmodules: []EquipmentFanModuleRelationship{), GenericInventoryHolders: []InventoryGenericInventoryHolderRelationship{openapiclient.inventory.GenericInventoryHolder.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Endpoint: "Endpoint_example", ComputeBlade: , ComputeRackUnit: , GenericInventory: []InventoryGenericInventoryRelationship{openapiclient.inventory.GenericInventory.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Key: "Key_example", Type: "Type_example", Value: "Value_example", InventoryGenericInventoryHolder: openapiclient.inventory.GenericInventoryHolder.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Endpoint: "Endpoint_example", ComputeBlade: , ComputeRackUnit: , GenericInventory: []InventoryGenericInventoryRelationship{openapiclient.inventory.GenericInventory.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Key: "Key_example", Type: "Type_example", Value: "Value_example", InventoryGenericInventoryHolder: , RegisteredDevice: }), RegisteredDevice: }, RegisteredDevice: }), RegisteredDevice: }), LocatorLed: , PciDevices: []PciDeviceRelationship{openapiclient.pci.Device.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", FirmwareVersion: "FirmwareVersion_example", Pid: "Pid_example", SlotId: "SlotId_example", ComputeBlade: , ComputeRackUnit: , RegisteredDevice: }), Psus: []EquipmentPsuRelationship{), RackEnclosureSlot: , RegisteredDevice: , SasExpanders: []StorageSasExpanderRelationship{), StorageEnclosures: []StorageEnclosureRelationship{), TopSystem: }, RegisteredDevice: }, Biosunits: []BiosUnitRelationship{), Bmc: , Board: , BootDeviceBootmode: openapiclient.boot.DeviceBootMode.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", ConfiguredBootMode: "ConfiguredBootMode_example", ComputeRackUnit: , RegisteredDevice: }, Fanmodules: []EquipmentFanModuleRelationship{), GenericInventoryHolders: []InventoryGenericInventoryHolderRelationship{), LocatorLed: , PciDevices: []PciDeviceRelationship{openapiclient.pci.Device.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", FirmwareVersion: "FirmwareVersion_example", Pid: "Pid_example", SlotId: "SlotId_example", ComputeBlade: , ComputeRackUnit: , RegisteredDevice: }), Psus: []EquipmentPsuRelationship{), RackEnclosureSlot: , RegisteredDevice: , SasExpanders: []StorageSasExpanderRelationship{), StorageEnclosures: []StorageEnclosureRelationship{), TopSystem: }, Controller: , ExtEthIfs: []AdapterExtEthInterfaceRelationship{openapiclient.adapter.ExtEthInterface.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", EpDn: "EpDn_example", ExtEthInterfaceId: "ExtEthInterfaceId_example", InterfaceType: "InterfaceType_example", MacAddress: "MacAddress_example", OperState: "OperState_example", PeerDn: "PeerDn_example", AdapterUnit: , RegisteredDevice: }), HostEthIfs: []AdapterHostEthInterfaceRelationship{openapiclient.adapter.HostEthInterface.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", EpDn: "EpDn_example", HostEthInterfaceId: int64(123), InterfaceType: "InterfaceType_example", MacAddress: "MacAddress_example", Name: "Name_example", OperState: "OperState_example", Operability: "Operability_example", OriginalMacAddress: "OriginalMacAddress_example", PciAddr: "PciAddr_example", PeerDn: "PeerDn_example", VirtualizationPreference: "VirtualizationPreference_example", VnicDn: "VnicDn_example", AdapterUnit: , RegisteredDevice: }), HostFcIfs: []AdapterHostFcInterfaceRelationship{openapiclient.adapter.HostFcInterface.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", EpDn: "EpDn_example", HostFcInterfaceId: int64(123), Name: "Name_example", OperState: "OperState_example", Operability: "Operability_example", OriginalWwnn: "OriginalWwnn_example", OriginalWwpn: "OriginalWwpn_example", PeerDn: "PeerDn_example", Wwnn: "Wwnn_example", Wwpn: "Wwpn_example", AdapterUnit: , RegisteredDevice: }), HostIscsiIfs: []AdapterHostIscsiInterfaceRelationship{openapiclient.adapter.HostIscsiInterface.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", EpDn: "EpDn_example", HostIscsiInterfaceId: int64(123), HostVisible: "HostVisible_example", MacAddress: "MacAddress_example", Name: "Name_example", OperState: "OperState_example", Operability: "Operability_example", PeerDn: "PeerDn_example", AdapterUnit: , RegisteredDevice: }), RegisteredDevice: }), BiosUnits: []BiosUnitRelationship{), Bmc: , Board: , EquipmentChassis: , EquipmentIoExpanders: []EquipmentIoExpanderRelationship{openapiclient.equipment.IoExpander.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", OperState: "OperState_example", Presence: "Presence_example", ComputeBlade: , RegisteredDevice: }), GenericInventoryHolders: []InventoryGenericInventoryHolderRelationship{), LocatorLed: , PciDevices: []PciDeviceRelationship{), RegisteredDevice: , StorageEnclosures: []StorageEnclosureRelationship{), TopSystem: }, ComputeRackUnit: , Controller: , ExtEthIfs: []AdapterExtEthInterfaceRelationship{openapiclient.adapter.ExtEthInterface.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", EpDn: "EpDn_example", ExtEthInterfaceId: "ExtEthInterfaceId_example", InterfaceType: "InterfaceType_example", MacAddress: "MacAddress_example", OperState: "OperState_example", PeerDn: "PeerDn_example", AdapterUnit: , RegisteredDevice: }), HostEthIfs: []AdapterHostEthInterfaceRelationship{openapiclient.adapter.HostEthInterface.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", EpDn: "EpDn_example", HostEthInterfaceId: int64(123), InterfaceType: "InterfaceType_example", MacAddress: "MacAddress_example", Name: "Name_example", OperState: "OperState_example", Operability: "Operability_example", OriginalMacAddress: "OriginalMacAddress_example", PciAddr: "PciAddr_example", PeerDn: "PeerDn_example", VirtualizationPreference: "VirtualizationPreference_example", VnicDn: "VnicDn_example", AdapterUnit: , RegisteredDevice: }), HostFcIfs: []AdapterHostFcInterfaceRelationship{openapiclient.adapter.HostFcInterface.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", EpDn: "EpDn_example", HostFcInterfaceId: int64(123), Name: "Name_example", OperState: "OperState_example", Operability: "Operability_example", OriginalWwnn: "OriginalWwnn_example", OriginalWwpn: "OriginalWwpn_example", PeerDn: "PeerDn_example", Wwnn: "Wwnn_example", Wwpn: "Wwpn_example", AdapterUnit: , RegisteredDevice: }), HostIscsiIfs: []AdapterHostIscsiInterfaceRelationship{openapiclient.adapter.HostIscsiInterface.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", EpDn: "EpDn_example", HostIscsiInterfaceId: int64(123), HostVisible: "HostVisible_example", MacAddress: "MacAddress_example", Name: "Name_example", OperState: "OperState_example", Operability: "Operability_example", PeerDn: "PeerDn_example", AdapterUnit: , RegisteredDevice: }), RegisteredDevice: }), BiosUnits: []BiosUnitRelationship{), Bmc: , Board: , EquipmentChassis: , EquipmentIoExpanders: []EquipmentIoExpanderRelationship{openapiclient.equipment.IoExpander.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", OperState: "OperState_example", Presence: "Presence_example", ComputeBlade: , RegisteredDevice: }), GenericInventoryHolders: []InventoryGenericInventoryHolderRelationship{), LocatorLed: , PciDevices: []PciDeviceRelationship{), RegisteredDevice: , StorageEnclosures: []StorageEnclosureRelationship{), TopSystem: }, ComputeRackUnit: , EquipmentTpms: []EquipmentTpmRelationship{openapiclient.equipment.Tpm.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ActivationStatus: "ActivationStatus_example", AdminState: "AdminState_example", Ownership: "Ownership_example", Presence: "Presence_example", TpmId: int64(123), Version: "Version_example", ComputeBoard: , RegisteredDevice: }), GraphicsCards: []GraphicsCardRelationship{), MemoryArrays: []MemoryArrayRelationship{), PciCoprocessorCards: []PciCoprocessorCardRelationship{openapiclient.pci.CoprocessorCard.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", CardId: int64(123), PciSlot: "PciSlot_example", ComputeBoard: , RegisteredDevice: }), PciSwitch: []PciSwitchRelationship{), PersistentMemoryConfiguration: , Processors: []ProcessorUnitRelationship{openapiclient.processor.Unit.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", Architecture: "Architecture_example", NumCores: int64(123), NumCoresEnabled: "NumCoresEnabled_example", NumThreads: "NumThreads_example", OperPowerState: "OperPowerState_example", OperState: "OperState_example", Operability: "Operability_example", Presence: "Presence_example", ProcessorId: int64(123), SocketDesignation: "SocketDesignation_example", Speed: 123, Stepping: "Stepping_example", Thermal: "Thermal_example", ComputeBoard: , RegisteredDevice: }), RegisteredDevice: , SecurityUnits: []SecurityUnitRelationship{openapiclient.security.Unit.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", OperState: "OperState_example", Operability: "Operability_example", PartNumber: "PartNumber_example", PciSlot: "PciSlot_example", Power: "Power_example", Presence: "Presence_example", Thermal: "Thermal_example", UnitId: int64(123), Vid: "Vid_example", Voltage: "Voltage_example", ComputeBoard: , RegisteredDevice: }), StorageControllers: []StorageControllerRelationship{), StorageFlexFlashControllers: []StorageFlexFlashControllerRelationship{), StorageFlexUtilControllers: []StorageFlexUtilControllerRelationship{)}, PersistentMemoryUnits: []MemoryPersistentMemoryUnitRelationship{), RegisteredDevice: , Units: []MemoryUnitRelationship{)} // MemoryArray | The 'memory.Array' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.PatchMemoryArray(context.Background(), moid, memoryArray).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.PatchMemoryArray``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMemoryArray`: MemoryArray
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.PatchMemoryArray`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMemoryArrayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryArray** | [**MemoryArray**](MemoryArray.md) | The &#39;memory.Array&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryArray**](memory.Array.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMemoryPersistentMemoryConfigResult

> MemoryPersistentMemoryConfigResult PatchMemoryPersistentMemoryConfigResult(ctx, moid).MemoryPersistentMemoryConfigResult(memoryPersistentMemoryConfigResult).IfMatch(ifMatch).Execute()

Update a 'memory.PersistentMemoryConfigResult' resource.

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
    memoryPersistentMemoryConfigResult := openapiclient.memory.PersistentMemoryConfigResult{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", ConfigErrorDesc: "ConfigErrorDesc_example", ConfigResult: "ConfigResult_example", ConfigSequenceNo: int64(123), ConfigState: "ConfigState_example", MemoryPersistentMemoryConfiguration: , PersistentMemoryNamespaceConfigResults: []MemoryPersistentMemoryNamespaceConfigResultRelationship{), RegisteredDevice: } // MemoryPersistentMemoryConfigResult | The 'memory.PersistentMemoryConfigResult' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.PatchMemoryPersistentMemoryConfigResult(context.Background(), moid, memoryPersistentMemoryConfigResult).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.PatchMemoryPersistentMemoryConfigResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMemoryPersistentMemoryConfigResult`: MemoryPersistentMemoryConfigResult
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.PatchMemoryPersistentMemoryConfigResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMemoryPersistentMemoryConfigResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryPersistentMemoryConfigResult** | [**MemoryPersistentMemoryConfigResult**](MemoryPersistentMemoryConfigResult.md) | The &#39;memory.PersistentMemoryConfigResult&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryPersistentMemoryConfigResult**](memory.PersistentMemoryConfigResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMemoryPersistentMemoryConfiguration

> MemoryPersistentMemoryConfiguration PatchMemoryPersistentMemoryConfiguration(ctx, moid).MemoryPersistentMemoryConfiguration(memoryPersistentMemoryConfiguration).IfMatch(ifMatch).Execute()

Update a 'memory.PersistentMemoryConfiguration' resource.

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
    memoryPersistentMemoryConfiguration := openapiclient.memory.PersistentMemoryConfiguration{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", MemoryCapacity: "MemoryCapacity_example", NumOfModules: "NumOfModules_example", NumOfRegions: "NumOfRegions_example", PersistentMemoryCapacity: "PersistentMemoryCapacity_example", ReservedCapacity: "ReservedCapacity_example", SecurityState: "SecurityState_example", TotalCapacity: "TotalCapacity_example", ComputeBoard: , PersistentMemoryConfigResult: , PersistentMemoryRegions: []MemoryPersistentMemoryRegionRelationship{), RegisteredDevice: } // MemoryPersistentMemoryConfiguration | The 'memory.PersistentMemoryConfiguration' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.PatchMemoryPersistentMemoryConfiguration(context.Background(), moid, memoryPersistentMemoryConfiguration).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.PatchMemoryPersistentMemoryConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMemoryPersistentMemoryConfiguration`: MemoryPersistentMemoryConfiguration
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.PatchMemoryPersistentMemoryConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMemoryPersistentMemoryConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryPersistentMemoryConfiguration** | [**MemoryPersistentMemoryConfiguration**](MemoryPersistentMemoryConfiguration.md) | The &#39;memory.PersistentMemoryConfiguration&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryPersistentMemoryConfiguration**](memory.PersistentMemoryConfiguration.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMemoryPersistentMemoryNamespace

> MemoryPersistentMemoryNamespace PatchMemoryPersistentMemoryNamespace(ctx, moid).MemoryPersistentMemoryNamespace(memoryPersistentMemoryNamespace).IfMatch(ifMatch).Execute()

Update a 'memory.PersistentMemoryNamespace' resource.

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
    memoryPersistentMemoryNamespace := openapiclient.memory.PersistentMemoryNamespace{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Capacity: "Capacity_example", HealthState: "HealthState_example", LabelVersion: "LabelVersion_example", Mode: "Mode_example", Name: "Name_example", Uuid: "Uuid_example", MemoryPersistentMemoryRegion: , RegisteredDevice: } // MemoryPersistentMemoryNamespace | The 'memory.PersistentMemoryNamespace' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.PatchMemoryPersistentMemoryNamespace(context.Background(), moid, memoryPersistentMemoryNamespace).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.PatchMemoryPersistentMemoryNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMemoryPersistentMemoryNamespace`: MemoryPersistentMemoryNamespace
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.PatchMemoryPersistentMemoryNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMemoryPersistentMemoryNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryPersistentMemoryNamespace** | [**MemoryPersistentMemoryNamespace**](MemoryPersistentMemoryNamespace.md) | The &#39;memory.PersistentMemoryNamespace&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryPersistentMemoryNamespace**](memory.PersistentMemoryNamespace.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMemoryPersistentMemoryNamespaceConfigResult

> MemoryPersistentMemoryNamespaceConfigResult PatchMemoryPersistentMemoryNamespaceConfigResult(ctx, moid).MemoryPersistentMemoryNamespaceConfigResult(memoryPersistentMemoryNamespaceConfigResult).IfMatch(ifMatch).Execute()

Update a 'memory.PersistentMemoryNamespaceConfigResult' resource.

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
    memoryPersistentMemoryNamespaceConfigResult := openapiclient.memory.PersistentMemoryNamespaceConfigResult{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", ConfigStatus: "ConfigStatus_example", Name: "Name_example", SocketId: "SocketId_example", SocketMemoryId: "SocketMemoryId_example", MemoryPersistentMemoryConfigResult: , RegisteredDevice: } // MemoryPersistentMemoryNamespaceConfigResult | The 'memory.PersistentMemoryNamespaceConfigResult' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.PatchMemoryPersistentMemoryNamespaceConfigResult(context.Background(), moid, memoryPersistentMemoryNamespaceConfigResult).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.PatchMemoryPersistentMemoryNamespaceConfigResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMemoryPersistentMemoryNamespaceConfigResult`: MemoryPersistentMemoryNamespaceConfigResult
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.PatchMemoryPersistentMemoryNamespaceConfigResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMemoryPersistentMemoryNamespaceConfigResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryPersistentMemoryNamespaceConfigResult** | [**MemoryPersistentMemoryNamespaceConfigResult**](MemoryPersistentMemoryNamespaceConfigResult.md) | The &#39;memory.PersistentMemoryNamespaceConfigResult&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryPersistentMemoryNamespaceConfigResult**](memory.PersistentMemoryNamespaceConfigResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMemoryPersistentMemoryPolicy

> MemoryPersistentMemoryPolicy PatchMemoryPersistentMemoryPolicy(ctx, moid).MemoryPersistentMemoryPolicy(memoryPersistentMemoryPolicy).IfMatch(ifMatch).Execute()

Update a 'memory.PersistentMemoryPolicy' resource.

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
    memoryPersistentMemoryPolicy := openapiclient.memory.PersistentMemoryPolicy{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Goals: []MemoryPersistentMemoryGoal{openapiclient.memory.PersistentMemoryGoal{ClassId: "ClassId_example", ObjectType: "ObjectType_example", MemoryModePercentage: 123, PersistentMemoryType: "PersistentMemoryType_example", SocketId: "SocketId_example"}), LocalSecurity: openapiclient.memory.PersistentMemoryLocalSecurity{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Enabled: false, IsSecurePassphraseSet: false, SecurePassphrase: "SecurePassphrase_example"}, LogicalNamespaces: []MemoryPersistentMemoryLogicalNamespace{openapiclient.memory.PersistentMemoryLogicalNamespace{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Capacity: 123, Mode: "Mode_example", Name: "Name_example", SocketId: 123, SocketMemoryId: "SocketMemoryId_example"}), ManagementMode: "ManagementMode_example", RetainNamespaces: false, Organization: , Profiles: []PolicyAbstractConfigProfileRelationship{)} // MemoryPersistentMemoryPolicy | The 'memory.PersistentMemoryPolicy' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.PatchMemoryPersistentMemoryPolicy(context.Background(), moid, memoryPersistentMemoryPolicy).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.PatchMemoryPersistentMemoryPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMemoryPersistentMemoryPolicy`: MemoryPersistentMemoryPolicy
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.PatchMemoryPersistentMemoryPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMemoryPersistentMemoryPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryPersistentMemoryPolicy** | [**MemoryPersistentMemoryPolicy**](MemoryPersistentMemoryPolicy.md) | The &#39;memory.PersistentMemoryPolicy&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryPersistentMemoryPolicy**](memory.PersistentMemoryPolicy.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMemoryPersistentMemoryRegion

> MemoryPersistentMemoryRegion PatchMemoryPersistentMemoryRegion(ctx, moid).MemoryPersistentMemoryRegion(memoryPersistentMemoryRegion).IfMatch(ifMatch).Execute()

Update a 'memory.PersistentMemoryRegion' resource.

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
    memoryPersistentMemoryRegion := openapiclient.memory.PersistentMemoryRegion{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", FreeCapacity: "FreeCapacity_example", HealthState: "HealthState_example", InterleavedSetId: "InterleavedSetId_example", LocaterIds: "LocaterIds_example", PersistentMemoryType: "PersistentMemoryType_example", RegionId: "RegionId_example", SocketId: "SocketId_example", SocketMemoryId: "SocketMemoryId_example", TotalCapacity: "TotalCapacity_example", MemoryPersistentMemoryConfiguration: , PersistentMemoryNamespaces: []MemoryPersistentMemoryNamespaceRelationship{), RegisteredDevice: } // MemoryPersistentMemoryRegion | The 'memory.PersistentMemoryRegion' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.PatchMemoryPersistentMemoryRegion(context.Background(), moid, memoryPersistentMemoryRegion).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.PatchMemoryPersistentMemoryRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMemoryPersistentMemoryRegion`: MemoryPersistentMemoryRegion
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.PatchMemoryPersistentMemoryRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMemoryPersistentMemoryRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryPersistentMemoryRegion** | [**MemoryPersistentMemoryRegion**](MemoryPersistentMemoryRegion.md) | The &#39;memory.PersistentMemoryRegion&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryPersistentMemoryRegion**](memory.PersistentMemoryRegion.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMemoryPersistentMemoryUnit

> MemoryPersistentMemoryUnit PatchMemoryPersistentMemoryUnit(ctx, moid).MemoryPersistentMemoryUnit(memoryPersistentMemoryUnit).IfMatch(ifMatch).Execute()

Update a 'memory.PersistentMemoryUnit' resource.

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
    memoryPersistentMemoryUnit := openapiclient.memory.PersistentMemoryUnit{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", ArrayId: int64(123), Bank: int64(123), Capacity: "Capacity_example", Clock: "Clock_example", FormFactor: "FormFactor_example", Latency: "Latency_example", Location: "Location_example", OperPowerState: "OperPowerState_example", OperState: "OperState_example", Operability: "Operability_example", Presence: "Presence_example", Set: int64(123), Speed: "Speed_example", Thermal: "Thermal_example", Type: "Type_example", Visibility: "Visibility_example", Width: "Width_example", AppDirectCapacity: "AppDirectCapacity_example", CountStatus: "CountStatus_example", FirmwareVersion: "FirmwareVersion_example", FrozenStatus: "FrozenStatus_example", HealthState: "HealthState_example", LockStatus: "LockStatus_example", MemoryCapacity: "MemoryCapacity_example", MemoryId: int64(123), PersistentMemoryCapacity: "PersistentMemoryCapacity_example", ReservedCapacity: "ReservedCapacity_example", SecurityStatus: "SecurityStatus_example", SocketId: "SocketId_example", SocketMemoryId: "SocketMemoryId_example", TotalCapacity: "TotalCapacity_example", Uid: "Uid_example", MemoryArray: , RegisteredDevice: } // MemoryPersistentMemoryUnit | The 'memory.PersistentMemoryUnit' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.PatchMemoryPersistentMemoryUnit(context.Background(), moid, memoryPersistentMemoryUnit).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.PatchMemoryPersistentMemoryUnit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMemoryPersistentMemoryUnit`: MemoryPersistentMemoryUnit
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.PatchMemoryPersistentMemoryUnit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMemoryPersistentMemoryUnitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryPersistentMemoryUnit** | [**MemoryPersistentMemoryUnit**](MemoryPersistentMemoryUnit.md) | The &#39;memory.PersistentMemoryUnit&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryPersistentMemoryUnit**](memory.PersistentMemoryUnit.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMemoryUnit

> MemoryUnit PatchMemoryUnit(ctx, moid).MemoryUnit(memoryUnit).IfMatch(ifMatch).Execute()

Update a 'memory.Unit' resource.

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
    memoryUnit := openapiclient.memory.Unit{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", ArrayId: int64(123), Bank: int64(123), Capacity: "Capacity_example", Clock: "Clock_example", FormFactor: "FormFactor_example", Latency: "Latency_example", Location: "Location_example", OperPowerState: "OperPowerState_example", OperState: "OperState_example", Operability: "Operability_example", Presence: "Presence_example", Set: int64(123), Speed: "Speed_example", Thermal: "Thermal_example", Type: "Type_example", Visibility: "Visibility_example", Width: "Width_example", MemoryId: int64(123), MemoryArray: , RegisteredDevice: } // MemoryUnit | The 'memory.Unit' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.PatchMemoryUnit(context.Background(), moid, memoryUnit).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.PatchMemoryUnit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMemoryUnit`: MemoryUnit
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.PatchMemoryUnit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMemoryUnitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryUnit** | [**MemoryUnit**](MemoryUnit.md) | The &#39;memory.Unit&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryUnit**](memory.Unit.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMemoryArray

> MemoryArray UpdateMemoryArray(ctx, moid).MemoryArray(memoryArray).IfMatch(ifMatch).Execute()

Update a 'memory.Array' resource.

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
    memoryArray := openapiclient.memory.Array{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", ArrayId: int64(123), CpuId: int64(123), CurrentCapacity: "CurrentCapacity_example", ErrorCorrection: "ErrorCorrection_example", MaxCapacity: "MaxCapacity_example", MaxDevices: "MaxDevices_example", OperPowerState: "OperPowerState_example", Presence: "Presence_example", ComputeBoard: , PersistentMemoryUnits: []MemoryPersistentMemoryUnitRelationship{), RegisteredDevice: , Units: []MemoryUnitRelationship{)} // MemoryArray | The 'memory.Array' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.UpdateMemoryArray(context.Background(), moid, memoryArray).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.UpdateMemoryArray``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMemoryArray`: MemoryArray
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.UpdateMemoryArray`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemoryArrayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryArray** | [**MemoryArray**](MemoryArray.md) | The &#39;memory.Array&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryArray**](memory.Array.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMemoryPersistentMemoryConfigResult

> MemoryPersistentMemoryConfigResult UpdateMemoryPersistentMemoryConfigResult(ctx, moid).MemoryPersistentMemoryConfigResult(memoryPersistentMemoryConfigResult).IfMatch(ifMatch).Execute()

Update a 'memory.PersistentMemoryConfigResult' resource.

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
    memoryPersistentMemoryConfigResult := openapiclient.memory.PersistentMemoryConfigResult{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", ConfigErrorDesc: "ConfigErrorDesc_example", ConfigResult: "ConfigResult_example", ConfigSequenceNo: int64(123), ConfigState: "ConfigState_example", MemoryPersistentMemoryConfiguration: , PersistentMemoryNamespaceConfigResults: []MemoryPersistentMemoryNamespaceConfigResultRelationship{), RegisteredDevice: } // MemoryPersistentMemoryConfigResult | The 'memory.PersistentMemoryConfigResult' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.UpdateMemoryPersistentMemoryConfigResult(context.Background(), moid, memoryPersistentMemoryConfigResult).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.UpdateMemoryPersistentMemoryConfigResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMemoryPersistentMemoryConfigResult`: MemoryPersistentMemoryConfigResult
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.UpdateMemoryPersistentMemoryConfigResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemoryPersistentMemoryConfigResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryPersistentMemoryConfigResult** | [**MemoryPersistentMemoryConfigResult**](MemoryPersistentMemoryConfigResult.md) | The &#39;memory.PersistentMemoryConfigResult&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryPersistentMemoryConfigResult**](memory.PersistentMemoryConfigResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMemoryPersistentMemoryConfiguration

> MemoryPersistentMemoryConfiguration UpdateMemoryPersistentMemoryConfiguration(ctx, moid).MemoryPersistentMemoryConfiguration(memoryPersistentMemoryConfiguration).IfMatch(ifMatch).Execute()

Update a 'memory.PersistentMemoryConfiguration' resource.

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
    memoryPersistentMemoryConfiguration := openapiclient.memory.PersistentMemoryConfiguration{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", MemoryCapacity: "MemoryCapacity_example", NumOfModules: "NumOfModules_example", NumOfRegions: "NumOfRegions_example", PersistentMemoryCapacity: "PersistentMemoryCapacity_example", ReservedCapacity: "ReservedCapacity_example", SecurityState: "SecurityState_example", TotalCapacity: "TotalCapacity_example", ComputeBoard: , PersistentMemoryConfigResult: , PersistentMemoryRegions: []MemoryPersistentMemoryRegionRelationship{), RegisteredDevice: } // MemoryPersistentMemoryConfiguration | The 'memory.PersistentMemoryConfiguration' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.UpdateMemoryPersistentMemoryConfiguration(context.Background(), moid, memoryPersistentMemoryConfiguration).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.UpdateMemoryPersistentMemoryConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMemoryPersistentMemoryConfiguration`: MemoryPersistentMemoryConfiguration
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.UpdateMemoryPersistentMemoryConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemoryPersistentMemoryConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryPersistentMemoryConfiguration** | [**MemoryPersistentMemoryConfiguration**](MemoryPersistentMemoryConfiguration.md) | The &#39;memory.PersistentMemoryConfiguration&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryPersistentMemoryConfiguration**](memory.PersistentMemoryConfiguration.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMemoryPersistentMemoryNamespace

> MemoryPersistentMemoryNamespace UpdateMemoryPersistentMemoryNamespace(ctx, moid).MemoryPersistentMemoryNamespace(memoryPersistentMemoryNamespace).IfMatch(ifMatch).Execute()

Update a 'memory.PersistentMemoryNamespace' resource.

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
    memoryPersistentMemoryNamespace := openapiclient.memory.PersistentMemoryNamespace{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Capacity: "Capacity_example", HealthState: "HealthState_example", LabelVersion: "LabelVersion_example", Mode: "Mode_example", Name: "Name_example", Uuid: "Uuid_example", MemoryPersistentMemoryRegion: , RegisteredDevice: } // MemoryPersistentMemoryNamespace | The 'memory.PersistentMemoryNamespace' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.UpdateMemoryPersistentMemoryNamespace(context.Background(), moid, memoryPersistentMemoryNamespace).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.UpdateMemoryPersistentMemoryNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMemoryPersistentMemoryNamespace`: MemoryPersistentMemoryNamespace
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.UpdateMemoryPersistentMemoryNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemoryPersistentMemoryNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryPersistentMemoryNamespace** | [**MemoryPersistentMemoryNamespace**](MemoryPersistentMemoryNamespace.md) | The &#39;memory.PersistentMemoryNamespace&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryPersistentMemoryNamespace**](memory.PersistentMemoryNamespace.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMemoryPersistentMemoryNamespaceConfigResult

> MemoryPersistentMemoryNamespaceConfigResult UpdateMemoryPersistentMemoryNamespaceConfigResult(ctx, moid).MemoryPersistentMemoryNamespaceConfigResult(memoryPersistentMemoryNamespaceConfigResult).IfMatch(ifMatch).Execute()

Update a 'memory.PersistentMemoryNamespaceConfigResult' resource.

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
    memoryPersistentMemoryNamespaceConfigResult := openapiclient.memory.PersistentMemoryNamespaceConfigResult{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", ConfigStatus: "ConfigStatus_example", Name: "Name_example", SocketId: "SocketId_example", SocketMemoryId: "SocketMemoryId_example", MemoryPersistentMemoryConfigResult: , RegisteredDevice: } // MemoryPersistentMemoryNamespaceConfigResult | The 'memory.PersistentMemoryNamespaceConfigResult' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.UpdateMemoryPersistentMemoryNamespaceConfigResult(context.Background(), moid, memoryPersistentMemoryNamespaceConfigResult).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.UpdateMemoryPersistentMemoryNamespaceConfigResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMemoryPersistentMemoryNamespaceConfigResult`: MemoryPersistentMemoryNamespaceConfigResult
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.UpdateMemoryPersistentMemoryNamespaceConfigResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemoryPersistentMemoryNamespaceConfigResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryPersistentMemoryNamespaceConfigResult** | [**MemoryPersistentMemoryNamespaceConfigResult**](MemoryPersistentMemoryNamespaceConfigResult.md) | The &#39;memory.PersistentMemoryNamespaceConfigResult&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryPersistentMemoryNamespaceConfigResult**](memory.PersistentMemoryNamespaceConfigResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMemoryPersistentMemoryPolicy

> MemoryPersistentMemoryPolicy UpdateMemoryPersistentMemoryPolicy(ctx, moid).MemoryPersistentMemoryPolicy(memoryPersistentMemoryPolicy).IfMatch(ifMatch).Execute()

Update a 'memory.PersistentMemoryPolicy' resource.

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
    memoryPersistentMemoryPolicy :=  // MemoryPersistentMemoryPolicy | The 'memory.PersistentMemoryPolicy' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.UpdateMemoryPersistentMemoryPolicy(context.Background(), moid, memoryPersistentMemoryPolicy).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.UpdateMemoryPersistentMemoryPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMemoryPersistentMemoryPolicy`: MemoryPersistentMemoryPolicy
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.UpdateMemoryPersistentMemoryPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemoryPersistentMemoryPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryPersistentMemoryPolicy** | [**MemoryPersistentMemoryPolicy**](MemoryPersistentMemoryPolicy.md) | The &#39;memory.PersistentMemoryPolicy&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryPersistentMemoryPolicy**](memory.PersistentMemoryPolicy.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMemoryPersistentMemoryRegion

> MemoryPersistentMemoryRegion UpdateMemoryPersistentMemoryRegion(ctx, moid).MemoryPersistentMemoryRegion(memoryPersistentMemoryRegion).IfMatch(ifMatch).Execute()

Update a 'memory.PersistentMemoryRegion' resource.

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
    memoryPersistentMemoryRegion := openapiclient.memory.PersistentMemoryRegion{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", FreeCapacity: "FreeCapacity_example", HealthState: "HealthState_example", InterleavedSetId: "InterleavedSetId_example", LocaterIds: "LocaterIds_example", PersistentMemoryType: "PersistentMemoryType_example", RegionId: "RegionId_example", SocketId: "SocketId_example", SocketMemoryId: "SocketMemoryId_example", TotalCapacity: "TotalCapacity_example", MemoryPersistentMemoryConfiguration: , PersistentMemoryNamespaces: []MemoryPersistentMemoryNamespaceRelationship{), RegisteredDevice: } // MemoryPersistentMemoryRegion | The 'memory.PersistentMemoryRegion' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.UpdateMemoryPersistentMemoryRegion(context.Background(), moid, memoryPersistentMemoryRegion).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.UpdateMemoryPersistentMemoryRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMemoryPersistentMemoryRegion`: MemoryPersistentMemoryRegion
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.UpdateMemoryPersistentMemoryRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemoryPersistentMemoryRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryPersistentMemoryRegion** | [**MemoryPersistentMemoryRegion**](MemoryPersistentMemoryRegion.md) | The &#39;memory.PersistentMemoryRegion&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryPersistentMemoryRegion**](memory.PersistentMemoryRegion.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMemoryPersistentMemoryUnit

> MemoryPersistentMemoryUnit UpdateMemoryPersistentMemoryUnit(ctx, moid).MemoryPersistentMemoryUnit(memoryPersistentMemoryUnit).IfMatch(ifMatch).Execute()

Update a 'memory.PersistentMemoryUnit' resource.

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
    memoryPersistentMemoryUnit := openapiclient.memory.PersistentMemoryUnit{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", ArrayId: int64(123), Bank: int64(123), Capacity: "Capacity_example", Clock: "Clock_example", FormFactor: "FormFactor_example", Latency: "Latency_example", Location: "Location_example", OperPowerState: "OperPowerState_example", OperState: "OperState_example", Operability: "Operability_example", Presence: "Presence_example", Set: int64(123), Speed: "Speed_example", Thermal: "Thermal_example", Type: "Type_example", Visibility: "Visibility_example", Width: "Width_example", AppDirectCapacity: "AppDirectCapacity_example", CountStatus: "CountStatus_example", FirmwareVersion: "FirmwareVersion_example", FrozenStatus: "FrozenStatus_example", HealthState: "HealthState_example", LockStatus: "LockStatus_example", MemoryCapacity: "MemoryCapacity_example", MemoryId: int64(123), PersistentMemoryCapacity: "PersistentMemoryCapacity_example", ReservedCapacity: "ReservedCapacity_example", SecurityStatus: "SecurityStatus_example", SocketId: "SocketId_example", SocketMemoryId: "SocketMemoryId_example", TotalCapacity: "TotalCapacity_example", Uid: "Uid_example", MemoryArray: , RegisteredDevice: } // MemoryPersistentMemoryUnit | The 'memory.PersistentMemoryUnit' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.UpdateMemoryPersistentMemoryUnit(context.Background(), moid, memoryPersistentMemoryUnit).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.UpdateMemoryPersistentMemoryUnit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMemoryPersistentMemoryUnit`: MemoryPersistentMemoryUnit
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.UpdateMemoryPersistentMemoryUnit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemoryPersistentMemoryUnitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryPersistentMemoryUnit** | [**MemoryPersistentMemoryUnit**](MemoryPersistentMemoryUnit.md) | The &#39;memory.PersistentMemoryUnit&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryPersistentMemoryUnit**](memory.PersistentMemoryUnit.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMemoryUnit

> MemoryUnit UpdateMemoryUnit(ctx, moid).MemoryUnit(memoryUnit).IfMatch(ifMatch).Execute()

Update a 'memory.Unit' resource.

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
    memoryUnit := openapiclient.memory.Unit{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DeviceMoId: "DeviceMoId_example", Dn: "Dn_example", Rn: "Rn_example", Model: "Model_example", Revision: "Revision_example", Serial: "Serial_example", Vendor: "Vendor_example", AdminState: "AdminState_example", ArrayId: int64(123), Bank: int64(123), Capacity: "Capacity_example", Clock: "Clock_example", FormFactor: "FormFactor_example", Latency: "Latency_example", Location: "Location_example", OperPowerState: "OperPowerState_example", OperState: "OperState_example", Operability: "Operability_example", Presence: "Presence_example", Set: int64(123), Speed: "Speed_example", Thermal: "Thermal_example", Type: "Type_example", Visibility: "Visibility_example", Width: "Width_example", MemoryId: int64(123), MemoryArray: , RegisteredDevice: } // MemoryUnit | The 'memory.Unit' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemoryApi.UpdateMemoryUnit(context.Background(), moid, memoryUnit).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemoryApi.UpdateMemoryUnit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMemoryUnit`: MemoryUnit
    fmt.Fprintf(os.Stdout, "Response from `MemoryApi.UpdateMemoryUnit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemoryUnitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryUnit** | [**MemoryUnit**](MemoryUnit.md) | The &#39;memory.Unit&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**MemoryUnit**](memory.Unit.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

