package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMoRef(t *testing.T) {
	tests := []struct {
		moref    string
		filter   string
		datatype string
		ok       bool
	}{
		{
			moref:    "MoRef[Name:isctl-test]",
			filter:   "Name eq 'isctl-test'",
			datatype: "",
			ok:       true,
		},
		{
			moref:    "MoRef[$filter:DeviceHostname eq 'XYZ' and PlatformType eq 'IWE']",
			filter:   "DeviceHostname eq 'XYZ' and PlatformType eq 'IWE'",
			datatype: "",
			ok:       true,
		},
		{
			moref:    "MoRef:KubernetesVirtualMachineInfraConfigPolicyRelationship[$filter:DeviceHostname eq 'XYZ' and PlatformType eq 'IWE']",
			filter:   "DeviceHostname eq 'XYZ' and PlatformType eq 'IWE'",
			datatype: "KubernetesVirtualMachineInfraConfigPolicyRelationship",
			ok:       true,
		},
		{
			moref:    "{\"ClassId\": \"Organization\"}",
			filter:   "",
			datatype: "",
			ok:       false,
		},
		{
			moref:    "cg-k8s-1.18.2",
			filter:   "Name eq 'cg-k8s-1.18.2'",
			datatype: "",
			ok:       true,
		},
		{
			moref:    "cg-k8s-1.18[2",
			filter:   "",
			datatype: "",
			ok:       false,
		},
		{
			moref:    "MoRef[default]",
			filter:   "Name eq 'default'",
			datatype: "",
			ok:       true,
		},
		{
			moref:    "MoRef:KubernetesVirtualMachineInfrastructureProviderRelationship[default]",
			filter:   "Name eq 'default'",
			datatype: "KubernetesVirtualMachineInfrastructureProviderRelationship",
			ok:       true,
		},
		{
			moref:    "MoRef:KubernetesInfrastructureProviderRelationship[Name:default]",
			filter:   "Name eq 'default'",
			datatype: "KubernetesInfrastructureProviderRelationship",
			ok:       true,
		},
		{
			moref:    "MoRef[Name:cg-k8s-1.18.2]",
			filter:   "Name eq 'cg-k8s-1.18.2'",
			datatype: "",
			ok:       true,
		},
		{
			moref:    "59c84e4a16267c0001c23428",
			filter:   "",
			datatype: "",
			ok:       false,
		},
	}

	for _, test := range tests {
		filter, datatype, ok := ParseMoRef(test.moref)
		assert.Equal(t, test.filter, filter)
		assert.Equal(t, test.datatype, datatype)
		assert.Equal(t, test.ok, ok)
	}
}

func TestGetMoid(t *testing.T) {
	var moid string
	var ok bool

	moid, ok = GetMoid(map[string]interface{}{
		"ObjectType": "ntp.Policy.List",
		"Results": []any{
			map[string]any{
				"Moid": "12345",
			},
		},
	})
	assert.Equal(t, true, ok)
	assert.Equal(t, "12345", moid)
}
