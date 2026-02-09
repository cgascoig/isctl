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

## Editing resources interactively

For more complex edits, or when you want to review the current state before making changes, use the `isctl edit` command. This opens the resource in your preferred text editor, similar to `kubectl edit`.

```
isctl edit ntp policy name isctl-test-1
```

This will:
1. Fetch the current resource from Intersight
2. Filter to show only editable properties (read-only fields like Moid and CreateTime are hidden)
3. Open it in your editor as YAML
4. After you save and close, apply any changes via an update operation

If you close the editor without making changes, no update is performed.

### Editor selection

The editor is determined by (in order of precedence):
1. `$EDITOR` environment variable
2. `$VISUAL` environment variable
3. Platform default (`vi` on Unix, `notepad` on Windows)

To use a specific editor:

```
EDITOR=nano isctl edit ntp policy name isctl-test-1
```

