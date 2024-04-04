package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceArgs(t *testing.T) {
	var s string
	var e error
	s, e = ReplaceArgs("api/{Moid}", []string{"12345"})
	assert.NoError(t, e)
	assert.Equal(t, "api/12345", s)
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

// func TestExpandTemplate(t *testing.T) {
// 	tests := []struct {
// 		in       map[string]any
// 		expected map[string]any
// 	}{
// 		{
// 			in: map[string]any{
// 				"Organization": "{{ MoRef }}",
// 			},
// 			expected: map[string]any{
// 				"Organization": map[string]any{
// 					"ClassId":    "mo.MoRef",
// 					"Moid":       "NewMoRef",
// 					"ObjectType": "organization.Organization",
// 				},
// 			},
// 		},
// 		{
// 			in: map[string]any{
// 				"Ancestors": []string{
// 					"{{ MoRef}}",
// 					"{{MoRef }}",
// 					"{{MoRef}}",
// 				},
// 			},
// 			expected: map[string]any{
// 				"Ancestors": []any{
// 					map[string]any{
// 						"ClassId":    "mo.MoRef",
// 						"Moid":       "NewMoRef",
// 						"ObjectType": "organization.Organization",
// 					},
// 					map[string]any{
// 						"ClassId":    "mo.MoRef",
// 						"Moid":       "NewMoRef",
// 						"ObjectType": "organization.Organization",
// 					},
// 					map[string]any{
// 						"ClassId":    "mo.MoRef",
// 						"Moid":       "NewMoRef",
// 						"ObjectType": "organization.Organization",
// 					},
// 				},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		ExpandTemplate(nil, &test.in)
// 		assert.Equal(t, test.expected, test.in)
// 	}
// }
