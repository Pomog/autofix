package main

import (
	"fmt"
	"os"

	"github.com/Pomog/autofix/functions"
)

func main() {
	checkArgs(len(os.Args))
	greetengs()

	initialFileName, resultFileName := os.Args[1], os.Args[2]

	initialStrings, errRead := functions.ReadFromFile(initialFileName)
	if errRead != nil {
		fmt.Println(errRead)
		os.Exit(1)
	}

	var resultStrings []string
	for _, str := range initialStrings {
		fixedStr := functions.ApplyAutoFixingFunctions(str)
		resultStrings = append(resultStrings, fixedStr)
	}

	if functions.WriteToFile(resultFileName, resultStrings) != nil {
		fmt.Println(errRead)
		os.Exit(1)
	}

	farewell(resultFileName)
}

func checkArgs(argsCount int) {
	if argsCount != 3 {
		fmt.Println("Wrong number of arguments!")
		fmt.Println("Usage: go run . <input file> <output file>")
		os.Exit(1)
	}
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
