TEST_SECTION="IKS"

TEST_NAME=isctl-bats-test-1

@test "${TEST_SECTION}: invoking isctl iks outputs help" {
    run ./build/isctl ${ISCTL_OPTIONS} iks
    assert_success
    assert_output --regexp "^Intersight Kubernetes Service commands"
}

@test "${TEST_SECTION}: invoking isctl iks cluster outputs help" {
    run ./build/isctl ${ISCTL_OPTIONS} iks cluster
    assert_success
    assert_output --regexp "^Intersight Kubernetes Service cluster commands"
}

@test "${TEST_SECTION}: get kubernetes version" {
    run ./build/isctl ${ISCTL_OPTIONS} get kubernetes version
    assert_success
}

@test "${TEST_SECTION}: create IP Pool" {
    run ./build/isctl ${ISCTL_OPTIONS} create ippool pool --Name "${TEST_NAME}" --IpV4Blocks '[{"From": "10.67.28.230", "To": "10.67.28.234"}]' --IpV4Config '{"Gateway": "10.67.28.129", "Netmask": "255.255.255.128", "PrimaryDns": "10.67.28.130"}' --Organization default
    assert_success
    GATEWAY=$(./build/isctl ${ISCTL_OPTIONS} get ippool pool --name "${TEST_NAME}" -o json | jq -r .IpV4Config.Gateway)
    assert_equal "${GATEWAY}" "10.67.28.129"
}

@test "${TEST_SECTION}: create Network Policy" {
    run ./build/isctl ${ISCTL_OPTIONS} create kubernetes networkpolicy --Name "${TEST_NAME}" --PodNetworkCidr 192.168.0.0/16 --ServiceCidr 10.13.1.0/24 --Organization default
    assert_success
    PODCIDR=$(./build/isctl ${ISCTL_OPTIONS} get kubernetes networkpolicy --name "${TEST_NAME}" -o json | jq -r .PodNetworkCidr)
    assert_equal "${PODCIDR}" "192.168.0.0/16"
}

@test "${TEST_SECTION}: create Node OS Config Policy" {
    run ./build/isctl ${ISCTL_OPTIONS} create kubernetes sysconfigpolicy --Name "${TEST_NAME}" --DnsDomainName cisco.com --DnsServers 8.8.8.8,8.8.4.4 --NtpServers ntp.esl.cisco.com --Timezone Australia/Sydney --Organization default
    assert_success
    NTP=$(./build/isctl ${ISCTL_OPTIONS} get kubernetes sysconfigpolicy --name "${TEST_NAME}" -o json | jq -r ".NtpServers[0]")
    assert_equal "${NTP}" "ntp.esl.cisco.com"
}

@test "${TEST_SECTION}: create VM Infra Config Policy" {
    run ./build/isctl ${ISCTL_OPTIONS} create kubernetes virtualmachineinfraconfigpolicy --Name "${TEST_NAME}" --VmConfig '{"ClassId": "kubernetes.EsxiVirtualMachineInfraConfig", "ObjectType": "kubernetes.EsxiVirtualMachineInfraConfig", "Cluster":"Melb-HX-Hybrid", "Datastore": "cgascoig-1", "ResourcePool": "", "Interfaces": ["vm-network-28"], "Passphrase": "C1sco123!!"}' --Organization default --Target "MoRef[DeviceHostname:10.67.17.125]"
    assert_success
    DS=$(./build/isctl ${ISCTL_OPTIONS} get kubernetes virtualmachineinfraconfigpolicy --name "${TEST_NAME}" -o json | jq -r ".VmConfig.Datastore")
    assert_equal "${DS}" "cgascoig-1"
}

@test "${TEST_SECTION}: create VM Instance Type" {
    run ./build/isctl ${ISCTL_OPTIONS} create kubernetes virtualmachineinstancetype --Name "${TEST_NAME}" --Cpu 2 --Memory 8192 --DiskSize 40 --Organization default
    assert_success
    DS=$(./build/isctl ${ISCTL_OPTIONS} get kubernetes virtualmachineinstancetype --name "${TEST_NAME}" -o json | jq -r ".DiskSize")
    assert_equal "${DS}" "40"
}

@test "${TEST_SECTION}: create Kubernetes Version Policy" {
    run ./build/isctl ${ISCTL_OPTIONS} create kubernetes versionpolicy --Name "${TEST_NAME}" --Version 1.19.5-iks-0 --Organization default
    assert_success
    V=$(./build/isctl ${ISCTL_OPTIONS} get kubernetes versionpolicy --name "${TEST_NAME}" -o json | jq -r ".Version.ObjectType")
    assert_equal "${V}" "kubernetes.Version"
}

@test "${TEST_SECTION}: create Kubernetes Cluster Profile" {
    run ./build/isctl ${ISCTL_OPTIONS} iks cluster create "${TEST_NAME}" --ip-pool "${TEST_NAME}" --network-policy "${TEST_NAME}" --sys-config-policy "${TEST_NAME}"  --vm-infra-config-policy "${TEST_NAME}"  --vm-instance-type-policy "${TEST_NAME}"  --version-policy "${TEST_NAME}"
    assert_success

    # check that the clusterprofile object exists and it has 2 nodegroupprofiles associated
    N=$(./build/isctl ${ISCTL_OPTIONS} get kubernetes clusterprofile --name "${TEST_NAME}" -o json | jq -r ".NodeGroups | length")
    assert_equal "${N}" "2"
}

@test "${TEST_SECTION}: test cluster visibility commands" {

    # check that new cluster appears in cluster list
    ./build/isctl ${ISCTL_OPTIONS} iks cluster list | grep "${TEST_NAME}"

    run ./build/isctl ${ISCTL_OPTIONS} iks cluster status "${TEST_NAME}"
    assert_success
    assert_line --index 0 "Cluster Name: ${TEST_NAME}"
    assert_line --index 1 "Status: Undeployed"
}

@test "${TEST_SECTION}: test cluster deployment" {
    if [ "${BATS_TEST_IKS_DEPLOY}" != "yes" ]; then 
        skip "Skipping actual cluster deployment, set BATS_TEST_IKS_DEPLOY=yes to include this test"
    fi

    ./build/isctl ${ISCTL_OPTIONS} iks cluster deploy "${TEST_NAME}" --wait-timeout 30
    run ./build/isctl ${ISCTL_OPTIONS} iks cluster status "${TEST_NAME}"
    assert_success
    assert_line --index 0 "Cluster Name: ${TEST_NAME}"
    assert_line --index 1 "Status: Ready"
}

@test "${TEST_SECTION}: test deleting all resources" {
    sleep 10

    # delete any test objects that already exist. 
    ./build/isctl ${ISCTL_OPTIONS} iks cluster undeploy "${TEST_NAME}"

    ./build/isctl ${ISCTL_OPTIONS} delete kubernetes nodegroupprofile name "${TEST_NAME}-worker"
    ./build/isctl ${ISCTL_OPTIONS} delete kubernetes nodegroupprofile name "${TEST_NAME}-ctrl"
    ./build/isctl ${ISCTL_OPTIONS} delete kubernetes clusterprofile name "${TEST_NAME}"

    ./build/isctl ${ISCTL_OPTIONS} delete kubernetes versionpolicy name "${TEST_NAME}"
    ./build/isctl ${ISCTL_OPTIONS} delete ippool pool name "${TEST_NAME}"
    ./build/isctl ${ISCTL_OPTIONS} delete kubernetes networkpolicy name "${TEST_NAME}"
    ./build/isctl ${ISCTL_OPTIONS} delete kubernetes sysconfigpolicy name "${TEST_NAME}"
    ./build/isctl ${ISCTL_OPTIONS} delete kubernetes virtualmachineinfraconfigpolicy name "${TEST_NAME}"
    ./build/isctl ${ISCTL_OPTIONS} delete kubernetes virtualmachineinstancetype name "${TEST_NAME}"

    for i in $(seq 1 20)
    do
        run ./build/isctl ${ISCTL_OPTIONS} get kubernetes clusterprofile --name "${TEST_NAME}"
        if [ "${#lines[@]}" -eq 0 ]
        then
            break
        fi
        sleep 30
    done
    assert_output ""

    run ./build/isctl ${ISCTL_OPTIONS} get kubernetes nodegroupprofile --name "${TEST_NAME}-worker"
    assert_output ""
    run ./build/isctl ${ISCTL_OPTIONS} get kubernetes nodegroupprofile --name "${TEST_NAME}-ctrl"
    assert_output ""

    run ./build/isctl ${ISCTL_OPTIONS} get kubernetes versionpolicy --name "${TEST_NAME}"
    assert_output ""
    run ./build/isctl ${ISCTL_OPTIONS} get ippool pool --name "${TEST_NAME}"
    assert_output ""
    run ./build/isctl ${ISCTL_OPTIONS} get kubernetes networkpolicy --name "${TEST_NAME}"
    assert_output ""
    run ./build/isctl ${ISCTL_OPTIONS} get kubernetes sysconfigpolicy --name "${TEST_NAME}"
    assert_output ""
    run ./build/isctl ${ISCTL_OPTIONS} get kubernetes virtualmachineinfraconfigpolicy --name "${TEST_NAME}"
    assert_output ""
    run ./build/isctl ${ISCTL_OPTIONS} get kubernetes virtualmachineinstancetype --name "${TEST_NAME}"
    assert_output ""
}


setup_file() {
    # delete any test objects that already exist. Don't check the exit code. 
    run ./build/isctl ${ISCTL_OPTIONS} iks cluster undeploy "${TEST_NAME}"
    run ./build/isctl ${ISCTL_OPTIONS} delete kubernetes nodegroupprofile name "${TEST_NAME}-worker"
    run ./build/isctl ${ISCTL_OPTIONS} delete kubernetes nodegroupprofile name "${TEST_NAME}-ctrl"
    run ./build/isctl ${ISCTL_OPTIONS} delete kubernetes clusterprofile name "${TEST_NAME}"

    run ./build/isctl ${ISCTL_OPTIONS} delete kubernetes versionpolicy name "${TEST_NAME}"

    run ./build/isctl ${ISCTL_OPTIONS} delete ippool pool name "${TEST_NAME}"
    run ./build/isctl ${ISCTL_OPTIONS} delete kubernetes networkpolicy name "${TEST_NAME}"
    run ./build/isctl ${ISCTL_OPTIONS} delete kubernetes sysconfigpolicy name "${TEST_NAME}"
    run ./build/isctl ${ISCTL_OPTIONS} delete kubernetes virtualmachineinfraconfigpolicy name "${TEST_NAME}"
    run ./build/isctl ${ISCTL_OPTIONS} delete kubernetes virtualmachineinstancetype name "${TEST_NAME}"
}

setup() {
    load 'test_helper/bats-support/load' # this is required by bats-assert!
    load 'test_helper/bats-assert/load'
}