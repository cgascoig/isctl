# Intersight Kubernetes Service

## Deploying a cluster

### Create the Kubernetes policies

First, create the required policies for your cluster(s). Ideally you will create these policies once and then reuse them across multiple clusters. You could also create the policies through the Intersight GUI since it should be a one-time operation. 

Create the IP Pool:
```
isctl create ippool pool --Name cg-isctl-test-1 --IpV4Blocks '[{"From": "10.67.28.230", "To": "10.67.28.234"}]' --IpV4Config '{"Gateway": "10.67.28.129", "Netmask": "255.255.255.128", "PrimaryDns": "10.67.28.130"}' --Organization default
```

Create the Kubernetes Network CIDR policy:
```
isctl create kubernetes networkpolicy --PodNetworkCidr 192.168.0.0/16 --ServiceCidr 10.13.1.0/24 --Name cg-isctl-test-1 --Organization default
```

Create the Kubernetes Node OS Configuration policy:
```
isctl create kubernetes sysconfigpolicy --DnsDomainName cisco.com --DnsServers 8.8.8.8,8.8.4.4 --NtpServers pool.ntp.org --Timezone Australia/Sydney --Name cg-isctl-test-1 --Organization default
```

Create the Kubernetes VM Infra Config policy:
```
isctl create kubernetes virtualmachineinfraconfigpolicy --Name cg-isctl-test-1 --VmConfig '{"ClassId": "kubernetes.EsxiVirtualMachineInfraConfig", "ObjectType": "kubernetes.EsxiVirtualMachineInfraConfig", "Cluster":"Melb-HX-Hybrid", "Datastore": "cgascoig-1", "ResourcePool": "", "Interfaces": ["vm-network-28"], "Passphrase": "C1sco123!!"}' --Organization default --Target "MoRef[DeviceHostname:10.67.17.125]"
```

Create the Kubernetes VM Instance Type policy:
```
isctl create kubernetes virtualmachineinstancetype --Name cg-isctl-test-1 --Cpu 2 --Memory 4096 --DiskSize 40 --Organization default
```

Create the Kubernetes Version policy:
```
isctl create kubernetes versionpolicy --Version 1.19.5-iks-0 --Name cg-isctl-test-1 --Organization default
```

### Deploy a new Kubernetes cluster

First, create a new Kubernetes Cluster Profile that references the previously created policies:
```
isctl iks cluster create cg-isctl-test-1 --ip-pool cg-isctl-test-1 --network-policy cg-isctl-test-1 --sys-config-policy cg-isctl-test-1  --vm-infra-config-policy cg-isctl-test-1  --vm-instance-type-policy cg-isctl-test-1  --version-policy cg-isctl-test-1
```

Then, start the cluster deployment:
```
isctl iks cluster deploy cg-isctl-test-1
```

By default, `isctl` will wait for deployment to complete and report success/failure, however, you can use `--no-wait` to just start the deployment and return. 

## Working with clusters

To list all existing IKS clusters within Intersight: 
```
isctl iks cluster list
```

To check the status of an existing cluster:
```
isctl iks cluster status <CLUSTER NAME>
```
> This will show information about the cluster such as deployment status, node group and node details:
```
Cluster Name: cg-iks-prod
Status: Ready
Node Groups:
  Node Group Name: cg-iks-prod-control
  Type: ControlPlane
  Desired Size: 1
  Nodes:
    Node Name: cg-iks-prod-control-4e97693a32
    Cloud Provider: external
    IP Address: 10.67.28.237
    -----------------
  -----------------
  Node Group Name: cg-iks-prod-worker_profile
  Type: Worker
  Desired Size: 1
  Nodes:
    Node Name: cg-iks-prod-worker-pr-79b95a983b
    Cloud Provider: external
    IP Address: 10.67.28.238
    -----------------
  -----------------
```

To get the credentials (kubeconfig file) for a deployed cluster:
```
isctl iks cluster get-credentials <CLUSTER NAME>
```

By default, the kubeconfig details will be printed to stdout, but you can save them into a file using `-f`. Therefore a typical way to get access to a deployed cluster is:
```
isctl iks cluster get-credentials cg-iks-prod -f ~/Downloads/cg-iks-prod.yaml
export KUBECONFIG=~/Downloads/cg-iks-prod.yaml
kubectl get nodes
...
```

## Destroying clusters

First, undeploy the cluster - `isctl` will wait for the process to complete before returning:
```
isctl iks cluster undeploy <CLUSTER NAME>
```

Next, delete the Node Groups:
```
isctl delete kubernetes nodegroupprofile name <CLUSTER NAME>-worker
isctl delete kubernetes nodegroupprofile name <CLUSTER NAME>-ctrl
```

Finally, delete the Cluster Profile:
```
isctl delete kubernetes clusterprofile name <CLUSTER NAME>
```

# Other tasks

## Script to repair IKS cluster and vCenter target linkage

[This](https://github.com/cgascoig/isctl/blob/devel/examples/scripts/iks-fix-vcenter-target.sh) example script will re-establish the linkage between a deployed IKS cluster and a vCenter target. This is useful in scenarios where the vCenter target was removed and re-added while an IKS cluster was deployed. 