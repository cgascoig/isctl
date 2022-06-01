package main

import (
	"encoding/json"
	"testing"

	openapi "github.com/cgascoig/intersight-go-sdk/intersight"
	"github.com/stretchr/testify/assert"
)

func TestApplyJSONPathFilter(t *testing.T) {
	var in openapi.NtpPolicyResponse

	inStr := `
	{
		"ObjectType": "ntp.Policy.List",
		"Results": [
		  {
			"ClassId": "ntp.Policy",
			"Enabled": true,
			"Moid": "1234567",
			"Name": "NtpTest",
			"NtpServers": [
			  "1.1.1.1",
			  "1.1.1.2"
			],
			"ObjectType": "ntp.Policy"
		  }
		]
	}
	`

	err := json.Unmarshal([]byte(inStr), &in)
	assert.NoError(t, err)

	var expected interface{}

	//Test with empty filter
	expectedSingleResult := map[string]interface{}{
		"ClassId":    "ntp.Policy",
		"ObjectType": "ntp.Policy",
		"Moid":       "1234567",
		"Name":       "NtpTest",
		"Enabled":    true,
		"NtpServers": []interface{}{"1.1.1.1", "1.1.1.2"},
	}
	expected = []interface{}{
		expectedSingleResult,
	}

	out, err := applyJSONPathFilter(in, "", false)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)

	out, err = applyJSONPathFilter(in, "", true)
	assert.NoError(t, err)
	assert.Equal(t, expectedSingleResult, out)

	//Test with common filter
	expected = []interface{}{"NtpTest"}

	out, err = applyJSONPathFilter(in, "$[*].Name", false)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)

	out, err = applyJSONPathFilter(in, "$.Name", true)
	assert.NoError(t, err)
	assert.Equal(t, "NtpTest", out)
}

type simplifyResultTestCase struct {
	In           interface{}
	Out          interface{}
	SingleResult bool
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

		// Single result should return just a slice of map[string] when singleResult=false
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
			},
			SingleResult: false,
		},

		// Single result should return just a map[string] when singleResult=true
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
			SingleResult: true,
		},
	}

	for _, testCase := range testCases {
		out := removeWrappers(testCase.In, testCase.SingleResult)

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
	expectedHeaders = []string{"Name", "Moid", "Enabled", "NtpServers"}

	outData, outHeaders = prepareResultTable(in, true)

	assert.Equal(t, expectedHeaders, outHeaders)
	// assert.ElementsMatch(t, expectedHeaders, outHeaders)
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
	outData, outHeaders = prepareResultTable(in, true)

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
	outData, outHeaders = prepareResultTable(in, true)

	assert.ElementsMatch(t, expectedHeaders, outHeaders)
	for i := range expectedData {
		assert.ElementsMatch(t, expectedData[i], outData[i])
	}
}

func TestPrepareResultTableCustomColums(t *testing.T) {
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

	expectedHeaders = []string{"NAME", "MOID", "ENABLED"}
	expectedData = [][]string{
		{"NtpTest", "1234567", "True"},
	}

	outData, outHeaders = prepareResultTableCustomColumns(in, "NAME:.Name,MOID:.Moid,ENABLED:.Enabled")

	assert.Equal(t, expectedHeaders, outHeaders)
	for i := range expectedData {
		assert.Equal(t, expectedData[i], outData[i])
	}
}

func TestRelaxedJSONPathExpression(t *testing.T) {
	assert.Equal(t, "$.Name", relaxedJSONPathExpression("$.Name"))
	assert.Equal(t, "$.Name", relaxedJSONPathExpression(".Name"))
	assert.Equal(t, "$.Name[0].Tags", relaxedJSONPathExpression("$.Name[0].Tags"))
}

func optionalStr(s string) *string {
	return &s
}

func optionalBool(b bool) *bool {
	return &b
}
