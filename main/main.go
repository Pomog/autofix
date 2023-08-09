package main

import (
	"fmt"

	"github.com/Pomog/autofix/functions"
)

func main() {
	fmt.Println("testCapitalizationWithNumber")
	str := "This is so exciting (up, 2)"
	functions.CapitalizationWithNumber(str)
}
