TEST_SECTION="Basic Commands"

@test "${TEST_SECTION}: invoking isctl version outputs version" {
    run ./build/isctl ${ISCTL_OPTIONS} version
    [ "$status" -eq 0 ]
    [[ "${lines[0]}" == "isctl version "* ]]
}

@test "${TEST_SECTION}: invoking isctl with invalid command exit code non-zero" {
    run ./build/isctl ${ISCTL_OPTIONS} invalid_command
    [ "$status" -ne 0 ]
}

@test "${TEST_SECTION}: invoking isctl with no command works" {
    run ./build/isctl ${ISCTL_OPTIONS}
    assert_success
    assert_line "Usage:"
    assert_line "  isctl [command]"
}

setup() {
    load 'test_helper/bats-support/load' # this is required by bats-assert!
    load 'test_helper/bats-assert/load'
}