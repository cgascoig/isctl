# Advanced Intersight Queries

## Select

To improve performance, the [Intersight API](https://intersight.com/apidocs/introduction/query/#select-query-option-selecting-properties-in-the-http-response) supports a `select` query parameter that tells Intersight which attributes to include in the returned object(s). To use this with `isctl`, use the `--select <attribute_list>` option. Where `<attribute_list>` is a comma separated list of attribute names. For example:

```
isctl get ntp policy --select Name,NtpServers
```
Output:

```
-------------  -----------------------------  ----------------------
        Name                           Moid              NtpServers 
-------------  -----------------------------  ----------------------
     NTPTest       626f1ad36275722d30486f94       ntp.esl.cisco.com 
    NtpTest1       626f4f476275722d304c406b        1.1.1.1, 2.2.2.2 
    NtpTest2       626f4f4f6275722d304c40df        1.1.1.1, 2.2.2.2 
```

## Expand

In the Intersight API, many objects contain references to other objects. For example, the Organization attribute is reference to another organization.Organization object:

```
isctl get ntp policy --name cg-tf-ntp-test  -o yaml --select Name,Organization
```
Output:
```
ClassId: ntp.Policy
Moid: 61e9f9696275722d31e74cfd
Name: cg-tf-ntp-test
ObjectType: ntp.Policy
Organization:
  ClassId: mo.MoRef
  Moid: 123456789012345678901234
  ObjectType: organization.Organization
  link: https://www.intersight.com/api/v1/organization/Organizations/5deed7fc6972652d33bc48d0
```

To improve performance and remove the need for multiple API calls to get information from these referenced objects, the API can [expand](https://intersight.com/apidocs/introduction/query/#expand-query-option-returning-related-resources) these objects inline. To use this capability in `isctl`, use the `--expand <attribute_list>` option. For example:

```
isctl get ntp policy --name cg-tf-ntp-test  -o yaml --select Name,Organization --expand Organization
```
Output:
```
ClassId: ntp.Policy
Moid: 61e9f9696275722d31e74cfd
Name: cg-tf-ntp-test
ObjectType: ntp.Policy
Organization:
  Account:
    ClassId: mo.MoRef
    Moid: 123456789012345678901234
    ObjectType: iam.Account
    link: https://www.intersight.com/api/v1/iam/Accounts/123456789012345678901234
  AccountMoid: 123456789012345678901234
  Ancestors: []
  ClassId: organization.Organization
  CreateTime: "2019-12-09T23:25:48.249Z"
  Description: All Equipment in Lab
  DomainGroupMoid: 5b25418d7a7662743465cf72
  ModTime: "2019-12-09T23:25:48.256Z"
  Moid: 5deed7fc6972652d33bc48d0
  Name: Example Organisation
  ObjectType: organization.Organization
  Owners:
  - 123456789012345678901234
  PermissionResources:
  - ClassId: mo.MoRef
    Moid: 123456789012345678901234
    ObjectType: organization.Organization
    link: https://www.intersight.com/api/v1/organization/Organizations/123456789012345678901234
  ResourceGroups:
  - ClassId: mo.MoRef
    Moid: 123456789012345678901234
    ObjectType: resource.Group
    link: https://www.intersight.com/api/v1/resource/Groups/123456789012345678901234
  SharedScope: ""
  Tags: []
```

Notice in the example above, the Organization attribute now includes the entire organization.Organization object nested in the result. These nested attributes can be used in the same way as any of the normal attributes, for example, they can be used in the custom-colums spec:

```
isctl get ntp policy --name cg-tf-ntp-test --select Name,Organization --expand Organization --output custom-columns=Name:.Name,OrgDescription:.Organization.Description
```
Output:
```
-------------------  -------------------------
              Name             OrgDescription
-------------------  -------------------------
    cg-tf-ntp-test       All Equipment in Lab
-------------------  -------------------------
```

## Pagination

By default, the Intersight API returns the first 100 results for a query. To enable pagination for larger result sets, the API supports [top, skip and orderBy](https://intersight.com/apidocs/introduction/query/#top-and-skip-query-options-pagination) parameters. To use this in `isctl`, use the `--top`, `--skip` and `--orderby` options. For example, to order the results set by name and get the first 2 results:

```
isctl get ntp policy --select Name --orderby Name --top 2
```
Output:
```
------------------------  -----------------------------
                   Name                           Moid
------------------------  -----------------------------
                 BA-NTP       5cc6f2426275722d30fbd461
    BulkTest-ntp-policy       61b022e96275722d3176387f
------------------------  -----------------------------
```

Then, to get the next 2 results:
```
isctl get ntp policy --select Name --orderby Name --top 2 --skip 2
```
Output:
```
-------------  -----------------------------
        Name                           Moid
-------------  -----------------------------
    CPOC-NTP       607fb51a6275722d3164130b
    CiscoNTP       5b21faf1666c663463a8341f
-------------  -----------------------------
```