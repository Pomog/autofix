package main

import (
	"fmt"

	"github.com/Pomog/autofix/functions"
)

func main() {
	fmt.Println("testCapitalizationWithNumber")
	str := "This is so exciting (up, -1). I can't wait to see  what happens next (up, 3). But I'm sure it will be (up, 3) exciting."
	fmt.Println(functions.CapitalizationWithNumber(str))
}
