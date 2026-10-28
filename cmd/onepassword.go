package main

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	onePasswordRefPrefix = "op://"
	opPathEnvName        = "ISCTL_OP_PATH"
)

// secretRefKeys are the config keys that may contain 1Password secret references
var secretRefKeys = []string{
	CKIntersightApiKeyId,
	CKIntersightSecretKey,
	CKIntersightClientId,
	CKIntersightClientSecret,
}

func isOnePasswordRef(s string) bool {
	return strings.HasPrefix(strings.TrimSpace(s), onePasswordRefPrefix)
}

func opBinary() string {
	if p := os.Getenv(opPathEnvName); p != "" {
		return p
	}
	return "op"
}

// resolveOnePasswordRefs uses a single invocation of the 1Password CLI (op inject) to resolve all of the
// secret references in refs (map of config key to reference), as each op invocation has significant overhead.
// It returns a map of config key to secret.
func resolveOnePasswordRefs(keys []string, refs map[string]string) (map[string]string, error) {
	nonce := make([]byte, 8)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	beginMarker := func(key string) string { return fmt.Sprintf("<<isctl-begin:%s:%x>>", key, nonce) }
	endMarker := func(key string) string { return fmt.Sprintf("<<isctl-end:%s:%x>>", key, nonce) }

	// the template only contains the references (not secrets), each secret is delimited by unique markers
	var tpl strings.Builder
	var refList []string
	for _, key := range keys {
		ref := strings.TrimSpace(refs[key])
		refList = append(refList, ref)
		fmt.Fprintf(&tpl, "%s{{ %s }}%s\n", beginMarker(key), ref, endMarker(key))
	}

	// pass the template as a file so that stdin remains available for op to prompt for sign in
	tplFile, err := os.CreateTemp("", "isctl-op-template-*")
	if err != nil {
		return nil, fmt.Errorf("unable to create 1Password template file: %v", err)
	}
	defer os.Remove(tplFile.Name())
	if _, err := tplFile.WriteString(tpl.String()); err != nil {
		tplFile.Close()
		return nil, fmt.Errorf("unable to write 1Password template file: %v", err)
	}
	tplFile.Close()

	var stdout bytes.Buffer
	cmd := exec.Command(opBinary(), "inject", "-i", tplFile.Name())
	// connect stdin/stderr so that op can prompt for sign in / biometric unlock
	cmd.Stdin = os.Stdin
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr

	log.Debugf("Running 1Password CLI: %s (refs: %s)", strings.Join(cmd.Args, " "), strings.Join(refList, ", "))
	start := time.Now()
	err = cmd.Run()
	log.Debugf("1Password CLI finished in %v (%d bytes returned, err: %v)", time.Since(start), stdout.Len(), err)

	if err != nil {
		if errors.Is(err, exec.ErrNotFound) || errors.Is(err, fs.ErrNotExist) {
			return nil, fmt.Errorf("1Password CLI (%s) not found, set %s if it is not in PATH", opBinary(), opPathEnvName)
		}
		return nil, err
	}

	out := stdout.String()
	secrets := make(map[string]string, len(keys))
	for _, key := range keys {
		_, rest, ok := strings.Cut(out, beginMarker(key))
		if !ok {
			return nil, fmt.Errorf("unexpected output from 1Password CLI: missing secret for %s", key)
		}
		secret, _, ok := strings.Cut(rest, endMarker(key))
		if !ok {
			return nil, fmt.Errorf("unexpected output from 1Password CLI: missing secret for %s", key)
		}
		secrets[key] = secret
	}
	return secrets, nil
}

// resolveSecretRefs replaces any 1Password secret references in the credential config items with the secret values
func resolveSecretRefs() error {
	var keys []string
	refs := map[string]string{}
	for _, key := range secretRefKeys {
		if val := gK.String(key); isOnePasswordRef(val) {
			keys = append(keys, key)
			refs[key] = val
		}
	}
	if len(keys) == 0 {
		return nil
	}

	log.Debugf("Resolving 1Password references for %s", strings.Join(keys, ", "))
	start := time.Now()
	secrets, err := resolveOnePasswordRefs(keys, refs)
	if err != nil {
		return fmt.Errorf("unable to resolve 1Password reference(s) for %s: %v", strings.Join(keys, ", "), err)
	}
	for _, key := range keys {
		gK.Set(key, secrets[key])
	}
	log.Debugf("Resolved %d 1Password reference(s) in %v", len(keys), time.Since(start))
	return nil
}
