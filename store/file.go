package store

import (
	"os"
)

// WriteFile writes a file containing the data into a file in the specified directory.
func WriteFile(path, file string, data []byte) error {
	err := os.MkdirAll(path, 0o755)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, data, 0o644)
	if err != nil {
		return err
	}
	return nil
}

// ReadFile reads a file and returns the data
func ReadFile(file string) ([]byte, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}
