TEST_SECTION="Contract Status"

@test "${TEST_SECTION}: isctl report contract-status --help works" {
    run ./build/isctl ${ISCTL_OPTIONS} report contract-status --help
    assert_success
    assert_line --partial "Display contract status"
}

@test "${TEST_SECTION}: isctl report contract-status -o json succeeds" {
    JSON=$(./build/isctl ${ISCTL_OPTIONS} report contract-status -o json 2>/dev/null)

    # Output should be a valid JSON array with at least one entry
    COUNT=$(echo "$JSON" | jq 'length')
    [ "$COUNT" -gt 0 ]
}

@test "${TEST_SECTION}: JSON output contains expected fields" {
    JSON=$(./build/isctl ${ISCTL_OPTIONS} report contract-status -o json 2>/dev/null)

    # Check the first entry has all expected fields
    echo "$JSON" | jq -e '.[0].ContractStatus' > /dev/null
    echo "$JSON" | jq -e '.[0] | has("ContractStatusReason")' > /dev/null
    echo "$JSON" | jq -e '.[0] | has("DeviceId")' > /dev/null
    echo "$JSON" | jq -e '.[0] | has("DeviceType")' > /dev/null
    echo "$JSON" | jq -e '.[0] | has("ItemType")' > /dev/null
    echo "$JSON" | jq -e '.[0] | has("PlatformType")' > /dev/null
    echo "$JSON" | jq -e '.[0] | has("ProductNumber")' > /dev/null
}

@test "${TEST_SECTION}: default output works" {
    run ./build/isctl ${ISCTL_OPTIONS} report contract-status
    assert_success
    assert_line --partial "ContractStatus"
}

setup() {
    load 'test_helper/bats-support/load' # this is required by bats-assert!
    load 'test_helper/bats-assert/load'
}
