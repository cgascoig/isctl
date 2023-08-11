package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsKeyData(t *testing.T) {
	tests := []struct {
		in       string
		out      bool
		testName string
	}{
		{
			// Dummy key - not real
			in: `-----BEGIN RSA PRIVATE KEY-----
			MIIEogIBAAKCAQEAtfV97Tef7flu8khNdOvOyFzt847nTRH9zrbySOCw1BOjVsp/
			xzJvshZF5nE9wuQbWn0WHbK8E5S35qGC6viWHbq1LrYyqKc0YCjKCYnH++E+8PH9
			r8mp+ENaUYJEN6tx4nWZznwWj0eUEbtd3N5YKXwqi+HqBbtqwtM=
			-----END RSA PRIVATE KEY-----`,
			out:      true,
			testName: "RSA key",
		},
		{
			// Dummy key - not real
			in: `
			-----BEGIN RSA PRIVATE KEY-----
			MIIEogIBAAKCAQEAtfV97Tef7flu8khNdOvOyFzt847nTRH9zrbySOCw1BOjVsp/
			xzJvshZF5nE9wuQbWn0WHbK8E5S35qGC6viWHbq1LrYyqKc0YCjKCYnH++E+8PH9
			r8mp+ENaUYJEN6tx4nWZznwWj0eUEbtd3N5YKXwqi+HqBbtqwtM=
			-----END RSA PRIVATE KEY-----
			`,
			out:      true,
			testName: "RSA key with whitespace",
		},
		{
			in:       "~/intersight.pem",
			out:      false,
			testName: "Filename",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, isKeyData(test.in), test.testName)
	}
}
