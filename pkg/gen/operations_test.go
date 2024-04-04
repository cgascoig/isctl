package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUpdateOperationForClassID(t *testing.T) {
	op := GetUpdateOperationForClassID("ntp.Policy")
	assert.NotNil(t, op)
	assert.Equal(t, "UpdateNtpPolicy", op.operation.OperationID)
}

func TestGetCreateOperationForClassID(t *testing.T) {
	op := GetCreateOperationForClassID("ntp.Policy")
	assert.NotNil(t, op)
	assert.Equal(t, "CreateNtpPolicy", op.operation.OperationID)
}

func TestGetGetOperationForClassID(t *testing.T) {
	op := GetGetOperationForClassID("ntp.Policy")
	assert.NotNil(t, op)
	assert.Equal(t, "GetNtpPolicyList", op.operation.OperationID)
}

func TestGetDeleteOperationForClassID(t *testing.T) {
	op := GetDeleteOperationForClassID("ntp.Policy")
	assert.NotNil(t, op)
	assert.Equal(t, "DeleteNtpPolicy", op.operation.OperationID)
}

func TestGetReferencedClasses(t *testing.T) {
	op := GetUpdateOperationForClassID("ntp.Policy")

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

	op = GetUpdateOperationForClassID("kubernetes.ClusterProfile")

	refClasses, err = op.GetReferencedClasses(map[string]interface{}{
		"ClusterIpPools": []interface{}{"MoRef[ip-pool-1]"},
	})

	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"ippool.Pool"}, refClasses)

	op = GetUpdateOperationForClassID("server.ProfileTemplate")

	refClasses, err = op.GetReferencedClasses(map[string]interface{}{
		"PolicyBucket": []any{
			"MoRef:BiosPolicyRelationship[cgascoig-bios-policy]",
			"MoRef:IamLdapPolicyRelationship[cgascoig-ldap-policy]",
		},
		"Organization": "default",
	})

	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"organization.Organization", "bios.Policy", "iam.LdapPolicy"}, refClasses)

	op = GetUpdateOperationForClassID("bulk.MoCloner")
	refClasses, err = op.GetReferencedClasses(map[string]any{
		"ClassId":      "bulk.MoCloner",
		"Organization": "default",
		"Sources":      []any{"MoRef:ServerProfileTemplateRelationship[OCP-BM]"},
	})

	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{
		"organization.Organization",
		"server.ProfileTemplate",
	}, refClasses)
}
