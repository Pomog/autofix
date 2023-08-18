package functions

import (
	"os"
	"testing"
)

func runTest(t *testing.T, testFunction func(string) string, tests []struct{ input, expected string }) {
	for _, test := range tests {
		result := testFunction(test.input)
		if result != test.expected {
			t.Errorf("For input \n'%s',\n expected '%s',\n but got '%s'", test.input, test.expected, result)
		}
	}
}

func TestReplaceHexWithDecimal(t *testing.T) {
	tests := []struct {
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
	runTest(t, ReplaceHexWithDecimal, tests)
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
	runTest(t, ReplaceBinWithDecimal, tests)
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
	runTest(t, ToUppercase, tests)
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
	runTest(t, ToLowercase, tests)
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
	runTest(t, Capitalization, tests)
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
	runTest(t, CapitalizationWithNumber, tests)
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
	runTest(t, ToUppercaseWithNumber, tests)
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
		// Only if there is at least one word before preceded a FLAG it matches regex pattern
		{"(low, 0) No words should be converted (low, 0).", "No words should be converted ."},
		{"(low, 0) No words should be converted.", "(low, 0) No words should be converted."},
		{"(low, 1) No words Should Be Converted (low, 2).", "(low, 1) No words Should be converted."},
		{"(low, 1) (low, 2) No words Should Be Converted (low, 2).", "(low, 1) (low, 2) No words Should be converted."},
	}
	runTest(t, ToLowercaseWithNumber, tests)
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
	runTest(t, CorrectPunctuationsSpaces, tests)
}

func TestCorrectApostrophesSpaces(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"I am exactly how they describe me: ' awesome person '", "I am exactly how they describe me: 'awesome person'"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
	}
	runTest(t, CorrectApostrophesSpaces, tests)
}

func TestArticlesCorrection(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"There it was. A amazing rock, an amazing! It was a amazing day. An amazing day.", "There it was. An amazing rock, an amazing! It was an amazing day. An amazing day."},
		{"A dog barked at a man. The man had an umbrella.", "A dog barked at a man. The man had an umbrella."},
		{"She saw a elephant and an mouse at the zoo.", "She saw an elephant and an mouse at the zoo."},
		{"I have a university degree and an MBA.", "I have an university degree and an MBA."},
		{"A unicorn is a mythical creature. An unicorn has a horn on its forehead.", "An unicorn is a mythical creature. An unicorn has an horn on its forehead."},
	}
	runTest(t, ArticlesCorrection, tests)
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

func TestReadFromFile(t *testing.T) {
	// Create a temporary file for testing
	content := "Line 1\nLine 2\nLine 3"
	tmpFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	lines, err := ReadFromFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("ReadFromFile returned an error: %v", err)
	}

	expectedLines := []string{"Line 1", "Line 2", "Line 3"}
	if len(lines) != len(expectedLines) {
		t.Fatalf("Expected %d lines, but got %d lines", len(expectedLines), len(lines))
	}

	for i, line := range lines {
		if line != expectedLines[i] {
			t.Errorf("Line %d: Expected '%s', but got '%s'", i+1, expectedLines[i], line)
		}
	}
}
