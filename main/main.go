package main

import (
	"fmt"
	"strings"

	"github.com/Pomog/autofix/functions"
)

func main() {
	test := "1E   (hex) files were added"
	output1 := functions.ReplaceHexWithDecimal(test)
	fmt.Println(output1)
	fmt.Println(strings.ToUpper("123@23"))

	fmt.Println("testCapitalizationWithNumber")
	str := "This is so exciting (up, 3)"
	functions.CapitalizationWithNumber(str)

}
