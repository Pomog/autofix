package functions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
replaces hexadecimal values followed by "(hex)"
in the input string with their corresponding decimal equivalents.
*/
func ReplaceHexWithDecimal(input string) string {
	re := regexp.MustCompile(`\b([0-9A-Fa-f]+)\s*\(hex\)`)
	return replace(re, input, parseStrHexToDec)
}

/*
replaces binary values followed by "(bin)"
in the input string with their corresponding decimal equivalents.
*/
func ReplaceBinWithDecimal(input string) string {
	re := regexp.MustCompile(`\b([01]+)\s*\(bin\)`)
	return replace(re, input, parseStrBinToDec)
}

/*
converts the word followed by "(up)"
with the Uppercase version
*/
func ToUppercase(input string) string {
	re := regexp.MustCompile(`\b([a-zA-Z0-9]+)\s*\(up\)`)
	return replace(re, input, convertToUpper)
}

/*
converts the word followed by "(up)"
with the Uppercase version
*/
func ToLowercase(input string) string {
	re := regexp.MustCompile(`\b([a-zA-Z0-9]+)\s*\(low\)`)
	return replace(re, input, convertToLower)
}

/*
converts the word followed by "(cap)"
with the capitalized version of it
*/
func Capitalization(input string) string {
	re := regexp.MustCompile(`\b([a-zA-Z0-9]+)\s*\(cap\)`)
	return replace(re, input, convertToCap)
}

/*
search and replaces matched patterns in the input string
using the provided conversion function.
*/
func replace(re *regexp.Regexp, input string, conversionFunc func(foundStr string, match string) string) string {
	output := re.ReplaceAllStringFunc(input, func(match string) string {
		foundStr := re.FindStringSubmatch(match)[1]
		return conversionFunc(foundStr, match)
	})
	return output
}

/*
parseStrHexToDec converts a hexadecimal string to its decimal equivalent.
If the conversion fails, it returns the original match.
*/
func parseStrHexToDec(hexStr string, match string) string {
	return parseStrWithBaseToDec(hexStr, match, 16, 64)
}

/*
parseStrBinToDec converts a binary string to its decimal equivalent.
If the conversion fails, it returns the original match.
*/
func parseStrBinToDec(binStr string, match string) string {
	return parseStrWithBaseToDec(binStr, match, 2, 64)
}

func parseStrWithBaseToDec(binStr, match string, base, bitsize int) string {
	decimalValue, err := strconv.ParseInt(binStr, base, bitsize)
	if err != nil {
		return match
	}
	return fmt.Sprintf("%d", decimalValue)
}

/*
converts a string to its uppercase equivalent
function could potentially be combined with the convertToLower and convertToCap functions since it's just capitalizing the first letter.
*/
func convertToUpper(str string, match string) string {
	return strings.ToUpper(str)
}

/*
converts a string to its lowercase equivalent
function could potentially be combined with the convertToUpper and convertToCap functions since it's just capitalizing the first letter.
*/
func convertToLower(str string, match string) string {
	return strings.ToLower(str)
}

/*
converts a string to its capitalized equivalent
function could potentially be combined with the convertToUpper and convertToLower functions since it's just capitalizing the first letter.
*/
func convertToCap(str string, match string) string {
	return strings.Title(str)
}
