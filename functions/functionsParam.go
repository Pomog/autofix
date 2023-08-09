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
	shouldReturn, returnValue := noMatchCheck(re, input)
	if shouldReturn {
		return returnValue
	}

	parsedNumber := parseInt(re.FindStringSubmatch(input)[1])

	repeatedPattern := strings.Repeat(`(\b[\w]+\s*)`, parsedNumber) + `\(up(?:,\s*(\d+))?\)`
	reRep := regexp.MustCompile(repeatedPattern)

	var fixedString string
	for i, result := range reRep.FindAllStringSubmatch(input, -1)[0][1 : parsedNumber+1] {
		fixedString = fixedString + convertToCap(result, result)
		fmt.Printf("index: %d, result: %s\n", i, result)
		fmt.Printf("fixedString: %s\n", fixedString)
	}

	return replaceParam(reRep, input, convertToCap)
}

/*
No match found, return input unchanged
*/
func noMatchCheck(re *regexp.Regexp, input string) (bool, string) {
	match := re.FindStringSubmatch(input)
	if match == nil {
		return true, input
	}
	return false, ""
}

/*
search and replaces matched patterns in the input string
using the provided conversion function.
*/
func replaceParam(reRep *regexp.Regexp, input string, conversionFunc func(foundStr string, match string) string) string {
	output := reRep.ReplaceAllStringFunc(input, func(match string) string {
		foundStr := reRep.FindStringSubmatch(match)[1]
		return conversionFunc(foundStr, match)
	})
	return output
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

// if len(reRep.FindAllStringSubmatch(input, -1)) != 0 {
// 	for i, word := range reRep.FindAllStringSubmatch(input, -1)[0][1:] {
// 		fmt.Printf("index: %d, word: %s\n", i, word)
// 	}
// } else {
// 	fmt.Println("no match")
// }
// result := replaceParam(reRep, input, convertToCap)
// fmt.Println("result")
// fmt.Println(result)
// return result
