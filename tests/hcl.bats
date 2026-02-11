TEST_SECTION="HCL Status"

@test "${TEST_SECTION}: isctl report --help works" {
    run ./build/isctl ${ISCTL_OPTIONS} report --help
    assert_success
    assert_line --partial "Report commands"
}

@test "${TEST_SECTION}: isctl report hcl --help works" {
    run ./build/isctl ${ISCTL_OPTIONS} report hcl --help
    assert_success
    assert_line --partial "Display HCL status"
}

#bats test_tags=noci
@test "${TEST_SECTION}: isctl report hcl -o json succeeds" {
    # Use subshell to separate stdout (JSON) from stderr (logs)
    JSON=$(./build/isctl ${ISCTL_OPTIONS} report hcl -o json 2>/dev/null)

    # Output should be a valid JSON array with at least one entry
    COUNT=$(echo "$JSON" | jq 'length')
    [ "$COUNT" -gt 0 ]
}

#bats test_tags=noci
@test "${TEST_SECTION}: JSON output contains expected fields" {
    JSON=$(./build/isctl ${ISCTL_OPTIONS} report hcl -o json 2>/dev/null)

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
    JSON=$(./build/isctl ${ISCTL_OPTIONS} report hcl -o json 2>/dev/null)

    # Every Status value should be one of the known HCL statuses
    INVALID=$(echo "$JSON" | jq '[.[].Status] | map(select(. != "Validated" and . != "Not-Listed" and . != "Incomplete" and . != "Not-Evaluated" and . != "")) | length')
    [ "$INVALID" -eq 0 ]
}

@test "${TEST_SECTION}: default output works" {
    run ./build/isctl ${ISCTL_OPTIONS} report hcl
    assert_success
    assert_line --partial "Status"
}

setup() {
    load 'test_helper/bats-support/load' # this is required by bats-assert!
    load 'test_helper/bats-assert/load'
}
