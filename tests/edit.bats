TEST_EDIT_POLICY_NAME=isctl-bats-test-edit-policy

TEST_SECTION="Edit Command"

@test "${TEST_SECTION}: help shows available commands" {
    run ./build/isctl ${ISCTL_OPTIONS} edit --help
    assert_success
    assert_line --partial "Edit an Intersight resource"
    assert_line --partial "edit ntp policy name my-ntp-policy"
}

@test "${TEST_SECTION}: edit ntp policy shows subcommands" {
    run ./build/isctl ${ISCTL_OPTIONS} edit ntp policy --help
    assert_success
    assert_line --partial "moid"
    assert_line --partial "name"
}

@test "${TEST_SECTION}: edit with EDITOR=cat shows no changes" {
    # Create a test policy first
    ./build/isctl ${ISCTL_OPTIONS} create ntp policy --Name "${TEST_EDIT_POLICY_NAME}" --NtpServers 1.1.1.1 --Organization default

    # Edit with cat as editor (no changes possible)
    run env EDITOR=cat ./build/isctl ${ISCTL_OPTIONS} edit ntp policy name "${TEST_EDIT_POLICY_NAME}"
    assert_success
    assert_line --partial "no changes made"
}

@test "${TEST_SECTION}: edit output contains only editable properties" {
    # Use cat to dump what would be shown in editor
    run env EDITOR=cat ./build/isctl ${ISCTL_OPTIONS} edit ntp policy name "${TEST_EDIT_POLICY_NAME}"
    assert_success
    
    # Should contain editable properties
    assert_line --partial "Name"
    assert_line --partial "Enabled"
    
    # Should match user's customized behavior (Moid is now included)
    assert_line --partial "Moid:"
    assert_line --partial "ClassId:"
    assert_line --partial "ObjectType:"
    refute_line --partial "CreateTime:"
}
@test "${TEST_SECTION}: edit actually updates the resource" {
    # First ensure the policy is enabled
    ./build/isctl ${ISCTL_OPTIONS} update ntp policy name "${TEST_EDIT_POLICY_NAME}" --Enabled=true
    
    # Verify it's enabled
    ENABLED=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_EDIT_POLICY_NAME}" -o json | jq -r .Enabled )
    assert_equal "${ENABLED}" "true"
    
    # Create a wrapper script that uses sed to change Enabled from true to false
    EDITOR_SCRIPT=$(mktemp)
    cat > "${EDITOR_SCRIPT}" << 'EOF'
#!/bin/bash
sed -i.bak 's/Enabled: true/Enabled: false/' "$1"
EOF
    chmod +x "${EDITOR_SCRIPT}"
    
    # Run edit with our sed wrapper as the editor
    run env EDITOR="${EDITOR_SCRIPT}" ./build/isctl ${ISCTL_OPTIONS} edit ntp policy name "${TEST_EDIT_POLICY_NAME}"
    rm -f "${EDITOR_SCRIPT}"
    
    assert_success
    assert_line --partial "Edit applied successfully"
    
    # Verify the change was actually applied
    ENABLED=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_EDIT_POLICY_NAME}" -o json | jq -r .Enabled )
    assert_equal "${ENABLED}" "false"
}

@test "${TEST_SECTION}: cleanup test policy" {
    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy name "${TEST_EDIT_POLICY_NAME}"
    assert_success
}

setup_file() {
    # Clean up any leftover test policy
    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy name "${TEST_EDIT_POLICY_NAME}"
}

setup() {
    load 'test_helper/bats-support/load'
    load 'test_helper/bats-assert/load'
}

@test "${TEST_SECTION}: edit multiple objects via yaml-editable export" {
    # Setup: Create 2 start-up NTP policies
    POLICY_PREFIX="isctl-bats-edit-test"
    POLICY1="${POLICY_PREFIX}-1"
    POLICY2="${POLICY_PREFIX}-2"
    
    # Ensure clean state
    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy --filter "startswith(Name, '${POLICY_PREFIX}')"
    
    ./build/isctl ${ISCTL_OPTIONS} create ntp policy --Name "${POLICY1}" --NtpServers 1.1.1.1 --Organization default
    ./build/isctl ${ISCTL_OPTIONS} create ntp policy --Name "${POLICY2}" --NtpServers 2.2.2.2 --Organization default
    
    # Action: Export with filter and -o yaml-editable
    EXPORT_FILE=$(mktemp)
    ./build/isctl ${ISCTL_OPTIONS} get ntp policy --filter "startswith(Name, '${POLICY_PREFIX}')" -o yaml-editable > "${EXPORT_FILE}"
    
    # Verify export contains both policies
    grep -q "${POLICY1}" "${EXPORT_FILE}"
    grep -q "${POLICY2}" "${EXPORT_FILE}"
    
    echo "--- Exported YAML (before edit) ---"
    cat "${EXPORT_FILE}"
    
    # Action: Modify the exported YAML
    # Change NTP servers to 3.3.3.3
    # Use sed to replace the specific lines
    sed -i.bak 's/- 1.1.1.1/- 3.3.3.3/' "${EXPORT_FILE}"
    sed -i.bak 's/- 2.2.2.2/- 3.3.3.3/' "${EXPORT_FILE}"
    
    echo "--- Exported YAML (after edit) ---"
    cat "${EXPORT_FILE}"
    cat "${EXPORT_FILE}"
    
    # Action: Apply the modified YAML
    run ./build/isctl ${ISCTL_OPTIONS} apply -f "${EXPORT_FILE}"
    assert_success
    rm -f "${EXPORT_FILE}" "${EXPORT_FILE}.bak"
    
    # Verification: Check updates on both objects
    SERVER1=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${POLICY1}" -o json | jq -r '.NtpServers[0]' )
    assert_equal "${SERVER1}" "3.3.3.3"
    
    SERVER2=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${POLICY2}" -o json | jq -r '.NtpServers[0]' )
    assert_equal "${SERVER2}" "3.3.3.3"
    
    # Cleanup
    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy --filter "startswith(Name, '${POLICY_PREFIX}')"
}
