package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("problem opening provided file")
	}

	// Ensure the file is closed when the function exits
	defer file.Close()

	// Create a new Scanner to read the file
	scanner := bufio.NewScanner(file)

	var lines []string
	// Read the file line by line
	for scanner.Scan() {
		// Append each line to the lines slice
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		file.Close()
		return nil, err
	}

	return lines, nil
}
