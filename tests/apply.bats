
#this is the name used for created test objects. Must match the names in YAML files in tests/data/...
TEST_NAME=isctl-bats-test   

TEST_SECTION="Apply"

@test "${TEST_SECTION}: apply all-in-one YAML file" {
    ./build/isctl ${ISCTL_OPTIONS} apply -f tests/data/apply-yaml-all-in-one/test.yaml
}

@test "${TEST_SECTION}: get NTP policy list should include new policy" {
    ./build/isctl ${ISCTL_OPTIONS} get ntp policy | grep "${TEST_NAME}"
    ./build/isctl ${ISCTL_OPTIONS} get organization organization | grep "${TEST_NAME}"
}

@test "${TEST_SECTION}: get NTP policy by name and check NTP servers correct" {
    NTPSERVERS=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NAME}" -o json | jq -r -c .NtpServers )
    
    [ "${NTPSERVERS}" == "[\"1.1.1.1\"]" ]
}

@test "${TEST_SECTION}: apply directoy of YAML files" {
    ./build/isctl ${ISCTL_OPTIONS} apply -f tests/data/apply-yaml
}

@test "${TEST_SECTION}: check updated NTP servers correct" {
    NTPSERVERS=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NAME}" -o json | jq -r -c .NtpServers )
    
    [ "${NTPSERVERS}" == "[\"1.1.1.1\",\"2.2.2.2\"]" ]
}

@test "${TEST_SECTION}: apply YAML file with duplicated NTP Policy name in different org" {
    # Apply YAML with same NTP policy name but different org (also creates the org)
    ./build/isctl ${ISCTL_OPTIONS} apply -f tests/data/test-duplicate-name-cross-orgs.yaml

    # Get the Moid of new org
    ORG_MOID=$(./build/isctl ${ISCTL_OPTIONS} get organization organization name isctl-bats-dup-test -o jsonpath=.Moid)
    [ ! -z "${ORG_MOID}" ]

    # Check if new policy in org was created successfully
    NTP_MOID=$(./build/isctl ${ISCTL_OPTIONS} get ntp policy --filter "Name eq '${TEST_NAME}' and Organization/Moid eq '${ORG_MOID}'" -o 'jsonpath=$[*].Moid')
    [ ! -z "${NTP_MOID}" ]
    # ./build/isctl ${ISCTL_OPTIONS} get ntp policy --filter "Name eq '${TEST_NAME}' and Organization/Moid eq '${ORG_MOID}'" | grep "${TEST_NAME}"

    # ./build/isctl 

    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid ${NTP_MOID}
    run ./build/isctl ${ISCTL_OPTIONS} delete organization organization moid ${ORG_MOID}
}


@test "${TEST_SECTION}: delete NTP policy" {
    ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NAME}" -o jsonpath='$.Moid')
    sleep 5
    ./build/isctl ${ISCTL_OPTIONS} delete organization organization moid $(./build/isctl ${ISCTL_OPTIONS} get organization organization --name "${TEST_NAME}" -o jsonpath='$.Moid')

    # check that the policy no longer exists
    ! ./build/isctl ${ISCTL_OPTIONS} get ntp policy | grep "${TEST_NAME}"
    ! ./build/isctl ${ISCTL_OPTIONS} get organization organization | grep "${TEST_NAME}"
}

@test "${TEST_SECTION}: bulk.MoCloners" {
    ./build/isctl ${ISCTL_OPTIONS} create server profiletemplate --Name "${TEST_NAME}" --Organization default

    ./build/isctl ${ISCTL_OPTIONS} apply -f tests/data/test-bulk-mo-cloner.yaml

    run ./build/isctl ${ISCTL_OPTIONS} delete server profile name "${TEST_NAME}-1"
    run ./build/isctl ${ISCTL_OPTIONS} delete server profiletemplate name "${TEST_NAME}" 
}


setup_file() {
    # delete the test objects if they already exist. Don't check the exit code. 
    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NAME}" -o jsonpath='$.Moid')
    run ./build/isctl ${ISCTL_OPTIONS} delete organization organization moid $(./build/isctl ${ISCTL_OPTIONS} get organization organization --name "${TEST_NAME}" -o jsonpath='$.Moid')
}