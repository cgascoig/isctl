TEST_NTP_POLICY_NAME=isctl-bats-test-ntp-policy-1

@test "create new NTP policy" {
    ./build/isctl create ntp policy --Name "${TEST_NTP_POLICY_NAME}" --NtpServers 1.1.1.1,2.2.2.2 --Organization default
}

# Intersight API now seems to allow this without error
# @test "create duplicate NTP policy should fail" {
#     run ./build/isctl create ntp policy --Name "${TEST_NTP_POLICY_NAME}" --NtpServers 1.1.1.1,2.2.2.2 --Organization default
#     [ "$status" -ne 0 ]
# }

@test "get NTP policy list should include new policy" {
    ./build/isctl get ntp policy | grep "${TEST_NTP_POLICY_NAME}"
}

@test "get NTP policy by name and MOID" {
    MOID=$( ./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Moid )
    MOID2=$( ./build/isctl get ntp policy moid "${MOID}" -o json | jq -r .Moid )

    [ "${MOID}" == "${MOID2}" ]
}

@test "get NTP policy by name and check enabled by default" {
    ENABLED=$( ./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Enabled )
    
    [ "${ENABLED}" == "true" ]
}

@test "update NTP policy" {
    # update the NTP policy with Enabled=False
    ./build/isctl update ntp policy moid $(./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" --jsonpath '$.Moid') --Enabled=False

    # Check Enabled=False
    ENABLED=$( ./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Enabled )

    [ "${ENABLED}" == "false" ]
}

@test "delete NTP policy" {
    ./build/isctl delete ntp policy moid $(./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" --jsonpath '$.Moid')

    # check that the policy no longer exists
    ! ./build/isctl get ntp policy | grep "${TEST_NTP_POLICY_NAME}"
}


setup_file() {
    # delete the test policy if it already exists. Don't check the exit code. 
    run ./build/isctl delete ntp policy moid $(./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" --jsonpath '$.Moid')
}