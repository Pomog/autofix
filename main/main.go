package main

import (
	"fmt"

	"github.com/Pomog/autofix/functions"
)

func main() {
	fmt.Println("testCapitalizationWithNumber")
	str := "This is so exciting (cap,   2). I can't wait to see  what happens next (cap, 3). But I'm sure it will be (cap, 3) exciting."
	fmt.Println(functions.CapitalizationWithNumber(str))
}
