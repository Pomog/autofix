package main

import (
	"fmt"

	"github.com/Pomog/autofix/functions"
)

func main() {
	fmt.Println("testCapitalizationWithNumber")
	str := "This is so exciting (cap,   2). I can't wait to see  what happens next (cap, 3). But I'm sure it will be (cap, 3) exciting."
	fmt.Println(functions.CapitalizationWithNumber(str))

	fmt.Println("testToUppercaseWithNumber")
	str = "(up, 0) This is so exciting (up,   2). I can't wait to see  what happens next (up, 3). But I'm sure it will be (up, 3) exciting."
	fmt.Println(functions.ToUppercaseWithNumber(str))

	fmt.Println("testToloWercaseWithNumber")
	str = "(low, 1) This is so Exciting (low,   0)."
	fmt.Println(functions.ToLowercaseWithNumber(str))

	fmt.Println("CorrectSpaces")
	str = "This is so exciting ...  I can't wait to see  what happens next.But I'm sure it will be exciting !"
	output := functions.CorrectPunctuationsSpaces(str)
	fmt.Println(output)

}
