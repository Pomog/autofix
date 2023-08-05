package functions

import (
	"fmt"
)

var RegisteredFunctions = []struct {
	Name string
	Fn   func()
}{
	{"Funct1", Funct1},
}

func Funct1() {
	fmt.Println("Funct1 is called")
}
