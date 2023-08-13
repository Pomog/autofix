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
	str = "This is so exciting (up,   2). I can't wait to see  what happens next (up, 3). But I'm sure it will be (up, 3) exciting."
	fmt.Println(functions.ToUppercaseWithNumber(str))
}
