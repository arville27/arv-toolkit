package utils

import (
	"io"
	"os"
)

func WriteToFile(filePath string, data []byte) error {
	// Open the file for writing. If the file doesn't exist, it will be created.
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write data to the file
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(filePath string) ([]byte, error) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file's content
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}
