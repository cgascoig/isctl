# Basic Intersight Queries

To retrieve resources from Intersight, use the `isctl get ...` commands. 

For example, to query all objects of a certain type:

```
isctl get ntp policy
```
Output:
```
------------  -----------------------------  ------------------------  -----------------------------
    Enabled                           Moid                      Name                     NtpServers
------------  -----------------------------  ------------------------  -----------------------------
      False       5ecb069c6275722d3143e668           test-ntp-policy

       True       5ee1aa076275722d3122a944          cg-tf-ntp-test-1       10.10.10.10, 10.10.10.12
------------  -----------------------------  ------------------------  -----------------------------
```
Notice that all objects are returned in a human-readable table. 

You can query a single object based on its MoId:

```
isctl get ntp policy moid 5ee1aa076275722d3122a944
```

Or its name:
```
isctl get ntp policy --name cg-tf-ntp-test-1
```

Output:
```
2020/06/25 17:40:04 Single result, falling back to vertical output. NOTE: this is not valid YAML; use --output yaml to get valid YAML.
Enabled: true
Moid: 5ee1aa076275722d3122a944
Name: cg-tf-ntp-test-1
NtpServers:
- 10.10.10.10
- 10.10.10.12
```
Notice that when displaying a single object, table format is not used. 

The default outputs above automatically hide many of the boilerplate attributes of the objects (e.g. `ClassId`, `Organization`, etc.). If you specify JSON or YAML as the output format all attributes will be included. For example:

```
isctl get ntp policy moid 5ee1aa076275722d3122a944 --output json
```
Output:
```json
{
  "AccountMoid": "123456789012345678901234",
  "Ancestors": [],
  "ClassId": "ntp.Policy",
  "CreateTime": "2020-06-11T03:50:32.062Z",
  "Description": "",
  "DomainGroupMoid": "123456789012345678901234",
  "Enabled": true,
  "ModTime": "2020-06-11T03:50:32.063Z",
  "Moid": "5ee1aa076275722d3122a944",
  "Name": "cg-tf-ntp-test-1",
  "NtpServers": [
    "10.10.10.10",
    "10.10.10.12"
  ],
  "ObjectType": "ntp.Policy",
  "Organization": {
    "ClassId": "mo.MoRef",
    "Moid": "123456789012345678901234",
    "ObjectType": "organization.Organization",
    "link": "https://www.intersight.com/api/v1/organization/Organizations/123456789012345678901234"
  },
  "Owners": [
    "123456789012345678901234"
  ],
  "PermissionResources": [
    {
      "ClassId": "mo.MoRef",
      "Moid": "123456789012345678901234",
      "ObjectType": "organization.Organization",
      "link": "https://www.intersight.com/api/v1/organization/Organizations/123456789012345678901234"
    }
  ],
  "Profiles": [],
  "SharedScope": "",
  "Tags": []
}
```

You can use JSONPath to filter/transform the returned data - to output just the name attribute of all the returned objects:

```
isctl get ntp policy -o yaml --jsonpath '$[*].Name'
```
Output:
```
- test-ntp-policy
- cg-tf-ntp-test-1
```

Or even to get multiple attributes from the returned objects:

```
isctl get ntp policy moid 5ee1aa076275722d3122a944 --jsonpath '$["Name","Enabled"]'
```
Output:
```
cg-tf-ntp-test-1
True
```