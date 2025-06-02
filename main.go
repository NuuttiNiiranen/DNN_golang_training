// This class is entry point for the app. It has training and evaluation.
package main

import (
	"fmt"
	"os"
)

func main() {
	// Creating a simple neural network model
	// neuralNetwork := models.FirstModel()
	// Running the first model with input of 64 1s
	// neuralNetwork.Predict([]float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})

	// Adjusting single weights and biases of neural network
	// neuralNetwork.AdjustWeights(2, 1, 1, 2.5, 2)
	// neuralNetwork.AdjustWeights(2, 1, 2, 2.5, 2)

	// Next I need to save the created model so I can actually run the network
	// jsonText := neuralNetwork.Jsonify()
	// err := SaveNetworkToFile("FirstModel", jsonText)
	// if err != nil {
	//	fmt.Print("Something went wrong :(")
	// } else {
	// 	fmt.Print("Saving succesful!")
	// }
	byteText, err := LoadFileIfExists("models/FirstModel")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(byteText))
}
func SaveNetworkToFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// LoadFileIfExists checks if a file exists, and returns its content as []byte if it does
func LoadFileIfExists(filename string) ([]byte, error) {
	// Check if file exists
	// Underscore means that we ignore the first return value coming from os.stat.
	// Having os.stat before ; means that we execute that part of the code before if statement
	// so here we first check what errors we get from os.stat and then use if clause depending on if we got IsNotExist error
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s does not exist", filename)
	}

	// Read file into []byte
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	fmt.Print("Data loaded succesfully!")

	return data, nil
}
