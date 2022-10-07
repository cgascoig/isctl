TEST_NTP_POLICY_NAME=isctl-bats-test-policy

JSON_POLICY_NAME="${TEST_NTP_POLICY_NAME}-json"
JSON_NTP_POLICY="{\"Name\": \"${JSON_POLICY_NAME}\",\"Enabled\": true, \"NtpServers\": [\"1.1.1.1\"], \"Organization\": \"MoRef[Name:default]\"}"

YAML_POLICY_NAME="${TEST_NTP_POLICY_NAME}-yaml"
YAML_NTP_POLICY="ClassId: ntp.Policy\nObjectType: ntp.Policy\nName: ${YAML_POLICY_NAME}\nEnabled: true\nNtpServers: \n    - 1.1.1.1\nOrganization: MoRef[Name:default]"

TEST_SECTION="Passing body as YAML/JSON"

##################################################################
# JSON Tests
##################################################################

@test "${TEST_SECTION}: create new NTP policy - JSON" {
    echo -e "${JSON_NTP_POLICY}" | ./build/isctl ${ISCTL_OPTIONS} create ntp policy --bodyformat json
}

@test "${TEST_SECTION}: get NTP policy list should include new policy - JSON" {
    ./build/isctl ${ISCTL_OPTIONS} get ntp policy | grep "${JSON_POLICY_NAME}"
}

# @test "${TEST_SECTION}: update NTP policy - JSON" {
#     # update the NTP policy with Enabled=False
#     ./build/isctl ${ISCTL_OPTIONS} update ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o jsonpath='$.Moid') --Enabled=False

#     # Check Enabled=False
#     ENABLED=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Enabled )

#     [ "${ENABLED}" == "false" ]
# }

@test "${TEST_SECTION}: delete NTP policy - JSON" {
    ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${JSON_POLICY_NAME}" -o jsonpath='$.Moid')

    # check that the policy no longer exists
    ! ./build/isctl ${ISCTL_OPTIONS} get ntp policy | grep "${JSON_POLICY_NAME}"
}

##################################################################
# YAML Tests
##################################################################

@test "${TEST_SECTION}: create new NTP policy - YAML" {
    echo -e "${YAML_NTP_POLICY}" | ./build/isctl ${ISCTL_OPTIONS} create ntp policy --bodyformat yaml
}

@test "${TEST_SECTION}: get NTP policy list should include new policy - YAML" {
    ./build/isctl ${ISCTL_OPTIONS} get ntp policy | grep "${YAML_POLICY_NAME}"
}

# @test "${TEST_SECTION}: update NTP policy - YAML" {
#     # update the NTP policy with Enabled=False
#     ./build/isctl ${ISCTL_OPTIONS} update ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o jsonpath='$.Moid') --Enabled=False

#     # Check Enabled=False
#     ENABLED=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Enabled )

#     [ "${ENABLED}" == "false" ]
# }

@test "${TEST_SECTION}: delete NTP policy - YAML" {
    ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${YAML_POLICY_NAME}" -o jsonpath='$.Moid')

    # check that the policy no longer exists
    ! ./build/isctl ${ISCTL_OPTIONS} get ntp policy | grep "${YAML_POLICY_NAME}"
}


setup_file() {
    # delete the test policy if it already exists. Don't check the exit code. 
    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${JSON_POLICY_NAME}" -o jsonpath='$.Moid')
    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${YAML_POLICY_NAME}" -o jsonpath='$.Moid')
}