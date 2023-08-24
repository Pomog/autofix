package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Pomog/autofix/functions"
)

func main() {
	checkArgs(len(os.Args)) // there should be 2 arguments added: the input file name and the output file name
	greetings()

	initialFileName, resultFileName := os.Args[1], os.Args[2]

	initialStrings, errRead := functions.ReadFromFile(initialFileName)
	if errRead != nil {
		log.Fatal(errRead)
	}

	var resultStrings []string
	for _, str := range initialStrings {
		fixedStr := functions.ApplyAutoFixingFunctions(str) // apply all functions form the map which returns by getAutoFixingFunctions()
		resultStrings = append(resultStrings, fixedStr)
	}

	if witeToFileEWrror := functions.WriteToFile(resultFileName, resultStrings); witeToFileEWrror != nil {
		log.Fatal(witeToFileEWrror)
	}

	farewell(resultFileName)
}

func checkArgs(argsCount int) {
	if argsCount != 3 {
		log.Fatal("Wrong number of arguments!\nUsage: go run . <input file> <output file>")
	}
}

func greetings() {
	message := fmt.Sprintf("Welcome to the AutoFix tool!\nThe file --> %s <-- will be fixed\n", os.Args[1])
	fmt.Println(message)
}

func farewell(resultFileName string) {
	message := fmt.Sprintf("Finished. No errors. Thanks for using.\nThe result is in the file --> %s <--\nGoodbye!\n", resultFileName)
	fmt.Println(message)
}
