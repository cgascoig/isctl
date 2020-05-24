# \SoftwareApi

All URIs are relative to *https://intersight.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSoftwareHclMeta**](SoftwareApi.md#CreateSoftwareHclMeta) | **Post** /software/HclMeta | Create a &#39;software.HclMeta&#39; resource.
[**CreateSoftwareHyperflexDistributable**](SoftwareApi.md#CreateSoftwareHyperflexDistributable) | **Post** /software/HyperflexDistributables | Create a &#39;software.HyperflexDistributable&#39; resource.
[**CreateSoftwareSolutionDistributable**](SoftwareApi.md#CreateSoftwareSolutionDistributable) | **Post** /software/SolutionDistributables | Create a &#39;software.SolutionDistributable&#39; resource.
[**DeleteSoftwareHclMeta**](SoftwareApi.md#DeleteSoftwareHclMeta) | **Delete** /software/HclMeta/{Moid} | Delete a &#39;software.HclMeta&#39; resource.
[**DeleteSoftwareHyperflexDistributable**](SoftwareApi.md#DeleteSoftwareHyperflexDistributable) | **Delete** /software/HyperflexDistributables/{Moid} | Delete a &#39;software.HyperflexDistributable&#39; resource.
[**DeleteSoftwareSolutionDistributable**](SoftwareApi.md#DeleteSoftwareSolutionDistributable) | **Delete** /software/SolutionDistributables/{Moid} | Delete a &#39;software.SolutionDistributable&#39; resource.
[**GetSoftwareHclMetaByMoid**](SoftwareApi.md#GetSoftwareHclMetaByMoid) | **Get** /software/HclMeta/{Moid} | Read a &#39;software.HclMeta&#39; resource.
[**GetSoftwareHclMetaList**](SoftwareApi.md#GetSoftwareHclMetaList) | **Get** /software/HclMeta | Read a &#39;software.HclMeta&#39; resource.
[**GetSoftwareHyperflexDistributableByMoid**](SoftwareApi.md#GetSoftwareHyperflexDistributableByMoid) | **Get** /software/HyperflexDistributables/{Moid} | Read a &#39;software.HyperflexDistributable&#39; resource.
[**GetSoftwareHyperflexDistributableList**](SoftwareApi.md#GetSoftwareHyperflexDistributableList) | **Get** /software/HyperflexDistributables | Read a &#39;software.HyperflexDistributable&#39; resource.
[**GetSoftwareSolutionDistributableByMoid**](SoftwareApi.md#GetSoftwareSolutionDistributableByMoid) | **Get** /software/SolutionDistributables/{Moid} | Read a &#39;software.SolutionDistributable&#39; resource.
[**GetSoftwareSolutionDistributableList**](SoftwareApi.md#GetSoftwareSolutionDistributableList) | **Get** /software/SolutionDistributables | Read a &#39;software.SolutionDistributable&#39; resource.
[**PatchSoftwareHclMeta**](SoftwareApi.md#PatchSoftwareHclMeta) | **Patch** /software/HclMeta/{Moid} | Update a &#39;software.HclMeta&#39; resource.
[**PatchSoftwareHyperflexDistributable**](SoftwareApi.md#PatchSoftwareHyperflexDistributable) | **Patch** /software/HyperflexDistributables/{Moid} | Update a &#39;software.HyperflexDistributable&#39; resource.
[**PatchSoftwareSolutionDistributable**](SoftwareApi.md#PatchSoftwareSolutionDistributable) | **Patch** /software/SolutionDistributables/{Moid} | Update a &#39;software.SolutionDistributable&#39; resource.
[**UpdateSoftwareHclMeta**](SoftwareApi.md#UpdateSoftwareHclMeta) | **Post** /software/HclMeta/{Moid} | Update a &#39;software.HclMeta&#39; resource.
[**UpdateSoftwareHyperflexDistributable**](SoftwareApi.md#UpdateSoftwareHyperflexDistributable) | **Post** /software/HyperflexDistributables/{Moid} | Update a &#39;software.HyperflexDistributable&#39; resource.
[**UpdateSoftwareSolutionDistributable**](SoftwareApi.md#UpdateSoftwareSolutionDistributable) | **Post** /software/SolutionDistributables/{Moid} | Update a &#39;software.SolutionDistributable&#39; resource.



## CreateSoftwareHclMeta

> SoftwareHclMeta CreateSoftwareHclMeta(ctx).SoftwareHclMeta(softwareHclMeta).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'software.HclMeta' resource.

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
    softwareHclMeta := openapiclient.software.HclMeta{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{openapiclient.mo.Tag{Key: "Key_example", Value: "Value_example"}), VersionContext: openapiclient.mo.VersionContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", InterestedMos: []MoMoRef{openapiclient.mo.MoRef{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example"}), RefMo: openapiclient.mo.MoRef{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example"}, Timestamp: "TODO", Version: "Version_example", VersionType: "VersionType_example"}, Ancestors: []MoBaseMoRelationship{openapiclient.mo.BaseMo.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{openapiclient.mo.Tag{Key: "Key_example", Value: "Value_example"}), VersionContext: openapiclient.mo.VersionContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", InterestedMos: []MoMoRef{), RefMo: , Timestamp: "TODO", Version: "Version_example", VersionType: "VersionType_example"}, Ancestors: []MoBaseMoRelationship{openapiclient.mo.BaseMo.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }}), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }}), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", DownloadCount: int64(123), ImportAction: "ImportAction_example", ImportState: "ImportState_example", ImportedTime: "TODO", LastAccessTime: "TODO", Md5sum: "Md5sum_example", Name: "Name_example", ReleaseDate: "TODO", Sha512sum: "Sha512sum_example", Size: int64(123), SoftwareAdvisoryUrl: "SoftwareAdvisoryUrl_example", Source: openapiclient.softwarerepository.FileServer{ClassId: "ClassId_example", ObjectType: "ObjectType_example"}, Version: "Version_example", BundleType: "BundleType_example", Guid: "Guid_example", Mdfid: "Mdfid_example", Model: "Model_example", PlatformType: "PlatformType_example", RecommendedBuild: "RecommendedBuild_example", ReleaseNotesUrl: "ReleaseNotesUrl_example", SoftwareTypeId: "SoftwareTypeId_example", SupportedModels: []string{"SupportedModels_example"), Vendor: "Vendor_example", ContentType: "ContentType_example", Catalog: openapiclient.softwarerepository.Catalog.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Organization: openapiclient.organization.Organization.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: openapiclient.iam.Account.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Status: "Status_example", AppRegistrations: []IamAppRegistrationRelationship{openapiclient.iam.AppRegistration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientId: "ClientId_example", ClientName: "ClientName_example", ClientSecret: "ClientSecret_example", ClientType: "ClientType_example", Description: "Description_example", GrantTypes: []string{"GrantTypes_example"), RedirectUris: []string{"RedirectUris_example"), RenewClientSecret: false, ResponseTypes: []string{"ResponseTypes_example"), RevocationTimestamp: "TODO", Revoke: false, Account: openapiclient.iam.Account.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Status: "Status_example", AppRegistrations: []IamAppRegistrationRelationship{openapiclient.iam.AppRegistration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientId: "ClientId_example", ClientName: "ClientName_example", ClientSecret: "ClientSecret_example", ClientType: "ClientType_example", Description: "Description_example", GrantTypes: []string{"GrantTypes_example"), RedirectUris: []string{"RedirectUris_example"), RenewClientSecret: false, ResponseTypes: []string{"ResponseTypes_example"), RevocationTimestamp: "TODO", Revoke: false, Account: , OauthTokens: []IamOAuthTokenRelationship{openapiclient.iam.OAuthToken.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccessExpirationTime: "TODO", ClientId: "ClientId_example", ClientIpAddress: "ClientIpAddress_example", ClientName: "ClientName_example", ExpirationTime: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", TokenId: "TokenId_example", UserMeta: openapiclient.iam.ClientMeta{ClassId: "ClassId_example", ObjectType: "ObjectType_example", DeviceModel: "DeviceModel_example", UserAgent: "UserAgent_example"}, AppRegistration: , Permission: openapiclient.iam.Permission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , EndPointRoles: []IamEndPointRoleRelationship{openapiclient.iam.EndPointRole.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", RoleType: "RoleType_example", Type: "Type_example", Account: , EndPointPrivileges: []IamEndPointPrivilegeRelationship{openapiclient.iam.EndPointPrivilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", System: openapiclient.iam.System.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointPrivileges: []IamEndPointPrivilegeRelationship{openapiclient.iam.EndPointPrivilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", System: openapiclient.iam.System.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointPrivileges: []IamEndPointPrivilegeRelationship{), EndPointRoles: []IamEndPointRoleRelationship{openapiclient.iam.EndPointRole.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", RoleType: "RoleType_example", Type: "Type_example", Account: , EndPointPrivileges: []IamEndPointPrivilegeRelationship{), System: }), Idp: openapiclient.iam.Idp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", Metadata: "Metadata_example", Name: "Name_example", Type: "Type_example", Account: , LdapPolicy: openapiclient.iam.LdapPolicy.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", BaseProperties: openapiclient.iam.LdapBaseProperties{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Attribute: "Attribute_example", BaseDn: "BaseDn_example", BindDn: "BindDn_example", BindMethod: "BindMethod_example", Domain: "Domain_example", EnableEncryption: false, EnableGroupAuthorization: false, Filter: "Filter_example", GroupAttribute: "GroupAttribute_example", IsPasswordSet: false, NestedGroupSearchDepth: int64(123), Password: "Password_example", Timeout: int64(123)}, DnsParameters: openapiclient.iam.LdapDnsParameters{ClassId: "ClassId_example", ObjectType: "ObjectType_example", SearchDomain: "SearchDomain_example", SearchForest: "SearchForest_example", Source: "Source_example"}, EnableDns: false, Enabled: false, UserSearchPrecedence: "UserSearchPrecedence_example", Var0Idp: openapiclient.iam.Idp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", Metadata: "Metadata_example", Name: "Name_example", Type: "Type_example", Account: , LdapPolicy: openapiclient.iam.LdapPolicy.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", BaseProperties: openapiclient.iam.LdapBaseProperties{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Attribute: "Attribute_example", BaseDn: "BaseDn_example", BindDn: "BindDn_example", BindMethod: "BindMethod_example", Domain: "Domain_example", EnableEncryption: false, EnableGroupAuthorization: false, Filter: "Filter_example", GroupAttribute: "GroupAttribute_example", IsPasswordSet: false, NestedGroupSearchDepth: int64(123), Password: "Password_example", Timeout: int64(123)}, DnsParameters: openapiclient.iam.LdapDnsParameters{ClassId: "ClassId_example", ObjectType: "ObjectType_example", SearchDomain: "SearchDomain_example", SearchForest: "SearchForest_example", Source: "Source_example"}, EnableDns: false, Enabled: false, UserSearchPrecedence: "UserSearchPrecedence_example", Var0Idp: , ApplianceAccount: , Groups: []IamLdapGroupRelationship{openapiclient.iam.LdapGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Domain: "Domain_example", Name: "Name_example", EndPointRole: []IamEndPointRoleRelationship{), LdapPolicy: }), Organization: openapiclient.organization.Organization.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , ResourceGroups: []ResourceGroupRelationship{openapiclient.resource.Group.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", PerTypeCombinedSelector: []ResourcePerTypeCombinedSelector{openapiclient.resource.PerTypeCombinedSelector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", CombinedSelector: "CombinedSelector_example", EmptyFilter: false, SelectorObjectType: "SelectorObjectType_example"}), Qualifier: "Qualifier_example", Selectors: []ResourceSelector{openapiclient.resource.Selector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Selector: "Selector_example"}), Account: , Organizations: []OrganizationOrganizationRelationship{)})}, Profiles: []PolicyAbstractConfigProfileRelationship{openapiclient.policy.AbstractConfigProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: openapiclient.policy.AbstractProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: openapiclient.policy.AbstractProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: }}, Action: "Action_example", ConfigContext: openapiclient.policy.ConfigContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", ConfigState: "ConfigState_example", ControlAction: "ControlAction_example", ErrorState: "ErrorState_example", OperState: "OperState_example"}}), Providers: []IamLdapProviderRelationship{openapiclient.iam.LdapProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Port: int64(123), Server: "Server_example", LdapPolicy: })}, System: , UserPreferences: []IamUserPreferenceRelationship{openapiclient.iam.UserPreference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Preference: 123, Idp: , IdpReference: openapiclient.iam.IdpReference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", MultiFactorAuthentication: false, Name: "Name_example", Account: , Idp: , UserPreferences: []IamUserPreferenceRelationship{openapiclient.iam.UserPreference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Preference: 123, Idp: , IdpReference: openapiclient.iam.IdpReference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", MultiFactorAuthentication: false, Name: "Name_example", Account: , Idp: , UserPreferences: []IamUserPreferenceRelationship{), Usergroups: []IamUserGroupRelationship{openapiclient.iam.UserGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Idp: , Idpreference: , Permissions: []IamPermissionRelationship{openapiclient.iam.Permission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , EndPointRoles: []IamEndPointRoleRelationship{), ResourceRoles: []IamResourceRolesRelationship{openapiclient.iam.ResourceRoles.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointRoles: []IamEndPointRoleRelationship{), Permission: , Resource: , Roles: []IamRoleRelationship{openapiclient.iam.Role.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , PrivilegeSets: []IamPrivilegeSetRelationship{openapiclient.iam.PrivilegeSet.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , AssociatedPrivilegeSets: []IamPrivilegeSetRelationship{openapiclient.iam.PrivilegeSet.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , AssociatedPrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{openapiclient.iam.Privilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HostnamePrefix: "HostnamePrefix_example", Method: "Method_example", Name: "Name_example", RestPath: "RestPath_example", UrlPrefix: "UrlPrefix_example", Account: , System: }), System: }), Privileges: []IamPrivilegeRelationship{openapiclient.iam.Privilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HostnamePrefix: "HostnamePrefix_example", Method: "Method_example", Name: "Name_example", RestPath: "RestPath_example", UrlPrefix: "UrlPrefix_example", Account: , System: }), System: }), System: })}), Roles: []IamRoleRelationship{openapiclient.iam.Role.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , PrivilegeSets: []IamPrivilegeSetRelationship{), System: }), SessionLimits: openapiclient.iam.SessionLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, IdleTimeOut: int64(123), MaximumLimit: int64(123), PerUserLimit: int64(123), SessionTimeOut: int64(123), Account: , Permission: }, UserGroups: []IamUserGroupRelationship{openapiclient.iam.UserGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Idp: , Idpreference: , Permissions: []IamPermissionRelationship{), Qualifier: openapiclient.iam.Qualifier.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Value: []string{"Value_example"), Usergroup: }, Users: []IamUserRelationship{openapiclient.iam.User.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientIpAddress: "ClientIpAddress_example", Email: "Email_example", FirstName: "FirstName_example", LastLoginTime: "TODO", LastName: "LastName_example", Name: "Name_example", UserIdOrEmail: "UserIdOrEmail_example", UserType: "UserType_example", ApiKeys: []IamApiKeyRelationship{openapiclient.iam.ApiKey.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HashAlgorithm: "HashAlgorithm_example", KeySpec: openapiclient.pkix.KeyGenerationSpec{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Name: "Name_example"}, PrivateKey: "PrivateKey_example", Purpose: "Purpose_example", SigningAlgorithm: "SigningAlgorithm_example", Permission: , User: openapiclient.iam.User.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientIpAddress: "ClientIpAddress_example", Email: "Email_example", FirstName: "FirstName_example", LastLoginTime: "TODO", LastName: "LastName_example", Name: "Name_example", UserIdOrEmail: "UserIdOrEmail_example", UserType: "UserType_example", ApiKeys: []IamApiKeyRelationship{openapiclient.iam.ApiKey.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HashAlgorithm: "HashAlgorithm_example", KeySpec: openapiclient.pkix.KeyGenerationSpec{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Name: "Name_example"}, PrivateKey: "PrivateKey_example", Purpose: "Purpose_example", SigningAlgorithm: "SigningAlgorithm_example", Permission: , User: }), AppRegistrations: []IamAppRegistrationRelationship{), Idp: , Idpreference: , LocalUserPassword: openapiclient.iam.LocalUserPassword.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, CurrentPassword: "CurrentPassword_example", NewPassword: "NewPassword_example", Password: 123, User: }, OauthTokens: []IamOAuthTokenRelationship{openapiclient.iam.OAuthToken.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccessExpirationTime: "TODO", ClientId: "ClientId_example", ClientIpAddress: "ClientIpAddress_example", ClientName: "ClientName_example", ExpirationTime: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", TokenId: "TokenId_example", UserMeta: openapiclient.iam.ClientMeta{ClassId: "ClassId_example", ObjectType: "ObjectType_example", DeviceModel: "DeviceModel_example", UserAgent: "UserAgent_example"}, AppRegistration: , Permission: , User: }), Permissions: []IamPermissionRelationship{), Sessions: []IamSessionRelationship{openapiclient.iam.Session.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccountPermissions: []IamAccountPermissions{openapiclient.iam.AccountPermissions{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountIdentifier: "AccountIdentifier_example", AccountName: "AccountName_example", AccountStatus: "AccountStatus_example", Permissions: []IamPermissionReference{openapiclient.iam.PermissionReference{ClassId: "ClassId_example", ObjectType: "ObjectType_example", PermissionIdentifier: "PermissionIdentifier_example", PermissionName: "PermissionName_example"})}), ClientIpAddress: "ClientIpAddress_example", Expiration: "TODO", IdleTimeExpiration: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", Permission: , User: })}}), AppRegistrations: []IamAppRegistrationRelationship{), Idp: , Idpreference: , LocalUserPassword: openapiclient.iam.LocalUserPassword.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, CurrentPassword: "CurrentPassword_example", NewPassword: "NewPassword_example", Password: 123, User: }, OauthTokens: []IamOAuthTokenRelationship{), Permissions: []IamPermissionRelationship{), Sessions: []IamSessionRelationship{openapiclient.iam.Session.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccountPermissions: []IamAccountPermissions{openapiclient.iam.AccountPermissions{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountIdentifier: "AccountIdentifier_example", AccountName: "AccountName_example", AccountStatus: "AccountStatus_example", Permissions: []IamPermissionReference{openapiclient.iam.PermissionReference{ClassId: "ClassId_example", ObjectType: "ObjectType_example", PermissionIdentifier: "PermissionIdentifier_example", PermissionName: "PermissionName_example"})}), ClientIpAddress: "ClientIpAddress_example", Expiration: "TODO", IdleTimeExpiration: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", Permission: , User: })})}), Users: []IamUserRelationship{)}), Qualifier: openapiclient.iam.Qualifier.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Value: []string{"Value_example"), Usergroup: }, Users: []IamUserRelationship{)}), Users: []IamUserRelationship{)}}), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}}), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, ApplianceAccount: , Groups: []IamLdapGroupRelationship{openapiclient.iam.LdapGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Domain: "Domain_example", Name: "Name_example", EndPointRole: []IamEndPointRoleRelationship{), LdapPolicy: }), Organization: , Profiles: []PolicyAbstractConfigProfileRelationship{openapiclient.policy.AbstractConfigProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: , Action: "Action_example", ConfigContext: openapiclient.policy.ConfigContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", ConfigState: "ConfigState_example", ControlAction: "ControlAction_example", ErrorState: "ErrorState_example", OperState: "OperState_example"}}), Providers: []IamLdapProviderRelationship{openapiclient.iam.LdapProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Port: int64(123), Server: "Server_example", LdapPolicy: })}, System: , UserPreferences: []IamUserPreferenceRelationship{), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), Roles: []IamRoleRelationship{), ServiceProvider: openapiclient.iam.ServiceProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EntityId: "EntityId_example", Metadata: "Metadata_example", Name: "Name_example", System: }}}), EndPointRoles: []IamEndPointRoleRelationship{), Idp: , PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), Roles: []IamRoleRelationship{), ServiceProvider: openapiclient.iam.ServiceProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EntityId: "EntityId_example", Metadata: "Metadata_example", Name: "Name_example", System: }}}), System: }), ResourceRoles: []IamResourceRolesRelationship{openapiclient.iam.ResourceRoles.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointRoles: []IamEndPointRoleRelationship{), Permission: , Resource: , Roles: []IamRoleRelationship{)}), Roles: []IamRoleRelationship{), SessionLimits: openapiclient.iam.SessionLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, IdleTimeOut: int64(123), MaximumLimit: int64(123), PerUserLimit: int64(123), SessionTimeOut: int64(123), Account: , Permission: }, UserGroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, User: }), Permission: , Roles: []IamRoleRelationship{), User: }), DomainGroups: []IamDomainGroupRelationship{openapiclient.iam.DomainGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Partition1: int64(123), Partition2: int64(123), Partition3: int64(123), PartitionKey: "PartitionKey_example", Usage: int64(123), Account: }), EndPointRoles: []IamEndPointRoleRelationship{), Idpreferences: []IamIdpReferenceRelationship{), Idps: []IamIdpRelationship{), Permissions: []IamPermissionRelationship{), PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), ResourceLimits: openapiclient.iam.ResourceLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PerAccountUserLimit: int64(123), Account: }, Roles: []IamRoleRelationship{), SecurityHolder: openapiclient.iam.SecurityHolder.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Account: , ResourcePermissions: []IamResourcePermissionRelationship{openapiclient.iam.ResourcePermission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PermissionRoles: []IamPermissionToRoles{openapiclient.iam.PermissionToRoles{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Permission: openapiclient.cmrf.CmRf{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example"}, Roles: []CmrfCmRf{openapiclient.cmrf.CmRf{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example"})}), TargetApp: "TargetApp_example", Holder: openapiclient.iam.SecurityHolder.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Account: , ResourcePermissions: []IamResourcePermissionRelationship{openapiclient.iam.ResourcePermission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PermissionRoles: []IamPermissionToRoles{openapiclient.iam.PermissionToRoles{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Permission: , Roles: []CmrfCmRf{)}), TargetApp: "TargetApp_example", Holder: , Resource: })}, Resource: })}, SessionLimits: }, OauthTokens: []IamOAuthTokenRelationship{), Permission: , Roles: []IamRoleRelationship{), User: }), DomainGroups: []IamDomainGroupRelationship{openapiclient.iam.DomainGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Partition1: int64(123), Partition2: int64(123), Partition3: int64(123), PartitionKey: "PartitionKey_example", Usage: int64(123), Account: }), EndPointRoles: []IamEndPointRoleRelationship{), Idpreferences: []IamIdpReferenceRelationship{), Idps: []IamIdpRelationship{), Permissions: []IamPermissionRelationship{), PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), ResourceLimits: openapiclient.iam.ResourceLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PerAccountUserLimit: int64(123), Account: }, Roles: []IamRoleRelationship{), SecurityHolder: , SessionLimits: }, ResourceGroups: []ResourceGroupRelationship{openapiclient.resource.Group.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", PerTypeCombinedSelector: []ResourcePerTypeCombinedSelector{openapiclient.resource.PerTypeCombinedSelector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", CombinedSelector: "CombinedSelector_example", EmptyFilter: false, SelectorObjectType: "SelectorObjectType_example"}), Qualifier: "Qualifier_example", Selectors: []ResourceSelector{openapiclient.resource.Selector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Selector: "Selector_example"}), Account: , Organizations: []OrganizationOrganizationRelationship{)})}, System: }} // SoftwareHclMeta | The 'software.HclMeta' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SoftwareApi.CreateSoftwareHclMeta(context.Background(), softwareHclMeta).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.CreateSoftwareHclMeta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSoftwareHclMeta`: SoftwareHclMeta
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.CreateSoftwareHclMeta`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSoftwareHclMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **softwareHclMeta** | [**SoftwareHclMeta**](SoftwareHclMeta.md) | The &#39;software.HclMeta&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**SoftwareHclMeta**](software.HclMeta.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSoftwareHyperflexDistributable

> SoftwareHyperflexDistributable CreateSoftwareHyperflexDistributable(ctx).SoftwareHyperflexDistributable(softwareHyperflexDistributable).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'software.HyperflexDistributable' resource.

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
    softwareHyperflexDistributable := openapiclient.software.HyperflexDistributable{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", DownloadCount: int64(123), ImportAction: "ImportAction_example", ImportState: "ImportState_example", ImportedTime: "TODO", LastAccessTime: "TODO", Md5sum: "Md5sum_example", Name: "Name_example", ReleaseDate: "TODO", Sha512sum: "Sha512sum_example", Size: int64(123), SoftwareAdvisoryUrl: "SoftwareAdvisoryUrl_example", Source: openapiclient.softwarerepository.FileServer{ClassId: "ClassId_example", ObjectType: "ObjectType_example"}, Version: "Version_example", BundleType: "BundleType_example", Guid: "Guid_example", Mdfid: "Mdfid_example", Model: "Model_example", PlatformType: "PlatformType_example", RecommendedBuild: "RecommendedBuild_example", ReleaseNotesUrl: "ReleaseNotesUrl_example", SoftwareTypeId: "SoftwareTypeId_example", SupportedModels: []string{"SupportedModels_example"), Vendor: "Vendor_example", Catalog: openapiclient.softwarerepository.Catalog.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Organization: , System: }} // SoftwareHyperflexDistributable | The 'software.HyperflexDistributable' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SoftwareApi.CreateSoftwareHyperflexDistributable(context.Background(), softwareHyperflexDistributable).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.CreateSoftwareHyperflexDistributable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSoftwareHyperflexDistributable`: SoftwareHyperflexDistributable
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.CreateSoftwareHyperflexDistributable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSoftwareHyperflexDistributableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **softwareHyperflexDistributable** | [**SoftwareHyperflexDistributable**](SoftwareHyperflexDistributable.md) | The &#39;software.HyperflexDistributable&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**SoftwareHyperflexDistributable**](software.HyperflexDistributable.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSoftwareSolutionDistributable

> SoftwareSolutionDistributable CreateSoftwareSolutionDistributable(ctx).SoftwareSolutionDistributable(softwareSolutionDistributable).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'software.SolutionDistributable' resource.

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
    softwareSolutionDistributable := openapiclient.software.SolutionDistributable{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", DownloadCount: int64(123), ImportAction: "ImportAction_example", ImportState: "ImportState_example", ImportedTime: "TODO", LastAccessTime: "TODO", Md5sum: "Md5sum_example", Name: "Name_example", ReleaseDate: "TODO", Sha512sum: "Sha512sum_example", Size: int64(123), SoftwareAdvisoryUrl: "SoftwareAdvisoryUrl_example", Source: , Version: "Version_example", BundleType: "BundleType_example", Guid: "Guid_example", Mdfid: "Mdfid_example", Model: "Model_example", PlatformType: "PlatformType_example", RecommendedBuild: "RecommendedBuild_example", ReleaseNotesUrl: "ReleaseNotesUrl_example", SoftwareTypeId: "SoftwareTypeId_example", SupportedModels: []string{"SupportedModels_example"), Vendor: "Vendor_example", FilePath: "FilePath_example", SolutionName: "SolutionName_example", SubType: "SubType_example", Catalog: } // SoftwareSolutionDistributable | The 'software.SolutionDistributable' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SoftwareApi.CreateSoftwareSolutionDistributable(context.Background(), softwareSolutionDistributable).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.CreateSoftwareSolutionDistributable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSoftwareSolutionDistributable`: SoftwareSolutionDistributable
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.CreateSoftwareSolutionDistributable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSoftwareSolutionDistributableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **softwareSolutionDistributable** | [**SoftwareSolutionDistributable**](SoftwareSolutionDistributable.md) | The &#39;software.SolutionDistributable&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**SoftwareSolutionDistributable**](software.SolutionDistributable.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSoftwareHclMeta

> DeleteSoftwareHclMeta(ctx, moid).Execute()

Delete a 'software.HclMeta' resource.

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
    resp, r, err := api_client.SoftwareApi.DeleteSoftwareHclMeta(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.DeleteSoftwareHclMeta``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSoftwareHclMetaRequest struct via the builder pattern


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


## DeleteSoftwareHyperflexDistributable

> DeleteSoftwareHyperflexDistributable(ctx, moid).Execute()

Delete a 'software.HyperflexDistributable' resource.

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
    resp, r, err := api_client.SoftwareApi.DeleteSoftwareHyperflexDistributable(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.DeleteSoftwareHyperflexDistributable``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSoftwareHyperflexDistributableRequest struct via the builder pattern


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


## DeleteSoftwareSolutionDistributable

> DeleteSoftwareSolutionDistributable(ctx, moid).Execute()

Delete a 'software.SolutionDistributable' resource.

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
    resp, r, err := api_client.SoftwareApi.DeleteSoftwareSolutionDistributable(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.DeleteSoftwareSolutionDistributable``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSoftwareSolutionDistributableRequest struct via the builder pattern


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


## GetSoftwareHclMetaByMoid

> SoftwareHclMeta GetSoftwareHclMetaByMoid(ctx, moid).Execute()

Read a 'software.HclMeta' resource.

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
    resp, r, err := api_client.SoftwareApi.GetSoftwareHclMetaByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.GetSoftwareHclMetaByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoftwareHclMetaByMoid`: SoftwareHclMeta
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.GetSoftwareHclMetaByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoftwareHclMetaByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SoftwareHclMeta**](software.HclMeta.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoftwareHclMetaList

> SoftwareHclMetaResponse GetSoftwareHclMetaList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'software.HclMeta' resource.

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
    resp, r, err := api_client.SoftwareApi.GetSoftwareHclMetaList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.GetSoftwareHclMetaList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoftwareHclMetaList`: SoftwareHclMetaResponse
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.GetSoftwareHclMetaList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSoftwareHclMetaListRequest struct via the builder pattern


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

[**SoftwareHclMetaResponse**](software.HclMeta.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoftwareHyperflexDistributableByMoid

> SoftwareHyperflexDistributable GetSoftwareHyperflexDistributableByMoid(ctx, moid).Execute()

Read a 'software.HyperflexDistributable' resource.

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
    resp, r, err := api_client.SoftwareApi.GetSoftwareHyperflexDistributableByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.GetSoftwareHyperflexDistributableByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoftwareHyperflexDistributableByMoid`: SoftwareHyperflexDistributable
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.GetSoftwareHyperflexDistributableByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoftwareHyperflexDistributableByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SoftwareHyperflexDistributable**](software.HyperflexDistributable.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoftwareHyperflexDistributableList

> SoftwareHyperflexDistributableResponse GetSoftwareHyperflexDistributableList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'software.HyperflexDistributable' resource.

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
    resp, r, err := api_client.SoftwareApi.GetSoftwareHyperflexDistributableList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.GetSoftwareHyperflexDistributableList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoftwareHyperflexDistributableList`: SoftwareHyperflexDistributableResponse
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.GetSoftwareHyperflexDistributableList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSoftwareHyperflexDistributableListRequest struct via the builder pattern


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

[**SoftwareHyperflexDistributableResponse**](software.HyperflexDistributable.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoftwareSolutionDistributableByMoid

> SoftwareSolutionDistributable GetSoftwareSolutionDistributableByMoid(ctx, moid).Execute()

Read a 'software.SolutionDistributable' resource.

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
    resp, r, err := api_client.SoftwareApi.GetSoftwareSolutionDistributableByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.GetSoftwareSolutionDistributableByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoftwareSolutionDistributableByMoid`: SoftwareSolutionDistributable
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.GetSoftwareSolutionDistributableByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoftwareSolutionDistributableByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SoftwareSolutionDistributable**](software.SolutionDistributable.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoftwareSolutionDistributableList

> SoftwareSolutionDistributableResponse GetSoftwareSolutionDistributableList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'software.SolutionDistributable' resource.

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
    resp, r, err := api_client.SoftwareApi.GetSoftwareSolutionDistributableList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.GetSoftwareSolutionDistributableList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoftwareSolutionDistributableList`: SoftwareSolutionDistributableResponse
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.GetSoftwareSolutionDistributableList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSoftwareSolutionDistributableListRequest struct via the builder pattern


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

[**SoftwareSolutionDistributableResponse**](software.SolutionDistributable.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSoftwareHclMeta

> SoftwareHclMeta PatchSoftwareHclMeta(ctx, moid).SoftwareHclMeta(softwareHclMeta).IfMatch(ifMatch).Execute()

Update a 'software.HclMeta' resource.

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
    softwareHclMeta := openapiclient.software.HclMeta{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", DownloadCount: int64(123), ImportAction: "ImportAction_example", ImportState: "ImportState_example", ImportedTime: "TODO", LastAccessTime: "TODO", Md5sum: "Md5sum_example", Name: "Name_example", ReleaseDate: "TODO", Sha512sum: "Sha512sum_example", Size: int64(123), SoftwareAdvisoryUrl: "SoftwareAdvisoryUrl_example", Source: , Version: "Version_example", BundleType: "BundleType_example", Guid: "Guid_example", Mdfid: "Mdfid_example", Model: "Model_example", PlatformType: "PlatformType_example", RecommendedBuild: "RecommendedBuild_example", ReleaseNotesUrl: "ReleaseNotesUrl_example", SoftwareTypeId: "SoftwareTypeId_example", SupportedModels: []string{"SupportedModels_example"), Vendor: "Vendor_example", ContentType: "ContentType_example", Catalog: } // SoftwareHclMeta | The 'software.HclMeta' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SoftwareApi.PatchSoftwareHclMeta(context.Background(), moid, softwareHclMeta).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.PatchSoftwareHclMeta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSoftwareHclMeta`: SoftwareHclMeta
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.PatchSoftwareHclMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSoftwareHclMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **softwareHclMeta** | [**SoftwareHclMeta**](SoftwareHclMeta.md) | The &#39;software.HclMeta&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**SoftwareHclMeta**](software.HclMeta.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSoftwareHyperflexDistributable

> SoftwareHyperflexDistributable PatchSoftwareHyperflexDistributable(ctx, moid).SoftwareHyperflexDistributable(softwareHyperflexDistributable).IfMatch(ifMatch).Execute()

Update a 'software.HyperflexDistributable' resource.

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
    softwareHyperflexDistributable := openapiclient.software.HyperflexDistributable{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", DownloadCount: int64(123), ImportAction: "ImportAction_example", ImportState: "ImportState_example", ImportedTime: "TODO", LastAccessTime: "TODO", Md5sum: "Md5sum_example", Name: "Name_example", ReleaseDate: "TODO", Sha512sum: "Sha512sum_example", Size: int64(123), SoftwareAdvisoryUrl: "SoftwareAdvisoryUrl_example", Source: , Version: "Version_example", BundleType: "BundleType_example", Guid: "Guid_example", Mdfid: "Mdfid_example", Model: "Model_example", PlatformType: "PlatformType_example", RecommendedBuild: "RecommendedBuild_example", ReleaseNotesUrl: "ReleaseNotesUrl_example", SoftwareTypeId: "SoftwareTypeId_example", SupportedModels: []string{"SupportedModels_example"), Vendor: "Vendor_example", Catalog: } // SoftwareHyperflexDistributable | The 'software.HyperflexDistributable' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SoftwareApi.PatchSoftwareHyperflexDistributable(context.Background(), moid, softwareHyperflexDistributable).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.PatchSoftwareHyperflexDistributable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSoftwareHyperflexDistributable`: SoftwareHyperflexDistributable
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.PatchSoftwareHyperflexDistributable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSoftwareHyperflexDistributableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **softwareHyperflexDistributable** | [**SoftwareHyperflexDistributable**](SoftwareHyperflexDistributable.md) | The &#39;software.HyperflexDistributable&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**SoftwareHyperflexDistributable**](software.HyperflexDistributable.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSoftwareSolutionDistributable

> SoftwareSolutionDistributable PatchSoftwareSolutionDistributable(ctx, moid).SoftwareSolutionDistributable(softwareSolutionDistributable).IfMatch(ifMatch).Execute()

Update a 'software.SolutionDistributable' resource.

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
    softwareSolutionDistributable := openapiclient.software.SolutionDistributable{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", DownloadCount: int64(123), ImportAction: "ImportAction_example", ImportState: "ImportState_example", ImportedTime: "TODO", LastAccessTime: "TODO", Md5sum: "Md5sum_example", Name: "Name_example", ReleaseDate: "TODO", Sha512sum: "Sha512sum_example", Size: int64(123), SoftwareAdvisoryUrl: "SoftwareAdvisoryUrl_example", Source: , Version: "Version_example", BundleType: "BundleType_example", Guid: "Guid_example", Mdfid: "Mdfid_example", Model: "Model_example", PlatformType: "PlatformType_example", RecommendedBuild: "RecommendedBuild_example", ReleaseNotesUrl: "ReleaseNotesUrl_example", SoftwareTypeId: "SoftwareTypeId_example", SupportedModels: []string{"SupportedModels_example"), Vendor: "Vendor_example", FilePath: "FilePath_example", SolutionName: "SolutionName_example", SubType: "SubType_example", Catalog: } // SoftwareSolutionDistributable | The 'software.SolutionDistributable' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SoftwareApi.PatchSoftwareSolutionDistributable(context.Background(), moid, softwareSolutionDistributable).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.PatchSoftwareSolutionDistributable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSoftwareSolutionDistributable`: SoftwareSolutionDistributable
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.PatchSoftwareSolutionDistributable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSoftwareSolutionDistributableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **softwareSolutionDistributable** | [**SoftwareSolutionDistributable**](SoftwareSolutionDistributable.md) | The &#39;software.SolutionDistributable&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**SoftwareSolutionDistributable**](software.SolutionDistributable.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSoftwareHclMeta

> SoftwareHclMeta UpdateSoftwareHclMeta(ctx, moid).SoftwareHclMeta(softwareHclMeta).IfMatch(ifMatch).Execute()

Update a 'software.HclMeta' resource.

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
    softwareHclMeta :=  // SoftwareHclMeta | The 'software.HclMeta' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SoftwareApi.UpdateSoftwareHclMeta(context.Background(), moid, softwareHclMeta).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.UpdateSoftwareHclMeta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSoftwareHclMeta`: SoftwareHclMeta
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.UpdateSoftwareHclMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSoftwareHclMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **softwareHclMeta** | [**SoftwareHclMeta**](SoftwareHclMeta.md) | The &#39;software.HclMeta&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**SoftwareHclMeta**](software.HclMeta.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSoftwareHyperflexDistributable

> SoftwareHyperflexDistributable UpdateSoftwareHyperflexDistributable(ctx, moid).SoftwareHyperflexDistributable(softwareHyperflexDistributable).IfMatch(ifMatch).Execute()

Update a 'software.HyperflexDistributable' resource.

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
    softwareHyperflexDistributable :=  // SoftwareHyperflexDistributable | The 'software.HyperflexDistributable' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SoftwareApi.UpdateSoftwareHyperflexDistributable(context.Background(), moid, softwareHyperflexDistributable).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.UpdateSoftwareHyperflexDistributable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSoftwareHyperflexDistributable`: SoftwareHyperflexDistributable
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.UpdateSoftwareHyperflexDistributable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSoftwareHyperflexDistributableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **softwareHyperflexDistributable** | [**SoftwareHyperflexDistributable**](SoftwareHyperflexDistributable.md) | The &#39;software.HyperflexDistributable&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**SoftwareHyperflexDistributable**](software.HyperflexDistributable.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSoftwareSolutionDistributable

> SoftwareSolutionDistributable UpdateSoftwareSolutionDistributable(ctx, moid).SoftwareSolutionDistributable(softwareSolutionDistributable).IfMatch(ifMatch).Execute()

Update a 'software.SolutionDistributable' resource.

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
    softwareSolutionDistributable :=  // SoftwareSolutionDistributable | The 'software.SolutionDistributable' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SoftwareApi.UpdateSoftwareSolutionDistributable(context.Background(), moid, softwareSolutionDistributable).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.UpdateSoftwareSolutionDistributable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSoftwareSolutionDistributable`: SoftwareSolutionDistributable
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.UpdateSoftwareSolutionDistributable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSoftwareSolutionDistributableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **softwareSolutionDistributable** | [**SoftwareSolutionDistributable**](SoftwareSolutionDistributable.md) | The &#39;software.SolutionDistributable&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**SoftwareSolutionDistributable**](software.SolutionDistributable.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

