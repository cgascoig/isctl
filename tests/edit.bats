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
    
    # Should NOT contain read-only properties
    refute_line --partial "Moid:"
    refute_line --partial "ClassId:"
    refute_line --partial "ObjectType:"
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
