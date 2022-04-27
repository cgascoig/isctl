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