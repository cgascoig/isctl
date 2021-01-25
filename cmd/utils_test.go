package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMoRef(t *testing.T) {
	tests := []struct {
		moref  string
		filter string
		ok     bool
	}{
		{
			moref:  "MoRef[Name:isctl-test]",
			filter: "Name eq 'isctl-test'",
			ok:     true,
		},
		{
			moref:  "{\"ClassId\": \"Organization\"}",
			filter: "",
			ok:     false,
		},
		{
			moref:  "default",
			filter: "Name eq 'default'",
			ok:     true,
		},
		{
			moref:  "MoRef[default]",
			filter: "Name eq 'default'",
			ok:     true,
		},
	}

	for _, test := range tests {
		filter, ok := parseMoRef(test.moref)
		assert.Equal(t, test.filter, filter)
		assert.Equal(t, test.ok, ok)
	}
}
