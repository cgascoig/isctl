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
```

