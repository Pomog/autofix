package main

import (
	"fmt"
	"os"

	"github.com/Pomog/autofix/functions"
)

func main() {
	greetengs()

	initialFileName, resultFileName := os.Args[1], os.Args[2]

	initialStrings, errRead := functions.ReadFromFile(initialFileName)
	if errRead != nil {
		fmt.Println(errRead)
		os.Exit(1)
	}

	var resultStrings []string
	for _, str := range initialStrings {
		for _, fixingFunc := range functions.GetAutoFixingFunctions() {
			str = fixingFunc(str)
		}
		resultStrings = append(resultStrings, str)
	}

	if functions.WriteToFile(resultFileName, resultStrings) != nil {
		fmt.Println(errRead)
		os.Exit(1)
	}

	farewell(resultFileName)
}

func greetengs() {
	fmt.Println("Welcome to the AutoFix tool!")
	fmt.Printf("The file --> %s <-- will be fixed\n", os.Args[1])
}

func farewell(resultFileName string) {
	fmt.Println("Finished. No errors. Thanks for using.")
	fmt.Printf("The result is in the file --> %s <--\n", resultFileName)
	fmt.Println("Goodbye!")
}
