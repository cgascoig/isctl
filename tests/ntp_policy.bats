TEST_NTP_POLICY_NAME=isctl-bats-test-ntp-policy-1
TEST_NTP_POLICY_NAME_2=isctl-bats-test-ntp-policy-2
TEST_NTP_POLICY_NAME_BASE=isctl-bats-test-ntp-policy-

TEST_SECTION="NTP Policy CRUD"

@test "${TEST_SECTION}: create new NTP policy" {
    ./build/isctl ${ISCTL_OPTIONS} create ntp policy --Name "${TEST_NTP_POLICY_NAME}" --NtpServers 1.1.1.1,2.2.2.2 --Organization default --Tags '[{"Key":"tag1", "Value":"value1"},{"Key":"tag2", "Value":"value2"}]'
}

# Intersight API now seems to allow this without error
# @test "${TEST_SECTION}: create duplicate NTP policy should fail" {
#     run ./build/isctl ${ISCTL_OPTIONS} create ntp policy --Name "${TEST_NTP_POLICY_NAME}" --NtpServers 1.1.1.1,2.2.2.2 --Organization default
#     [ "$status" -ne 0 ]
# }

@test "${TEST_SECTION}: get NTP policy list should include new policy" {
    ./build/isctl ${ISCTL_OPTIONS} get ntp policy | grep "${TEST_NTP_POLICY_NAME}"
}

@test "${TEST_SECTION}: get NTP policy by name and MOID" {
    MOID=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Moid )
    MOID2=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy moid "${MOID}" -o json | jq -r .Moid )

    [ "${MOID}" == "${MOID2}" ]
}

@test "${TEST_SECTION}: get NTP policy by name and check enabled by default" {
    ENABLED=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Enabled )
    
    [ "${ENABLED}" == "true" ]
}

@test "${TEST_SECTION}: jsonpath" {
    run ./build/isctl ${ISCTL_OPTIONS} get ntp policy name "${TEST_NTP_POLICY_NAME}" --jsonpath '$.Name'
    assert_output "${TEST_NTP_POLICY_NAME}"

    run ./build/isctl ${ISCTL_OPTIONS} get ntp policy --filter "Name eq '${TEST_NTP_POLICY_NAME}'" --jsonpath '$[*].Name'
    assert_output "${TEST_NTP_POLICY_NAME}"
}

@test "${TEST_SECTION}: update NTP policy by moid" {
    # update the NTP policy with Enabled=False
    ./build/isctl ${ISCTL_OPTIONS} update ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NTP_POLICY_NAME}" --jsonpath '$.Moid') --Enabled=False

    # Check Enabled=False
    ENABLED=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Enabled )

    [ "${ENABLED}" == "false" ]
}

@test "${TEST_SECTION}: update NTP policy by name" {
    # update the NTP policy with Enabled=True
    ./build/isctl ${ISCTL_OPTIONS} update ntp policy name "${TEST_NTP_POLICY_NAME}" --Enabled=True

    # Check Enabled=False
    ENABLED=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Enabled )

    [ "${ENABLED}" == "true" ]
}

@test "${TEST_SECTION}: orderby, top and skip options" {
    # Create a second policy to test with
    ./build/isctl ${ISCTL_OPTIONS} create ntp policy --Name "${TEST_NTP_POLICY_NAME_2}" --NtpServers 1.1.1.1,2.2.2.2 --Organization default --Tags '[{"Key":"tag1", "Value":"value1"},{"Key":"tag2", "Value":"value2"}]'

    run ./build/isctl ${ISCTL_OPTIONS} get ntp policy --filter "startswith(Name, '${TEST_NTP_POLICY_NAME_BASE}')" -o table --select Name --orderby Name --top 1
    assert_success
    assert_line --index 3 --regexp "^ +${TEST_NTP_POLICY_NAME} +"

    run ./build/isctl ${ISCTL_OPTIONS} get ntp policy --filter "startswith(Name, '${TEST_NTP_POLICY_NAME_BASE}')" -o table --select Name --orderby Name --top 1 --skip 1
    assert_success
    assert_line --index 3 --regexp "^ +${TEST_NTP_POLICY_NAME_2} +"
}

@test "${TEST_SECTION}: filter NTP policy and select" {
    run ./build/isctl ${ISCTL_OPTIONS} get ntp policy --filter "Name eq '${TEST_NTP_POLICY_NAME}'" -o table --select Name,Enabled
    assert_line --index 1 --regexp '^ +Name +Moid +Enabled *$'
    assert_line --index 3 --regexp "^ +${TEST_NTP_POLICY_NAME} +[0-9a-f]{24} +True *$"
}

@test "${TEST_SECTION}: expand organisation name" {
  ORG=$(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json --expand 'Organization($select=Name)' | jq -r .Organization.Name)
  assert_equal "${ORG}" "default"
}

@test "${TEST_SECTION}: custom-columns" {
    run ./build/isctl ${ISCTL_OPTIONS} get ntp policy --filter "Name eq '${TEST_NTP_POLICY_NAME}'" -o custom-columns='NAME:.Name,ENABLED:.Enabled'
    assert_success
    assert_line --index 1 --regexp '^ +NAME +ENABLED *$'
    assert_line --index 3 --regexp "^ +${TEST_NTP_POLICY_NAME} +True *$"

    run ./build/isctl ${ISCTL_OPTIONS} get ntp policy --filter "Name eq '${TEST_NTP_POLICY_NAME}'" -o custom-columns='NAME:.Name,ENABLED:.Enabled'
    assert_success
    assert_line --index 1 --regexp '^ +NAME +ENABLED *$'
    assert_line --index 3 --regexp "^ +${TEST_NTP_POLICY_NAME} +True *$"

    run ./build/isctl ${ISCTL_OPTIONS} get ntp policy --filter "Name eq '${TEST_NTP_POLICY_NAME}'" -o custom-columns='NAME:.Name,ENABLED:.Enabled'
    assert_success
    assert_line --index 1 --regexp '^ +NAME +ENABLED *$'
    assert_line --index 3 --regexp "^ +${TEST_NTP_POLICY_NAME} +True *$"
}

@test "${TEST_SECTION}: tag output formatted correctly" {
    run ./build/isctl ${ISCTL_OPTIONS} get ntp policy --filter "Name eq '${TEST_NTP_POLICY_NAME}'" -o custom-columns='NAME:.Name,TAGS:.Tags'
    assert_success
    assert_line --index 1 --regexp '^ +NAME +TAGS *$'
    assert_line --index 3 --regexp "^ +${TEST_NTP_POLICY_NAME} +tag1: value1, tag2: value2 *$"
}

@test "${TEST_SECTION}: delete NTP policy" {
    ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NTP_POLICY_NAME}" --jsonpath '$.Moid')

    # check that the policy no longer exists
    ! ./build/isctl ${ISCTL_OPTIONS} get ntp policy | grep "${TEST_NTP_POLICY_NAME}"
}

@test "${TEST_SECTION}: test delete NTP policy by name" {
    DELBYNAME_POLICY_NAME="${TEST_NTP_POLICY_NAME}-delbyname"

    # create the policy
    ./build/isctl ${ISCTL_OPTIONS} create ntp policy --Name "${DELBYNAME_POLICY_NAME}" --NtpServers 1.1.1.1,2.2.2.2 --Organization default

    # check that the policy exists
    ./build/isctl ${ISCTL_OPTIONS} get ntp policy | grep "${DELBYNAME_POLICY_NAME}"

    # delete by name
    ./build/isctl ${ISCTL_OPTIONS} delete ntp policy name "${DELBYNAME_POLICY_NAME}"

    # check that the policy no longer exists
    ! ./build/isctl ${ISCTL_OPTIONS} get ntp policy | grep "${DELBYNAME_POLICY_NAME}"
}


setup_file() {
    # delete the test policy if it already exists. Don't check the exit code. 
    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NTP_POLICY_NAME}" --jsonpath '$.Moid')
    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NTP_POLICY_NAME_2}" --jsonpath '$.Moid')
}

setup() {
    load 'test_helper/bats-support/load' # this is required by bats-assert!
    load 'test_helper/bats-assert/load'
}