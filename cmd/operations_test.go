package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUpdateOperationForClassID(t *testing.T) {
	op := getUpdateOperationForClassID("ntp.Policy")
	_, ok := op.(*UpdateNtpPolicy)
	assert.True(t, ok)
}

func TestGetCreateOperationForClassID(t *testing.T) {
	op := getCreateOperationForClassID("ntp.Policy")
	_, ok := op.(*CreateNtpPolicy)
	assert.True(t, ok)
}

func TestGetGetOperationForClassID(t *testing.T) {
	op := getGetOperationForClassID("ntp.Policy")
	_, ok := op.(*GetNtpPolicyList)
	assert.True(t, ok)
}

func TestGetDeleteOperationForClassID(t *testing.T) {
	op := getDeleteOperationForClassID("ntp.Policy")
	_, ok := op.(*DeleteNtpPolicy)
	assert.True(t, ok)
}

func TestGetReferencedClasses(t *testing.T) {
	op := getUpdateOperationForClassID("ntp.Policy")

	refClasses, err := op.GetReferencedClasses(map[string]interface{}{
		"Organization": "MoRef[Name:default]",
	})

	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"organization.Organization"}, refClasses)

	refClasses, err = op.GetReferencedClasses(map[string]interface{}{
		"Organization":     "MoRef[Name:default]",
		"ApplianceAccount": "MoRef[Name:default]",
	})

	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"organization.Organization", "iam.Account"}, refClasses)

	refClasses, err = op.GetReferencedClasses(map[string]interface{}{
		"Organization": "MoRef[default]",
	})

	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"organization.Organization"}, refClasses)

	op = getUpdateOperationForClassID("kubernetes.ClusterProfile")

	refClasses, err = op.GetReferencedClasses(map[string]interface{}{
		"ClusterIpPools": []interface{}{"MoRef[ip-pool-1]"},
	})

	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"ippool.Pool"}, refClasses)
}
