package main

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyJSONPathFilter(t *testing.T) {
	var in map[string]any

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
		{"1234567", "NtpTest", "true", "1.1.1.1, 1.1.1.2"},
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
		{"NtpTest", "1234567", "true"},
	}

	outData, outHeaders = prepareResultTableCustomColumns(in, "NAME:.Name,MOID:.Moid,ENABLED:.Enabled")

	assert.Equal(t, expectedHeaders, outHeaders)
	for i := range expectedData {
		assert.Equal(t, expectedData[i], outData[i])
	}

	outData, outHeaders = prepareResultTableCustomColumns(in, "")
	expectedHeaders = []string{"Name", "Moid", "Enabled", "NtpServers"}
	expectedData = [][]string{{"NtpTest", "1234567", "true", "1.1.1.1, 1.1.1.2"}}

	assert.Equal(t, expectedHeaders, outHeaders)
	for i := range expectedData {
		assert.Equal(t, expectedData[i], outData[i])
	}

	// Test empty list with no template — must not panic
	outData, outHeaders = prepareResultTableCustomColumns([]interface{}{}, "")
	assert.Empty(t, outHeaders)
	assert.Empty(t, outData)
}

func TestRelaxedJSONPathExpression(t *testing.T) {
	assert.Equal(t, "$.Name", relaxedJSONPathExpression("$.Name"))
	assert.Equal(t, "$.Name", relaxedJSONPathExpression(".Name"))
	assert.Equal(t, "$.Name[0].Tags", relaxedJSONPathExpression("$.Name[0].Tags"))
	assert.Equal(t, "$[*].Tags", relaxedJSONPathExpression("[*].Tags"))
}

func TestSortHeaders(t *testing.T) {
	tests := []struct {
		in            []string
		priorityNames []string
		out           []string
	}{
		{
			in:            []string{"Test", "ATest", "Name", "Moid", "Compute_Serial"},
			priorityNames: []string{"Compute_Serial", "Name", "Moid"},
			out:           []string{"Compute_Serial", "Name", "Moid", "ATest", "Test"},
		},
		{
			in:            []string{"Test", "ATest", "Name", "Moid"},
			priorityNames: []string{"Compute_Serial", "Name", "Moid"},
			out:           []string{"Name", "Moid", "ATest", "Test"},
		},
	}

	for _, test := range tests {
		in := test.in[0:]
		SortHeaders(in, test.priorityNames)
		assert.Equal(t, test.out, in)
	}
}

func TestPrintResultGoTemplate(t *testing.T) {
	// Helper to capture stdout
	captureOutput := func(f func()) string {
		// Save original stdout
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		f()

		w.Close()
		os.Stdout = oldStdout

		var buf bytes.Buffer
		buf.ReadFrom(r)
		return buf.String()
	}

	tests := []struct {
		name     string
		result   interface{}
		template string
		expected string
	}{
		{
			name: "simple field access",
			result: map[string]interface{}{
				"Name":    "TestPolicy",
				"Enabled": true,
				"Moid":    "1234567",
			},
			template: "{{.Name}}",
			expected: "TestPolicy\n",
		},
		{
			name: "multiple fields",
			result: map[string]interface{}{
				"Name":    "TestPolicy",
				"Enabled": true,
			},
			template: "Name: {{.Name}}, Enabled: {{.Enabled}}",
			expected: "Name: TestPolicy, Enabled: true\n",
		},
		{
			name: "range over slice",
			result: []interface{}{
				map[string]interface{}{"Name": "Policy1"},
				map[string]interface{}{"Name": "Policy2"},
			},
			template: "{{range .}}{{.Name}}\n{{end}}",
			expected: "Policy1\nPolicy2\n\n",
		},
		{
			name: "sprig upper function",
			result: map[string]interface{}{
				"Name": "TestPolicy",
			},
			template: "{{.Name | upper}}",
			expected: "TESTPOLICY\n",
		},
		{
			name: "sprig trim function",
			result: map[string]interface{}{
				"Name": "  TestPolicy  ",
			},
			template: "{{.Name | trim}}",
			expected: "TestPolicy\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				printResultGoTemplate(tt.result, tt.template)
			})
			assert.Equal(t, tt.expected, output)
		})
	}
}

func TestPrintResultGoTemplateWithConditions(t *testing.T) {
	result := map[string]interface{}{
		"Name":    "TestPolicy",
		"Enabled": true,
	}

	// Helper to capture stdout
	captureOutput := func(f func()) string {
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		f()

		w.Close()
		os.Stdout = oldStdout

		var buf bytes.Buffer
		buf.ReadFrom(r)
		return buf.String()
	}

	output := captureOutput(func() {
		printResultGoTemplate(result, "{{if .Enabled}}ENABLED{{else}}DISABLED{{end}}")
	})

	assert.True(t, strings.Contains(output, "ENABLED"))
}

func TestPrintCSVLine(t *testing.T) {
	captureOutput := func(f func()) string {
		var buf bytes.Buffer
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		f()

		w.Close()
		os.Stdout = old
		buf.ReadFrom(r)
		return buf.String()
	}

	t.Run("plain values", func(t *testing.T) {
		output := captureOutput(func() {
			printCSVLine([]string{"hello", "world"})
		})
		assert.Equal(t, "\"hello\",\"world\"\n", output)
	})

	t.Run("escapes embedded double quotes", func(t *testing.T) {
		output := captureOutput(func() {
			printCSVLine([]string{`say "hello"`})
		})
		assert.Equal(t, `"say ""hello"""`+"\n", output)
	})
}
