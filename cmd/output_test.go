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

	var expected interface{}

	//Test with empty filter
	expected = map[string]interface{}{
		"ClassId":    "ntp.Policy",
		"ObjectType": "ntp.Policy",
		"Moid":       "1234567",
		"Name":       "NtpTest",
		"Enabled":    true,
		"NtpServers": []interface{}{"1.1.1.1", "1.1.1.2"},
	}

	out, err := applyJSONPathFilter(in, "")
	assert.NoError(t, err)

	assert.Equal(t, expected, out)

	//Test with common filter

	expected = "NtpTest"

	out, err = applyJSONPathFilter(in, "$.Name")
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

type simplifyResultTestCase struct {
	In  interface{}
	Out interface{}
}

func TestRemoveWrapper(t *testing.T) {
	testCases := []simplifyResultTestCase{

		// Multiple results should return a slice of map[string]
		{
			In: map[string]interface{}{
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
			},
			Out: []interface{}{
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

		// Single result should return just a map[string]
		{
			In: map[string]interface{}{
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
			},
			Out: map[string]interface{}{
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
	}

	for _, testCase := range testCases {
		out := removeWrappers(testCase.In)

		assert.Equal(t, testCase.Out, out)
	}
}

func TestFilterAttributes(t *testing.T) {
	testCases := []simplifyResultTestCase{
		{
			In: []interface{}{
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
			Out: []interface{}{
				map[string]interface{}{
					"Moid":       "1234567",
					"Name":       "NtpTest",
					"Enabled":    true,
					"NtpServers": []interface{}{"1.1.1.1", "1.1.1.2"},
					"TestRef":    "MoRef[TestRefType/7654321]",
				},
			},
		},
	}

	for _, testCase := range testCases {
		out := filterAttributes(testCase.In)

		assert.Equal(t, testCase.Out, out)
	}
}

func TestPrepareResultTable(t *testing.T) {
	var in interface{}
	var expectedData, outData [][]string
	var expectedHeaders, outHeaders []string

	// Test normal scenario - result is a slice of maps
	in = []interface{}{
		map[string]interface{}{
			"Moid":       "1234567",
			"Name":       "NtpTest",
			"Enabled":    true,
			"NtpServers": []interface{}{"1.1.1.1", "1.1.1.2"},
		},
	}

	expectedData = [][]string{
		{"1234567", "NtpTest", "True", "1.1.1.1, 1.1.1.2"},
	}
	expectedHeaders = []string{"Moid", "Name", "Enabled", "NtpServers"}

	outData, outHeaders = prepareResultTable(in)

	assert.ElementsMatch(t, expectedHeaders, outHeaders)
	for i := range expectedData {
		assert.ElementsMatch(t, expectedData[i], outData[i])
	}

	// Test JSON path result scenario - result is a slice of strings
	in = []interface{}{
		"123",
		"456",
	}
	expectedData = [][]string{
		{"123"},
		{"456"},
	}
	expectedHeaders = []string{}
	outData, outHeaders = prepareResultTable(in)

	assert.ElementsMatch(t, expectedHeaders, outHeaders)
	for i := range expectedData {
		assert.ElementsMatch(t, expectedData[i], outData[i])
	}

	// Test JSON path result scenario - result is a string
	in = "123"
	expectedData = [][]string{
		{"123"},
	}
	expectedHeaders = []string{}
	outData, outHeaders = prepareResultTable(in)

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
