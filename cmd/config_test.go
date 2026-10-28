package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"strings"
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

func TestNormalizeKeyData(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(t, err)
	multiLine := strings.TrimSpace(string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})))

	tests := []struct {
		in       string
		testName string
	}{
		{multiLine, "multi-line key unchanged"},
		{multiLine + "\n", "multi-line key with trailing newline"},
		{strings.ReplaceAll(multiLine, "\n", " "), "newlines converted to spaces"},
		{" " + strings.ReplaceAll(multiLine, "\n", " ") + " ", "newlines converted to spaces with surrounding whitespace"},
		{strings.ReplaceAll(multiLine, "\n", ""), "newlines removed"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			out := normalizeKeyData(tt.in)
			block, _ := pem.Decode([]byte(out))
			if assert.NotNil(t, block, "PEM decode failed for %q", out) {
				_, err := x509.ParsePKCS1PrivateKey(block.Bytes)
				assert.NoError(t, err)
			}
		})
	}

	assert.Equal(t, multiLine, normalizeKeyData(strings.ReplaceAll(multiLine, "\n", " ")))
	assert.Equal(t, "~/intersight.pem", normalizeKeyData("~/intersight.pem"))
}
