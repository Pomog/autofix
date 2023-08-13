package functions

import (
	"fmt"
	"regexp"
	"strings"
)

type CustomLogicFunc func(reRep *regexp.Regexp, input string, parsedNumber int) string

/*
converts the previously specified <number> of words followed by "(cap, <number>)"
with the capitalizex version of it
*/
func CapitalizationWithNumber(input string) string {
	flag := "cap"
	input = fixingWithNumber(input, flag, getCapitalizedString)
	return input
}

/*
converts the previously specified <number> of words followed by "(flag, <number>)"
with the version of it obtained by applying the customLogic function
*/
func fixingWithNumber(input, flag string, customLogic CustomLogicFunc) string {
	re := regexp.MustCompile(fmt.Sprintf(`\b[\w]+\s*\(%s(?:,\s*(\d+))?\)`, flag))
	for noMatchCheck(re, input) {
		parsedNumber := parseInt(re.FindStringSubmatch(input)[1])
		repeatedPattern := strings.Repeat(`(\b[\w]+\s*)`, parsedNumber) + `\(` + flag + `(?:,\s*(\d+))?\)`
		reRep := regexp.MustCompile(repeatedPattern)

		fixedString := customLogic(reRep, input, parsedNumber)

		input = replaceStrings(reRep, input, parsedNumber, fixedString)
	}
	return input
}

/*
replace the pattern form the input String by fixedString String
*/
func replaceStrings(reRep *regexp.Regexp, input string, parsedNumber int, fixedString string) string {
	machedString := reRep.FindStringSubmatch(input)[0]
	if parsedNumber != 0 {
		input = strings.Replace(input, machedString, strings.TrimSpace(fixedString), 1)
	} else {
		input = strings.TrimSpace(strings.Replace(input, machedString, "", 1))
	}
	return input
}

/*
converts a String to its capitalized equivalent
*/
func getCapitalizedString(reRep *regexp.Regexp, input string, parsedNumber int) string {
	var fixedString string
	for _, result := range reRep.FindStringSubmatch(input)[1 : parsedNumber+1] {
		fixedString = fixedString + getCapitalizedWord(result)
	}
	return fixedString
}

/*
converts a word to its capitalized equivalent
*/
func getCapitalizedWord(input string) string {
	return strings.Title(input)
}

/*
No match found, return input unchanged
*/
func noMatchCheck(re *regexp.Regexp, input string) bool {
	var isMatchFound bool
	if re.FindStringSubmatch(input) != nil {
		isMatchFound = true
	}
	return isMatchFound
}

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
