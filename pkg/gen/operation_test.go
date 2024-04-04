package gen

import (
	"testing"

	"github.com/cgascoig/isctl/pkg/oapi"
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
	refClasses := GetReferencedClasses(map[string]interface{}{
		"Organization": &oapi.MoRef{
			Filter:           "Name eq 'default",
			RelationshipType: "organization.Organization",
		},
	})

	assert.ElementsMatch(t, []string{"organization.Organization"}, refClasses)

	mo := map[string]interface{}{
		"Organization":     "MoRef[Name:default]",
		"ApplianceAccount": "MoRef[Name:default]",
	}
	oapi.CanonicaliseMoRefs(&mo, "ntp.Policy")
	refClasses = GetReferencedClasses(mo)

	assert.ElementsMatch(t, []string{"organization.Organization", "iam.Account"}, refClasses)

	mo = map[string]interface{}{
		"ClusterIpPools": []interface{}{"MoRef[ip-pool-1]"},
	}
	oapi.CanonicaliseMoRefs(&mo, "kubernetes.ClusterProfile")

	refClasses = GetReferencedClasses(mo)

	assert.ElementsMatch(t, []string{"ippool.Pool"}, refClasses)

	mo = map[string]interface{}{
		"PolicyBucket": []any{
			"MoRef:BiosPolicyRelationship[cgascoig-bios-policy]",
			"MoRef:IamLdapPolicyRelationship[cgascoig-ldap-policy]",
		},
		"Organization": "default",
	}
	oapi.CanonicaliseMoRefs(&mo, "server.ProfileTemplate")

	refClasses = GetReferencedClasses(mo)

	assert.ElementsMatch(t, []string{"organization.Organization", "bios.Policy", "iam.LdapPolicy"}, refClasses)

	mo = map[string]any{
		"ClassId":      "bulk.MoCloner",
		"Organization": "default",
		"Sources":      []any{"MoRef:ServerProfileTemplateRelationship[OCP-BM]"},
	}
	oapi.CanonicaliseMoRefs(&mo, "bulk.MoCloner")
	refClasses = GetReferencedClasses(mo)

	assert.ElementsMatch(t, []string{
		"organization.Organization",
		"server.ProfileTemplate",
	}, refClasses)
}
