TEST_SECTION="IKS"

TEST_NAME=isctl-bats-test-1

@test "${TEST_SECTION}: invoking isctl iks outputs help" {
    run ./build/isctl iks
    [ "$status" -eq 0 ]
    [[ "${lines[0]}" == "Intersight Kubernetes Service commands"* ]]
}

@test "${TEST_SECTION}: invoking isctl iks cluster outputs help" {
    run ./build/isctl iks cluster
    [ "$status" -eq 0 ]
    [[ "${lines[0]}" == "Intersight Kubernetes Service cluster commands"* ]]
}

@test "${TEST_SECTION}: create IP Pool" {
    run ./build/isctl create ippool pool --Name "${TEST_NAME}" --IpV4Blocks '[{"From": "10.67.28.230", "To": "10.67.28.234"}]' --IpV4Config '{"Gateway": "10.67.28.129", "Netmask": "255.255.255.128", "PrimaryDns": "10.67.28.130"}' --Organization default
    [ "$status" -eq 0 ]
    GATEWAY=$(./build/isctl get ippool pool --name "${TEST_NAME}" -o json | jq -r .IpV4Config.Gateway)
    [ "${GATEWAY}" == "10.67.28.129" ]
}

@test "${TEST_SECTION}: create Network Policy" {
    run ./build/isctl create kubernetes networkpolicy --Name "${TEST_NAME}" --PodNetworkCidr 192.168.0.0/16 --ServiceCidr 10.13.1.0/24 --Organization default
    [ "$status" -eq 0 ]
    PODCIDR=$(./build/isctl get kubernetes networkpolicy --name "${TEST_NAME}" -o json | jq -r .PodNetworkCidr)
    [ "${PODCIDR}" == "192.168.0.0/16" ]
}

@test "${TEST_SECTION}: create Node OS Config Policy" {
    run ./build/isctl create kubernetes sysconfigpolicy --Name "${TEST_NAME}" --DnsDomainName cisco.com --DnsServers 8.8.8.8,8.8.4.4 --NtpServers ntp.esl.cisco.com --Timezone Australia/Sydney --Organization default
    [ "$status" -eq 0 ]
    NTP=$(./build/isctl get kubernetes sysconfigpolicy --name "${TEST_NAME}" -o json | jq -r ".NtpServers[0]")
    [ "${NTP}" == "ntp.esl.cisco.com" ]
}

@test "${TEST_SECTION}: create VM Infra Config Policy" {
    run ./build/isctl create kubernetes virtualmachineinfraconfigpolicy --Name "${TEST_NAME}" --VmConfig '{"ClassId": "kubernetes.EsxiVirtualMachineInfraConfig", "ObjectType": "kubernetes.EsxiVirtualMachineInfraConfig", "Cluster":"Melb-HX-Hybrid", "Datastore": "cgascoig-1", "ResourcePool": "", "Interfaces": ["vm-network-28"], "Passphrase": "C1sco123!!"}' --Organization default --Target "MoRef[DeviceHostname:10.67.17.125]"
    [ "$status" -eq 0 ]
    DS=$(./build/isctl get kubernetes virtualmachineinfraconfigpolicy --name "${TEST_NAME}" -o json | jq -r ".VmConfig.Datastore")
    [ "${DS}" == "cgascoig-1" ]
}

@test "${TEST_SECTION}: create VM Instance Type" {
    run ./build/isctl create kubernetes virtualmachineinstancetype --Name "${TEST_NAME}" --Cpu 2 --Memory 8192 --DiskSize 40 --Organization default
    [ "$status" -eq 0 ]
    DS=$(./build/isctl get kubernetes virtualmachineinstancetype --name "${TEST_NAME}" -o json | jq -r ".DiskSize")
    [ "${DS}" == "40" ]
}

@test "${TEST_SECTION}: create Kubernetes Version Policy" {
    run ./build/isctl create kubernetes versionpolicy --Name "${TEST_NAME}" --Version 1.19.5-iks-0 --Organization default
    [ "$status" -eq 0 ]
    V=$(./build/isctl get kubernetes versionpolicy --name "${TEST_NAME}" -o json | jq -r ".Version.ObjectType")
    [ "${V}" == "kubernetes.Version" ]
}

@test "${TEST_SECTION}: create Kubernetes Cluster Profile" {
    run ./build/isctl iks cluster create "${TEST_NAME}" --ip-pool "${TEST_NAME}" --network-policy "${TEST_NAME}" --sys-config-policy "${TEST_NAME}"  --vm-infra-config-policy "${TEST_NAME}"  --vm-instance-type-policy "${TEST_NAME}"  --version-policy "${TEST_NAME}"
    [ "$status" -eq 0 ]

    # check that the clusterprofile object exists and it has 2 nodegroupprofiles associated
    N=$(./build/isctl get kubernetes clusterprofile --name "${TEST_NAME}" -o json | jq -r ".NodeGroups | length")
    [ "${N}" == "2" ]
}

@test "${TEST_SECTION}: test cluster visibility commands" {

    # check that new cluster appears in cluster list
    ./build/isctl iks cluster list | grep "${TEST_NAME}"

    run ./build/isctl iks cluster status "${TEST_NAME}"
    [ "$status" -eq 0 ]
    [ "${lines[0]}" == "Cluster Name: ${TEST_NAME}" ]
    [ "${lines[1]}" == "Status: Undeployed" ]
}

@test "${TEST_SECTION}: test cluster deployment" {
    if [ "${BATS_TEST_IKS_DEPLOY}" != "yes" ]; then 
        skip "Skipping actual cluster deployment, set BATS_TEST_IKS_DEPLOY=yes to include this test"
    fi

    ./build/isctl iks cluster deploy "${TEST_NAME}" --wait-timeout 30
    run ./build/isctl iks cluster status "${TEST_NAME}"
    [ "$status" -eq 0 ]
    [ "${lines[0]}" == "Cluster Name: ${TEST_NAME}" ]
    [ "${lines[1]}" == "Status: Ready" ]
}

@test "${TEST_SECTION}: test deleting all resources" {
    sleep 10

    # delete any test objects that already exist. 
    ./build/isctl iks cluster undeploy "${TEST_NAME}"

    ./build/isctl delete kubernetes nodegroupprofile name "${TEST_NAME}-worker"
    ./build/isctl delete kubernetes nodegroupprofile name "${TEST_NAME}-ctrl"
    ./build/isctl delete kubernetes clusterprofile name "${TEST_NAME}"

    ./build/isctl delete kubernetes versionpolicy name "${TEST_NAME}"
    ./build/isctl delete ippool pool name "${TEST_NAME}"
    ./build/isctl delete kubernetes networkpolicy name "${TEST_NAME}"
    ./build/isctl delete kubernetes sysconfigpolicy name "${TEST_NAME}"
    ./build/isctl delete kubernetes virtualmachineinfraconfigpolicy name "${TEST_NAME}"
    ./build/isctl delete kubernetes virtualmachineinstancetype name "${TEST_NAME}"

    run ./build/isctl get kubernetes clusterprofile --name "${TEST_NAME}"
    [ "${#lines[@]}" -eq 0 ]
    run ./build/isctl get kubernetes nodegroupprofile --name "${TEST_NAME}-worker"
    [ "${#lines[@]}" -eq 0 ]
    run ./build/isctl get kubernetes nodegroupprofile --name "${TEST_NAME}-ctrl"
    [ "${#lines[@]}" -eq 0 ]

    run ./build/isctl get kubernetes versionpolicy --name "${TEST_NAME}"
    [ "${#lines[@]}" -eq 0 ]
    run ./build/isctl get ippool pool --name "${TEST_NAME}"
    [ "${#lines[@]}" -eq 0 ]
    run ./build/isctl get kubernetes networkpolicy --name "${TEST_NAME}"
    [ "${#lines[@]}" -eq 0 ]
    run ./build/isctl get kubernetes sysconfigpolicy --name "${TEST_NAME}"
    [ "${#lines[@]}" -eq 0 ]
    run ./build/isctl get kubernetes virtualmachineinfraconfigpolicy --name "${TEST_NAME}"
    [ "${#lines[@]}" -eq 0 ]
    run ./build/isctl get kubernetes virtualmachineinstancetype --name "${TEST_NAME}"
    [ "${#lines[@]}" -eq 0 ]
}


setup_file() {
    # delete any test objects that already exist. Don't check the exit code. 
    run ./build/isctl iks cluster undeploy "${TEST_NAME}"
    run ./build/isctl delete kubernetes nodegroupprofile name "${TEST_NAME}-worker"
    run ./build/isctl delete kubernetes nodegroupprofile name "${TEST_NAME}-ctrl"
    run ./build/isctl delete kubernetes clusterprofile name "${TEST_NAME}"

    run ./build/isctl delete kubernetes versionpolicy name "${TEST_NAME}"

    run ./build/isctl delete ippool pool name "${TEST_NAME}"
    run ./build/isctl delete kubernetes networkpolicy name "${TEST_NAME}"
    run ./build/isctl delete kubernetes sysconfigpolicy name "${TEST_NAME}"
    run ./build/isctl delete kubernetes virtualmachineinfraconfigpolicy name "${TEST_NAME}"
    run ./build/isctl delete kubernetes virtualmachineinstancetype name "${TEST_NAME}"
}