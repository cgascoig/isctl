package main

import (
	"testing"

	"github.com/cgascoig/isctl/openapi"
	"github.com/stretchr/testify/assert"
)

func TestApplyJSONPathFilter(t *testing.T) {
	in := openapi.NtpPolicyResponse{
		NtpPolicyList: &openapi.NtpPolicyList{
			MoBaseResponse: openapi.MoBaseResponse{
				ObjectType: "ntp.Policy.List",
			},
			Results: &[]openapi.NtpPolicy{
				{
					PolicyAbstractPolicy: openapi.PolicyAbstractPolicy{
						MoBaseMo: openapi.MoBaseMo{
							Moid:       optionalStr("1234567"),
							ClassId:    "ntp.Policy",
							ObjectType: "ntp.Policy",
						},
						Name: optionalStr("NtpTest"),
					},
					Enabled:    optionalBool(true),
					NtpServers: &[]string{"1.1.1.1", "1.1.1.2"},
				},
			},
		},
	}

	//Test with empty filter

	var expected interface{}
	expected = map[string]interface{}{
		"MoDocumentCount": nil,
		"NtpPolicyList": map[string]interface{}{
			"ObjectType": "ntp.Policy.List",
			"Results": []interface{}{
				map[string]interface{}{
					"ClassId":    "ntp.Policy",
					"ObjectType": "ntp.Policy",
					"Moid":       "1234567",
					"Name":       "NtpTest",
					"Enabled":    true,
					"NtpServers": []interface{}{"1.1.1.1", "1.1.1.2"},
				},
			},
		},
	}

	out, err := applyJSONPathFilter(in, "")
	assert.NoError(t, err)

	assert.Equal(t, expected, out)

	//Test with common filter

	expected = []interface{}{
		"NtpTest",
	}

	out, err = applyJSONPathFilter(in, "$.NtpPolicyList.Results[*].Name")
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestSimplifyResult(t *testing.T) {
	var in interface{}

	in = map[string]interface{}{
		"MoDocumentCount": nil,
		"NtpPolicyList": map[string]interface{}{
			"ObjectType": "ntp.Policy.List",
			"Results": []interface{}{
				map[string]interface{}{
					"ClassId":    "ntp.Policy",
					"ObjectType": "ntp.Policy",
					"Moid":       "1234567",
					"Name":       "NtpTest",
					"Enabled":    true,
					"NtpServers": []interface{}{"1.1.1.1", "1.1.1.2"},
					"TestRef": map[string]interface{}{
						"ClassId":    "mo.MoRef",
						"ObjectType": "TestRefType",
						"Moid":       "7654321",
					},
				},
			},
		},
	}

	expected := []interface{}{
		map[string]interface{}{
			"Moid":       "1234567",
			"Name":       "NtpTest",
			"Enabled":    true,
			"NtpServers": []interface{}{"1.1.1.1", "1.1.1.2"},
			"TestRef":    "MoRef[TestRefType/7654321]",
		},
	}

	out := simplifyResult(in)

	assert.Equal(t, expected, out)
}

func TestPrepareResultTable(t *testing.T) {
	var in interface{}

	in = []interface{}{
		map[string]interface{}{
			"Moid":       "1234567",
			"Name":       "NtpTest",
			"Enabled":    true,
			"NtpServers": []interface{}{"1.1.1.1", "1.1.1.2"},
		},
	}

	expectedData := [][]string{
		{"1234567", "NtpTest", "True", "1.1.1.1, 1.1.1.2"},
	}
	expectedHeaders := []string{"Moid", "Name", "Enabled", "NtpServers"}

	outData, outHeaders := prepareResultTable(in)

	assert.ElementsMatch(t, expectedHeaders, outHeaders)
	for i := range expectedData {
		assert.ElementsMatch(t, expectedData[i], outData[i])
	}
}

func optionalStr(s string) *string {
	return &s
}

func optionalBool(b bool) *bool {
	return &b
}
