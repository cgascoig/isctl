load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'

setup() {
    # common_setup # common setup not found in tests/apply.bats, seemingly not used or implicit?
    # actually apply.bats doesn't load common. It defines setup() locally.
    # I'll just use what I need.
    :
}

teardown() {
    # Clean up the test object if it exists
    ./build/isctl ${ISCTL_OPTIONS} apply -d -f "$BATS_TEST_TMPDIR/template.yaml" --var myname="template-test" || true
}

@test "apply creates object using templates and variable precedence" {
    # cleanup any stale objects
    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy name template-test

    # 1. Create a template file
    cat <<EOF > "$BATS_TEST_TMPDIR/template.yaml"
ClassId: ntp.Policy
Name: {{ .Vars.myname }}
Organization: default
Description: "Created by {{ .Vars.owner | default "unknown" }} from {{ .Vars.source | upper }} Env Test: {{ .Vars.envtest }}"
NtpServers:
  - 1.2.3.4
EOF

    # 2. Create a variables file
    cat <<EOF > "$BATS_TEST_TMPDIR/vars.yaml"
myname: file-name
owner: file-owner
source: file-source
EOF

    # 3. Apply with flags overriding file, and file overriding env
    # ISCTL_VAR_owner=env-owner (overridden by file)
    # --var-file (provides defaults)
    # --var myname=template-test (overrides file)
    
    export ISCTL_VAR_owner="env-owner"
    export ISCTL_VAR_source="env-source"
    export ISCTL_VAR_envtest="env-test"
    
    run ./build/isctl ${ISCTL_OPTIONS} apply -f "$BATS_TEST_TMPDIR/template.yaml" \
        --var-file "$BATS_TEST_TMPDIR/vars.yaml" \
        --var myname="template-test"

    assert_success
    assert_output --partial "Performing create operation on new MO"
    assert_output --partial "Name: template-test"

    # 4. Verify the object was created with correct values
    # Name should be "template-test" (from flag)
    # Description should check precedence: 
    #   owner: "file-owner" (file overrides env)
    #   source: "file-source" (file overrides env) -> upper -> "FILE-SOURCE"
    
    run ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name template-test -o yaml
    assert_success
    assert_output --partial "Name: template-test"
    assert_output --partial "Description: 'Created by file-owner from FILE-SOURCE Env Test: env-test'"
    
    unset ISCTL_VAR_owner
    unset ISCTL_VAR_source
    unset ISCTL_VAR_envtest
}

@test "apply fails with missing variable and no default" {
    cat <<EOF > "$BATS_TEST_TMPDIR/invalid_template.yaml"
ClassId: ntp.Policy
Name: {{ .Vars.missing }}
EOF

    run ./build/isctl ${ISCTL_OPTIONS} apply -f "$BATS_TEST_TMPDIR/invalid_template.yaml"
    assert_failure
    # Template processing succeeds (empty value), but API validation fails because Name is empty/invalid
    assert_output --partial "Cannot set the property 'policy.AbstractPolicy.Name'"
}

@test "apply loads variables from isctl.vars.yaml in directory" {
    # cleanup any stale objects
    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy name dir-var-test

    # 1. Create a directory structure
    mkdir -p "$BATS_TEST_TMPDIR/subdir"

    # 2. Create a template file in the directory
    cat <<EOF > "$BATS_TEST_TMPDIR/subdir/template.yaml"
ClassId: ntp.Policy
Name: {{ .Vars.name }}
Organization: default
Description: "Loaded from {{ .Vars.source }}"
NtpServers:
  - 1.2.3.4
EOF

    # 3. Create isctl.vars.yaml in the directory
    cat <<EOF > "$BATS_TEST_TMPDIR/subdir/isctl.vars.yaml"
name: dir-var-test
source: dir-file
EOF

    # 4. Apply the directory
    run ./build/isctl ${ISCTL_OPTIONS} apply -f "$BATS_TEST_TMPDIR/subdir"
    assert_success
    assert_output --partial "Performing create operation on new MO"
    assert_output --partial "Name: dir-var-test"

    # 5. Verify object
    run ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name dir-var-test -o yaml
    assert_success
    assert_output --partial "Name: dir-var-test"
    assert_output --partial "Description: Loaded from dir-file"

    # 6. Verify precedence (env > dir)
    export ISCTL_VAR_source="env-override"
    run ./build/isctl ${ISCTL_OPTIONS} apply -f "$BATS_TEST_TMPDIR/subdir"
    assert_success
    
    run ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name dir-var-test -o yaml
    assert_success
    assert_output --partial "Description: Loaded from env-override"
    
    unset ISCTL_VAR_source
    rm -rf "$BATS_TEST_TMPDIR/subdir"
}
