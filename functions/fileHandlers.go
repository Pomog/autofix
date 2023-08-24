package functions

import (
	"bufio"
	"os"
)

/*
ReadFromFile reads lines from a file and returns them as a slice of strings.
*/
func ReadFromFile(fileName string) ([]string, error) {
	file, errRead := os.Open(fileName)
	if errRead != nil {
		return nil, errRead
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	if errScan := fileScanner.Err(); errScan != nil {
		return nil, errScan
	}
	file.Close() //deferred closure may not be executed if an error occurs before the defer statement
	return fileLines, nil
}

/*
WriteToFile writes lines to a file
*/
func WriteToFile(fileName string, lines []string) error { // TODO: unit tests
	file, errWrite := os.Create(fileName)
	if errWrite != nil {
		return errWrite
	}
	defer file.Close()

	fileWriter := bufio.NewWriter(file)
	for _, line := range lines {
		fileWriter.WriteString(line + "\n")
	}

	// Flush any buffered data to the file
	if errFlush := fileWriter.Flush(); errFlush != nil {
		return errFlush
	}
	file.Close() //deferred closure may not be executed if an error occurs before the defer statement
	return nil
}
