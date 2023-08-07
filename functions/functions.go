package functions

import (
	"fmt"
	"regexp"
	"strconv"
)

func ReplaceHexWithDecimal(input string) string {
	re := regexp.MustCompile(`\b([0-9A-Fa-f]+)\s*\(hex\)`)
	output := re.ReplaceAllStringFunc(input, func(match string) string {
		hexStr := re.FindStringSubmatch(match)[1]
		decimalValue, err := strconv.ParseInt(hexStr, 16, 64)
		if err != nil {
			// If conversion fails, return the original match
			return match
		}
		return fmt.Sprintf("%d", decimalValue)
	})
	return output
}

func ReplaceBinWithDecimal(input string) string {
	re := regexp.MustCompile(`\b([01]+)\s*\(bin\)`)
	output := re.ReplaceAllStringFunc(input, func(match string) string {
		binStr := re.FindStringSubmatch(match)[1]
		decimalValue, err := strconv.ParseInt(binStr, 2, 64)
		if err != nil {
			return match
		}
		return fmt.Sprintf("%d", decimalValue)
	})
	return output
}
