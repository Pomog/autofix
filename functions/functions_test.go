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
		{"1E (hex) files were added", "30 files were added"},
		{"A(hex)", "10"},
		{"A (hex)", "10"},
		{" 1F  (hex)   ", " 31   "},
		{"Invalid (hex)", "Invalid (hex)"},
		{"Invalid  (hex)", "Invalid  (hex)"},
		{"ABC (hex) DEF (hex)(hex)", "2748 3567(hex)"},
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
		{"Invalid (bin)", "Invalid (bin)"},
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

func TestToUpper(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello (up)", "HELLO"},
		{"he1llo (up)", "HE1LLO"},
		{"world  (up) ", "WORLD "},
		{"@world (up)", "@WORLD"},
		{"hello@world (up)", "hello@WORLD"},
		{"hello1world (up)", "HELLO1WORLD"},
		{"go (up) programming", "GO programming"},
		{"Ready, set, go (up) !", "Ready, set, GO !"},
		{"Ready, set, Go (up) !", "Ready, set, GO !"},
	}

	for _, test := range tests {
		output := ToUppercase(test.input)
		if output != test.expected {
			t.Errorf("Input: %s\nExpected: %s\nGot: %s", test.input, test.expected, output)
		}
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HELLO (low)", "hello"},
		{"HE1LLO (low)", "he1llo"},
		{"WORLD  (low) ", "world "},
		{"@WORLD (low)", "@world"},
		{"hello@WORLD (low)", "hello@world"},
		{"HELLO1WORLD (low)", "hello1world"},
		{"GO (low) programming", "go programming"},
		{"Ready, set, GO (low) !", "Ready, set, go !"},
		{"Ready, set, Go (low) !", "Ready, set, go !"},
	}

	for _, test := range tests {
		output := ToLowercase(test.input)
		if output != test.expected {
			t.Errorf("Input: %s\nExpected: %s\nGot: %s", test.input, test.expected, output)
		}
	}
}

func TestCapitalization(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello (cap)", "Hello"},
		{"hello", "hello"},
		{"Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"wORLD  (cap) ", "WORLD "},
		{"@go (cap)", "@Go"},
		{"go@lang (cap)", "go@Lang"},
		{"hello1world (cap)", "Hello1world"},
		{"Ready, set, go (cap) !", "Ready, set, Go !"},
	}

	for _, test := range tests {
		output := Capitalization(test.input)
		if output != test.expected {
			t.Errorf("Input: %s\nExpected: %s\nGot: %s", test.input, test.expected, output)
		}
	}
}

func TestCapitalizationWithNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"This is so exciting (up, 1)", "This is so Exciting"},
		{"This is so exciting (up, 0).", "This is so exciting ."},
		{"This is so exciting", "This is so exciting"},
		{"This is (up, 1) so exciting (up, 1)", "This Is so Exciting"},
		{"this is (up, 2) so exciting (up, 1)", "This Is so Exciting"},
		{"This is so exciting (up, 2)", "This is So Exciting"},
		{"This is so exciting (up, 3)", "This Is So Exciting"},
		{"This is so exciting (up, 4)", "This Is So Exciting"},
		{"This is so (up, 1) exciting", "This is So exciting"},
		{"This is so  (up, 2) exciting", "This Is So exciting"},
		{"But I'm sure it will be (up, 3) exciting.", "But I'm sure It Will Be exciting."},
	}

	for _, test := range tests {
		output := CapitalizationWithNumber(test.input)
		if output != test.expected {
			t.Errorf("Input: %s\nExpected: %s\nGot: %s", test.input, test.expected, output)
		}
	}
}
