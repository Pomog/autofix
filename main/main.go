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
}
