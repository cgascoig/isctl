
TEST_POLICY_NAME=isctl-bats-test-direct-moid-policy
TEST_ORG_NAME=default

TEST_SECTION="Direct Moid MoRef"

setup_file() {
    # Get the default org Moid once for all tests
    ORG_MOID=$(./build/isctl ${ISCTL_OPTIONS} get organization organization --name "${TEST_ORG_NAME}" -o jsonpath='$.Moid')
    export ORG_MOID
}

@test "${TEST_SECTION}: create NTP policy with bare hex Moid for Organization" {
    ./build/isctl ${ISCTL_OPTIONS} create ntp policy \
        --Name "${TEST_POLICY_NAME}" \
        --NtpServers 1.1.1.1 \
        --Organization "${ORG_MOID}"
}

@test "${TEST_SECTION}: verify policy was created" {
    ./build/isctl ${ISCTL_OPTIONS} get ntp policy | grep "${TEST_POLICY_NAME}"
}

@test "${TEST_SECTION}: update policy using MoRef[<hex>] for Organization" {
    ./build/isctl ${ISCTL_OPTIONS} update ntp policy \
        moid "$(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_POLICY_NAME}" -o jsonpath='$.Moid')" \
        --Organization "MoRef[${ORG_MOID}]"
}

@test "${TEST_SECTION}: update policy using MoRef[Moid:<hex>] for Organization" {
    ./build/isctl ${ISCTL_OPTIONS} update ntp policy \
        moid "$(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_POLICY_NAME}" -o jsonpath='$.Moid')" \
        --Organization "MoRef[Moid:${ORG_MOID}]"
}

@test "${TEST_SECTION}: apply YAML with MoRef[Moid:<hex>] for Organization" {
    TMPFILE=$(mktemp /tmp/isctl-test-XXXXXX.yaml)
    sed "s/PLACEHOLDER_ORG_MOID/${ORG_MOID}/g" tests/data/test-direct-moid.yaml.template > "${TMPFILE}"
    ./build/isctl ${ISCTL_OPTIONS} apply -f "${TMPFILE}"
    rm -f "${TMPFILE}"
}

@test "${TEST_SECTION}: verify policy still exists after apply" {
    ./build/isctl ${ISCTL_OPTIONS} get ntp policy | grep "${TEST_POLICY_NAME}"
}

@test "${TEST_SECTION}: update policy using MoRef:organization.Organization[Moid:<hex>] for Organization" {
    ./build/isctl ${ISCTL_OPTIONS} update ntp policy \
        moid "$(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_POLICY_NAME}" -o jsonpath='$.Moid')" \
        --Organization "MoRef:organization.Organization[Moid:${ORG_MOID}]"
}

teardown_file() {
    POLICY_MOID=$(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_POLICY_NAME}" -o jsonpath='$.Moid' 2>/dev/null || true)
    if [ -n "${POLICY_MOID}" ]; then
        ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid "${POLICY_MOID}" || true
    fi
}
