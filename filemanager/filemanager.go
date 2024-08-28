package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputPath string, outputPath string) (FileManager, error) {

	if inputPath == "" || outputPath == "" {
		return FileManager{}, errors.New("an input path and output path are required")
	}

	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}, nil
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("problem opening provided file")
	}

	// Ensure the file is closed when the function exits
	defer file.Close()

	// Create a new Scanner to read the file
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// 'any' can also be written as interface{}
func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("unable to create file")
	}

	// NewEncoder expects an io.Writer type input.
	// The file output by os.Create satisfies this.
	encoder := json.NewEncoder(file)
	defer file.Close()

	err = encoder.Encode(data)
	if err != nil {
		return errors.New("unable to encode data to JSON")
	}

	return nil
}
