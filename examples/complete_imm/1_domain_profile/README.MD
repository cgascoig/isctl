## Pre Requisites
Make sure you have [isctl](https://github.com/cgascoig/isctl) installed and configured.

## Steps
1. Claim FI Targets, this is a manual UI process for now.

2. Edit all the yamls file to match your environment

3. Apply all the config
```
isctl apply -f .
```

4. Assign the Domain Profile
Through the UI OR 

isctl update fabric switchprofile name IMM-Domain-01-A --AssignedSwitch <switch-moid>
isctl update fabric switchprofile name IMM-Domain-01-B --AssignedSwitch <switch-moid>


5. Deploy the Domain Profiles
```
isctl update fabric switchprofile name IMM-Domain-01-A --Action Deploy
isctl update fabric switchprofile name IMM-Domain-01-B --Action Deploy
```

## Clean up
1. Unassign Domain Profile
```
isctl update fabric switchprofile name IMM-Domain-01-A --Action Unassign
isctl update fabric switchprofile name IMM-Domain-01-B --Action Unassign
```

2. remove the config
To clean up just rust the `isctl apply` commands but add a `-d` at the end.
```
isctl apply -f . -d
```