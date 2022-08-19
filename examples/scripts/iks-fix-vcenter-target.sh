#!/bin/bash
#
# This example script will re-establish the linkage between a deployed IKS cluster and a vCenter target.
# This is useful in scenarios where the vCenter target was removed and re-added while an IKS cluster was deployed. 
#
# Contributed by Boris Aelen and Rob Moss
#

# for debugging
#set -x

config="--config .isctl.yaml"

function help () {
    echo "Usage:"
    echo "$0 -v <vcenter> -c <cluster> [-w]"
    echo "-w don't wait for the cluster to be connected"
    exit
}

while getopts ':h:c:v:w' arg
do
    case ${arg} in
        h) help
        ;;

        c) CLUSTER=${OPTARG}
        ;;

        v) VCENTER=${OPTARG}
        ;;

        w) NOWAIT=1
        ;;

        *) help
        ;;

    esac
done


if [ -z $CLUSTER ] || [ -z $VCENTER ]; then
    help
else
    echo "Reconnecting cluster $CLUSTER with vCenter $VCENTER target"
fi

echo -ne "* MOID of vCenter target with the name $VCENTER...\t\t"
TARGET_MOID=$(isctl ${config} get asset target name ${VCENTER} -o jsonpath='$.RegisteredDevice.Moid')
# if blank exit
if [ -z $TARGET_MOID ]; then
    echo "Target MOID is blank, exiting"
    exit
fi
echo "$TARGET_MOID"

echo -ne "* MOID of IKS Cluster Profile with the name $CLUSTER...\t"
CLUSTERPROFILEMOID=$(isctl ${config} get kubernetes clusterprofile name $CLUSTER -o jsonpath='$.Moid')
# if blank exit
if [ -z $CLUSTERPROFILEMOID ]; then
    echo "Cluster Profile MOID is blank, exiting"
    exit
fi
echo "$CLUSTERPROFILEMOID"

echo -ne "* Number of VM's in your IKS Cluster...\t\t\t\t"
NODEPROFILEMOIDS=$(isctl ${config} get kubernetes virtualmachinenodeprofile --filter="Ancestors/any(a:a/Moid eq '$CLUSTERPROFILEMOID')" -o jsonpath='$[*].Moid')
echo `echo $NODEPROFILEMOIDS | wc -w`
for node in $NODEPROFILEMOIDS
do
    echo -ne "\tName of the VM...\t\t\t\t\t"
    vmname=$(isctl ${config} get kubernetes virtualmachinenodeprofile moid $node -o jsonpath='$.Name')
    echo "$vmname"
    echo -ne "\tMOID of the VM in our inventory...\t\t\t"
    vmmoid=$(isctl ${config} get virtualization virtualmachine name $vmname -o jsonpath='$.Moid')
    echo "$vmmoid"
    
    echo -en "\tPatching the correct VM in the inventory...\t\t"
    isctl ${config} -o yaml update kubernetes virtualmachinenodeprofile moid $node --Target "MoRef[Moid:${TARGET_MOID}]"  > /dev/null 
    isctl ${config} -o yaml update kubernetes virtualmachinenodeprofile moid $node --VirtualMachine "MoRef[Moid:${vmmoid}]" > /dev/null
    echo "done"
done

echo -ne "* Number of Nodegroups...\t\t\t\t\t"
INFRAPROVIDERMOIDS=$(isctl ${config} get kubernetes clusterprofile --name $CLUSTER --expand="NodeGroups" -o jsonpath='$.NodeGroups[*].InfraProvider.Moid')
echo `echo $INFRAPROVIDERMOIDS | wc -w`
for provider in $INFRAPROVIDERMOIDS
do
    echo -e "\tMOID of VM InfrastructureProvider...\t\t\t$provider"
    echo -ne "\tPatching VirtualMachineInfrastructureProvider...\t"
    isctl ${config} -o yaml update kubernetes virtualmachineinfrastructureprovider moid $provider --Target "MoRef[Moid:${TARGET_MOID}]" > /dev/null
    echo "done"

    echo -ne "\tMOID of InfraConfig Policy...\t\t\t\t"
    VMINFRACONFIGPOLICY=$(isctl ${config} get kubernetes virtualmachineinfrastructureprovider moid $provider -o jsonpath='$.InfraConfigPolicy.Moid')
    echo "$VMINFRACONFIGPOLICY"

    if [ ! -z $VMINFRACONFIGPOLICY ]; then
        echo -ne "\tPatching InfraConfig Policy...\t\t\t\t"
        isctl ${config} -o yaml update kubernetes virtualmachineinfraconfigpolicy moid ${VMINFRACONFIGPOLICY} --Target "MoRef[Moid:${TARGET_MOID}]" > /dev/null 
        echo "done"
    fi
done

echo -ne "Patching IKS Cluster...\t"
isctl ${config} -o yaml update kubernetes cluster name $CLUSTER --RegisteredDevices "[{\"ClassId\":\"mo.MoRef\", \"ObjectType\":\"asset.DeviceRegistration\", \"Moid\":\"${TARGET_MOID}\"}]"  
echo "done"

if [ -z $NOWAIT ]; then
    seconds=0
    echo "* Waiting for the cluster to come back online, this can take 10 minutes"
    while true;
    do
        STATUS=$(isctl ${config} get kubernetes cluster name ${CLUSTER} -o jsonpath='$.ConnectionStatus') 
        if [ -z $STATUS ]; then    
            sleep 10
            seconds=$(($seconds + 10))
            echo "Cluster still offline, should come online in 600 seconds. Been trying now for $seconds seconds"
        else
            echo "Cluster $CLUSTER is now $STATUS"
            break
        fi
    done
fi