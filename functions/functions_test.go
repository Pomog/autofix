package functions

import (
	"testing"
)

func TestReplaceHexWithDecimal(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "1E (hex) files were added",
			expected: "30 files were added",
		},
		{
			input:    "A (hex)",
			expected: "10",
		},
		{
			input:    " 1F  (hex)   ",
			expected: " 31   ",
		},
		{
			input:    "Invalid (hex)",
			expected: "Invalid (hex)",
		},
		{
			input:    "Invalid  (hex)",
			expected: "Invalid  (hex)",
		},
		{
			input:    "ABC (hex) DEF (hex)(hex)",
			expected: "2748 3567(hex)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ReplaceHexWithDecimal(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %s, Got: %s", tc.expected, result)
			}
		})
	}
}
