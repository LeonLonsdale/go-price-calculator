package filemanager

import (
	"bufio"
	"encoding/json"
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

// 'any' can also be written as interface{}
func WriteJSON(path string, data any) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("unable to create file")
	}

	// NewEncoder expects an io.Writer type input.
	// The file output by os.Create satisfies this.
	encoder := json.NewEncoder(file)
	defer file.Close()

	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("unable to encode data to JSON")
	}

	return nil
}
