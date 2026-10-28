package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
)

const fakeOpScript = `#!/usr/bin/env bash
# fake op: supports "op inject -i <file>"
[ "$1" = "inject" ] && [ "$2" = "-i" ] || { echo "unsupported: $*" >&2; exit 1; }
echo "$*" >> "$(dirname "$0")/invocations"
tpl=$(cat "$3")
sub() { local p="{{ $1 }}"; tpl=${tpl//"$p"/"$2"}; }
sub op://vault/item/key_id resolved_key_id
sub op://vault/item/secret resolved_secret
sub op://vault/item/multiline $'line1\nline2\nline3'
if [[ "$tpl" == *"{{"* ]]; then echo "item not found" >&2; exit 1; fi
printf '%s\n' "$tpl"
`

func setupFakeOp(t *testing.T) string {
	dir := t.TempDir()
	opPath := filepath.Join(dir, "op")
	assert.NoError(t, os.WriteFile(opPath, []byte(fakeOpScript), 0755))
	t.Setenv(opPathEnvName, opPath)

	oldK := gK
	gK = koanf.New(".")
	t.Cleanup(func() { gK = oldK })
	return dir
}

func TestIsOnePasswordRef(t *testing.T) {
	assert.True(t, isOnePasswordRef("op://vault/item/field"))
	assert.True(t, isOnePasswordRef("  op://vault/item/field\n"))
	assert.False(t, isOnePasswordRef("~/intersight.pem"))
	assert.False(t, isOnePasswordRef("-----BEGIN RSA PRIVATE KEY-----"))
	assert.False(t, isOnePasswordRef(""))
}

func TestResolveSecretRefs(t *testing.T) {
	dir := setupFakeOp(t)
	gK.Set(CKIntersightApiKeyId, "op://vault/item/key_id")
	gK.Set(CKIntersightSecretKey, "op://vault/item/multiline")
	gK.Set(CKIntersightClientId, "my_client_id")
	gK.Set(CKIntersightClientSecret, " op://vault/item/secret ")

	assert.NoError(t, resolveSecretRefs())
	assert.Equal(t, "resolved_key_id", gK.String(CKIntersightApiKeyId))
	assert.Equal(t, "line1\nline2\nline3", gK.String(CKIntersightSecretKey))
	assert.Equal(t, "my_client_id", gK.String(CKIntersightClientId))
	assert.Equal(t, "resolved_secret", gK.String(CKIntersightClientSecret))

	// all references should be resolved with a single op invocation
	invocations, err := os.ReadFile(filepath.Join(dir, "invocations"))
	assert.NoError(t, err)
	assert.Equal(t, 1, strings.Count(string(invocations), "\n"))
}

func TestResolveSecretRefsNoRefs(t *testing.T) {
	dir := setupFakeOp(t)
	gK.Set(CKIntersightApiKeyId, "key_id")
	gK.Set(CKIntersightSecretKey, "~/intersight.pem")

	assert.NoError(t, resolveSecretRefs())
	assert.Equal(t, "key_id", gK.String(CKIntersightApiKeyId))
	assert.Equal(t, "~/intersight.pem", gK.String(CKIntersightSecretKey))

	// op should not be invoked when there are no references
	_, err := os.Stat(filepath.Join(dir, "invocations"))
	assert.True(t, os.IsNotExist(err))
}

func TestResolveSecretRefsError(t *testing.T) {
	setupFakeOp(t)
	gK.Set(CKIntersightSecretKey, "op://vault/item/missing")

	err := resolveSecretRefs()
	assert.ErrorContains(t, err, CKIntersightSecretKey)
}

func TestResolveSecretRefsOpNotFound(t *testing.T) {
	setupFakeOp(t)
	t.Setenv(opPathEnvName, filepath.Join(t.TempDir(), "does-not-exist"))
	gK.Set(CKIntersightSecretKey, "op://vault/item/secret")

	err := resolveSecretRefs()
	assert.ErrorContains(t, err, "not found")
}
