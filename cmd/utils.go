package main

import (
	"fmt"
	"regexp"
)

//Parse the MoRef string and return a filter or nil
func parseMoRef(moref string) (string, bool) {
	r := regexp.MustCompile(`MoRef\[(\w+):([0-9A-Za-z_-]+)\]`)

	m := r.FindStringSubmatch(moref)
	if m != nil {
		return fmt.Sprintf("%s eq '%s'", m[1], m[2]), true
	}

	r = regexp.MustCompile(`^\s*([0-9A-Za-z_-]+)\s*$`)
	m = r.FindStringSubmatch(moref)
	if m != nil {
		return fmt.Sprintf("Name eq '%s'", m[1]), true
	}

	return "", false
}
