package functions

import (
	"fmt"
	"regexp"
	"strings"
)

type CustomLogicFunc func(reRep *regexp.Regexp, input string, parsedNumber int) string
type StringModificationFunc func(input string) string

/*
implementation of the fixingWithNumber function
converts the previously specified <number> of words followed by "(cap, <number>)"
with the capitalized version of it
*/
func CapitalizationWithNumber(input string) string {
	flag := "cap"
	input = fixingWithNumber(input, flag, fixText(capitalization))
	return input
}

/*
implementation of the fixingWithNumber function
converts the previously specified <number> of words followed by "(up, <number>)"
with the UpperCase version of it
*/
func ToUppercaseWithNumber(input string) string {
	flag := "up"
	input = fixingWithNumber(input, flag, fixText(strings.ToUpper))
	return input
}

/*
converts the previously specified <number> of words followed by "(flag, <number>)"
with the version of it obtained by applying the customLogic function
*/
func fixingWithNumber(input, flag string, customLogic CustomLogicFunc) string {
	re := regexp.MustCompile(fmt.Sprintf(`\b[\w]+\s*\(%s(?:,\s*(\d+))?\)`, flag))
	return processPatternMatches(input, re, flag, customLogic)
}

/*
processes the input string to find and replace patterns
based on the provided regular expression and custom logic function.

It iterates through the input string, detects pattern matches using the
regular expression 're', applies the custom logic function to generate fixed
replacements, and updates the input string accordingly. The process continues
until no more pattern matches are found.

The modified input string is returned after processing.
*/
func processPatternMatches(input string, re *regexp.Regexp, flag string, customLogic CustomLogicFunc) string {
	for foundPatternMatch(re, input) {
		parsedNumber := parseInt(re.FindStringSubmatch(input)[1])
		reRep := createRepeatedPatternRegexp(parsedNumber, flag)

		fixedString := customLogic(reRep, input, parsedNumber)

		input = replaceStrings(reRep, input, parsedNumber, fixedString)
	}
	return input
}

/*
returns the compiled regular expression for repeated patterns
*/
func createRepeatedPatternRegexp(parsedNumber int, flag string) *regexp.Regexp {
	repeatedPattern := strings.Repeat(`(\b[\w]+\s*)`, parsedNumber) + `\(` + flag + `(?:,\s*(\d+))?\)`
	reRep := regexp.MustCompile(repeatedPattern)
	return reRep
}

/*
replace the pattern from the input String by fixedString String
*/
func replaceStrings(reRep *regexp.Regexp, input string, parsedNumber int, fixedString string) string {
	matchedString := reRep.FindStringSubmatch(input)[0]
	if parsedNumber != 0 {
		input = strings.Replace(input, matchedString, strings.TrimSpace(fixedString), 1)
	} else {
		input = strings.TrimSpace(strings.Replace(input, matchedString, "", 1))
	}
	return input
}

/*
converts a String using the provided modification function
*/
func getModifiedString(modificationFunc StringModificationFunc, reRep *regexp.Regexp, input string, parsedNumber int) string {
	var fixedString string
	for _, result := range reRep.FindStringSubmatch(input)[1 : parsedNumber+1] {
		fixedString = fixedString + modificationFunc(result)
	}
	return fixedString
}

/*
converts a word to its capitalized equivalent
*/
func capitalization(input string) string {
	return strings.Title(input)
}

/*
Checks if the input string matches the regular expression 're'.
*/
func foundPatternMatch(re *regexp.Regexp, input string) bool {
	return re.MatchString(input)
}

/*
If the input string is not a valid integer, it returns 0.
*/
func parseInt(s string) int {
	var num int
	if s == "" {
		return 0
	}
	_, err := fmt.Sscanf(s, "%d", &num)
	if err != nil {
		return 0
	}
	return num
}

/*
Applies the provided modification function to the input text
*/
func fixText(modificationFunc StringModificationFunc) CustomLogicFunc {
	return func(reRep *regexp.Regexp, input string, parsedNumber int) string {
		return getModifiedString(modificationFunc, reRep, input, parsedNumber)
	}
}
