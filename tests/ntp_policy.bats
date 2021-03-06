TEST_NTP_POLICY_NAME=isctl-bats-test-ntp-policy-1

TEST_SECTION="NTP Policy CRUD"

@test "${TEST_SECTION}: create new NTP policy" {
    ./build/isctl create ntp policy --Name "${TEST_NTP_POLICY_NAME}" --NtpServers 1.1.1.1,2.2.2.2 --Organization default
}

# Intersight API now seems to allow this without error
# @test "${TEST_SECTION}: create duplicate NTP policy should fail" {
#     run ./build/isctl create ntp policy --Name "${TEST_NTP_POLICY_NAME}" --NtpServers 1.1.1.1,2.2.2.2 --Organization default
#     [ "$status" -ne 0 ]
# }

@test "${TEST_SECTION}: get NTP policy list should include new policy" {
    ./build/isctl get ntp policy | grep "${TEST_NTP_POLICY_NAME}"
}

@test "${TEST_SECTION}: get NTP policy by name and MOID" {
    MOID=$( ./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Moid )
    MOID2=$( ./build/isctl get ntp policy moid "${MOID}" -o json | jq -r .Moid )

    [ "${MOID}" == "${MOID2}" ]
}

@test "${TEST_SECTION}: get NTP policy by name and check enabled by default" {
    ENABLED=$( ./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Enabled )
    
    [ "${ENABLED}" == "true" ]
}

@test "${TEST_SECTION}: update NTP policy by moid" {
    # update the NTP policy with Enabled=False
    ./build/isctl update ntp policy moid $(./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" --jsonpath '$.Moid') --Enabled=False

    # Check Enabled=False
    ENABLED=$( ./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Enabled )

    [ "${ENABLED}" == "false" ]
}

@test "${TEST_SECTION}: update NTP policy by name" {
    # update the NTP policy with Enabled=True
    ./build/isctl update ntp policy name "${TEST_NTP_POLICY_NAME}" --Enabled=True

    # Check Enabled=False
    ENABLED=$( ./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Enabled )

    [ "${ENABLED}" == "true" ]
}

@test "${TEST_SECTION}: delete NTP policy" {
    ./build/isctl delete ntp policy moid $(./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" --jsonpath '$.Moid')

    # check that the policy no longer exists
    ! ./build/isctl get ntp policy | grep "${TEST_NTP_POLICY_NAME}"
}


setup_file() {
    # delete the test policy if it already exists. Don't check the exit code. 
    run ./build/isctl delete ntp policy moid $(./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" --jsonpath '$.Moid')
}