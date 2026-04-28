
#this is the name used for created test objects. Must match the names in YAML files in tests/data/...
TEST_NAME=isctl-org-test   

TEST_SECTION="Organization"

@test "${TEST_SECTION}: manually create duplicate endpointuser policies" {
    echo "Creating test IAM endpointuser in default org"
    ./build/isctl ${ISCTL_OPTIONS} create iam endpointuser --Name "${TEST_NAME}" --Organization "default"
    echo "Creating test IAM endpointuser in test org"
    ./build/isctl ${ISCTL_OPTIONS} create iam endpointuser --Name "${TEST_NAME}" --Organization "${TEST_NAME}"
}

@test "${TEST_SECTION}: manually reference an mo in a specific org" {

    echo "Creating test IAM endpointuserpolicy in test org"
    ./build/isctl ${ISCTL_OPTIONS} create iam endpointuserpolicy --Name "${TEST_NAME}" --Organization "${TEST_NAME}"

    echo "Creating test IAM endpointuserrole in test org (references endpointuser and endpointuserpolicy in specific org)"
    ./build/isctl ${ISCTL_OPTIONS} create iam endpointuserrole \
        --EndPointRole "[\"MoRef[\$filter:Name eq 'admin' and Type eq 'IMC']\"]" \
        --EndPointUser "MoRef[${TEST_NAME}\\${TEST_NAME}]" \
        --EndPointUserPolicy "MoRef[${TEST_NAME}\\${TEST_NAME}]" \
        --Password "P@ssw0rd!!!"

    echo "Deleting test IAM endpointuserrole in test org"
    ORG_MOID=$(./build/isctl ${ISCTL_OPTIONS} get organization organization --name "${TEST_NAME}" -o jsonpath='$.Moid'|| echo "")

    # Don't need to delete endpointuserrole - it is automatically deleted when the policy is deleted
    # ./build/isctl ${ISCTL_OPTIONS} delete iam endpointuserrole moid $(./build/isctl ${ISCTL_OPTIONS} get iam endpointuserrole --filter "Name eq '${TEST_NAME}' and Organization/Moid eq '${ORG_MOID}'" -o 'jsonpath=$[*].Moid'|| echo "")

    echo "Deleting test IAM endpointuserpolicy in test org"
    ./build/isctl ${ISCTL_OPTIONS} delete iam endpointuserpolicy moid $(./build/isctl ${ISCTL_OPTIONS} get iam endpointuserpolicy --filter "Name eq '${TEST_NAME}' and Organization/Moid eq '${ORG_MOID}'" -o 'jsonpath=$[*].Moid'|| echo "")
}

# This test is disabled as it the API no longer allows NTP policies with the same name in different orgs
# @test "${TEST_SECTION}: manually create duplicate ntp policies" {
#     echo "Creating test NTP policy in default org"
#     ./build/isctl ${ISCTL_OPTIONS} create ntp policy --Name "${TEST_NAME}" --Organization "default" --NtpServers 1.1.1.1

#     echo "Creating test NTP policy in test org"
#     ./build/isctl ${ISCTL_OPTIONS} create ntp policy --Name "${TEST_NAME}" --Organization "${TEST_NAME}" --NtpServers 1.1.1.1
# }

@test "${TEST_SECTION}: apply YAML file with duplicated endpointuser name in different org" {
    # delete the endpointuser in test org if it exists
    ORG_MOID=$(./build/isctl ${ISCTL_OPTIONS} get organization organization --name "${TEST_NAME}" -o jsonpath='$.Moid'|| echo "")
    run ./build/isctl ${ISCTL_OPTIONS} delete iam endpointuser moid $(./build/isctl ${ISCTL_OPTIONS} get iam endpointuser --filter "Name eq '${TEST_NAME}' and Organization/Moid eq '${ORG_MOID}'" -o 'jsonpath=$[*].Moid'|| echo "")
    
    # Apply YAML with duplicate endpointuser in test org
    ./build/isctl ${ISCTL_OPTIONS} apply -f tests/data/test-duplicate-name-cross-orgs.yaml

    # Check if new endpointuser in test org was created successfully
    EPU_MOID=$(./build/isctl ${ISCTL_OPTIONS} get iam endpointuser --filter "Name eq '${TEST_NAME}' and Organization/Moid eq '${ORG_MOID}'" -o 'jsonpath=$[*].Moid')
    [ ! -z "${EPU_MOID}" ]
}

setup() {
    load 'test_helper/bats-support/load' # this is required by bats-assert!
    load 'test_helper/bats-assert/load'
}

setup_file() {
    # # delete the test objects if they already exist. Don't check the exit code. 
    teardown_file

    echo "Creating test organization"
    ./build/isctl ${ISCTL_OPTIONS} create organization organization --Name "${TEST_NAME}"
}

teardown_file() {
    echo "Cleaning up objects"
    # Delete the objects in default org
    echo "Deleting objects in default org.."
    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NAME}" -o jsonpath='$.Moid'|| echo "")
    run ./build/isctl ${ISCTL_OPTIONS} delete iam endpointuser moid $(./build/isctl ${ISCTL_OPTIONS} get iam endpointuser --name "${TEST_NAME}" -o jsonpath='$.Moid'|| echo "")

    # Delete the objects in new org
    echo "Deleting objects in test org.."
    ORG_MOID=$(./build/isctl ${ISCTL_OPTIONS} get organization organization --name "${TEST_NAME}" -o jsonpath='$.Moid'|| echo "")
    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --filter "Name eq '${TEST_NAME}' and Organization/Moid eq '${ORG_MOID}'" -o 'jsonpath=$[*].Moid'|| echo "")
    run ./build/isctl ${ISCTL_OPTIONS} delete iam endpointuser moid $(./build/isctl ${ISCTL_OPTIONS} get iam endpointuser --filter "Name eq '${TEST_NAME}' and Organization/Moid eq '${ORG_MOID}'" -o 'jsonpath=$[*].Moid'|| echo "")
    run ./build/isctl ${ISCTL_OPTIONS} delete iam endpointuserpolicy moid $(./build/isctl ${ISCTL_OPTIONS} get iam endpointuserpolicy --filter "Name eq '${TEST_NAME}' and Organization/Moid eq '${ORG_MOID}'" -o 'jsonpath=$[*].Moid'|| echo "")

    # Delete org
    if ! [ -z "${ORG_MOID}" ] 
    then
        echo "Deleting org.."
        ./build/isctl ${ISCTL_OPTIONS} delete organization organization moid ${ORG_MOID}
    fi
}