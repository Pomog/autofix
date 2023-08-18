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
