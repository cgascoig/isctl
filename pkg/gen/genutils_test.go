package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cgascoig/isctl/pkg/oapi"
)

func TestGetMoMoRefDirectMoid(t *testing.T) {
	moref := &oapi.MoRef{Moid: "59c84e4a16267c0001c23428"}
	result, err := GetMoMoRef(nil, moref)
	assert.NoError(t, err)
	assert.Equal(t, map[string]any{
		"ClassId": "mo.MoRef",
		"Moid":    "59c84e4a16267c0001c23428",
	}, result)
	_, hasObjectType := result["ObjectType"]
	assert.False(t, hasObjectType)
}

func TestGetMoMoRefDirectMoidWithType(t *testing.T) {
	moref := &oapi.MoRef{Moid: "59c84e4a16267c0001c23428", RelationshipType: "server.Profile"}
	result, err := GetMoMoRef(nil, moref)
	assert.NoError(t, err)
	assert.Equal(t, map[string]any{
		"ClassId":    "mo.MoRef",
		"Moid":       "59c84e4a16267c0001c23428",
		"ObjectType": "server.Profile",
	}, result)
}

func TestReplaceArgs(t *testing.T) {
	var s string
	var e error
	s, e = ReplaceArgs("api/{Moid}", []string{"12345"})
	assert.NoError(t, e)
	assert.Equal(t, "api/12345", s)
}

func TestGetOperationForRelationshipAbstractClass(t *testing.T) {
	// Abstract classes should have no direct GET operation, which triggers the fallback path
	op := GetOperationForRelationship("resource.AbstractResourceQualificationPolicy.Relationship")
	assert.Nil(t, op, "abstract class should not have a direct GET operation")

	// Concrete classes should have a GET operation
	op = GetOperationForRelationship("ntp.Policy.Relationship")
	assert.NotNil(t, op, "concrete class ntp.Policy should have a GET operation")
}

func TestAppendResults(t *testing.T) {
	in := map[string]any{
		"Results": []any{
			map[string]any{
				"a": 1,
			},
		},
	}
	new := map[string]any{
		"Results": []any{
			map[string]any{
				"b": 2,
			},
		},
	}
	expected := map[string]any{
		"Results": []any{
			map[string]any{
				"a": 1,
			},
			map[string]any{
				"b": 2,
			},
		},
	}

	out, count, err := appendResults(in, new)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
	assert.Equal(t, 1, count)
}
