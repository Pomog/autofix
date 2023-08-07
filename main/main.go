package main

import (
	"fmt"

	"github.com/Pomog/autofix/functions"
)

func main() {
	test := "1E   (hex) files were added"
	output1 := functions.ReplaceHexWithDecimal(test)
	fmt.Println(output1)
}
