package oapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBodyParamVars(t *testing.T) {
	assert.Equal(t, []*Var{}, getBodyParamVars("ntp.Policy"))
}

func TestFindOperation(t *testing.T) {
	op := FindOperation("get", "ntp.Policy")
	assert.NotNil(t, op)
	assert.Equal(t, "GetNtpPolicyList", op.OperationID)

	op = FindOperation("post", "ntp.Policy")
	assert.NotNil(t, op)
	assert.Equal(t, "CreateNtpPolicy", op.OperationID)

	op = FindOperation("put", "ntp.Policy")
	assert.NotNil(t, op)
	assert.Equal(t, "UpdateNtpPolicy", op.OperationID)

	op = FindOperation("delete", "ntp.Policy")
	assert.NotNil(t, op)
	assert.Equal(t, "DeleteNtpPolicy", op.OperationID)
}
