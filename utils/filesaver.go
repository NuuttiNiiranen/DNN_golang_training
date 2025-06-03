package utils

import (
	"dnn_golang_training/network"
	"os"
)

// Function to save neural network model to file
func SaveNetworkToFile(filename string, neuralNetwork network.NeuralNetwork) error {
	// First we transform network into JSON data for easy transport
	jsonText := neuralNetwork.Jsonify()
	// JSON is written into a file :)
	err := os.WriteFile(filename, jsonText, 0644)
	if err != nil {
		return err
	}
	return nil
}
