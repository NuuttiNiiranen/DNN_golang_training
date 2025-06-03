package utils

import (
	"fmt"
	"os"
)

// LoadFileIfExists checks if a file exists, and returns its content as []byte if it does
func LoadFileIfExists(filename string) ([]byte, error) {
	// Check if file exists
	// Underscore means that we ignore the first return value coming from os.stat.
	// Having os.stat before ; means that we execute that part of the code before if statement
	// so here we first check what errors we get from os.stat and then use if clause depending on if we got IsNotExist error
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s does not exist - message from fileloader.go", filename)
	}

	// Read file into []byte
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w - message from fileloader.go", filename, err)
	}

	fmt.Print("Data loaded succesfully! - message from fileloader.go")

	return data, nil
}
