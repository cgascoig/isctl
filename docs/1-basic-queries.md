# Basic Intersight Queries

## Basic queries

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

## Filtering results

The Intersight API can filter results before returning them to the client. To specify the filter query, use `--filter <filter_expression>`. See the [Intersight API Guide](https://intersight.com/apidocs/introduction/query/#filter-query-option-filtering-the-resources) for details on the `<filter_expression>` syntax. Some examples are below.

Filter on exact name:
```
isctl get ntp policy --filter "Name eq 'cg-tf-ntp-test'"
```

Filter objects where name starts with "cg-":
```
isctl get ntp policy --filter "startsWith(Name, 'cg-')"
```

## Output customisation

By default, `isctl` will try to produce human readable output. Typically, this will be a table format with boilerplate attributes (e.g. `ClassId`, `Organization`, etc.) hidden for brevity. If the output table would have too many columns, `isctl` will fall back to a "vertical" output (essentially YAML). If you want to disable this behaviour and force the table format (even though it will be very wide), use `--output table`. 

There are a number of other options to customise the output described below.

### JSON/YAML

Using `--output json` or `--output yaml` will cause the output to use the chosen format and will disable the hiding of any "boilerplate" attributes:

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

This can be useful to feed the structured output to another tool, for example `jq`:

```
isctl get ntp policy -o json | jq -r '.[].Name'
```
Output:
```
NTP_ESL
CiscoNTP
BA-NTP
```

### JSONPath
You can use [JSONPath](https://goessner.net/articles/JsonPath/) to filter/transform the returned data. For example, to output just the name attribute of all the returned objects:

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

### Cusom Columns
To customise the table output, "custom-columns" output can be used. The syntax is `--output custom-columns=<spec>` where `<spec>` is a comma separated list of `<column_title>:<value_path>` column specifications. The `<value_path>` is a simplified JSONPath expression that specifies how to extract the column value from a returned item. For example:

```
isctl get ntp policy --output custom-columns=NAME:.Name,ENABLED:.Enabled
```
Output:
```
-------------  ------------
        NAME       ENABLED
-------------  ------------
     NTP_ESL          True
    CiscoNTP          True
      BA-NTP          True
-------------  ------------
```

### CSV
The output can be formatted as a CSV for easy importing into spreadsheets or other applications. The syntax for CSV output is similar to custom-columns above - `--output csv=<spec>`. For example:

```
isctl get ntp policy --output csv=NAME:.Name,ENABLED:.Enabled
```
Output:
```
"NAME","ENABLED"
"NTP_ESL","True"
"CiscoNTP","True"
"BA-NTP","True"
```

A common way to use this with the default spreadsheet application is:

```
isctl get ntp policy --output csv=NAME:.Name,ENABLED:.Enabled  > /tmp/ntp.csv && open /tmp/ntp.csv
```

