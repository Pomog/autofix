package functions

import (
	"fmt"
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

func TestReplaceBinWithDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"101 (bin)", "5"},
		{"1101 (bin)", "13"},
		{" 0 (bin)", " 0"},
		{"11111111  (bin)", "255"},
		{"Hello (bin)", "Hello (bin)"},
		{"It has been 10 (bin) years", "It has been 2 years"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %s", test.input), func(t *testing.T) {
			output := ReplaceBinWithDecimal(test.input)
			if output != test.expected {
				t.Errorf("Expected: %s, Got: %s", test.expected, output)
			}
		})
	}
}
