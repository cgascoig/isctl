## Pre Requisites
Make sure you have [isctl](https://github.com/cgascoig/isctl) installed and configured.

## Steps
1. Update the yaml files with your environment settings

2. Apply the config
```
isctl apply -f .
```

3. Deploy the Server Profiles
```
isctl update server profile name IMM-RHEL-01 --ScheduledActions '[{"Action":"Deploy","ProceedOnReboot":false},{"Action":"Activate","ProceedOnReboot":true}]'
isctl update server profile name IMM-RHEL-02 --ScheduledActions '[{"Action":"Deploy","ProceedOnReboot":false},{"Action":"Activate","ProceedOnReboot":true}]'
isctl update server profile name IMM-RHEL-03 --ScheduledActions '[{"Action":"Deploy","ProceedOnReboot":false},{"Action":"Activate","ProceedOnReboot":true}]'
isctl update server profile name IMM-RHEL-04 --ScheduledActions '[{"Action":"Deploy","ProceedOnReboot":false},{"Action":"Activate","ProceedOnReboot":true}]'
isctl update server profile name IMM-RHEL-05 --ScheduledActions '[{"Action":"Deploy","ProceedOnReboot":false},{"Action":"Activate","ProceedOnReboot":true}]'
isctl update server profile name IMM-RHEL-06 --ScheduledActions '[{"Action":"Deploy","ProceedOnReboot":false},{"Action":"Activate","ProceedOnReboot":true}]'
isctl update server profile name IMM-RHEL-07 --ScheduledActions '[{"Action":"Deploy","ProceedOnReboot":false},{"Action":"Activate","ProceedOnReboot":true}]'
isctl update server profile name IMM-RHEL-08 --ScheduledActions '[{"Action":"Deploy","ProceedOnReboot":false},{"Action":"Activate","ProceedOnReboot":true}]'
```

## Clean up
1. Unassign & delete Server Profiles
```
isctl update server profile name IMM-RHEL-01 --Action Unassign
isctl update server profile name IMM-RHEL-02 --Action Unassign
isctl update server profile name IMM-RHEL-03 --Action Unassign
isctl update server profile name IMM-RHEL-04 --Action Unassign
isctl update server profile name IMM-RHEL-05 --Action Unassign
isctl update server profile name IMM-RHEL-06 --Action Unassign
isctl update server profile name IMM-RHEL-07 --Action Unassign
isctl update server profile name IMM-RHEL-08 --Action Unassign

# Wait for the undeploy workflow to finish

isctl delete server profile name IMM-RHEL-01
isctl delete server profile name IMM-RHEL-02
isctl delete server profile name IMM-RHEL-03
isctl delete server profile name IMM-RHEL-04
isctl delete server profile name IMM-RHEL-05
isctl delete server profile name IMM-RHEL-06
isctl delete server profile name IMM-RHEL-07
isctl delete server profile name IMM-RHEL-08
```

2. remove the config
To clean up just rust the `isctl apply` commands but add a `-d` at the end.
```
isctl apply -f . -d
```