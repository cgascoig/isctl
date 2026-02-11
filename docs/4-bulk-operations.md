# Bulk Operations

You can manage (create/update/destroy) a collection of Intersight resources in bulk using the `apply` command. 

First, create your resources in a YAML file (or files) (see [here](examples) for more examples). For example, the following manages an Organization and NTP Policy:

```
ClassId: ntp.Policy
ObjectType: ntp.Policy
Name: isctl-test
Enabled: true
NtpServers: 
    - 1.1.1.1
Organization: MoRef[Name:isctl-test]
---
ClassId: organization.Organization
ObjectType: organization.Organization
Name: isctl-test
```

The order is not important - `isctl` will try to apply the individual resources in the correct order based on any dependencies between them. 

To apply (create or update) these to Intersight, run: 

```
isctl apply -f examples/ntp_and_organization/all-in-one/ntp_organization.yaml
```

You can specify individual files or you can specify a directory and all the contained YAML files will be applied automatically:

```
isctl apply -f examples/ntp_and_organization/file-per-resource/
```

`isctl` will automatically create/update each resource:

```
INFO[0001] Performing update operation on existing MO (Name: isctl-test, Moid: 60138f066972652d33286fa0, ClassId: organization.Organization)
INFO[0002] Performing create operation on new MO (Name: isctl-test, ClassId: ntp.Policy)
Apply completed successfully
```

You can also bulk destroy resources from the YAML definitions:

```
isctl apply --delete -f examples/ntp_and_organization/file-per-resource/
```

Output:
```
INFO[0000] Performing delete operation on existing MO (Name: isctl-test, Moid: 60138f386275722d31241642, ClassId: ntp.Policy)
INFO[0002] Performing delete operation on existing MO (Name: isctl-test, Moid: 60138f066972652d33286fa0, ClassId: organization.Organization)
Destroy completed successfully
```

## Specifying relationships

In the bulk YAML files or when creating/updating individual resources, there are a few ways to specify relationships between resources. 

You can set the relationship explicitly with the ClassId and Moid:

```
Organization:
  ClassId: mo.MoRef
  Moid: 123456789012345678901234
```

Or you can specify the relationship using the MoRef syntax. For example, the following will find the Organization with Name "isctl-test" and automatically fill in the Moid:

```
Organization: MoRef[Name:isctl-test]
```

There are a few variants of the MoRef syntax to allow more flexibility - the following are all equivalent to each other:

```
# if the field is a relationship, isctl will assume a bare string is the name to query
Organization: isctl-test

# explicitly specify this is an MoRef
Organization: MoRef[isctl-test] 

# explicitly specify that the field to match the destination by is Name (you can use other fields)
Organization: MoRef[Name:isctl-test]

# this is only needed where the field has a relationship to an abstract class and you need to specify which concrete class
Organization: MoRef:OrganizationOrganizationRelationship[Name:isctl-test] 

# a common use-case for this is the PolicyBucket attribute which is a list of references to different policy types. For example, to manage a Server Profile Template:
Name: cgascoig-isctl-test
ClassId: server.ProfileTemplate
Organization: default
PolicyBucket:
  - MoRef:BiosPolicyRelationship[cgascoig-bios-policy]
  - MoRef:IamLdapPolicyRelationship[cgascoig-ldap-policy] 
  - MoRef:DeviceconnectorPolicyRelationship[cgascoig-deviceconnector-policy] 
  - MoRef:IpmioverlanPolicyRelationship[cgascoig-ipmi-policy]
```

## Renaming
For each object in your YAML file, the Name attribute is used to check if the object already exists and determine whether to apply an update or create operation for that object. For this reason, you cannot rename an existing object - isctl will not know about the original name for the object and will simply create a new object and leave the original untouched. 

## Templating and Variables

YAML files processed by `isctl apply` are treated as [Go templates](https://pkg.go.dev/text/template). This allows you to parameterize your manifests and inject values at runtime.

### Variables

You can access variables in your templates using the `{{ .Vars.VariableName }}` syntax.

Variables can be provided from three sources, in order of precedence (highest to lowest):

1.  **Command-line flags**: `--var key=value`
2.  **Variable file**: `--var-file path/to/vars.yaml`
3.  **Environment variables**: `ISCTL_VAR_variableName`
4.  **Directory variables**: `isctl.vars.yaml` or `isctl.vars.yml` in directories passed to `-f`

### Functions

[Sprig](http://masterminds.github.io/sprig/) functions are available for use in your templates. This includes functions for string manipulation, math, dates, and more.

### Example

**template.yaml**:
```yaml
ClassId: ntp.Policy
Name: {{ .Vars.policyName }}
Organization: default
# Use Sprig's default function for optional variables
Description: "Managed by {{ .Vars.owner | default "admin" | upper }}"
NtpServers:
  - {{ .Vars.ntpServer }}
```

**Apply with variables**:
```bash
# Set owner via environment variable
export ISCTL_VAR_owner="devops"

# Apply with var-file and flags
isctl apply -f template.yaml \
  --var-file common-vars.yaml \
  --var policyName="my-dynamic-policy" \
  --var ntpServer="10.0.0.1"
```

### Loops and Ranges

You can use standard Go template loops to create multiple resources from a single template definition. This is useful for creating a set of similar resources, such as a sequence of policies or profiles.

**Example: Creating multiple NTP Policies**

```yaml
{{- range $i := list 1 2 3 }}
ClassId: ntp.Policy
ObjectType: ntp.Policy
Name: ntp-policy-{{ $i }}
Enabled: true
NtpServers: 
    - 1.1.1.1
Organization: default
---
{{- end }}
```

In this example:

*   `range $i := list 1 2 3`: Iterates through the list `[1, 2, 3]`, assigning values to `$i`.
*   `Name: ntp-policy-{{ $i }}`: Dynamically sets the name for each policy (e.g., `ntp-policy-1`, `ntp-policy-2`).
*   `---`: Separates each generated YAML document. This is critical when generating multiple resources in a single output stream.
*   `{{- end }}`: Closes the loop.