package functions

import (
	"fmt"
	"regexp"
	"strings"
)

/*
converts the word followed by "(cap)"
with the capitalized version of it
if a number appears next to it, like so: (cap, <number>)
it turns the previously specified number of words in capitalized form
*/
func CapitalizationWithNumber(input string) string {
	re := regexp.MustCompile(`\b[\w]+\s*\(up(?:,\s*(\d+))?\)`)
	for noMatchCheck(re, input) {
		parsedNumber := parseInt(re.FindStringSubmatch(input)[1])
		repeatedPattern := strings.Repeat(`(\b[\w]+\s*)`, parsedNumber) + `\(up(?:,\s*(\d+))?\)`
		reRep := regexp.MustCompile(repeatedPattern)

		fixedString := getCapitalizedString(reRep, input, parsedNumber)

		input = replaceStrings(reRep, input, parsedNumber, fixedString)
	}
	return input
}

func replaceStrings(reRep *regexp.Regexp, input string, parsedNumber int, fixedString string) string {
	machedString := reRep.FindStringSubmatch(input)[0]
	if parsedNumber != 0 {
		input = strings.Replace(input, machedString, strings.TrimSpace(fixedString), 1)
	} else {
		input = strings.TrimSpace(strings.Replace(input, machedString, "", 1))
	}
	return input
}

func getCapitalizedString(reRep *regexp.Regexp, input string, parsedNumber int) string {
	var fixedString string
	for _, result := range reRep.FindStringSubmatch(input)[1 : parsedNumber+1] {
		fixedString = fixedString + makeCapitallization(result)
	}
	return fixedString
}

/*
converts a string to its capitalized equivalent
*/
func makeCapitallization(input string) string {
	return strings.Title(input)
}

/*
No match found, return input unchanged
*/
func noMatchCheck(re *regexp.Regexp, input string) bool {
	var answer bool
	if re.FindStringSubmatch(input) != nil {
		answer = true
	}
	return answer
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
