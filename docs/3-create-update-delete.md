# Creating, Updating and Deleting Intersight Objects
## Creating Intersight resources

To create resources in Intersight, use the `isctl create ...` commands. 

The attributes for the new object can be passed in as command line flags:

```
isctl create ntp policy --Name isctl-test-1 --NtpServers 1.1.1.1,2.2.2.2 --Organization default
```

Output:
```
2020/07/19 13:02:13 Single result, falling back to vertical output. NOTE: this is not valid YAML; use --output yaml to get valid YAML.
Enabled: true
Moid: 5f13b7b56275722d31133e63
Name: isctl-test-1
NtpServers:
- 1.1.1.1
- 2.2.2.2
```

The output of the `create` commands is the resource after it was created in Intersight, so the Moid, etc. are included.

You can also define the attributes of the new object via standard input:

```
isctl create ntp policy --bodyformat json
Waiting for JSON body:
{
    "Name": "isctl-ntp-test",
    "Enabled": true,
    "NtpServers": [
        "1.1.1.1"
    ],
    "Organization": {
        "Moid": "123456789012345678901234",
        "ClassId": "mo.MoRef"
    }
}
```

This is useful when you have your object defined as a JSON or YAML file, for example:

```
cat ntp_policy.json | isctl create ntp policy --bodyformat json
```

## Deleting resources

Deleting resources is quite simple if you know the Moid:

```
isctl delete ntp policy moid 123456789012345678901234
```

Or the Name:

```
isctl delete ntp policy name isctl-test-1
```

## Updating resources

Resources can also be updated by Moid or Name. For example, if you want to set the `Enabled` attribute in an NTP policy to `False`:

```
isctl update ntp policy moid 123456789012345678901234 --Enabled=False
```

Or by Name:

```
isctl update ntp policy name isctl-test-1 --Enabled=False
```

