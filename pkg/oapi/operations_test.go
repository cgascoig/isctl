package oapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBodyParamVars(t *testing.T) {
	vars := getBodyParamVars("ntp.Policy")
	names := []string{}
	for _, v := range vars {
		names = append(names, v.Name)
	}
	expected := []string{
		"SharedScope", "Ancestors", "DomainGroupMoid", "Owners", "Tags", "DisplayNames", "ObjectType",
		"ModTime", "Moid", "VersionContext", "Parent", "PermissionResources", "ClassId", "AccountMoid",
		"CreateTime", "ClassId", "ObjectType", "Description", "Name", "ApplianceAccount", "Organization",
		"Profiles", "ClassId", "ObjectType", "Timezone", "AuthenticatedNtpServers", "Enabled", "NtpServers"}
	assert.ElementsMatch(t, names, expected)
}

func TestFindOperation(t *testing.T) {
	op := FindOperation("get", "ntp.Policy")
	assert.NotNil(t, op)
	assert.Equal(t, "GetNtpPolicyList", op.OperationID)

	op = FindOperation("post", "ntp.Policy")
	assert.NotNil(t, op)
	assert.Equal(t, "CreateNtpPolicy", op.OperationID)

	op = FindOperation("patch", "ntp.Policy")
	assert.NotNil(t, op)
	assert.Equal(t, "UpdateNtpPolicy", op.OperationID)

	op = FindOperation("delete", "ntp.Policy")
	assert.NotNil(t, op)
	assert.Equal(t, "DeleteNtpPolicy", op.OperationID)
}
