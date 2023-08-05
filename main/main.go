package main

import (
	"fmt"

	"github.com/Pomog/autofix/functions"
)

var functionMap = make(map[string]func())

func addToFunctionMap(fnName string, fn func()) {
	functionMap[fnName] = fn
}

func main() {
	// Add registered functions to the functionMap
	for _, fn := range functions.RegisteredFunctions {
		addToFunctionMap(fn.Name, fn.Fn)
	}

	condition := "Funct1"

	if chosenFunction, exists := functionMap[condition]; exists {
		chosenFunction()
	} else {
		fmt.Println("No function found for the given condition")
	}
}
