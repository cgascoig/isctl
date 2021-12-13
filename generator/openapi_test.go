package main

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestToCamelCase(t *testing.T) {
	assert.Equal(t, "AaaAuditRecord", ToCamelCase("aaa.AuditRecord"))
	assert.Equal(t, "AaaAuditRecord", ToCamelCase("AaaAuditRecord"))
	assert.Equal(t, "AaaAuditRecord", ToCamelCase("Aaa Audit Record"))
	assert.Equal(t, "AaaAuditRecord", ToCamelCase("aaa    audit record"))
}

func TestSchemaRefToType(t *testing.T) {
	assert.Equal(t, "", SchemaRefToType(&openapi3.SchemaRef{}))
	assert.Equal(t, "", SchemaRefToType(nil))
	assert.Equal(t, "VnicEthNetworkPolicy", SchemaRefToType(&openapi3.SchemaRef{
		Ref: "#/components/schemas/vnic.EthNetworkPolicy",
	}))
}
