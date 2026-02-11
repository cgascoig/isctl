TEST_SECTION="HCL Status"

@test "${TEST_SECTION}: isctl hcl --help works" {
    run ./build/isctl ${ISCTL_OPTIONS} hcl --help
    assert_success
    assert_line --partial "Hardware Compatibility List"
}

@test "${TEST_SECTION}: isctl hcl status --help works" {
    run ./build/isctl ${ISCTL_OPTIONS} hcl status --help
    assert_success
    assert_line --partial "Display HCL status"
}

@test "${TEST_SECTION}: isctl hcl status -o json succeeds" {
    # Use subshell to separate stdout (JSON) from stderr (logs)
    JSON=$(./build/isctl ${ISCTL_OPTIONS} hcl status -o json 2>/dev/null)

    # Output should be a valid JSON array with at least one entry
    COUNT=$(echo "$JSON" | jq 'length')
    [ "$COUNT" -gt 0 ]
}

@test "${TEST_SECTION}: JSON output contains expected fields" {
    JSON=$(./build/isctl ${ISCTL_OPTIONS} hcl status -o json 2>/dev/null)

    # Check the first entry has all expected fields
    echo "$JSON" | jq -e '.[0].Status' > /dev/null
    echo "$JSON" | jq -e '.[0].HclModel' > /dev/null
    echo "$JSON" | jq -e '.[0].HclFirmwareVersion' > /dev/null
    echo "$JSON" | jq -e '.[0].HclOsVendor' > /dev/null
    echo "$JSON" | jq -e '.[0].HclOsVersion' > /dev/null
    echo "$JSON" | jq -e '.[0].HclProcessor' > /dev/null
    echo "$JSON" | jq -e '.[0].ComponentStatus' > /dev/null
    echo "$JSON" | jq -e '.[0].HardwareStatus' > /dev/null
    echo "$JSON" | jq -e '.[0].SoftwareStatus' > /dev/null
    echo "$JSON" | jq -e '.[0].ServerReason' > /dev/null
    echo "$JSON" | jq -e '.[0] | has("Details")' > /dev/null
}

@test "${TEST_SECTION}: Status field has valid values" {
    JSON=$(./build/isctl ${ISCTL_OPTIONS} hcl status -o json 2>/dev/null)

    # Every Status value should be one of the known HCL statuses
    INVALID=$(echo "$JSON" | jq '[.[].Status] | map(select(. != "Validated" and . != "Not-Listed" and . != "Incomplete" and . != "Not-Evaluated" and . != "")) | length')
    [ "$INVALID" -eq 0 ]
}

@test "${TEST_SECTION}: default output works" {
    run ./build/isctl ${ISCTL_OPTIONS} hcl status
    assert_success
    assert_line --partial "Status"
}

setup() {
    load 'test_helper/bats-support/load' # this is required by bats-assert!
    load 'test_helper/bats-assert/load'
}
