package functions

import (
	"fmt"
	"regexp"
	"strconv"
)

var RegisteredFunctions = []struct {
	Name string
	Fn   func()
}{
	{"Funct1", Funct1},
	{"Funct2", WrapReplaceHexWithDecimal},
}

func Funct1() {
	fmt.Println("Funct1 is called")
}

func WrapReplaceHexWithDecimal() {
	input := "1E (hex) files were added. 2A (hex) items are available."
	output := ReplaceHexWithDecimal(input)
	fmt.Println(output)
}

func ReplaceHexWithDecimal(input string) string {
	re := regexp.MustCompile(`([0-9A-Fa-f]+) \(hex\)`)
	output := re.ReplaceAllStringFunc(input, func(match string) string {
		hexStr := re.FindStringSubmatch(match)[1]
		decimalValue, err := strconv.ParseInt(hexStr, 16, 64)
		if err != nil {
			return match
		}
		return fmt.Sprintf("%d", decimalValue)
	})
	return output
}
