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
		{"This is so exciting (cap, 1)", "This is so Exciting"},
		{"This is so exciting (cap, -1)", "This is so exciting (cap, -1)"},
		{"This is so exciting (cap, 0).", "This is so exciting ."},
		{"This is so exciting", "This is so exciting"},
		{"This is (cap, 1) so exciting (cap, 1)", "This Is so Exciting"},
		{"this is (cap, 2) so exciting (cap, 1)", "This Is so Exciting"},
		{"This is so exciting (cap, 2)", "This is So Exciting"},
		{"This is so exciting (cap, 3)", "This Is So Exciting"},
		{"This is so exciting (cap, 4)", "This Is So Exciting"},
		{"This is so (cap, 1) exciting", "This is So exciting"},
		{"This is so  (cap, 2) exciting", "This Is So exciting"},
		{"But I'm sure it will be (cap, 3) exciting.", "But I'm sure It Will Be exciting."},
	}

	for _, test := range tests {
		output := CapitalizationWithNumber(test.input)
		if output != test.expected {
			t.Errorf("Input: %s\nExpected: %s\nGot: %s", test.input, test.expected, output)
		}
	}
}

func TestToUppercaseWithNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"This is so exciting (up, 1).", "This is so EXCITING."},
		{"I can't wait to see  what happens next (up, 3).", "I can't wait to see  WHAT HAPPENS NEXT."},
		{"This is so exciting (up, 0).", "This is so exciting ."},
		{"This is so exciting", "This is so exciting"},
		{"This is (up, 1) so exciting (up, 1)", "This IS so EXCITING"},
		{"But I'm sure it will be (up, 3) exciting.", "But I'm sure IT WILL BE exciting."},
		{"This is so (up, 1) exciting", "This is SO exciting"},
		{"This is so  (up, 2) exciting", "This IS SO exciting"},
		{"But I'm sure it will be (up, 3) exciting.", "But I'm sure IT WILL BE exciting."},
		{"This is so EXCITINg (up, 1).", "This is so EXCITING."},
	}

	for _, test := range tests {
		output := ToUppercaseWithNumber(test.input)
		if output != test.expected {
			t.Errorf("Input: %s\nExpected: %s\nGot: %s", test.input, test.expected, output)
		}
	}
}

func TestToLowercaseWithNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"This is so EXCITING (low, 1)", "This is so exciting"},
		{"I CAN'T WAIT TO SEE  WHAT HAPPENS NEXT (low, 3).", "I CAN'T WAIT TO SEE  what happens next."},
		{"This is so EXCITINg (low, 0).", "This is so EXCITINg ."},
		{"This is so EXCITINg (low, 1).", "This is so exciting."},
		{"Testing (low, 1) uppercase (low, 2) lowercase (low, 3) transformation", "testing uppercase lowercase transformation"},
		{"(low, 1) ALL CAPS SENTENCE (low, 2) with (low, 3) multiple (low, 4) words", "(low, 1) ALL caps sentence with multiple words"},
		{"(low, 10) This sentence has fewer words than specified", "(low, 10) This sentence has fewer words than specified"},
		{"(low, 1) (low, 2) (low, 3) Multiple conversions in a row", "(low, 1) (low, 2) (low, 3) Multiple conversions in a row"},

		// These are the most incomprehensible cases, as the assignment defined in fuzzy way. The primary function for these processes is replaceStrings().
		{"(low, 0) No words should be converted (low, 0).", "No words should be converted ."},
		{"(low, 0) No words should be converted.", "(low, 0) No words should be converted."},
		{"(low, 1) No words Should Be Converted (low, 2).", "(low, 1) No words Should be converted."},
		{"(low, 1) (low, 2) No words Should Be Converted (low, 2).", "(low, 1) (low, 2) No words Should be converted."},
	}

	for _, test := range tests {
		output := ToLowercaseWithNumber(test.input)
		if output != test.expected {
			t.Errorf("Input: %s\nExpected: %s\nGot: %s", test.input, test.expected, output)
		}
	}
}

func TestCorrectPunctuationsSpaces(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"Hello,world.How are you?I am good.", "Hello, world. How are you? I am good."},
		{"No spaces here:", "No spaces here:"},
		{"Punctuation at!the end:", "Punctuation at! the end:"},
		{"This is a test: testing, testing!", "This is a test: testing, testing!"},
		{"Hello...world.", "Hello... world."},
		{"Spaces  should be corrected .", "Spaces  should be corrected."},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := CorrectPunctuationsSpaces(test.input)
			if result != test.expected {
				t.Errorf("Expected: %s\nGot: %s", test.expected, result)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{"-456", -456},
		{"0", 0},
		{"abc", 0}, // Invalid input
		{"", 0},    // Empty input
	}

	for _, test := range tests {
		result := parseInt(test.input)
		if result != test.expected {
			t.Errorf("Expected: %d, Got: %d", test.expected, result)
		}

	}
}
