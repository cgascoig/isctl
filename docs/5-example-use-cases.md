# Example Use Cases

## Check TPM status on all servers

```
isctl get equipment tpm  --expand 'ComputeBoard($expand=ComputeRackUnit,ComputeBlade)' -o custom-columns=Serial:.Serial,BladeSerial:.ComputeBoard.ComputeBlade.Serial,RackSerial:.ComputeBoard.ComputeRackUnit.Serial,ActivationStatus:.ActivationStatus,AdminState:.AdminState
```
Output:
```
----------------  ----------------  ----------------  ---------------------  ---------------
         Serial       BladeSerial        RackSerial       ActivationStatus       AdminState
----------------  ----------------  ----------------  ---------------------  ---------------
             NA            <none>       FCH2035V0CP                unknown          unknown
    FCH1807J9V6       FCH18097X4Z            <none>                unknown          unknown
    FCH1803JEVA       FCH18097X4Z            <none>              activated          enabled
             NA            <none>       FCH2002V21L                unknown          unknown
    FCH221173Z3            <none>       WZP21520P1Y                     NA          enabled
             NA            <none>       FCH1810V1C4                unknown          unknown
    FCH2041774F       FCH2046736C            <none>              activated          enabled
    FCH20337FE5       FCH210977GZ            <none>              activated          enabled
             NA            <none>       FCH2002V21J                unknown          unknown
             NA            <none>       FCH2002V1NL                unknown          unknown
             NA            <none>       FCH2132V15A                unknown          unknown
             NA            <none>       WZP22440MUW                unknown          unknown
             NA            <none>       WZP22441AT6                unknown          unknown
             NA            <none>       FCH2108V0NU                unknown          unknown
             NA            <none>       FCH2002V1ZP                unknown          unknown
    FCH1906J0U6            <none>       FCH1902V1FX              activated          enabled
----------------  ----------------  ----------------  ---------------------  --------------- 
```

## Check version and support status of Kubernetes clusters

```
isctl get kubernetes clusterprofile --expand 'NodeGroups($expand=KubernetesVersion($expand=Version))' -o custom-columns='Name:.Name,Version:.NodeGroups[0].KubernetesVersion.Version.Name,Tags:.NodeGroups[0].KubernetesVersion.Version.Tags'
```
Output:
```
----------------------------  ------------------------------------  ---------------------------------------------------------
                       Name                               Version                                                       Tags
----------------------------  ------------------------------------  ---------------------------------------------------------
             cisg-team-call                         1.19.16-iks.0       intersight.kubernetes.SupportStatus: UpgradeRequired
            grscarle-cpoc-1                         1.21.10-iks.0            intersight.kubernetes.SupportStatus: Deprecated
         tf-iks-cpoc-demo-1                         1.20.14-iks.4
               cg-hcrs-demo                         1.21.11-iks.2
```

## Performing consistent updates on many managed objects

This example shows how to iterate over all NTP Policy managed objects (using `--filter` to select a subset is probably advised) and update the NtpServers attribute on all of them:
```
for moid in $(isctl get ntp policy --jsonpath '$[*].Moid'); do 
  isctl update ntp policy moid ${moid} --NtpServers 1.1.1.1
done
```

## Script to repair IKS cluster and vCenter target linkage

[This](https://github.com/cgascoig/isctl/blob/devel/examples/scripts/iks-fix-vcenter-target.sh) example script will re-establish the linkage between a deployed IKS cluster and a vCenter target. This is useful in scenarios where the vCenter target was removed and re-added while an IKS cluster was deployed. 