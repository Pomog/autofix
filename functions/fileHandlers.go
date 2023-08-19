package functions

import (
	"bufio"
	"os"
)

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

	return fileLines, nil
}

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

	return nil
}
