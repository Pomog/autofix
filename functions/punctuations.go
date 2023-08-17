package functions

import (
	"fmt"
	"regexp"
	"strings"
)

/*
CorrectPunctuationsSpaces takes an input string and adds spaces around punctuation marks
based on the specified rules duscribed by inner function applyPunctuationRules.
Punctuation marks are placed close to the previous word
and separated by a space from the following content. Exceptions include groups of
punctuation like '...' or '!?'.
*/
func CorrectPunctuationsSpaces(input string) string {
	punctuation := []string{".", ",", "!", "?", ":", ";"} // TODO: add configurable punctuation, by config file or env var
	punctuationPattern := strings.Join(punctuation, "|\\")
	strPattern := `(\b[\w]+)\s*([\` + punctuationPattern + `]+)\s*` // shoud be two groups, where the first element is the word, and the second is the punctuation mark

	re := regexp.MustCompile(strPattern)

	return correctString(re, input)
}

/*
applies the punctuation spacing rules to the input string.
It uses a regular expression to match and replace substrings based on the rules defined in applyPunctuationRules.
*/
func correctString(re *regexp.Regexp, input string) string {
	return re.ReplaceAllStringFunc(input, func(match string) string {
		return applyPunctuationRules(re, match, input)
	})
}

/*
applies the punctuation spacing rules to the matched substring.
It takes the matched substring, extracts the word and punctuation mark preseeded by it,
and places them together with or without following space based on whether it's the end of the string.
*/
func applyPunctuationRules(re *regexp.Regexp, match string, input string) string {
	submatches := re.FindStringSubmatch(match)
	if strings.HasSuffix(input, match) {
		return fmt.Sprintf("%s%s", submatches[1], submatches[2]) // end of the string
	}
	return fmt.Sprintf("%s%s ", submatches[1], submatches[2]) // not the end of the string
}

/*
The punctuation mark ' will always be found with another instance of it
and they should be placed to the right and left of the word(s) in the middle of them,
without any spaces.
*/
func CorrectApostrophesSpaces(input string) string {
	symbolsInsideApostrophes := []string{"\\w", "-"} // TODO: add configurable punctuation, by config file or env var
	symbolsPattern := strings.Join(symbolsInsideApostrophes, "")
	strPattern := `'\s*([` + symbolsPattern + `]+(\s*[` + symbolsPattern + `]+)*)\s*'`
	re := regexp.MustCompile(strPattern)

	return re.ReplaceAllString(input, "'$1'")
}
