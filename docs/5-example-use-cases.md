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

## Performing consistent updates on many managed objects

This example shows how to iterate over all NTP Policy managed objects (using `--filter` to select a subset is probably advised) and update the NtpServers attribute on all of them:
```
for moid in $(isctl get ntp policy --jsonpath '$[*].Moid'); do 
  isctl update ntp policy moid ${moid} --NtpServers 1.1.1.1
done
```
