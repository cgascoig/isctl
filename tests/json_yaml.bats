TEST_NTP_POLICY_NAME=isctl-bats-test-policy

JSON_POLICY_NAME="${TEST_NTP_POLICY_NAME}-json"
JSON_NTP_POLICY="{\"Name\": \"${JSON_POLICY_NAME}\",\"Enabled\": true, \"NtpServers\": [\"1.1.1.1\"], \"Organization\": \"MoRef[Name:default]\"}"

YAML_POLICY_NAME="${TEST_NTP_POLICY_NAME}-yaml"
YAML_NTP_POLICY="ClassId: ntp.Policy\nObjectType: ntp.Policy\nName: ${YAML_POLICY_NAME}\nEnabled: true\nNtpServers: \n    - 1.1.1.1\nOrganization: MoRef[Name:default]"

##################################################################
# JSON Tests
##################################################################

@test "create new NTP policy - JSON" {
    echo -e "${JSON_NTP_POLICY}" | ./build/isctl create ntp policy --bodyformat json
}

@test "get NTP policy list should include new policy - JSON" {
    ./build/isctl get ntp policy | grep "${JSON_POLICY_NAME}"
}

# @test "update NTP policy - JSON" {
#     # update the NTP policy with Enabled=False
#     ./build/isctl update ntp policy moid $(./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" --jsonpath '$.Moid') --Enabled=False

#     # Check Enabled=False
#     ENABLED=$( ./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Enabled )

#     [ "${ENABLED}" == "false" ]
# }

@test "delete NTP policy - JSON" {
    ./build/isctl delete ntp policy moid $(./build/isctl get ntp policy --name "${JSON_POLICY_NAME}" --jsonpath '$.Moid')

    # check that the policy no longer exists
    ! ./build/isctl get ntp policy | grep "${JSON_POLICY_NAME}"
}

##################################################################
# YAML Tests
##################################################################

@test "create new NTP policy - YAML" {
    echo -e "${YAML_NTP_POLICY}" | ./build/isctl create ntp policy --bodyformat yaml
}

@test "get NTP policy list should include new policy - YAML" {
    ./build/isctl get ntp policy | grep "${YAML_POLICY_NAME}"
}

# @test "update NTP policy - YAML" {
#     # update the NTP policy with Enabled=False
#     ./build/isctl update ntp policy moid $(./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" --jsonpath '$.Moid') --Enabled=False

#     # Check Enabled=False
#     ENABLED=$( ./build/isctl get ntp policy --name "${TEST_NTP_POLICY_NAME}" -o json | jq -r .Enabled )

#     [ "${ENABLED}" == "false" ]
# }

@test "delete NTP policy - YAML" {
    ./build/isctl delete ntp policy moid $(./build/isctl get ntp policy --name "${YAML_POLICY_NAME}" --jsonpath '$.Moid')

    # check that the policy no longer exists
    ! ./build/isctl get ntp policy | grep "${YAML_POLICY_NAME}"
}


setup_file() {
    # delete the test policy if it already exists. Don't check the exit code. 
    run ./build/isctl delete ntp policy moid $(./build/isctl get ntp policy --name "${JSON_POLICY_NAME}" --jsonpath '$.Moid')
    run ./build/isctl delete ntp policy moid $(./build/isctl get ntp policy --name "${YAML_POLICY_NAME}" --jsonpath '$.Moid')
}