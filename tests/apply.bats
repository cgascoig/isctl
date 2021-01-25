
#this is the name used for created test objects. Must match the names in YAML files in tests/data/...
TEST_NAME=isctl-bats-test   

TEST_SECTION="Apply"

@test "${TEST_SECTION}: apply all-in-one YAML file" {
    ./build/isctl apply -f tests/data/apply-yaml-all-in-one/test.yaml
}

@test "${TEST_SECTION}: get NTP policy list should include new policy" {
    ./build/isctl get ntp policy | grep "${TEST_NAME}"
    ./build/isctl get organization organization | grep "${TEST_NAME}"
}

@test "${TEST_SECTION}: get NTP policy by name and check NTP servers correct" {
    NTPSERVERS=$( ./build/isctl get ntp policy --name "${TEST_NAME}" -o json | jq -r -c .NtpServers )
    
    [ "${NTPSERVERS}" == "[\"1.1.1.1\"]" ]
}

@test "${TEST_SECTION}: apply directoy of YAML files" {
    ./build/isctl apply -f tests/data/apply-yaml
}

@test "${TEST_SECTION}: check updated NTP servers correct" {
    NTPSERVERS=$( ./build/isctl get ntp policy --name "${TEST_NAME}" -o json | jq -r -c .NtpServers )
    
    [ "${NTPSERVERS}" == "[\"1.1.1.1\",\"2.2.2.2\"]" ]
}


@test "${TEST_SECTION}: delete NTP policy" {
    ./build/isctl delete ntp policy moid $(./build/isctl get ntp policy --name "${TEST_NAME}" --jsonpath '$.Moid')
    ./build/isctl delete organization organization moid $(./build/isctl get organization organization --name "${TEST_NAME}" --jsonpath '$.Moid')

    # check that the policy no longer exists
    ! ./build/isctl get ntp policy | grep "${TEST_NAME}"
    ! ./build/isctl get organization organization | grep "${TEST_NAME}"
}


setup_file() {
    # delete the test objects if they already exist. Don't check the exit code. 
    run ./build/isctl delete ntp policy moid $(./build/isctl get ntp policy --name "${TEST_NAME}" --jsonpath '$.Moid')
    run ./build/isctl delete organization organization moid $(./build/isctl get organization organization --name "${TEST_NAME}" --jsonpath '$.Moid')
}