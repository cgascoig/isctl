@test "invoking isctl version outputs version" {
    run ./build/isctl version
    [ "$status" -eq 0 ]
    [[ "${lines[0]}" == "isctl version "* ]]
}

@test "invoking isctl with invalid command exit code non-zero" {
    run ./build/isctl invalid_command
    [ "$status" -ne 0 ]
}