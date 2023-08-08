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
	re := regexp.MustCompile(`\b([a-z]+)\s*\(up\)`)
	return replace(re, input, convertToUpper)
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
parseStrBinToDec converts a binary string to its decimal equivalent.
If the conversion fails, it returns the original match.
*/
func parseStrBinToDec(binStr string, match string) string {
	decimalValue, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		return match
	}
	return fmt.Sprintf("%d", decimalValue)
}

/*
parseStrHexToDec converts a hexadecimal string to its decimal equivalent.
If the conversion fails, it returns the original match.
*/
func parseStrHexToDec(hexStr string, match string) string {
	decimalValue, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return match
	}
	return fmt.Sprintf("%d", decimalValue)
}

/*
converts a string to its uppercase equivalent
*/
func convertToUpper(str string, match string) string {
	return strings.ToUpper(str)
}
